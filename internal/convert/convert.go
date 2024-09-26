// Package convert converts between structs from different sources (DB, API, Protobuf, ...)
package convert

import (
	"strings"
	"time"

	"github.com/stnokott/spacetrader-server/internal/api"
	"github.com/stnokott/spacetrader-server/internal/db/query"
	"github.com/stnokott/spacetrader-server/internal/graph/model"
)

// Converter converts between API responses and protobuf structs.
// goverter:converter
// goverter:output:package github.com/stnokott/spacetrader-server/internal/convert
// goverter:output:file ./convert.gen.go
// goverter:output:format function
// goverter:ignoreUnexported yes
// goverter:enum:unknown @error
// goverter:extend IntTo.*
// goverter:extend Int64To.*
// goverter:extend Parse.*
// goverter:extend TimeToTime
type Converter interface {
	// goverter:map LastReset.Time LastReset
	// goverter:map Resets.Next NextReset
	// goverter:map Statistics Stats
	ConvertServerStatus(source *api.Status) *model.Server

	// goverter:map Symbol Name
	// goverter:map Headquarters Hq
	ConvertAgent(source *api.Agent) *model.Agent

	// goverter:map Registration.Name Name
	// goverter:map Registration.Role Role
	// goverter:map Nav.Status Status
	// goverter:map Nav.SystemSymbol SystemID
	// goverter:map Nav.WaypointSymbol WaypointID
	// goverter:ignore System
	// goverter:ignore Waypoint
	ConvertShip(source *api.Ship) (*model.Ship, error)
	ConvertShips(source []*api.Ship) ([]*model.Ship, error)

	// goverter:map Symbol Name
	// goverter:ignore Waypoints
	// goverter:map Factions Factions | ParseFactions
	// goverter:ignore HasJumpgates
	ConvertSystem(source query.System) *model.System
	ConvertSystems(source []query.System) []*model.System

	// goverter:map Symbol Name
	// goverter:map System SystemID
	// goverter:ignore System
	// goverter:ignore ConnectedTo
	ConvertWaypoint(source query.Waypoint) *model.Waypoint
	ConvertWaypoints(source []query.Waypoint) []*model.Waypoint
}

// ParseFactions converts the concatenated factions from DB into a string slice.
func ParseFactions(concat string) []api.FactionSymbol {
	split := strings.Split(concat, ",")
	if len(split) == 1 && split[0] == "" {
		return []api.FactionSymbol{}
	}
	out := make([]api.FactionSymbol, len(split))
	for i, x := range split {
		out[i] = api.FactionSymbol(x)
	}
	return out
}

// TimeToTime is simple, but required since time.Time contains unexported fields which goverter
// cannot convert.
func TimeToTime(source time.Time) time.Time {
	return source
}
