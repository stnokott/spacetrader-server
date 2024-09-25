package mocks

import (
	"github.com/jinzhu/copier"
	"github.com/stnokott/spacetrader-server/internal/api"
	"github.com/stnokott/spacetrader-server/internal/graph/model"
)

func newDefaultShip() *model.Ship {
	ship := new(model.Ship)
	if err := copier.CopyWithOption(ship, defaultShip, copier.Option{DeepCopy: true}); err != nil {
		// only called during testing, so panicking is ok
		panic(err)
	}
	return ship
}

func NewShipInSystem(name string, sys *model.System) *model.Ship {
	ship := newDefaultShip()
	ship.Name = name
	ship.System = sys
	ship.SystemID = sys.Name
	wp := sys.Waypoints[0]
	ship.Waypoint = wp
	ship.WaypointID = wp.Name
	return ship
}

var defaultShip = &model.Ship{
	Name:       "STNOKOTT-1",
	Role:       api.ShipRoleCOMMAND,
	System:     nil, // should be filled when copying
	Waypoint:   nil, // should be filled when copying
	SystemID:   "",  // should be filled when copying
	WaypointID: "",  // should be filled when copying
}
