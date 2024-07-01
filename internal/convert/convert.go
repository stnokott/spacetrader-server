// Package convert converts between API responses and protobuf structs.
package convert

import (
	"github.com/stnokott/spacetrader/internal/api"

	pb "github.com/stnokott/spacetrader/internal/proto"
)

//go:generate go run github.com/jmattheis/goverter/cmd/goverter@v1.5.0 gen ./...

// Converter converts between API responses and protobuf structs.
// goverter:converter
// goverter:output:file ./convert.gen.go
// goverter:output:package github.com/stnokott/spacetrader/internal/convert
// goverter:output:format function
// goverter:ignoreUnexported yes
type Converter interface {
	// goverter:map LastReset.Time LastReset | google.golang.org/protobuf/types/known/timestamppb:New
	// goverter:map Resets.Next NextReset | google.golang.org/protobuf/types/known/timestamppb:New
	// goverter:map Statistics GlobalStats
	ConvertStatus(source *api.Status) *pb.ServerStatusReply
}
