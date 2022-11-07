package mod

import (
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"

	"image/color"
	"test/consts"
)

// Текстовая сущность для контента

var Label = widget.NewLabel(consts.LABEL_TEXT)
var Label2 = widget.NewLabel(consts.LABEL_TEXT_TWO)

var ColorText *canvas.Text
var COLOR_FOR_ColorText = color.NRGBA{R: 100, G: 150, B: 210, A: 255}
func CreateColorLabel() {
	colortext := canvas.NewText("Text", COLOR_FOR_ColorText)
	ColorText = colortext
}
