package main

import (
	"fmt"
)

var pc [256]byte //p[i] - количество единичных битов в i

func main() {
	for i := range pc {
		fmt.Printf("pc[%d]= pc[%d/2](%d) + byte(%d&1)(%d)\t{", i, i, pc[i/2], i, byte(i&1))
		pc[i] = pc[i/2] + byte(i&1)
		fmt.Printf("%d}\n", pc[i])
	}

	fmt.Println(PopCount(4))
}

// PopCount возвращает степень заполнения
// (количество установленных битов) значения x.
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
