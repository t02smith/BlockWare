package pages

import (
	"fmt"
	"image/color"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/t02smith/part-iii-project/toolkit/model/games"
	model "github.com/t02smith/part-iii-project/toolkit/model/net"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

func UploadPage() fyne.CanvasObject {

	// * GAME INPUT FIELDS
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

	// * GAME CREATION PROGRESS
	progressBar := widget.NewProgressBar()
	count, countText, total, totalText, status, statusText := 0, widget.NewLabel("0"), 0, widget.NewLabel("0"), "waiting...", widget.NewLabel("waiting...")

	progressContainer := container.NewHBox(countText, canvas.NewText("/", color.White), totalText, layout.NewSpacer(), statusText)

	countBind, totalBind, statusBind := binding.BindInt(&count), binding.BindInt(&total), binding.BindString(&status)
	countBind.AddListener(binding.NewDataListener(func() {
		countText.SetText(fmt.Sprint(count))
	}))
	totalBind.AddListener(binding.NewDataListener(func() {
		totalText.SetText(fmt.Sprint(total))
	}))
	statusBind.AddListener(binding.NewDataListener(func() {
		statusText.SetText(status)
	}))

	// * INPUT FORM
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
			// * disable inputs
			for _, x := range []*widget.Entry{title, version, domain, day, month, year, shardSize} {
				x.Disable()
				defer x.Enable()
				defer x.SetText("")
			}

			rootDirBtn.Disable()
			defer rootDirBtn.Enable()

			// ? parse inputs
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

			// * setup progress bar
			progress := make(chan int)
			go func() {
				totalBind.Set(<-progress)
				count = 0

				statusText.SetText("Sharding directory...")
				for count < total {
					<-progress
					countBind.Set(count + 1)

					progressBar.SetValue(float64(count) / float64(total))
				}
				statusText.SetText("Outputting to file...")

			}()

			// * create game
			p := model.GetPeerInstance()
			g, err := games.CreateGame(title.Text, version.Text, date, domain.Text, rootDir, uint(ss), progress)
			if err != nil {
				return
			}

			p.GetLibrary().AddGame(g)
			games.OutputToFile(g)
			statusText.SetText("Complete")

			libBinding.Set(true)
		},
	}

	content := widget.NewCard(
		"Upload a new game",
		"Fill in the details below and hit upload to share your game",
		container.NewVBox(form, progressContainer, progressBar),
	)

	return content
}
