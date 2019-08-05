// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages/messages.proto

package messages

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RequestStatus int32

const (
	RequestStatus_REQUEST_SUCCESSFUL RequestStatus = 0
	RequestStatus_REQUEST_FAILED     RequestStatus = 1
)

var RequestStatus_name = map[int32]string{
	0: "REQUEST_SUCCESSFUL",
	1: "REQUEST_FAILED",
}

var RequestStatus_value = map[string]int32{
	"REQUEST_SUCCESSFUL": 0,
	"REQUEST_FAILED":     1,
}

func (x RequestStatus) String() string {
	return proto.EnumName(RequestStatus_name, int32(x))
}

func (RequestStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_83994550f81e9f35, []int{0}
}

type LaunchLogStatus int32

const (
	LaunchLogStatus_CREATED     LaunchLogStatus = 0
	LaunchLogStatus_RETRIED     LaunchLogStatus = 1
	LaunchLogStatus_PENDING     LaunchLogStatus = 2
	LaunchLogStatus_SUCCESS     LaunchLogStatus = 3
	LaunchLogStatus_FAILED      LaunchLogStatus = 4
	LaunchLogStatus_SIGN_FAILED LaunchLogStatus = 5
)

var LaunchLogStatus_name = map[int32]string{
	0: "CREATED",
	1: "RETRIED",
	2: "PENDING",
	3: "SUCCESS",
	4: "FAILED",
	5: "SIGN_FAILED",
}

var LaunchLogStatus_value = map[string]int32{
	"CREATED":     0,
	"RETRIED":     1,
	"PENDING":     2,
	"SUCCESS":     3,
	"FAILED":      4,
	"SIGN_FAILED": 5,
}

func (x LaunchLogStatus) String() string {
	return proto.EnumName(LaunchLogStatus_name, int32(x))
}

func (LaunchLogStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_83994550f81e9f35, []int{1}
}

// The request message containing the user's name.
type CreateMessage struct {
	From                 string   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   string   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Data                 []byte   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	GasPrice             string   `protobuf:"bytes,5,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
	GasLimit             uint32   `protobuf:"varint,6,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	ItemType             string   `protobuf:"bytes,7,opt,name=item_type,json=itemType,proto3" json:"item_type,omitempty"`
	ItemId               string   `protobuf:"bytes,8,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMessage) Reset()         { *m = CreateMessage{} }
func (m *CreateMessage) String() string { return proto.CompactTextString(m) }
func (*CreateMessage) ProtoMessage()    {}
func (*CreateMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_83994550f81e9f35, []int{0}
}

func (m *CreateMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMessage.Unmarshal(m, b)
}
func (m *CreateMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMessage.Marshal(b, m, deterministic)
}
func (m *CreateMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMessage.Merge(m, src)
}
func (m *CreateMessage) XXX_Size() int {
	return xxx_messageInfo_CreateMessage.Size(m)
}
func (m *CreateMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMessage.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMessage proto.InternalMessageInfo

func (m *CreateMessage) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *CreateMessage) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *CreateMessage) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *CreateMessage) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *CreateMessage) GetGasPrice() string {
	if m != nil {
		return m.GasPrice
	}
	return ""
}

func (m *CreateMessage) GetGasLimit() uint32 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

func (m *CreateMessage) GetItemType() string {
	if m != nil {
		return m.ItemType
	}
	return ""
}

func (m *CreateMessage) GetItemId() string {
	if m != nil {
		return m.ItemId
	}
	return ""
}

// The response message containing the greetings
type CreateReply struct {
	Status               RequestStatus `protobuf:"varint,1,opt,name=status,proto3,enum=messages.RequestStatus" json:"status,omitempty"`
	ErrMsg               string        `protobuf:"bytes,2,opt,name=err_msg,json=errMsg,proto3" json:"err_msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateReply) Reset()         { *m = CreateReply{} }
func (m *CreateReply) String() string { return proto.CompactTextString(m) }
func (*CreateReply) ProtoMessage()    {}
func (*CreateReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_83994550f81e9f35, []int{1}
}

func (m *CreateReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateReply.Unmarshal(m, b)
}
func (m *CreateReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateReply.Marshal(b, m, deterministic)
}
func (m *CreateReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateReply.Merge(m, src)
}
func (m *CreateReply) XXX_Size() int {
	return xxx_messageInfo_CreateReply.Size(m)
}
func (m *CreateReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateReply.DiscardUnknown(m)
}

var xxx_messageInfo_CreateReply proto.InternalMessageInfo

func (m *CreateReply) GetStatus() RequestStatus {
	if m != nil {
		return m.Status
	}
	return RequestStatus_REQUEST_SUCCESSFUL
}

func (m *CreateReply) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

type HelloMessage struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloMessage) Reset()         { *m = HelloMessage{} }
func (m *HelloMessage) String() string { return proto.CompactTextString(m) }
func (*HelloMessage) ProtoMessage()    {}
func (*HelloMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_83994550f81e9f35, []int{2}
}

func (m *HelloMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloMessage.Unmarshal(m, b)
}
func (m *HelloMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloMessage.Marshal(b, m, deterministic)
}
func (m *HelloMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloMessage.Merge(m, src)
}
func (m *HelloMessage) XXX_Size() int {
	return xxx_messageInfo_HelloMessage.Size(m)
}
func (m *HelloMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloMessage.DiscardUnknown(m)
}

var xxx_messageInfo_HelloMessage proto.InternalMessageInfo

type HelloReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_83994550f81e9f35, []int{3}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

type NotifyMessage struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Status               bool     `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotifyMessage) Reset()         { *m = NotifyMessage{} }
func (m *NotifyMessage) String() string { return proto.CompactTextString(m) }
func (*NotifyMessage) ProtoMessage()    {}
func (*NotifyMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_83994550f81e9f35, []int{4}
}

func (m *NotifyMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyMessage.Unmarshal(m, b)
}
func (m *NotifyMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyMessage.Marshal(b, m, deterministic)
}
func (m *NotifyMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyMessage.Merge(m, src)
}
func (m *NotifyMessage) XXX_Size() int {
	return xxx_messageInfo_NotifyMessage.Size(m)
}
func (m *NotifyMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyMessage.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyMessage proto.InternalMessageInfo

func (m *NotifyMessage) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *NotifyMessage) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

type NotifyReply struct {
	Status               RequestStatus `protobuf:"varint,1,opt,name=status,proto3,enum=messages.RequestStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *NotifyReply) Reset()         { *m = NotifyReply{} }
func (m *NotifyReply) String() string { return proto.CompactTextString(m) }
func (*NotifyReply) ProtoMessage()    {}
func (*NotifyReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_83994550f81e9f35, []int{5}
}

func (m *NotifyReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyReply.Unmarshal(m, b)
}
func (m *NotifyReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyReply.Marshal(b, m, deterministic)
}
func (m *NotifyReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyReply.Merge(m, src)
}
func (m *NotifyReply) XXX_Size() int {
	return xxx_messageInfo_NotifyReply.Size(m)
}
func (m *NotifyReply) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyReply.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyReply proto.InternalMessageInfo

func (m *NotifyReply) GetStatus() RequestStatus {
	if m != nil {
		return m.Status
	}
	return RequestStatus_REQUEST_SUCCESSFUL
}

type Log struct {
	Hash                 string          `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	ItemType             string          `protobuf:"bytes,2,opt,name=item_type,json=itemType,proto3" json:"item_type,omitempty"`
	ItemId               string          `protobuf:"bytes,3,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Status               LaunchLogStatus `protobuf:"varint,4,opt,name=status,proto3,enum=messages.LaunchLogStatus" json:"status,omitempty"`
	GasPrice             string          `protobuf:"bytes,5,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
	GasLimit             string          `protobuf:"bytes,6,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Log) Reset()         { *m = Log{} }
func (m *Log) String() string { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()    {}
func (*Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_83994550f81e9f35, []int{6}
}

func (m *Log) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Log.Unmarshal(m, b)
}
func (m *Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Log.Marshal(b, m, deterministic)
}
func (m *Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Log.Merge(m, src)
}
func (m *Log) XXX_Size() int {
	return xxx_messageInfo_Log.Size(m)
}
func (m *Log) XXX_DiscardUnknown() {
	xxx_messageInfo_Log.DiscardUnknown(m)
}

var xxx_messageInfo_Log proto.InternalMessageInfo

func (m *Log) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Log) GetItemType() string {
	if m != nil {
		return m.ItemType
	}
	return ""
}

func (m *Log) GetItemId() string {
	if m != nil {
		return m.ItemId
	}
	return ""
}

func (m *Log) GetStatus() LaunchLogStatus {
	if m != nil {
		return m.Status
	}
	return LaunchLogStatus_CREATED
}

func (m *Log) GetGasPrice() string {
	if m != nil {
		return m.GasPrice
	}
	return ""
}

func (m *Log) GetGasLimit() string {
	if m != nil {
		return m.GasLimit
	}
	return ""
}

type GetMessage struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	ItemType             string   `protobuf:"bytes,2,opt,name=item_type,json=itemType,proto3" json:"item_type,omitempty"`
	ItemId               string   `protobuf:"bytes,3,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMessage) Reset()         { *m = GetMessage{} }
func (m *GetMessage) String() string { return proto.CompactTextString(m) }
func (*GetMessage) ProtoMessage()    {}
func (*GetMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_83994550f81e9f35, []int{7}
}

func (m *GetMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMessage.Unmarshal(m, b)
}
func (m *GetMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMessage.Marshal(b, m, deterministic)
}
func (m *GetMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMessage.Merge(m, src)
}
func (m *GetMessage) XXX_Size() int {
	return xxx_messageInfo_GetMessage.Size(m)
}
func (m *GetMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMessage.DiscardUnknown(m)
}

var xxx_messageInfo_GetMessage proto.InternalMessageInfo

func (m *GetMessage) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *GetMessage) GetItemType() string {
	if m != nil {
		return m.ItemType
	}
	return ""
}

func (m *GetMessage) GetItemId() string {
	if m != nil {
		return m.ItemId
	}
	return ""
}

type GetReply struct {
	Status               RequestStatus `protobuf:"varint,1,opt,name=status,proto3,enum=messages.RequestStatus" json:"status,omitempty"`
	Data                 []*Log        `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetReply) Reset()         { *m = GetReply{} }
func (m *GetReply) String() string { return proto.CompactTextString(m) }
func (*GetReply) ProtoMessage()    {}
func (*GetReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_83994550f81e9f35, []int{8}
}

func (m *GetReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetReply.Unmarshal(m, b)
}
func (m *GetReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetReply.Marshal(b, m, deterministic)
}
func (m *GetReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetReply.Merge(m, src)
}
func (m *GetReply) XXX_Size() int {
	return xxx_messageInfo_GetReply.Size(m)
}
func (m *GetReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetReply proto.InternalMessageInfo

func (m *GetReply) GetStatus() RequestStatus {
	if m != nil {
		return m.Status
	}
	return RequestStatus_REQUEST_SUCCESSFUL
}

func (m *GetReply) GetData() []*Log {
	if m != nil {
		return m.Data
	}
	return nil
}

type SubscribeMessage struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	ItemType             string   `protobuf:"bytes,2,opt,name=item_type,json=itemType,proto3" json:"item_type,omitempty"`
	ItemId               string   `protobuf:"bytes,3,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeMessage) Reset()         { *m = SubscribeMessage{} }
func (m *SubscribeMessage) String() string { return proto.CompactTextString(m) }
func (*SubscribeMessage) ProtoMessage()    {}
func (*SubscribeMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_83994550f81e9f35, []int{9}
}

func (m *SubscribeMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeMessage.Unmarshal(m, b)
}
func (m *SubscribeMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeMessage.Marshal(b, m, deterministic)
}
func (m *SubscribeMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeMessage.Merge(m, src)
}
func (m *SubscribeMessage) XXX_Size() int {
	return xxx_messageInfo_SubscribeMessage.Size(m)
}
func (m *SubscribeMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeMessage proto.InternalMessageInfo

func (m *SubscribeMessage) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *SubscribeMessage) GetItemType() string {
	if m != nil {
		return m.ItemType
	}
	return ""
}

func (m *SubscribeMessage) GetItemId() string {
	if m != nil {
		return m.ItemId
	}
	return ""
}

type SubscribeReply struct {
	Status               LaunchLogStatus `protobuf:"varint,1,opt,name=status,proto3,enum=messages.LaunchLogStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SubscribeReply) Reset()         { *m = SubscribeReply{} }
func (m *SubscribeReply) String() string { return proto.CompactTextString(m) }
func (*SubscribeReply) ProtoMessage()    {}
func (*SubscribeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_83994550f81e9f35, []int{10}
}

func (m *SubscribeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeReply.Unmarshal(m, b)
}
func (m *SubscribeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeReply.Marshal(b, m, deterministic)
}
func (m *SubscribeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeReply.Merge(m, src)
}
func (m *SubscribeReply) XXX_Size() int {
	return xxx_messageInfo_SubscribeReply.Size(m)
}
func (m *SubscribeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeReply.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeReply proto.InternalMessageInfo

func (m *SubscribeReply) GetStatus() LaunchLogStatus {
	if m != nil {
		return m.Status
	}
	return LaunchLogStatus_CREATED
}

func init() {
	proto.RegisterEnum("messages.RequestStatus", RequestStatus_name, RequestStatus_value)
	proto.RegisterEnum("messages.LaunchLogStatus", LaunchLogStatus_name, LaunchLogStatus_value)
	proto.RegisterType((*CreateMessage)(nil), "messages.CreateMessage")
	proto.RegisterType((*CreateReply)(nil), "messages.CreateReply")
	proto.RegisterType((*HelloMessage)(nil), "messages.HelloMessage")
	proto.RegisterType((*HelloReply)(nil), "messages.HelloReply")
	proto.RegisterType((*NotifyMessage)(nil), "messages.NotifyMessage")
	proto.RegisterType((*NotifyReply)(nil), "messages.NotifyReply")
	proto.RegisterType((*Log)(nil), "messages.Log")
	proto.RegisterType((*GetMessage)(nil), "messages.GetMessage")
	proto.RegisterType((*GetReply)(nil), "messages.GetReply")
	proto.RegisterType((*SubscribeMessage)(nil), "messages.SubscribeMessage")
	proto.RegisterType((*SubscribeReply)(nil), "messages.SubscribeReply")
}

func init() { proto.RegisterFile("messages/messages.proto", fileDescriptor_83994550f81e9f35) }

var fileDescriptor_83994550f81e9f35 = []byte{
	// 606 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0xda, 0x40,
	0x10, 0xc6, 0x36, 0x38, 0x66, 0x08, 0xc4, 0x1a, 0xa5, 0x89, 0x4b, 0x2f, 0xd4, 0x27, 0x94, 0x43,
	0xd2, 0xa4, 0x87, 0x4a, 0x89, 0x54, 0x29, 0x22, 0x8e, 0x8b, 0xe4, 0xa0, 0xd4, 0x26, 0xed, 0xa5,
	0x2a, 0x32, 0xb0, 0x31, 0x96, 0xa0, 0xa6, 0xde, 0xa5, 0x12, 0x4f, 0xd6, 0x4b, 0x5f, 0xa1, 0xef,
	0x54, 0xed, 0xda, 0xc6, 0xc6, 0x4a, 0x5b, 0x15, 0xf5, 0xb6, 0x33, 0xdf, 0xfc, 0x7e, 0x33, 0xb3,
	0x70, 0xbc, 0x20, 0x94, 0xfa, 0x01, 0xa1, 0x67, 0xd9, 0xe3, 0x74, 0x19, 0x47, 0x2c, 0x42, 0x2d,
	0x93, 0xcd, 0x9f, 0x12, 0x34, 0x7b, 0x31, 0xf1, 0x19, 0xb9, 0x4b, 0x54, 0x88, 0x50, 0x7d, 0x8c,
	0xa3, 0x85, 0x21, 0x75, 0xa4, 0x6e, 0xdd, 0x15, 0x6f, 0x6c, 0x81, 0xcc, 0x22, 0x43, 0x16, 0x1a,
	0x99, 0x45, 0x78, 0x08, 0xb5, 0x6f, 0xfe, 0x7c, 0x45, 0x0c, 0x45, 0xa8, 0x12, 0x81, 0x7b, 0x4e,
	0x7d, 0xe6, 0x1b, 0xd5, 0x8e, 0xd4, 0xdd, 0x77, 0xc5, 0x1b, 0x5f, 0x40, 0x3d, 0xf0, 0xe9, 0x68,
	0x19, 0x87, 0x13, 0x62, 0xd4, 0x84, 0xb5, 0x16, 0xf8, 0xf4, 0x9e, 0xcb, 0x19, 0x38, 0x0f, 0x17,
	0x21, 0x33, 0xd4, 0x8e, 0xd4, 0x6d, 0x0a, 0xd0, 0xe1, 0x32, 0x07, 0x43, 0x46, 0x16, 0x23, 0xb6,
	0x5e, 0x12, 0x63, 0x2f, 0xf1, 0xe4, 0x8a, 0xe1, 0x7a, 0x49, 0xf0, 0x18, 0xf6, 0x04, 0x18, 0x4e,
	0x0d, 0x4d, 0x40, 0x2a, 0x17, 0xfb, 0x53, 0xf3, 0x23, 0x34, 0x92, 0x76, 0x5c, 0xb2, 0x9c, 0xaf,
	0xf1, 0x0c, 0x54, 0xca, 0x7c, 0xb6, 0xa2, 0xa2, 0x9d, 0xd6, 0xc5, 0xf1, 0xe9, 0x86, 0x09, 0x97,
	0x7c, 0x5d, 0x11, 0xca, 0x3c, 0x01, 0xbb, 0xa9, 0x19, 0x0f, 0x4c, 0xe2, 0x78, 0xb4, 0xa0, 0x41,
	0xda, 0xae, 0x4a, 0xe2, 0xf8, 0x8e, 0x06, 0x66, 0x0b, 0xf6, 0xdf, 0x91, 0xf9, 0x3c, 0x4a, 0x69,
	0x32, 0xf7, 0x01, 0x84, 0x2c, 0xf2, 0x98, 0x57, 0xd0, 0x1c, 0x44, 0x2c, 0x7c, 0x5c, 0x17, 0x58,
	0x9c, 0xf9, 0x74, 0x96, 0xb1, 0xc8, 0xdf, 0x78, 0xb4, 0x29, 0x86, 0x87, 0xd6, 0xb2, 0x9c, 0xe6,
	0x5b, 0x68, 0x24, 0xce, 0xbb, 0xd5, 0x6c, 0xfe, 0x90, 0x40, 0x71, 0xa2, 0xe0, 0xc9, 0x9c, 0x5b,
	0x2c, 0xca, 0xbf, 0x67, 0x51, 0x29, 0xb2, 0x88, 0xe7, 0x9b, 0x12, 0xaa, 0xa2, 0x84, 0xe7, 0x79,
	0x09, 0x8e, 0xbf, 0xfa, 0x32, 0x99, 0x39, 0x51, 0x50, 0x22, 0xee, 0xdf, 0x06, 0x5d, 0xcf, 0x07,
	0x6d, 0x7e, 0x00, 0xb0, 0x09, 0xfb, 0x13, 0x71, 0x3b, 0x35, 0x61, 0x7e, 0x06, 0xcd, 0x26, 0x6c,
	0xc7, 0x3d, 0x78, 0x99, 0xee, 0xb2, 0xdc, 0x51, 0xba, 0x8d, 0x8b, 0x66, 0xa1, 0xff, 0x28, 0x48,
	0x56, 0xdb, 0xfc, 0x04, 0xba, 0xb7, 0x1a, 0xd3, 0x49, 0x1c, 0x8e, 0xc9, 0xff, 0xaf, 0xbe, 0x07,
	0xad, 0x4d, 0xf4, 0xa4, 0x87, 0xf3, 0x52, 0x0f, 0x7f, 0x1f, 0xca, 0xc9, 0x15, 0x34, 0xb7, 0xda,
	0xc3, 0x23, 0x40, 0xd7, 0x7a, 0xff, 0x60, 0x79, 0xc3, 0x91, 0xf7, 0xd0, 0xeb, 0x59, 0x9e, 0x77,
	0xfb, 0xe0, 0xe8, 0x15, 0x44, 0x68, 0x65, 0xfa, 0xdb, 0xeb, 0xbe, 0x63, 0xdd, 0xe8, 0xd2, 0xc9,
	0x18, 0x0e, 0x4a, 0x71, 0xb1, 0x01, 0x7b, 0x3d, 0xd7, 0xba, 0x1e, 0x5a, 0x37, 0x7a, 0x85, 0x0b,
	0xae, 0x35, 0x74, 0xfb, 0xdc, 0x98, 0x0b, 0xf7, 0xd6, 0xe0, 0xa6, 0x3f, 0xb0, 0x75, 0x99, 0x0b,
	0x69, 0x74, 0x5d, 0x41, 0x00, 0x35, 0x0d, 0x59, 0xc5, 0x03, 0x68, 0x78, 0x7d, 0x7b, 0x90, 0xe5,
	0xa8, 0x5d, 0x7c, 0x97, 0x41, 0x4b, 0x92, 0x90, 0x18, 0x2f, 0x41, 0x4d, 0x6e, 0x17, 0x0b, 0xe3,
	0xd9, 0xfa, 0x9c, 0xda, 0xcf, 0xca, 0x40, 0x72, 0x7e, 0x15, 0x3c, 0x07, 0xc5, 0x26, 0x0c, 0x0f,
	0x73, 0x3c, 0xdf, 0xa9, 0x36, 0x6e, 0x69, 0x33, 0x97, 0x4b, 0x50, 0x93, 0xb3, 0x2b, 0xa6, 0xdb,
	0xba, 0xe2, 0x62, 0xba, 0xc2, 0x85, 0x9a, 0x15, 0x7c, 0x03, 0x35, 0x71, 0xfd, 0x78, 0x94, 0x5b,
	0x14, 0xbf, 0x87, 0xf6, 0x61, 0x49, 0x9f, 0x39, 0xda, 0x50, 0xdf, 0x8c, 0x15, 0xdb, 0xb9, 0x51,
	0x79, 0x93, 0xda, 0xc6, 0x13, 0x58, 0x1a, 0xa4, 0x2b, 0xbd, 0x92, 0xc6, 0xaa, 0xf8, 0xc9, 0x5f,
	0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xc4, 0xd2, 0x8f, 0xb3, 0xe4, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LauncherClient is the client API for Launcher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LauncherClient interface {
	Create(ctx context.Context, in *CreateMessage, opts ...grpc.CallOption) (*CreateReply, error)
	Get(ctx context.Context, in *GetMessage, opts ...grpc.CallOption) (*GetReply, error)
	Notify(ctx context.Context, in *NotifyMessage, opts ...grpc.CallOption) (*NotifyReply, error)
	Hello(ctx context.Context, in *HelloMessage, opts ...grpc.CallOption) (*HelloReply, error)
	// will push every changed launch log info to the subscriber
	Subscribe(ctx context.Context, opts ...grpc.CallOption) (Launcher_SubscribeClient, error)
}

type launcherClient struct {
	cc *grpc.ClientConn
}

func NewLauncherClient(cc *grpc.ClientConn) LauncherClient {
	return &launcherClient{cc}
}

func (c *launcherClient) Create(ctx context.Context, in *CreateMessage, opts ...grpc.CallOption) (*CreateReply, error) {
	out := new(CreateReply)
	err := c.cc.Invoke(ctx, "/messages.Launcher/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *launcherClient) Get(ctx context.Context, in *GetMessage, opts ...grpc.CallOption) (*GetReply, error) {
	out := new(GetReply)
	err := c.cc.Invoke(ctx, "/messages.Launcher/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *launcherClient) Notify(ctx context.Context, in *NotifyMessage, opts ...grpc.CallOption) (*NotifyReply, error) {
	out := new(NotifyReply)
	err := c.cc.Invoke(ctx, "/messages.Launcher/Notify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *launcherClient) Hello(ctx context.Context, in *HelloMessage, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/messages.Launcher/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *launcherClient) Subscribe(ctx context.Context, opts ...grpc.CallOption) (Launcher_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Launcher_serviceDesc.Streams[0], "/messages.Launcher/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &launcherSubscribeClient{stream}
	return x, nil
}

type Launcher_SubscribeClient interface {
	Send(*SubscribeMessage) error
	Recv() (*SubscribeReply, error)
	grpc.ClientStream
}

type launcherSubscribeClient struct {
	grpc.ClientStream
}

func (x *launcherSubscribeClient) Send(m *SubscribeMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *launcherSubscribeClient) Recv() (*SubscribeReply, error) {
	m := new(SubscribeReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LauncherServer is the server API for Launcher service.
type LauncherServer interface {
	Create(context.Context, *CreateMessage) (*CreateReply, error)
	Get(context.Context, *GetMessage) (*GetReply, error)
	Notify(context.Context, *NotifyMessage) (*NotifyReply, error)
	Hello(context.Context, *HelloMessage) (*HelloReply, error)
	// will push every changed launch log info to the subscriber
	Subscribe(Launcher_SubscribeServer) error
}

// UnimplementedLauncherServer can be embedded to have forward compatible implementations.
type UnimplementedLauncherServer struct {
}

func (*UnimplementedLauncherServer) Create(ctx context.Context, req *CreateMessage) (*CreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedLauncherServer) Get(ctx context.Context, req *GetMessage) (*GetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedLauncherServer) Notify(ctx context.Context, req *NotifyMessage) (*NotifyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Notify not implemented")
}
func (*UnimplementedLauncherServer) Hello(ctx context.Context, req *HelloMessage) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (*UnimplementedLauncherServer) Subscribe(srv Launcher_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}

func RegisterLauncherServer(s *grpc.Server, srv LauncherServer) {
	s.RegisterService(&_Launcher_serviceDesc, srv)
}

func _Launcher_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LauncherServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messages.Launcher/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LauncherServer).Create(ctx, req.(*CreateMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Launcher_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LauncherServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messages.Launcher/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LauncherServer).Get(ctx, req.(*GetMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Launcher_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LauncherServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messages.Launcher/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LauncherServer).Notify(ctx, req.(*NotifyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Launcher_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LauncherServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messages.Launcher/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LauncherServer).Hello(ctx, req.(*HelloMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Launcher_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LauncherServer).Subscribe(&launcherSubscribeServer{stream})
}

type Launcher_SubscribeServer interface {
	Send(*SubscribeReply) error
	Recv() (*SubscribeMessage, error)
	grpc.ServerStream
}

type launcherSubscribeServer struct {
	grpc.ServerStream
}

func (x *launcherSubscribeServer) Send(m *SubscribeReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *launcherSubscribeServer) Recv() (*SubscribeMessage, error) {
	m := new(SubscribeMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Launcher_serviceDesc = grpc.ServiceDesc{
	ServiceName: "messages.Launcher",
	HandlerType: (*LauncherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Launcher_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Launcher_Get_Handler,
		},
		{
			MethodName: "Notify",
			Handler:    _Launcher_Notify_Handler,
		},
		{
			MethodName: "Hello",
			Handler:    _Launcher_Hello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _Launcher_Subscribe_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "messages/messages.proto",
}
