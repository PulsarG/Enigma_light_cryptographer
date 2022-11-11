package accord

import (
	"fyne.io/fyne/v2/widget"

	"test/consts"
)

func CreateAccordeon(btn *widget.Button) *widget.Accordion {
	accordeonItemOne := widget.NewAccordionItem(consts.ACCORDEON_ONE_TITLE, btn)
	accordeon := widget.NewAccordion(accordeonItemOne)
	return accordeon
}
