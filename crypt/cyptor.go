package crypt

import (
	"fmt"
	"image/color"
	"io"
	"reflect"
	"strconv"

	/* "strings" */

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"test/anna"
	"test/elem"
	/* 	"test/menu" */
	/* "test/window" */
	/* 	"test/apps" */
	"test/consts"
	/* "test/elem" */)

type Cryptor struct {
	textField   *widget.Entry
	keyWord     widget.Entry
	button      widget.Button
	progressBar widget.ProgressBarInfinite
	mainwindow  fyne.Window

	App      fyne.App
	Resulter *Resulter
}

func NewCryptor(app fyne.App) *Cryptor {
	return &Cryptor{
		textField: createMultiLineEntry(),
		App:       app,
		Resulter:  NewResulter(app),
	}
}


func createMultiLineEntry() *widget.Entry {
	multiLineTextField := widget.NewMultiLineEntry()
	return multiLineTextField
}

func (c *Cryptor) converToFloat() {
	if s, err := strconv.ParseFloat(c.textField.Text, 64); err == nil {
		fmt.Println(s)
		fmt.Println(reflect.TypeOf(s))
	} else {
		fmt.Println("Не верно введены цифры", err)
	}
}

func (c *Cryptor) startCrypt() {
	if c.checkKey() {
		c.progressBar.Show()

		code, ready := anna.StartCrypt(c.textField.Text, c.keyWord.Text)

		if ready {
			c.Resulter.openWindowResult(code)
			c.progressBar.Hide()
		}
	} else {
		c.showDialogKeyEmpty()
	}
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
	return &c.button
}

func (c *Cryptor) GetColorButtonWithSize(w, h float32) *fyne.Container {
	color := color.RGBA{11, 78, 150, 1}
	btn := container.New(
		layout.NewMaxLayout(),
		canvas.NewRectangle(color),
		elem.NewButton(consts.BUTTON_TEXT, c.startCrypt),
	)
	container := container.NewWithoutLayout(btn)
	btn.Resize(fyne.NewSize(w, h))
	btn.Move(fyne.NewPos(-20, 0))
	return container
}

func (c *Cryptor) GetProgressBar() *widget.ProgressBarInfinite {
	return &c.progressBar
}

func (c *Cryptor) showDialogKeyEmpty() {

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
	fmt.Println("3")
	dialog.ShowCustom(
		consts.DIALOG_KEY_WINDOW_TITILE,
		consts.DIALOG_KEY_BTN_TITLE,
		widget.NewLabel(consts.DIALOG_KEY_ERROR_TEXT),
		c.mainwindow)
	fmt.Println("4")

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
