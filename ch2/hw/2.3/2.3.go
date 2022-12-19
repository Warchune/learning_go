package main

import (
	"fmt"
	"os"
	"strconv"
)

var pc [256]byte //p[i] - количество единичных битов в i

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	for _, arg := range os.Args[1:] {
		num, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println(err)
			fmt.Println("использование: 2.3 <целое число> ...")
			break
		}
		fmt.Println(PopCountLoop(uint64(num)))
		fmt.Println(PopCount(uint64(num)))
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCount возвращает степень заполнения
// (количество установленных битов) значения x.
func PopCountLoop(x uint64) int {
	var res int
	for i := 0; i != 8; i++ {
		res += int(pc[byte(x>>(i*8))])
	}
	return res
}
