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
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"test/anna"
	/* 	"test/elem" */
	/* 	"test/apps" */
	"test/consts"
	/* "test/elem" */)

type Cryptor struct {
	textField   *widget.Entry
	keyWord     widget.Entry
	button      widget.Button
	progressBar widget.ProgressBarInfinite

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
	c.button.Text = consts.BUTTON_TEXT
	c.button.OnTapped = func() { c.testCrypt() }
}

func (c *Cryptor) converToFloat() {
	if s, err := strconv.ParseFloat(c.textField.Text, 64); err == nil {
		fmt.Println(s)
		fmt.Println(reflect.TypeOf(s))
	} else {
		fmt.Println("Не верно введены цифры", err)
	}
}

func (c *Cryptor) testCrypt() {
	c.progressBar.Show()

	code, ready := anna.StartCrypt(c.textField.Text, c.keyWord.Text)

	if ready {
		c.Resulter.openWindowResult(code)
		c.progressBar.Hide()
	}
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

/* func (c *Cryptor) createWindowResult() fyne.Window {
	windowR := c.App.NewWindow(consts.NAME_WINDOW_RESULT)
	windowR.Resize(fyne.NewSize(consts.WINDOW_WEIGHT, consts.WINDOW_HEIGHT))

	windowR.SetContent(c.createContainersForWindowResult())

	return windowR
} */

/* func (c *Cryptor) openWindowResult() {
	if c.windowResult != nil {
		c.windowResult.Close()
	}
	c.windowResult = c.createWindowResult()
	c.windowResult.Show()
} */
/*
func (c *Cryptor) createContainersForWindowResult() *fyne.Container {
	containerWithResult := container.NewGridWithRows(1, &c.label)

	containerWithSaveButton := container.NewGridWithColumns(2, elem.SaveButton, elem.SaveButton)
	containerWithOpenButton := container.NewGridWithColumns(2, elem.OpenButton, elem.OpenButton)

	containatWithButtons := container.NewGridWithRows(2, containerWithSaveButton, containerWithOpenButton)

	container := container.NewGridWithRows(2, containerWithResult, containatWithButtons)
	return container
} */
