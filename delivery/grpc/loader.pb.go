// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: loader.proto

package grpc

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type LoaderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *LoaderResponse) Reset() {
	*x = LoaderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loader_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoaderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoaderResponse) ProtoMessage() {}

func (x *LoaderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_loader_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoaderResponse.ProtoReflect.Descriptor instead.
func (*LoaderResponse) Descriptor() ([]byte, []int) {
	return file_loader_proto_rawDescGZIP(), []int{0}
}

func (x *LoaderResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_loader_proto protoreflect.FileDescriptor

var file_loader_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x0e, 0x4c,
	0x6f, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x73, 0x0a, 0x0d, 0x4c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x4c, 0x6f, 0x61, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x3b,
	0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_loader_proto_rawDescOnce sync.Once
	file_loader_proto_rawDescData = file_loader_proto_rawDesc
)

func file_loader_proto_rawDescGZIP() []byte {
	file_loader_proto_rawDescOnce.Do(func() {
		file_loader_proto_rawDescData = protoimpl.X.CompressGZIP(file_loader_proto_rawDescData)
	})
	return file_loader_proto_rawDescData
}

var file_loader_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_loader_proto_goTypes = []interface{}{
	(*LoaderResponse)(nil), // 0: LoaderResponse
	(*empty.Empty)(nil),    // 1: google.protobuf.Empty
}
var file_loader_proto_depIdxs = []int32{
	1, // 0: LoaderService.Start:input_type -> google.protobuf.Empty
	1, // 1: LoaderService.Check:input_type -> google.protobuf.Empty
	0, // 2: LoaderService.Start:output_type -> LoaderResponse
	0, // 3: LoaderService.Check:output_type -> LoaderResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_loader_proto_init() }
func file_loader_proto_init() {
	if File_loader_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_loader_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoaderResponse); i {
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
			RawDescriptor: file_loader_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_loader_proto_goTypes,
		DependencyIndexes: file_loader_proto_depIdxs,
		MessageInfos:      file_loader_proto_msgTypes,
	}.Build()
	File_loader_proto = out.File
	file_loader_proto_rawDesc = nil
	file_loader_proto_goTypes = nil
	file_loader_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LoaderServiceClient is the client API for LoaderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoaderServiceClient interface {
	Start(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*LoaderResponse, error)
	Check(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*LoaderResponse, error)
}

type loaderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLoaderServiceClient(cc grpc.ClientConnInterface) LoaderServiceClient {
	return &loaderServiceClient{cc}
}

func (c *loaderServiceClient) Start(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*LoaderResponse, error) {
	out := new(LoaderResponse)
	err := c.cc.Invoke(ctx, "/LoaderService/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loaderServiceClient) Check(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*LoaderResponse, error) {
	out := new(LoaderResponse)
	err := c.cc.Invoke(ctx, "/LoaderService/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoaderServiceServer is the server API for LoaderService service.
type LoaderServiceServer interface {
	Start(context.Context, *empty.Empty) (*LoaderResponse, error)
	Check(context.Context, *empty.Empty) (*LoaderResponse, error)
}

// UnimplementedLoaderServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLoaderServiceServer struct {
}

func (*UnimplementedLoaderServiceServer) Start(context.Context, *empty.Empty) (*LoaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (*UnimplementedLoaderServiceServer) Check(context.Context, *empty.Empty) (*LoaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}

func RegisterLoaderServiceServer(s *grpc.Server, srv LoaderServiceServer) {
	s.RegisterService(&_LoaderService_serviceDesc, srv)
}

func _LoaderService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoaderServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LoaderService/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoaderServiceServer).Start(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoaderService_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoaderServiceServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LoaderService/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoaderServiceServer).Check(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoaderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "LoaderService",
	HandlerType: (*LoaderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _LoaderService_Start_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _LoaderService_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "loader.proto",
}