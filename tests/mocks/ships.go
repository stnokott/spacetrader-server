package mocks

import (
	"time"

	"github.com/jinzhu/copier"
	pb "github.com/stnokott/spacetrader-server/internal/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	ship.CurrentLocation.System = sys
	wp := sys.Waypoints[0]
	ship.CurrentLocation.Waypoint = wp
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
	FlightMode:      pb.Ship_CRUISE,
	Crew: &pb.Ship_Crew{
		Current:  57,
		Capacity: 80,
		Required: 57,
		Rotation: pb.Ship_Crew_STRICT,
		Morale:   100,
		Wages:    0,
	},
	Fuel: &pb.Ship_Fuel{
		Current:  400,
		Capacity: 400,
		Consumed: &pb.Ship_Fuel_Consumption{
			Amount:    0,
			Timestamp: timestamppb.New(time.Date(2024, 6, 30, 22, 14, 48, 907, time.UTC)),
		},
	},
	Cooldown: &pb.Ship_Cooldown{
		TotalSeconds:     0,
		RemainingSeconds: 0,
	},
	Frame: &pb.FrameComponent{
		Type:        pb.FrameComponent_FRAME_FRIGATE,
		Name:        "Frigate",
		Description: "A medium-sized, multi-purpose spacecraft, often used for combat, transport, or support operations.",
		Requirements: &pb.ModuleRequirements{
			Power: ptr(int32(8)),
			Crew:  ptr(int32(25)),
		},
		Degradation: &pb.ModuleDegradable{
			Condition: 1,
			Integrity: 1,
		},
		ModuleSlots:    8,
		MountingPoints: 5,
		FuelCapacity:   400,
	},
	Reactor: &pb.ReactorComponent{
		Type:        pb.ReactorComponent_REACTOR_FISSION_I,
		Name:        "Fission Reactor I",
		Description: "A basic fission power reactor, used to generate electricity from nuclear fission reactions.",
		Requirements: &pb.ModuleRequirements{
			Crew: ptr(int32(8)),
		},
		Degradation: &pb.ModuleDegradable{
			Condition: 1,
			Integrity: 1,
		},
		PowerOutput: 31,
	},
	Engine: &pb.EngineComponent{
		Type:        pb.EngineComponent_ENGINE_ION_DRIVE_II,
		Name:        "Ion Drive II",
		Description: "An advanced propulsion system that uses ionized particles to generate high-speed, low-thrust acceleration, with improved efficiency and performance.",
		Requirements: &pb.ModuleRequirements{
			Power: ptr(int32(6)),
			Crew:  ptr(int32(8)),
		},
		Degradation: &pb.ModuleDegradable{
			Condition: 1,
			Integrity: 1,
		},
		Speed: 30,
	},
	Modules: []*pb.Module{
		{
			Type:        pb.Module_MODULE_CARGO_HOLD_II,
			Name:        "Expanded Cargo Hold",
			Description: "An expanded cargo hold module that provides more efficient storage space for a ship's cargo.",
			Requirements: &pb.ModuleRequirements{
				Crew:  ptr(int32(2)),
				Power: ptr(int32(2)),
				Slots: ptr(int32(2)),
			},
			Capacity: ptr(int32(40)),
		},
		{
			Type:        pb.Module_MODULE_CREW_QUARTERS_I,
			Name:        "Crew Quarters",
			Description: "A module that provides living space and amenities for the crew.",
			Requirements: &pb.ModuleRequirements{
				Crew:  ptr(int32(2)),
				Power: ptr(int32(1)),
				Slots: ptr(int32(1)),
			},
			Capacity: ptr(int32(40)),
		},
		{
			Type:        pb.Module_MODULE_CREW_QUARTERS_I,
			Name:        "Crew Quarters",
			Description: "A module that provides living space and amenities for the crew.",
			Requirements: &pb.ModuleRequirements{
				Crew:  ptr(int32(2)),
				Power: ptr(int32(1)),
				Slots: ptr(int32(1)),
			},
			Capacity: ptr(int32(40)),
		},
		{
			Type:        pb.Module_MODULE_MINERAL_PROCESSOR_I,
			Name:        "Mineral Processor",
			Description: "Crushes and processes extracted minerals and ores into their component parts, filters out impurities, and containerizes them into raw storage units.",
			Requirements: &pb.ModuleRequirements{
				Crew:  ptr(int32(0)),
				Power: ptr(int32(1)),
				Slots: ptr(int32(2)),
			},
		},
		{
			Type:        pb.Module_MODULE_GAS_PROCESSOR_I,
			Name:        "Gas Processor",
			Description: "Filters and processes extracted gases into their component parts, filters out impurities, and containerizes them into raw storage units.",
			Requirements: &pb.ModuleRequirements{
				Crew:  ptr(int32(0)),
				Power: ptr(int32(1)),
				Slots: ptr(int32(2)),
			},
		},
	},
	Mounts: []*pb.Mount{
		{
			Type:        pb.Mount_MOUNT_SENSOR_ARRAY_II,
			Name:        "Sensor Array II",
			Description: ptr("An advanced sensor array that improves a ship's ability to detect and track other objects in space with greater accuracy and range."),
			Requirements: &pb.ModuleRequirements{
				Crew:  ptr(int32(2)),
				Power: ptr(int32(2)),
			},
			Strength: ptr(int32(4)),
		},
		{
			Type:        pb.Mount_MOUNT_GAS_SIPHON_II,
			Name:        "Gas Siphon II",
			Description: ptr("An advanced gas siphon that can extract gas from gas giants and other gas-rich bodies more efficiently and at a higher rate."),
			Requirements: &pb.ModuleRequirements{
				Crew:  ptr(int32(2)),
				Power: ptr(int32(2)),
			},
			Strength: ptr(int32(20)),
		},
		{
			Type:        pb.Mount_MOUNT_MINING_LASER_II,
			Name:        "Mining Laser II",
			Description: ptr("An advanced mining laser that is more efficient and effective at extracting valuable minerals from asteroids and other space objects."),
			Requirements: &pb.ModuleRequirements{
				Crew:  ptr(int32(2)),
				Power: ptr(int32(2)),
			},
			Strength: ptr(int32(5)),
		},
		{
			Type:        pb.Mount_MOUNT_SURVEYOR_II,
			Name:        "Surveyor II",
			Description: ptr("An advanced survey probe that can be used to gather information about a mineral deposit with greater accuracy."),
			Requirements: &pb.ModuleRequirements{
				Crew:  ptr(int32(4)),
				Power: ptr(int32(3)),
			},
			Strength: ptr(int32(2)),
			Deposits: []pb.TradeItem{
				pb.TradeItem_QUARTZ_SAND,
				pb.TradeItem_SILICON_CRYSTALS,
				pb.TradeItem_PRECIOUS_STONES,
				pb.TradeItem_ICE_WATER,
				pb.TradeItem_AMMONIA_ICE,
				pb.TradeItem_IRON_ORE,
				pb.TradeItem_COPPER_ORE,
				pb.TradeItem_SILVER_ORE,
				pb.TradeItem_ALUMINUM_ORE,
				pb.TradeItem_GOLD_ORE,
				pb.TradeItem_PLATINUM_ORE,
				pb.TradeItem_DIAMONDS,
				pb.TradeItem_URANITE_ORE,
			},
		},
	},
	Cargo: &pb.Ship_Cargo{
		Capacity:  40,
		Units:     0,
		Inventory: []*pb.Ship_Cargo_InventoryItem{},
	},
}
