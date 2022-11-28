package elem

import (
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"

	"image/color"
	"test/consts"
)

var Label = widget.NewLabel(consts.LABEL_RESULT)
var LabelRules = widget.NewLabel(consts.LABEL_RULES)

var ColorText *canvas.Text
var COLOR_FOR_ColorText = color.NRGBA{R: 100, G: 150, B: 210, A: 255}
func CreateColorLabel() {
	colortext := canvas.NewText("Text", COLOR_FOR_ColorText)
	ColorText = colortext
}
