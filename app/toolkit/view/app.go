package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/t02smith/part-iii-project/toolkit/util"
	"github.com/t02smith/part-iii-project/toolkit/view/components"
	"github.com/t02smith/part-iii-project/toolkit/view/pages"
)

const ()

func StartApp() {
	app := app.New()
	window := setupWindow(app)

	util.Logger.Info("Starting GUI")
	window.ShowAndRun()
}

func setupWindow(app fyne.App) fyne.Window {
	util.Logger.Info("Setting up GUI")
	window := app.NewWindow(WINDOW_TITLE)

	window.Resize(fyne.NewSize(WINDOW_WIDTH, WINDOW_HEIGHT))
	window.SetMaster()
	window.CenterOnScreen()

	title := components.Title()
	nav := components.Navbar(
		&components.NavbarTab{
			Name:    "Home",
			Content: pages.HomePage(),
		},
		&components.NavbarTab{
			Name:    "Games",
			Content: pages.GamesPage(),
		},
		&components.NavbarTab{
			Name:    "Upload",
			Content: pages.UploadPage(),
		},
		&components.NavbarTab{
			Name:    "Peers",
			Content: pages.PeersPage(),
		},
	)

	window.SetContent(container.NewBorder(title, nil, nil, nil, nav))
	return window
}
