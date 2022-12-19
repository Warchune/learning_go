package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

const (
	width, height = 600, 320 // размер канвы в пикселях
	cells         = 100      // количество ячеек сетки
	xyrange       = 30.0     // диапазон осей
	// (-xyrange..+ xyrange)
	xyscale = width / 2 / xyrange // Пикселей в единице x или y
	zscale  = height * 0.5        // Пикселей в единице z
	angle   = math.Pi             // Углф осей x, y (=30)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) //sin(30), cos(30)

type zFunc func(x, y float64) float64

func main() {
	usage := "использование: 3.2 saddle|eggbox"
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}

	var f zFunc
	switch os.Args[1] {
	case "saddle":
		f = saddle
	case "eggbox":
		f = eggbox
	default:
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}
	svg(os.Stdout, f)
}

func saddle(x, y float64) float64 {
	a := 25.0
	b := 17.0
	a2 := a * a
	b2 := b * b
	return (y*y/a2 - x*x/b2)
}
func eggbox(x, y float64) float64 {
	return 0.2 * (math.Cos(x) + math.Cos(y))
}

func corner(i, j int, f zFunc) (float64, float64) {
	// Ищем угловую точку (x,y) ячейки (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Вычисляем высоту повернхоности z
	z := f(x, y)
	// Изометрически проецируем (x, y, z) на двумерную канву SVG (sx, xy)
	sx := width/2 + (x+y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func svg(w io.Writer, f zFunc) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, f)
			bx, by := corner(i, j, f)
			cx, cy := corner(i, j+1, f)
			dx, dy := corner(i+1, j+1, f)
			if math.IsNaN(ax) || math.IsNaN(ay) || math.IsNaN(bx) || math.IsNaN(by) || math.IsNaN(cx) || math.IsNaN(cy) || math.IsNaN(dx) || math.IsNaN(dy) {
				continue
			}
			fmt.Fprintf(w, "<polygon style='stroke: %s; fill: #222222' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				"#666666", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(w, "</svg>")
}
