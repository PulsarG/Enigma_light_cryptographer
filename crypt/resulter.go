package crypt

import (
	/* "fmt" */
	"io"

	"fyne.io/fyne/v2"
/* 	"fyne.io/fyne/v2/container" */
	"fyne.io/fyne/v2/dialog"
	/* "fyne.io/fyne/v2/layout" */
	"fyne.io/fyne/v2/widget"

	"test/consts"
	/* "test/elem" */
)

type Resulter struct {
	code   string
	/* button widget.Button */

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

func (r *Resulter) saveFile() {
	dialog.ShowFileSave(
		func(uc fyne.URIWriteCloser, err error) {
			io.WriteString(uc, r.code)
		}, r.windowResult,
	)

}