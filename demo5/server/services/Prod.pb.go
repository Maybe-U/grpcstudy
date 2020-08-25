// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Prod.proto

package services

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

type ProdRequest struct {
	ProdId               int32    `protobuf:"varint,1,opt,name=prod_id,json=prodId,proto3" json:"prod_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProdRequest) Reset()         { *m = ProdRequest{} }
func (m *ProdRequest) String() string { return proto.CompactTextString(m) }
func (*ProdRequest) ProtoMessage()    {}
func (*ProdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b02cd6816510a0e, []int{0}
}

func (m *ProdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProdRequest.Unmarshal(m, b)
}
func (m *ProdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProdRequest.Marshal(b, m, deterministic)
}
func (m *ProdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProdRequest.Merge(m, src)
}
func (m *ProdRequest) XXX_Size() int {
	return xxx_messageInfo_ProdRequest.Size(m)
}
func (m *ProdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProdRequest proto.InternalMessageInfo

func (m *ProdRequest) GetProdId() int32 {
	if m != nil {
		return m.ProdId
	}
	return 0
}

type ProdResponse struct {
	ProdStock            int32    `protobuf:"varint,1,opt,name=prod_stock,json=prodStock,proto3" json:"prod_stock,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProdResponse) Reset()         { *m = ProdResponse{} }
func (m *ProdResponse) String() string { return proto.CompactTextString(m) }
func (*ProdResponse) ProtoMessage()    {}
func (*ProdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b02cd6816510a0e, []int{1}
}

func (m *ProdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProdResponse.Unmarshal(m, b)
}
func (m *ProdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProdResponse.Marshal(b, m, deterministic)
}
func (m *ProdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProdResponse.Merge(m, src)
}
func (m *ProdResponse) XXX_Size() int {
	return xxx_messageInfo_ProdResponse.Size(m)
}
func (m *ProdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProdResponse proto.InternalMessageInfo

func (m *ProdResponse) GetProdStock() int32 {
	if m != nil {
		return m.ProdStock
	}
	return 0
}

func init() {
	proto.RegisterType((*ProdRequest)(nil), "services.ProdRequest")
	proto.RegisterType((*ProdResponse)(nil), "services.ProdResponse")
}

func init() { proto.RegisterFile("Prod.proto", fileDescriptor_8b02cd6816510a0e) }

var fileDescriptor_8b02cd6816510a0e = []byte{
	// 149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x0a, 0x28, 0xca, 0x4f,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x28, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x2d,
	0x56, 0x52, 0xe3, 0xe2, 0x06, 0x89, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x89, 0x73,
	0xb1, 0x17, 0x14, 0xe5, 0xa7, 0xc4, 0x67, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0xb1,
	0x81, 0xb8, 0x9e, 0x29, 0x4a, 0xba, 0x5c, 0x3c, 0x10, 0x75, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9,
	0x42, 0xb2, 0x5c, 0x5c, 0x60, 0x85, 0xc5, 0x25, 0xf9, 0xc9, 0xd9, 0x50, 0xb5, 0x9c, 0x20, 0x91,
	0x60, 0x90, 0x80, 0x91, 0x0f, 0xc4, 0xd8, 0x60, 0x88, 0x35, 0x42, 0xb6, 0x5c, 0x3c, 0xee, 0xa9,
	0x25, 0x01, 0x30, 0x69, 0x21, 0x51, 0x3d, 0x98, 0x03, 0xf4, 0x90, 0x6c, 0x97, 0x12, 0x43, 0x17,
	0x86, 0x58, 0x96, 0xc4, 0x06, 0x76, 0xb5, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x78, 0x03,
	0x6e, 0xc3, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProdServiceClient is the client API for ProdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProdServiceClient interface {
	GetProdStock(ctx context.Context, in *ProdRequest, opts ...grpc.CallOption) (*ProdResponse, error)
}

type prodServiceClient struct {
	cc *grpc.ClientConn
}

func NewProdServiceClient(cc *grpc.ClientConn) ProdServiceClient {
	return &prodServiceClient{cc}
}

func (c *prodServiceClient) GetProdStock(ctx context.Context, in *ProdRequest, opts ...grpc.CallOption) (*ProdResponse, error) {
	out := new(ProdResponse)
	err := c.cc.Invoke(ctx, "/services.ProdService/GetProdStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProdServiceServer is the server API for ProdService service.
type ProdServiceServer interface {
	GetProdStock(context.Context, *ProdRequest) (*ProdResponse, error)
}

// UnimplementedProdServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProdServiceServer struct {
}

func (*UnimplementedProdServiceServer) GetProdStock(ctx context.Context, req *ProdRequest) (*ProdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProdStock not implemented")
}

func RegisterProdServiceServer(s *grpc.Server, srv ProdServiceServer) {
	s.RegisterService(&_ProdService_serviceDesc, srv)
}

func _ProdService_GetProdStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProdServiceServer).GetProdStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.ProdService/GetProdStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProdServiceServer).GetProdStock(ctx, req.(*ProdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProdService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.ProdService",
	HandlerType: (*ProdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProdStock",
			Handler:    _ProdService_GetProdStock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Prod.proto",
}
