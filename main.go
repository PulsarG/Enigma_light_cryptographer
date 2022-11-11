package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"test/accord"
	"test/consts"
	"test/crypt"
	"test/menu"
	"test/mod"
)

func main() {
	app := app.New()
	window := app.NewWindow(consts.NAME_WINDOW)

	/* multiLineTextField := widget.NewMultiLineEntry() */

	Cryptor := crypt.NewCryptor()
	Cryptor.SetWidgetsInCryptor()

	saveButton := widget.NewButton("Сохранить в файл .txt", func() {
		file, _ := os.Create("text.txt")

		defer file.Close()

		file.WriteString(Cryptor.GetTextFild().Text)
	})

	/* Container.Resize(fyne.NewSize(200, 200))
	Container.Move(fyne.NewPos(50, 50))
	ccc := container.NewWithoutLayout(Container) */

	mainMenu := menu.CreateMenu(Cryptor.GetTextFild(), Cryptor.GetLabel())
	accordeon := accord.CreateAccordeon(saveButton)

	window.Resize(fyne.NewSize(consts.WINDOW_WEIGHT, consts.WINDOW_HEIGHT))
	window.SetMainMenu(mainMenu)
	container := container.NewGridWithRows(5, mod.Label2, Cryptor.GetTextFild(), Cryptor.GetColorButton(), accordeon, Cryptor.GetLabel())
	window.SetContent(container)
	window.ShowAndRun()
}
