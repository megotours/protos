// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: like_service.proto

package like

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

// LikeServiceClient is the client API for LikeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LikeServiceClient interface {
	Exists(ctx context.Context, in *ExistsRequest, opts ...grpc.CallOption) (*ExistsResponse, error)
	Like(ctx context.Context, in *LikeRequest, opts ...grpc.CallOption) (*LikeResponse, error)
	Count(ctx context.Context, in *CountRequest, opts ...grpc.CallOption) (*CountResponse, error)
}

type likeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLikeServiceClient(cc grpc.ClientConnInterface) LikeServiceClient {
	return &likeServiceClient{cc}
}

func (c *likeServiceClient) Exists(ctx context.Context, in *ExistsRequest, opts ...grpc.CallOption) (*ExistsResponse, error) {
	out := new(ExistsResponse)
	err := c.cc.Invoke(ctx, "/like.LikeService/Exists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *likeServiceClient) Like(ctx context.Context, in *LikeRequest, opts ...grpc.CallOption) (*LikeResponse, error) {
	out := new(LikeResponse)
	err := c.cc.Invoke(ctx, "/like.LikeService/Like", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *likeServiceClient) Count(ctx context.Context, in *CountRequest, opts ...grpc.CallOption) (*CountResponse, error) {
	out := new(CountResponse)
	err := c.cc.Invoke(ctx, "/like.LikeService/Count", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LikeServiceServer is the server API for LikeService service.
// All implementations must embed UnimplementedLikeServiceServer
// for forward compatibility
type LikeServiceServer interface {
	Exists(context.Context, *ExistsRequest) (*ExistsResponse, error)
	Like(context.Context, *LikeRequest) (*LikeResponse, error)
	Count(context.Context, *CountRequest) (*CountResponse, error)
	mustEmbedUnimplementedLikeServiceServer()
}

// UnimplementedLikeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLikeServiceServer struct {
}

func (UnimplementedLikeServiceServer) Exists(context.Context, *ExistsRequest) (*ExistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exists not implemented")
}
func (UnimplementedLikeServiceServer) Like(context.Context, *LikeRequest) (*LikeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Like not implemented")
}
func (UnimplementedLikeServiceServer) Count(context.Context, *CountRequest) (*CountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Count not implemented")
}
func (UnimplementedLikeServiceServer) mustEmbedUnimplementedLikeServiceServer() {}

// UnsafeLikeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LikeServiceServer will
// result in compilation errors.
type UnsafeLikeServiceServer interface {
	mustEmbedUnimplementedLikeServiceServer()
}

func RegisterLikeServiceServer(s grpc.ServiceRegistrar, srv LikeServiceServer) {
	s.RegisterService(&LikeService_ServiceDesc, srv)
}

func _LikeService_Exists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServiceServer).Exists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/like.LikeService/Exists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServiceServer).Exists(ctx, req.(*ExistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LikeService_Like_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServiceServer).Like(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/like.LikeService/Like",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServiceServer).Like(ctx, req.(*LikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LikeService_Count_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServiceServer).Count(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/like.LikeService/Count",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServiceServer).Count(ctx, req.(*CountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LikeService_ServiceDesc is the grpc.ServiceDesc for LikeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LikeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "like.LikeService",
	HandlerType: (*LikeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Exists",
			Handler:    _LikeService_Exists_Handler,
		},
		{
			MethodName: "Like",
			Handler:    _LikeService_Like_Handler,
		},
		{
			MethodName: "Count",
			Handler:    _LikeService_Count_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "like_service.proto",
}