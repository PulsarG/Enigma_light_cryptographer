package elem

import (
	"fyne.io/fyne/v2/widget"
)

func NewButton(label string, f func()) *widget.Button {
	newButton := widget.NewButton(label, f)
	return newButton
}
