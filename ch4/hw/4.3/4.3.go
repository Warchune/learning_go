package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3}
	reverse(&arr)
	fmt.Println(arr)
}

func reverse(s *[3]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
