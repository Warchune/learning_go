package main

import (
	"fmt"
)

func main() {
	map1 := map[string]int{
		"a": 1,
		"b": 2,
	}
	map2 := make(map[string]int)
	map2["a"] = 1
	map2["b"] = 2

	map3 := map[string]int{
		"a": 0,
		"b": 1,
	}
	fmt.Printf("%v\n", equal(map1, map2))
	fmt.Printf("%v\n", equal(map3, map2))
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
