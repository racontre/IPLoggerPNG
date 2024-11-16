package utils

import (
	"image"
	"image/color"
	
	//"image/png"
	//"os"
  
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
  
  func GenerateImage(data [10]string) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, 400, 100))

	for index, ip := range data{
        addLabel(img, 10 + 150 * (index / 5), 15 + (index % 5) * 15, ip) //надо оборачивать в скобки / и % операции а то выйдет отстой
    }
  
	return img
	
	/*f, err := os.Create("data.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err := png.Encode(f, img); err != nil {
		panic(err)
	}*/
  }
/*
  func main() {
	GenerateImage()
  }*/