package mocks

import (
	"github.com/jinzhu/copier"
	pb "github.com/stnokott/spacetrader-server/internal/proto"
)

func newDefaultShip() *pb.Ship {
	ship := new(pb.Ship)
	if err := copier.CopyWithOption(ship, defaultShip, copier.Option{DeepCopy: true}); err != nil {
		// only called during testing, so panicking is ok
		panic(err)
	}
	return ship
}

func NewShipInSystem(name string, sys *pb.System) *pb.Ship {
	ship := newDefaultShip()
	ship.Name = name
	ship.CurrentLocation.System = sys.Id
	wp := sys.Waypoints[0]
	ship.CurrentLocation.Waypoint = wp.Id
	ship.Route.Origin = wp
	ship.Route.Destination = wp
	return ship
}

var defaultShip = &pb.Ship{
	Id:              "STNOKOTT-1",
	Name:            "STNOKOTT-1",
	Role:            pb.Ship_COMMAND,
	CurrentLocation: &pb.Ship_Location{}, // should be filled when copying
	Route:           &pb.Ship_Route{},    // should be filled when copying
	Status:          pb.Ship_DOCKED,
}
