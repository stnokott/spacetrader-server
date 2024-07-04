// Package main provides the frontend application.
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	pb "github.com/stnokott/spacetrader/internal/proto"
)

const _version = "v0.0.0" // TODO: from build

var currentTheme = &Theme{}

func main() {
	a := app.NewWithID("github.com/stnokott/spacetrader")
	a.Settings().SetTheme(currentTheme)

	w := a.NewWindow("SpaceTrader " + _version)

	serverInfoBinding := NewTypedBinding[*pb.ServerStatusReply]()
	headerBindings := HeaderWidgetBindings{
		ServerStatus: serverInfoBinding,
	}

	header := NewHeaderWidget(headerBindings)
	footer := NewFooterWidget()

	root := container.NewBorder(header, footer, nil, nil)

	workerBindings := WorkerBindings{
		ServerInfo: serverInfoBinding,
	}
	worker := NewWorker("localhost:55555", workerBindings) // TODO: from config
	defer worker.Close()

	w.SetContent(root)
	w.Resize(fyne.NewSize(800, 600))
	w.CenterOnScreen()
	w.SetMaster()

	ShowStartupSplash(a, w, worker)
	a.Run()
}
