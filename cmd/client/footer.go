package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// FooterWidget contains metadata about the application like the current version.
type FooterWidget struct {
	widget.BaseWidget

	root *fyne.Container
}

// FooterWidgetBindings holds all bindings relevant for FooterWidget.
type FooterWidgetBindings struct {
	Server ServerWidgetBindings
}

// NewFooterWidget constructs a new footer, containing the current version
// of the application and metadata about the current server connection.
func NewFooterWidget(bindings FooterWidgetBindings, parent fyne.Window) *FooterWidget {
	server := NewServerWidget(bindings.Server, parent)

	version := widget.NewLabel(_version)
	version.Alignment = fyne.TextAlignTrailing
	version.Importance = widget.LowImportance

	box := container.NewHBox(
		server,
		layout.NewSpacer(),
		version,
	)

	f := &FooterWidget{
		root: box,
	}
	f.ExtendBaseWidget(f)
	return f
}

// CreateRenderer is required for our custom widget.
func (w *FooterWidget) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(w.root)
}
