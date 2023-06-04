// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: grpc/proto/base_ball_api.proto

package protobuf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FetchPlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
}

func (x *FetchPlayerRequest) Reset() {
	*x = FetchPlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_base_ball_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchPlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchPlayerRequest) ProtoMessage() {}

func (x *FetchPlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_base_ball_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchPlayerRequest.ProtoReflect.Descriptor instead.
func (*FetchPlayerRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_base_ball_api_proto_rawDescGZIP(), []int{0}
}

func (x *FetchPlayerRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

type FetchPlayerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player *Player `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
}

func (x *FetchPlayerResponse) Reset() {
	*x = FetchPlayerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_base_ball_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchPlayerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchPlayerResponse) ProtoMessage() {}

func (x *FetchPlayerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_base_ball_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchPlayerResponse.ProtoReflect.Descriptor instead.
func (*FetchPlayerResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_base_ball_api_proto_rawDescGZIP(), []int{1}
}

func (x *FetchPlayerResponse) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

type CreatePlayersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	UniformNumber int64  `protobuf:"varint,2,opt,name=uniform_number,json=uniformNumber,proto3" json:"uniform_number,omitempty"`
	AtBats        int64  `protobuf:"varint,3,opt,name=at_bats,json=atBats,proto3" json:"at_bats,omitempty"`
	Hits          int64  `protobuf:"varint,4,opt,name=hits,proto3" json:"hits,omitempty"`
	Walks         int64  `protobuf:"varint,5,opt,name=walks,proto3" json:"walks,omitempty"`
	HomeRuns      int64  `protobuf:"varint,6,opt,name=home_runs,json=homeRuns,proto3" json:"home_runs,omitempty"`
	RunsBattedIn  int64  `protobuf:"varint,7,opt,name=runs_batted_in,json=runsBattedIn,proto3" json:"runs_batted_in,omitempty"`
}

func (x *CreatePlayersRequest) Reset() {
	*x = CreatePlayersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_base_ball_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePlayersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePlayersRequest) ProtoMessage() {}

func (x *CreatePlayersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_base_ball_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePlayersRequest.ProtoReflect.Descriptor instead.
func (*CreatePlayersRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_base_ball_api_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePlayersRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreatePlayersRequest) GetUniformNumber() int64 {
	if x != nil {
		return x.UniformNumber
	}
	return 0
}

func (x *CreatePlayersRequest) GetAtBats() int64 {
	if x != nil {
		return x.AtBats
	}
	return 0
}

func (x *CreatePlayersRequest) GetHits() int64 {
	if x != nil {
		return x.Hits
	}
	return 0
}

func (x *CreatePlayersRequest) GetWalks() int64 {
	if x != nil {
		return x.Walks
	}
	return 0
}

func (x *CreatePlayersRequest) GetHomeRuns() int64 {
	if x != nil {
		return x.HomeRuns
	}
	return 0
}

func (x *CreatePlayersRequest) GetRunsBattedIn() int64 {
	if x != nil {
		return x.RunsBattedIn
	}
	return 0
}

type CreatePlayerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player *Player `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
}

func (x *CreatePlayerResponse) Reset() {
	*x = CreatePlayerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_base_ball_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePlayerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePlayerResponse) ProtoMessage() {}

func (x *CreatePlayerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_base_ball_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePlayerResponse.ProtoReflect.Descriptor instead.
func (*CreatePlayerResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_base_ball_api_proto_rawDescGZIP(), []int{3}
}

func (x *CreatePlayerResponse) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	UniformNumber int64  `protobuf:"varint,3,opt,name=uniform_number,json=uniformNumber,proto3" json:"uniform_number,omitempty"`
	AtBats        int64  `protobuf:"varint,4,opt,name=at_bats,json=atBats,proto3" json:"at_bats,omitempty"`
	Hits          int64  `protobuf:"varint,5,opt,name=hits,proto3" json:"hits,omitempty"`
	Walks         int64  `protobuf:"varint,6,opt,name=walks,proto3" json:"walks,omitempty"`
	HomeRuns      int64  `protobuf:"varint,7,opt,name=home_runs,json=homeRuns,proto3" json:"home_runs,omitempty"`
	RunsBattedIn  int64  `protobuf:"varint,8,opt,name=runs_batted_in,json=runsBattedIn,proto3" json:"runs_batted_in,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_base_ball_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_base_ball_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_grpc_proto_base_ball_api_proto_rawDescGZIP(), []int{4}
}

func (x *Player) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Player) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Player) GetUniformNumber() int64 {
	if x != nil {
		return x.UniformNumber
	}
	return 0
}

func (x *Player) GetAtBats() int64 {
	if x != nil {
		return x.AtBats
	}
	return 0
}

func (x *Player) GetHits() int64 {
	if x != nil {
		return x.Hits
	}
	return 0
}

func (x *Player) GetWalks() int64 {
	if x != nil {
		return x.Walks
	}
	return 0
}

func (x *Player) GetHomeRuns() int64 {
	if x != nil {
		return x.HomeRuns
	}
	return 0
}

func (x *Player) GetRunsBattedIn() int64 {
	if x != nil {
		return x.RunsBattedIn
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_base_ball_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_base_ball_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_grpc_proto_base_ball_api_proto_rawDescGZIP(), []int{5}
}

var File_grpc_proto_base_ball_api_proto protoreflect.FileDescriptor

var file_grpc_proto_base_ball_api_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x5f, 0x62, 0x61, 0x6c, 0x6c, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a,
	0x12, 0x46, 0x65, 0x74, 0x63, 0x68, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x36, 0x0a, 0x13, 0x46, 0x65, 0x74, 0x63, 0x68, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0xd7, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x75, 0x6e, 0x69, 0x66, 0x6f, 0x72, 0x6d,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x75,
	0x6e, 0x69, 0x66, 0x6f, 0x72, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07,
	0x61, 0x74, 0x5f, 0x62, 0x61, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61,
	0x74, 0x42, 0x61, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x74, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x68, 0x69, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x61, 0x6c,
	0x6b, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x77, 0x61, 0x6c, 0x6b, 0x73, 0x12,
	0x1b, 0x0a, 0x09, 0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x72, 0x75, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x68, 0x6f, 0x6d, 0x65, 0x52, 0x75, 0x6e, 0x73, 0x12, 0x24, 0x0a, 0x0e,
	0x72, 0x75, 0x6e, 0x73, 0x5f, 0x62, 0x61, 0x74, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x72, 0x75, 0x6e, 0x73, 0x42, 0x61, 0x74, 0x74, 0x65, 0x64,
	0x49, 0x6e, 0x22, 0x37, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0xd9, 0x01, 0x0a, 0x06,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x75, 0x6e,
	0x69, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0d, 0x75, 0x6e, 0x69, 0x66, 0x6f, 0x72, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x74, 0x5f, 0x62, 0x61, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x61, 0x74, 0x42, 0x61, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69,
	0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x68, 0x69, 0x74, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x77, 0x61, 0x6c, 0x6b, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x77,
	0x61, 0x6c, 0x6b, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x72, 0x75, 0x6e,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x68, 0x6f, 0x6d, 0x65, 0x52, 0x75, 0x6e,
	0x73, 0x12, 0x24, 0x0a, 0x0e, 0x72, 0x75, 0x6e, 0x73, 0x5f, 0x62, 0x61, 0x74, 0x74, 0x65, 0x64,
	0x5f, 0x69, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x72, 0x75, 0x6e, 0x73, 0x42,
	0x61, 0x74, 0x74, 0x65, 0x64, 0x49, 0x6e, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x32, 0xc6, 0x01, 0x0a, 0x0b, 0x42, 0x61, 0x73, 0x65, 0x42, 0x61, 0x6c, 0x6c, 0x41, 0x70, 0x69,
	0x12, 0x3a, 0x0a, 0x0b, 0x46, 0x65, 0x74, 0x63, 0x68, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12,
	0x13, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0c,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x07,
	0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_proto_base_ball_api_proto_rawDescOnce sync.Once
	file_grpc_proto_base_ball_api_proto_rawDescData = file_grpc_proto_base_ball_api_proto_rawDesc
)

func file_grpc_proto_base_ball_api_proto_rawDescGZIP() []byte {
	file_grpc_proto_base_ball_api_proto_rawDescOnce.Do(func() {
		file_grpc_proto_base_ball_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_proto_base_ball_api_proto_rawDescData)
	})
	return file_grpc_proto_base_ball_api_proto_rawDescData
}

var file_grpc_proto_base_ball_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_grpc_proto_base_ball_api_proto_goTypes = []interface{}{
	(*FetchPlayerRequest)(nil),   // 0: FetchPlayerRequest
	(*FetchPlayerResponse)(nil),  // 1: FetchPlayerResponse
	(*CreatePlayersRequest)(nil), // 2: CreatePlayersRequest
	(*CreatePlayerResponse)(nil), // 3: CreatePlayerResponse
	(*Player)(nil),               // 4: Player
	(*Empty)(nil),                // 5: Empty
	(*emptypb.Empty)(nil),        // 6: google.protobuf.Empty
}
var file_grpc_proto_base_ball_api_proto_depIdxs = []int32{
	4, // 0: FetchPlayerResponse.player:type_name -> Player
	4, // 1: CreatePlayerResponse.player:type_name -> Player
	0, // 2: BaseBallApi.FetchPlayer:input_type -> FetchPlayerRequest
	2, // 3: BaseBallApi.CreatePlayer:input_type -> CreatePlayersRequest
	6, // 4: BaseBallApi.Crawler:input_type -> google.protobuf.Empty
	1, // 5: BaseBallApi.FetchPlayer:output_type -> FetchPlayerResponse
	3, // 6: BaseBallApi.CreatePlayer:output_type -> CreatePlayerResponse
	6, // 7: BaseBallApi.Crawler:output_type -> google.protobuf.Empty
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_grpc_proto_base_ball_api_proto_init() }
func file_grpc_proto_base_ball_api_proto_init() {
	if File_grpc_proto_base_ball_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_proto_base_ball_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchPlayerRequest); i {
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
		file_grpc_proto_base_ball_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchPlayerResponse); i {
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
		file_grpc_proto_base_ball_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePlayersRequest); i {
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
		file_grpc_proto_base_ball_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePlayerResponse); i {
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
		file_grpc_proto_base_ball_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
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
		file_grpc_proto_base_ball_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_grpc_proto_base_ball_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_proto_base_ball_api_proto_goTypes,
		DependencyIndexes: file_grpc_proto_base_ball_api_proto_depIdxs,
		MessageInfos:      file_grpc_proto_base_ball_api_proto_msgTypes,
	}.Build()
	File_grpc_proto_base_ball_api_proto = out.File
	file_grpc_proto_base_ball_api_proto_rawDesc = nil
	file_grpc_proto_base_ball_api_proto_goTypes = nil
	file_grpc_proto_base_ball_api_proto_depIdxs = nil
}