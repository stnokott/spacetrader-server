package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

func newSplashWindow(app fyne.App) fyne.Window {
	drv, ok := app.Driver().(desktop.Driver)
	if !ok {
		panic("needs to be compiled for desktop")
	}
	return drv.CreateSplashWindow()
}
