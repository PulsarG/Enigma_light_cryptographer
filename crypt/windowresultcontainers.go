package crypt

import (
	/* 	"fmt" */
	/* "io" */

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	/* "fyne.io/fyne/v2/dialog" */
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	/* 	"test/consts" */
	"test/elem"
)

func (r *Resulter) createContainersForWindowResult() *fyne.Container {

	labelResultCode := widget.NewLabel(r.code)
	labelResultCode.Wrapping = fyne.TextWrapWord
	containerLabelResult := container.NewScroll(labelResultCode)

	containatWithButtons := r.createContainerWithButtons()

	container := container.NewGridWithRows(2, containerLabelResult, containatWithButtons)
	return container
}

func (r *Resulter) createContainerWithButtons() *fyne.Container {
	saveFileButton := elem.NewButton("Сохранить в файл", r.saveFile)

	txtBound := binding.NewString()
	txtBound.Set(r.code)

	copyFileButton := widget.NewButtonWithIcon("Скопировать", theme.ContentCopyIcon(), func() {
		if content, err := txtBound.Get(); err == nil {
			r.windowResult.Clipboard().SetContent(content)
		}
	})
	saveBtn := container.NewWithoutLayout(saveFileButton)
	saveFileButton.Resize(fyne.NewSize(200, 40))
	saveFileButton.Move(fyne.NewPos(50, 0))
	copyBtn := container.NewWithoutLayout(copyFileButton)
	copyFileButton.Resize(fyne.NewSize(200, 40))
	copyFileButton.Move(fyne.NewPos(-50, 0))
	
	containatWithButtons := container.NewGridWithColumns(3,
		saveBtn,
		layout.NewSpacer(),
		copyBtn,
	)
	return containatWithButtons
}
