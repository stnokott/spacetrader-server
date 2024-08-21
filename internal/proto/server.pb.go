// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: server.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ServerStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version       string                       `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	LastReset     *timestamppb.Timestamp       `protobuf:"bytes,2,opt,name=last_reset,json=lastReset,proto3" json:"last_reset,omitempty"`
	NextReset     *timestamppb.Timestamp       `protobuf:"bytes,3,opt,name=next_reset,json=nextReset,proto3" json:"next_reset,omitempty"`
	GlobalStats   *ServerStatus_GlobalStats    `protobuf:"bytes,4,opt,name=global_stats,json=globalStats,proto3" json:"global_stats,omitempty"`
	Announcements []*ServerStatus_Announcement `protobuf:"bytes,5,rep,name=announcements,proto3" json:"announcements,omitempty"`
}

func (x *ServerStatus) Reset() {
	*x = ServerStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerStatus) ProtoMessage() {}

func (x *ServerStatus) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerStatus.ProtoReflect.Descriptor instead.
func (*ServerStatus) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{0}
}

func (x *ServerStatus) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ServerStatus) GetLastReset() *timestamppb.Timestamp {
	if x != nil {
		return x.LastReset
	}
	return nil
}

func (x *ServerStatus) GetNextReset() *timestamppb.Timestamp {
	if x != nil {
		return x.NextReset
	}
	return nil
}

func (x *ServerStatus) GetGlobalStats() *ServerStatus_GlobalStats {
	if x != nil {
		return x.GlobalStats
	}
	return nil
}

func (x *ServerStatus) GetAnnouncements() []*ServerStatus_Announcement {
	if x != nil {
		return x.Announcements
	}
	return nil
}

type Agent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Credits         int64   `protobuf:"varint,2,opt,name=credits,proto3" json:"credits,omitempty"`
	Headquarters    string  `protobuf:"bytes,3,opt,name=headquarters,proto3" json:"headquarters,omitempty"`
	ShipCount       int32   `protobuf:"varint,4,opt,name=shipCount,proto3" json:"shipCount,omitempty"`
	StartingFaction Faction `protobuf:"varint,5,opt,name=startingFaction,proto3,enum=proto.Faction" json:"startingFaction,omitempty"`
}

func (x *Agent) Reset() {
	*x = Agent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Agent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Agent) ProtoMessage() {}

func (x *Agent) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Agent.ProtoReflect.Descriptor instead.
func (*Agent) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{1}
}

func (x *Agent) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Agent) GetCredits() int64 {
	if x != nil {
		return x.Credits
	}
	return 0
}

func (x *Agent) GetHeadquarters() string {
	if x != nil {
		return x.Headquarters
	}
	return ""
}

func (x *Agent) GetShipCount() int32 {
	if x != nil {
		return x.ShipCount
	}
	return 0
}

func (x *Agent) GetStartingFaction() Faction {
	if x != nil {
		return x.StartingFaction
	}
	return Faction_UNKNOWN_FACTION
}

type Fleet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ships []*Ship `protobuf:"bytes,1,rep,name=ships,proto3" json:"ships,omitempty"`
}

func (x *Fleet) Reset() {
	*x = Fleet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fleet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fleet) ProtoMessage() {}

func (x *Fleet) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fleet.ProtoReflect.Descriptor instead.
func (*Fleet) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{2}
}

func (x *Fleet) GetShips() []*Ship {
	if x != nil {
		return x.Ships
	}
	return nil
}

type GetShipCountInSystemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SystemName string `protobuf:"bytes,1,opt,name=system_name,json=systemName,proto3" json:"system_name,omitempty"`
}

func (x *GetShipCountInSystemRequest) Reset() {
	*x = GetShipCountInSystemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShipCountInSystemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShipCountInSystemRequest) ProtoMessage() {}

func (x *GetShipCountInSystemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShipCountInSystemRequest.ProtoReflect.Descriptor instead.
func (*GetShipCountInSystemRequest) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{3}
}

func (x *GetShipCountInSystemRequest) GetSystemName() string {
	if x != nil {
		return x.SystemName
	}
	return ""
}

type GetShipCountInSystemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	N int32 `protobuf:"varint,1,opt,name=n,proto3" json:"n,omitempty"`
}

func (x *GetShipCountInSystemResponse) Reset() {
	*x = GetShipCountInSystemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShipCountInSystemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShipCountInSystemResponse) ProtoMessage() {}

func (x *GetShipCountInSystemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShipCountInSystemResponse.ProtoReflect.Descriptor instead.
func (*GetShipCountInSystemResponse) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{4}
}

func (x *GetShipCountInSystemResponse) GetN() int32 {
	if x != nil {
		return x.N
	}
	return 0
}

type GetShipCoordinatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShipName string `protobuf:"bytes,1,opt,name=ship_name,json=shipName,proto3" json:"ship_name,omitempty"`
}

func (x *GetShipCoordinatesRequest) Reset() {
	*x = GetShipCoordinatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShipCoordinatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShipCoordinatesRequest) ProtoMessage() {}

func (x *GetShipCoordinatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShipCoordinatesRequest.ProtoReflect.Descriptor instead.
func (*GetShipCoordinatesRequest) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{5}
}

func (x *GetShipCoordinatesRequest) GetShipName() string {
	if x != nil {
		return x.ShipName
	}
	return ""
}

type GetShipCoordinatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *GetShipCoordinatesResponse) Reset() {
	*x = GetShipCoordinatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShipCoordinatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShipCoordinatesResponse) ProtoMessage() {}

func (x *GetShipCoordinatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShipCoordinatesResponse.ProtoReflect.Descriptor instead.
func (*GetShipCoordinatesResponse) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{6}
}

func (x *GetShipCoordinatesResponse) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *GetShipCoordinatesResponse) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type GetSystemsInRectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	System    *System `protobuf:"bytes,1,opt,name=system,proto3" json:"system,omitempty"`
	ShipCount int32   `protobuf:"varint,2,opt,name=ship_count,json=shipCount,proto3" json:"ship_count,omitempty"`
}

func (x *GetSystemsInRectResponse) Reset() {
	*x = GetSystemsInRectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSystemsInRectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSystemsInRectResponse) ProtoMessage() {}

func (x *GetSystemsInRectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSystemsInRectResponse.ProtoReflect.Descriptor instead.
func (*GetSystemsInRectResponse) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{7}
}

func (x *GetSystemsInRectResponse) GetSystem() *System {
	if x != nil {
		return x.System
	}
	return nil
}

func (x *GetSystemsInRectResponse) GetShipCount() int32 {
	if x != nil {
		return x.ShipCount
	}
	return 0
}

type Vector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Vector) Reset() {
	*x = Vector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vector) ProtoMessage() {}

func (x *Vector) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vector.ProtoReflect.Descriptor instead.
func (*Vector) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{8}
}

func (x *Vector) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Vector) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type Rect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start *Vector `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End   *Vector `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *Rect) Reset() {
	*x = Rect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rect) ProtoMessage() {}

func (x *Rect) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rect.ProtoReflect.Descriptor instead.
func (*Rect) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{9}
}

func (x *Rect) GetStart() *Vector {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *Rect) GetEnd() *Vector {
	if x != nil {
		return x.End
	}
	return nil
}

type ServerStatus_GlobalStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Agents    int64 `protobuf:"varint,1,opt,name=agents,proto3" json:"agents,omitempty"`
	Ships     int64 `protobuf:"varint,2,opt,name=ships,proto3" json:"ships,omitempty"`
	Waypoints int64 `protobuf:"varint,3,opt,name=waypoints,proto3" json:"waypoints,omitempty"`
	Systems   int64 `protobuf:"varint,4,opt,name=systems,proto3" json:"systems,omitempty"`
}

func (x *ServerStatus_GlobalStats) Reset() {
	*x = ServerStatus_GlobalStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerStatus_GlobalStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerStatus_GlobalStats) ProtoMessage() {}

func (x *ServerStatus_GlobalStats) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerStatus_GlobalStats.ProtoReflect.Descriptor instead.
func (*ServerStatus_GlobalStats) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ServerStatus_GlobalStats) GetAgents() int64 {
	if x != nil {
		return x.Agents
	}
	return 0
}

func (x *ServerStatus_GlobalStats) GetShips() int64 {
	if x != nil {
		return x.Ships
	}
	return 0
}

func (x *ServerStatus_GlobalStats) GetWaypoints() int64 {
	if x != nil {
		return x.Waypoints
	}
	return 0
}

func (x *ServerStatus_GlobalStats) GetSystems() int64 {
	if x != nil {
		return x.Systems
	}
	return 0
}

type ServerStatus_Announcement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Body  string `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *ServerStatus_Announcement) Reset() {
	*x = ServerStatus_Announcement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerStatus_Announcement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerStatus_Announcement) ProtoMessage() {}

func (x *ServerStatus_Announcement) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerStatus_Announcement.ProtoReflect.Descriptor instead.
func (*ServerStatus_Announcement) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ServerStatus_Announcement) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ServerStatus_Announcement) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

var File_server_proto protoreflect.FileDescriptor

var file_server_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0c, 0x67, 0x61, 0x6c, 0x61, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x03,
	0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x72, 0x65, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x65, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x65,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x65, 0x74, 0x12, 0x42,
	0x0a, 0x0c, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x0b, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x12, 0x46, 0x0a, 0x0d, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x41,
	0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0d, 0x61, 0x6e, 0x6e,
	0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x73, 0x0a, 0x0b, 0x47, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x69, 0x70, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x73, 0x68, 0x69, 0x70, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x77, 0x61, 0x79, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x77, 0x61, 0x79, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x1a,
	0x38, 0x0a, 0x0c, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0xb1, 0x01, 0x0a, 0x05, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x73, 0x12, 0x22, 0x0a, 0x0c, 0x68, 0x65, 0x61, 0x64, 0x71, 0x75, 0x61, 0x72, 0x74, 0x65, 0x72,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x68, 0x65, 0x61, 0x64, 0x71, 0x75, 0x61,
	0x72, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x68, 0x69, 0x70, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x68, 0x69, 0x70, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x46,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x46, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2a, 0x0a,
	0x05, 0x46, 0x6c, 0x65, 0x65, 0x74, 0x12, 0x21, 0x0a, 0x05, 0x73, 0x68, 0x69, 0x70, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x68,
	0x69, 0x70, 0x52, 0x05, 0x73, 0x68, 0x69, 0x70, 0x73, 0x22, 0x3e, 0x0a, 0x1b, 0x47, 0x65, 0x74,
	0x53, 0x68, 0x69, 0x70, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2c, 0x0a, 0x1c, 0x47, 0x65, 0x74,
	0x53, 0x68, 0x69, 0x70, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x6e, 0x22, 0x38, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x53, 0x68,
	0x69, 0x70, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x69, 0x70, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x38, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x53, 0x68, 0x69, 0x70, 0x43, 0x6f, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a,
	0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x22, 0x60, 0x0a, 0x18, 0x47,
	0x65, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x49, 0x6e, 0x52, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x73, 0x68, 0x69, 0x70, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x24, 0x0a,
	0x06, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x01, 0x79, 0x22, 0x4c, 0x0a, 0x04, 0x52, 0x65, 0x63, 0x74, 0x12, 0x23, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x12, 0x1f, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x03, 0x65, 0x6e,
	0x64, 0x32, 0x8f, 0x03, 0x0a, 0x0b, 0x53, 0x70, 0x61, 0x63, 0x65, 0x74, 0x72, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x36, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3e, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x37, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x12, 0x30, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x65, 0x65, 0x74, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46,
	0x6c, 0x65, 0x65, 0x74, 0x12, 0x59, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x68, 0x69, 0x70, 0x43,
	0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x69, 0x70, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x69, 0x70, 0x43, 0x6f, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x42, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x49, 0x6e, 0x52,
	0x65, 0x63, 0x74, 0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x63, 0x74,
	0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x73, 0x49, 0x6e, 0x52, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x30, 0x01, 0x42, 0x49, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x74, 0x6e, 0x6f, 0x6b, 0x6f, 0x74, 0x74, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x74, 0x72, 0x61, 0x64, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x0f, 0x47,
	0x72, 0x70, 0x63, 0x53, 0x70, 0x61, 0x63, 0x65, 0x74, 0x72, 0x61, 0x64, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_proto_rawDescOnce sync.Once
	file_server_proto_rawDescData = file_server_proto_rawDesc
)

func file_server_proto_rawDescGZIP() []byte {
	file_server_proto_rawDescOnce.Do(func() {
		file_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_proto_rawDescData)
	})
	return file_server_proto_rawDescData
}

var file_server_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_server_proto_goTypes = []any{
	(*ServerStatus)(nil),                 // 0: proto.ServerStatus
	(*Agent)(nil),                        // 1: proto.Agent
	(*Fleet)(nil),                        // 2: proto.Fleet
	(*GetShipCountInSystemRequest)(nil),  // 3: proto.GetShipCountInSystemRequest
	(*GetShipCountInSystemResponse)(nil), // 4: proto.GetShipCountInSystemResponse
	(*GetShipCoordinatesRequest)(nil),    // 5: proto.GetShipCoordinatesRequest
	(*GetShipCoordinatesResponse)(nil),   // 6: proto.GetShipCoordinatesResponse
	(*GetSystemsInRectResponse)(nil),     // 7: proto.GetSystemsInRectResponse
	(*Vector)(nil),                       // 8: proto.Vector
	(*Rect)(nil),                         // 9: proto.Rect
	(*ServerStatus_GlobalStats)(nil),     // 10: proto.ServerStatus.GlobalStats
	(*ServerStatus_Announcement)(nil),    // 11: proto.ServerStatus.Announcement
	(*timestamppb.Timestamp)(nil),        // 12: google.protobuf.Timestamp
	(Faction)(0),                         // 13: proto.Faction
	(*Ship)(nil),                         // 14: proto.Ship
	(*System)(nil),                       // 15: proto.System
	(*emptypb.Empty)(nil),                // 16: google.protobuf.Empty
}
var file_server_proto_depIdxs = []int32{
	12, // 0: proto.ServerStatus.last_reset:type_name -> google.protobuf.Timestamp
	12, // 1: proto.ServerStatus.next_reset:type_name -> google.protobuf.Timestamp
	10, // 2: proto.ServerStatus.global_stats:type_name -> proto.ServerStatus.GlobalStats
	11, // 3: proto.ServerStatus.announcements:type_name -> proto.ServerStatus.Announcement
	13, // 4: proto.Agent.startingFaction:type_name -> proto.Faction
	14, // 5: proto.Fleet.ships:type_name -> proto.Ship
	15, // 6: proto.GetSystemsInRectResponse.system:type_name -> proto.System
	8,  // 7: proto.Rect.start:type_name -> proto.Vector
	8,  // 8: proto.Rect.end:type_name -> proto.Vector
	16, // 9: proto.Spacetrader.Ping:input_type -> google.protobuf.Empty
	16, // 10: proto.Spacetrader.GetServerStatus:input_type -> google.protobuf.Empty
	16, // 11: proto.Spacetrader.GetCurrentAgent:input_type -> google.protobuf.Empty
	16, // 12: proto.Spacetrader.GetFleet:input_type -> google.protobuf.Empty
	5,  // 13: proto.Spacetrader.GetShipCoordinates:input_type -> proto.GetShipCoordinatesRequest
	9,  // 14: proto.Spacetrader.GetSystemsInRect:input_type -> proto.Rect
	16, // 15: proto.Spacetrader.Ping:output_type -> google.protobuf.Empty
	0,  // 16: proto.Spacetrader.GetServerStatus:output_type -> proto.ServerStatus
	1,  // 17: proto.Spacetrader.GetCurrentAgent:output_type -> proto.Agent
	2,  // 18: proto.Spacetrader.GetFleet:output_type -> proto.Fleet
	6,  // 19: proto.Spacetrader.GetShipCoordinates:output_type -> proto.GetShipCoordinatesResponse
	7,  // 20: proto.Spacetrader.GetSystemsInRect:output_type -> proto.GetSystemsInRectResponse
	15, // [15:21] is the sub-list for method output_type
	9,  // [9:15] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_server_proto_init() }
func file_server_proto_init() {
	if File_server_proto != nil {
		return
	}
	file_ship_proto_init()
	file_galaxy_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_server_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ServerStatus); i {
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
		file_server_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Agent); i {
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
		file_server_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Fleet); i {
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
		file_server_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetShipCountInSystemRequest); i {
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
		file_server_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetShipCountInSystemResponse); i {
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
		file_server_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetShipCoordinatesRequest); i {
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
		file_server_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetShipCoordinatesResponse); i {
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
		file_server_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetSystemsInRectResponse); i {
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
		file_server_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*Vector); i {
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
		file_server_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*Rect); i {
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
		file_server_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*ServerStatus_GlobalStats); i {
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
		file_server_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*ServerStatus_Announcement); i {
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
			RawDescriptor: file_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_proto_goTypes,
		DependencyIndexes: file_server_proto_depIdxs,
		MessageInfos:      file_server_proto_msgTypes,
	}.Build()
	File_server_proto = out.File
	file_server_proto_rawDesc = nil
	file_server_proto_goTypes = nil
	file_server_proto_depIdxs = nil
}
