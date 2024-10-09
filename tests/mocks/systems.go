package mocks

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/stnokott/spacetrader-server/internal/api"
)

// GenerateSystems creates an artifial list of systems with a specified length.
// The coordinates of each system are randomly generated between coordMin and coordMax.
func GenerateSystems(count int, coordMin int, coordMax int) []*api.System {
	systems := make([]*api.System, count)
	for i := range systems {
		systemName := GenerateSystemName(count, i)
		systems[i] = &api.System{
			Symbol:   systemName,
			X:        coordMin + rand.Intn(coordMax-coordMin),
			Y:        coordMin + rand.Intn(coordMax-coordMin),
			Type:     api.SystemTypeBLACKHOLE,
			Factions: []api.SystemFaction{},
			Waypoints: []api.SystemWaypoint{
				{
					Symbol: systemName + "-WP",
					X:      0,
					Y:      0,
				},
			},
		}
		if rand.Float64() > 0.75 {
			systems[i].Factions = []api.SystemFaction{
				{Symbol: api.FactionSymbolASTRO},
			}
		}
	}
	return systems
}

// GenerateSystemName generates a system name with letters 0-9, with added padding for
// the total number of systems expected.
func GenerateSystemName(total int, i int) string {
	fmtString := fmt.Sprintf("SYSTEM-%%0%dd", len(strconv.Itoa(total)))
	return fmt.Sprintf(fmtString, i)
}
