package elem

import (
	"fmt"
	"fyne.io/fyne/v2/widget"
	"os"

	"test/consts"
)

var Btn = widget.NewButton(consts.BUTTON_TEXT, func() {
	fmt.Println(Entry.Text)
	Label.SetText(Entry.Text)
})

var SaveButton = widget.NewButton("Сохранить в файл .txt", func() {
	file, _ := os.Create("text.txt")
	defer file.Close()
	/* file.WriteString(Cryptor.GetTextFild().Text) */
})

