// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Users.proto

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

type UserScoreRequest struct {
	Users                []*UserInfo `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UserScoreRequest) Reset()         { *m = UserScoreRequest{} }
func (m *UserScoreRequest) String() string { return proto.CompactTextString(m) }
func (*UserScoreRequest) ProtoMessage()    {}
func (*UserScoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_59e1cb5f5c9ed3f9, []int{0}
}

func (m *UserScoreRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserScoreRequest.Unmarshal(m, b)
}
func (m *UserScoreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserScoreRequest.Marshal(b, m, deterministic)
}
func (m *UserScoreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserScoreRequest.Merge(m, src)
}
func (m *UserScoreRequest) XXX_Size() int {
	return xxx_messageInfo_UserScoreRequest.Size(m)
}
func (m *UserScoreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserScoreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserScoreRequest proto.InternalMessageInfo

func (m *UserScoreRequest) GetUsers() []*UserInfo {
	if m != nil {
		return m.Users
	}
	return nil
}

type UserScoreResponse struct {
	Users                []*UserInfo `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UserScoreResponse) Reset()         { *m = UserScoreResponse{} }
func (m *UserScoreResponse) String() string { return proto.CompactTextString(m) }
func (*UserScoreResponse) ProtoMessage()    {}
func (*UserScoreResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_59e1cb5f5c9ed3f9, []int{1}
}

func (m *UserScoreResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserScoreResponse.Unmarshal(m, b)
}
func (m *UserScoreResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserScoreResponse.Marshal(b, m, deterministic)
}
func (m *UserScoreResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserScoreResponse.Merge(m, src)
}
func (m *UserScoreResponse) XXX_Size() int {
	return xxx_messageInfo_UserScoreResponse.Size(m)
}
func (m *UserScoreResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserScoreResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserScoreResponse proto.InternalMessageInfo

func (m *UserScoreResponse) GetUsers() []*UserInfo {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*UserScoreRequest)(nil), "services.UserScoreRequest")
	proto.RegisterType((*UserScoreResponse)(nil), "services.UserScoreResponse")
}

func init() { proto.RegisterFile("Users.proto", fileDescriptor_59e1cb5f5c9ed3f9) }

var fileDescriptor_59e1cb5f5c9ed3f9 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x0e, 0x2d, 0x4e, 0x2d,
	0x2a, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x28, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e,
	0x2d, 0x96, 0xe2, 0xf1, 0xcd, 0x4f, 0x49, 0xcd, 0x81, 0x8a, 0x2b, 0xd9, 0x70, 0x09, 0x80, 0x94,
	0x05, 0x27, 0xe7, 0x17, 0xa5, 0x06, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x69, 0x70, 0xb1,
	0x96, 0x82, 0xb4, 0x4a, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x1b, 0x09, 0xe9, 0xc1, 0xf4, 0xea, 0x81,
	0x94, 0x7a, 0xe6, 0xa5, 0xe5, 0x07, 0x41, 0x14, 0x28, 0xd9, 0x72, 0x09, 0x22, 0xe9, 0x2e, 0x2e,
	0xc8, 0xcf, 0x2b, 0x4e, 0x25, 0x5e, 0xbb, 0x51, 0x0f, 0x13, 0xc4, 0x91, 0xc1, 0x10, 0x05, 0x42,
	0xee, 0x5c, 0x3c, 0xee, 0xa9, 0x25, 0x70, 0x13, 0x85, 0xa4, 0x50, 0xb5, 0x22, 0x3b, 0x52, 0x4a,
	0x1a, 0xab, 0x1c, 0xd4, 0x09, 0x91, 0x5c, 0x52, 0xc8, 0x06, 0x39, 0x55, 0x82, 0x6c, 0x48, 0x2d,
	0x0a, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0x25, 0xdb, 0x58, 0x25, 0x06, 0x03, 0x46, 0x4c, 0xa3, 0x9d,
	0x73, 0x32, 0x53, 0xf3, 0x4a, 0x28, 0x36, 0x5a, 0x83, 0x31, 0x89, 0x0d, 0x1c, 0x25, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x42, 0x2f, 0x7e, 0xb9, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	GetUserScore(ctx context.Context, in *UserScoreRequest, opts ...grpc.CallOption) (*UserScoreResponse, error)
	GetUserScoreByServerStream(ctx context.Context, in *UserScoreRequest, opts ...grpc.CallOption) (UserService_GetUserScoreByServerStreamClient, error)
	GetUserScoreByClientStream(ctx context.Context, opts ...grpc.CallOption) (UserService_GetUserScoreByClientStreamClient, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUserScore(ctx context.Context, in *UserScoreRequest, opts ...grpc.CallOption) (*UserScoreResponse, error) {
	out := new(UserScoreResponse)
	err := c.cc.Invoke(ctx, "/services.UserService/GetUserScore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserScoreByServerStream(ctx context.Context, in *UserScoreRequest, opts ...grpc.CallOption) (UserService_GetUserScoreByServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_UserService_serviceDesc.Streams[0], "/services.UserService/GetUserScoreByServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceGetUserScoreByServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UserService_GetUserScoreByServerStreamClient interface {
	Recv() (*UserScoreResponse, error)
	grpc.ClientStream
}

type userServiceGetUserScoreByServerStreamClient struct {
	grpc.ClientStream
}

func (x *userServiceGetUserScoreByServerStreamClient) Recv() (*UserScoreResponse, error) {
	m := new(UserScoreResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userServiceClient) GetUserScoreByClientStream(ctx context.Context, opts ...grpc.CallOption) (UserService_GetUserScoreByClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_UserService_serviceDesc.Streams[1], "/services.UserService/GetUserScoreByClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceGetUserScoreByClientStreamClient{stream}
	return x, nil
}

type UserService_GetUserScoreByClientStreamClient interface {
	Send(*UserScoreRequest) error
	CloseAndRecv() (*UserScoreResponse, error)
	grpc.ClientStream
}

type userServiceGetUserScoreByClientStreamClient struct {
	grpc.ClientStream
}

func (x *userServiceGetUserScoreByClientStreamClient) Send(m *UserScoreRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *userServiceGetUserScoreByClientStreamClient) CloseAndRecv() (*UserScoreResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UserScoreResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	GetUserScore(context.Context, *UserScoreRequest) (*UserScoreResponse, error)
	GetUserScoreByServerStream(*UserScoreRequest, UserService_GetUserScoreByServerStreamServer) error
	GetUserScoreByClientStream(UserService_GetUserScoreByClientStreamServer) error
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) GetUserScore(ctx context.Context, req *UserScoreRequest) (*UserScoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserScore not implemented")
}
func (*UnimplementedUserServiceServer) GetUserScoreByServerStream(req *UserScoreRequest, srv UserService_GetUserScoreByServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUserScoreByServerStream not implemented")
}
func (*UnimplementedUserServiceServer) GetUserScoreByClientStream(srv UserService_GetUserScoreByClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUserScoreByClientStream not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_GetUserScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.UserService/GetUserScore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserScore(ctx, req.(*UserScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserScoreByServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UserScoreRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserServiceServer).GetUserScoreByServerStream(m, &userServiceGetUserScoreByServerStreamServer{stream})
}

type UserService_GetUserScoreByServerStreamServer interface {
	Send(*UserScoreResponse) error
	grpc.ServerStream
}

type userServiceGetUserScoreByServerStreamServer struct {
	grpc.ServerStream
}

func (x *userServiceGetUserScoreByServerStreamServer) Send(m *UserScoreResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _UserService_GetUserScoreByClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UserServiceServer).GetUserScoreByClientStream(&userServiceGetUserScoreByClientStreamServer{stream})
}

type UserService_GetUserScoreByClientStreamServer interface {
	SendAndClose(*UserScoreResponse) error
	Recv() (*UserScoreRequest, error)
	grpc.ServerStream
}

type userServiceGetUserScoreByClientStreamServer struct {
	grpc.ServerStream
}

func (x *userServiceGetUserScoreByClientStreamServer) SendAndClose(m *UserScoreResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *userServiceGetUserScoreByClientStreamServer) Recv() (*UserScoreRequest, error) {
	m := new(UserScoreRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserScore",
			Handler:    _UserService_GetUserScore_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetUserScoreByServerStream",
			Handler:       _UserService_GetUserScoreByServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetUserScoreByClientStream",
			Handler:       _UserService_GetUserScoreByClientStream_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "Users.proto",
}
