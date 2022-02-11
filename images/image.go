package images

import (
	"fmt"
	"image/color"
	"os"
	"strconv"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

func PNG(n int) ([]byte, error) {
	fpath := fmt.Sprintf("/tmp/%d.png", n)
	data, err := os.ReadFile(fpath)
	if err == nil {
		return data, nil
	}

	imgWidth := 100
	imgHeight := 100
	dc := gg.NewContext(imgWidth, imgHeight)
	dc.SetRGB(0, 0, 0)
	dc.DrawRectangle(0, 0, 100, 100)
	dc.Fill()
	dc.SetColor(color.White)

	ft, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return nil, err
	}
	fontSize := 20
	face := truetype.NewFace(ft, &truetype.Options{
		Size: float64(fontSize),
	})
	dc.SetFontFace(face)
	dc.DrawStringAnchored(strconv.Itoa(n), 10, float64(imgHeight-fontSize)/2, 0, 0.5)
	err = dc.SavePNG(fpath)
	if err != nil {
		return nil, err
	}

	return os.ReadFile(fpath)
}
