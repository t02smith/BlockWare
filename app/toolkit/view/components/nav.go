package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type NavbarTab struct {
	Name    string
	Content fyne.CanvasObject
}

func Navbar(tabs ...*NavbarTab) *container.AppTabs {
	tabItems := []*container.TabItem{}
	for _, t := range tabs {
		tabItems = append(tabItems, container.NewTabItem(t.Name, t.Content))
	}

	nav := container.NewAppTabs(tabItems...)
	nav.SetTabLocation(container.TabLocationLeading)

	return nav

}
