package buff

import (
	"errors"
	"image"
	"image/color"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/vcaesar/gcv"
)

var BuffNotFound = errors.New("Buff not found")
var RobloxInactive = errors.New("Roblox window not active")

// SidePx is the number of pixels wide/tall of a buff.
const SidePx = 38

// buffBarY is the y position of the buff bar.
const buffBarY = 58

var _, screenWidth = robotgo.GetScreenSize()

// Duration gets the current amount of time left on the buff.
//
// Will return BuffNotFound if the buff is inactive. Additionally, returns RobloxInactive if the roblox window is not active.
func (b buff) Duration() (time.Duration, error) {
	currentBuff, err := b.screenshotBuff()
	if err != nil {
		return 0, err
	}

	for i := range SidePx {
		color := currentBuff.At(0, i).(color.NRGBA)
		if colorsEqual(color, b.color) {
			// inverse of the regular height; top pixel is SidePx, bottom is 0
			pixelY := SidePx - i
			duration := float64(pixelY) / float64(SidePx) * float64(b.duration)
			return time.Duration(duration), nil
		}
	}

	return 0, nil
}

// screenshotBuff screenshots the buff in the image buff. See [buff.find] for details on errors returned.
func (b buff) screenshotBuff() (image.Image, error) {
	point, err := b.find()
	if err != nil {
		return nil, err
	}

	buffImg, err := captureRoblox(point.X, buffBarY, SidePx, SidePx)
	if err != nil {
		return nil, err
	}

	return buffImg, nil
}

// Find finds the buff in the image buff. Assumes the buff is of the correct size of [SidePx] by [SidePx].
//
// Returns [BuffNotFound] if the buff could not be found, and returns [RobloxInactive] if the roblox window is not active at call time.
// May return other errors, e.g. if screenshotting fails.
func (b buff) find() (image.Point, error) {
	buffBar, err := screenshotBar()
	if err != nil {
		return image.Point{}, err
	}

	_, closeness, _, point := gcv.FindImg(b.img, buffBar)

	/*
		-1 < closeness < 1 (higher the better)
		From my testing, when a buff such as Precision first activates, it has a relatively high value (0.95+) but rapidly declines as its stack increases.
		The stack number seems to greatly affect the value, so use 0.5 for safety
	*/
	if closeness < 0.5 {
		return image.Point{}, BuffNotFound
	}

	return point, nil
}

// ScreenshotBar takes a screenshot of the buff bar.
func screenshotBar() (image.Image, error) {
	// The buff bar is always 58 pixels below the top of the screen.
	ss, err := captureRoblox(0, buffBarY, screenWidth, SidePx)

	return ss, err
}

func captureRoblox(x, y, w, h int) (image.Image, error) {
	if robotgo.GetTitle() != "Roblox" {
		return nil, RobloxInactive
	}

	return robotgo.CaptureImg(x, y, w, h)
}

// colorsEqual returns whether two colors are almost the same (within + or - five of each channel).
// Ignores the alpha channel.
func colorsEqual(c1, c2 color.NRGBA) bool {
	eql := func(n1 uint8, n2 uint8) bool {
		if n1 > n2 {
			return n1-n2 <= 5
		}

		return n2-n1 <= 5
	}

	return eql(c1.R, c2.R) && eql(c1.G, c2.G) && eql(c1.B, c2.B)
}
