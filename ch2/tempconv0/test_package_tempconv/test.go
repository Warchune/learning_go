package main

import (
	"fmt"
	"tempconv"
)

func main() {
	var c tempconv.Celsius = 0
	fmt.Println(tempconv.CToF(c))
}
