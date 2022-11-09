package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"test/consts"
)

type Cryptor struct {
	TextField       *widget.Entry
	Label           widget.Label
	ButtonTextField widget.Button
}

func main() {
	app := app.New()
	window := app.NewWindow(consts.NAME_WINDOW)

	multiLineTextField := widget.NewMultiLineEntry()

	Cryptor := NewCryptor(multiLineTextField)
	Cryptor.setWidgetsInCryptor()

	window.SetContent(
		container.NewGridWithRows(3, Cryptor.TextField, &Cryptor.ButtonTextField, &Cryptor.Label))
	window.Resize(fyne.NewSize(consts.WINDOW_WEIGHT, consts.WINDOW_HEIGHT))
	window.ShowAndRun()
}

func NewCryptor(multiLineTextField *widget.Entry) *Cryptor {
	return &Cryptor{
		TextField: multiLineTextField,
	}
}

func (c *Cryptor) setWidgetsInCryptor() {
	c.Label.Text = consts.LABEL_TEXT
	c.Label.Wrapping = fyne.TextWrapWord
	c.TextField.SetPlaceHolder("Enter here...")
	c.TextField.Wrapping = fyne.TextWrapWord
	c.ButtonTextField.Text = consts.BUTTON_TEXT
	c.ButtonTextField.OnTapped = func() { c.setTextFromButton() }
}

func (c *Cryptor) setTextFromButton() {
	c.Label.SetText(c.TextField.Text)
}
