// Пакет tempconv выполняет вычислеия температур
// по Цельсию
package main

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

func (c Celsius) String() string {
	return fmt.Sprintf("%gC", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%gF", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%gK", k)
}

const (
	AbsoluteZeroC Celsius = -273.15
	AbsoluteZeroK Kelvin  = 0
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func main() {
	fmt.Println(CToF(0))
	fmt.Println(FToC(32.0))
	c := FToC(212.0)
	fmt.Println(c)

}

func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9.0/5.0 + 32.0)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32.0) * 5.0 / 9.0)
}
