package main

import (
	"fyne.io/fyne/v2/app"

	"test/crypt"
	"test/window"
)

func main() {
	app := app.New()

	Cryptor := crypt.NewCryptor(app)
	Cryptor.SetWidgetsInCryptor()

	windowMain := window.CreateMainWindow(app, Cryptor)

	windowMain.Show()
	windowMain.SetMaster()
	app.Run()
}
