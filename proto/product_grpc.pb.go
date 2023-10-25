// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/product.proto

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

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductServiceClient interface {
	// simple RPC (unary RPC)
	GetProduct(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*Product, error)
	GetProductById(ctx context.Context, in *ProductIdRequest, opts ...grpc.CallOption) (*Product, error)
	// server streaming RPC
	GetProductServerStreaming(ctx context.Context, in *ProductIdList, opts ...grpc.CallOption) (ProductService_GetProductServerStreamingClient, error)
	// client streaming RPC
	// rpc GetProductClientStreaming(stream ProductIdList) returns (ProductList);
	GetProductClientStreaming(ctx context.Context, opts ...grpc.CallOption) (ProductService_GetProductClientStreamingClient, error)
	// bidirectional RPC
	GetProductBidirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (ProductService_GetProductBidirectionalStreamingClient, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) GetProduct(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/product_service.ProductService/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetProductById(ctx context.Context, in *ProductIdRequest, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/product_service.ProductService/GetProductById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetProductServerStreaming(ctx context.Context, in *ProductIdList, opts ...grpc.CallOption) (ProductService_GetProductServerStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProductService_ServiceDesc.Streams[0], "/product_service.ProductService/GetProductServerStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &productServiceGetProductServerStreamingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProductService_GetProductServerStreamingClient interface {
	Recv() (*ProductIdResponse, error)
	grpc.ClientStream
}

type productServiceGetProductServerStreamingClient struct {
	grpc.ClientStream
}

func (x *productServiceGetProductServerStreamingClient) Recv() (*ProductIdResponse, error) {
	m := new(ProductIdResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *productServiceClient) GetProductClientStreaming(ctx context.Context, opts ...grpc.CallOption) (ProductService_GetProductClientStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProductService_ServiceDesc.Streams[1], "/product_service.ProductService/GetProductClientStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &productServiceGetProductClientStreamingClient{stream}
	return x, nil
}

type ProductService_GetProductClientStreamingClient interface {
	Send(*ProductIdRequest) error
	CloseAndRecv() (*ProductIdList, error)
	grpc.ClientStream
}

type productServiceGetProductClientStreamingClient struct {
	grpc.ClientStream
}

func (x *productServiceGetProductClientStreamingClient) Send(m *ProductIdRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *productServiceGetProductClientStreamingClient) CloseAndRecv() (*ProductIdList, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ProductIdList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *productServiceClient) GetProductBidirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (ProductService_GetProductBidirectionalStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProductService_ServiceDesc.Streams[2], "/product_service.ProductService/GetProductBidirectionalStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &productServiceGetProductBidirectionalStreamingClient{stream}
	return x, nil
}

type ProductService_GetProductBidirectionalStreamingClient interface {
	Send(*ProductIdRequest) error
	Recv() (*ProductIdResponse, error)
	grpc.ClientStream
}

type productServiceGetProductBidirectionalStreamingClient struct {
	grpc.ClientStream
}

func (x *productServiceGetProductBidirectionalStreamingClient) Send(m *ProductIdRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *productServiceGetProductBidirectionalStreamingClient) Recv() (*ProductIdResponse, error) {
	m := new(ProductIdResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProductServiceServer is the server API for ProductService service.
// All implementations must embed UnimplementedProductServiceServer
// for forward compatibility
type ProductServiceServer interface {
	// simple RPC (unary RPC)
	GetProduct(context.Context, *NoParam) (*Product, error)
	GetProductById(context.Context, *ProductIdRequest) (*Product, error)
	// server streaming RPC
	GetProductServerStreaming(*ProductIdList, ProductService_GetProductServerStreamingServer) error
	// client streaming RPC
	// rpc GetProductClientStreaming(stream ProductIdList) returns (ProductList);
	GetProductClientStreaming(ProductService_GetProductClientStreamingServer) error
	// bidirectional RPC
	GetProductBidirectionalStreaming(ProductService_GetProductBidirectionalStreamingServer) error
	mustEmbedUnimplementedProductServiceServer()
}

// UnimplementedProductServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (UnimplementedProductServiceServer) GetProduct(context.Context, *NoParam) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedProductServiceServer) GetProductById(context.Context, *ProductIdRequest) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductById not implemented")
}
func (UnimplementedProductServiceServer) GetProductServerStreaming(*ProductIdList, ProductService_GetProductServerStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method GetProductServerStreaming not implemented")
}
func (UnimplementedProductServiceServer) GetProductClientStreaming(ProductService_GetProductClientStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method GetProductClientStreaming not implemented")
}
func (UnimplementedProductServiceServer) GetProductBidirectionalStreaming(ProductService_GetProductBidirectionalStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method GetProductBidirectionalStreaming not implemented")
}
func (UnimplementedProductServiceServer) mustEmbedUnimplementedProductServiceServer() {}

// UnsafeProductServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServiceServer will
// result in compilation errors.
type UnsafeProductServiceServer interface {
	mustEmbedUnimplementedProductServiceServer()
}

func RegisterProductServiceServer(s grpc.ServiceRegistrar, srv ProductServiceServer) {
	s.RegisterService(&ProductService_ServiceDesc, srv)
}

func _ProductService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product_service.ProductService/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProduct(ctx, req.(*NoParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetProductById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProductById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product_service.ProductService/GetProductById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProductById(ctx, req.(*ProductIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetProductServerStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ProductIdList)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProductServiceServer).GetProductServerStreaming(m, &productServiceGetProductServerStreamingServer{stream})
}

type ProductService_GetProductServerStreamingServer interface {
	Send(*ProductIdResponse) error
	grpc.ServerStream
}

type productServiceGetProductServerStreamingServer struct {
	grpc.ServerStream
}

func (x *productServiceGetProductServerStreamingServer) Send(m *ProductIdResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ProductService_GetProductClientStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProductServiceServer).GetProductClientStreaming(&productServiceGetProductClientStreamingServer{stream})
}

type ProductService_GetProductClientStreamingServer interface {
	SendAndClose(*ProductIdList) error
	Recv() (*ProductIdRequest, error)
	grpc.ServerStream
}

type productServiceGetProductClientStreamingServer struct {
	grpc.ServerStream
}

func (x *productServiceGetProductClientStreamingServer) SendAndClose(m *ProductIdList) error {
	return x.ServerStream.SendMsg(m)
}

func (x *productServiceGetProductClientStreamingServer) Recv() (*ProductIdRequest, error) {
	m := new(ProductIdRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ProductService_GetProductBidirectionalStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProductServiceServer).GetProductBidirectionalStreaming(&productServiceGetProductBidirectionalStreamingServer{stream})
}

type ProductService_GetProductBidirectionalStreamingServer interface {
	Send(*ProductIdResponse) error
	Recv() (*ProductIdRequest, error)
	grpc.ServerStream
}

type productServiceGetProductBidirectionalStreamingServer struct {
	grpc.ServerStream
}

func (x *productServiceGetProductBidirectionalStreamingServer) Send(m *ProductIdResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *productServiceGetProductBidirectionalStreamingServer) Recv() (*ProductIdRequest, error) {
	m := new(ProductIdRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProductService_ServiceDesc is the grpc.ServiceDesc for ProductService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product_service.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProduct",
			Handler:    _ProductService_GetProduct_Handler,
		},
		{
			MethodName: "GetProductById",
			Handler:    _ProductService_GetProductById_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetProductServerStreaming",
			Handler:       _ProductService_GetProductServerStreaming_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetProductClientStreaming",
			Handler:       _ProductService_GetProductClientStreaming_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetProductBidirectionalStreaming",
			Handler:       _ProductService_GetProductBidirectionalStreaming_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/product.proto",
}
