package crypt

import (
	"fmt"
	"image/color"
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
	"test/menu"
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

func (c *Cryptor) SetWidgetsInCryptor() {
	c.textField.SetPlaceHolder(consts.PLACEHOLDER_TEXTFIELD)
	c.keyWord.SetPlaceHolder("Ключ-слово")
	c.textField.Wrapping = fyne.TextWrapWord
	c.button = *elem.NewButton(consts.BUTTON_TEXT, c.startCrypt)
	c.mainwindow = c.createMainWindow()
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
	}
}

func (c *Cryptor) checkKey() bool {
	if c.keyWord.Text == "" {
		c.showDialogKeyEmpty()
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
		&c.button,
	)
	container := container.NewWithoutLayout(btn)
	btn.Resize(fyne.NewSize(w, h))
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

	dialog.ShowCustom("Hello", "Press", widget.NewLabel("BB"), c.mainwindow)
}

func (c *Cryptor) createMainWindow() fyne.Window {
	window := c.App.NewWindow(consts.NAME_WINDOW_MAIN)

	mainMenu := menu.CreateMenu(c.GetTextFild())

	elem.LabelRules.TextStyle = fyne.TextStyle{Italic: true}

	containerProgressbar := container.NewWithoutLayout(c.GetProgressBar())
	c.GetProgressBar().Resize(fyne.NewSize(500, 10))
	c.GetProgressBar().Hide()

	window.SetMainMenu(mainMenu)
	window.SetContent(c.createContainers(containerProgressbar))
	window.Resize(fyne.NewSize(consts.WINDOW_WEIGHT, consts.WINDOW_HEIGHT))
	return window
}

func (c *Cryptor) createContainers(pb *fyne.Container) *fyne.Container {
	containerWithKeyAndButtonStart := container.NewGridWithColumns(2, c.GetKeyWordWithSize(200, 40), c.GetColorButtonWithSize(200, 40))
	containerFull := container.NewGridWithRows(4, elem.LabelRules, c.GetTextFild(), pb, containerWithKeyAndButtonStart)
	return containerFull
}
