package buff

import (
	"image"
	"image/color"
	"time"

	"github.com/carnivorecookies/cookiesmacro/assets"
	"github.com/disintegration/imaging"
)

var Precision = newBuff(time.Minute, assets.PrecisionImg)

// Buff represents a buff that is findable using buffbar, and its duration is possible to retrieve.
type buff struct {
	duration time.Duration
	color    color.NRGBA
	img      *image.NRGBA
}

func newBuff(duration time.Duration, img *image.NRGBA) buff {
	return buff{
		duration: duration,
		color:    img.At(0, 0).(color.NRGBA),
		img:      buffImg(img),
	}
}

// buffImg returns the [image.Image] equivalent of bytes (which must be a png).
// Panics if an error occurs.
// Scales to buff size (38x38); assumes the passed bytes represent a square image.
func buffImg(img *image.NRGBA) *image.NRGBA {
	return imaging.Resize(img, SidePx, SidePx, imaging.Linear)
}
