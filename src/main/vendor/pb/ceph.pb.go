// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ceph.proto

package pb

/*
SRC_DIR=./src/pb ; protoc -I=$SRC_DIR --go_out=plugins=grpc:$SRC_DIR ./src/pb/ceph.proto
*/

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

type CephRequest struct {
	Instruction          string   `protobuf:"bytes,2,opt,name=instruction,proto3" json:"instruction,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CephRequest) Reset()         { *m = CephRequest{} }
func (m *CephRequest) String() string { return proto.CompactTextString(m) }
func (*CephRequest) ProtoMessage()    {}
func (*CephRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ceph_a5b76a92560c3f78, []int{0}
}
func (m *CephRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CephRequest.Unmarshal(m, b)
}
func (m *CephRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CephRequest.Marshal(b, m, deterministic)
}
func (dst *CephRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CephRequest.Merge(dst, src)
}
func (m *CephRequest) XXX_Size() int {
	return xxx_messageInfo_CephRequest.Size(m)
}
func (m *CephRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CephRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CephRequest proto.InternalMessageInfo

func (m *CephRequest) GetInstruction() string {
	if m != nil {
		return m.Instruction
	}
	return ""
}

type CephResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CephResponse) Reset()         { *m = CephResponse{} }
func (m *CephResponse) String() string { return proto.CompactTextString(m) }
func (*CephResponse) ProtoMessage()    {}
func (*CephResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ceph_a5b76a92560c3f78, []int{1}
}
func (m *CephResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CephResponse.Unmarshal(m, b)
}
func (m *CephResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CephResponse.Marshal(b, m, deterministic)
}
func (dst *CephResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CephResponse.Merge(dst, src)
}
func (m *CephResponse) XXX_Size() int {
	return xxx_messageInfo_CephResponse.Size(m)
}
func (m *CephResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CephResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CephResponse proto.InternalMessageInfo

func (m *CephResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type CephHeartBeatRequest struct {
	Ping                 string   `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CephHeartBeatRequest) Reset()         { *m = CephHeartBeatRequest{} }
func (m *CephHeartBeatRequest) String() string { return proto.CompactTextString(m) }
func (*CephHeartBeatRequest) ProtoMessage()    {}
func (*CephHeartBeatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ceph_a5b76a92560c3f78, []int{2}
}
func (m *CephHeartBeatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CephHeartBeatRequest.Unmarshal(m, b)
}
func (m *CephHeartBeatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CephHeartBeatRequest.Marshal(b, m, deterministic)
}
func (dst *CephHeartBeatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CephHeartBeatRequest.Merge(dst, src)
}
func (m *CephHeartBeatRequest) XXX_Size() int {
	return xxx_messageInfo_CephHeartBeatRequest.Size(m)
}
func (m *CephHeartBeatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CephHeartBeatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CephHeartBeatRequest proto.InternalMessageInfo

func (m *CephHeartBeatRequest) GetPing() string {
	if m != nil {
		return m.Ping
	}
	return ""
}

type CephHeartBeatResponse struct {
	Pong                 string   `protobuf:"bytes,2,opt,name=pong,proto3" json:"pong,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CephHeartBeatResponse) Reset()         { *m = CephHeartBeatResponse{} }
func (m *CephHeartBeatResponse) String() string { return proto.CompactTextString(m) }
func (*CephHeartBeatResponse) ProtoMessage()    {}
func (*CephHeartBeatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ceph_a5b76a92560c3f78, []int{3}
}
func (m *CephHeartBeatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CephHeartBeatResponse.Unmarshal(m, b)
}
func (m *CephHeartBeatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CephHeartBeatResponse.Marshal(b, m, deterministic)
}
func (dst *CephHeartBeatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CephHeartBeatResponse.Merge(dst, src)
}
func (m *CephHeartBeatResponse) XXX_Size() int {
	return xxx_messageInfo_CephHeartBeatResponse.Size(m)
}
func (m *CephHeartBeatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CephHeartBeatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CephHeartBeatResponse proto.InternalMessageInfo

func (m *CephHeartBeatResponse) GetPong() string {
	if m != nil {
		return m.Pong
	}
	return ""
}

func init() {
	proto.RegisterType((*CephRequest)(nil), "pb.CephRequest")
	proto.RegisterType((*CephResponse)(nil), "pb.CephResponse")
	proto.RegisterType((*CephHeartBeatRequest)(nil), "pb.CephHeartBeatRequest")
	proto.RegisterType((*CephHeartBeatResponse)(nil), "pb.CephHeartBeatResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CephServiceClient is the client API for CephService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CephServiceClient interface {
	GetClusterStatus(ctx context.Context, in *CephRequest, opts ...grpc.CallOption) (*CephResponse, error)
	HeartBeat(ctx context.Context, in *CephHeartBeatRequest, opts ...grpc.CallOption) (*CephHeartBeatResponse, error)
}

type cephServiceClient struct {
	cc *grpc.ClientConn
}

func NewCephServiceClient(cc *grpc.ClientConn) CephServiceClient {
	return &cephServiceClient{cc}
}

func (c *cephServiceClient) GetClusterStatus(ctx context.Context, in *CephRequest, opts ...grpc.CallOption) (*CephResponse, error) {
	out := new(CephResponse)
	err := c.cc.Invoke(ctx, "/pb.CephService/GetClusterStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cephServiceClient) HeartBeat(ctx context.Context, in *CephHeartBeatRequest, opts ...grpc.CallOption) (*CephHeartBeatResponse, error) {
	out := new(CephHeartBeatResponse)
	err := c.cc.Invoke(ctx, "/pb.CephService/HeartBeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CephServiceServer is the server API for CephService service.
type CephServiceServer interface {
	GetClusterStatus(context.Context, *CephRequest) (*CephResponse, error)
	HeartBeat(context.Context, *CephHeartBeatRequest) (*CephHeartBeatResponse, error)
}

func RegisterCephServiceServer(s *grpc.Server, srv CephServiceServer) {
	s.RegisterService(&_CephService_serviceDesc, srv)
}

func _CephService_GetClusterStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CephRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CephServiceServer).GetClusterStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CephService/GetClusterStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CephServiceServer).GetClusterStatus(ctx, req.(*CephRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CephService_HeartBeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CephHeartBeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CephServiceServer).HeartBeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CephService/HeartBeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CephServiceServer).HeartBeat(ctx, req.(*CephHeartBeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CephService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CephService",
	HandlerType: (*CephServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetClusterStatus",
			Handler:    _CephService_GetClusterStatus_Handler,
		},
		{
			MethodName: "HeartBeat",
			Handler:    _CephService_HeartBeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ceph.proto",
}

func init() { proto.RegisterFile("ceph.proto", fileDescriptor_ceph_a5b76a92560c3f78) }

var fileDescriptor_ceph_a5b76a92560c3f78 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x51, 0x4b, 0x87, 0x30,
	0x14, 0xc5, 0xf1, 0x4f, 0x08, 0x5e, 0x83, 0x64, 0x54, 0x98, 0x4f, 0xb2, 0x87, 0x88, 0x02, 0x83,
	0xa2, 0xf7, 0xc8, 0x87, 0x7a, 0xd6, 0x4f, 0xa0, 0x72, 0xd1, 0x81, 0x6c, 0x6b, 0xbb, 0xeb, 0x23,
	0xf4, 0xb9, 0xc3, 0x39, 0x4b, 0xa4, 0xb7, 0xbb, 0x7b, 0x7e, 0x3b, 0xe7, 0x6c, 0x00, 0x03, 0xea,
	0xa9, 0xd2, 0x46, 0x91, 0x62, 0x27, 0xdd, 0xf3, 0x47, 0x48, 0x6b, 0xd4, 0x53, 0x83, 0x9f, 0x0e,
	0x2d, 0xb1, 0x12, 0x52, 0x21, 0x2d, 0x19, 0x37, 0x90, 0x50, 0x32, 0x3f, 0x95, 0xd1, 0x5d, 0xd2,
	0xec, 0x57, 0xfc, 0x16, 0xce, 0xd7, 0x0b, 0x56, 0x2b, 0x69, 0x91, 0x5d, 0x43, 0x6c, 0xd0, 0xba,
	0x99, 0xf2, 0xc8, 0xc3, 0xe1, 0xc4, 0xef, 0xe1, 0x72, 0xe1, 0x3e, 0xb0, 0x33, 0xf4, 0x86, 0x1d,
	0x6d, 0x09, 0x0c, 0xce, 0xb4, 0x90, 0x63, 0xa0, 0xfd, 0xcc, 0x1f, 0xe0, 0xea, 0xc0, 0x06, 0xf3,
	0x05, 0x56, 0x72, 0x0c, 0x3d, 0xfc, 0xfc, 0xf4, 0x1d, 0xad, 0x95, 0x5b, 0x34, 0x5f, 0x62, 0x40,
	0xf6, 0x02, 0xd9, 0x3b, 0x52, 0x3d, 0x3b, 0x4b, 0x68, 0x5a, 0xea, 0xc8, 0x59, 0x76, 0x51, 0xe9,
	0xbe, 0xda, 0xbd, 0xab, 0xc8, 0xfe, 0x16, 0xc1, 0xfa, 0x15, 0x92, 0xdf, 0x3c, 0x96, 0x6f, 0xf2,
	0xb1, 0x6e, 0x71, 0xf3, 0x8f, 0xb2, 0x3a, 0xf4, 0xb1, 0xff, 0xc5, 0xe7, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x8c, 0x86, 0xcd, 0x20, 0x53, 0x01, 0x00, 0x00,
}
