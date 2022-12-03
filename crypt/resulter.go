package crypt

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"test/consts"
	"test/elem"
)

type Resulter struct {
	code   string
	button widget.Button

	App          fyne.App
	windowResult fyne.Window
}

func NewResulter(app fyne.App) *Resulter {
	return &Resulter{
		App: app,
	}
}

func createLabel(s string) *widget.Label {
	l := widget.NewLabel(s)
	return l
}

func (r *Resulter) createWindowResult() *fyne.Window {
	windowR := r.App.NewWindow(consts.NAME_WINDOW_RESULT)
	windowR.Resize(fyne.NewSize(consts.WINDOW_WEIGHT, consts.WINDOW_HEIGHT))

	windowR.SetContent(r.createContainersForWindowResult())

	return &windowR
}

func (r *Resulter) openWindowResult(code string) {
	r.code = code
	if r.windowResult != nil {
		r.windowResult.Close()
	}
	r.windowResult = *r.createWindowResult()
	r.windowResult.Show()
}

func (r *Resulter) createContainersForWindowResult() *fyne.Container {

	labelResultCode := widget.NewLabel(r.code)
	labelResultCode.Wrapping = fyne.TextWrapWord
	containerWithResult := container.NewGridWithRows(1, labelResultCode)

	saveFileButton := elem.NewButton("Сохранить в файл .txt", saveFile)
	openFileButton := elem.NewButton("Открыть файл", openFile)
	containerWithSaveButton := container.NewGridWithColumns(2, saveFileButton, saveFileButton)
	containerWithOpenButton := container.NewGridWithColumns(2, openFileButton, openFileButton)

	containatWithButtons := container.NewGridWithRows(2, containerWithSaveButton, containerWithOpenButton)

	container := container.NewGridWithRows(2, containerWithResult, containatWithButtons)
	return container
}

func (r *Resulter) closeResultWindow() {
	r.windowResult.Close()
}

func saveFile() {
	fmt.Println("File is Save")
}

func openFile() {
	fmt.Println("File is Open")
}
