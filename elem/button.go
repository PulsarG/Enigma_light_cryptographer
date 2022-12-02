package elem

import (
	"fmt"
	"fyne.io/fyne/v2/widget"
	/* "os"

	"test/consts" */
)

/* var Btn = widget.NewButton(consts.BUTTON_TEXT, func() {
	fmt.Println(Entry.Text)
	Label.SetText(Entry.Text)
}) */

var SaveButton = widget.NewButton("Сохранить в файл .txt", func() {
	fmt.Println("777")
})

var OpenButton = widget.NewButton("Открыть файл", func() {
	fmt.Println("000")
})
