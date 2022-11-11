package mod

import (
	"fmt"
	"fyne.io/fyne/v2/widget"

	"test/consts"
)


var Btn = widget.NewButton(consts.BUTTON_TEXT, func() {
	fmt.Println(Entry.Text)
	Label.SetText(Entry.Text)
})
