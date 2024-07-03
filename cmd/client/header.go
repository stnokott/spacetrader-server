package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// Header represents the header section of the application.
type Header struct {
	widget.BaseWidget

	agentName    *widget.Label
	agentCredits *canvas.Text
	root         *fyne.Container
}

// NewHeaderWidget creates a new widget to be displayed in the header, containing
// metadata about the current agent.
// TODO: binding (see server.go)
func NewHeaderWidget() *Header {
	title := widget.NewLabel("SpaceTrader")
	title.TextStyle.Bold = true

	agentName := widget.NewLabel("AGENTNAME")
	agentName.TextStyle.Bold = true

	agentCredits := canvas.NewText("0", _colorCredits)
	agentCredits.Alignment = fyne.TextAlignTrailing
	agentCreditsBox := container.NewHBox(
		layout.NewSpacer(),
		agentCredits,
		canvas.NewText("₡", _colorCredits),
	)
	agentDetails := container.NewVBox(
		agentName,
		agentCreditsBox,
	)

	box := container.NewHBox(
		title,
		layout.NewSpacer(),
		agentDetails,
	)
	h := &Header{
		agentName:    agentName,
		agentCredits: agentCredits,
		root:         box,
	}
	h.ExtendBaseWidget(h)
	return h
}

// CreateRenderer is required for our custom widget.
func (h *Header) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(h.root)
}

// SetAgentName sets the displayed agent name.
func (h *Header) SetAgentName(s string) {
	h.agentName.SetText(s)
}

// SetCredits sets the displayed agent credits.
func (h *Header) SetCredits(n int) {
	h.agentCredits.Text = fmtInt(n)
	h.agentCredits.Refresh()
}
