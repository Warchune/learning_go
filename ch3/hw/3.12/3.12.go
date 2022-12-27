package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("use: 3.12 string1 string2")
		return
	}

	fmt.Println(isAnagramma(os.Args[1], os.Args[2]))
}

func isAnagramma(s1, s2 string) int {
	/*if len(s1) != len(s2) {
		return -1
	}
	n := len(s1)
	for i := 0; i < n; i++ {
		if s1[i] != s2[n-1-i] {
			return -1
		}
	}*/
	aFreq := make(map[rune]int)
	for _, c := range s1 {
		aFreq[c]++
	}
	bFreq := make(map[rune]int)
	for _, c := range s2 {
		bFreq[c]++
	}
	for k, v := range aFreq {
		if bFreq[k] != v {
			return -1
		}
	}
	for k, v := range bFreq {
		if aFreq[k] != v {
			return -1
		}
	}
	return 1
}
