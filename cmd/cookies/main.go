package main

import (
	"errors"
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	adwtheme "fyne.io/x/fyne/theme"
	"github.com/carnivorecookies/cookiesmacro/buff"
	"github.com/carnivorecookies/cookiesmacro/ui"
)

func errPopup(win fyne.Window, err error) func() {
	label := widget.NewLabel(fmt.Sprintf("Unexpected error occurred: %v", err))
	popup := widget.NewPopUp(label, win.Canvas())

	return popup.Show
}

func main() {
	a := app.New()
	a.Settings().SetTheme(adwtheme.AdwaitaTheme())
	w := a.NewWindow("Toolbar Widget")

	label := widget.NewLabel("Unknown")
	w.SetContent(container.NewStack(label))
	go func() {
		for {
			time.Sleep(2 * time.Second)

			d, err := buff.Precision.Duration()
			switch {
			case errors.Is(err, buff.BuffNotFound):
				fyne.Do(func() { label.SetText("Buff not found") })
				continue
			case errors.Is(err, buff.RobloxInactive):
				fyne.Do(func() { label.SetText("Roblox window inactive") })
				continue
			}

			fyne.Do(func() {
				label.SetText(fmt.Sprintf("Precision: %v seconds", d.Seconds()))
			})

		}
	}()
	w.Show()
	ui.AlwaysOnTop(w)
	a.Run()
}
