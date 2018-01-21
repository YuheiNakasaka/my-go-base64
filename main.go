package main

import (
	"fmt"

	"github.com/YuheiNakasaka/my-go-base64/base64"
)

func main() {
	s := "ABCDEFGHIJKLMN"
	e := b64.Encode(s)
	fmt.Println(e)
}
