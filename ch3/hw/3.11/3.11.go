package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("-1234567890.123456"))
}

func comma(s string) string {
	var buff bytes.Buffer
	n := len(s)
	wholeFirst, wholeLast := 0, n
	if s[0] == '-' || s[0] == '+' {
		buff.WriteByte(s[0])
		wholeFirst = 1
	}
	if i := strings.Index(s, "."); i != -1 {
		wholeLast = i
	}
	buff.WriteString(commaRec(s[wholeFirst:wholeLast]))
	if wholeLast != n {
		buff.WriteByte('.')
		buff.WriteString(commaRec(s[wholeLast+1:]))
	}
	return buff.String()
}

func commaRec(s string) string {
	var buf bytes.Buffer
	if n := len(s); n > 3 {
		first := n % 3
		if first == 0 {
			first = 3
		}
		for i, j := 0, first; j <= n; j += 3 {
			buf.WriteString(s[i:j])
			if j != n {
				buf.WriteByte(',')
			}
			i = j
		}
	}
	return buf.String()
}
