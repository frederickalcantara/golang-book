package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

var green = color.RGBA{0x00, 0xFF, 0x00, 0xFF}
var cyan = color.RGBA{0x00, 0xFF, 0xFF, 0xFF}
var magenta = color.RGBA{0xFF, 0x00, 0xFF, 0xFF}
var blue = color.RGBA{0x00, 0x00, 0xFF, 0xFF}
var red = color.RGBA{0xFF, 0x00, 0x00, 0xFF}
var yellow = color.RGBA{0xFF, 0xFF, 0x00, 0xFF}
var black = color.RGBA{0x00, 0x00, 0x00, 0xFF}
var white = color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}
var multiColorPalette = []color.Color{white, green, cyan, black, magenta, blue, red, yellow}

const (
	foregroundIndex = 0
	backgroundIndex = 1
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	cycles := 5
	res := 0.001
	size := 100
	nframes := 64
	delay := 8

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for param, values := range r.Form {
		var value = strings.Join(values, "")
		// values is a array of strings so it has to be converted into 1 string to be parsed based off of parameter

		switch param {
		case "cycles":
			if c, err := strconv.Atoi(value); err == nil {
				cycles = c
			}
		case "res":
			if r, err := strconv.ParseFloat(value, 64); err == nil {
				res = r
			}
		case "size":
			if s, err := strconv.Atoi(value); err == nil {
				size = s
			}
		case "nframes":
			if n, err := strconv.Atoi(value); err == nil {
				nframes = n
			}
		case "delay":
			if d, err := strconv.Atoi(value); err != nil {
				delay = d
			}
		}
	}

	lissajous(w, cycles, res, size, nframes, delay)
}

func lissajous(out io.Writer, cycles int, res float64, size int, nframes int, delay int) {
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, multiColorPalette)
		for t := 0.0; t < float64(cycles*2)*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), uint8(i))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
