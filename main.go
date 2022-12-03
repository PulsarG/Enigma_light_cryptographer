package main

import (
	"fyne.io/fyne/v2/app"

	"test/crypt"
)

func main() {
	app := app.New()

	Cryptor := crypt.NewCryptor(app)
	Cryptor.SetWidgetsInCryptor()

	windowMain := Cryptor.GetWindowMain()

	windowMain.Show()
	windowMain.SetMaster()
	app.Run()
}
