package menu

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"

	"test/consts"
)

func CreateMenu(text *widget.Entry, label *widget.Label) *fyne.MainMenu {

	buttonSubmenuOne := fyne.NewMenuItem(consts.MENU_SUBMENU_ONE_BTN_ONE_TITLE, func() {
		label.SetText(text.Text)
	})
	buttonSubmenuTwo := createButtonsSubMenu(consts.MENU_SUBMENU_ONE_BTN_TWO_TITLE)
	buttonSubmenuTree := fyne.NewMenuItem(consts.MENU_SUBMENU_ONE_BTN_TREE_TITLE, nil)
	buttonSubmenuTree.ChildMenu = fyne.NewMenu(
		"",
		fyne.NewMenuItem("text1", nil),
		fyne.NewMenuItem("Text2", nil),
	)

	subemenuOne := fyne.NewMenu(consts.TITLE_SUBMENU_ONE, buttonSubmenuOne, buttonSubmenuTwo, buttonSubmenuTree)

	mainMenu := fyne.NewMainMenu(subemenuOne)

	return mainMenu
}

func createButtonsSubMenu(title string) *fyne.MenuItem {
	buttonSubmenu := fyne.NewMenuItem(title, func() {
		fmt.Println(title)
	})
	return buttonSubmenu
}
