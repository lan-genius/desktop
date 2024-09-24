package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.NewWithID("top.langenius.app")
	w := a.NewWindow("mainWindow")
	w.Resize(fyne.NewSize(200, 200))
	w.SetMaster()

	t := widget.NewLabel("hello中国")
	w.SetContent(t)
	w.ShowAndRun()
}
