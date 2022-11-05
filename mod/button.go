package mod

import (
	"fmt"
	"fyne.io/fyne/v2/widget"

	"test/consts"
)

// Создание виджета кнопки
var Btn = widget.NewButton(consts.BUTTON_TEXT, func() {
	fmt.Println(Entry.Text)

	// Присвоение Лейблу-тексту значения из поля ввода
	Label.SetText(Entry.Text)
	Label2.SetText(Entry.Text)
})
