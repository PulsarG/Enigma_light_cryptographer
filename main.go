package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	/* "test/accord" */
	"test/consts"
	"test/crypt"
	"test/menu"
	"test/mod"
)

func main() {
	app := app.New()
	window := app.NewWindow(consts.NAME_WINDOW)

	Cryptor := crypt.NewCryptor()
	Cryptor.SetWidgetsInCryptor()

	mainMenu := menu.CreateMenu(Cryptor.GetTextFild(), Cryptor.GetLabel())
	mod.LabelRules.TextStyle = fyne.TextStyle{Italic: true}
	saveButton := createSaveButton(consts.BUTTON_SAVE_WEIGHT, consts.BUTTON_SAVE_HEIGHT)

	window.Resize(fyne.NewSize(consts.WINDOW_WEIGHT, consts.WINDOW_HEIGHT))
	window.SetMainMenu(mainMenu)
	window.SetContent(createContainers(Cryptor, saveButton))

	window.ShowAndRun()
}

func createContainers(Cryptor *crypt.Cryptor, saveButton *fyne.Container) *fyne.Container {
	containerWithKeyAndButtonStart := container.NewGridWithColumns(2, Cryptor.GetKeyWordWithSize(200, 40), Cryptor.GetColorButtonWithSize(200, 40))
	containerS := container.NewGridWithRows(5, mod.LabelRules, Cryptor.GetTextFild(), containerWithKeyAndButtonStart, Cryptor.GetLabel(), saveButton)
	/* fullContainer := container.NewGridWithRows(3, mod.LabelRules, containerS, Cryptor.GetLabel(), saveButton) */
	return containerS
}

func createSaveButton(w, h float32) *fyne.Container {
	saveButton := widget.NewButton("Сохранить в файл .txt", func() {
		file, _ := os.Create("text.txt")
		defer file.Close()
		/* file.WriteString(Cryptor.GetTextFild().Text) */
	})
	sb := container.NewWithoutLayout(saveButton)
	saveButton.Resize(fyne.NewSize(w, h))
	return sb
}
