package crypt

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"

	"test/consts"
)

type Cryptor struct {
	textField *widget.Entry
	label     widget.Label
	button    widget.Button
}

func NewCryptor(multiLineTextField *widget.Entry) *Cryptor {
	return &Cryptor{
		textField: multiLineTextField,
	}
}

func (c *Cryptor) SetWidgetsInCryptor() {
	c.label.Text = consts.LABEL_TEXT
	c.label.Wrapping = fyne.TextWrapWord
	c.textField.SetPlaceHolder("Enter here...")
	c.textField.Wrapping = fyne.TextWrapWord
	c.button.Text = consts.BUTTON_TEXT
	c.button.OnTapped = func() { c.setTextFromButton() }
}

func (c *Cryptor) setTextFromButton() {
	c.label.SetText(c.textField.Text)
}

func (c *Cryptor) GetTextFild() *widget.Entry {
	return c.textField
}

func (c *Cryptor) GetLabel() *widget.Label {
	return &c.label
}

func (c *Cryptor) GetButton() *widget.Button {
	return &c.button
}
