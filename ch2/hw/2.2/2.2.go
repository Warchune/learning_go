package main

import (
	"fmt"
	"os"
	"strconv"
)

type Metr float64
type Santimetr float64

func main() {
	if os.Args[1] != "" {
		for _, a := range os.Args[1:] {
			b, err := strconv.Atoi(a)
			if err != nil {
				fmt.Println("ощибка аргумента")
				return
			}
			fmt.Println(sToM(Santimetr(b)))
		}
	} else {
		fmt.Printf("Нет аргументов\n")
	}
}
func (m Metr) String() string {
	return fmt.Sprintf("%gm", m)
}

func (s Santimetr) String() string {
	return fmt.Sprintf("%gm", s)
}

func sToM(s Santimetr) Metr {
	return Metr(s / 100)
}

func mToS(m Metr) Santimetr {
	return Santimetr(m * 100)
}
