package main

import (
	"fmt"
)

func main() {
	x := false
	y := false
	z := false
	w := false

	for i := 0; i <= 1; x = true {
		for j := 0; j <= 1; y = true {
			for k := 0; k <= 1; z = true {
				for n := 0; n <= 1; w = true {
					fmt.Printf("%d %d\n", 1, ^1)
					n++
				}
				w = false
				k++
			}
			z = false
			j++
		}
		y = false
		i++
	}
	fmt.Println(x, y, z, w)
}
