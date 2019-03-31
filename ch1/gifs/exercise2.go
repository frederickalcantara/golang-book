package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var green = color.RGBA{0x00, 0xFF, 0x00, 0xFF}
var cyan = color.RGBA{0x00, 0xFF, 0xFF, 0xFF}
var magenta = color.RGBA{0xFF, 0x00, 0xFF, 0xFF}
var blue = color.RGBA{0x00, 0x00, 0xFF, 0xFF}
var red = color.RGBA{0xFF, 0x00, 0x00, 0xFF}
var yellow = color.RGBA{0xFF, 0xFF, 0x00, 0xFF}
var black = color.RGBA{0x00, 0x00, 0x00, 0xFF}
var multiColorPalette = []color.Color{black, green, cyan, magenta, blue, red, yellow}

const (
	foregroundIndex = 0
	backgroundIndex = 1
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, multiColorPalette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(i))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
