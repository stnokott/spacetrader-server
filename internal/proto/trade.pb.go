// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: trade.proto

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

type TradeItem int32

const (
	TradeItem_UNKNOWN_TRADEITEM          TradeItem = 0
	TradeItem_PRECIOUS_STONES            TradeItem = 1
	TradeItem_QUARTZ_SAND                TradeItem = 2
	TradeItem_SILICON_CRYSTALS           TradeItem = 3
	TradeItem_AMMONIA_ICE                TradeItem = 4
	TradeItem_LIQUID_HYDROGEN            TradeItem = 5
	TradeItem_LIQUID_NITROGEN            TradeItem = 6
	TradeItem_ICE_WATER                  TradeItem = 7
	TradeItem_EXOTIC_MATTER              TradeItem = 8
	TradeItem_ADVANCED_CIRCUITRY         TradeItem = 9
	TradeItem_GRAVITON_EMITTERS          TradeItem = 10
	TradeItem_IRON                       TradeItem = 11
	TradeItem_IRON_ORE                   TradeItem = 12
	TradeItem_COPPER                     TradeItem = 13
	TradeItem_COPPER_ORE                 TradeItem = 14
	TradeItem_ALUMINUM                   TradeItem = 15
	TradeItem_ALUMINUM_ORE               TradeItem = 16
	TradeItem_SILVER                     TradeItem = 17
	TradeItem_SILVER_ORE                 TradeItem = 18
	TradeItem_GOLD                       TradeItem = 19
	TradeItem_GOLD_ORE                   TradeItem = 20
	TradeItem_PLATINUM                   TradeItem = 21
	TradeItem_PLATINUM_ORE               TradeItem = 22
	TradeItem_DIAMONDS                   TradeItem = 23
	TradeItem_URANITE                    TradeItem = 24
	TradeItem_URANITE_ORE                TradeItem = 25
	TradeItem_MERITIUM                   TradeItem = 26
	TradeItem_MERITIUM_ORE               TradeItem = 27
	TradeItem_HYDROCARBON                TradeItem = 28
	TradeItem_ANTIMATTER                 TradeItem = 29
	TradeItem_FAB_MATS                   TradeItem = 30
	TradeItem_FERTILIZERS                TradeItem = 31
	TradeItem_FABRICS                    TradeItem = 32
	TradeItem_FOOD                       TradeItem = 33
	TradeItem_JEWELRY                    TradeItem = 34
	TradeItem_MACHINERY                  TradeItem = 35
	TradeItem_FIREARMS                   TradeItem = 36
	TradeItem_ASSAULT_RIFLES             TradeItem = 37
	TradeItem_MILITARY_EQUIPMENT         TradeItem = 38
	TradeItem_EXPLOSIVES                 TradeItem = 39
	TradeItem_LAB_INSTRUMENTS            TradeItem = 40
	TradeItem_AMMUNITION                 TradeItem = 41
	TradeItem_ELECTRONICS                TradeItem = 42
	TradeItem_SHIP_PLATING               TradeItem = 43
	TradeItem_SHIP_PARTS                 TradeItem = 44
	TradeItem_EQUIPMENT                  TradeItem = 45
	TradeItem_FUEL                       TradeItem = 46
	TradeItem_MEDICINE                   TradeItem = 47
	TradeItem_DRUGS                      TradeItem = 48
	TradeItem_CLOTHING                   TradeItem = 49
	TradeItem_MICROPROCESSORS            TradeItem = 50
	TradeItem_PLASTICS                   TradeItem = 51
	TradeItem_POLYNUCLEOTIDES            TradeItem = 52
	TradeItem_BIOCOMPOSITES              TradeItem = 53
	TradeItem_QUANTUM_STABILIZERS        TradeItem = 54
	TradeItem_NANOBOTS                   TradeItem = 55
	TradeItem_AI_MAINFRAMES              TradeItem = 56
	TradeItem_QUANTUM_DRIVES             TradeItem = 57
	TradeItem_ROBOTIC_DRONES             TradeItem = 58
	TradeItem_CYBER_IMPLANTS             TradeItem = 59
	TradeItem_GENE_THERAPEUTICS          TradeItem = 60
	TradeItem_NEURAL_CHIPS               TradeItem = 61
	TradeItem_MOOD_REGULATORS            TradeItem = 62
	TradeItem_VIRAL_AGENTS               TradeItem = 63
	TradeItem_MICRO_FUSION_GENERATORS    TradeItem = 64
	TradeItem_SUPERGRAINS                TradeItem = 65
	TradeItem_LASER_RIFLES               TradeItem = 66
	TradeItem_HOLOGRAPHICS               TradeItem = 67
	TradeItem_SHIP_SALVAGE               TradeItem = 68
	TradeItem_RELIC_TECH                 TradeItem = 69
	TradeItem_NOVEL_LIFEFORMS            TradeItem = 70
	TradeItem_BOTANICAL_SPECIMENS        TradeItem = 71
	TradeItem_CULTURAL_ARTIFACTS         TradeItem = 72
	TradeItem_FRAME_PROBE                TradeItem = 73
	TradeItem_FRAME_DRONE                TradeItem = 74
	TradeItem_FRAME_INTERCEPTOR          TradeItem = 75
	TradeItem_FRAME_RACER                TradeItem = 76
	TradeItem_FRAME_FIGHTER              TradeItem = 77
	TradeItem_FRAME_FRIGATE              TradeItem = 78
	TradeItem_FRAME_SHUTTLE              TradeItem = 79
	TradeItem_FRAME_EXPLORER             TradeItem = 80
	TradeItem_FRAME_MINER                TradeItem = 81
	TradeItem_FRAME_LIGHT_FREIGHTER      TradeItem = 82
	TradeItem_FRAME_HEAVY_FREIGHTER      TradeItem = 83
	TradeItem_FRAME_TRANSPORT            TradeItem = 84
	TradeItem_FRAME_DESTROYER            TradeItem = 85
	TradeItem_FRAME_CRUISER              TradeItem = 86
	TradeItem_FRAME_CARRIER              TradeItem = 87
	TradeItem_REACTOR_SOLAR_I            TradeItem = 88
	TradeItem_REACTOR_FUSION_I           TradeItem = 89
	TradeItem_REACTOR_FISSION_I          TradeItem = 90
	TradeItem_REACTOR_CHEMICAL_I         TradeItem = 91
	TradeItem_REACTOR_ANTIMATTER_I       TradeItem = 92
	TradeItem_ENGINE_IMPULSE_DRIVE_I     TradeItem = 93
	TradeItem_ENGINE_ION_DRIVE_I         TradeItem = 94
	TradeItem_ENGINE_ION_DRIVE_II        TradeItem = 95
	TradeItem_ENGINE_HYPER_DRIVE_I       TradeItem = 96
	TradeItem_MODULE_MINERAL_PROCESSOR_I TradeItem = 97
	TradeItem_MODULE_GAS_PROCESSOR_I     TradeItem = 98
	TradeItem_MODULE_CARGO_HOLD_I        TradeItem = 99
	TradeItem_MODULE_CARGO_HOLD_II       TradeItem = 100
	TradeItem_MODULE_CARGO_HOLD_III      TradeItem = 101
	TradeItem_MODULE_CREW_QUARTERS_I     TradeItem = 102
	TradeItem_MODULE_ENVOY_QUARTERS_I    TradeItem = 103
	TradeItem_MODULE_PASSENGER_CABIN_I   TradeItem = 104
	TradeItem_MODULE_MICRO_REFINERY_I    TradeItem = 105
	TradeItem_MODULE_SCIENCE_LAB_I       TradeItem = 106
	TradeItem_MODULE_JUMP_DRIVE_I        TradeItem = 107
	TradeItem_MODULE_JUMP_DRIVE_II       TradeItem = 108
	TradeItem_MODULE_JUMP_DRIVE_III      TradeItem = 109
	TradeItem_MODULE_WARP_DRIVE_I        TradeItem = 110
	TradeItem_MODULE_WARP_DRIVE_II       TradeItem = 111
	TradeItem_MODULE_WARP_DRIVE_III      TradeItem = 112
	TradeItem_MODULE_SHIELD_GENERATOR_I  TradeItem = 113
	TradeItem_MODULE_SHIELD_GENERATOR_II TradeItem = 114
	TradeItem_MODULE_ORE_REFINERY_I      TradeItem = 115
	TradeItem_MODULE_FUEL_REFINERY_I     TradeItem = 116
	TradeItem_MOUNT_GAS_SIPHON_I         TradeItem = 117
	TradeItem_MOUNT_GAS_SIPHON_II        TradeItem = 118
	TradeItem_MOUNT_GAS_SIPHON_III       TradeItem = 119
	TradeItem_MOUNT_SURVEYOR_I           TradeItem = 120
	TradeItem_MOUNT_SURVEYOR_II          TradeItem = 121
	TradeItem_MOUNT_SURVEYOR_III         TradeItem = 122
	TradeItem_MOUNT_SENSOR_ARRAY_I       TradeItem = 123
	TradeItem_MOUNT_SENSOR_ARRAY_II      TradeItem = 124
	TradeItem_MOUNT_SENSOR_ARRAY_III     TradeItem = 125
	TradeItem_MOUNT_MINING_LASER_I       TradeItem = 126
	TradeItem_MOUNT_MINING_LASER_II      TradeItem = 127
	TradeItem_MOUNT_MINING_LASER_III     TradeItem = 128
	TradeItem_MOUNT_LASER_CANNON_I       TradeItem = 129
	TradeItem_MOUNT_MISSILE_LAUNCHER_I   TradeItem = 130
	TradeItem_MOUNT_TURRET_I             TradeItem = 131
	TradeItem_SHIP_PROBE                 TradeItem = 132
	TradeItem_SHIP_MINING_DRONE          TradeItem = 133
	TradeItem_SHIP_SIPHON_DRONE          TradeItem = 134
	TradeItem_SHIP_INTERCEPTOR           TradeItem = 135
	TradeItem_SHIP_LIGHT_HAULER          TradeItem = 136
	TradeItem_SHIP_COMMAND_FRIGATE       TradeItem = 137
	TradeItem_SHIP_EXPLORER              TradeItem = 138
	TradeItem_SHIP_HEAVY_FREIGHTER       TradeItem = 139
	TradeItem_SHIP_LIGHT_SHUTTLE         TradeItem = 140
	TradeItem_SHIP_ORE_HOUND             TradeItem = 141
	TradeItem_SHIP_REFINING_FREIGHTER    TradeItem = 142
	TradeItem_SHIP_SURVEYOR              TradeItem = 143
)

// Enum value maps for TradeItem.
var (
	TradeItem_name = map[int32]string{
		0:   "UNKNOWN_TRADEITEM",
		1:   "PRECIOUS_STONES",
		2:   "QUARTZ_SAND",
		3:   "SILICON_CRYSTALS",
		4:   "AMMONIA_ICE",
		5:   "LIQUID_HYDROGEN",
		6:   "LIQUID_NITROGEN",
		7:   "ICE_WATER",
		8:   "EXOTIC_MATTER",
		9:   "ADVANCED_CIRCUITRY",
		10:  "GRAVITON_EMITTERS",
		11:  "IRON",
		12:  "IRON_ORE",
		13:  "COPPER",
		14:  "COPPER_ORE",
		15:  "ALUMINUM",
		16:  "ALUMINUM_ORE",
		17:  "SILVER",
		18:  "SILVER_ORE",
		19:  "GOLD",
		20:  "GOLD_ORE",
		21:  "PLATINUM",
		22:  "PLATINUM_ORE",
		23:  "DIAMONDS",
		24:  "URANITE",
		25:  "URANITE_ORE",
		26:  "MERITIUM",
		27:  "MERITIUM_ORE",
		28:  "HYDROCARBON",
		29:  "ANTIMATTER",
		30:  "FAB_MATS",
		31:  "FERTILIZERS",
		32:  "FABRICS",
		33:  "FOOD",
		34:  "JEWELRY",
		35:  "MACHINERY",
		36:  "FIREARMS",
		37:  "ASSAULT_RIFLES",
		38:  "MILITARY_EQUIPMENT",
		39:  "EXPLOSIVES",
		40:  "LAB_INSTRUMENTS",
		41:  "AMMUNITION",
		42:  "ELECTRONICS",
		43:  "SHIP_PLATING",
		44:  "SHIP_PARTS",
		45:  "EQUIPMENT",
		46:  "FUEL",
		47:  "MEDICINE",
		48:  "DRUGS",
		49:  "CLOTHING",
		50:  "MICROPROCESSORS",
		51:  "PLASTICS",
		52:  "POLYNUCLEOTIDES",
		53:  "BIOCOMPOSITES",
		54:  "QUANTUM_STABILIZERS",
		55:  "NANOBOTS",
		56:  "AI_MAINFRAMES",
		57:  "QUANTUM_DRIVES",
		58:  "ROBOTIC_DRONES",
		59:  "CYBER_IMPLANTS",
		60:  "GENE_THERAPEUTICS",
		61:  "NEURAL_CHIPS",
		62:  "MOOD_REGULATORS",
		63:  "VIRAL_AGENTS",
		64:  "MICRO_FUSION_GENERATORS",
		65:  "SUPERGRAINS",
		66:  "LASER_RIFLES",
		67:  "HOLOGRAPHICS",
		68:  "SHIP_SALVAGE",
		69:  "RELIC_TECH",
		70:  "NOVEL_LIFEFORMS",
		71:  "BOTANICAL_SPECIMENS",
		72:  "CULTURAL_ARTIFACTS",
		73:  "FRAME_PROBE",
		74:  "FRAME_DRONE",
		75:  "FRAME_INTERCEPTOR",
		76:  "FRAME_RACER",
		77:  "FRAME_FIGHTER",
		78:  "FRAME_FRIGATE",
		79:  "FRAME_SHUTTLE",
		80:  "FRAME_EXPLORER",
		81:  "FRAME_MINER",
		82:  "FRAME_LIGHT_FREIGHTER",
		83:  "FRAME_HEAVY_FREIGHTER",
		84:  "FRAME_TRANSPORT",
		85:  "FRAME_DESTROYER",
		86:  "FRAME_CRUISER",
		87:  "FRAME_CARRIER",
		88:  "REACTOR_SOLAR_I",
		89:  "REACTOR_FUSION_I",
		90:  "REACTOR_FISSION_I",
		91:  "REACTOR_CHEMICAL_I",
		92:  "REACTOR_ANTIMATTER_I",
		93:  "ENGINE_IMPULSE_DRIVE_I",
		94:  "ENGINE_ION_DRIVE_I",
		95:  "ENGINE_ION_DRIVE_II",
		96:  "ENGINE_HYPER_DRIVE_I",
		97:  "MODULE_MINERAL_PROCESSOR_I",
		98:  "MODULE_GAS_PROCESSOR_I",
		99:  "MODULE_CARGO_HOLD_I",
		100: "MODULE_CARGO_HOLD_II",
		101: "MODULE_CARGO_HOLD_III",
		102: "MODULE_CREW_QUARTERS_I",
		103: "MODULE_ENVOY_QUARTERS_I",
		104: "MODULE_PASSENGER_CABIN_I",
		105: "MODULE_MICRO_REFINERY_I",
		106: "MODULE_SCIENCE_LAB_I",
		107: "MODULE_JUMP_DRIVE_I",
		108: "MODULE_JUMP_DRIVE_II",
		109: "MODULE_JUMP_DRIVE_III",
		110: "MODULE_WARP_DRIVE_I",
		111: "MODULE_WARP_DRIVE_II",
		112: "MODULE_WARP_DRIVE_III",
		113: "MODULE_SHIELD_GENERATOR_I",
		114: "MODULE_SHIELD_GENERATOR_II",
		115: "MODULE_ORE_REFINERY_I",
		116: "MODULE_FUEL_REFINERY_I",
		117: "MOUNT_GAS_SIPHON_I",
		118: "MOUNT_GAS_SIPHON_II",
		119: "MOUNT_GAS_SIPHON_III",
		120: "MOUNT_SURVEYOR_I",
		121: "MOUNT_SURVEYOR_II",
		122: "MOUNT_SURVEYOR_III",
		123: "MOUNT_SENSOR_ARRAY_I",
		124: "MOUNT_SENSOR_ARRAY_II",
		125: "MOUNT_SENSOR_ARRAY_III",
		126: "MOUNT_MINING_LASER_I",
		127: "MOUNT_MINING_LASER_II",
		128: "MOUNT_MINING_LASER_III",
		129: "MOUNT_LASER_CANNON_I",
		130: "MOUNT_MISSILE_LAUNCHER_I",
		131: "MOUNT_TURRET_I",
		132: "SHIP_PROBE",
		133: "SHIP_MINING_DRONE",
		134: "SHIP_SIPHON_DRONE",
		135: "SHIP_INTERCEPTOR",
		136: "SHIP_LIGHT_HAULER",
		137: "SHIP_COMMAND_FRIGATE",
		138: "SHIP_EXPLORER",
		139: "SHIP_HEAVY_FREIGHTER",
		140: "SHIP_LIGHT_SHUTTLE",
		141: "SHIP_ORE_HOUND",
		142: "SHIP_REFINING_FREIGHTER",
		143: "SHIP_SURVEYOR",
	}
	TradeItem_value = map[string]int32{
		"UNKNOWN_TRADEITEM":          0,
		"PRECIOUS_STONES":            1,
		"QUARTZ_SAND":                2,
		"SILICON_CRYSTALS":           3,
		"AMMONIA_ICE":                4,
		"LIQUID_HYDROGEN":            5,
		"LIQUID_NITROGEN":            6,
		"ICE_WATER":                  7,
		"EXOTIC_MATTER":              8,
		"ADVANCED_CIRCUITRY":         9,
		"GRAVITON_EMITTERS":          10,
		"IRON":                       11,
		"IRON_ORE":                   12,
		"COPPER":                     13,
		"COPPER_ORE":                 14,
		"ALUMINUM":                   15,
		"ALUMINUM_ORE":               16,
		"SILVER":                     17,
		"SILVER_ORE":                 18,
		"GOLD":                       19,
		"GOLD_ORE":                   20,
		"PLATINUM":                   21,
		"PLATINUM_ORE":               22,
		"DIAMONDS":                   23,
		"URANITE":                    24,
		"URANITE_ORE":                25,
		"MERITIUM":                   26,
		"MERITIUM_ORE":               27,
		"HYDROCARBON":                28,
		"ANTIMATTER":                 29,
		"FAB_MATS":                   30,
		"FERTILIZERS":                31,
		"FABRICS":                    32,
		"FOOD":                       33,
		"JEWELRY":                    34,
		"MACHINERY":                  35,
		"FIREARMS":                   36,
		"ASSAULT_RIFLES":             37,
		"MILITARY_EQUIPMENT":         38,
		"EXPLOSIVES":                 39,
		"LAB_INSTRUMENTS":            40,
		"AMMUNITION":                 41,
		"ELECTRONICS":                42,
		"SHIP_PLATING":               43,
		"SHIP_PARTS":                 44,
		"EQUIPMENT":                  45,
		"FUEL":                       46,
		"MEDICINE":                   47,
		"DRUGS":                      48,
		"CLOTHING":                   49,
		"MICROPROCESSORS":            50,
		"PLASTICS":                   51,
		"POLYNUCLEOTIDES":            52,
		"BIOCOMPOSITES":              53,
		"QUANTUM_STABILIZERS":        54,
		"NANOBOTS":                   55,
		"AI_MAINFRAMES":              56,
		"QUANTUM_DRIVES":             57,
		"ROBOTIC_DRONES":             58,
		"CYBER_IMPLANTS":             59,
		"GENE_THERAPEUTICS":          60,
		"NEURAL_CHIPS":               61,
		"MOOD_REGULATORS":            62,
		"VIRAL_AGENTS":               63,
		"MICRO_FUSION_GENERATORS":    64,
		"SUPERGRAINS":                65,
		"LASER_RIFLES":               66,
		"HOLOGRAPHICS":               67,
		"SHIP_SALVAGE":               68,
		"RELIC_TECH":                 69,
		"NOVEL_LIFEFORMS":            70,
		"BOTANICAL_SPECIMENS":        71,
		"CULTURAL_ARTIFACTS":         72,
		"FRAME_PROBE":                73,
		"FRAME_DRONE":                74,
		"FRAME_INTERCEPTOR":          75,
		"FRAME_RACER":                76,
		"FRAME_FIGHTER":              77,
		"FRAME_FRIGATE":              78,
		"FRAME_SHUTTLE":              79,
		"FRAME_EXPLORER":             80,
		"FRAME_MINER":                81,
		"FRAME_LIGHT_FREIGHTER":      82,
		"FRAME_HEAVY_FREIGHTER":      83,
		"FRAME_TRANSPORT":            84,
		"FRAME_DESTROYER":            85,
		"FRAME_CRUISER":              86,
		"FRAME_CARRIER":              87,
		"REACTOR_SOLAR_I":            88,
		"REACTOR_FUSION_I":           89,
		"REACTOR_FISSION_I":          90,
		"REACTOR_CHEMICAL_I":         91,
		"REACTOR_ANTIMATTER_I":       92,
		"ENGINE_IMPULSE_DRIVE_I":     93,
		"ENGINE_ION_DRIVE_I":         94,
		"ENGINE_ION_DRIVE_II":        95,
		"ENGINE_HYPER_DRIVE_I":       96,
		"MODULE_MINERAL_PROCESSOR_I": 97,
		"MODULE_GAS_PROCESSOR_I":     98,
		"MODULE_CARGO_HOLD_I":        99,
		"MODULE_CARGO_HOLD_II":       100,
		"MODULE_CARGO_HOLD_III":      101,
		"MODULE_CREW_QUARTERS_I":     102,
		"MODULE_ENVOY_QUARTERS_I":    103,
		"MODULE_PASSENGER_CABIN_I":   104,
		"MODULE_MICRO_REFINERY_I":    105,
		"MODULE_SCIENCE_LAB_I":       106,
		"MODULE_JUMP_DRIVE_I":        107,
		"MODULE_JUMP_DRIVE_II":       108,
		"MODULE_JUMP_DRIVE_III":      109,
		"MODULE_WARP_DRIVE_I":        110,
		"MODULE_WARP_DRIVE_II":       111,
		"MODULE_WARP_DRIVE_III":      112,
		"MODULE_SHIELD_GENERATOR_I":  113,
		"MODULE_SHIELD_GENERATOR_II": 114,
		"MODULE_ORE_REFINERY_I":      115,
		"MODULE_FUEL_REFINERY_I":     116,
		"MOUNT_GAS_SIPHON_I":         117,
		"MOUNT_GAS_SIPHON_II":        118,
		"MOUNT_GAS_SIPHON_III":       119,
		"MOUNT_SURVEYOR_I":           120,
		"MOUNT_SURVEYOR_II":          121,
		"MOUNT_SURVEYOR_III":         122,
		"MOUNT_SENSOR_ARRAY_I":       123,
		"MOUNT_SENSOR_ARRAY_II":      124,
		"MOUNT_SENSOR_ARRAY_III":     125,
		"MOUNT_MINING_LASER_I":       126,
		"MOUNT_MINING_LASER_II":      127,
		"MOUNT_MINING_LASER_III":     128,
		"MOUNT_LASER_CANNON_I":       129,
		"MOUNT_MISSILE_LAUNCHER_I":   130,
		"MOUNT_TURRET_I":             131,
		"SHIP_PROBE":                 132,
		"SHIP_MINING_DRONE":          133,
		"SHIP_SIPHON_DRONE":          134,
		"SHIP_INTERCEPTOR":           135,
		"SHIP_LIGHT_HAULER":          136,
		"SHIP_COMMAND_FRIGATE":       137,
		"SHIP_EXPLORER":              138,
		"SHIP_HEAVY_FREIGHTER":       139,
		"SHIP_LIGHT_SHUTTLE":         140,
		"SHIP_ORE_HOUND":             141,
		"SHIP_REFINING_FREIGHTER":    142,
		"SHIP_SURVEYOR":              143,
	}
)

func (x TradeItem) Enum() *TradeItem {
	p := new(TradeItem)
	*p = x
	return p
}

func (x TradeItem) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TradeItem) Descriptor() protoreflect.EnumDescriptor {
	return file_trade_proto_enumTypes[0].Descriptor()
}

func (TradeItem) Type() protoreflect.EnumType {
	return &file_trade_proto_enumTypes[0]
}

func (x TradeItem) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TradeItem.Descriptor instead.
func (TradeItem) EnumDescriptor() ([]byte, []int) {
	return file_trade_proto_rawDescGZIP(), []int{0}
}

var File_trade_proto protoreflect.FileDescriptor

var file_trade_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xcd, 0x17, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x64, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x54, 0x52,
	0x41, 0x44, 0x45, 0x49, 0x54, 0x45, 0x4d, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x52, 0x45,
	0x43, 0x49, 0x4f, 0x55, 0x53, 0x5f, 0x53, 0x54, 0x4f, 0x4e, 0x45, 0x53, 0x10, 0x01, 0x12, 0x0f,
	0x0a, 0x0b, 0x51, 0x55, 0x41, 0x52, 0x54, 0x5a, 0x5f, 0x53, 0x41, 0x4e, 0x44, 0x10, 0x02, 0x12,
	0x14, 0x0a, 0x10, 0x53, 0x49, 0x4c, 0x49, 0x43, 0x4f, 0x4e, 0x5f, 0x43, 0x52, 0x59, 0x53, 0x54,
	0x41, 0x4c, 0x53, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x4d, 0x4d, 0x4f, 0x4e, 0x49, 0x41,
	0x5f, 0x49, 0x43, 0x45, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x49, 0x51, 0x55, 0x49, 0x44,
	0x5f, 0x48, 0x59, 0x44, 0x52, 0x4f, 0x47, 0x45, 0x4e, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x4c,
	0x49, 0x51, 0x55, 0x49, 0x44, 0x5f, 0x4e, 0x49, 0x54, 0x52, 0x4f, 0x47, 0x45, 0x4e, 0x10, 0x06,
	0x12, 0x0d, 0x0a, 0x09, 0x49, 0x43, 0x45, 0x5f, 0x57, 0x41, 0x54, 0x45, 0x52, 0x10, 0x07, 0x12,
	0x11, 0x0a, 0x0d, 0x45, 0x58, 0x4f, 0x54, 0x49, 0x43, 0x5f, 0x4d, 0x41, 0x54, 0x54, 0x45, 0x52,
	0x10, 0x08, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x44, 0x56, 0x41, 0x4e, 0x43, 0x45, 0x44, 0x5f, 0x43,
	0x49, 0x52, 0x43, 0x55, 0x49, 0x54, 0x52, 0x59, 0x10, 0x09, 0x12, 0x15, 0x0a, 0x11, 0x47, 0x52,
	0x41, 0x56, 0x49, 0x54, 0x4f, 0x4e, 0x5f, 0x45, 0x4d, 0x49, 0x54, 0x54, 0x45, 0x52, 0x53, 0x10,
	0x0a, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x52, 0x4f, 0x4e, 0x10, 0x0b, 0x12, 0x0c, 0x0a, 0x08, 0x49,
	0x52, 0x4f, 0x4e, 0x5f, 0x4f, 0x52, 0x45, 0x10, 0x0c, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4f, 0x50,
	0x50, 0x45, 0x52, 0x10, 0x0d, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x50, 0x50, 0x45, 0x52, 0x5f,
	0x4f, 0x52, 0x45, 0x10, 0x0e, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x4c, 0x55, 0x4d, 0x49, 0x4e, 0x55,
	0x4d, 0x10, 0x0f, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x4c, 0x55, 0x4d, 0x49, 0x4e, 0x55, 0x4d, 0x5f,
	0x4f, 0x52, 0x45, 0x10, 0x10, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x49, 0x4c, 0x56, 0x45, 0x52, 0x10,
	0x11, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x49, 0x4c, 0x56, 0x45, 0x52, 0x5f, 0x4f, 0x52, 0x45, 0x10,
	0x12, 0x12, 0x08, 0x0a, 0x04, 0x47, 0x4f, 0x4c, 0x44, 0x10, 0x13, 0x12, 0x0c, 0x0a, 0x08, 0x47,
	0x4f, 0x4c, 0x44, 0x5f, 0x4f, 0x52, 0x45, 0x10, 0x14, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x4c, 0x41,
	0x54, 0x49, 0x4e, 0x55, 0x4d, 0x10, 0x15, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x4c, 0x41, 0x54, 0x49,
	0x4e, 0x55, 0x4d, 0x5f, 0x4f, 0x52, 0x45, 0x10, 0x16, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x49, 0x41,
	0x4d, 0x4f, 0x4e, 0x44, 0x53, 0x10, 0x17, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x52, 0x41, 0x4e, 0x49,
	0x54, 0x45, 0x10, 0x18, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x52, 0x41, 0x4e, 0x49, 0x54, 0x45, 0x5f,
	0x4f, 0x52, 0x45, 0x10, 0x19, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x45, 0x52, 0x49, 0x54, 0x49, 0x55,
	0x4d, 0x10, 0x1a, 0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x45, 0x52, 0x49, 0x54, 0x49, 0x55, 0x4d, 0x5f,
	0x4f, 0x52, 0x45, 0x10, 0x1b, 0x12, 0x0f, 0x0a, 0x0b, 0x48, 0x59, 0x44, 0x52, 0x4f, 0x43, 0x41,
	0x52, 0x42, 0x4f, 0x4e, 0x10, 0x1c, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x4e, 0x54, 0x49, 0x4d, 0x41,
	0x54, 0x54, 0x45, 0x52, 0x10, 0x1d, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x41, 0x42, 0x5f, 0x4d, 0x41,
	0x54, 0x53, 0x10, 0x1e, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x45, 0x52, 0x54, 0x49, 0x4c, 0x49, 0x5a,
	0x45, 0x52, 0x53, 0x10, 0x1f, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x41, 0x42, 0x52, 0x49, 0x43, 0x53,
	0x10, 0x20, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x4f, 0x4f, 0x44, 0x10, 0x21, 0x12, 0x0b, 0x0a, 0x07,
	0x4a, 0x45, 0x57, 0x45, 0x4c, 0x52, 0x59, 0x10, 0x22, 0x12, 0x0d, 0x0a, 0x09, 0x4d, 0x41, 0x43,
	0x48, 0x49, 0x4e, 0x45, 0x52, 0x59, 0x10, 0x23, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x49, 0x52, 0x45,
	0x41, 0x52, 0x4d, 0x53, 0x10, 0x24, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x53, 0x53, 0x41, 0x55, 0x4c,
	0x54, 0x5f, 0x52, 0x49, 0x46, 0x4c, 0x45, 0x53, 0x10, 0x25, 0x12, 0x16, 0x0a, 0x12, 0x4d, 0x49,
	0x4c, 0x49, 0x54, 0x41, 0x52, 0x59, 0x5f, 0x45, 0x51, 0x55, 0x49, 0x50, 0x4d, 0x45, 0x4e, 0x54,
	0x10, 0x26, 0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x58, 0x50, 0x4c, 0x4f, 0x53, 0x49, 0x56, 0x45, 0x53,
	0x10, 0x27, 0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x41, 0x42, 0x5f, 0x49, 0x4e, 0x53, 0x54, 0x52, 0x55,
	0x4d, 0x45, 0x4e, 0x54, 0x53, 0x10, 0x28, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x4d, 0x4d, 0x55, 0x4e,
	0x49, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x29, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x4c, 0x45, 0x43, 0x54,
	0x52, 0x4f, 0x4e, 0x49, 0x43, 0x53, 0x10, 0x2a, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x48, 0x49, 0x50,
	0x5f, 0x50, 0x4c, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x2b, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x48,
	0x49, 0x50, 0x5f, 0x50, 0x41, 0x52, 0x54, 0x53, 0x10, 0x2c, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x51,
	0x55, 0x49, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x2d, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x55, 0x45,
	0x4c, 0x10, 0x2e, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x45, 0x44, 0x49, 0x43, 0x49, 0x4e, 0x45, 0x10,
	0x2f, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x52, 0x55, 0x47, 0x53, 0x10, 0x30, 0x12, 0x0c, 0x0a, 0x08,
	0x43, 0x4c, 0x4f, 0x54, 0x48, 0x49, 0x4e, 0x47, 0x10, 0x31, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x49,
	0x43, 0x52, 0x4f, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x4f, 0x52, 0x53, 0x10, 0x32, 0x12,
	0x0c, 0x0a, 0x08, 0x50, 0x4c, 0x41, 0x53, 0x54, 0x49, 0x43, 0x53, 0x10, 0x33, 0x12, 0x13, 0x0a,
	0x0f, 0x50, 0x4f, 0x4c, 0x59, 0x4e, 0x55, 0x43, 0x4c, 0x45, 0x4f, 0x54, 0x49, 0x44, 0x45, 0x53,
	0x10, 0x34, 0x12, 0x11, 0x0a, 0x0d, 0x42, 0x49, 0x4f, 0x43, 0x4f, 0x4d, 0x50, 0x4f, 0x53, 0x49,
	0x54, 0x45, 0x53, 0x10, 0x35, 0x12, 0x17, 0x0a, 0x13, 0x51, 0x55, 0x41, 0x4e, 0x54, 0x55, 0x4d,
	0x5f, 0x53, 0x54, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x5a, 0x45, 0x52, 0x53, 0x10, 0x36, 0x12, 0x0c,
	0x0a, 0x08, 0x4e, 0x41, 0x4e, 0x4f, 0x42, 0x4f, 0x54, 0x53, 0x10, 0x37, 0x12, 0x11, 0x0a, 0x0d,
	0x41, 0x49, 0x5f, 0x4d, 0x41, 0x49, 0x4e, 0x46, 0x52, 0x41, 0x4d, 0x45, 0x53, 0x10, 0x38, 0x12,
	0x12, 0x0a, 0x0e, 0x51, 0x55, 0x41, 0x4e, 0x54, 0x55, 0x4d, 0x5f, 0x44, 0x52, 0x49, 0x56, 0x45,
	0x53, 0x10, 0x39, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x4f, 0x42, 0x4f, 0x54, 0x49, 0x43, 0x5f, 0x44,
	0x52, 0x4f, 0x4e, 0x45, 0x53, 0x10, 0x3a, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x59, 0x42, 0x45, 0x52,
	0x5f, 0x49, 0x4d, 0x50, 0x4c, 0x41, 0x4e, 0x54, 0x53, 0x10, 0x3b, 0x12, 0x15, 0x0a, 0x11, 0x47,
	0x45, 0x4e, 0x45, 0x5f, 0x54, 0x48, 0x45, 0x52, 0x41, 0x50, 0x45, 0x55, 0x54, 0x49, 0x43, 0x53,
	0x10, 0x3c, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x45, 0x55, 0x52, 0x41, 0x4c, 0x5f, 0x43, 0x48, 0x49,
	0x50, 0x53, 0x10, 0x3d, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x4f, 0x4f, 0x44, 0x5f, 0x52, 0x45, 0x47,
	0x55, 0x4c, 0x41, 0x54, 0x4f, 0x52, 0x53, 0x10, 0x3e, 0x12, 0x10, 0x0a, 0x0c, 0x56, 0x49, 0x52,
	0x41, 0x4c, 0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x53, 0x10, 0x3f, 0x12, 0x1b, 0x0a, 0x17, 0x4d,
	0x49, 0x43, 0x52, 0x4f, 0x5f, 0x46, 0x55, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x47, 0x45, 0x4e, 0x45,
	0x52, 0x41, 0x54, 0x4f, 0x52, 0x53, 0x10, 0x40, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x55, 0x50, 0x45,
	0x52, 0x47, 0x52, 0x41, 0x49, 0x4e, 0x53, 0x10, 0x41, 0x12, 0x10, 0x0a, 0x0c, 0x4c, 0x41, 0x53,
	0x45, 0x52, 0x5f, 0x52, 0x49, 0x46, 0x4c, 0x45, 0x53, 0x10, 0x42, 0x12, 0x10, 0x0a, 0x0c, 0x48,
	0x4f, 0x4c, 0x4f, 0x47, 0x52, 0x41, 0x50, 0x48, 0x49, 0x43, 0x53, 0x10, 0x43, 0x12, 0x10, 0x0a,
	0x0c, 0x53, 0x48, 0x49, 0x50, 0x5f, 0x53, 0x41, 0x4c, 0x56, 0x41, 0x47, 0x45, 0x10, 0x44, 0x12,
	0x0e, 0x0a, 0x0a, 0x52, 0x45, 0x4c, 0x49, 0x43, 0x5f, 0x54, 0x45, 0x43, 0x48, 0x10, 0x45, 0x12,
	0x13, 0x0a, 0x0f, 0x4e, 0x4f, 0x56, 0x45, 0x4c, 0x5f, 0x4c, 0x49, 0x46, 0x45, 0x46, 0x4f, 0x52,
	0x4d, 0x53, 0x10, 0x46, 0x12, 0x17, 0x0a, 0x13, 0x42, 0x4f, 0x54, 0x41, 0x4e, 0x49, 0x43, 0x41,
	0x4c, 0x5f, 0x53, 0x50, 0x45, 0x43, 0x49, 0x4d, 0x45, 0x4e, 0x53, 0x10, 0x47, 0x12, 0x16, 0x0a,
	0x12, 0x43, 0x55, 0x4c, 0x54, 0x55, 0x52, 0x41, 0x4c, 0x5f, 0x41, 0x52, 0x54, 0x49, 0x46, 0x41,
	0x43, 0x54, 0x53, 0x10, 0x48, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x52, 0x41, 0x4d, 0x45, 0x5f, 0x50,
	0x52, 0x4f, 0x42, 0x45, 0x10, 0x49, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x52, 0x41, 0x4d, 0x45, 0x5f,
	0x44, 0x52, 0x4f, 0x4e, 0x45, 0x10, 0x4a, 0x12, 0x15, 0x0a, 0x11, 0x46, 0x52, 0x41, 0x4d, 0x45,
	0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x43, 0x45, 0x50, 0x54, 0x4f, 0x52, 0x10, 0x4b, 0x12, 0x0f,
	0x0a, 0x0b, 0x46, 0x52, 0x41, 0x4d, 0x45, 0x5f, 0x52, 0x41, 0x43, 0x45, 0x52, 0x10, 0x4c, 0x12,
	0x11, 0x0a, 0x0d, 0x46, 0x52, 0x41, 0x4d, 0x45, 0x5f, 0x46, 0x49, 0x47, 0x48, 0x54, 0x45, 0x52,
	0x10, 0x4d, 0x12, 0x11, 0x0a, 0x0d, 0x46, 0x52, 0x41, 0x4d, 0x45, 0x5f, 0x46, 0x52, 0x49, 0x47,
	0x41, 0x54, 0x45, 0x10, 0x4e, 0x12, 0x11, 0x0a, 0x0d, 0x46, 0x52, 0x41, 0x4d, 0x45, 0x5f, 0x53,
	0x48, 0x55, 0x54, 0x54, 0x4c, 0x45, 0x10, 0x4f, 0x12, 0x12, 0x0a, 0x0e, 0x46, 0x52, 0x41, 0x4d,
	0x45, 0x5f, 0x45, 0x58, 0x50, 0x4c, 0x4f, 0x52, 0x45, 0x52, 0x10, 0x50, 0x12, 0x0f, 0x0a, 0x0b,
	0x46, 0x52, 0x41, 0x4d, 0x45, 0x5f, 0x4d, 0x49, 0x4e, 0x45, 0x52, 0x10, 0x51, 0x12, 0x19, 0x0a,
	0x15, 0x46, 0x52, 0x41, 0x4d, 0x45, 0x5f, 0x4c, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x46, 0x52, 0x45,
	0x49, 0x47, 0x48, 0x54, 0x45, 0x52, 0x10, 0x52, 0x12, 0x19, 0x0a, 0x15, 0x46, 0x52, 0x41, 0x4d,
	0x45, 0x5f, 0x48, 0x45, 0x41, 0x56, 0x59, 0x5f, 0x46, 0x52, 0x45, 0x49, 0x47, 0x48, 0x54, 0x45,
	0x52, 0x10, 0x53, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x52, 0x41, 0x4d, 0x45, 0x5f, 0x54, 0x52, 0x41,
	0x4e, 0x53, 0x50, 0x4f, 0x52, 0x54, 0x10, 0x54, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x52, 0x41, 0x4d,
	0x45, 0x5f, 0x44, 0x45, 0x53, 0x54, 0x52, 0x4f, 0x59, 0x45, 0x52, 0x10, 0x55, 0x12, 0x11, 0x0a,
	0x0d, 0x46, 0x52, 0x41, 0x4d, 0x45, 0x5f, 0x43, 0x52, 0x55, 0x49, 0x53, 0x45, 0x52, 0x10, 0x56,
	0x12, 0x11, 0x0a, 0x0d, 0x46, 0x52, 0x41, 0x4d, 0x45, 0x5f, 0x43, 0x41, 0x52, 0x52, 0x49, 0x45,
	0x52, 0x10, 0x57, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x45, 0x41, 0x43, 0x54, 0x4f, 0x52, 0x5f, 0x53,
	0x4f, 0x4c, 0x41, 0x52, 0x5f, 0x49, 0x10, 0x58, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x45, 0x41, 0x43,
	0x54, 0x4f, 0x52, 0x5f, 0x46, 0x55, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x10, 0x59, 0x12, 0x15,
	0x0a, 0x11, 0x52, 0x45, 0x41, 0x43, 0x54, 0x4f, 0x52, 0x5f, 0x46, 0x49, 0x53, 0x53, 0x49, 0x4f,
	0x4e, 0x5f, 0x49, 0x10, 0x5a, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x45, 0x41, 0x43, 0x54, 0x4f, 0x52,
	0x5f, 0x43, 0x48, 0x45, 0x4d, 0x49, 0x43, 0x41, 0x4c, 0x5f, 0x49, 0x10, 0x5b, 0x12, 0x18, 0x0a,
	0x14, 0x52, 0x45, 0x41, 0x43, 0x54, 0x4f, 0x52, 0x5f, 0x41, 0x4e, 0x54, 0x49, 0x4d, 0x41, 0x54,
	0x54, 0x45, 0x52, 0x5f, 0x49, 0x10, 0x5c, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x4e, 0x47, 0x49, 0x4e,
	0x45, 0x5f, 0x49, 0x4d, 0x50, 0x55, 0x4c, 0x53, 0x45, 0x5f, 0x44, 0x52, 0x49, 0x56, 0x45, 0x5f,
	0x49, 0x10, 0x5d, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x4e, 0x47, 0x49, 0x4e, 0x45, 0x5f, 0x49, 0x4f,
	0x4e, 0x5f, 0x44, 0x52, 0x49, 0x56, 0x45, 0x5f, 0x49, 0x10, 0x5e, 0x12, 0x17, 0x0a, 0x13, 0x45,
	0x4e, 0x47, 0x49, 0x4e, 0x45, 0x5f, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x52, 0x49, 0x56, 0x45, 0x5f,
	0x49, 0x49, 0x10, 0x5f, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x4e, 0x47, 0x49, 0x4e, 0x45, 0x5f, 0x48,
	0x59, 0x50, 0x45, 0x52, 0x5f, 0x44, 0x52, 0x49, 0x56, 0x45, 0x5f, 0x49, 0x10, 0x60, 0x12, 0x1e,
	0x0a, 0x1a, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x4d, 0x49, 0x4e, 0x45, 0x52, 0x41, 0x4c,
	0x5f, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x4f, 0x52, 0x5f, 0x49, 0x10, 0x61, 0x12, 0x1a,
	0x0a, 0x16, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x47, 0x41, 0x53, 0x5f, 0x50, 0x52, 0x4f,
	0x43, 0x45, 0x53, 0x53, 0x4f, 0x52, 0x5f, 0x49, 0x10, 0x62, 0x12, 0x17, 0x0a, 0x13, 0x4d, 0x4f,
	0x44, 0x55, 0x4c, 0x45, 0x5f, 0x43, 0x41, 0x52, 0x47, 0x4f, 0x5f, 0x48, 0x4f, 0x4c, 0x44, 0x5f,
	0x49, 0x10, 0x63, 0x12, 0x18, 0x0a, 0x14, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x43, 0x41,
	0x52, 0x47, 0x4f, 0x5f, 0x48, 0x4f, 0x4c, 0x44, 0x5f, 0x49, 0x49, 0x10, 0x64, 0x12, 0x19, 0x0a,
	0x15, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x43, 0x41, 0x52, 0x47, 0x4f, 0x5f, 0x48, 0x4f,
	0x4c, 0x44, 0x5f, 0x49, 0x49, 0x49, 0x10, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x4d, 0x4f, 0x44, 0x55,
	0x4c, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x57, 0x5f, 0x51, 0x55, 0x41, 0x52, 0x54, 0x45, 0x52, 0x53,
	0x5f, 0x49, 0x10, 0x66, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x45,
	0x4e, 0x56, 0x4f, 0x59, 0x5f, 0x51, 0x55, 0x41, 0x52, 0x54, 0x45, 0x52, 0x53, 0x5f, 0x49, 0x10,
	0x67, 0x12, 0x1c, 0x0a, 0x18, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x50, 0x41, 0x53, 0x53,
	0x45, 0x4e, 0x47, 0x45, 0x52, 0x5f, 0x43, 0x41, 0x42, 0x49, 0x4e, 0x5f, 0x49, 0x10, 0x68, 0x12,
	0x1b, 0x0a, 0x17, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x4d, 0x49, 0x43, 0x52, 0x4f, 0x5f,
	0x52, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x52, 0x59, 0x5f, 0x49, 0x10, 0x69, 0x12, 0x18, 0x0a, 0x14,
	0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x53, 0x43, 0x49, 0x45, 0x4e, 0x43, 0x45, 0x5f, 0x4c,
	0x41, 0x42, 0x5f, 0x49, 0x10, 0x6a, 0x12, 0x17, 0x0a, 0x13, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45,
	0x5f, 0x4a, 0x55, 0x4d, 0x50, 0x5f, 0x44, 0x52, 0x49, 0x56, 0x45, 0x5f, 0x49, 0x10, 0x6b, 0x12,
	0x18, 0x0a, 0x14, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x4a, 0x55, 0x4d, 0x50, 0x5f, 0x44,
	0x52, 0x49, 0x56, 0x45, 0x5f, 0x49, 0x49, 0x10, 0x6c, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x4f, 0x44,
	0x55, 0x4c, 0x45, 0x5f, 0x4a, 0x55, 0x4d, 0x50, 0x5f, 0x44, 0x52, 0x49, 0x56, 0x45, 0x5f, 0x49,
	0x49, 0x49, 0x10, 0x6d, 0x12, 0x17, 0x0a, 0x13, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x57,
	0x41, 0x52, 0x50, 0x5f, 0x44, 0x52, 0x49, 0x56, 0x45, 0x5f, 0x49, 0x10, 0x6e, 0x12, 0x18, 0x0a,
	0x14, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x57, 0x41, 0x52, 0x50, 0x5f, 0x44, 0x52, 0x49,
	0x56, 0x45, 0x5f, 0x49, 0x49, 0x10, 0x6f, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x4f, 0x44, 0x55, 0x4c,
	0x45, 0x5f, 0x57, 0x41, 0x52, 0x50, 0x5f, 0x44, 0x52, 0x49, 0x56, 0x45, 0x5f, 0x49, 0x49, 0x49,
	0x10, 0x70, 0x12, 0x1d, 0x0a, 0x19, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x53, 0x48, 0x49,
	0x45, 0x4c, 0x44, 0x5f, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x49, 0x10,
	0x71, 0x12, 0x1e, 0x0a, 0x1a, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x53, 0x48, 0x49, 0x45,
	0x4c, 0x44, 0x5f, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x49, 0x49, 0x10,
	0x72, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x4f, 0x52, 0x45, 0x5f,
	0x52, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x52, 0x59, 0x5f, 0x49, 0x10, 0x73, 0x12, 0x1a, 0x0a, 0x16,
	0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x46, 0x55, 0x45, 0x4c, 0x5f, 0x52, 0x45, 0x46, 0x49,
	0x4e, 0x45, 0x52, 0x59, 0x5f, 0x49, 0x10, 0x74, 0x12, 0x16, 0x0a, 0x12, 0x4d, 0x4f, 0x55, 0x4e,
	0x54, 0x5f, 0x47, 0x41, 0x53, 0x5f, 0x53, 0x49, 0x50, 0x48, 0x4f, 0x4e, 0x5f, 0x49, 0x10, 0x75,
	0x12, 0x17, 0x0a, 0x13, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x47, 0x41, 0x53, 0x5f, 0x53, 0x49,
	0x50, 0x48, 0x4f, 0x4e, 0x5f, 0x49, 0x49, 0x10, 0x76, 0x12, 0x18, 0x0a, 0x14, 0x4d, 0x4f, 0x55,
	0x4e, 0x54, 0x5f, 0x47, 0x41, 0x53, 0x5f, 0x53, 0x49, 0x50, 0x48, 0x4f, 0x4e, 0x5f, 0x49, 0x49,
	0x49, 0x10, 0x77, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x53, 0x55, 0x52,
	0x56, 0x45, 0x59, 0x4f, 0x52, 0x5f, 0x49, 0x10, 0x78, 0x12, 0x15, 0x0a, 0x11, 0x4d, 0x4f, 0x55,
	0x4e, 0x54, 0x5f, 0x53, 0x55, 0x52, 0x56, 0x45, 0x59, 0x4f, 0x52, 0x5f, 0x49, 0x49, 0x10, 0x79,
	0x12, 0x16, 0x0a, 0x12, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x53, 0x55, 0x52, 0x56, 0x45, 0x59,
	0x4f, 0x52, 0x5f, 0x49, 0x49, 0x49, 0x10, 0x7a, 0x12, 0x18, 0x0a, 0x14, 0x4d, 0x4f, 0x55, 0x4e,
	0x54, 0x5f, 0x53, 0x45, 0x4e, 0x53, 0x4f, 0x52, 0x5f, 0x41, 0x52, 0x52, 0x41, 0x59, 0x5f, 0x49,
	0x10, 0x7b, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x53, 0x45, 0x4e, 0x53,
	0x4f, 0x52, 0x5f, 0x41, 0x52, 0x52, 0x41, 0x59, 0x5f, 0x49, 0x49, 0x10, 0x7c, 0x12, 0x1a, 0x0a,
	0x16, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x53, 0x45, 0x4e, 0x53, 0x4f, 0x52, 0x5f, 0x41, 0x52,
	0x52, 0x41, 0x59, 0x5f, 0x49, 0x49, 0x49, 0x10, 0x7d, 0x12, 0x18, 0x0a, 0x14, 0x4d, 0x4f, 0x55,
	0x4e, 0x54, 0x5f, 0x4d, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x4c, 0x41, 0x53, 0x45, 0x52, 0x5f,
	0x49, 0x10, 0x7e, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x4d, 0x49, 0x4e,
	0x49, 0x4e, 0x47, 0x5f, 0x4c, 0x41, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x49, 0x10, 0x7f, 0x12, 0x1b,
	0x0a, 0x16, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x4d, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x4c,
	0x41, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x49, 0x49, 0x10, 0x80, 0x01, 0x12, 0x19, 0x0a, 0x14, 0x4d,
	0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x4c, 0x41, 0x53, 0x45, 0x52, 0x5f, 0x43, 0x41, 0x4e, 0x4e, 0x4f,
	0x4e, 0x5f, 0x49, 0x10, 0x81, 0x01, 0x12, 0x1d, 0x0a, 0x18, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f,
	0x4d, 0x49, 0x53, 0x53, 0x49, 0x4c, 0x45, 0x5f, 0x4c, 0x41, 0x55, 0x4e, 0x43, 0x48, 0x45, 0x52,
	0x5f, 0x49, 0x10, 0x82, 0x01, 0x12, 0x13, 0x0a, 0x0e, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x54,
	0x55, 0x52, 0x52, 0x45, 0x54, 0x5f, 0x49, 0x10, 0x83, 0x01, 0x12, 0x0f, 0x0a, 0x0a, 0x53, 0x48,
	0x49, 0x50, 0x5f, 0x50, 0x52, 0x4f, 0x42, 0x45, 0x10, 0x84, 0x01, 0x12, 0x16, 0x0a, 0x11, 0x53,
	0x48, 0x49, 0x50, 0x5f, 0x4d, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x44, 0x52, 0x4f, 0x4e, 0x45,
	0x10, 0x85, 0x01, 0x12, 0x16, 0x0a, 0x11, 0x53, 0x48, 0x49, 0x50, 0x5f, 0x53, 0x49, 0x50, 0x48,
	0x4f, 0x4e, 0x5f, 0x44, 0x52, 0x4f, 0x4e, 0x45, 0x10, 0x86, 0x01, 0x12, 0x15, 0x0a, 0x10, 0x53,
	0x48, 0x49, 0x50, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x43, 0x45, 0x50, 0x54, 0x4f, 0x52, 0x10,
	0x87, 0x01, 0x12, 0x16, 0x0a, 0x11, 0x53, 0x48, 0x49, 0x50, 0x5f, 0x4c, 0x49, 0x47, 0x48, 0x54,
	0x5f, 0x48, 0x41, 0x55, 0x4c, 0x45, 0x52, 0x10, 0x88, 0x01, 0x12, 0x19, 0x0a, 0x14, 0x53, 0x48,
	0x49, 0x50, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x46, 0x52, 0x49, 0x47, 0x41,
	0x54, 0x45, 0x10, 0x89, 0x01, 0x12, 0x12, 0x0a, 0x0d, 0x53, 0x48, 0x49, 0x50, 0x5f, 0x45, 0x58,
	0x50, 0x4c, 0x4f, 0x52, 0x45, 0x52, 0x10, 0x8a, 0x01, 0x12, 0x19, 0x0a, 0x14, 0x53, 0x48, 0x49,
	0x50, 0x5f, 0x48, 0x45, 0x41, 0x56, 0x59, 0x5f, 0x46, 0x52, 0x45, 0x49, 0x47, 0x48, 0x54, 0x45,
	0x52, 0x10, 0x8b, 0x01, 0x12, 0x17, 0x0a, 0x12, 0x53, 0x48, 0x49, 0x50, 0x5f, 0x4c, 0x49, 0x47,
	0x48, 0x54, 0x5f, 0x53, 0x48, 0x55, 0x54, 0x54, 0x4c, 0x45, 0x10, 0x8c, 0x01, 0x12, 0x13, 0x0a,
	0x0e, 0x53, 0x48, 0x49, 0x50, 0x5f, 0x4f, 0x52, 0x45, 0x5f, 0x48, 0x4f, 0x55, 0x4e, 0x44, 0x10,
	0x8d, 0x01, 0x12, 0x1c, 0x0a, 0x17, 0x53, 0x48, 0x49, 0x50, 0x5f, 0x52, 0x45, 0x46, 0x49, 0x4e,
	0x49, 0x4e, 0x47, 0x5f, 0x46, 0x52, 0x45, 0x49, 0x47, 0x48, 0x54, 0x45, 0x52, 0x10, 0x8e, 0x01,
	0x12, 0x12, 0x0a, 0x0d, 0x53, 0x48, 0x49, 0x50, 0x5f, 0x53, 0x55, 0x52, 0x56, 0x45, 0x59, 0x4f,
	0x52, 0x10, 0x8f, 0x01, 0x42, 0x49, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x6e, 0x6f, 0x6b, 0x6f, 0x74, 0x74, 0x2f, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x74, 0x72, 0x61, 0x64, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x0f,
	0x47, 0x72, 0x70, 0x63, 0x53, 0x70, 0x61, 0x63, 0x65, 0x74, 0x72, 0x61, 0x64, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_trade_proto_rawDescOnce sync.Once
	file_trade_proto_rawDescData = file_trade_proto_rawDesc
)

func file_trade_proto_rawDescGZIP() []byte {
	file_trade_proto_rawDescOnce.Do(func() {
		file_trade_proto_rawDescData = protoimpl.X.CompressGZIP(file_trade_proto_rawDescData)
	})
	return file_trade_proto_rawDescData
}

var file_trade_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_trade_proto_goTypes = []any{
	(TradeItem)(0), // 0: proto.TradeItem
}
var file_trade_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_trade_proto_init() }
func file_trade_proto_init() {
	if File_trade_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_trade_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_trade_proto_goTypes,
		DependencyIndexes: file_trade_proto_depIdxs,
		EnumInfos:         file_trade_proto_enumTypes,
	}.Build()
	File_trade_proto = out.File
	file_trade_proto_rawDesc = nil
	file_trade_proto_goTypes = nil
	file_trade_proto_depIdxs = nil
}
