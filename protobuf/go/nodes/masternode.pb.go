// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: protobuf/proto/masternode.proto

package nodes

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

//*
// enum Action is a rule of conduct the managers
type Action int32

const (
	Action_NOTHING Action = 0
	Action_UP      Action = 1
	Action_DOWN    Action = 2
)

// Enum value maps for Action.
var (
	Action_name = map[int32]string{
		0: "NOTHING",
		1: "UP",
		2: "DOWN",
	}
	Action_value = map[string]int32{
		"NOTHING": 0,
		"UP":      1,
		"DOWN":    2,
	}
)

func (x Action) Enum() *Action {
	p := new(Action)
	*p = x
	return p
}

func (x Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Action) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuf_proto_masternode_proto_enumTypes[0].Descriptor()
}

func (Action) Type() protoreflect.EnumType {
	return &file_protobuf_proto_masternode_proto_enumTypes[0]
}

func (x Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Action.Descriptor instead.
func (Action) EnumDescriptor() ([]byte, []int) {
	return file_protobuf_proto_masternode_proto_rawDescGZIP(), []int{0}
}

//*
// enum Type is a type of ruling objects
type Type int32

const (
	Type_TYPE_UNKNOWN Type = 0
	Type_TYPE_WORKER  Type = 1
	Type_TYPE_CHANNEL Type = 2
	Type_TYPE_MANAGER Type = 3
	Type_TYPE_NODE    Type = 4
	Type_TYPE_CLUSTER Type = 5
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "TYPE_UNKNOWN",
		1: "TYPE_WORKER",
		2: "TYPE_CHANNEL",
		3: "TYPE_MANAGER",
		4: "TYPE_NODE",
		5: "TYPE_CLUSTER",
	}
	Type_value = map[string]int32{
		"TYPE_UNKNOWN": 0,
		"TYPE_WORKER":  1,
		"TYPE_CHANNEL": 2,
		"TYPE_MANAGER": 3,
		"TYPE_NODE":    4,
		"TYPE_CLUSTER": 5,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuf_proto_masternode_proto_enumTypes[1].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_protobuf_proto_masternode_proto_enumTypes[1]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_protobuf_proto_masternode_proto_rawDescGZIP(), []int{1}
}

//*
// enum ChanType  is a type of channel
type ChanType int32

const (
	ChanType_CHAN_UNKNOWN        ChanType = 0 // when channel is not is existed
	ChanType_CHAN_STD_GO         ChanType = 1 // core go library channel aka make(chan <type>, length)
	ChanType_CHAN_STACK          ChanType = 2 //
	ChanType_CHAN_PRIORITY_QUEUE ChanType = 3 //
)

// Enum value maps for ChanType.
var (
	ChanType_name = map[int32]string{
		0: "CHAN_UNKNOWN",
		1: "CHAN_STD_GO",
		2: "CHAN_STACK",
		3: "CHAN_PRIORITY_QUEUE",
	}
	ChanType_value = map[string]int32{
		"CHAN_UNKNOWN":        0,
		"CHAN_STD_GO":         1,
		"CHAN_STACK":          2,
		"CHAN_PRIORITY_QUEUE": 3,
	}
)

func (x ChanType) Enum() *ChanType {
	p := new(ChanType)
	*p = x
	return p
}

func (x ChanType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChanType) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuf_proto_masternode_proto_enumTypes[2].Descriptor()
}

func (ChanType) Type() protoreflect.EnumType {
	return &file_protobuf_proto_masternode_proto_enumTypes[2]
}

func (x ChanType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChanType.Descriptor instead.
func (ChanType) EnumDescriptor() ([]byte, []int) {
	return file_protobuf_proto_masternode_proto_rawDescGZIP(), []int{2}
}

type ManagerAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   Type   `protobuf:"varint,1,opt,name=Type,json=type,proto3,enum=nodes.Type" json:"Type,omitempty"`
	Action Action `protobuf:"varint,2,opt,name=Action,json=action,proto3,enum=nodes.Action" json:"Action,omitempty"`
	Delta  int32  `protobuf:"varint,3,opt,name=Delta,json=number,proto3" json:"Delta,omitempty"`
}

func (x *ManagerAction) Reset() {
	*x = ManagerAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_proto_masternode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManagerAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagerAction) ProtoMessage() {}

func (x *ManagerAction) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_proto_masternode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagerAction.ProtoReflect.Descriptor instead.
func (*ManagerAction) Descriptor() ([]byte, []int) {
	return file_protobuf_proto_masternode_proto_rawDescGZIP(), []int{0}
}

func (x *ManagerAction) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_TYPE_UNKNOWN
}

func (x *ManagerAction) GetAction() Action {
	if x != nil {
		return x.Action
	}
	return Action_NOTHING
}

func (x *ManagerAction) GetDelta() int32 {
	if x != nil {
		return x.Delta
	}
	return 0
}

//*
// To be or not to be
type SimpleResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OK bool `protobuf:"varint,1,opt,name=OK,json=ok,proto3" json:"OK,omitempty"`
}

func (x *SimpleResult) Reset() {
	*x = SimpleResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_proto_masternode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleResult) ProtoMessage() {}

func (x *SimpleResult) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_proto_masternode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleResult.ProtoReflect.Descriptor instead.
func (*SimpleResult) Descriptor() ([]byte, []int) {
	return file_protobuf_proto_masternode_proto_rawDescGZIP(), []int{1}
}

func (x *SimpleResult) GetOK() bool {
	if x != nil {
		return x.OK
	}
	return false
}

//*
// message ChanData contents the information about single channel
type ChanData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsExisted  bool     `protobuf:"varint,1,opt,name=IsExisted,proto3" json:"IsExisted,omitempty"`           // last manager has false for out channel
	Type       ChanType `protobuf:"varint,2,opt,name=Type,proto3,enum=nodes.ChanType" json:"Type,omitempty"` //  type of channel. last manager has UNKNOWN for out channel
	Length     uint32   `protobuf:"varint,3,opt,name=Length,proto3" json:"Length,omitempty"`                 // length of channel
	NumberInCh uint32   `protobuf:"varint,4,opt,name=NumberInCh,proto3" json:"NumberInCh,omitempty"`         // number of items in channel right now
	// number of workers which are serving the channel right now. Usually it equals the WorkersData.Number
	NumberOfWorkers uint32 `protobuf:"varint,5,opt,name=NumberOfWorkers,proto3" json:"NumberOfWorkers,omitempty"`
}

func (x *ChanData) Reset() {
	*x = ChanData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_proto_masternode_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChanData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChanData) ProtoMessage() {}

func (x *ChanData) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_proto_masternode_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChanData.ProtoReflect.Descriptor instead.
func (*ChanData) Descriptor() ([]byte, []int) {
	return file_protobuf_proto_masternode_proto_rawDescGZIP(), []int{2}
}

func (x *ChanData) GetIsExisted() bool {
	if x != nil {
		return x.IsExisted
	}
	return false
}

func (x *ChanData) GetType() ChanType {
	if x != nil {
		return x.Type
	}
	return ChanType_CHAN_UNKNOWN
}

func (x *ChanData) GetLength() uint32 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *ChanData) GetNumberInCh() uint32 {
	if x != nil {
		return x.NumberInCh
	}
	return 0
}

func (x *ChanData) GetNumberOfWorkers() uint32 {
	if x != nil {
		return x.NumberOfWorkers
	}
	return 0
}

//*
// message WorkersData contents the information about workers for single manager
type WorkersData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Min    uint32 `protobuf:"varint,1,opt,name=Min,proto3" json:"Min,omitempty"`       // available minimum of number of workers
	Max    uint32 `protobuf:"varint,2,opt,name=Max,proto3" json:"Max,omitempty"`       // available maximum of number of workers
	Number uint32 `protobuf:"varint,3,opt,name=Number,proto3" json:"Number,omitempty"` // total number of workers right now
	Active uint32 `protobuf:"varint,4,opt,name=Active,proto3" json:"Active,omitempty"` // number of workers these are active right now
}

func (x *WorkersData) Reset() {
	*x = WorkersData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_proto_masternode_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkersData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkersData) ProtoMessage() {}

func (x *WorkersData) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_proto_masternode_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkersData.ProtoReflect.Descriptor instead.
func (*WorkersData) Descriptor() ([]byte, []int) {
	return file_protobuf_proto_masternode_proto_rawDescGZIP(), []int{3}
}

func (x *WorkersData) GetMin() uint32 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *WorkersData) GetMax() uint32 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *WorkersData) GetNumber() uint32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *WorkersData) GetActive() uint32 {
	if x != nil {
		return x.Active
	}
	return 0
}

//*
// message ManagerData contents the information about single manager
type ManagerData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string               `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	ID         string               `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
	Created    *timestamp.Timestamp `protobuf:"bytes,3,opt,name=Created,proto3" json:"Created,omitempty"` // time of getting into
	Delta      uint64               `protobuf:"varint,4,opt,name=Delta,proto3" json:"Delta,omitempty"`    // number of microseconds from last getting into
	Workers    *WorkersData         `protobuf:"bytes,5,opt,name=Workers,proto3" json:"Workers,omitempty"`
	ChanBefore []*ChanData          `protobuf:"bytes,6,rep,name=ChanBefore,proto3" json:"ChanBefore,omitempty"`
	ChanAfter  []*ChanData          `protobuf:"bytes,7,rep,name=ChanAfter,proto3" json:"ChanAfter,omitempty"`
}

func (x *ManagerData) Reset() {
	*x = ManagerData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_proto_masternode_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManagerData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagerData) ProtoMessage() {}

func (x *ManagerData) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_proto_masternode_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagerData.ProtoReflect.Descriptor instead.
func (*ManagerData) Descriptor() ([]byte, []int) {
	return file_protobuf_proto_masternode_proto_rawDescGZIP(), []int{4}
}

func (x *ManagerData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ManagerData) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *ManagerData) GetCreated() *timestamp.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *ManagerData) GetDelta() uint64 {
	if x != nil {
		return x.Delta
	}
	return 0
}

func (x *ManagerData) GetWorkers() *WorkersData {
	if x != nil {
		return x.Workers
	}
	return nil
}

func (x *ManagerData) GetChanBefore() []*ChanData {
	if x != nil {
		return x.ChanBefore
	}
	return nil
}

func (x *ManagerData) GetChanAfter() []*ChanData {
	if x != nil {
		return x.ChanAfter
	}
	return nil
}

//*
// SlaveNodeInfoRequest is a request with slave node data to master node
type SlaveNodeInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterID        string         `protobuf:"bytes,1,opt,name=ClusterID,proto3" json:"ClusterID,omitempty"`               // unique name of cluster. No specific info
	NodeID           string         `protobuf:"bytes,2,opt,name=NodeID,proto3" json:"NodeID,omitempty"`                     // unique name of node. No specific info
	ManagerData      []*ManagerData `protobuf:"bytes,3,rep,name=ManagerData,proto3" json:"ManagerData,omitempty"`           //
	FinalManagerData []*ManagerData `protobuf:"bytes,4,rep,name=FinalManagerData,proto3" json:"FinalManagerData,omitempty"` //
	ErrorManagerData []*ManagerData `protobuf:"bytes,5,rep,name=ErrorManagerData,proto3" json:"ErrorManagerData,omitempty"` //
}

func (x *SlaveNodeInfoRequest) Reset() {
	*x = SlaveNodeInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_proto_masternode_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlaveNodeInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlaveNodeInfoRequest) ProtoMessage() {}

func (x *SlaveNodeInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_proto_masternode_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlaveNodeInfoRequest.ProtoReflect.Descriptor instead.
func (*SlaveNodeInfoRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_proto_masternode_proto_rawDescGZIP(), []int{5}
}

func (x *SlaveNodeInfoRequest) GetClusterID() string {
	if x != nil {
		return x.ClusterID
	}
	return ""
}

func (x *SlaveNodeInfoRequest) GetNodeID() string {
	if x != nil {
		return x.NodeID
	}
	return ""
}

func (x *SlaveNodeInfoRequest) GetManagerData() []*ManagerData {
	if x != nil {
		return x.ManagerData
	}
	return nil
}

func (x *SlaveNodeInfoRequest) GetFinalManagerData() []*ManagerData {
	if x != nil {
		return x.FinalManagerData
	}
	return nil
}

func (x *SlaveNodeInfoRequest) GetErrorManagerData() []*ManagerData {
	if x != nil {
		return x.ErrorManagerData
	}
	return nil
}

var File_protobuf_proto_masternode_proto protoreflect.FileDescriptor

var file_protobuf_proto_masternode_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x0d, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x73,
	0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x6e, 0x6f,
	0x64, 0x65, 0x73, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x15, 0x0a, 0x05, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x1e, 0x0a, 0x0c, 0x53, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x4b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0xaf, 0x01, 0x0a, 0x08, 0x43, 0x68,
	0x61, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x73, 0x45, 0x78, 0x69, 0x73,
	0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x49, 0x73, 0x45, 0x78, 0x69,
	0x73, 0x74, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x4c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x4c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x43, 0x68, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x43,
	0x68, 0x12, 0x28, 0x0a, 0x0f, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x57, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x4f, 0x66, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x22, 0x61, 0x0a, 0x0b, 0x57,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x4d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03,
	0x4d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x4d, 0x61, 0x78, 0x12, 0x16,
	0x0a, 0x06, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x8b,
	0x02, 0x0a, 0x0b, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x34, 0x0a, 0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x44, 0x65, 0x6c, 0x74,
	0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x12, 0x2c,
	0x0a, 0x07, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x07, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x12, 0x2f, 0x0a, 0x0a,
	0x43, 0x68, 0x61, 0x6e, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x0a, 0x43, 0x68, 0x61, 0x6e, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x2d, 0x0a,
	0x09, 0x43, 0x68, 0x61, 0x6e, 0x41, 0x66, 0x74, 0x65, 0x72, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x09, 0x43, 0x68, 0x61, 0x6e, 0x41, 0x66, 0x74, 0x65, 0x72, 0x22, 0x82, 0x02, 0x0a,
	0x14, 0x53, 0x6c, 0x61, 0x76, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x12, 0x34, 0x0a, 0x0b, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x0b, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x3e, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6e, 0x6f,
	0x64, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x10, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x3e, 0x0a, 0x10, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6e, 0x6f,
	0x64, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x10, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x2a, 0x27, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x4e,
	0x4f, 0x54, 0x48, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x55, 0x50, 0x10, 0x01,
	0x12, 0x08, 0x0a, 0x04, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x02, 0x2a, 0x6e, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x57, 0x4f, 0x52,
	0x4b, 0x45, 0x52, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x48,
	0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x52, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x4e, 0x4f, 0x44, 0x45, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x10, 0x05, 0x2a, 0x56, 0x0a, 0x08, 0x43, 0x68,
	0x61, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x48, 0x41, 0x4e, 0x5f, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x48, 0x41, 0x4e,
	0x5f, 0x53, 0x54, 0x44, 0x5f, 0x47, 0x4f, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x48, 0x41,
	0x4e, 0x5f, 0x53, 0x54, 0x41, 0x43, 0x4b, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x48, 0x41,
	0x4e, 0x5f, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x51, 0x55, 0x45, 0x55, 0x45,
	0x10, 0x03, 0x32, 0x50, 0x0a, 0x0a, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65,
	0x12, 0x42, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x1b, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x53, 0x6c, 0x61, 0x76, 0x65,
	0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x42, 0x0a, 0x5a, 0x08, 0x67, 0x6f, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_proto_masternode_proto_rawDescOnce sync.Once
	file_protobuf_proto_masternode_proto_rawDescData = file_protobuf_proto_masternode_proto_rawDesc
)

func file_protobuf_proto_masternode_proto_rawDescGZIP() []byte {
	file_protobuf_proto_masternode_proto_rawDescOnce.Do(func() {
		file_protobuf_proto_masternode_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_proto_masternode_proto_rawDescData)
	})
	return file_protobuf_proto_masternode_proto_rawDescData
}

var file_protobuf_proto_masternode_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_protobuf_proto_masternode_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protobuf_proto_masternode_proto_goTypes = []interface{}{
	(Action)(0),                  // 0: nodes.Action
	(Type)(0),                    // 1: nodes.Type
	(ChanType)(0),                // 2: nodes.ChanType
	(*ManagerAction)(nil),        // 3: nodes.ManagerAction
	(*SimpleResult)(nil),         // 4: nodes.SimpleResult
	(*ChanData)(nil),             // 5: nodes.ChanData
	(*WorkersData)(nil),          // 6: nodes.WorkersData
	(*ManagerData)(nil),          // 7: nodes.ManagerData
	(*SlaveNodeInfoRequest)(nil), // 8: nodes.SlaveNodeInfoRequest
	(*timestamp.Timestamp)(nil),  // 9: google.protobuf.Timestamp
}
var file_protobuf_proto_masternode_proto_depIdxs = []int32{
	1,  // 0: nodes.ManagerAction.Type:type_name -> nodes.Type
	0,  // 1: nodes.ManagerAction.Action:type_name -> nodes.Action
	2,  // 2: nodes.ChanData.Type:type_name -> nodes.ChanType
	9,  // 3: nodes.ManagerData.Created:type_name -> google.protobuf.Timestamp
	6,  // 4: nodes.ManagerData.Workers:type_name -> nodes.WorkersData
	5,  // 5: nodes.ManagerData.ChanBefore:type_name -> nodes.ChanData
	5,  // 6: nodes.ManagerData.ChanAfter:type_name -> nodes.ChanData
	7,  // 7: nodes.SlaveNodeInfoRequest.ManagerData:type_name -> nodes.ManagerData
	7,  // 8: nodes.SlaveNodeInfoRequest.FinalManagerData:type_name -> nodes.ManagerData
	7,  // 9: nodes.SlaveNodeInfoRequest.ErrorManagerData:type_name -> nodes.ManagerData
	8,  // 10: nodes.MasterNode.UpdateNodeInfo:input_type -> nodes.SlaveNodeInfoRequest
	4,  // 11: nodes.MasterNode.UpdateNodeInfo:output_type -> nodes.SimpleResult
	11, // [11:12] is the sub-list for method output_type
	10, // [10:11] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_protobuf_proto_masternode_proto_init() }
func file_protobuf_proto_masternode_proto_init() {
	if File_protobuf_proto_masternode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_proto_masternode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManagerAction); i {
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
		file_protobuf_proto_masternode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleResult); i {
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
		file_protobuf_proto_masternode_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChanData); i {
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
		file_protobuf_proto_masternode_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkersData); i {
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
		file_protobuf_proto_masternode_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManagerData); i {
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
		file_protobuf_proto_masternode_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlaveNodeInfoRequest); i {
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
			RawDescriptor: file_protobuf_proto_masternode_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_proto_masternode_proto_goTypes,
		DependencyIndexes: file_protobuf_proto_masternode_proto_depIdxs,
		EnumInfos:         file_protobuf_proto_masternode_proto_enumTypes,
		MessageInfos:      file_protobuf_proto_masternode_proto_msgTypes,
	}.Build()
	File_protobuf_proto_masternode_proto = out.File
	file_protobuf_proto_masternode_proto_rawDesc = nil
	file_protobuf_proto_masternode_proto_goTypes = nil
	file_protobuf_proto_masternode_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MasterNodeClient is the client API for MasterNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MasterNodeClient interface {
	// rpc UpdateNodeInfo sends info about current details of slave node to master node
	UpdateNodeInfo(ctx context.Context, in *SlaveNodeInfoRequest, opts ...grpc.CallOption) (*SimpleResult, error)
}

type masterNodeClient struct {
	cc grpc.ClientConnInterface
}

func NewMasterNodeClient(cc grpc.ClientConnInterface) MasterNodeClient {
	return &masterNodeClient{cc}
}

func (c *masterNodeClient) UpdateNodeInfo(ctx context.Context, in *SlaveNodeInfoRequest, opts ...grpc.CallOption) (*SimpleResult, error) {
	out := new(SimpleResult)
	err := c.cc.Invoke(ctx, "/nodes.MasterNode/UpdateNodeInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterNodeServer is the server API for MasterNode service.
type MasterNodeServer interface {
	// rpc UpdateNodeInfo sends info about current details of slave node to master node
	UpdateNodeInfo(context.Context, *SlaveNodeInfoRequest) (*SimpleResult, error)
}

// UnimplementedMasterNodeServer can be embedded to have forward compatible implementations.
type UnimplementedMasterNodeServer struct {
}

func (*UnimplementedMasterNodeServer) UpdateNodeInfo(context.Context, *SlaveNodeInfoRequest) (*SimpleResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNodeInfo not implemented")
}

func RegisterMasterNodeServer(s *grpc.Server, srv MasterNodeServer) {
	s.RegisterService(&_MasterNode_serviceDesc, srv)
}

func _MasterNode_UpdateNodeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SlaveNodeInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterNodeServer).UpdateNodeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodes.MasterNode/UpdateNodeInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterNodeServer).UpdateNodeInfo(ctx, req.(*SlaveNodeInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MasterNode_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nodes.MasterNode",
	HandlerType: (*MasterNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateNodeInfo",
			Handler:    _MasterNode_UpdateNodeInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/proto/masternode.proto",
}