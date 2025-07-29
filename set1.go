package main

import (
	"fmt"
	"os"
	"strings"
)

func main () {
	fmt.Println("Challenge 1 - Convert hex to base64")
	
	hex := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	fmt.Printf("hex : %s\n", hex)

	// decode hex to string
	decoded := decodeHex(hex)

	fmt.Printf("hex decoded : %s\n",decoded)

	// encode to base64
	encoded := base64Encode(decoded)
	fmt.Printf("base64 encoded : %s\n", encoded)

	fmt.Println("\nChallenge 2 - Fixed XOR")

	op1:= []byte("1c0111001f010100061a024b53535009181c")
	op2:= []byte("686974207468652062756c6c277320657965")

	decodedOp1 := decodeHex(op1)
	decodedOp2 := decodeHex(op2)
	
	// fmt.Println(len(decodedOp1)==len(decodedOp2))
	fmt.Printf("op1 decoded : %x\nop2 decoded : %x\n", decodedOp1,decodedOp2)

	xorRes := fixedXOR(decodedOp1, decodedOp2)
	xorResEncoded:= encodeHex(xorRes)
	fmt.Printf("XOR result : %s\n",xorResEncoded)

	fmt.Println("\nChallenge 3 - Single-byte XOR cipher")

	cipherHex := []byte("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	cipherBytes := decodeHex(cipherHex)
	fmt.Printf("hex decoded cipher text : %d\n",cipherBytes)

	_, bestDecoded, _ := findBestSingleByteXOR(cipherBytes)
	fmt.Printf("decoded text: %s\n", bestDecoded)

	fmt.Println("\nChallenge 4 - Detect single-character XOR")

	file, err := os.ReadFile("./data/set1-challenge4.txt")
	if err != nil {
		panic(fmt.Sprintf("Failed to read file: %s\n", err))
	}
	lines:= strings.Split(string(file),"\n")
	
	bestScore := 0.0
	bestDecoded = []byte{}
	for _, line := range lines {
		if len(line) > 0 {
			cipherBytes := decodeHex([]byte(line))
			_,decoded,score := findBestSingleByteXOR(cipherBytes)
			
			if score > bestScore {
				bestScore = score
				bestDecoded = decoded
			}
		}
	}

	fmt.Printf("Best decoded text: %s\n", bestDecoded)

	fmt.Println("\nChallenge 5 - Implement repeating-key XOR")

	plaintext := []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")
	key := []byte("ICE")
	fmt.Printf("Plaintext: %s\nKey: %s\n\n", plaintext, key)

	ciphertext := repeatingKeyXOR(plaintext, key)
	fmt.Printf("Ciphertext: %x\n", ciphertext)
}