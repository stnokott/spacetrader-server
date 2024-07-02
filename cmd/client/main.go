// Package main provides the frontend application.
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const version = "v0.0.0" // TODO: from build

func main() {
	a := app.NewWithID("github.com/stnokott/spacetrader")
	a.Settings().SetTheme(&Theme{})

	w := a.NewWindow("SpaceTrader " + version)
	w.Resize(fyne.NewSize(800, 600))

	header := newHeader()

	version := widget.NewLabel(version)
	version.Alignment = fyne.TextAlignTrailing
	version.Importance = widget.LowImportance

	statusLabel := widget.NewLabel("Status:")
	statusLabel.Alignment = fyne.TextAlignTrailing
	statusLabel.Importance = widget.LowImportance

	status := widget.NewLabel("Idle")
	statusLabel.Alignment = fyne.TextAlignLeading

	footer := container.NewHBox(
		statusLabel,
		status,
		layout.NewSpacer(),
		version,
	)
	root := container.NewBorder(header, footer, nil, nil)

	w.SetContent(root)
	w.CenterOnScreen()
	w.SetMaster()
	w.ShowAndRun()
}
