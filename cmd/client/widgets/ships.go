package widgets

import (
	"sync"

	"fyne.io/fyne/v2"
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
			// defer since we don't update the underyling slice untilt the end of this function
			defer l.list.RefreshItem(i)
		} else {
			ships[i] = l.items[i]
		}
	}
	l.m.Lock()
	defer l.m.Unlock()
	l.items = ships
}

func (l *ShipList) length() int {
	l.m.RLock()
	defer l.m.RUnlock()
	return len(l.items)
}

func (l *ShipList) createItem() fyne.CanvasObject {
	return widget.NewLabel("n/a")
}

func (l *ShipList) updateItem(i widget.ListItemID, o fyne.CanvasObject) {
	l.m.RLock()
	o.(*widget.Label).SetText(l.items[i].Name)
	l.m.RUnlock()
}
