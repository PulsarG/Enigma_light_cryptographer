package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"

	"test/consts"
	"test/mod"
)

func main() {
	app := app.New()
	window := app.NewWindow(consts.NAME_WINDOW)

	mod.CreateChecks()
	mod.CreateRadio()
	mod.CreateColorLabel()

	circle := canvas.NewCircle(color.NRGBA{R: 100, G: 150, B: 210, A: 255})
	circle.StrokeColor = color.NRGBA{R: 190, G: 110, B: 110, A: 255}
	circle.StrokeWidth = 10

	
	window = setContentInWindow(window, circle)

	window.Resize(fyne.NewSize(consts.WINDOW_WEIGHT, consts.WINDOW_HEIGHT))
	window.ShowAndRun()
}

func setContentInWindow(window fyne.Window, Circle *canvas.Circle) fyne.Window {
	Containar := container.NewVBox(
		mod.Label,
		mod.Label2,
		mod.Entry,
		mod.Btn,
		mod.Checks,
		mod.Radio,
	)

	window.SetContent(
		container.NewGridWithColumns(4, Circle, mod.ColorText, mod.Image, Containar))

	return window
}
