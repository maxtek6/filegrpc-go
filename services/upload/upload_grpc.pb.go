// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: services/upload/upload.proto

package upload

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UploadClient is the client API for Upload service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UploadClient interface {
	Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*InitResponse, error)
	Send(ctx context.Context, opts ...grpc.CallOption) (Upload_SendClient, error)
}

type uploadClient struct {
	cc grpc.ClientConnInterface
}

func NewUploadClient(cc grpc.ClientConnInterface) UploadClient {
	return &uploadClient{cc}
}

func (c *uploadClient) Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*InitResponse, error) {
	out := new(InitResponse)
	err := c.cc.Invoke(ctx, "/filegrpc.Upload/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uploadClient) Send(ctx context.Context, opts ...grpc.CallOption) (Upload_SendClient, error) {
	stream, err := c.cc.NewStream(ctx, &Upload_ServiceDesc.Streams[0], "/filegrpc.Upload/Send", opts...)
	if err != nil {
		return nil, err
	}
	x := &uploadSendClient{stream}
	return x, nil
}

type Upload_SendClient interface {
	Send(*SendRequest) error
	CloseAndRecv() (*SendResponse, error)
	grpc.ClientStream
}

type uploadSendClient struct {
	grpc.ClientStream
}

func (x *uploadSendClient) Send(m *SendRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *uploadSendClient) CloseAndRecv() (*SendResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SendResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UploadServer is the server API for Upload service.
// All implementations must embed UnimplementedUploadServer
// for forward compatibility
type UploadServer interface {
	Init(context.Context, *InitRequest) (*InitResponse, error)
	Send(Upload_SendServer) error
	mustEmbedUnimplementedUploadServer()
}

// UnimplementedUploadServer must be embedded to have forward compatible implementations.
type UnimplementedUploadServer struct {
}

func (UnimplementedUploadServer) Init(context.Context, *InitRequest) (*InitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (UnimplementedUploadServer) Send(Upload_SendServer) error {
	return status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedUploadServer) mustEmbedUnimplementedUploadServer() {}

// UnsafeUploadServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UploadServer will
// result in compilation errors.
type UnsafeUploadServer interface {
	mustEmbedUnimplementedUploadServer()
}

func RegisterUploadServer(s grpc.ServiceRegistrar, srv UploadServer) {
	s.RegisterService(&Upload_ServiceDesc, srv)
}

func _Upload_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploadServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filegrpc.Upload/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploadServer).Init(ctx, req.(*InitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Upload_Send_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UploadServer).Send(&uploadSendServer{stream})
}

type Upload_SendServer interface {
	SendAndClose(*SendResponse) error
	Recv() (*SendRequest, error)
	grpc.ServerStream
}

type uploadSendServer struct {
	grpc.ServerStream
}

func (x *uploadSendServer) SendAndClose(m *SendResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *uploadSendServer) Recv() (*SendRequest, error) {
	m := new(SendRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Upload_ServiceDesc is the grpc.ServiceDesc for Upload service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Upload_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "filegrpc.Upload",
	HandlerType: (*UploadServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _Upload_Init_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Send",
			Handler:       _Upload_Send_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "services/upload/upload.proto",
}
