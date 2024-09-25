package mocks

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/stnokott/spacetrader-server/internal/api"
	"github.com/stnokott/spacetrader-server/internal/graph/model"
)

// GenerateSystems creates an artifial list of systems with a specified length.
// The coordinates of each system are randomly generated between coordMin and coordMax.
func GenerateSystems(count int, coordMin int, coordMax int) []*model.System {
	systems := make([]*model.System, count)
	for i := range systems {
		systemName := GenerateSystemName(count, i)
		systems[i] = &model.System{
			Name:     systemName,
			X:        coordMin + rand.Intn(coordMax-coordMin),
			Y:        coordMin + rand.Intn(coordMax-coordMin),
			Type:     api.SystemTypeBLACKHOLE,
			Factions: []api.FactionSymbol{},
			Waypoints: []*model.Waypoint{
				{
					Name:   systemName + "-WP",
					System: nil,
					X:      0,
					Y:      0,
				},
			},
		}
		if rand.Float64() > 0.75 {
			systems[i].Factions = []api.FactionSymbol{api.FactionSymbolASTRO}
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
