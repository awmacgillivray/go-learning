package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

func (im Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (im Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 128, 128)
}

func (im Image) At(x, y int) color.Color {
	r := 100
	g := 25
	b := x^y
	return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
}

func nine() {
	m := Image{}
	pic.ShowImage(m)
}
