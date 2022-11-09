package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"test/consts"
	"test/crypt"
)

func main() {
	app := app.New()
	window := app.NewWindow(consts.NAME_WINDOW)

	multiLineTextField := widget.NewMultiLineEntry()

	Cryptor := crypt.NewCryptor(multiLineTextField)
	Cryptor.SetWidgetsInCryptor()

	Container := container.NewGridWithRows(3, Cryptor.GetTextFild(), Cryptor.GetButton(), Cryptor.GetLabel())
	window.SetContent(Container)
	window.Resize(fyne.NewSize(consts.WINDOW_WEIGHT, consts.WINDOW_HEIGHT))
	window.ShowAndRun()
}
