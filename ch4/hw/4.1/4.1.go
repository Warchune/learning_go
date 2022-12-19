package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	var n int
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	for i := 0; i < len(c1); i++ {
		s := c1[i] ^ c2[i]
		fmt.Printf("%9b = %9b ^ %9b\n", s, c1[i], c2[i])
		nc := 0
		for s != 0 {
			fmt.Printf("%9b & %9b = ", s, s-1)
			s = s & (s - 1)
			fmt.Printf("%9b //", s)
			nc++
		}
		n += nc
		fmt.Println()
	}

	fmt.Printf("Different bit count is %d\n", n)
}
