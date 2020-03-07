// Code generated by protoc-gen-go. DO NOT EDIT.
// source: idProduce/idProduce.proto

package idProduce

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

type IdProduceRequest struct {
	Len                  uint32   `protobuf:"varint,1,opt,name=len,proto3" json:"len,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdProduceRequest) Reset()         { *m = IdProduceRequest{} }
func (m *IdProduceRequest) String() string { return proto.CompactTextString(m) }
func (*IdProduceRequest) ProtoMessage()    {}
func (*IdProduceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_533d7d59978f5e65, []int{0}
}

func (m *IdProduceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdProduceRequest.Unmarshal(m, b)
}
func (m *IdProduceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdProduceRequest.Marshal(b, m, deterministic)
}
func (m *IdProduceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdProduceRequest.Merge(m, src)
}
func (m *IdProduceRequest) XXX_Size() int {
	return xxx_messageInfo_IdProduceRequest.Size(m)
}
func (m *IdProduceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IdProduceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IdProduceRequest proto.InternalMessageInfo

func (m *IdProduceRequest) GetLen() uint32 {
	if m != nil {
		return m.Len
	}
	return 0
}

type IdProduceResponse struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Ids                  []uint64 `protobuf:"varint,2,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdProduceResponse) Reset()         { *m = IdProduceResponse{} }
func (m *IdProduceResponse) String() string { return proto.CompactTextString(m) }
func (*IdProduceResponse) ProtoMessage()    {}
func (*IdProduceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_533d7d59978f5e65, []int{1}
}

func (m *IdProduceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdProduceResponse.Unmarshal(m, b)
}
func (m *IdProduceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdProduceResponse.Marshal(b, m, deterministic)
}
func (m *IdProduceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdProduceResponse.Merge(m, src)
}
func (m *IdProduceResponse) XXX_Size() int {
	return xxx_messageInfo_IdProduceResponse.Size(m)
}
func (m *IdProduceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IdProduceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IdProduceResponse proto.InternalMessageInfo

func (m *IdProduceResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *IdProduceResponse) GetIds() []uint64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func init() {
	proto.RegisterType((*IdProduceRequest)(nil), "idProduce.IdProduceRequest")
	proto.RegisterType((*IdProduceResponse)(nil), "idProduce.IdProduceResponse")
}

func init() { proto.RegisterFile("idProduce/idProduce.proto", fileDescriptor_533d7d59978f5e65) }

var fileDescriptor_533d7d59978f5e65 = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0x4c, 0x09, 0x28,
	0xca, 0x4f, 0x29, 0x4d, 0x4e, 0xd5, 0x87, 0xb3, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x38,
	0xe1, 0x02, 0x4a, 0x2a, 0x5c, 0x02, 0x9e, 0x30, 0x4e, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89,
	0x90, 0x00, 0x17, 0x73, 0x4e, 0x6a, 0x9e, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x6f, 0x10, 0x88, 0xa9,
	0x64, 0xc9, 0x25, 0x88, 0xa4, 0xaa, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x88, 0x8b, 0x25,
	0x39, 0x3f, 0x25, 0x15, 0xaa, 0x0e, 0xcc, 0x06, 0x69, 0xcd, 0x4c, 0x29, 0x96, 0x60, 0x52, 0x60,
	0xd6, 0x60, 0x09, 0x02, 0x31, 0x8d, 0xa2, 0xb8, 0xd8, 0x3c, 0x52, 0x73, 0x0a, 0x52, 0x8b, 0x84,
	0x02, 0xb8, 0xf8, 0xdd, 0x53, 0x4b, 0x5c, 0x32, 0x8b, 0x4b, 0x8a, 0x32, 0x93, 0x4a, 0x4b, 0x52,
	0x3d, 0x53, 0x84, 0xa4, 0xf5, 0x10, 0x4e, 0x43, 0x77, 0x86, 0x94, 0x0c, 0x76, 0x49, 0x88, 0xed,
	0x1a, 0x8c, 0x06, 0x8c, 0x4e, 0xe2, 0xa7, 0x98, 0x44, 0x7c, 0x33, 0x93, 0x8b, 0xf2, 0x03, 0x40,
	0xde, 0x8a, 0x81, 0x2b, 0x4a, 0x62, 0x03, 0xfb, 0xd3, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xd4,
	0xe7, 0x52, 0xa3, 0x04, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelperClient is the client API for Helper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelperClient interface {
	GetDistributeId(ctx context.Context, opts ...grpc.CallOption) (Helper_GetDistributeIdClient, error)
}

type helperClient struct {
	cc *grpc.ClientConn
}

func NewHelperClient(cc *grpc.ClientConn) HelperClient {
	return &helperClient{cc}
}

func (c *helperClient) GetDistributeId(ctx context.Context, opts ...grpc.CallOption) (Helper_GetDistributeIdClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Helper_serviceDesc.Streams[0], "/idProduce.Helper/GetDistributeId", opts...)
	if err != nil {
		return nil, err
	}
	x := &helperGetDistributeIdClient{stream}
	return x, nil
}

type Helper_GetDistributeIdClient interface {
	Send(*IdProduceRequest) error
	Recv() (*IdProduceResponse, error)
	grpc.ClientStream
}

type helperGetDistributeIdClient struct {
	grpc.ClientStream
}

func (x *helperGetDistributeIdClient) Send(m *IdProduceRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helperGetDistributeIdClient) Recv() (*IdProduceResponse, error) {
	m := new(IdProduceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelperServer is the server API for Helper service.
type HelperServer interface {
	GetDistributeId(Helper_GetDistributeIdServer) error
}

// UnimplementedHelperServer can be embedded to have forward compatible implementations.
type UnimplementedHelperServer struct {
}

func (*UnimplementedHelperServer) GetDistributeId(srv Helper_GetDistributeIdServer) error {
	return status.Errorf(codes.Unimplemented, "method GetDistributeId not implemented")
}

func RegisterHelperServer(s *grpc.Server, srv HelperServer) {
	s.RegisterService(&_Helper_serviceDesc, srv)
}

func _Helper_GetDistributeId_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelperServer).GetDistributeId(&helperGetDistributeIdServer{stream})
}

type Helper_GetDistributeIdServer interface {
	Send(*IdProduceResponse) error
	Recv() (*IdProduceRequest, error)
	grpc.ServerStream
}

type helperGetDistributeIdServer struct {
	grpc.ServerStream
}

func (x *helperGetDistributeIdServer) Send(m *IdProduceResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helperGetDistributeIdServer) Recv() (*IdProduceRequest, error) {
	m := new(IdProduceRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Helper_serviceDesc = grpc.ServiceDesc{
	ServiceName: "idProduce.Helper",
	HandlerType: (*HelperServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetDistributeId",
			Handler:       _Helper_GetDistributeId_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "idProduce/idProduce.proto",
}
