package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"test/consts"
	"test/mod"
)

func main() {
	app := app.New()
	window := app.NewWindow(consts.NAME_WINDOW)
	
	mod.CreateChecks()
	mod.CreateRadio()

	window.SetContent(
		container.NewVBox(
			mod.Label,
			mod.Label2,
			mod.Entry,
			mod.Btn,
			mod.Checks,
			mod.Radio,
		))

	window.Resize(fyne.NewSize(consts.WINDOW_WEIGHT, consts.WINDOW_HEIGHT))
	window.ShowAndRun()
}
