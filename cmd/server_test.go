package main

import (
	"testing"

	pb "github.com/stnokott/spacetrader-server/internal/proto"
	"github.com/stretchr/testify/assert"
)

func TestServerShipsPerSystem(t *testing.T) {
	server := &Server{
		fleetCache: FleetCache{
			[]*pb.Ship{
				{CurrentLocation: &pb.Ship_Location{System: "Foo"}},
				{CurrentLocation: &pb.Ship_Location{System: "Bar"}},
				{CurrentLocation: &pb.Ship_Location{System: "Foo"}},
				{CurrentLocation: &pb.Ship_Location{System: "Fuzz"}},
			},
		},
	}

	assert := assert.New(t)
	got := server.shipsPerSystem()

	assert.Equal(3, len(got))
	assert.Equal(2, got["Foo"])
	assert.Equal(1, got["Bar"])
	assert.Equal(1, got["Fuzz"])
}
