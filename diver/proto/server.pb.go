// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server.proto

package diver

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

type GetTotalNumParams struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTotalNumParams) Reset()         { *m = GetTotalNumParams{} }
func (m *GetTotalNumParams) String() string { return proto.CompactTextString(m) }
func (*GetTotalNumParams) ProtoMessage()    {}
func (*GetTotalNumParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{0}
}

func (m *GetTotalNumParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTotalNumParams.Unmarshal(m, b)
}
func (m *GetTotalNumParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTotalNumParams.Marshal(b, m, deterministic)
}
func (m *GetTotalNumParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTotalNumParams.Merge(m, src)
}
func (m *GetTotalNumParams) XXX_Size() int {
	return xxx_messageInfo_GetTotalNumParams.Size(m)
}
func (m *GetTotalNumParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTotalNumParams.DiscardUnknown(m)
}

var xxx_messageInfo_GetTotalNumParams proto.InternalMessageInfo

type AddNumParams struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddNumParams) Reset()         { *m = AddNumParams{} }
func (m *AddNumParams) String() string { return proto.CompactTextString(m) }
func (*AddNumParams) ProtoMessage()    {}
func (*AddNumParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{1}
}

func (m *AddNumParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddNumParams.Unmarshal(m, b)
}
func (m *AddNumParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddNumParams.Marshal(b, m, deterministic)
}
func (m *AddNumParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddNumParams.Merge(m, src)
}
func (m *AddNumParams) XXX_Size() int {
	return xxx_messageInfo_AddNumParams.Size(m)
}
func (m *AddNumParams) XXX_DiscardUnknown() {
	xxx_messageInfo_AddNumParams.DiscardUnknown(m)
}

var xxx_messageInfo_AddNumParams proto.InternalMessageInfo

func (m *AddNumParams) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type TotalNum struct {
	Total                int32    `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TotalNum) Reset()         { *m = TotalNum{} }
func (m *TotalNum) String() string { return proto.CompactTextString(m) }
func (*TotalNum) ProtoMessage()    {}
func (*TotalNum) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{2}
}

func (m *TotalNum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TotalNum.Unmarshal(m, b)
}
func (m *TotalNum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TotalNum.Marshal(b, m, deterministic)
}
func (m *TotalNum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TotalNum.Merge(m, src)
}
func (m *TotalNum) XXX_Size() int {
	return xxx_messageInfo_TotalNum.Size(m)
}
func (m *TotalNum) XXX_DiscardUnknown() {
	xxx_messageInfo_TotalNum.DiscardUnknown(m)
}

var xxx_messageInfo_TotalNum proto.InternalMessageInfo

func (m *TotalNum) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterType((*GetTotalNumParams)(nil), "diver.getTotalNumParams")
	proto.RegisterType((*AddNumParams)(nil), "diver.addNumParams")
	proto.RegisterType((*TotalNum)(nil), "diver.totalNum")
}

func init() { proto.RegisterFile("server.proto", fileDescriptor_ad098daeda4239f7) }

var fileDescriptor_ad098daeda4239f7 = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0x2d, 0x2a,
	0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0xc9, 0x2c, 0x4b, 0x2d, 0x52,
	0x12, 0xe6, 0x12, 0x4c, 0x4f, 0x2d, 0x09, 0xc9, 0x2f, 0x49, 0xcc, 0xf1, 0x2b, 0xcd, 0x0d, 0x48,
	0x2c, 0x4a, 0xcc, 0x2d, 0x56, 0x52, 0xe3, 0xe2, 0x49, 0x4c, 0x49, 0x81, 0xf3, 0x85, 0xc4, 0xb8,
	0xd8, 0xf2, 0x4a, 0x73, 0x93, 0x52, 0x8b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0xa0, 0x3c,
	0x25, 0x05, 0x2e, 0x8e, 0x12, 0xa8, 0x4e, 0x21, 0x11, 0x2e, 0x56, 0x30, 0x1b, 0xaa, 0x04, 0xc2,
	0x31, 0xaa, 0xe5, 0xe2, 0x85, 0x98, 0x14, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0x64, 0xc0,
	0xc5, 0x06, 0x11, 0x10, 0x12, 0xd6, 0x03, 0xbb, 0x40, 0x0f, 0xd9, 0x26, 0x29, 0x7e, 0xa8, 0x20,
	0xcc, 0x58, 0x25, 0x06, 0x21, 0x2b, 0x2e, 0x6e, 0x24, 0x17, 0x0a, 0x49, 0x40, 0x55, 0x60, 0xb8,
	0x1a, 0x8b, 0xde, 0x24, 0x36, 0xb0, 0x5f, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x40, 0x68,
	0xec, 0xe7, 0xfb, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AddNumServiceClient is the client API for AddNumService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AddNumServiceClient interface {
	AddNum(ctx context.Context, in *AddNumParams, opts ...grpc.CallOption) (*TotalNum, error)
	GetTotalNum(ctx context.Context, in *GetTotalNumParams, opts ...grpc.CallOption) (*TotalNum, error)
}

type addNumServiceClient struct {
	cc *grpc.ClientConn
}

func NewAddNumServiceClient(cc *grpc.ClientConn) AddNumServiceClient {
	return &addNumServiceClient{cc}
}

func (c *addNumServiceClient) AddNum(ctx context.Context, in *AddNumParams, opts ...grpc.CallOption) (*TotalNum, error) {
	out := new(TotalNum)
	err := c.cc.Invoke(ctx, "/diver.addNumService/addNum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addNumServiceClient) GetTotalNum(ctx context.Context, in *GetTotalNumParams, opts ...grpc.CallOption) (*TotalNum, error) {
	out := new(TotalNum)
	err := c.cc.Invoke(ctx, "/diver.addNumService/getTotalNum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddNumServiceServer is the server API for AddNumService service.
type AddNumServiceServer interface {
	AddNum(context.Context, *AddNumParams) (*TotalNum, error)
	GetTotalNum(context.Context, *GetTotalNumParams) (*TotalNum, error)
}

// UnimplementedAddNumServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAddNumServiceServer struct {
}

func (*UnimplementedAddNumServiceServer) AddNum(ctx context.Context, req *AddNumParams) (*TotalNum, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNum not implemented")
}
func (*UnimplementedAddNumServiceServer) GetTotalNum(ctx context.Context, req *GetTotalNumParams) (*TotalNum, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTotalNum not implemented")
}

func RegisterAddNumServiceServer(s *grpc.Server, srv AddNumServiceServer) {
	s.RegisterService(&_AddNumService_serviceDesc, srv)
}

func _AddNumService_AddNum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNumParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddNumServiceServer).AddNum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/diver.addNumService/AddNum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddNumServiceServer).AddNum(ctx, req.(*AddNumParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddNumService_GetTotalNum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTotalNumParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddNumServiceServer).GetTotalNum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/diver.addNumService/GetTotalNum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddNumServiceServer).GetTotalNum(ctx, req.(*GetTotalNumParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _AddNumService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "diver.addNumService",
	HandlerType: (*AddNumServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "addNum",
			Handler:    _AddNumService_AddNum_Handler,
		},
		{
			MethodName: "getTotalNum",
			Handler:    _AddNumService_GetTotalNum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}
