package b64

import (
	"fmt"
	"strconv"
	"strings"
)

const encodeStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func stringToBin(s string) string {
	binStr := ""
	for _, c := range s {
		b := fmt.Sprintf("%.8b", c)
		binStr = fmt.Sprintf("%s%s", binStr, b)
	}
	return binStr
}

func split6bits(s string) []string {
	l := (len(s) / 6) + 1
	ss := make([]string, l)
	i := 1
	str := ""
	for _, c := range s {
		str = str + string(c)
		if i%6 == 0 {
			ss[(i/6)-1] = str
			str = ""
		}
		i++
	}

	if str != "" {
		pad := 6 - len(str)
		pads := strings.Repeat("0", pad)
		ss[l-1] = str + pads
	}

	return ss
}

func convert(ss []string) string {
	s := ""
	for i := 0; i < len(ss); i++ {
		i2, _ := strconv.ParseInt(ss[i], 2, 0)
		s = fmt.Sprintf("%s%s", s, string(encodeStr[i2]))
	}

	l := (4 - len(s)%4) % 4
	pad := strings.Repeat("=", l)

	return s + pad
}

// Encode message with [A-Z][a-z][1-9]+/
func Encode(s string) string {
	bins := stringToBin(s)
	bin6s := split6bits(bins)
	resp := convert(bin6s)
	return resp
}
