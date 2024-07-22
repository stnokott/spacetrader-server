// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.20.1
// source: galaxy.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Faction int32

const (
	Faction_UNKNOWN_FACTION Faction = 0
	Faction_COSMIC          Faction = 1
	Faction_VOID            Faction = 2
	Faction_GALACTIC        Faction = 3
	Faction_QUANTUM         Faction = 4
	Faction_DOMINION        Faction = 5
	Faction_ASTRO           Faction = 6
	Faction_CORSAIRS        Faction = 7
	Faction_OBSIDIAN        Faction = 8
	Faction_AEGIS           Faction = 9
	Faction_UNITED          Faction = 10
	Faction_SOLITARY        Faction = 11
	Faction_COBALD          Faction = 12
	Faction_OMEGA           Faction = 13
	Faction_ECHO            Faction = 14
	Faction_LORDS           Faction = 15
	Faction_CULT            Faction = 16
	Faction_ANCIENTS        Faction = 17
	Faction_SHADOW          Faction = 18
	Faction_ETHEREAL        Faction = 19
)

// Enum value maps for Faction.
var (
	Faction_name = map[int32]string{
		0:  "UNKNOWN_FACTION",
		1:  "COSMIC",
		2:  "VOID",
		3:  "GALACTIC",
		4:  "QUANTUM",
		5:  "DOMINION",
		6:  "ASTRO",
		7:  "CORSAIRS",
		8:  "OBSIDIAN",
		9:  "AEGIS",
		10: "UNITED",
		11: "SOLITARY",
		12: "COBALD",
		13: "OMEGA",
		14: "ECHO",
		15: "LORDS",
		16: "CULT",
		17: "ANCIENTS",
		18: "SHADOW",
		19: "ETHEREAL",
	}
	Faction_value = map[string]int32{
		"UNKNOWN_FACTION": 0,
		"COSMIC":          1,
		"VOID":            2,
		"GALACTIC":        3,
		"QUANTUM":         4,
		"DOMINION":        5,
		"ASTRO":           6,
		"CORSAIRS":        7,
		"OBSIDIAN":        8,
		"AEGIS":           9,
		"UNITED":          10,
		"SOLITARY":        11,
		"COBALD":          12,
		"OMEGA":           13,
		"ECHO":            14,
		"LORDS":           15,
		"CULT":            16,
		"ANCIENTS":        17,
		"SHADOW":          18,
		"ETHEREAL":        19,
	}
)

func (x Faction) Enum() *Faction {
	p := new(Faction)
	*p = x
	return p
}

func (x Faction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Faction) Descriptor() protoreflect.EnumDescriptor {
	return file_galaxy_proto_enumTypes[0].Descriptor()
}

func (Faction) Type() protoreflect.EnumType {
	return &file_galaxy_proto_enumTypes[0]
}

func (x Faction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Faction.Descriptor instead.
func (Faction) EnumDescriptor() ([]byte, []int) {
	return file_galaxy_proto_rawDescGZIP(), []int{0}
}

type System_Type int32

const (
	System_UNKNOWN_SYSTEMTYPE System_Type = 0
	System_NEUTRON_STAR       System_Type = 1
	System_RED_STAR           System_Type = 2
	System_ORANGE_STAR        System_Type = 3
	System_BLUE_STAR          System_Type = 4
	System_YOUNG_STAR         System_Type = 5
	System_WHITE_DWARF        System_Type = 6
	System_BLACK_HOLE         System_Type = 7
	System_HYPERGIANT         System_Type = 8
	System_NEBULA             System_Type = 9
	System_UNSTABLE           System_Type = 10
)

// Enum value maps for System_Type.
var (
	System_Type_name = map[int32]string{
		0:  "UNKNOWN_SYSTEMTYPE",
		1:  "NEUTRON_STAR",
		2:  "RED_STAR",
		3:  "ORANGE_STAR",
		4:  "BLUE_STAR",
		5:  "YOUNG_STAR",
		6:  "WHITE_DWARF",
		7:  "BLACK_HOLE",
		8:  "HYPERGIANT",
		9:  "NEBULA",
		10: "UNSTABLE",
	}
	System_Type_value = map[string]int32{
		"UNKNOWN_SYSTEMTYPE": 0,
		"NEUTRON_STAR":       1,
		"RED_STAR":           2,
		"ORANGE_STAR":        3,
		"BLUE_STAR":          4,
		"YOUNG_STAR":         5,
		"WHITE_DWARF":        6,
		"BLACK_HOLE":         7,
		"HYPERGIANT":         8,
		"NEBULA":             9,
		"UNSTABLE":           10,
	}
)

func (x System_Type) Enum() *System_Type {
	p := new(System_Type)
	*p = x
	return p
}

func (x System_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (System_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_galaxy_proto_enumTypes[1].Descriptor()
}

func (System_Type) Type() protoreflect.EnumType {
	return &file_galaxy_proto_enumTypes[1]
}

func (x System_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use System_Type.Descriptor instead.
func (System_Type) EnumDescriptor() ([]byte, []int) {
	return file_galaxy_proto_rawDescGZIP(), []int{0, 0}
}

type WaypointBase_Type int32

const (
	WaypointBase_UNKNOWN_WAYPOINTTYPE    WaypointBase_Type = 0
	WaypointBase_PLANET                  WaypointBase_Type = 1
	WaypointBase_GAS_GIANT               WaypointBase_Type = 2
	WaypointBase_MOON                    WaypointBase_Type = 3
	WaypointBase_ORBITAL_STATION         WaypointBase_Type = 4
	WaypointBase_JUMP_GATE               WaypointBase_Type = 5
	WaypointBase_ASTEROID_FIELD          WaypointBase_Type = 6
	WaypointBase_ASTEROID                WaypointBase_Type = 7
	WaypointBase_ENGINEERED_ASTEROID     WaypointBase_Type = 8
	WaypointBase_ASTEROID_BASE           WaypointBase_Type = 9
	WaypointBase_NEBULA                  WaypointBase_Type = 10
	WaypointBase_DEBRIS_FIELD            WaypointBase_Type = 11
	WaypointBase_GRAVITY_WELL            WaypointBase_Type = 12
	WaypointBase_ARTIFICIAL_GRAVITY_WELL WaypointBase_Type = 13
	WaypointBase_FUEL_STATION            WaypointBase_Type = 14
)

// Enum value maps for WaypointBase_Type.
var (
	WaypointBase_Type_name = map[int32]string{
		0:  "UNKNOWN_WAYPOINTTYPE",
		1:  "PLANET",
		2:  "GAS_GIANT",
		3:  "MOON",
		4:  "ORBITAL_STATION",
		5:  "JUMP_GATE",
		6:  "ASTEROID_FIELD",
		7:  "ASTEROID",
		8:  "ENGINEERED_ASTEROID",
		9:  "ASTEROID_BASE",
		10: "NEBULA",
		11: "DEBRIS_FIELD",
		12: "GRAVITY_WELL",
		13: "ARTIFICIAL_GRAVITY_WELL",
		14: "FUEL_STATION",
	}
	WaypointBase_Type_value = map[string]int32{
		"UNKNOWN_WAYPOINTTYPE":    0,
		"PLANET":                  1,
		"GAS_GIANT":               2,
		"MOON":                    3,
		"ORBITAL_STATION":         4,
		"JUMP_GATE":               5,
		"ASTEROID_FIELD":          6,
		"ASTEROID":                7,
		"ENGINEERED_ASTEROID":     8,
		"ASTEROID_BASE":           9,
		"NEBULA":                  10,
		"DEBRIS_FIELD":            11,
		"GRAVITY_WELL":            12,
		"ARTIFICIAL_GRAVITY_WELL": 13,
		"FUEL_STATION":            14,
	}
)

func (x WaypointBase_Type) Enum() *WaypointBase_Type {
	p := new(WaypointBase_Type)
	*p = x
	return p
}

func (x WaypointBase_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WaypointBase_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_galaxy_proto_enumTypes[2].Descriptor()
}

func (WaypointBase_Type) Type() protoreflect.EnumType {
	return &file_galaxy_proto_enumTypes[2]
}

func (x WaypointBase_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WaypointBase_Type.Descriptor instead.
func (WaypointBase_Type) EnumDescriptor() ([]byte, []int) {
	return file_galaxy_proto_rawDescGZIP(), []int{1, 0}
}

type System struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Sector    string          `protobuf:"bytes,2,opt,name=sector,proto3" json:"sector,omitempty"`
	Type      System_Type     `protobuf:"varint,3,opt,name=type,proto3,enum=proto.System_Type" json:"type,omitempty"`
	X         int32           `protobuf:"varint,4,opt,name=x,proto3" json:"x,omitempty"`
	Y         int32           `protobuf:"varint,5,opt,name=y,proto3" json:"y,omitempty"`
	Waypoints []*WaypointBase `protobuf:"bytes,6,rep,name=waypoints,proto3" json:"waypoints,omitempty"`
	// Factions that control this system.
	Factions []Faction `protobuf:"varint,7,rep,packed,name=factions,proto3,enum=proto.Faction" json:"factions,omitempty"`
}

func (x *System) Reset() {
	*x = System{}
	if protoimpl.UnsafeEnabled {
		mi := &file_galaxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *System) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*System) ProtoMessage() {}

func (x *System) ProtoReflect() protoreflect.Message {
	mi := &file_galaxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use System.ProtoReflect.Descriptor instead.
func (*System) Descriptor() ([]byte, []int) {
	return file_galaxy_proto_rawDescGZIP(), []int{0}
}

func (x *System) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *System) GetSector() string {
	if x != nil {
		return x.Sector
	}
	return ""
}

func (x *System) GetType() System_Type {
	if x != nil {
		return x.Type
	}
	return System_UNKNOWN_SYSTEMTYPE
}

func (x *System) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *System) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *System) GetWaypoints() []*WaypointBase {
	if x != nil {
		return x.Waypoints
	}
	return nil
}

func (x *System) GetFactions() []Faction {
	if x != nil {
		return x.Factions
	}
	return nil
}

type WaypointBase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	System string            `protobuf:"bytes,2,opt,name=system,proto3" json:"system,omitempty"`
	Type   WaypointBase_Type `protobuf:"varint,3,opt,name=type,proto3,enum=proto.WaypointBase_Type" json:"type,omitempty"`
	X      int32             `protobuf:"varint,4,opt,name=x,proto3" json:"x,omitempty"`
	Y      int32             `protobuf:"varint,5,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *WaypointBase) Reset() {
	*x = WaypointBase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_galaxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WaypointBase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WaypointBase) ProtoMessage() {}

func (x *WaypointBase) ProtoReflect() protoreflect.Message {
	mi := &file_galaxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WaypointBase.ProtoReflect.Descriptor instead.
func (*WaypointBase) Descriptor() ([]byte, []int) {
	return file_galaxy_proto_rawDescGZIP(), []int{1}
}

func (x *WaypointBase) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WaypointBase) GetSystem() string {
	if x != nil {
		return x.System
	}
	return ""
}

func (x *WaypointBase) GetType() WaypointBase_Type {
	if x != nil {
		return x.Type
	}
	return WaypointBase_UNKNOWN_WAYPOINTTYPE
}

func (x *WaypointBase) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *WaypointBase) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

var File_galaxy_proto protoreflect.FileDescriptor

var file_galaxy_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x67, 0x61, 0x6c, 0x61, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x03, 0x0a, 0x06, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x26, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c,
	0x0a, 0x01, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x12, 0x31, 0x0a, 0x09,
	0x77, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x09, 0x77, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12,
	0x2a, 0x0a, 0x08, 0x66, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x0e, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x08, 0x66, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xb9, 0x01, 0x0a, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f,
	0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x54, 0x59, 0x50, 0x45, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c,
	0x4e, 0x45, 0x55, 0x54, 0x52, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x10, 0x01, 0x12, 0x0c,
	0x0a, 0x08, 0x52, 0x45, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b,
	0x4f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x10, 0x03, 0x12, 0x0d, 0x0a,
	0x09, 0x42, 0x4c, 0x55, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a,
	0x59, 0x4f, 0x55, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b,
	0x57, 0x48, 0x49, 0x54, 0x45, 0x5f, 0x44, 0x57, 0x41, 0x52, 0x46, 0x10, 0x06, 0x12, 0x0e, 0x0a,
	0x0a, 0x42, 0x4c, 0x41, 0x43, 0x4b, 0x5f, 0x48, 0x4f, 0x4c, 0x45, 0x10, 0x07, 0x12, 0x0e, 0x0a,
	0x0a, 0x48, 0x59, 0x50, 0x45, 0x52, 0x47, 0x49, 0x41, 0x4e, 0x54, 0x10, 0x08, 0x12, 0x0a, 0x0a,
	0x06, 0x4e, 0x45, 0x42, 0x55, 0x4c, 0x41, 0x10, 0x09, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x4e, 0x53,
	0x54, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x0a, 0x22, 0x99, 0x03, 0x0a, 0x0c, 0x57, 0x61, 0x79, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x42, 0x61, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x12, 0x2c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x42,
	0x61, 0x73, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0c,
	0x0a, 0x01, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x22, 0x96, 0x02, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x57,
	0x41, 0x59, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x54, 0x59, 0x50, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x50, 0x4c, 0x41, 0x4e, 0x45, 0x54, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x47, 0x41, 0x53,
	0x5f, 0x47, 0x49, 0x41, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x4f, 0x4f, 0x4e,
	0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x4f, 0x52, 0x42, 0x49, 0x54, 0x41, 0x4c, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x4a, 0x55, 0x4d, 0x50, 0x5f,
	0x47, 0x41, 0x54, 0x45, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x53, 0x54, 0x45, 0x52, 0x4f,
	0x49, 0x44, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x53,
	0x54, 0x45, 0x52, 0x4f, 0x49, 0x44, 0x10, 0x07, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x4e, 0x47, 0x49,
	0x4e, 0x45, 0x45, 0x52, 0x45, 0x44, 0x5f, 0x41, 0x53, 0x54, 0x45, 0x52, 0x4f, 0x49, 0x44, 0x10,
	0x08, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x53, 0x54, 0x45, 0x52, 0x4f, 0x49, 0x44, 0x5f, 0x42, 0x41,
	0x53, 0x45, 0x10, 0x09, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x45, 0x42, 0x55, 0x4c, 0x41, 0x10, 0x0a,
	0x12, 0x10, 0x0a, 0x0c, 0x44, 0x45, 0x42, 0x52, 0x49, 0x53, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44,
	0x10, 0x0b, 0x12, 0x10, 0x0a, 0x0c, 0x47, 0x52, 0x41, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x57, 0x45,
	0x4c, 0x4c, 0x10, 0x0c, 0x12, 0x1b, 0x0a, 0x17, 0x41, 0x52, 0x54, 0x49, 0x46, 0x49, 0x43, 0x49,
	0x41, 0x4c, 0x5f, 0x47, 0x52, 0x41, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x57, 0x45, 0x4c, 0x4c, 0x10,
	0x0d, 0x12, 0x10, 0x0a, 0x0c, 0x46, 0x55, 0x45, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x10, 0x0e, 0x2a, 0x87, 0x02, 0x0a, 0x07, 0x46, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x13, 0x0a, 0x0f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x46, 0x41, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4f, 0x53, 0x4d, 0x49, 0x43, 0x10, 0x01,
	0x12, 0x08, 0x0a, 0x04, 0x56, 0x4f, 0x49, 0x44, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x47, 0x41,
	0x4c, 0x41, 0x43, 0x54, 0x49, 0x43, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x51, 0x55, 0x41, 0x4e,
	0x54, 0x55, 0x4d, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x4f, 0x4d, 0x49, 0x4e, 0x49, 0x4f,
	0x4e, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x53, 0x54, 0x52, 0x4f, 0x10, 0x06, 0x12, 0x0c,
	0x0a, 0x08, 0x43, 0x4f, 0x52, 0x53, 0x41, 0x49, 0x52, 0x53, 0x10, 0x07, 0x12, 0x0c, 0x0a, 0x08,
	0x4f, 0x42, 0x53, 0x49, 0x44, 0x49, 0x41, 0x4e, 0x10, 0x08, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x45,
	0x47, 0x49, 0x53, 0x10, 0x09, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x4e, 0x49, 0x54, 0x45, 0x44, 0x10,
	0x0a, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x4f, 0x4c, 0x49, 0x54, 0x41, 0x52, 0x59, 0x10, 0x0b, 0x12,
	0x0a, 0x0a, 0x06, 0x43, 0x4f, 0x42, 0x41, 0x4c, 0x44, 0x10, 0x0c, 0x12, 0x09, 0x0a, 0x05, 0x4f,
	0x4d, 0x45, 0x47, 0x41, 0x10, 0x0d, 0x12, 0x08, 0x0a, 0x04, 0x45, 0x43, 0x48, 0x4f, 0x10, 0x0e,
	0x12, 0x09, 0x0a, 0x05, 0x4c, 0x4f, 0x52, 0x44, 0x53, 0x10, 0x0f, 0x12, 0x08, 0x0a, 0x04, 0x43,
	0x55, 0x4c, 0x54, 0x10, 0x10, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x4e, 0x43, 0x49, 0x45, 0x4e, 0x54,
	0x53, 0x10, 0x11, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x48, 0x41, 0x44, 0x4f, 0x57, 0x10, 0x12, 0x12,
	0x0c, 0x0a, 0x08, 0x45, 0x54, 0x48, 0x45, 0x52, 0x45, 0x41, 0x4c, 0x10, 0x13, 0x42, 0x42, 0x5a,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x6e, 0x6f,
	0x6b, 0x6f, 0x74, 0x74, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x74, 0x72, 0x61, 0x64, 0x65, 0x72,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa,
	0x02, 0x0f, 0x47, 0x72, 0x70, 0x63, 0x53, 0x70, 0x61, 0x63, 0x65, 0x74, 0x72, 0x61, 0x64, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_galaxy_proto_rawDescOnce sync.Once
	file_galaxy_proto_rawDescData = file_galaxy_proto_rawDesc
)

func file_galaxy_proto_rawDescGZIP() []byte {
	file_galaxy_proto_rawDescOnce.Do(func() {
		file_galaxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_galaxy_proto_rawDescData)
	})
	return file_galaxy_proto_rawDescData
}

var file_galaxy_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_galaxy_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_galaxy_proto_goTypes = []any{
	(Faction)(0),           // 0: proto.Faction
	(System_Type)(0),       // 1: proto.System.Type
	(WaypointBase_Type)(0), // 2: proto.WaypointBase.Type
	(*System)(nil),         // 3: proto.System
	(*WaypointBase)(nil),   // 4: proto.WaypointBase
}
var file_galaxy_proto_depIdxs = []int32{
	1, // 0: proto.System.type:type_name -> proto.System.Type
	4, // 1: proto.System.waypoints:type_name -> proto.WaypointBase
	0, // 2: proto.System.factions:type_name -> proto.Faction
	2, // 3: proto.WaypointBase.type:type_name -> proto.WaypointBase.Type
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_galaxy_proto_init() }
func file_galaxy_proto_init() {
	if File_galaxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_galaxy_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*System); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_galaxy_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*WaypointBase); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_galaxy_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_galaxy_proto_goTypes,
		DependencyIndexes: file_galaxy_proto_depIdxs,
		EnumInfos:         file_galaxy_proto_enumTypes,
		MessageInfos:      file_galaxy_proto_msgTypes,
	}.Build()
	File_galaxy_proto = out.File
	file_galaxy_proto_rawDesc = nil
	file_galaxy_proto_goTypes = nil
	file_galaxy_proto_depIdxs = nil
}