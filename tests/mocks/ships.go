package mocks

import (
	"github.com/jinzhu/copier"
	"github.com/stnokott/spacetrader-server/internal/api"
)

func newDefaultShip() *api.Ship {
	ship := new(api.Ship)
	if err := copier.CopyWithOption(ship, defaultShip, copier.Option{DeepCopy: true}); err != nil {
		// only called during testing, so panicking is ok
		panic(err)
	}
	return ship
}

func NewShipInSystem(name string, sys *api.System) *api.Ship {
	ship := newDefaultShip()
	ship.Symbol = name
	ship.Nav.SystemSymbol = sys.Symbol
	wp := sys.Waypoints[0]
	ship.Nav.WaypointSymbol = wp.Symbol
	return ship
}

var defaultShip = &api.Ship{
	Symbol: "STNOKOTT-1",
	Registration: api.ShipRegistration{
		Role: api.ShipRoleCOMMAND,
	},
	Nav: api.ShipNav{
		SystemSymbol:   "", // should be filled when copying
		WaypointSymbol: "", // should be filled when copying
		Status:         api.DOCKED,
	},
}
