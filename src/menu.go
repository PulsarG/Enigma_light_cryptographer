package src

import (
	/* "fmt" */

	"fyne.io/fyne/v2"
	/* "fyne.io/fyne/v2/widget" */

	"enigma/consts"
)

func CreateMenu(c *Cryptor) *fyne.MainMenu {
	btnText := consts.MENU_SUBMENU_ONE_BTN_ONE_TITLE

	buttonSubmenuOne := fyne.NewMenuItem(btnText, func() {
		/* c.SetLang()
		consts.ChangeLang(c.GetLang())

		// ?????????????????????

		c.button.SetText(consts.BUTTON_TEXT)
		c.buttonOpenFile.SetText(consts.BUTTON_OPEN_FILE)
		c.keyWord.SetPlaceHolder(consts.KEY_WORD_TITLE)
		c.textField.SetPlaceHolder(consts.LABEL_RULES) */
	})

	buttonSubmenuOne.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem(consts.MENU_SUBMENU_ONE_BTN_ENGL, func() {
			consts.ChangeLang(false)

			c.button.SetText(consts.BUTTON_TEXT)
			c.buttonOpenFile.SetText(consts.BUTTON_OPEN_FILE)
			c.keyWord.SetPlaceHolder(consts.KEY_WORD_TITLE)
			c.textField.SetPlaceHolder(consts.LABEL_RULES)
		}),

		fyne.NewMenuItem(consts.MENU_SUBMENU_ONE_BTN_RUSSIAN, func() {
			consts.ChangeLang(true)

			c.button.SetText(consts.BUTTON_TEXT)
			c.buttonOpenFile.SetText(consts.BUTTON_OPEN_FILE)
			c.keyWord.SetPlaceHolder(consts.KEY_WORD_TITLE)
			c.textField.SetPlaceHolder(consts.LABEL_RULES)
		}))

	subemenuOne := fyne.NewMenu(consts.TITLE_SUBMENU_ONE, buttonSubmenuOne)
	mainMenu := fyne.NewMainMenu(subemenuOne)
	return mainMenu
}
