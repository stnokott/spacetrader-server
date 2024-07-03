package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"

	pb "github.com/stnokott/spacetrader/internal/proto"
)

// Server displays the current server connection status.
// It also offers a popup with additional server information while connected.
type Server struct {
	widget.DisableableWidget

	bg     *canvas.Rectangle
	status *widget.Label

	serverInfoDialog *dialog.CustomDialog
	gameVersion      *widget.Label
	agents           *widget.Label
	ships            *widget.Label
	waypoints        *widget.Label
	systems          *widget.Label

	root *fyne.Container
}

// NewServerWidget constructs a new widget which displays the current server connection status
// and some additional metadata about it in a popup if data is populated.
func NewServerWidget(data *TypedBinding[*pb.ServerStatusReply], parent fyne.Window) *Server {
	bg := canvas.NewRectangle(color.Transparent)
	status := widget.NewLabel("unknown")
	stack := container.NewStack(
		bg,
		status,
	)

	gameVersion := widget.NewLabel("n/a")
	agents := widget.NewLabel("n/a")
	ships := widget.NewLabel("n/a")
	waypoints := widget.NewLabel("n/a")
	systems := widget.NewLabel("n/a")

	data.AddListener(func(data *pb.ServerStatusReply) {
		gameVersion.SetText(data.Version)
		agents.SetText(fmtInt(int(data.GlobalStats.Agents)))
		ships.SetText(fmtInt(int(data.GlobalStats.Ships)))
		waypoints.SetText(fmtInt(int(data.GlobalStats.Waypoints)))
		systems.SetText(fmtInt(int(data.GlobalStats.Systems)))
	})

	serverInfo := dialog.NewCustom(
		"Server information",
		"Close",
		container.NewGridWithColumns(
			2,
			widget.NewLabel("Version:"), gameVersion,
			widget.NewLabel("Agents:"), agents,
			widget.NewLabel("Ships:"), ships,
			widget.NewLabel("Waypoints:"), waypoints,
			widget.NewLabel("Systems:"), systems,
		),
		parent,
	)

	s := &Server{
		bg:               bg,
		status:           status,
		serverInfoDialog: serverInfo,
		gameVersion:      gameVersion,
		agents:           agents,
		ships:            ships,
		waypoints:        waypoints,
		systems:          systems,
		root:             stack,
	}
	s.ExtendBaseWidget(s)

	s.SetConnected(false)
	return s
}

// CreateRenderer is required for our custom widget.
func (s *Server) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(s.root)
}

// Tapped implements the fyne.Tappable interface.
func (s *Server) Tapped(_ *fyne.PointEvent) {
	if s.Disabled() {
		return
	}
	s.serverInfoDialog.Show()
}

// Cursor implements the desktop.Cursorable interface.
func (s *Server) Cursor() desktop.Cursor {
	if !s.Disabled() {
		return desktop.PointerCursor
	}
	return desktop.DefaultCursor
}

// SetConnected defines how the widget looks and whether the button for the
// popup with additional information is displayed (only when connected).
func (s *Server) SetConnected(c bool) {
	if c {
		s.bg.FillColor = _colorSuccess
		s.status.SetText("Connected")
		s.Enable()
	} else {
		s.bg.FillColor = _colorError
		s.status.SetText("Disconnected")
		s.serverInfoDialog.Hide()
		s.Disable()
	}
}
