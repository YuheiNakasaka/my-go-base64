package main

import (
	"encoding/base64"
	"fmt"

	"github.com/YuheiNakasaka/my-go-base64/base64"
)

// Example
func main() {
	s := "ABCDEFG"
	fmt.Println("Src message:     ", s)

	// original base64
	enc1 := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println("Original Encode: ", enc1)
	dec1, _ := base64.StdEncoding.DecodeString(string(enc1))
	fmt.Println("Original Decode: ", string(dec1))

	// custom base64
	enc2 := b64.Encode(s)
	dec2 := b64.Decode(enc2)
	fmt.Println("Custom Encode:   ", enc2)
	fmt.Println("Custom Decode:   ", dec2)
}
