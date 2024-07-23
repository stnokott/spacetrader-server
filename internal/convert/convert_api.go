package convert

import (
	"fmt"
	"time"

	"github.com/stnokott/spacetrader/internal/api"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/stnokott/spacetrader/internal/proto"
)

// ConverterAPI converts between API responses and protobuf structs.
// goverter:converter
// goverter:output:file ./convert_api.gen.go
// goverter:output:package github.com/stnokott/spacetrader/internal/convert
// goverter:output:format function
// goverter:ignoreUnexported yes
// goverter:extend IntTo.*
// goverter:extend Parse.*
type ConverterAPI interface {
	// goverter:map LastReset.Time LastReset
	// goverter:map Resets.Next NextReset
	// goverter:map Statistics GlobalStats
	ConvertStatus(source *api.Status) *pb.ServerStatus

	// goverter:map Symbol Name
	ConvertAgent(source *api.Agent) (*pb.Agent, error)

	// goverter:map SystemSymbol System
	// goverter:map WaypointSymbol Waypoint
	ConvertShipLocation(source api.ShipNav) *pb.Ship_Location

	// goverter:map Symbol Id
	// goverter:map SystemSymbol System
	ConvertNavWaypoint(source api.ShipNavRouteWaypoint) (*pb.WaypointBase, error)

	// goverter:map Arrival ArrivalTime
	ConvertNavRoute(source api.ShipNavRoute) (*pb.Ship_Route, error)

	// goverter:map Symbol Type
	// goverter:map . Degradation
	ConvertShipFrame(source api.ShipFrame) (*pb.FrameComponent, error)
	// goverter:map Symbol Type
	// goverter:map . Degradation
	ConvertShipReactor(source api.ShipReactor) (*pb.ReactorComponent, error)
	// goverter:map Symbol Type
	// goverter:map . Degradation
	ConvertShipEngine(source api.ShipEngine) (*pb.EngineComponent, error)

	// goverter:map Symbol Type
	ConvertShipMount(source api.ShipMount) (*pb.Mount, error)
	// goverter:map Symbol Type
	ConvertShipModule(source api.ShipModule) (*pb.Module, error)

	// goverter:enum:transform regex ShipRole(\w+) Ship_${1}
	// goverter:enum:unknown @error
	ConvertShipRole(source api.ShipRole) (pb.Ship_Role, error)

	// goverter:enum:map DOCKED Ship_DOCKED
	// goverter:enum:map INORBIT Ship_IN_ORBIT
	// goverter:enum:map INTRANSIT Ship_IN_TRANSIT
	// goverter:enum:unknown @error
	ConvertShipFlightStatus(s api.ShipNavStatus) (pb.Ship_FlightStatus, error)

	// goverter:enum:transform regex (.+) Ship_Crew_${1}
	// goverter:enum:unknown @error
	ConvertShipCrewRotation(s api.ShipCrewRotation) (pb.Ship_Crew_Rotation, error)

	// goverter:useZeroValueOnPointerInconsistency yes
	ConvertShipMountDeposits(source *[]api.ShipMountDeposits) ([]pb.TradeItem, error)

	// goverter:map Symbol Type
	ConvertShipCargoItem(source api.ShipCargoItem) (*pb.Ship_Cargo_InventoryItem, error)

	// goverter:map Symbol Id
	// goverter:map Registration.Name Name
	// goverter:map Registration.Role Role
	// goverter:map Nav CurrentLocation
	// goverter:map Nav.Route Route
	// goverter:map Nav.Status Status
	// goverter:map Nav.FlightMode FlightMode
	ConvertShip(source *api.Ship) (*pb.Ship, error)
}

// ParseTimestamp converts a Go time.Time to a protobuf timestamp.
func ParseTimestamp(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}

// ParseTimestampOptional converts a Go time.Time pointer to a protobuf timestamp.
// If t is nil, nil is returned.
func ParseTimestampOptional(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}
	return timestamppb.New(*t)
}

// ParseFaction parses a faction string into its enum equivalent.
func ParseFaction(s string) (pb.Faction, error) {
	if f, ok := pb.Faction_value[s]; ok {
		return pb.Faction(f), nil
	}
	return pb.Faction_UNKNOWN_FACTION, fmt.Errorf("invalid faction '%s'", s)
}

// ParseShipRoute returns nil if the ship is not in transit and uses
// a regular conversion otherwise.
func ParseShipRoute(c ConverterAPI, source api.ShipNavRoute) (*pb.Ship_Route, error) {
	if source.Origin.SystemSymbol == source.Destination.SystemSymbol &&
		source.Origin.Symbol == source.Destination.Symbol {
		// no route present as ship is not moving (not in transit)
		return nil, nil
	}

	return c.ConvertNavRoute(source)
}

// ParseWaypointType parses a waypoint type string into its enum pb equivalent.
func ParseWaypointType(s api.WaypointType) (pb.WaypointBase_Type, error) {
	if t, ok := pb.WaypointBase_Type_value[string(s)]; ok {
		return pb.WaypointBase_Type(t), nil
	}
	return pb.WaypointBase_UNKNOWN_WAYPOINTTYPE, fmt.Errorf("invalid waypoint type '%s'", s)
}

// ParseShipFlightMode parses a ship flight mode into its enum equivalent.
func ParseShipFlightMode(s api.ShipNavFlightMode) (pb.Ship_FlightMode, error) {
	if mode, ok := pb.Ship_FlightMode_value[string(s)]; ok {
		return pb.Ship_FlightMode(mode), nil
	}
	return pb.Ship_UNKNOWN_FLIGHTMODE, fmt.Errorf("invalid ship flight mode '%s'", s)
}

// ParseShipEngineType parses a ship engine type into its enum equivalent.
func ParseShipEngineType(s api.ShipEngineSymbol) (pb.EngineComponent_Type, error) {
	if t, ok := pb.EngineComponent_Type_value[string(s)]; ok {
		return pb.EngineComponent_Type(t), nil
	}
	return pb.EngineComponent_ENGINE_UNKNOWN, fmt.Errorf("invalid ship engine type '%s'", s)
}

// ParseShipFrameType parses a ship frame type into its enum equivalent.
func ParseShipFrameType(s api.ShipFrameSymbol) (pb.FrameComponent_Type, error) {
	if t, ok := pb.FrameComponent_Type_value[string(s)]; ok {
		return pb.FrameComponent_Type(t), nil
	}
	return pb.FrameComponent_FRAME_UNKNOWN, fmt.Errorf("invalid ship frame type '%s'", s)
}

// ParseShipModuleType parses a ship module type into its enum equivalent.
func ParseShipModuleType(s api.ShipModuleSymbol) (pb.Module_Type, error) {
	if t, ok := pb.Module_Type_value[string(s)]; ok {
		return pb.Module_Type(t), nil
	}
	return pb.Module_MODULE_UNKNOWN, fmt.Errorf("invalid ship module type '%s'", s)
}

// ParseShipMountType parses a ship mount type into its enum equivalent.
func ParseShipMountType(s api.ShipMountSymbol) (pb.Mount_Type, error) {
	if t, ok := pb.Mount_Type_value[string(s)]; ok {
		return pb.Mount_Type(t), nil
	}
	return pb.Mount_MOUNT_UNKNOWN, fmt.Errorf("invalid ship mount type '%s'", s)
}

// ParseShipMountDeposit parses a ship mount deposit into its enum equivalent.
func ParseShipMountDeposit(s api.ShipMountDeposits) (pb.TradeItem, error) {
	if d, ok := pb.TradeItem_value[string(s)]; ok {
		return pb.TradeItem(d), nil
	}
	return pb.TradeItem_UNKNOWN_TRADEITEM, fmt.Errorf("invalid ship mount deposit '%s'", s)
}

// ParseShipReactorType parses a ship reactor type into its enum equivalent.
func ParseShipReactorType(s api.ShipReactorSymbol) (pb.ReactorComponent_Type, error) {
	if t, ok := pb.ReactorComponent_Type_value[string(s)]; ok {
		return pb.ReactorComponent_Type(t), nil
	}
	return pb.ReactorComponent_REACTOR_UNKNOWN, fmt.Errorf("invalid ship reactor type '%s'", s)
}

// ParseTradeItem parses a trade item into its enum equivalent.
func ParseTradeItem(s api.TradeSymbol) (pb.TradeItem, error) {
	if item, ok := pb.TradeItem_value[string(s)]; ok {
		return pb.TradeItem(item), nil
	}
	return pb.TradeItem_UNKNOWN_TRADEITEM, fmt.Errorf("invalid trade item '%s'", s)
}
