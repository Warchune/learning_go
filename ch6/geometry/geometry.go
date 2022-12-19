package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }
type Path []Point

func main() {
	a := Point{1, 1}
	b := Point{1, 4}
	c := Point{5, 1}
	fmt.Println(a.Distance(b))
	fmt.Println(a.Distance(c))
	fmt.Println(b.Distance(c))

	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}

	fmt.Println(perim.Distance())

	z := Point{1, 2}
	(&z).ScaleBy(3)
	z.ScaleBy(3)
	fmt.Println(z)
}

// расстояние между двумя точками
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// Distance возвращает дину пути.
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
