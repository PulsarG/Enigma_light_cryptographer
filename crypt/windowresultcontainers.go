package crypt

import (
/* 	"fmt" */
	/* "io" */

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	/* "fyne.io/fyne/v2/dialog" */
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

/* 	"test/consts" */
	"test/elem"
)

func (r *Resulter) createContainersForWindowResult() *fyne.Container {

	labelResultCode := widget.NewLabel(r.code)
	labelResultCode.Wrapping = fyne.TextWrapWord
	containerLabelResult := container.NewVBox(labelResultCode)
	/* labelResultCode.Resize(fyne.NewSize(500, 400)) */
	/* containerWithResult := container.NewGridWithRows(1, containerLabelResult) */

	containatWithButtons := r.createContainerWithButtons()

	container := container.NewVBox(containerLabelResult, containatWithButtons)
	return container
}

func (r *Resulter) createContainerWithButtons() *fyne.Container {
	saveFileButton := elem.NewButton("Сохранить в файл", r.saveFile)
	/* openFileButton := elem.NewButton("Открыть файл", r.openFile) */

	containerSaveButton := container.NewWithoutLayout(saveFileButton)
	/* saveFileButton.Resize(fyne.NewSize(200, 50)) */
	saveFileButton.Resize(fyne.NewSize(200, 40))
	saveFileButton.Move(fyne.NewPos(-50, 10))
	/* containerOpenButton := container.NewWithoutLayout(openFileButton)
	openFileButton.Resize(fyne.NewSize(200, 50)) */

	containatWithButtons := container.NewHBox(layout.NewSpacer(), containerSaveButton)
	return containatWithButtons
}

/* func (r *Resulter) closeResultWindow() {
	r.windowResult.Close()
} */