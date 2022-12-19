package main

import "fmt"

func main() {
	fmt.Println(2/3, 2%3, 3/2, 3%2)
	fmt.Println(gcd(5, 4))
}

func gcd(x, y int) int {
	for y != 0 {
		fmt.Println(x % y)
		x, y = y, x%y
	}
	return x
}
