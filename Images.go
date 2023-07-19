package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{}

func (Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)
}

func (Image) At(x, y int) color.Color {
	return color.RGBA{B: 255, A: 255}
}

func main() {
	m2 := Image{}
	pic.ShowImage(m2)
}
