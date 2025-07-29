package main

import (
	"fmt"
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
}