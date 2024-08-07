// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.3
// source: socket.proto

package socket

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type Response_Error_Code int32

const (
	Response_Error_UNKNOWN             Response_Error_Code = 0
	Response_Error_INVALID_REQUEST     Response_Error_Code = 1
	Response_Error_UNSUPPORTED_REQUEST Response_Error_Code = 2
)

// Enum value maps for Response_Error_Code.
var (
	Response_Error_Code_name = map[int32]string{
		0: "UNKNOWN",
		1: "INVALID_REQUEST",
		2: "UNSUPPORTED_REQUEST",
	}
	Response_Error_Code_value = map[string]int32{
		"UNKNOWN":             0,
		"INVALID_REQUEST":     1,
		"UNSUPPORTED_REQUEST": 2,
	}
)

func (x Response_Error_Code) Enum() *Response_Error_Code {
	p := new(Response_Error_Code)
	*p = x
	return p
}

func (x Response_Error_Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Response_Error_Code) Descriptor() protoreflect.EnumDescriptor {
	return file_socket_proto_enumTypes[0].Descriptor()
}

func (Response_Error_Code) Type() protoreflect.EnumType {
	return &file_socket_proto_enumTypes[0]
}

func (x Response_Error_Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Response_Error_Code.Descriptor instead.
func (Response_Error_Code) EnumDescriptor() ([]byte, []int) {
	return file_socket_proto_rawDescGZIP(), []int{1, 1, 0}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*Request_PlayerMove
	Type isRequest_Type `protobuf_oneof:"type"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_socket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_socket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_socket_proto_rawDescGZIP(), []int{0}
}

func (m *Request) GetType() isRequest_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Request) GetPlayerMove() *Pos {
	if x, ok := x.GetType().(*Request_PlayerMove); ok {
		return x.PlayerMove
	}
	return nil
}

type isRequest_Type interface {
	isRequest_Type()
}

type Request_PlayerMove struct {
	// Player moves
	PlayerMove *Pos `protobuf:"bytes,1,opt,name=player_move,json=playerMove,proto3,oneof"`
}

func (*Request_PlayerMove) isRequest_Type() {}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*Response_Bulk_
	//	*Response_Error_
	//	*Response_Info_
	//	*Response_PlayerJoin_
	//	*Response_PlayerLeave
	//	*Response_PlayerMove_
	Type isResponse_Type `protobuf_oneof:"type"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_socket_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_socket_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_socket_proto_rawDescGZIP(), []int{1}
}

func (m *Response) GetType() isResponse_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Response) GetBulk() *Response_Bulk {
	if x, ok := x.GetType().(*Response_Bulk_); ok {
		return x.Bulk
	}
	return nil
}

func (x *Response) GetError() *Response_Error {
	if x, ok := x.GetType().(*Response_Error_); ok {
		return x.Error
	}
	return nil
}

func (x *Response) GetInfo() *Response_Info {
	if x, ok := x.GetType().(*Response_Info_); ok {
		return x.Info
	}
	return nil
}

func (x *Response) GetPlayerJoin() *Response_PlayerJoin {
	if x, ok := x.GetType().(*Response_PlayerJoin_); ok {
		return x.PlayerJoin
	}
	return nil
}

func (x *Response) GetPlayerLeave() uint64 {
	if x, ok := x.GetType().(*Response_PlayerLeave); ok {
		return x.PlayerLeave
	}
	return 0
}

func (x *Response) GetPlayerMove() *Response_PlayerMove {
	if x, ok := x.GetType().(*Response_PlayerMove_); ok {
		return x.PlayerMove
	}
	return nil
}

type isResponse_Type interface {
	isResponse_Type()
}

type Response_Bulk_ struct {
	// Bulk responses
	Bulk *Response_Bulk `protobuf:"bytes,1,opt,name=bulk,proto3,oneof"`
}

type Response_Error_ struct {
	// Error
	Error *Response_Error `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

type Response_Info_ struct {
	// Info
	Info *Response_Info `protobuf:"bytes,3,opt,name=info,proto3,oneof"`
}

type Response_PlayerJoin_ struct {
	// Player joins game
	PlayerJoin *Response_PlayerJoin `protobuf:"bytes,10,opt,name=player_join,json=playerJoin,proto3,oneof"`
}

type Response_PlayerLeave struct {
	// Player leaves game
	PlayerLeave uint64 `protobuf:"varint,11,opt,name=player_leave,json=playerLeave,proto3,oneof"`
}

type Response_PlayerMove_ struct {
	// Player moves
	PlayerMove *Response_PlayerMove `protobuf:"bytes,20,opt,name=player_move,json=playerMove,proto3,oneof"`
}

func (*Response_Bulk_) isResponse_Type() {}

func (*Response_Error_) isResponse_Type() {}

func (*Response_Info_) isResponse_Type() {}

func (*Response_PlayerJoin_) isResponse_Type() {}

func (*Response_PlayerLeave) isResponse_Type() {}

func (*Response_PlayerMove_) isResponse_Type() {}

type Pos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float64 `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float64 `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Pos) Reset() {
	*x = Pos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_socket_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pos) ProtoMessage() {}

func (x *Pos) ProtoReflect() protoreflect.Message {
	mi := &file_socket_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pos.ProtoReflect.Descriptor instead.
func (*Pos) Descriptor() ([]byte, []int) {
	return file_socket_proto_rawDescGZIP(), []int{2}
}

func (x *Pos) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Pos) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

type Response_Bulk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Responses []*Response `protobuf:"bytes,1,rep,name=responses,proto3" json:"responses,omitempty"`
}

func (x *Response_Bulk) Reset() {
	*x = Response_Bulk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_socket_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response_Bulk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response_Bulk) ProtoMessage() {}

func (x *Response_Bulk) ProtoReflect() protoreflect.Message {
	mi := &file_socket_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response_Bulk.ProtoReflect.Descriptor instead.
func (*Response_Bulk) Descriptor() ([]byte, []int) {
	return file_socket_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Response_Bulk) GetResponses() []*Response {
	if x != nil {
		return x.Responses
	}
	return nil
}

type Response_Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Response_Error_Code `protobuf:"varint,1,opt,name=code,proto3,enum=socket.Response_Error_Code" json:"code,omitempty"`
	Message string              `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response_Error) Reset() {
	*x = Response_Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_socket_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response_Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response_Error) ProtoMessage() {}

func (x *Response_Error) ProtoReflect() protoreflect.Message {
	mi := &file_socket_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response_Error.ProtoReflect.Descriptor instead.
func (*Response_Error) Descriptor() ([]byte, []int) {
	return file_socket_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Response_Error) GetCode() Response_Error_Code {
	if x != nil {
		return x.Code
	}
	return Response_Error_UNKNOWN
}

func (x *Response_Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Response_Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId uint64                 `protobuf:"varint,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	Players  []*Response_PlayerJoin `protobuf:"bytes,2,rep,name=players,proto3" json:"players,omitempty"`
}

func (x *Response_Info) Reset() {
	*x = Response_Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_socket_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response_Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response_Info) ProtoMessage() {}

func (x *Response_Info) ProtoReflect() protoreflect.Message {
	mi := &file_socket_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response_Info.ProtoReflect.Descriptor instead.
func (*Response_Info) Descriptor() ([]byte, []int) {
	return file_socket_proto_rawDescGZIP(), []int{1, 2}
}

func (x *Response_Info) GetPlayerId() uint64 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

func (x *Response_Info) GetPlayers() []*Response_PlayerJoin {
	if x != nil {
		return x.Players
	}
	return nil
}

type Response_PlayerJoin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *Response_PlayerJoin) Reset() {
	*x = Response_PlayerJoin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_socket_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response_PlayerJoin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response_PlayerJoin) ProtoMessage() {}

func (x *Response_PlayerJoin) ProtoReflect() protoreflect.Message {
	mi := &file_socket_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response_PlayerJoin.ProtoReflect.Descriptor instead.
func (*Response_PlayerJoin) Descriptor() ([]byte, []int) {
	return file_socket_proto_rawDescGZIP(), []int{1, 3}
}

func (x *Response_PlayerJoin) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Response_PlayerJoin) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type Response_PlayerMove struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId uint64 `protobuf:"varint,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	Pos      *Pos   `protobuf:"bytes,2,opt,name=pos,proto3" json:"pos,omitempty"`
}

func (x *Response_PlayerMove) Reset() {
	*x = Response_PlayerMove{}
	if protoimpl.UnsafeEnabled {
		mi := &file_socket_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response_PlayerMove) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response_PlayerMove) ProtoMessage() {}

func (x *Response_PlayerMove) ProtoReflect() protoreflect.Message {
	mi := &file_socket_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response_PlayerMove.ProtoReflect.Descriptor instead.
func (*Response_PlayerMove) Descriptor() ([]byte, []int) {
	return file_socket_proto_rawDescGZIP(), []int{1, 4}
}

func (x *Response_PlayerMove) GetPlayerId() uint64 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

func (x *Response_PlayerMove) GetPos() *Pos {
	if x != nil {
		return x.Pos
	}
	return nil
}

var File_socket_proto protoreflect.FileDescriptor

var file_socket_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x50, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0b, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x4d, 0x6f, 0x76, 0x65, 0x42, 0x0b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x03, 0xf8, 0x42,
	0x01, 0x22, 0x94, 0x07, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35,
	0x0a, 0x04, 0x62, 0x75, 0x6c, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73,
	0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x42,
	0x75, 0x6c, 0x6b, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52,
	0x04, 0x62, 0x75, 0x6c, 0x6b, 0x12, 0x38, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x35, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x49, 0x6e, 0x66, 0x6f, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x48, 0x00,
	0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x48, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x5f, 0x6a, 0x6f, 0x69, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x6f,
	0x63, 0x6b, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x4a, 0x6f, 0x69, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02,
	0x10, 0x01, 0x48, 0x00, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4a, 0x6f, 0x69, 0x6e,
	0x12, 0x2c, 0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x6c, 0x65, 0x61, 0x76, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x48,
	0x00, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x12, 0x48,
	0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4d, 0x6f, 0x76, 0x65,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x0a, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x4d, 0x6f, 0x76, 0x65, 0x1a, 0x47, 0x0a, 0x04, 0x42, 0x75, 0x6c, 0x6b,
	0x12, 0x3f, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0f, 0xfa, 0x42, 0x0c, 0x92, 0x01, 0x09, 0x08, 0x01, 0x22,
	0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x09, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x73, 0x1a, 0xa8, 0x01, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x39, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x73, 0x6f, 0x63, 0x6b,
	0x65, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x41, 0x0a, 0x04, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x13,
	0x0a, 0x0f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53,
	0x54, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54,
	0x45, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x02, 0x1a, 0x72, 0x0a, 0x04,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00,
	0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x44, 0x0a, 0x07, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x6f,
	0x63, 0x6b, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x4a, 0x6f, 0x69, 0x6e, 0x42, 0x0d, 0xfa, 0x42, 0x0a, 0x92, 0x01, 0x07,
	0x22, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73,
	0x1a, 0x4c, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x17,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32,
	0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04,
	0x10, 0x01, 0x18, 0x20, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x5b,
	0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4d, 0x6f, 0x76, 0x65, 0x12, 0x24, 0x0a, 0x09,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x27, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x42, 0x0b, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x22, 0x21, 0x0a, 0x03, 0x50, 0x6f, 0x73, 0x12,
	0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a,
	0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x79, 0x42, 0x10, 0x5a, 0x0e, 0x70,
	0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_socket_proto_rawDescOnce sync.Once
	file_socket_proto_rawDescData = file_socket_proto_rawDesc
)

func file_socket_proto_rawDescGZIP() []byte {
	file_socket_proto_rawDescOnce.Do(func() {
		file_socket_proto_rawDescData = protoimpl.X.CompressGZIP(file_socket_proto_rawDescData)
	})
	return file_socket_proto_rawDescData
}

var file_socket_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_socket_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_socket_proto_goTypes = []interface{}{
	(Response_Error_Code)(0),    // 0: socket.Response.Error.Code
	(*Request)(nil),             // 1: socket.Request
	(*Response)(nil),            // 2: socket.Response
	(*Pos)(nil),                 // 3: socket.Pos
	(*Response_Bulk)(nil),       // 4: socket.Response.Bulk
	(*Response_Error)(nil),      // 5: socket.Response.Error
	(*Response_Info)(nil),       // 6: socket.Response.Info
	(*Response_PlayerJoin)(nil), // 7: socket.Response.PlayerJoin
	(*Response_PlayerMove)(nil), // 8: socket.Response.PlayerMove
}
var file_socket_proto_depIdxs = []int32{
	3,  // 0: socket.Request.player_move:type_name -> socket.Pos
	4,  // 1: socket.Response.bulk:type_name -> socket.Response.Bulk
	5,  // 2: socket.Response.error:type_name -> socket.Response.Error
	6,  // 3: socket.Response.info:type_name -> socket.Response.Info
	7,  // 4: socket.Response.player_join:type_name -> socket.Response.PlayerJoin
	8,  // 5: socket.Response.player_move:type_name -> socket.Response.PlayerMove
	2,  // 6: socket.Response.Bulk.responses:type_name -> socket.Response
	0,  // 7: socket.Response.Error.code:type_name -> socket.Response.Error.Code
	7,  // 8: socket.Response.Info.players:type_name -> socket.Response.PlayerJoin
	3,  // 9: socket.Response.PlayerMove.pos:type_name -> socket.Pos
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_socket_proto_init() }
func file_socket_proto_init() {
	if File_socket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_socket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_socket_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_socket_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pos); i {
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
		file_socket_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response_Bulk); i {
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
		file_socket_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response_Error); i {
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
		file_socket_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response_Info); i {
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
		file_socket_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response_PlayerJoin); i {
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
		file_socket_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response_PlayerMove); i {
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
	file_socket_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Request_PlayerMove)(nil),
	}
	file_socket_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Response_Bulk_)(nil),
		(*Response_Error_)(nil),
		(*Response_Info_)(nil),
		(*Response_PlayerJoin_)(nil),
		(*Response_PlayerLeave)(nil),
		(*Response_PlayerMove_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_socket_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_socket_proto_goTypes,
		DependencyIndexes: file_socket_proto_depIdxs,
		EnumInfos:         file_socket_proto_enumTypes,
		MessageInfos:      file_socket_proto_msgTypes,
	}.Build()
	File_socket_proto = out.File
	file_socket_proto_rawDesc = nil
	file_socket_proto_goTypes = nil
	file_socket_proto_depIdxs = nil
}
