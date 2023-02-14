package pages

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	model "github.com/t02smith/part-iii-project/toolkit/model/net"
)

var libBinding binding.ExternalBool
var libUpdated bool = false

func GamesPage() fyne.CanvasObject {
	gs := [][]string{}

	// ? watch for changes in the game
	lib := model.GetPeerInstance().GetLibrary()
	libBinding = binding.BindBool(&libUpdated)
	libBinding.AddListener(binding.NewDataListener(func() {
		if changed, err := libBinding.Get(); !changed || err != nil {
			return
		}

		gamesLs := lib.GetOwnedGames()

		gs = [][]string{}
		for _, g := range gamesLs {
			gs = append(gs, []string{g.Title, g.Version, g.Developer, g.ReleaseDate})
		}

		libBinding.Set(false)
	}))

	//
	games := widget.NewTable(
		func() (int, int) {
			return len(gs), 4
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(tci widget.TableCellID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(gs[tci.Row][tci.Col])
		},
	)

	libBinding.Set(true)
	return widget.NewCard("Your Library", "these are the games that are available to you", games)
}
