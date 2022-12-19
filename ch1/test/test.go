package main

import (
	"fmt"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RUB
)

func main() {
	sumbol := [...]string{USD: "$", EUR: "E", GBP: "G", RUB: "P"}
	fmt.Println(RUB, sumbol[RUB])
}
