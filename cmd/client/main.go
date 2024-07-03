// Package main provides the frontend application.
package main

import (
	"context"
	"log"

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
	w.Resize(fyne.NewSize(800, 600))

	serverInfoBinding := NewTypedBinding[*pb.ServerStatusReply]()

	footerBindings := FooterWidgetBindings{
		Server: ServerWidgetBindings{
			Info: serverInfoBinding,
		},
	}

	header := NewHeaderWidget()
	footer := NewFooterWidget(footerBindings, w)

	root := container.NewBorder(header, footer, nil, nil)

	workerBindings := WorkerBindings{
		ServerInfo: serverInfoBinding,
	}
	worker := NewWorker("localhost:55555", workerBindings)
	defer worker.Close()

	// TODO: splash screen during connection
	if !worker.CheckAppServer(context.TODO()) {
		// TODO: popup
		log.Println("app server not available")
		return
	}
	if !worker.CheckGameServer(context.TODO()) {
		// TODO: popup
		log.Println("game server not available")
		return
	}

	w.SetContent(root)
	w.CenterOnScreen()
	w.SetMaster()
	w.ShowAndRun()
}
