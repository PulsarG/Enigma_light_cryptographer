package menu

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"

	"test/consts"
)

func CreateMenu(textForLabel *widget.Entry, label *widget.Label) *fyne.MainMenu {

	buttonSubmenuOne := fyne.NewMenuItem(consts.MENU_SUBMENU_ONE_BTN_ONE_TITLE, func() {
		label.SetText(textForLabel.Text)
	})
	buttonSubmenuTwo := createButtonsSubMenu(consts.MENU_SUBMENU_ONE_BTN_TWO_TITLE)
	buttonSubmenuTwo2 := createButtonsSubMenu("test menu")
	buttonSubmenuTree := fyne.NewMenuItem(consts.MENU_SUBMENU_ONE_BTN_TREE_TITLE, nil)
	buttonSubmenuTree.ChildMenu = fyne.NewMenu(
		"",
		fyne.NewMenuItem(consts.MENU_SUBMENU_ONEBTN_SETTINGS_TITLE, nil),
		fyne.NewMenuItem("Text2", nil),
	)

	subemenuOne := fyne.NewMenu(consts.TITLE_SUBMENU_ONE, buttonSubmenuOne, buttonSubmenuTwo, buttonSubmenuTwo2, buttonSubmenuTree)

	mainMenu := fyne.NewMainMenu(subemenuOne)

	return mainMenu
}

func createButtonsSubMenu(title string) *fyne.MenuItem {
	buttonSubmenu := fyne.NewMenuItem(title, func() {
		fmt.Println(title)
	})
	return buttonSubmenu
}
