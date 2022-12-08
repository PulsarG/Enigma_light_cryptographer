package crypt

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

func (c *Cryptor) Start() /* fyne.Window */ {
	c.textField.SetPlaceHolder(consts.LABEL_RULES)
	c.keyWord.SetPlaceHolder(consts.KEY_WORD_TITLE)
	c.textField.Wrapping = fyne.TextWrapWord

	window := c.App.NewWindow(consts.NAME_WINDOW_MAIN)
	c.mainwindow = window

	/* mainMenu := menu.CreateMenu() */

	containerProgressbar := container.NewVBox(c.GetProgressBar())
	c.GetProgressBar().Hide()

	/* window.SetMainMenu(mainMenu) */
	window.SetContent(c.createContainers(containerProgressbar))
	window.Resize(fyne.NewSize(consts.WINDOW_WEIGHT, consts.WINDOW_HEIGHT))

	window.Show()
	window.SetMaster()

}

func (c *Cryptor) createContainers(pb *fyne.Container) *fyne.Container {
	containerTextField := container.New(layout.NewMaxLayout(), c.textField)

	containerWithKeyAndButtonStart := container.NewVBox(
		container.NewHBox(
			c.GetKeyWordWithSize(200, 40),
			layout.NewSpacer(),
			c.GetColorButtonWithSize(250, 40),
		),
	)

	btnOpen := elem.NewButton("Открыть файл и расшифровать", c.openFile)

	containerBtnOpen := container.NewWithoutLayout(btnOpen)
	btnOpen.Resize(fyne.NewSize(300, 40))
	btnOpen.Move(fyne.NewPos(-30, -20))

	containerWithOpenButton := container.NewGridWithColumns(2,
		layout.NewSpacer(),
		containerBtnOpen,
	)

	containerWithAllButton := container.NewGridWithRows(4,
		pb,
		containerWithKeyAndButtonStart,
		layout.NewSpacer(),
		containerWithOpenButton,
	)

	containerFull := container.NewGridWithRows(2,
		containerTextField,
		/* pb, */
		containerWithAllButton)

	return containerFull
}
