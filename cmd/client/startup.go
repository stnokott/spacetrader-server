package main

import (
	"context"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"google.golang.org/grpc/status"
)

// ShowStartupSplash creates and displays a splashscreen while the application initializes.
//
// Once startup has completed, the splash window will close and the next window will be opened.
func ShowStartupSplash(app fyne.App, next fyne.Window, worker *Worker) {
	w := newSplashWindow(app)

	// display placeholder while attempting to connect
	card := widget.NewCard(
		"SpaceTrader "+_version,
		"Initializing...",
		container.NewVBox(
			widget.NewProgressBarInfinite(),
			layout.NewSpacer(),
		),
	)

	w.SetContent(card)
	w.Show()

	// define function to be called when error occurs
	// since there is not only one possible error, we allow specifying an error along with its type.
	onError := func(typ string, err error) {
		label := widget.NewLabel(
			status.Convert(err).Message(),
		)
		label.Importance = widget.WarningImportance
		label.Wrapping = fyne.TextWrapBreak
		card := widget.NewCard(
			"SpaceTrader "+_version,
			typ+" error occured:",
			label,
		)
		box := container.NewVBox(
			card,
			widget.NewButton(
				"Close",
				w.Close,
			),
		)
		// hijack current splash window, replace content with error display
		w.SetContent(box)
		// close app when window is closed
		w.SetOnClosed(app.Quit)
		w.Resize(box.MinSize())
		w.CenterOnScreen()
	}

	// start goroutine which attempts to connect to both app and game servers.
	// only when both connections are successful, we continue with the main application.
	go func() {
		// for some unknown reason, we need to give the Fyne app a little time to initialize.
		// if we dont do this, the splash window is not properly closed.
		time.Sleep(1 * time.Second)
		if err := worker.CheckAppServer(context.TODO()); err != nil {
			onError("App connection", err)
			return
		}
		if err := worker.UpdateServerInfo(context.TODO()); err != nil {
			onError("Game connection", err)
			return
		}
		if err := worker.UpdateCurrentAgent(context.TODO()); err != nil {
			onError("Querying agent info", err)
			return
		}
		w.Close()
		next.Show()
	}()
}
