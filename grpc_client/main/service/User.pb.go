// 指定protobuf的版本，proto3是最新的语法版本

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: User.proto

package service

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

type UserScoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*UserInfo `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *UserScoreRequest) Reset() {
	*x = UserScoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_User_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserScoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserScoreRequest) ProtoMessage() {}

func (x *UserScoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_User_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserScoreRequest.ProtoReflect.Descriptor instead.
func (*UserScoreRequest) Descriptor() ([]byte, []int) {
	return file_User_proto_rawDescGZIP(), []int{0}
}

func (x *UserScoreRequest) GetUsers() []*UserInfo {
	if x != nil {
		return x.Users
	}
	return nil
}

type UserScoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*UserInfo `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *UserScoreResponse) Reset() {
	*x = UserScoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_User_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserScoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserScoreResponse) ProtoMessage() {}

func (x *UserScoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_User_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserScoreResponse.ProtoReflect.Descriptor instead.
func (*UserScoreResponse) Descriptor() ([]byte, []int) {
	return file_User_proto_rawDescGZIP(), []int{1}
}

func (x *UserScoreResponse) GetUsers() []*UserInfo {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_User_proto protoreflect.FileDescriptor

var file_User_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x22, 0x3c, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x32, 0xd9,
	0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45,
	0x0a, 0x0c, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x19,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x1a, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x79, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x55, 0x0a, 0x1a,
	0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x79, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x28, 0x01, 0x12, 0x55, 0x0a, 0x18, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x42, 0x79, 0x42, 0x6f, 0x74, 0x68, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12,
	0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_User_proto_rawDescOnce sync.Once
	file_User_proto_rawDescData = file_User_proto_rawDesc
)

func file_User_proto_rawDescGZIP() []byte {
	file_User_proto_rawDescOnce.Do(func() {
		file_User_proto_rawDescData = protoimpl.X.CompressGZIP(file_User_proto_rawDescData)
	})
	return file_User_proto_rawDescData
}

var file_User_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_User_proto_goTypes = []interface{}{
	(*UserScoreRequest)(nil),  // 0: service.UserScoreRequest
	(*UserScoreResponse)(nil), // 1: service.UserScoreResponse
	(*UserInfo)(nil),          // 2: service.UserInfo
}
var file_User_proto_depIdxs = []int32{
	2, // 0: service.UserScoreRequest.users:type_name -> service.UserInfo
	2, // 1: service.UserScoreResponse.users:type_name -> service.UserInfo
	0, // 2: service.UserService.getUserScore:input_type -> service.UserScoreRequest
	0, // 3: service.UserService.getUserScoreByServerStream:input_type -> service.UserScoreRequest
	0, // 4: service.UserService.getUserScoreByClientStream:input_type -> service.UserScoreRequest
	0, // 5: service.UserService.getUserScoreByBothStream:input_type -> service.UserScoreRequest
	1, // 6: service.UserService.getUserScore:output_type -> service.UserScoreResponse
	1, // 7: service.UserService.getUserScoreByServerStream:output_type -> service.UserScoreResponse
	1, // 8: service.UserService.getUserScoreByClientStream:output_type -> service.UserScoreResponse
	1, // 9: service.UserService.getUserScoreByBothStream:output_type -> service.UserScoreResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_User_proto_init() }
func file_User_proto_init() {
	if File_User_proto != nil {
		return
	}
	file_Models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_User_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserScoreRequest); i {
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
		file_User_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserScoreResponse); i {
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
			RawDescriptor: file_User_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_User_proto_goTypes,
		DependencyIndexes: file_User_proto_depIdxs,
		MessageInfos:      file_User_proto_msgTypes,
	}.Build()
	File_User_proto = out.File
	file_User_proto_rawDesc = nil
	file_User_proto_goTypes = nil
	file_User_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	GetUserScore(ctx context.Context, in *UserScoreRequest, opts ...grpc.CallOption) (*UserScoreResponse, error)
	//   服务端流
	GetUserScoreByServerStream(ctx context.Context, in *UserScoreRequest, opts ...grpc.CallOption) (UserService_GetUserScoreByServerStreamClient, error)
	//   客户端流
	GetUserScoreByClientStream(ctx context.Context, opts ...grpc.CallOption) (UserService_GetUserScoreByClientStreamClient, error)
	//  双向流
	GetUserScoreByBothStream(ctx context.Context, opts ...grpc.CallOption) (UserService_GetUserScoreByBothStreamClient, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUserScore(ctx context.Context, in *UserScoreRequest, opts ...grpc.CallOption) (*UserScoreResponse, error) {
	out := new(UserScoreResponse)
	err := c.cc.Invoke(ctx, "/service.UserService/getUserScore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserScoreByServerStream(ctx context.Context, in *UserScoreRequest, opts ...grpc.CallOption) (UserService_GetUserScoreByServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_UserService_serviceDesc.Streams[0], "/service.UserService/getUserScoreByServerStream", opts...)
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
	stream, err := c.cc.NewStream(ctx, &_UserService_serviceDesc.Streams[1], "/service.UserService/getUserScoreByClientStream", opts...)
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

func (c *userServiceClient) GetUserScoreByBothStream(ctx context.Context, opts ...grpc.CallOption) (UserService_GetUserScoreByBothStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_UserService_serviceDesc.Streams[2], "/service.UserService/getUserScoreByBothStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceGetUserScoreByBothStreamClient{stream}
	return x, nil
}

type UserService_GetUserScoreByBothStreamClient interface {
	Send(*UserScoreRequest) error
	Recv() (*UserScoreResponse, error)
	grpc.ClientStream
}

type userServiceGetUserScoreByBothStreamClient struct {
	grpc.ClientStream
}

func (x *userServiceGetUserScoreByBothStreamClient) Send(m *UserScoreRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *userServiceGetUserScoreByBothStreamClient) Recv() (*UserScoreResponse, error) {
	m := new(UserScoreResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	GetUserScore(context.Context, *UserScoreRequest) (*UserScoreResponse, error)
	//   服务端流
	GetUserScoreByServerStream(*UserScoreRequest, UserService_GetUserScoreByServerStreamServer) error
	//   客户端流
	GetUserScoreByClientStream(UserService_GetUserScoreByClientStreamServer) error
	//  双向流
	GetUserScoreByBothStream(UserService_GetUserScoreByBothStreamServer) error
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) GetUserScore(context.Context, *UserScoreRequest) (*UserScoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserScore not implemented")
}
func (*UnimplementedUserServiceServer) GetUserScoreByServerStream(*UserScoreRequest, UserService_GetUserScoreByServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUserScoreByServerStream not implemented")
}
func (*UnimplementedUserServiceServer) GetUserScoreByClientStream(UserService_GetUserScoreByClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUserScoreByClientStream not implemented")
}
func (*UnimplementedUserServiceServer) GetUserScoreByBothStream(UserService_GetUserScoreByBothStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUserScoreByBothStream not implemented")
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
		FullMethod: "/service.UserService/GetUserScore",
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

func _UserService_GetUserScoreByBothStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UserServiceServer).GetUserScoreByBothStream(&userServiceGetUserScoreByBothStreamServer{stream})
}

type UserService_GetUserScoreByBothStreamServer interface {
	Send(*UserScoreResponse) error
	Recv() (*UserScoreRequest, error)
	grpc.ServerStream
}

type userServiceGetUserScoreByBothStreamServer struct {
	grpc.ServerStream
}

func (x *userServiceGetUserScoreByBothStreamServer) Send(m *UserScoreResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *userServiceGetUserScoreByBothStreamServer) Recv() (*UserScoreRequest, error) {
	m := new(UserScoreRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getUserScore",
			Handler:    _UserService_GetUserScore_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "getUserScoreByServerStream",
			Handler:       _UserService_GetUserScoreByServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "getUserScoreByClientStream",
			Handler:       _UserService_GetUserScoreByClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "getUserScoreByBothStream",
			Handler:       _UserService_GetUserScoreByBothStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "User.proto",
}
