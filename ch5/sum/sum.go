package main

func main() {

	println(sum(1, 2, 3, 5))
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}
