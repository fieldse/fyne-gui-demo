package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
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
