package main

import (
	"fmt"
)

func main () {
	fmt.Println("Challenge 1 - Convert hex to base64")
	
	hex := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	fmt.Printf("hex : %s\n", hex)

	// decode hex to string
	decoded, err := decodeHex(hex)
	if err != nil {
		fmt.Printf("Failed to decode hex %s\n",err)
	}
	fmt.Printf("hex decoded : %s\n",decoded)

	// encode to base64
	encoded := base64Encode(decoded)
	fmt.Printf("base64 encoded : %s\n", encoded)

	fmt.Print("Challenge 1 complete\n\n")


}