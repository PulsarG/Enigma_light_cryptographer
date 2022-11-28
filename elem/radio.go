package elem

import (
	"fyne.io/fyne/v2/widget"
)

var Radio *widget.RadioGroup

func CreateRadio() {
	var radios = widget.NewRadioGroup([]string{"Radio 1", "Radio 2", "Radio 3"}, func(s string) {})
	Radio = radios
}
