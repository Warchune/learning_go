// Lissajous генерирует анимированный GIF из случайных
// фигур Лиссажу.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.RGBA{0, 0, 0, 1}, color.RGBA{0, 255, 0, 1}}

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 10    // количество полных колебаний x
		res     = 0.001 // угловое разрешение
		size    = 100   // Канва изображения охватывает [size..+size]
		nframes = 64    // Количество кадров анимации
		delay   = 8     // Задержка между кадрами (единица - 10мс)
	)

	var rr, gg, bb uint8
	var randColorRGB color.RGBA

	rand.Seed(time.Now().UTC().UnixNano())

	freq := rand.Float64() * 3.0 //Относительная частота колебаний y
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5),
				size+int(y*size+0.5),
				uint8(i))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)

		rr = RandColorRGB()
		gg = RandColorRGB()
		bb = RandColorRGB()
		randColorRGB = color.RGBA{rr, gg, bb, 1}

		palette = append(palette, randColorRGB)
	}
	gif.EncodeAll(out, &anim) // Примечание: игнорируем ошибки
}

func RandColorRGB() uint8 {
	rand.Seed(time.Now().UTC().UnixNano())

	var a uint8
	for {
		a = uint8(rand.Int() % 1000)
		if a > (255) {
			continue
		} else {
			break
		}
	}
	return a
}
