package utils

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/gif"
	"log"
	"net/http"
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
  
  // this image is returned on /{page}.png route
  func GenerateImage(data []string, parser *GeoIPParser) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, 400, 100))

	for index, ip := range data{
		x_margin := 10;
		y_margin := 5;
		x := 160 * (index / 5)	 + x_margin
		y := (index % 5) * 15	 + y_margin
		txt_x_offset := 25	// offseting to the right of the flag
		txt_y_offset := 10
		
        addLabel(img, x + txt_x_offset, y + txt_y_offset, strconv.Itoa(index + 1) + ") " + ip)
		if parser != nil { 
			countryCode, err := parser.GetCountry_DB(ip)
			if err != nil {log.Println("failed to draw flag: ", ip, countryCode, err)}

			if countryCode != "Unknown" {DrawFlag(img, -x, -y, countryCode)}
		}
	}
	return img
  }

  // shouldn't draw anything if there's no appropriate isocode
  func DrawFlag(img *image.RGBA, x, y int, flagCode string) error {
	flag, err := loadImageFromURL(fmt.Sprintf("https://www.translatorscafe.com/cafe/images/flags/%s.gif", flagCode))
	if err != nil {return err}

	draw.Draw(img, img.Bounds(), flag, image.Point{x, y}, draw.Src)

	return nil
  }

  func loadImageFromURL(URL string) (image.Image, error) {
    response, err := http.Get(URL)
    if err != nil { return nil, err }
    defer response.Body.Close()

    if response.StatusCode != 200 {
        return nil, errors.New("received non-200 response code")
    }

    img, _, err := image.Decode(response.Body)
    if err != nil { return nil, err }

    return img, nil
}

func loadImageLocal(path string) (image.Image, error){
	return nil, nil
}