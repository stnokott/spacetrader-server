// Package main provides the frontend application.
package main

import (
	"context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/stnokott/spacetrader/cmd/client/theme"
	"github.com/stnokott/spacetrader/cmd/client/widgets"
	pb "github.com/stnokott/spacetrader/internal/proto"
)

const _version = "v0.0.0" // TODO: from build

var currentTheme = &theme.Theme{}

func main() {
	a := app.NewWithID("github.com/stnokott/spacetrader")
	a.Settings().SetTheme(currentTheme)

	w := a.NewWindow("SpaceTrader " + _version)

	serverBinding := widgets.NewTypedBinding[*pb.ServerStatus]()
	agentBinding := widgets.NewTypedBinding[*pb.Agent]()
	headerBindings := HeaderWidgetBindings{
		ServerStatus: serverBinding,
		AgentInfo:    agentBinding,
	}

	header := NewHeaderWidget(headerBindings)
	footer := NewFooterWidget()

	fleetBinding := widgets.NewTypedBinding[*pb.Fleet]()
	shipListBindings := widgets.ShipListBindings{
		Fleet: fleetBinding,
	}
	left := widgets.NewShipList(shipListBindings)

	split1 := container.NewHSplit(
		left, widget.NewLabel("Center"),
	)
	split1.SetOffset(0.3)
	split2 := container.NewHSplit(
		split1, widget.NewLabel("Right"),
	)
	split2.SetOffset(0.8)
	mainLayout := container.NewBorder(header, footer, nil, nil, split2)

	loadingOverlay := NewLoadingOverlay()
	root := container.NewStack(
		mainLayout,
		loadingOverlay,
	)

	workerBindings := WorkerBindings{
		Server: serverBinding,
		Agent:  agentBinding,
		Fleet:  fleetBinding,
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
		Job{"retrieving fleet", func() error {
			return worker.UpdateFleet(context.TODO())
		}},
	)
	a.Run()
}
