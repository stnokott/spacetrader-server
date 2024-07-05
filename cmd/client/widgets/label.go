package widgets

import (
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// UpdatingLabel refreshes its displayed text at specific intervals by calling a user-defined function on
// an underlying value.
//
// This can be useful for displaying text which depends on an external value that is only indirectly
// part of the text, e.g. the current time.
type UpdatingLabel[T any] struct {
	widget.BaseWidget
	label *widget.Label

	v T
	m sync.RWMutex
}

// NewUpdatingLabel creates a new instance of UpdatingLabel.
//
// At every interval, setFunc will be called, allowing you to update the label itself to your liking.
// The binding is required to notify the label of changes in the underlying value.
func NewUpdatingLabel[T any](
	value T,
	interval time.Duration,
	updateFunc func(*widget.Label, T),
) *UpdatingLabel[T] {
	l := &UpdatingLabel[T]{
		label: widget.NewLabel(""),
		v:     value,
	}
	l.ExtendBaseWidget(l)

	go func() {
		for range time.Tick(interval) {
			l.m.RLock()
			updateFunc(l.label, l.v)
			l.m.RUnlock()
		}
	}()
	return l
}

// CreateRenderer creates a renderer for UpdatingLabel.
func (l *UpdatingLabel[T]) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(l.label)
}

// SetValue updates the underlying value that is passed to the update function every interval.
func (l *UpdatingLabel[T]) SetValue(v T) {
	l.m.Lock()
	l.v = v
	l.m.Unlock()
}
