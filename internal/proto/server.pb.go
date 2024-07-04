// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.20.1
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

type ServerStatusReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version       string                            `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	LastReset     *timestamppb.Timestamp            `protobuf:"bytes,2,opt,name=last_reset,json=lastReset,proto3" json:"last_reset,omitempty"`
	NextReset     *timestamppb.Timestamp            `protobuf:"bytes,3,opt,name=next_reset,json=nextReset,proto3" json:"next_reset,omitempty"`
	GlobalStats   *ServerStatusReply_GlobalStats    `protobuf:"bytes,4,opt,name=global_stats,json=globalStats,proto3" json:"global_stats,omitempty"`
	Announcements []*ServerStatusReply_Announcement `protobuf:"bytes,5,rep,name=announcements,proto3" json:"announcements,omitempty"`
}

func (x *ServerStatusReply) Reset() {
	*x = ServerStatusReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerStatusReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerStatusReply) ProtoMessage() {}

func (x *ServerStatusReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ServerStatusReply.ProtoReflect.Descriptor instead.
func (*ServerStatusReply) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{0}
}

func (x *ServerStatusReply) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ServerStatusReply) GetLastReset() *timestamppb.Timestamp {
	if x != nil {
		return x.LastReset
	}
	return nil
}

func (x *ServerStatusReply) GetNextReset() *timestamppb.Timestamp {
	if x != nil {
		return x.NextReset
	}
	return nil
}

func (x *ServerStatusReply) GetGlobalStats() *ServerStatusReply_GlobalStats {
	if x != nil {
		return x.GlobalStats
	}
	return nil
}

func (x *ServerStatusReply) GetAnnouncements() []*ServerStatusReply_Announcement {
	if x != nil {
		return x.Announcements
	}
	return nil
}

type CurrentAgentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Credits      int64  `protobuf:"varint,2,opt,name=credits,proto3" json:"credits,omitempty"`
	Headquarters string `protobuf:"bytes,3,opt,name=headquarters,proto3" json:"headquarters,omitempty"`
	ShipCount    int64  `protobuf:"varint,4,opt,name=shipCount,proto3" json:"shipCount,omitempty"`
}

func (x *CurrentAgentReply) Reset() {
	*x = CurrentAgentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrentAgentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentAgentReply) ProtoMessage() {}

func (x *CurrentAgentReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CurrentAgentReply.ProtoReflect.Descriptor instead.
func (*CurrentAgentReply) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{1}
}

func (x *CurrentAgentReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CurrentAgentReply) GetCredits() int64 {
	if x != nil {
		return x.Credits
	}
	return 0
}

func (x *CurrentAgentReply) GetHeadquarters() string {
	if x != nil {
		return x.Headquarters
	}
	return ""
}

func (x *CurrentAgentReply) GetShipCount() int64 {
	if x != nil {
		return x.ShipCount
	}
	return 0
}

type ServerStatusReply_GlobalStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Agents    int64 `protobuf:"varint,1,opt,name=agents,proto3" json:"agents,omitempty"`
	Ships     int64 `protobuf:"varint,2,opt,name=ships,proto3" json:"ships,omitempty"`
	Waypoints int64 `protobuf:"varint,3,opt,name=waypoints,proto3" json:"waypoints,omitempty"`
	Systems   int64 `protobuf:"varint,4,opt,name=systems,proto3" json:"systems,omitempty"`
}

func (x *ServerStatusReply_GlobalStats) Reset() {
	*x = ServerStatusReply_GlobalStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerStatusReply_GlobalStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerStatusReply_GlobalStats) ProtoMessage() {}

func (x *ServerStatusReply_GlobalStats) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ServerStatusReply_GlobalStats.ProtoReflect.Descriptor instead.
func (*ServerStatusReply_GlobalStats) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ServerStatusReply_GlobalStats) GetAgents() int64 {
	if x != nil {
		return x.Agents
	}
	return 0
}

func (x *ServerStatusReply_GlobalStats) GetShips() int64 {
	if x != nil {
		return x.Ships
	}
	return 0
}

func (x *ServerStatusReply_GlobalStats) GetWaypoints() int64 {
	if x != nil {
		return x.Waypoints
	}
	return 0
}

func (x *ServerStatusReply_GlobalStats) GetSystems() int64 {
	if x != nil {
		return x.Systems
	}
	return 0
}

type ServerStatusReply_Announcement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Body  string `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *ServerStatusReply_Announcement) Reset() {
	*x = ServerStatusReply_Announcement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerStatusReply_Announcement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerStatusReply_Announcement) ProtoMessage() {}

func (x *ServerStatusReply_Announcement) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ServerStatusReply_Announcement.ProtoReflect.Descriptor instead.
func (*ServerStatusReply_Announcement) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ServerStatusReply_Announcement) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ServerStatusReply_Announcement) GetBody() string {
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
	0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x03, 0x0a, 0x11, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x65,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x65, 0x74, 0x12, 0x39,
	0x0a, 0x0a, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x6e, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x65, 0x74, 0x12, 0x47, 0x0a, 0x0c, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x0b, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x12, 0x4b, 0x0a, 0x0d, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x0d, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a,
	0x73, 0x0a, 0x0b, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x69, 0x70, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x68, 0x69, 0x70, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x77, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x77, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x73, 0x1a, 0x38, 0x0a, 0x0c, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x83,
	0x01, 0x0a, 0x11, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x64,
	0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x68, 0x65, 0x61, 0x64, 0x71, 0x75, 0x61, 0x72, 0x74, 0x65,
	0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x68, 0x65, 0x61, 0x64, 0x71, 0x75,
	0x61, 0x72, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x68, 0x69, 0x70, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x68, 0x69, 0x70, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x32, 0xd7, 0x01, 0x0a, 0x13, 0x53, 0x70, 0x61, 0x63, 0x65, 0x54, 0x72,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x04,
	0x50, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x43, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x43, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x30,
	0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x6e,
	0x6f, 0x6b, 0x6f, 0x74, 0x74, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x74, 0x72, 0x61, 0x64, 0x65,
	0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_server_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_server_proto_goTypes = []any{
	(*ServerStatusReply)(nil),              // 0: proto.ServerStatusReply
	(*CurrentAgentReply)(nil),              // 1: proto.CurrentAgentReply
	(*ServerStatusReply_GlobalStats)(nil),  // 2: proto.ServerStatusReply.GlobalStats
	(*ServerStatusReply_Announcement)(nil), // 3: proto.ServerStatusReply.Announcement
	(*timestamppb.Timestamp)(nil),          // 4: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),                  // 5: google.protobuf.Empty
}
var file_server_proto_depIdxs = []int32{
	4, // 0: proto.ServerStatusReply.last_reset:type_name -> google.protobuf.Timestamp
	4, // 1: proto.ServerStatusReply.next_reset:type_name -> google.protobuf.Timestamp
	2, // 2: proto.ServerStatusReply.global_stats:type_name -> proto.ServerStatusReply.GlobalStats
	3, // 3: proto.ServerStatusReply.announcements:type_name -> proto.ServerStatusReply.Announcement
	5, // 4: proto.SpaceTradersService.Ping:input_type -> google.protobuf.Empty
	5, // 5: proto.SpaceTradersService.GetServerStatus:input_type -> google.protobuf.Empty
	5, // 6: proto.SpaceTradersService.GetCurrentAgent:input_type -> google.protobuf.Empty
	5, // 7: proto.SpaceTradersService.Ping:output_type -> google.protobuf.Empty
	0, // 8: proto.SpaceTradersService.GetServerStatus:output_type -> proto.ServerStatusReply
	1, // 9: proto.SpaceTradersService.GetCurrentAgent:output_type -> proto.CurrentAgentReply
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_server_proto_init() }
func file_server_proto_init() {
	if File_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ServerStatusReply); i {
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
			switch v := v.(*CurrentAgentReply); i {
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
			switch v := v.(*ServerStatusReply_GlobalStats); i {
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
			switch v := v.(*ServerStatusReply_Announcement); i {
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
			NumMessages:   4,
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
