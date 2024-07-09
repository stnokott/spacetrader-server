package widgets

import (
	"testing"

	"github.com/stretchr/testify/assert"

	pb "github.com/stnokott/spacetrader/internal/proto"
	"github.com/stnokott/spacetrader/tests/mocks"
)

func TestShipsList(t *testing.T) {
	fleetBinding := NewTypedBinding[*pb.Fleet]()
	shipListBindings := ShipListBindings{
		Fleet: fleetBinding,
	}

	shipList := NewShipList(shipListBindings)

	expectLength := func(n int) {
		assert.Len(t, shipList.items, n)
		assert.Equal(t, shipList.list.Length(), n)
	}
	expectItemTextAt := func(i int, text string) {
		lbl := NewShipListItem()
		shipList.list.UpdateItem(i, lbl)
		assert.Equal(t, text, lbl.GetShipName())
	}

	// items should be empty when first created
	expectLength(0)

	// add one item
	ship1 := mocks.NewDefaultShip()
	ship1.Name = "Foo Ship"
	shipListBindings.Fleet.Set(&pb.Fleet{
		Ships: []*pb.Ship{ship1},
	})
	// length should be 1
	expectLength(1)
	expectItemTextAt(0, "Foo Ship")

	// add one more item
	ship2 := mocks.NewDefaultShip()
	ship2.Name = "Bar Ship"
	shipListBindings.Fleet.Set(&pb.Fleet{
		Ships: []*pb.Ship{ship1, ship2},
	})
	// length should be 2
	expectLength(2)
	expectItemTextAt(0, "Foo Ship")
	expectItemTextAt(1, "Bar Ship")

	// reverse order
	shipListBindings.Fleet.Set(&pb.Fleet{
		Ships: []*pb.Ship{ship2, ship1},
	})
	// nothing should change
	expectLength(2)
	expectItemTextAt(0, "Bar Ship")
	expectItemTextAt(1, "Foo Ship")

	// remove first item
	shipListBindings.Fleet.Set(&pb.Fleet{
		Ships: []*pb.Ship{ship1},
	})
	// length should be 1
	expectLength(1)
	expectItemTextAt(0, "Foo Ship")

	// clear items
	shipListBindings.Fleet.Set(&pb.Fleet{
		Ships: []*pb.Ship{},
	})
	// items should be empty once again
	expectLength(0)
}
