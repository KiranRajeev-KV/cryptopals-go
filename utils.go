package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
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