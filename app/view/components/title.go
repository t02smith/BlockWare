package components

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func Title() fyne.CanvasObject {
	appTitle := canvas.NewText("BlockWare", color.White)
	appTitle.TextSize = 18

	author := canvas.NewText("Tom Smith", color.Opaque)
	author.TextSize = 18

	return container.New(layout.NewHBoxLayout(), appTitle, layout.NewSpacer(), author)
}
