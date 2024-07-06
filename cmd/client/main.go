// Package main provides the frontend application.
package main

import (
	"context"

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

	serverInfoBinding := NewTypedBinding[*pb.ServerStatus]()
	agentInfoBinding := NewTypedBinding[*pb.Agent]()
	headerBindings := HeaderWidgetBindings{
		ServerStatus: serverInfoBinding,
		AgentInfo:    agentInfoBinding,
	}

	header := NewHeaderWidget(headerBindings)
	footer := NewFooterWidget()

	loadingOverlay := NewLoadingOverlay()
	root := container.NewStack(
		container.NewBorder(header, footer, nil, nil),
		loadingOverlay,
	)

	workerBindings := WorkerBindings{
		ServerInfo: serverInfoBinding,
		AgentInfo:  agentInfoBinding,
	}
	worker := NewWorker("localhost:55555", workerBindings) // TODO: from config
	defer worker.Close()

	w.SetContent(root)
	w.Resize(fyne.NewSize(800, 600))
	w.CenterOnScreen()
	w.SetMaster()
	w.Show()

	loadingOverlay.AddJobs(
		Job{"connecting to app server", func() error {
			return worker.CheckAppServer(context.TODO())
		}},
		Job{"connecting to game server", func() error {
			return worker.UpdateServerInfo(context.TODO())
		}},
		Job{"retrieving agent info", func() error {
			return worker.UpdateCurrentAgent(context.TODO())
		}},
	)
	a.Run()
}
