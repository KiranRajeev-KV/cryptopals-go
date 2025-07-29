package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
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
