package main

import (
	"fmt"
)

func main() {
	println(max(1, 3, 9, 5))
	fmt.Println(max())
}

func max(vals ...int) (int, error) {
	if len(vals) < 1 {
		return 0, fmt.Errorf("max: non values")
	}
	max := vals[0]
	for _, val := range vals {
		if max < val {
			max = val
		}
	}
	return max, nil
}
