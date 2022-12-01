package window

import (
	/* "fmt" */

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"test/consts"
	"test/crypt"
	"test/elem"
	"test/menu"
)

func CreateMainWindow(app fyne.App, Cryptor *crypt.Cryptor) fyne.Window {
	window := app.NewWindow(consts.NAME_WINDOW_MAIN)

	mainMenu := menu.CreateMenu(Cryptor.GetTextFild(), Cryptor.GetLabel())

	elem.LabelRules.TextStyle = fyne.TextStyle{Italic: true}

	progressBarInfinite := widget.NewProgressBarInfinite()
	containerProgressbar := container.NewWithoutLayout(progressBarInfinite)
	progressBarInfinite.Resize(fyne.NewSize(500, 10))

	window.SetMainMenu(mainMenu)
	window.SetContent(createContainers(Cryptor, containerProgressbar))
	window.Resize(fyne.NewSize(consts.WINDOW_WEIGHT, consts.WINDOW_HEIGHT))
	return window
}

func createContainers(Cryptor *crypt.Cryptor, pb *fyne.Container) *fyne.Container {
	containerWithKeyAndButtonStart := container.NewGridWithColumns(2, Cryptor.GetKeyWordWithSize(200, 40), Cryptor.GetColorButtonWithSize(200, 40))
	containerFull := container.NewGridWithRows(4, elem.LabelRules, Cryptor.GetTextFild(), pb, containerWithKeyAndButtonStart)
	return containerFull
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
