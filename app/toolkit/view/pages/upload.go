package pages

import (
	"image/color"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/t02smith/part-iii-project/toolkit/model/games"
	model "github.com/t02smith/part-iii-project/toolkit/model/net"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

func UploadPage() fyne.CanvasObject {

	title := widget.NewEntry()
	version := widget.NewEntry()
	domain := widget.NewEntry()
	shardSize := widget.NewEntry()

	rootDir := ""
	rootDirBtn := widget.NewButton("Choose directory", func() {
		dialog := dialog.NewFolderOpen(
			func(lu fyne.ListableURI, err error) {
				rootDir = lu.Path()
			},
			fyne.CurrentApp().Driver().AllWindows()[0],
		)

		dialog.Show()
	})

	day := widget.NewEntry()
	month := widget.NewEntry()
	year := widget.NewEntry()

	date := container.NewHBox(day, canvas.NewText("/", color.White), month, canvas.NewText("/", color.White), year)

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Game Title", Widget: title},
			{Text: "Version", Widget: version},
			{Text: "Your domain", Widget: domain},
			{Text: "Release date", Widget: date},
			{Text: "Game root directory", Widget: rootDirBtn},
			{Text: "Shard size", Widget: shardSize},
		},
		OnSubmit: func() {

			ss, err := strconv.ParseUint(shardSize.Text, 10, 32)
			if err != nil {
				util.Logger.Errorf("Invalid shard size %s", err)
				return
			}

			y, err := strconv.ParseInt(year.Text, 10, 32)
			if err != nil {
				util.Logger.Errorf("Invalid year %s", err)
			}

			m, err := strconv.ParseInt(month.Text, 10, 32)
			if err != nil {
				util.Logger.Errorf("Invalid month %s", err)
			}

			d, err := strconv.ParseInt(day.Text, 10, 32)
			if err != nil {
				util.Logger.Errorf("Invalid day %s", err)
			}

			date := time.Date(int(y), time.Month(m), int(d), 0, 0, 0, 0, time.UTC).String()

			p := model.GetPeerInstance()
			g, err := games.CreateGame(title.Text, version.Text, date, domain.Text, rootDir, uint(ss))
			if err != nil {
				return
			}

			p.GetLibrary().AddGame(g)
			games.OutputToFile(g)

			libBinding.Set(true)
		},
	}

	content := widget.NewCard(
		"Upload a new game",
		"Fill in the details below and hit upload to share your game",
		form,
	)

	return content
}
