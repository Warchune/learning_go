// ex3.3 prints an svg image, coloring its vertices based on their height.
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

const (
	width, height = 600, 320            // размер канвы в пикселях
	cells         = 150                 // количество ячеек сетки
	xyrange       = 20.0                // диапазон осей (-xyrange..+ xyrange)
	xyscale       = width / 2 / xyrange // Пикселей в единице x или y
	zscale        = height * 0.5        // Пикселей в единице z
	angle         = math.Pi / 6         // Угол осей x, y (=30)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	svg(os.Stdout)
}

func svg(w io.Writer) {
	zmin, zmax := minmaxConv()
	var k float64

	if math.Abs(zmin)+math.Abs(zmax) == 0 {
		k = 255 / 0.0001
	} else {
		k = 255./math.Abs(zmin) + math.Abs(zmax)
	}
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			if math.IsNaN(ax) || math.IsNaN(ay) || math.IsNaN(bx) || math.IsNaN(by) || math.IsNaN(cx) || math.IsNaN(cy) || math.IsNaN(dx) || math.IsNaN(dy) {
				continue
			}
			fmt.Fprintf(w, "<polygon style='stroke: #%s; fill: #%s' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				color(i, j, k), color(i, j, k), ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

// minmax возвращает минимальное и максимальное значения для z с учетом минимального/максимального значения x
// и y и предполагая квадратную область.
func minmaxConv() (min float64, max float64) {
	min = math.NaN()
	max = math.NaN()
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			minmaxPoligon(i, j, &min, &max)
		}
	}
	return
}

func corner(i, j int) (float64, float64) {
	// Находит точку(x,y) в углу ячейки (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Вычислет высоту поверхности z.
	z := f(x, y)

	// Спроецируем (x, y, z) изометрически на двухмерный холст SVG (sx, sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func color(i, j int, k float64) string {
	var hexColor, color string
	var err error
	min := math.NaN()
	max := math.NaN()
	minmaxPoligon(i, j, &min, &max)
	if math.Abs(max) > math.Abs(min) {
		intHexColor := int(math.Abs(min * k / 4))
		if intHexColor > 255 {
			intHexColor = 255
		}
		hexColor, err = ConvertInt(strconv.Itoa(intHexColor), 10, 16)
		if err != nil {
			log.Fatal(err)
		}
		color = fmt.Sprintf("%02s0000", hexColor)
	} else {
		intHexColor := int(math.Abs(max * k))
		hexColor, err = ConvertInt(strconv.Itoa(intHexColor), 10, 16)
		if err != nil {
			log.Fatal(err)
		}
		color = fmt.Sprintf("0000%02s", hexColor)
	}
	return color
}

func minmaxPoligon(i int, j int, min *float64, max *float64) {
	for xoff := 0; xoff <= 1; xoff++ {
		for yoff := 0; yoff <= 1; yoff++ {
			x := xyrange * (float64(i+xoff)/cells - 0.5)
			y := xyrange * (float64(j+yoff)/cells - 0.5)
			z := f(x, y)
			if math.IsNaN(*min) || z < *min {
				*min = z
			}
			if math.IsNaN(*max) || z > *max {
				*max = z
			}
		}
	}
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // дистанция от (0,0)
	return math.Sin(r) / r
}

func ConvertInt(val string, base, toBase int) (string, error) {
	i, err := strconv.ParseInt(val, base, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, toBase), nil
}
