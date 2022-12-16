package src

import (
	/* "fmt"
	"image/color"
	"io"
	"reflect"
	"strconv"
	*/
	/* "strings" */

	"fyne.io/fyne/v2"
	/* 	"fyne.io/fyne/v2/canvas" */
	"fyne.io/fyne/v2/container"
	/* "fyne.io/fyne/v2/dialog" */
	"fyne.io/fyne/v2/layout"
	/* 	"fyne.io/fyne/v2/widget" */

	/* "enigma/anna" */
	/* "enigma/menu" */
	/* "enigma/window" */
	/* 	"enigma/apps" */
	"enigma/consts"
	"enigma/elem"
)

func (c *Cryptor) Start() {
	consts.ChangeLang(c.isRus)
	c.setPlaceHolder()

	c.mainwindow = c.App.NewWindow(consts.NAME_WINDOW_MAIN)

	containerProgressbar := container.NewVBox(c.GetProgressBar())
	c.GetProgressBar().Hide()

	c.content = c.createContainers(containerProgressbar)
	c.menu = CreateMenu(c)
	c.mainwindow.SetMainMenu(c.menu)
	c.mainwindow.SetContent(c.content)

	c.mainwindow.Resize(fyne.NewSize(consts.WINDOW_WEIGHT, consts.WINDOW_HEIGHT))

	c.mainwindow.Show()
	c.mainwindow.SetMaster()

}

func (c *Cryptor) createContainers(pb *fyne.Container) *fyne.Container {
	containerTextField := container.New(layout.NewMaxLayout(), c.textField)

	c.buttonOpenFile = elem.NewButton(consts.BUTTON_OPEN_FILE, c.openFile)

	containerBtnOpen := container.NewWithoutLayout(c.buttonOpenFile)
	c.buttonOpenFile.Resize(fyne.NewSize(200, 40))
	c.buttonOpenFile.Move(fyne.NewPos(20, 20))

	containerWithOpenButton := container.NewHBox(
		containerBtnOpen,
	)

	/* c.button = elem.NewButton(consts.BUTTON_TEXT, c.startCrypt)
	containerBtn := container.NewWithoutLayout(c.button)
	c.button.Resize(fyne.NewSize(250, 40))
	c.button.Move(fyne.NewPos(0, 0)) */

	containerWithKeyAndButtonStart := container.NewVBox(
		container.NewHBox(
			c.GetKeyWordWithSize(200, 40),
			layout.NewSpacer(),
			/* containerBtn, */
			c.GetColorButtonWithSize(200, 40),
		),
		containerWithOpenButton,
	)

	containerWithAllButton := container.NewGridWithRows(2,
		pb,
		containerWithKeyAndButtonStart,
	)

	containerFull := container.NewGridWithRows(2,
		containerTextField,
		containerWithAllButton)

	return containerFull
}

func (c *Cryptor) setPlaceHolder() {
	c.textField.SetPlaceHolder(consts.LABEL_RULES)
	c.textField.Wrapping = fyne.TextWrapWord

	c.keyWord.SetPlaceHolder(consts.KEY_WORD_TITLE)
}
