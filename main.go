package main

import (
	"encoding/base64"
	"fmt"

	"github.com/YuheiNakasaka/my-go-base64/base64"
)

func main() {
	s := "ABCDEFG"

	// original
	enc1 := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println("Original: ", enc1)

	// custom
	enc2 := b64.Encode(s)
	fmt.Println("Custom:   ", enc2)
}
