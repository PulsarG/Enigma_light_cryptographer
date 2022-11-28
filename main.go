package main

import (
	/* "fmt" */

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"test/consts"
	"test/crypt"
	"test/elem"
	"test/menu"
)

func main() {
	app := app.New()
	window := app.NewWindow(consts.NAME_WINDOW)

	Cryptor := crypt.NewCryptor()
	Cryptor.SetWidgetsInCryptor()

	mainMenu := menu.CreateMenu(Cryptor.GetTextFild(), Cryptor.GetLabel())
	elem.LabelRules.TextStyle = fyne.TextStyle{Italic: true}
	saveButton := createSaveButton(elem.SaveButton, consts.BUTTON_SAVE_WEIGHT, consts.BUTTON_SAVE_HEIGHT)

	setHideSaveButton := widget.NewCheck("Show/Hide Save Button", func(b bool) {
		if b == true {
			saveButton.Hide()
		} else {
			saveButton.Show()
		}

		showDialogWindow(window)
	})

	window.Resize(fyne.NewSize(consts.WINDOW_WEIGHT, consts.WINDOW_HEIGHT))
	window.SetMainMenu(mainMenu)
	window.SetContent(createContainers(Cryptor, saveButton, setHideSaveButton))

	window.ShowAndRun()
}

func createContainers(Cryptor *crypt.Cryptor, saveButton *fyne.Container, setHideSaveButton *widget.Check) *fyne.Container {
	containerWithKeyAndButtonStart := container.NewGridWithColumns(2, Cryptor.GetKeyWordWithSize(200, 40), Cryptor.GetColorButtonWithSize(200, 40))
	containerS := container.NewGridWithRows(6, elem.LabelRules, Cryptor.GetTextFild(), containerWithKeyAndButtonStart, Cryptor.GetLabel(), saveButton, setHideSaveButton)
	return containerS
}

func createSaveButton(SB *widget.Button, w, h float32) *fyne.Container {
	sb := container.NewWithoutLayout(SB)
	SB.Resize(fyne.NewSize(w, h))
	return sb
}

func showDialogWindow(w fyne.Window) {
	/* dialog.ShowConfirm("Dialog Window", "Hello, World", func(b bool) {
		if b == true {
			fmt.Println("Yes")

		} else {
			fmt.Println("No")
		}
	}, w) */

	/* dialog.ShowCustomConfirm("Hello", "Да", "Нет", widget.NewLabel("Set answer"), func(b bool) {
		if b == true {
			fmt.Println("Yes")

		} else {
			fmt.Println("No")
		}
	}, w) */

	dialog.ShowCustom("Hello", "Press", widget.NewLabel("BB"), w)
}
