package main

import (
	"context"
	"fmt"

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
		"Connecting...",
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
			fmt.Sprintf("%s error: %s", typ, status.Convert(err).Message()),
		)
		label.Importance = widget.WarningImportance
		label.Wrapping = fyne.TextWrapBreak
		card := widget.NewCard(
			"SpaceTrader "+_version,
			"Connection error:",
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
		if err := worker.CheckAppServer(context.TODO()); err != nil {
			onError("App connection", err)
			return
		}
		if err := worker.CheckGameServer(context.TODO()); err != nil {
			onError("Game connection", err)
		}
		w.Close()
		next.Show()
	}()
}
