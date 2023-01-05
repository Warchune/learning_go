package main

import (
	"fmt"
)

func main() {
	s := []string{"one", "one", "two"}
	fmt.Println(unique(s))
}

func unique(strs []string) []string {
	w := 0
	for _, s := range strs {
		if strs[w] == s {
			continue
		}
		w++
		strs[w] = s
	}
	return strs[:w+1]
}
