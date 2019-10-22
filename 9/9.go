package main

import "image"
import "image/color"
import "golang.org/x/tour/pic"

type Image struct {
	w, h int
}

func (Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.w, img.h)
}

func (Image) At(x, y int) color.Color {
	return color.RGBA{0, 0, uint8(x - y), 255}
}

func main() {
	m := Image{100, 100}
	pic.ShowImage(m)
}
