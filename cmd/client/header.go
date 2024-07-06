package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/stnokott/spacetrader/cmd/client/widgets"
	pb "github.com/stnokott/spacetrader/internal/proto"
)

// HeaderWidget represents the header section of the application.
type HeaderWidget struct {
	widget.BaseWidget

	root *fyne.Container
}

// HeaderWidgetBindings contains all bindings for HeaderWidget.
type HeaderWidgetBindings struct {
	ServerStatus *TypedBinding[*pb.ServerStatus]
	AgentInfo    *TypedBinding[*pb.Agent]
}

// NewHeaderWidget creates a new widget to be displayed in the header, containing
// metadata about the current agent.
func NewHeaderWidget(bindings HeaderWidgetBindings) *HeaderWidget {
	gameVersion := widget.NewLabel("n/a")
	gameVersion.TextStyle.Bold = true
	nextReset := widgets.NewUpdatingLabel(
		time.Time{},
		1*time.Second,
		formatNextReset,
	)
	gameDetails := container.NewVBox(
		gameVersion,
		nextReset,
	)
	bindings.ServerStatus.AddListener(func(data *pb.ServerStatus) {
		gameVersion.SetText("Game API " + data.Version)
		nextReset.SetValue(data.NextReset.AsTime())
	})

	agentName := widget.NewLabel("n/a")
	agentName.TextStyle.Bold = true
	agentName.Alignment = fyne.TextAlignTrailing
	agentCredits := canvas.NewText("n/a", _colorCredits)
	agentCredits.Alignment = fyne.TextAlignTrailing
	agentDetails := container.NewVBox(
		agentName,
		container.NewHBox(
			layout.NewSpacer(),
			agentCredits,
			canvas.NewText("₡", _colorCredits),
		),
	)
	bindings.AgentInfo.AddListener(func(data *pb.Agent) {
		agentName.SetText(data.Name)
		agentCredits.Text = fmtInt(int(data.Credits))
		agentCredits.Refresh()
	})

	box := container.NewHBox(
		gameDetails,
		layout.NewSpacer(),
		agentDetails,
	)
	h := &HeaderWidget{
		root: box,
	}
	h.ExtendBaseWidget(h)
	return h
}

// CreateRenderer is required for our custom widget.
func (h *HeaderWidget) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(h.root)
}

func formatNextReset(label *widget.Label, nextReset time.Time) {
	d := time.Until(nextReset)
	label.Text = "Reset in " + fmtDuration(d)
	if d < (1 * time.Hour) {
		label.Importance = widget.HighImportance
		label.TextStyle.Bold = true
	} else {
		label.Importance = widget.MediumImportance
		label.TextStyle.Bold = false
	}
	label.Refresh()
}
