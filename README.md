# my-go-base64

My base64 implementation without a bit shift in Go.

Th performance is slower than the original encoding/base64 package.

# Example

```
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

// Original:  QUJDREVGRw==
// Custom:    QUJDREVGRw==
```
