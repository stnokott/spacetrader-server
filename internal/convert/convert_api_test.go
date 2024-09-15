package convert

import (
	"testing"
	"time"

	"github.com/jinzhu/copier"
	"github.com/stnokott/spacetrader-server/internal/api"
	pb "github.com/stnokott/spacetrader-server/internal/proto"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ptr[T any](x T) *T {
	return &x
}

var defaultAPIShip = api.Ship{
	Symbol: "STNOKOTT-1",
	Nav: api.ShipNav{
		SystemSymbol:   "X1-MB64",
		WaypointSymbol: "X1-MB64-A1",
		Route: api.ShipNavRoute{
			Origin: api.ShipNavRouteWaypoint{
				Symbol:       "X1-MB64-A1",
				SystemSymbol: "X1-MB64",
				Type:         api.PLANET,
				X:            22,
				Y:            -13,
			},
			Destination: api.ShipNavRouteWaypoint{
				Symbol:       "X1-MB64-A1",
				SystemSymbol: "X1-MB64",
				Type:         api.PLANET,
				X:            22,
				Y:            -13,
			},
			Arrival:       time.Date(2024, 6, 30, 22, 14, 48, 907, time.UTC),
			DepartureTime: time.Date(2024, 6, 30, 22, 14, 48, 907, time.UTC),
		},
		Status:     api.DOCKED,
		FlightMode: api.CRUISE,
	},
	Crew: api.ShipCrew{
		Current:  57,
		Capacity: 80,
		Required: 57,
		Rotation: api.STRICT,
		Morale:   100,
		Wages:    0,
	},
	Fuel: api.ShipFuel{
		Current:  400,
		Capacity: 400,
		Consumed: &struct {
			Amount    int       `json:"amount"`
			Timestamp time.Time `json:"timestamp"`
		}{
			Amount:    0,
			Timestamp: time.Date(2024, 6, 30, 22, 14, 48, 907, time.UTC),
		},
	},
	Cooldown: api.Cooldown{
		ShipSymbol:       "STNOKOTT-1",
		TotalSeconds:     0,
		RemainingSeconds: 0,
	},
	Frame: api.ShipFrame{
		Symbol:         api.FRAMEFRIGATE,
		Name:           "Frigate",
		Description:    "A medium-sized, multi-purpose spacecraft, often used for combat, transport, or support operations.",
		ModuleSlots:    8,
		MountingPoints: 5,
		FuelCapacity:   400,
		Condition:      1,
		Integrity:      1,
		Requirements: api.ShipRequirements{
			Power: ptr(8),
			Crew:  ptr(25),
		},
	},
	Reactor: api.ShipReactor{
		Symbol:      api.REACTORFISSIONI,
		Name:        "Fission Reactor I",
		Description: "A basic fission power reactor, used to generate electricity from nuclear fission reactions.",
		Condition:   1,
		Integrity:   1,
		PowerOutput: 31,
		Requirements: api.ShipRequirements{
			Crew: ptr(8),
		},
	},
	Engine: api.ShipEngine{
		Symbol:      api.ShipEngineSymbolENGINEIONDRIVEII,
		Name:        "Ion Drive II",
		Description: "An advanced propulsion system that uses ionized particles to generate high-speed, low-thrust acceleration, with improved efficiency and performance.",
		Condition:   1,
		Integrity:   1,
		Speed:       30,
		Requirements: api.ShipRequirements{
			Power: ptr(6),
			Crew:  ptr(8),
		},
	},
	Modules: []api.ShipModule{
		{
			Symbol:      api.MODULECARGOHOLDII,
			Name:        "Expanded Cargo Hold",
			Description: "An expanded cargo hold module that provides more efficient storage space for a ship's cargo.",
			Capacity:    ptr(40),
			Requirements: api.ShipRequirements{
				Crew:  ptr(2),
				Power: ptr(2),
				Slots: ptr(2),
			},
		},
		{
			Symbol:      api.MODULECREWQUARTERSI,
			Name:        "Crew Quarters",
			Description: "A module that provides living space and amenities for the crew.",
			Capacity:    ptr(40),
			Requirements: api.ShipRequirements{
				Crew:  ptr(2),
				Power: ptr(1),
				Slots: ptr(1),
			},
		},
		{
			Symbol:      api.MODULECREWQUARTERSI,
			Name:        "Crew Quarters",
			Description: "A module that provides living space and amenities for the crew.",
			Capacity:    ptr(40),
			Requirements: api.ShipRequirements{
				Crew:  ptr(2),
				Power: ptr(1),
				Slots: ptr(1),
			},
		},
		{
			Symbol:      api.MODULEMINERALPROCESSORI,
			Name:        "Mineral Processor",
			Description: "Crushes and processes extracted minerals and ores into their component parts, filters out impurities, and containerizes them into raw storage units.",
			Requirements: api.ShipRequirements{
				Crew:  ptr(0),
				Power: ptr(1),
				Slots: ptr(2),
			},
		},
		{
			Symbol:      api.MODULEGASPROCESSORI,
			Name:        "Gas Processor",
			Description: "Filters and processes extracted gases into their component parts, filters out impurities, and containerizes them into raw storage units.",
			Requirements: api.ShipRequirements{
				Crew:  ptr(0),
				Power: ptr(1),
				Slots: ptr(2),
			},
		},
	},
	Mounts: []api.ShipMount{
		{
			Symbol:      api.MOUNTSENSORARRAYII,
			Name:        "Sensor Array II",
			Description: ptr("An advanced sensor array that improves a ship's ability to detect and track other objects in space with greater accuracy and range."),
			Strength:    ptr(4),
			Requirements: api.ShipRequirements{
				Crew:  ptr(2),
				Power: ptr(2),
			},
		},
		{
			Symbol:      api.MOUNTGASSIPHONII,
			Name:        "Gas Siphon II",
			Description: ptr("An advanced gas siphon that can extract gas from gas giants and other gas-rich bodies more efficiently and at a higher rate."),
			Strength:    ptr(20),
			Requirements: api.ShipRequirements{
				Crew:  ptr(2),
				Power: ptr(2),
			},
		},
		{
			Symbol:      api.MOUNTMININGLASERII,
			Name:        "Mining Laser II",
			Description: ptr("An advanced mining laser that is more efficient and effective at extracting valuable minerals from asteroids and other space objects."),
			Strength:    ptr(5),
			Requirements: api.ShipRequirements{
				Crew:  ptr(2),
				Power: ptr(2),
			},
		},
		{
			Symbol:      api.MOUNTSURVEYORII,
			Name:        "Surveyor II",
			Description: ptr("An advanced survey probe that can be used to gather information about a mineral deposit with greater accuracy."),
			Strength:    ptr(2),
			Deposits: &[]api.ShipMountDeposits{
				api.QUARTZSAND,
				api.SILICONCRYSTALS,
				api.PRECIOUSSTONES,
				api.ICEWATER,
				api.AMMONIAICE,
				api.IRONORE,
				api.COPPERORE,
				api.SILVERORE,
				api.ALUMINUMORE,
				api.GOLDORE,
				api.PLATINUMORE,
				api.DIAMONDS,
				api.URANITEORE,
			},
			Requirements: api.ShipRequirements{
				Crew:  ptr(4),
				Power: ptr(3),
			},
		},
	},
	Registration: api.ShipRegistration{
		Name:          "STNOKOTT-1",
		FactionSymbol: "COSMIC",
		Role:          "COMMAND",
	},
	Cargo: api.ShipCargo{
		Capacity:  40,
		Inventory: []api.ShipCargoItem{},
		Units:     0,
	},
}

var defaultPbShip = pb.Ship{
	Id:   "STNOKOTT-1",
	Name: "STNOKOTT-1",
	Role: pb.Ship_COMMAND,
	CurrentLocation: &pb.Ship_Location{
		System:   "X1-MB64",
		Waypoint: "X1-MB64-A1",
	},
	Route: &pb.Ship_Route{
		Origin: &pb.WaypointBase{
			Id:     "X1-MB64-A1",
			System: "X1-MB64",
			Type:   pb.WaypointBase_PLANET,
			X:      22,
			Y:      -13,
		},
		Destination: &pb.WaypointBase{
			Id:     "X1-MB64-A1",
			System: "X1-MB64",
			Type:   pb.WaypointBase_PLANET,
			X:      22,
			Y:      -13,
		},
		ArrivalTime:   timestamppb.New(time.Date(2024, 6, 30, 22, 14, 48, 907, time.UTC)),
		DepartureTime: timestamppb.New(time.Date(2024, 6, 30, 22, 14, 48, 907, time.UTC)),
	},
	Status: pb.Ship_DOCKED,
}

func TestConvertShip(t *testing.T) {
	type modifierIn func(*api.Ship)
	type modifierExpected func(*pb.Ship)
	tests := []struct {
		name          string
		modifIn       modifierIn
		modifExpected modifierExpected
		wantErr       bool
	}{
		{
			name:          "default",
			modifIn:       nil,
			modifExpected: nil,
			wantErr:       false,
		},
		{
			name: "invalid ship status",
			modifIn: func(s *api.Ship) {
				s.Nav.Status = "DRIFTING"
			},
			modifExpected: nil,
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			source := new(api.Ship)
			if err := copier.CopyWithOption(source, &defaultAPIShip, copier.Option{DeepCopy: true}); err != nil {
				t.Fatalf("failed to copy: %v", err)
			}
			if tt.modifIn != nil {
				tt.modifIn(source)
			}
			expected := new(pb.Ship)
			if err := copier.CopyWithOption(expected, &defaultPbShip, copier.Option{DeepCopy: true}); err != nil {
				t.Fatalf("failed to copy: %v", err)
			}
			if tt.modifExpected != nil {
				tt.modifExpected(expected)
			}

			got, err := ConvertShip(source)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertShip() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if !proto.Equal(got, expected) {
				t.Log("ConvertShip() mismatch")
				t.Logf("got  = %#v", got)
				t.Logf("want = %#v", expected)
				t.Fail()
			}
		})
	}
}
