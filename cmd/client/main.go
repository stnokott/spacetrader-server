// Package main provides the frontend application.
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

const _version = "v0.0.0" // TODO: from build

var currentTheme = &Theme{}

func main() {
	a := app.NewWithID("github.com/stnokott/spacetrader")
	a.Settings().SetTheme(currentTheme)

	w := a.NewWindow("SpaceTrader " + _version)
	w.Resize(fyne.NewSize(800, 600))

	header := NewHeaderWidget()
	footer := NewFooterWidget(w)

	root := container.NewBorder(header, footer, nil, nil)

	w.SetContent(root)
	w.CenterOnScreen()
	w.SetMaster()
	w.ShowAndRun()
}
