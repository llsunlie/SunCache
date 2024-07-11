// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package chatTool

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

// ChatToolClient is the client API for ChatTool service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatToolClient interface {
	Get(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type chatToolClient struct {
	cc grpc.ClientConnInterface
}

func NewChatToolClient(cc grpc.ClientConnInterface) ChatToolClient {
	return &chatToolClient{cc}
}

func (c *chatToolClient) Get(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/chatTool.ChatTool/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatToolServer is the server API for ChatTool service.
// All implementations must embed UnimplementedChatToolServer
// for forward compatibility
type ChatToolServer interface {
	Get(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedChatToolServer()
}

// UnimplementedChatToolServer must be embedded to have forward compatible implementations.
type UnimplementedChatToolServer struct {
}

func (UnimplementedChatToolServer) Get(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedChatToolServer) mustEmbedUnimplementedChatToolServer() {}

// UnsafeChatToolServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatToolServer will
// result in compilation errors.
type UnsafeChatToolServer interface {
	mustEmbedUnimplementedChatToolServer()
}

func RegisterChatToolServer(s grpc.ServiceRegistrar, srv ChatToolServer) {
	s.RegisterService(&ChatTool_ServiceDesc, srv)
}

func _ChatTool_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatToolServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatTool.ChatTool/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatToolServer).Get(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatTool_ServiceDesc is the grpc.ServiceDesc for ChatTool service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatTool_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chatTool.ChatTool",
	HandlerType: (*ChatToolServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ChatTool_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chatTool.proto",
}