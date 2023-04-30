package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("GUI application")

	w.SetContent(widget.NewLabel("Hello world!"))
	w.ShowAndRun()
}
