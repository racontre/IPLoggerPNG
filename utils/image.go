package utils

import (
	"image"
	"image/color"
	"strconv"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

  func addLabel(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{200, 100, 0, 255}
	point := fixed.Point26_6{X: fixed.I(x), Y: fixed.I(y)}
  
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
  }
  
  func GenerateImage(data []string) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, 400, 100))

	for index, ip := range data{
		//rows := len(data) / 2
        addLabel(img, 10 + 150 * (index / 5), 15 + (index % 5) * 15, strconv.Itoa(index + 1) + ") " + ip) //надо оборачивать в скобки / и % операции иначе то выйдет float
    }
  
	return img
  }