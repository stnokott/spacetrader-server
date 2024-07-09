package widgets

import (
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	pb "github.com/stnokott/spacetrader/internal/proto"
)

// ShipList displays a list with all the current agent's ships along
// with their status and other helpful information.
type ShipList struct {
	widget.BaseWidget

	list  *widget.List
	items []*shipListShipModel
	m     sync.RWMutex
}

type shipListShipModel struct {
	Name string
}

func (m *shipListShipModel) Equal(m2 *shipListShipModel) bool {
	return m.Name == m2.Name
}

// ShipListBindings holds all bindings required for a ShipList instance.
type ShipListBindings struct {
	Fleet *TypedBinding[*pb.Fleet]
}

// NewShipList creates a new ShipList instance.
func NewShipList(bindings ShipListBindings) *ShipList {
	l := &ShipList{}
	l.ExtendBaseWidget(l)

	list := widget.NewList(
		l.length,
		l.createItem,
		l.updateItem,
	)
	l.list = list

	bindings.Fleet.AddListener(l.onFleetUpdate)
	return l
}

// CreateRenderer creates a new renderer for l.
func (l *ShipList) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(l.list)
}

func (l *ShipList) onFleetUpdate(data *pb.Fleet) {
	ships := make([]*shipListShipModel, len(data.Ships))
	for i, ship := range data.Ships {
		model := &shipListShipModel{
			Name: ship.Name,
		}
		// check if item changed
		l.m.RLock()
		isNewItem := i >= len(l.items)
		isUpdatedItem := !isNewItem && !model.Equal(l.items[i])
		l.m.RUnlock()
		if isNewItem || isUpdatedItem {
			ships[i] = model
		} else {
			ships[i] = l.items[i]
		}
	}
	l.m.Lock()
	l.items = ships
	l.m.Unlock()
	l.list.Refresh()
}
func (l *ShipList) length() int {
	l.m.RLock()
	defer l.m.RUnlock()
	return len(l.items)
}

func (l *ShipList) createItem() fyne.CanvasObject {
	return NewShipListItem()
}

func (l *ShipList) updateItem(i widget.ListItemID, o fyne.CanvasObject) {
	li := o.(*ShipListItem)
	l.m.RLock()
	li.SetShipName(l.items[i].Name)
	l.m.RUnlock()
}

// ShipListItem displays information about a single ship.
type ShipListItem struct {
	widget.BaseWidget

	icon  *widget.Icon
	label *widget.Label
}

// NewShipListItem creates a new ship list item with placeholder values.
func NewShipListItem() *ShipListItem {
	icon := widget.NewIcon(theme.ErrorIcon())
	label := widget.NewLabel("n/a")

	li := &ShipListItem{
		icon:  icon,
		label: label,
	}
	li.ExtendBaseWidget(li)
	return li
}

// CreateRenderer creates a new renderer for li.
func (li *ShipListItem) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(container.NewHBox(
		li.icon,
		li.label,
	))
}

// SetShipName sets the displayed name of the ship.
func (li *ShipListItem) SetShipName(s string) {
	li.label.SetText(s)
}
