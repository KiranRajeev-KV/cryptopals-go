package main

import (
	"encoding/base64"
	"encoding/hex"
)

func decodeHex (input []byte) ([]byte, error) {
	// hex.DecodedLen returns the exact length needed to hold the decoded bytes.
	decoded := make([]byte,hex.DecodedLen(len(input)))
	_, err := hex.Decode(decoded,input)
	if err != nil {
		return nil,err
	}
	return decoded,nil
}

func base64Encode(input []byte) ([]byte) {
	encoded := make([]byte, base64.StdEncoding.EncodedLen(len(input)))
	base64.StdEncoding.Encode(encoded, input)
	return encoded
}