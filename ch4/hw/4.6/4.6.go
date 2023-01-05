package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	s := []byte("one  two")
	fmt.Printf("%s\n", convertDbSpaceToOne(s))
}

func convertDbSpaceToOne(str []byte) []byte {
	w := 0
	countsp := 0
	for i := 0; i < len(str); i++ {
		r, size := utf8.DecodeRuneInString(string(str[i:]))
		fmt.Println(size)
		if unicode.IsSpace(r) {
			countsp++
			continue
		}
		if countsp > 0 {
			countsp = 0
			str[w] = byte(' ')
			w++
		}
		for j := i; j < i+size; j++ {
			str[w] = str[j]
			w++
		}
	}
	return str[:w-1]
}
