package mocks

import (
	"fmt"
	"math/rand"
	"strconv"

	pb "github.com/stnokott/spacetrader-server/internal/proto"
)

// GenerateSystems creates an artifial list of systems with a specified length.
// The coordinates of each system are randomly generated between coordMin and coordMax.
func GenerateSystems(count int, coordMin int, coordMax int) []*pb.System {
	systems := make([]*pb.System, count)
	for i := range systems {
		systems[i] = &pb.System{
			Id:       GenerateSystemName(count, i),
			X:        int32(coordMin + rand.Intn(coordMax-coordMin)),
			Y:        int32(coordMin + rand.Intn(coordMax-coordMin)),
			Type:     pb.System_Type(int(pb.System_UNKNOWN_SYSTEMTYPE) + rand.Intn(int(pb.System_UNSTABLE))),
			Factions: []pb.Faction{},
		}
		if rand.Float64() > 0.75 {
			systems[i].Factions = []pb.Faction{
				pb.Faction(int(pb.Faction_UNKNOWN_FACTION) + rand.Intn(int(pb.Faction_ETHEREAL))),
			}
		}
	}
	return systems
}

// GenerateSystemName generates a system name with letters 0-9, with added padding for
// the total number of systems expected.
func GenerateSystemName(total int, i int) string {
	fmtString := fmt.Sprintf("%%0%dd", len(strconv.Itoa(total)))
	return fmt.Sprintf(fmtString, i)
}
