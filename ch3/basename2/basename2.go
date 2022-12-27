package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(basename2("a/b/c.go"))
}

func basename2(s string) string {
	if slash := strings.LastIndex(s, "/"); slash != -1 {
		s = s[slash+1:]
	}
	if dot := strings.LastIndex(s, "."); dot > 0 {
		s = s[:dot]
	}
	return s
}
