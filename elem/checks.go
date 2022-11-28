package elem

import (
	"fyne.io/fyne/v2/widget"
)

var Checks *widget.CheckGroup

func CreateChecks() {
	var checkss = widget.NewCheckGroup([]string{"Check 1", "Check 2", "Check 3"}, func(s []string) {})
	Checks = checkss
}
