package components

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func Clock() fyne.CanvasObject {
	clock := widget.NewLabel("")

	formatted := time.Now().Format("Time : 03:04:05")
	clock.SetText(formatted)

	go func() {
		for range time.Tick(time.Second) {
			formatted = time.Now().Format("Time : 03:04:05")
			clock.SetText(formatted)
		}
	}()

	return clock

}
