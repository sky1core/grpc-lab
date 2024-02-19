// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: memo.proto

package proto

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

const (
	MemoService_CreateMemo_FullMethodName = "/memo.MemoService/CreateMemo"
	MemoService_GetMemo_FullMethodName    = "/memo.MemoService/GetMemo"
	MemoService_ListMemos_FullMethodName  = "/memo.MemoService/ListMemos"
	MemoService_UpdateMemo_FullMethodName = "/memo.MemoService/UpdateMemo"
	MemoService_DeleteMemo_FullMethodName = "/memo.MemoService/DeleteMemo"
)

// MemoServiceClient is the client API for MemoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MemoServiceClient interface {
	CreateMemo(ctx context.Context, in *CreateMemoRequest, opts ...grpc.CallOption) (*MemoResponse, error)
	GetMemo(ctx context.Context, in *GetMemoRequest, opts ...grpc.CallOption) (*MemoResponse, error)
	ListMemos(ctx context.Context, in *ListMemosRequest, opts ...grpc.CallOption) (*ListMemosResponse, error)
	UpdateMemo(ctx context.Context, in *UpdateMemoRequest, opts ...grpc.CallOption) (*MemoResponse, error)
	DeleteMemo(ctx context.Context, in *DeleteMemoRequest, opts ...grpc.CallOption) (*DeleteMemoResponse, error)
}

type memoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMemoServiceClient(cc grpc.ClientConnInterface) MemoServiceClient {
	return &memoServiceClient{cc}
}

func (c *memoServiceClient) CreateMemo(ctx context.Context, in *CreateMemoRequest, opts ...grpc.CallOption) (*MemoResponse, error) {
	out := new(MemoResponse)
	err := c.cc.Invoke(ctx, MemoService_CreateMemo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memoServiceClient) GetMemo(ctx context.Context, in *GetMemoRequest, opts ...grpc.CallOption) (*MemoResponse, error) {
	out := new(MemoResponse)
	err := c.cc.Invoke(ctx, MemoService_GetMemo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memoServiceClient) ListMemos(ctx context.Context, in *ListMemosRequest, opts ...grpc.CallOption) (*ListMemosResponse, error) {
	out := new(ListMemosResponse)
	err := c.cc.Invoke(ctx, MemoService_ListMemos_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memoServiceClient) UpdateMemo(ctx context.Context, in *UpdateMemoRequest, opts ...grpc.CallOption) (*MemoResponse, error) {
	out := new(MemoResponse)
	err := c.cc.Invoke(ctx, MemoService_UpdateMemo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memoServiceClient) DeleteMemo(ctx context.Context, in *DeleteMemoRequest, opts ...grpc.CallOption) (*DeleteMemoResponse, error) {
	out := new(DeleteMemoResponse)
	err := c.cc.Invoke(ctx, MemoService_DeleteMemo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MemoServiceServer is the server API for MemoService service.
// All implementations must embed UnimplementedMemoServiceServer
// for forward compatibility
type MemoServiceServer interface {
	CreateMemo(context.Context, *CreateMemoRequest) (*MemoResponse, error)
	GetMemo(context.Context, *GetMemoRequest) (*MemoResponse, error)
	ListMemos(context.Context, *ListMemosRequest) (*ListMemosResponse, error)
	UpdateMemo(context.Context, *UpdateMemoRequest) (*MemoResponse, error)
	DeleteMemo(context.Context, *DeleteMemoRequest) (*DeleteMemoResponse, error)
	mustEmbedUnimplementedMemoServiceServer()
}

// UnimplementedMemoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMemoServiceServer struct {
}

func (UnimplementedMemoServiceServer) CreateMemo(context.Context, *CreateMemoRequest) (*MemoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMemo not implemented")
}
func (UnimplementedMemoServiceServer) GetMemo(context.Context, *GetMemoRequest) (*MemoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMemo not implemented")
}
func (UnimplementedMemoServiceServer) ListMemos(context.Context, *ListMemosRequest) (*ListMemosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMemos not implemented")
}
func (UnimplementedMemoServiceServer) UpdateMemo(context.Context, *UpdateMemoRequest) (*MemoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMemo not implemented")
}
func (UnimplementedMemoServiceServer) DeleteMemo(context.Context, *DeleteMemoRequest) (*DeleteMemoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMemo not implemented")
}
func (UnimplementedMemoServiceServer) mustEmbedUnimplementedMemoServiceServer() {}

// UnsafeMemoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MemoServiceServer will
// result in compilation errors.
type UnsafeMemoServiceServer interface {
	mustEmbedUnimplementedMemoServiceServer()
}

func RegisterMemoServiceServer(s grpc.ServiceRegistrar, srv MemoServiceServer) {
	s.RegisterService(&MemoService_ServiceDesc, srv)
}

func _MemoService_CreateMemo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMemoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemoServiceServer).CreateMemo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MemoService_CreateMemo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemoServiceServer).CreateMemo(ctx, req.(*CreateMemoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemoService_GetMemo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMemoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemoServiceServer).GetMemo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MemoService_GetMemo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemoServiceServer).GetMemo(ctx, req.(*GetMemoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemoService_ListMemos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMemosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemoServiceServer).ListMemos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MemoService_ListMemos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemoServiceServer).ListMemos(ctx, req.(*ListMemosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemoService_UpdateMemo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMemoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemoServiceServer).UpdateMemo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MemoService_UpdateMemo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemoServiceServer).UpdateMemo(ctx, req.(*UpdateMemoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemoService_DeleteMemo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMemoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemoServiceServer).DeleteMemo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MemoService_DeleteMemo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemoServiceServer).DeleteMemo(ctx, req.(*DeleteMemoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MemoService_ServiceDesc is the grpc.ServiceDesc for MemoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MemoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "memo.MemoService",
	HandlerType: (*MemoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMemo",
			Handler:    _MemoService_CreateMemo_Handler,
		},
		{
			MethodName: "GetMemo",
			Handler:    _MemoService_GetMemo_Handler,
		},
		{
			MethodName: "ListMemos",
			Handler:    _MemoService_ListMemos_Handler,
		},
		{
			MethodName: "UpdateMemo",
			Handler:    _MemoService_UpdateMemo_Handler,
		},
		{
			MethodName: "DeleteMemo",
			Handler:    _MemoService_DeleteMemo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "memo.proto",
}
