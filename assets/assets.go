package assets

import (
	"bytes"
	_ "embed"
	"image"
	"image/draw"
	"image/png"
)

func img(b []byte) *image.NRGBA {
	img, err := png.Decode(bytes.NewReader(b))
	if err != nil {
		panic(err)
	}

	// NRGBA seems to be more common (and is used by the imaging library)
	nrgba := image.NewNRGBA(image.Rectangle{img.Bounds().Min, img.Bounds().Max})
	draw.Draw(nrgba, nrgba.Bounds(), img, img.Bounds().Min, draw.Src)

	return nrgba
}

//go:embed Precision.png
var precisionBytes []byte

var PrecisionImg = img(precisionBytes)
