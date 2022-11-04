package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"test/src"
)

func main() {
	// создание сущности приложения
	a := app.New()

	// создание окна
	w := a.NewWindow(src.NAME_WINDOW)

	// Текстовая сущность для контента
	label := widget.NewLabel(src.LABEL_TEXT)
	label2 := widget.NewLabel(src.LABEL_TEXT_TWO)

	// Создание поля ввода
	entry := widget.NewEntry()

	// Создание виджета кнопки
	btn := widget.NewButton(src.BUTTON_TEXT, func() {
		fmt.Println(entry.Text)


		// Присвоение Лейблу-тексту значения из поля ввода
		label.SetText(entry.Text)
	})

	// Задаем контент для окна
	w.SetContent(

		//  контейнер сущность для внедрения объектов
		container.NewVBox(

			// помещаем визуальный объект в контейнер
			label,
			label2,
			entry,
			btn,
		))

	// Вывод окна на экран
	w.ShowAndRun()
}
