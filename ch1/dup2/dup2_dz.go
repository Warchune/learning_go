// dup2 выводит текст каждой строки, которая появляется во
// входных данных более одного раза. Программа читает
// стандартный ввод или список именованных файлов.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	fmt.Printf("\n")
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, arg string) {
	var flag int8
	input := bufio.NewScanner(f)
	for input.Scan() {
		//		counts[input.Text()]++
		if (counts[input.Text()] + 1) >= 2 {
			flag = 1
		}
	}
	if flag == 1 {
		fmt.Printf("%s, ", arg)
	}

	// Примечание: игнорируем потенциальные ошибки
	// из input.Err()
}
