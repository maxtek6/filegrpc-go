// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: services/download/download.proto

package download

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

// DownloadClient is the client API for Download service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DownloadClient interface {
	Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*InitResponse, error)
	Recv(ctx context.Context, in *RecvRequest, opts ...grpc.CallOption) (Download_RecvClient, error)
}

type downloadClient struct {
	cc grpc.ClientConnInterface
}

func NewDownloadClient(cc grpc.ClientConnInterface) DownloadClient {
	return &downloadClient{cc}
}

func (c *downloadClient) Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*InitResponse, error) {
	out := new(InitResponse)
	err := c.cc.Invoke(ctx, "/filegrpc.Download/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *downloadClient) Recv(ctx context.Context, in *RecvRequest, opts ...grpc.CallOption) (Download_RecvClient, error) {
	stream, err := c.cc.NewStream(ctx, &Download_ServiceDesc.Streams[0], "/filegrpc.Download/Recv", opts...)
	if err != nil {
		return nil, err
	}
	x := &downloadRecvClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Download_RecvClient interface {
	Recv() (*RecvResponse, error)
	grpc.ClientStream
}

type downloadRecvClient struct {
	grpc.ClientStream
}

func (x *downloadRecvClient) Recv() (*RecvResponse, error) {
	m := new(RecvResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DownloadServer is the server API for Download service.
// All implementations must embed UnimplementedDownloadServer
// for forward compatibility
type DownloadServer interface {
	Init(context.Context, *InitRequest) (*InitResponse, error)
	Recv(*RecvRequest, Download_RecvServer) error
	mustEmbedUnimplementedDownloadServer()
}

// UnimplementedDownloadServer must be embedded to have forward compatible implementations.
type UnimplementedDownloadServer struct {
}

func (UnimplementedDownloadServer) Init(context.Context, *InitRequest) (*InitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (UnimplementedDownloadServer) Recv(*RecvRequest, Download_RecvServer) error {
	return status.Errorf(codes.Unimplemented, "method Recv not implemented")
}
func (UnimplementedDownloadServer) mustEmbedUnimplementedDownloadServer() {}

// UnsafeDownloadServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DownloadServer will
// result in compilation errors.
type UnsafeDownloadServer interface {
	mustEmbedUnimplementedDownloadServer()
}

func RegisterDownloadServer(s grpc.ServiceRegistrar, srv DownloadServer) {
	s.RegisterService(&Download_ServiceDesc, srv)
}

func _Download_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownloadServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filegrpc.Download/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownloadServer).Init(ctx, req.(*InitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Download_Recv_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RecvRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DownloadServer).Recv(m, &downloadRecvServer{stream})
}

type Download_RecvServer interface {
	Send(*RecvResponse) error
	grpc.ServerStream
}

type downloadRecvServer struct {
	grpc.ServerStream
}

func (x *downloadRecvServer) Send(m *RecvResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Download_ServiceDesc is the grpc.ServiceDesc for Download service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Download_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "filegrpc.Download",
	HandlerType: (*DownloadServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _Download_Init_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Recv",
			Handler:       _Download_Recv_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "services/download/download.proto",
}
