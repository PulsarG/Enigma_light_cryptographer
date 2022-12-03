package crypt

import (
	"fmt"
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
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
	containerLabelResult := container.NewWithoutLayout(labelResultCode)
	labelResultCode.Resize(fyne.NewSize(500, 400))
	/* containerWithResult := container.NewGridWithRows(1, containerLabelResult) */

	containatWithButtons := r.createContainerWithButtons()

	container := container.NewGridWithRows(2, containerLabelResult, containatWithButtons)
	return container
}

func (r *Resulter) createContainerWithButtons() *fyne.Container {
	saveFileButton := elem.NewButton("Сохранить в файл", r.saveFile)
	openFileButton := elem.NewButton("Открыть файл", r.openFile)

	containerSaveButton := container.NewWithoutLayout(saveFileButton)
	saveFileButton.Resize(fyne.NewSize(200, 50))
	containerOpenButton := container.NewWithoutLayout(openFileButton)
	openFileButton.Resize(fyne.NewSize(200, 50))

	containatWithButtons := container.NewGridWithColumns(2, containerSaveButton, containerOpenButton)
	return containatWithButtons
}

func (r *Resulter) closeResultWindow() {
	r.windowResult.Close()
}

func (r *Resulter) saveFile() {
	dialog.ShowFileSave(
		func(uc fyne.URIWriteCloser, err error) {
			io.WriteString(uc, r.code)
		}, r.windowResult,
	)
	fmt.Println("File is Save")
}

func (r *Resulter) openFile() {
	dialog.ShowFileOpen(
		func(uc fyne.URIReadCloser, err error) {
			data, _ := io.ReadAll(uc)

			d := string(data)
			r.openWindowResult(d)
		}, r.windowResult,
	)
	fmt.Println("File is Open")
}
