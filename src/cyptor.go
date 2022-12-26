package src

import (
	/* "fmt" */
	"image/color"
	"io"
	/* "reflect"
	"strconv"
	*/
	/* "strings" */

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"enigma/elem"
	"enigma/enigmasistem"

	/* 	"enigmasistem/menu" */
	/* "enigmasistem/window" */
	"enigma/consts"
)

type Cryptor struct {
	textField      *widget.Entry
	keyWord        widget.Entry
	button         *widget.Button
	btnContainer   fyne.Container
	buttonOpenFile *widget.Button
	progressBar    widget.ProgressBarInfinite
	mainwindow     fyne.Window

	App      fyne.App
	content  *fyne.Container
	menu     *fyne.MainMenu
	Resulter *Resulter

	isRus bool
}

func NewCryptor(app fyne.App) *Cryptor {
	return &Cryptor{
		textField: createMultiLineEntry(),
		App:       app,
		Resulter:  NewResulter(app),
		isRus:     true,
	}

}

func createMultiLineEntry() *widget.Entry {
	multiLineTextField := widget.NewMultiLineEntry()
	return multiLineTextField
}

func (c *Cryptor) startCrypt() {
	/* if c.checkKey() { */
		c.progressBar.Show()

		code, ready := enigmasistem.StartCrypt(c.textField.Text, c.keyWord.Text)

		if ready {
			c.Resulter.openWindowResult(code)
			c.progressBar.Hide()
		} else {
			c.progressBar.Hide()
			c.showDialogWrong(code)
			return
		}
/* 
	} else {
		c.showDialogKeyEmpty()
	} */
}

func (c *Cryptor) checkKey() bool {
	if c.keyWord.Text == "" {
		return false
	} else {
		return true
	}
}

func (c *Cryptor) GetWindowMain() fyne.Window {
	return c.mainwindow
}

func (c *Cryptor) GetTextFild() *widget.Entry {
	return c.textField
}

func (c *Cryptor) GetKeyWordWithSize(w, h float32) *fyne.Container {
	container := container.NewWithoutLayout(&c.keyWord)
	c.keyWord.Resize(fyne.NewSize(w, h))
	c.keyWord.Move(fyne.NewPos(20, 0))
	return container
}

func (c *Cryptor) GetButton() *widget.Button {
	c.button = elem.NewButton(consts.BUTTON_TEXT, c.startCrypt)
	return c.button
}

func (c *Cryptor) GetColorButtonWithSize(w, h float32) *fyne.Container {
	c.button = elem.NewButton(consts.BUTTON_TEXT, c.startCrypt)
	color := color.RGBA{11, 78, 150, 1}
	btn := container.New(
		layout.NewMaxLayout(),
		canvas.NewRectangle(color),
		c.button,
	)
	container := container.NewWithoutLayout(btn)
	c.button.Resize(fyne.NewSize(w, h))
	c.button.Move(fyne.NewPos(-30, 0))
	return container
}

func (c *Cryptor) GetProgressBar() *widget.ProgressBarInfinite {
	return &c.progressBar
}

func (c *Cryptor) showDialogKeyEmpty() {
	dialog.ShowCustom(
		consts.DIALOG_KEY_WINDOW_TITLE,
		consts.DIALOG_KEY_BTN_TITLE,
		widget.NewLabel(consts.DIALOG_KEY_ERROR_TEXT),
		c.mainwindow)
}

func (c *Cryptor) openFile() {
	var d string
	if c.checkKey() {
		dialog.ShowFileOpen(
			func(uc fyne.URIReadCloser, err error) {
				if uc != nil {
					data, _ := io.ReadAll(uc)
					d = string(data)
					c.textField.SetText(d)
				} else {
					return
				}
			}, c.mainwindow,
		)
	} else {
		c.showDialogKeyEmpty()
	}
}

func (c *Cryptor) showDialogWrong(s string) {
	dialog.ShowCustom(
		consts.DIALOG_KEY_WINDOW_TITLE,
		"ОК",
		widget.NewLabel(s),
		c.mainwindow)
}
