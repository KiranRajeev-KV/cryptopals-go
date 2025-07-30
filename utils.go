package main

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"maps"
	"sort"
	"unicode"
)

func decodeHex(input []byte) []byte {
	// hex.DecodedLen returns the exact length needed to hold the decoded bytes.
	decoded := make([]byte, hex.DecodedLen(len(input)))
	_, err := hex.Decode(decoded, input)
	if err != nil {
		panic(fmt.Sprintf("Failed to decode hex %s\n", err))
	}
	return decoded
}

func encodeHex(input []byte) []byte {
	encoded := make([]byte, hex.EncodedLen(len(input)))
	hex.Encode(encoded, input)
	return encoded
}

func base64Decode(input []byte) []byte {
	decoded := make([]byte, base64.StdEncoding.DecodedLen(len(input)))
	_, err := base64.StdEncoding.Decode(decoded, input)
	if err != nil {
		panic(fmt.Sprintf("Failed to decode base64 %s\n", err))
	}
	return decoded
}

func base64Encode(input []byte) []byte {
	encoded := make([]byte, base64.StdEncoding.EncodedLen(len(input)))
	base64.StdEncoding.Encode(encoded, input)
	return encoded
}

func fixedXOR(op1 []byte, op2 []byte) []byte {
	if len(op1) != len(op2) {
		panic(fmt.Sprintf("Operands must be of equal length (op1: %d bytes, op2: %d bytes)", len(op1), len(op2)))
	}
	res := make([]byte, len(op1))
	for i := range op1 {
		res[i] = op1[i] ^ op2[i]
	}
	return res
}

func getCharWeight(char byte) float64 {
	// got the freq from wikipedia ;)
	freq := map[byte]float64{
		' ': 0.1874,
		'a': 0.08167, 'b': 0.01492, 'c': 0.02782, 'd': 0.04253,
		'e': 0.12702, 'f': 0.02228, 'g': 0.02015, 'h': 0.06094,
		'i': 0.06966, 'j': 0.00153, 'k': 0.00772, 'l': 0.04025,
		'm': 0.02406, 'n': 0.06749, 'o': 0.07507, 'p': 0.01929,
		'q': 0.00095, 'r': 0.05987, 's': 0.06327, 't': 0.09056,
		'u': 0.02758, 'v': 0.00978, 'w': 0.02360, 'x': 0.00150,
		'y': 0.01974, 'z': 0.00074,
	}

	// Normalize char to lowercase if it's a letter
	lower := byte(unicode.ToLower(rune(char)))
	return freq[lower]
}

func scoreText(text []byte) float64 {
	score := 0.0
	// range text returns 2 things : index and char
	for _, char := range text {
		score += getCharWeight(char)
	}
	return score
}

func singleByteXOR(cipher []byte, key byte) []byte {
	res := make([]byte, len(cipher))
	for i := range cipher {
		res[i] = cipher[i] ^ key
	}
	return res
}

func findBestSingleByteXOR(cipher []byte) (byte, []byte, float64) {
	bestKey := byte(0)
	bestScore := 0.0
	bestDecoded := []byte{}

	for key := range 256 {
		decoded := singleByteXOR(cipher, byte(key))
		score := scoreText(decoded)

		if score > bestScore {
			bestScore = score
			bestKey = byte(key)
			bestDecoded = decoded
		}
	}

	return bestKey, bestDecoded, bestScore
}

func repeatingKeyXOR(plaintext []byte, key []byte) []byte {
	res := make([]byte, len(plaintext))
	keyLen := len(key)

	for i := range plaintext {
		// 0 % 3 = 0, 1 % 3 = 1, 2 % 3 = 2, 3 % 3 = 0, 4 % 3 = 1, ...
		res[i] = plaintext[i] ^ key[i%keyLen]
	}

	return res
}

func getHammingWeights() map[byte]int {
	weights := map[byte]int{0: 0}
	pow2 := byte(1)

	for range 8 {
		copyMap := make(map[byte]int)
		maps.Copy(copyMap, weights)
		for k, v := range copyMap {
			weights[k+pow2] = v + 1
		}
		pow2 <<= 1
	}
	return weights
}

func getHammingDistance(str1 []byte, str2 []byte) int {
	if len(str1) != len(str2) {
		panic(fmt.Sprintf("Strings must be of equal length (str1: %d bytes, str2: %d bytes)", len(str1), len(str2)))
	}
	res := fixedXOR(str1, str2)

	weights := getHammingWeights()
	distance := 0
	for _, b := range res {
		distance += weights[b]
	}
	return distance
}

type keySizeStrength struct {
	size               int
	normalizedDistance float64
}

func getKeySizes(ciphertext []byte) []keySizeStrength {
	allKeyLengths := make([]keySizeStrength, 0)

	for keySize := 2; keySize < 40; keySize++ {
		var totalDistance float64
		for i := range 4 {
			start := keySize * i
			end := start + keySize
			nextEnd := end + keySize
			if nextEnd > len(ciphertext) {
				break
			}
			totalDistance += float64(getHammingDistance(ciphertext[start:end], ciphertext[end:nextEnd]))
		}
		normalizedDistance := totalDistance / float64(keySize)
		allKeyLengths = append(allKeyLengths, keySizeStrength{size: keySize, normalizedDistance: normalizedDistance})
	}
	sort.Slice(allKeyLengths, func(i, j int) bool {
		return allKeyLengths[i].normalizedDistance < allKeyLengths[j].normalizedDistance
	})
	return allKeyLengths
}

func breakRepeatingKeyXOR(ciphertext []byte) ([]byte, []byte) {
	keyStrengths := getKeySizes(ciphertext)
	var bestPlaintext []byte
	var bestKey []byte
	var bestScore float64

	for i := range 5 {
		keySize := keyStrengths[i].size

		// Split ciphertext into keySize chunks, each chunk contains bytes encrypted with the same key byte
		blocks := make([][]byte, keySize)
		for j, b := range ciphertext {
			blocks[j%keySize] = append(blocks[j%keySize], b)
		}

		// Find key byte for each block
		key := make([]byte, keySize)
		for k := range blocks {
			kb, _, _ := findBestSingleByteXOR(blocks[k])
			key[k] = kb
		}

		// Decrypt ciphertext with repeating key
		plaintext := repeatingKeyXOR(ciphertext, key)
		score := scoreText(plaintext)
		if score > bestScore || i == 0 {
			bestPlaintext = plaintext
			bestKey = key
			bestScore = score
		}
	}
	return bestPlaintext, bestKey
}

func removeAESPadding(input []byte) []byte {
	if len(input) == 0 {
		return input
	}

	paddingLength := int(input[len(input)-1])
	if paddingLength == 0 || paddingLength > len(input) {
		panic("Invalid padding length")
	}

	for _, b := range input[len(input)-paddingLength:] {
		if b != byte(paddingLength) {
			panic("Invalid padding byte")
		}
	}
	return input[:len(input)-paddingLength]
}

func decryptECB(ciphertext []byte, key []byte) []byte {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		panic(fmt.Sprintf("Failed to create AES cipher: %s\n", err))
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		panic("Ciphertext is not a multiple of the block size")
	}

	paddedPlaintext := make([]byte, len(ciphertext))
	for bs, be := 0, aes.BlockSize; bs < len(ciphertext); bs, be = bs+aes.BlockSize, be+aes.BlockSize {
		cipher.Decrypt(paddedPlaintext[bs:be], ciphertext[bs:be])
	}
	plaintext := removeAESPadding(paddedPlaintext)

	return plaintext
}

func detectECB(input []byte) int {
	lines := bytes.SplitSeq(input, []byte{'\n'})

	for line := range lines {
		if len(line) == 0 {
			continue
		}

		ciphertext := decodeHex([]byte(line))
		blockCount := make(map[string]int)
		
		for i := 0; i+aes.BlockSize <= len(ciphertext); i += aes.BlockSize {
			block := ciphertext[i : i+aes.BlockSize]
			blockHex := string(encodeHex(block))
			blockCount[blockHex]++
			if blockCount[blockHex] > 1 {
				return i
			}
		}
	}
	return -1 // No ECB detected
}