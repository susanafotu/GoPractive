package main


import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)


type Image struct{
}


func (p Image) At(x, y int) color.Color{
	return color.RGBA{uint8(x),uint8(y),255,255}
}


//for bound method, choose arbitrary values for w, h
func (p Image) Bounds() image.Rectangle{
	return image.Rect(0,0,200,200)
}


func (p Image) ColorModel() color.Model{
	return color.RGBAModel
}


func main() {
	m := Image{}
	pic.ShowImage(m)
}

