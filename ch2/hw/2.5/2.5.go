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
		fmt.Println(PopCountShiftValue(uint64(num)))
		fmt.Println(PopCountRightmost(uint64(num)))
	}
}

func PopCountRightmost(x uint64) int {
	count := 0
	for x != 0 {
		x &= (x - 1)
		count++
	}
	return count
}

func PopCountShiftValue(x uint64) uint64 {
	mask := uint64(1)
	count := uint64(0)

	for i := 0; i < 64; i++ {
		if mask&x > 0 {
			count++
		}
		x >>= 1
	}
	return count
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
