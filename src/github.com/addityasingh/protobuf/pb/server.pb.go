// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/github.com/addityasingh/protobuf/pb/server.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EchoRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoRequest) Reset()         { *m = EchoRequest{} }
func (m *EchoRequest) String() string { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()    {}
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_server_9189e8ef8985863a, []int{0}
}
func (m *EchoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoRequest.Unmarshal(m, b)
}
func (m *EchoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoRequest.Marshal(b, m, deterministic)
}
func (dst *EchoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoRequest.Merge(dst, src)
}
func (m *EchoRequest) XXX_Size() int {
	return xxx_messageInfo_EchoRequest.Size(m)
}
func (m *EchoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EchoRequest proto.InternalMessageInfo

func (m *EchoRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *EchoRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type EchoResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoResponse) Reset()         { *m = EchoResponse{} }
func (m *EchoResponse) String() string { return proto.CompactTextString(m) }
func (*EchoResponse) ProtoMessage()    {}
func (*EchoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_server_9189e8ef8985863a, []int{1}
}
func (m *EchoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoResponse.Unmarshal(m, b)
}
func (m *EchoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoResponse.Marshal(b, m, deterministic)
}
func (dst *EchoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoResponse.Merge(dst, src)
}
func (m *EchoResponse) XXX_Size() int {
	return xxx_messageInfo_EchoResponse.Size(m)
}
func (m *EchoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EchoResponse proto.InternalMessageInfo

func (m *EchoResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*EchoRequest)(nil), "pb.EchoRequest")
	proto.RegisterType((*EchoResponse)(nil), "pb.EchoResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EchoClient is the client API for Echo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EchoClient interface {
	Reply(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
}

type echoClient struct {
	cc *grpc.ClientConn
}

func NewEchoClient(cc *grpc.ClientConn) EchoClient {
	return &echoClient{cc}
}

func (c *echoClient) Reply(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, "/pb.Echo/Reply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoServer is the server API for Echo service.
type EchoServer interface {
	Reply(context.Context, *EchoRequest) (*EchoResponse, error)
}

func RegisterEchoServer(s *grpc.Server, srv EchoServer) {
	s.RegisterService(&_Echo_serviceDesc, srv)
}

func _Echo_Reply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).Reply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Echo/Reply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).Reply(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Echo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Echo",
	HandlerType: (*EchoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Reply",
			Handler:    _Echo_Reply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/github.com/addityasingh/protobuf/pb/server.proto",
}

func init() {
	proto.RegisterFile("src/github.com/addityasingh/protobuf/pb/server.proto", fileDescriptor_server_9189e8ef8985863a)
}

var fileDescriptor_server_9189e8ef8985863a = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8e, 0x41, 0xca, 0xc2, 0x30,
	0x10, 0x85, 0xff, 0x86, 0x5f, 0xc5, 0x51, 0x54, 0xb2, 0x2a, 0xae, 0xa4, 0xab, 0x2e, 0x24, 0x01,
	0x2d, 0x78, 0x02, 0x2f, 0xd0, 0x1b, 0x34, 0xcd, 0xd8, 0x06, 0x6c, 0x13, 0x33, 0xa9, 0xd0, 0xdb,
	0x8b, 0x2d, 0x62, 0x5d, 0xce, 0xc7, 0x37, 0xef, 0x3d, 0xc8, 0xc8, 0x97, 0xb2, 0x32, 0xa1, 0xee,
	0x94, 0x28, 0x6d, 0x23, 0x0b, 0xad, 0x4d, 0xe8, 0x0b, 0x32, 0x6d, 0x55, 0x4b, 0xe7, 0x6d, 0xb0,
	0xaa, 0xbb, 0x49, 0xa7, 0x24, 0xa1, 0x7f, 0xa2, 0x17, 0x03, 0xe2, 0xcc, 0xa9, 0xe4, 0x02, 0xab,
	0x6b, 0x59, 0xdb, 0x1c, 0x1f, 0x1d, 0x52, 0xe0, 0x1b, 0x60, 0x46, 0xc7, 0xd1, 0x21, 0x4a, 0x97,
	0x39, 0x33, 0x9a, 0xc7, 0xb0, 0x68, 0x90, 0xa8, 0xa8, 0x30, 0x66, 0x03, 0xfc, 0x9c, 0x49, 0x0a,
	0xeb, 0xf1, 0x91, 0x9c, 0x6d, 0x09, 0xa7, 0x66, 0xf4, 0x63, 0x9e, 0x32, 0xf8, 0x7f, 0x9b, 0xfc,
	0x08, 0xb3, 0x1c, 0xdd, 0xbd, 0xe7, 0x5b, 0xe1, 0x94, 0x98, 0xb4, 0xee, 0x77, 0x5f, 0x30, 0xa6,
	0x25, 0x7f, 0x6a, 0x3e, 0x6c, 0x3c, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x69, 0xcb, 0x80,
	0xdb, 0x00, 0x00, 0x00,
}
