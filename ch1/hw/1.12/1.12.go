//server1 - минимальный "echo" - сервер со счетчиком запросов
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var palette = []color.Color{color.RGBA{0, 0, 0, 1}, color.RGBA{0, 255, 0, 1}}
var mu sync.Mutex
var count int

func main() {
	//fmt.Printf("%d, %s\n", os.Args, os.Args[1])
	handler := func(w http.ResponseWriter, r *http.Request) {
		var cycles int
		var err error
		cycles, err = strconv.Atoi(r.URL.Query().Get("cycles"))
		if err != nil {
			fmt.Println(err)
		}
		if cycles < 1 || cycles >= 100 {
			cycles = 10
		}
		lissajous(w, cycles)
		fmt.Printf("%d\n", os.Args)
	}
	http.HandleFunc("/", handler)

	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func lissajous(out io.Writer, cycles int) {
	var palette = []color.Color{color.RGBA{0, 0, 0, 1}, color.RGBA{0, 255, 0, 1}}

	const (
		//cycles  = 10    // количество полных колебаний x
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
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
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
	err := gif.EncodeAll(out, &anim) // Примечание: игнорируем ошибки
	if err != nil {
		log.Println(err)
	}
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
