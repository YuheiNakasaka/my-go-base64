package b64

import (
	"fmt"
	"strconv"
	"strings"
)

const encodeStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func msgToBin(s string) string {
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

func binToMsg(ss []string) (s string) {
	for _, c := range ss {
		i2, _ := strconv.ParseInt(string(c), 2, 0)
		s = fmt.Sprintf("%s%c", s, i2)
	}
	return s
}

func concat6bitTo8bit(ss6 []string) (ss8 []string) {
	s := strings.Join(ss6, "")
	ss8 = make([]string, (len(s) / 8))
	i := 1
	str := ""
	for _, c := range s {
		str = str + string(c)
		if i%8 == 0 {
			ss8[(i/8)-1] = str
			str = ""
		}
		i++
	}
	return ss8
}

func revert(s string) (ss []string) {
	s = strings.Replace(s, "=", "", -1)
	ss = make([]string, len(s))
	for i, c := range s {
		idx := strings.Index(encodeStr, string(c))
		b := fmt.Sprintf("%.6b", idx)
		ss[i] = b
	}
	return ss
}

// Encode message with [A-Z][a-z][1-9]+/
func Encode(s string) (resp string) {
	bins := msgToBin(s)
	bin6s := split6bits(bins)
	resp = convert(bin6s)
	return resp
}

// Decode base64 message
func Decode(s string) (resp string) {
	bin6s := revert(s)
	bin8s := concat6bitTo8bit(bin6s)
	resp = binToMsg(bin8s)
	return resp
}
