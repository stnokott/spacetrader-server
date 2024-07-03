package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"

	pb "github.com/stnokott/spacetrader/internal/proto"
)

// ServerWidget displays the current server connection status.
// It also offers a popup with additional server information while connected.
type ServerWidget struct {
	widget.BaseWidget

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

// ServerWidgetBindings holds all bindings which are relevant for the ServerWidget.
type ServerWidgetBindings struct {
	Info *TypedBinding[*pb.ServerStatusReply]
}

// NewServerWidget constructs a new widget which displays the current server connection status
// and some additional metadata about it in a popup if data is populated.
func NewServerWidget(bindings ServerWidgetBindings, parent fyne.Window) *ServerWidget {
	bg := canvas.NewRectangle(_colorSuccess)
	status := widget.NewLabel("Connected")
	stack := container.NewStack(
		bg,
		status,
	)

	newNameLabel := func(name string) *widget.Label {
		return widget.NewLabelWithStyle(name, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	}
	newValueLabel := func() *widget.Label {
		return widget.NewLabelWithStyle("n/a", fyne.TextAlignTrailing, fyne.TextStyle{})
	}
	gameVersion := newValueLabel()
	agents := newValueLabel()
	ships := newValueLabel()
	waypoints := newValueLabel()
	systems := newValueLabel()

	serverInfo := dialog.NewCustom(
		"Server information",
		"Close",
		container.NewGridWithColumns(
			2,
			newNameLabel("Version:"), gameVersion,
			newNameLabel("Agents:"), agents,
			newNameLabel("Ships:"), ships,
			newNameLabel("Waypoints:"), waypoints,
			newNameLabel("Systems:"), systems,
		),
		parent,
	)

	w := &ServerWidget{
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
	w.ExtendBaseWidget(w)

	bindings.Info.AddListener(func(data *pb.ServerStatusReply) {
		gameVersion.SetText(data.Version)
		agents.SetText(fmtInt(int(data.GlobalStats.Agents)))
		ships.SetText(fmtInt(int(data.GlobalStats.Ships)))
		waypoints.SetText(fmtInt(int(data.GlobalStats.Waypoints)))
		systems.SetText(fmtInt(int(data.GlobalStats.Systems)))
	})
	return w
}

// CreateRenderer is required for our custom widget.
func (w *ServerWidget) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(w.root)
}

// Tapped implements the fyne.Tappable interface.
func (w *ServerWidget) Tapped(_ *fyne.PointEvent) {
	w.serverInfoDialog.Show()
}

// Cursor implements the desktop.Cursorable interface.
func (w *ServerWidget) Cursor() desktop.Cursor {
	return desktop.PointerCursor
}
