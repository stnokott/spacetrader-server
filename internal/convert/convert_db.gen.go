// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.
//go:build !goverter

package convert

import (
	query "github.com/stnokott/spacetrader/internal/db/query"
	proto "github.com/stnokott/spacetrader/internal/proto"
)

func ConvertSystem(source *query.System) (*proto.System, error) {
	var pProtoSystem *proto.System
	if source != nil {
		var protoSystem proto.System
		protoSystem.Id = (*source).Symbol
		protoSystem_Type, err := ParseSystemType((*source).Type)
		if err != nil {
			return nil, err
		}
		protoSystem.Type = protoSystem_Type
		xint32, err := Int64ToInt32((*source).X)
		if err != nil {
			return nil, err
		}
		protoSystem.X = xint32
		xint322, err := Int64ToInt32((*source).Y)
		if err != nil {
			return nil, err
		}
		protoSystem.Y = xint322
		protoFactionList, err := ParseDBFactions((*source).Factions)
		if err != nil {
			return nil, err
		}
		protoSystem.Factions = protoFactionList
		pProtoSystem = &protoSystem
	}
	return pProtoSystem, nil
}
