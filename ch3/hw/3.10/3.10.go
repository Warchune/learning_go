package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("1234567890"))
	fmt.Println(commaRec("1234567890"))
	fmt.Println(commaBuffer("1234567890"))
}

func comma(s string) string {
	if n := len(s); n > 3 {
		for j := 3; j <= n; j += 3 {
			s = s[:n-j] + "," + s[n-j:]
		}
	}
	return s
}

func commaRec(s string) string {
	n := len(s)
	if n < 3 {
		return s
	}
	return commaRec(s[:n-3]) + "," + s[n-3:]
}

func commaBuffer(s string) string {
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
