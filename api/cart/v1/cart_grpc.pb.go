// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.6
// source: cart/v1/cart.proto

package v1

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
	Cart_AddCommodity_FullMethodName    = "/helloworld.v1.Cart/AddCommodity"
	Cart_UpdateCommodity_FullMethodName = "/helloworld.v1.Cart/UpdateCommodity"
	Cart_DeleteCommodity_FullMethodName = "/helloworld.v1.Cart/DeleteCommodity"
	Cart_GetCart_FullMethodName         = "/helloworld.v1.Cart/GetCart"
	Cart_ClearCart_FullMethodName       = "/helloworld.v1.Cart/ClearCart"
)

// CartClient is the client API for Cart service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CartClient interface {
	AddCommodity(ctx context.Context, in *AddCommodityRequest, opts ...grpc.CallOption) (*AddCommodityReply, error)
	UpdateCommodity(ctx context.Context, in *UpdateCommodityRequest, opts ...grpc.CallOption) (*UpdateCommodityReply, error)
	DeleteCommodity(ctx context.Context, in *DeleteCommodityRequest, opts ...grpc.CallOption) (*DeleteCommodityReply, error)
	GetCart(ctx context.Context, in *GetCartRequest, opts ...grpc.CallOption) (*GetCartReply, error)
	ClearCart(ctx context.Context, in *ClearCartRequest, opts ...grpc.CallOption) (*ClearCartReply, error)
}

type cartClient struct {
	cc grpc.ClientConnInterface
}

func NewCartClient(cc grpc.ClientConnInterface) CartClient {
	return &cartClient{cc}
}

func (c *cartClient) AddCommodity(ctx context.Context, in *AddCommodityRequest, opts ...grpc.CallOption) (*AddCommodityReply, error) {
	out := new(AddCommodityReply)
	err := c.cc.Invoke(ctx, Cart_AddCommodity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartClient) UpdateCommodity(ctx context.Context, in *UpdateCommodityRequest, opts ...grpc.CallOption) (*UpdateCommodityReply, error) {
	out := new(UpdateCommodityReply)
	err := c.cc.Invoke(ctx, Cart_UpdateCommodity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartClient) DeleteCommodity(ctx context.Context, in *DeleteCommodityRequest, opts ...grpc.CallOption) (*DeleteCommodityReply, error) {
	out := new(DeleteCommodityReply)
	err := c.cc.Invoke(ctx, Cart_DeleteCommodity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartClient) GetCart(ctx context.Context, in *GetCartRequest, opts ...grpc.CallOption) (*GetCartReply, error) {
	out := new(GetCartReply)
	err := c.cc.Invoke(ctx, Cart_GetCart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartClient) ClearCart(ctx context.Context, in *ClearCartRequest, opts ...grpc.CallOption) (*ClearCartReply, error) {
	out := new(ClearCartReply)
	err := c.cc.Invoke(ctx, Cart_ClearCart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartServer is the server API for Cart service.
// All implementations must embed UnimplementedCartServer
// for forward compatibility
type CartServer interface {
	AddCommodity(context.Context, *AddCommodityRequest) (*AddCommodityReply, error)
	UpdateCommodity(context.Context, *UpdateCommodityRequest) (*UpdateCommodityReply, error)
	DeleteCommodity(context.Context, *DeleteCommodityRequest) (*DeleteCommodityReply, error)
	GetCart(context.Context, *GetCartRequest) (*GetCartReply, error)
	ClearCart(context.Context, *ClearCartRequest) (*ClearCartReply, error)
	mustEmbedUnimplementedCartServer()
}

// UnimplementedCartServer must be embedded to have forward compatible implementations.
type UnimplementedCartServer struct {
}

func (UnimplementedCartServer) AddCommodity(context.Context, *AddCommodityRequest) (*AddCommodityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCommodity not implemented")
}
func (UnimplementedCartServer) UpdateCommodity(context.Context, *UpdateCommodityRequest) (*UpdateCommodityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCommodity not implemented")
}
func (UnimplementedCartServer) DeleteCommodity(context.Context, *DeleteCommodityRequest) (*DeleteCommodityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCommodity not implemented")
}
func (UnimplementedCartServer) GetCart(context.Context, *GetCartRequest) (*GetCartReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCart not implemented")
}
func (UnimplementedCartServer) ClearCart(context.Context, *ClearCartRequest) (*ClearCartReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearCart not implemented")
}
func (UnimplementedCartServer) mustEmbedUnimplementedCartServer() {}

// UnsafeCartServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CartServer will
// result in compilation errors.
type UnsafeCartServer interface {
	mustEmbedUnimplementedCartServer()
}

func RegisterCartServer(s grpc.ServiceRegistrar, srv CartServer) {
	s.RegisterService(&Cart_ServiceDesc, srv)
}

func _Cart_AddCommodity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCommodityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServer).AddCommodity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cart_AddCommodity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).AddCommodity(ctx, req.(*AddCommodityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cart_UpdateCommodity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCommodityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServer).UpdateCommodity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cart_UpdateCommodity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).UpdateCommodity(ctx, req.(*UpdateCommodityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cart_DeleteCommodity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommodityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServer).DeleteCommodity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cart_DeleteCommodity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).DeleteCommodity(ctx, req.(*DeleteCommodityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cart_GetCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServer).GetCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cart_GetCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).GetCart(ctx, req.(*GetCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cart_ClearCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServer).ClearCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cart_ClearCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).ClearCart(ctx, req.(*ClearCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Cart_ServiceDesc is the grpc.ServiceDesc for Cart service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cart_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.v1.Cart",
	HandlerType: (*CartServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCommodity",
			Handler:    _Cart_AddCommodity_Handler,
		},
		{
			MethodName: "UpdateCommodity",
			Handler:    _Cart_UpdateCommodity_Handler,
		},
		{
			MethodName: "DeleteCommodity",
			Handler:    _Cart_DeleteCommodity_Handler,
		},
		{
			MethodName: "GetCart",
			Handler:    _Cart_GetCart_Handler,
		},
		{
			MethodName: "ClearCart",
			Handler:    _Cart_ClearCart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cart/v1/cart.proto",
}