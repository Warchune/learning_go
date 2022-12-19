package main

import (
	"fmt"
)

type tree struct {
	value       int
	left, right *tree
}

func main() {
	var arr []int
	arr = append(arr, 3, 4, 0, 12, 112, 800, 545)
	fmt.Println(arr[:1])
	fmt.Println(arr)
	Sort(arr)
	fmt.Println(arr)
}

//Sort сортирует значения на месте
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues добавляет элементы t к values в требуемом порядке
// и возвращает результирующий срез.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Эквивалентно возврату &tree{value: value}
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
