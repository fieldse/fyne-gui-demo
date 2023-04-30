package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("GUI application")

	l := widget.NewLabel("Hello world!")

	w.SetContent(
		container.NewVBox(
			l,
			widget.NewButton("Hi!", func() {
				l.SetText("Welcome :)")
			}),
		))
	w.ShowAndRun()
}
