package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	pb "github.com/stnokott/spacetrader/internal/proto"
)

// Footer contains metadata about the application like the current version.
type Footer struct {
	widget.BaseWidget

	root *fyne.Container
}

// NewFooterWidget constructs a new footer, containing the current version
// of the application and metadata about the current server connection.
func NewFooterWidget(parent fyne.Window) *Footer {
	serverData := NewTypedBinding[*pb.ServerStatusReply]()
	server := NewServerWidget(serverData, parent)

	version := widget.NewLabel(_version)
	version.Alignment = fyne.TextAlignTrailing
	version.Importance = widget.LowImportance

	box := container.NewHBox(
		server,
		layout.NewSpacer(),
		version,
	)

	f := &Footer{
		root: box,
	}
	f.ExtendBaseWidget(f)
	return f
}

// CreateRenderer is required for our custom widget.
func (f *Footer) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(f.root)
}
