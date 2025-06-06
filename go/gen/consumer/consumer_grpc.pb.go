// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: consumer/consumer.proto

package consumer

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
	ConsumerService_GetCapabilities_FullMethodName = "/consumer.ConsumerService/GetCapabilities"
	ConsumerService_Configure_FullMethodName       = "/consumer.ConsumerService/Configure"
	ConsumerService_Consume_FullMethodName         = "/consumer.ConsumerService/Consume"
	ConsumerService_BatchConsume_FullMethodName    = "/consumer.ConsumerService/BatchConsume"
	ConsumerService_StreamConsume_FullMethodName   = "/consumer.ConsumerService/StreamConsume"
	ConsumerService_CheckHealth_FullMethodName     = "/consumer.ConsumerService/CheckHealth"
	ConsumerService_GetMetrics_FullMethodName      = "/consumer.ConsumerService/GetMetrics"
	ConsumerService_Control_FullMethodName         = "/consumer.ConsumerService/Control"
)

// ConsumerServiceClient is the client API for ConsumerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConsumerServiceClient interface {
	// Basic operations
	GetCapabilities(ctx context.Context, in *CapabilitiesRequest, opts ...grpc.CallOption) (*CapabilitiesResponse, error)
	Configure(ctx context.Context, in *ConfigureRequest, opts ...grpc.CallOption) (*ConfigureResponse, error)
	// Single message consumption
	Consume(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (*ConsumeResponse, error)
	// Batch message consumption
	BatchConsume(ctx context.Context, in *BatchConsumeRequest, opts ...grpc.CallOption) (*BatchConsumeResponse, error)
	// Streaming message consumption
	StreamConsume(ctx context.Context, opts ...grpc.CallOption) (ConsumerService_StreamConsumeClient, error)
	// Health and metrics
	CheckHealth(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	GetMetrics(ctx context.Context, in *MetricsRequest, opts ...grpc.CallOption) (*MetricsResponse, error)
	// Lifecycle management
	Control(ctx context.Context, in *LifecycleRequest, opts ...grpc.CallOption) (*LifecycleResponse, error)
}

type consumerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConsumerServiceClient(cc grpc.ClientConnInterface) ConsumerServiceClient {
	return &consumerServiceClient{cc}
}

func (c *consumerServiceClient) GetCapabilities(ctx context.Context, in *CapabilitiesRequest, opts ...grpc.CallOption) (*CapabilitiesResponse, error) {
	out := new(CapabilitiesResponse)
	err := c.cc.Invoke(ctx, ConsumerService_GetCapabilities_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerServiceClient) Configure(ctx context.Context, in *ConfigureRequest, opts ...grpc.CallOption) (*ConfigureResponse, error) {
	out := new(ConfigureResponse)
	err := c.cc.Invoke(ctx, ConsumerService_Configure_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerServiceClient) Consume(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (*ConsumeResponse, error) {
	out := new(ConsumeResponse)
	err := c.cc.Invoke(ctx, ConsumerService_Consume_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerServiceClient) BatchConsume(ctx context.Context, in *BatchConsumeRequest, opts ...grpc.CallOption) (*BatchConsumeResponse, error) {
	out := new(BatchConsumeResponse)
	err := c.cc.Invoke(ctx, ConsumerService_BatchConsume_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerServiceClient) StreamConsume(ctx context.Context, opts ...grpc.CallOption) (ConsumerService_StreamConsumeClient, error) {
	stream, err := c.cc.NewStream(ctx, &ConsumerService_ServiceDesc.Streams[0], ConsumerService_StreamConsume_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &consumerServiceStreamConsumeClient{stream}
	return x, nil
}

type ConsumerService_StreamConsumeClient interface {
	Send(*ConsumeRequest) error
	Recv() (*ConsumeResponse, error)
	grpc.ClientStream
}

type consumerServiceStreamConsumeClient struct {
	grpc.ClientStream
}

func (x *consumerServiceStreamConsumeClient) Send(m *ConsumeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *consumerServiceStreamConsumeClient) Recv() (*ConsumeResponse, error) {
	m := new(ConsumeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *consumerServiceClient) CheckHealth(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, ConsumerService_CheckHealth_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerServiceClient) GetMetrics(ctx context.Context, in *MetricsRequest, opts ...grpc.CallOption) (*MetricsResponse, error) {
	out := new(MetricsResponse)
	err := c.cc.Invoke(ctx, ConsumerService_GetMetrics_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerServiceClient) Control(ctx context.Context, in *LifecycleRequest, opts ...grpc.CallOption) (*LifecycleResponse, error) {
	out := new(LifecycleResponse)
	err := c.cc.Invoke(ctx, ConsumerService_Control_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConsumerServiceServer is the server API for ConsumerService service.
// All implementations must embed UnimplementedConsumerServiceServer
// for forward compatibility
type ConsumerServiceServer interface {
	// Basic operations
	GetCapabilities(context.Context, *CapabilitiesRequest) (*CapabilitiesResponse, error)
	Configure(context.Context, *ConfigureRequest) (*ConfigureResponse, error)
	// Single message consumption
	Consume(context.Context, *ConsumeRequest) (*ConsumeResponse, error)
	// Batch message consumption
	BatchConsume(context.Context, *BatchConsumeRequest) (*BatchConsumeResponse, error)
	// Streaming message consumption
	StreamConsume(ConsumerService_StreamConsumeServer) error
	// Health and metrics
	CheckHealth(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	GetMetrics(context.Context, *MetricsRequest) (*MetricsResponse, error)
	// Lifecycle management
	Control(context.Context, *LifecycleRequest) (*LifecycleResponse, error)
	mustEmbedUnimplementedConsumerServiceServer()
}

// UnimplementedConsumerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConsumerServiceServer struct {
}

func (UnimplementedConsumerServiceServer) GetCapabilities(context.Context, *CapabilitiesRequest) (*CapabilitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCapabilities not implemented")
}
func (UnimplementedConsumerServiceServer) Configure(context.Context, *ConfigureRequest) (*ConfigureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Configure not implemented")
}
func (UnimplementedConsumerServiceServer) Consume(context.Context, *ConsumeRequest) (*ConsumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Consume not implemented")
}
func (UnimplementedConsumerServiceServer) BatchConsume(context.Context, *BatchConsumeRequest) (*BatchConsumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchConsume not implemented")
}
func (UnimplementedConsumerServiceServer) StreamConsume(ConsumerService_StreamConsumeServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamConsume not implemented")
}
func (UnimplementedConsumerServiceServer) CheckHealth(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckHealth not implemented")
}
func (UnimplementedConsumerServiceServer) GetMetrics(context.Context, *MetricsRequest) (*MetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetrics not implemented")
}
func (UnimplementedConsumerServiceServer) Control(context.Context, *LifecycleRequest) (*LifecycleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Control not implemented")
}
func (UnimplementedConsumerServiceServer) mustEmbedUnimplementedConsumerServiceServer() {}

// UnsafeConsumerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConsumerServiceServer will
// result in compilation errors.
type UnsafeConsumerServiceServer interface {
	mustEmbedUnimplementedConsumerServiceServer()
}

func RegisterConsumerServiceServer(s grpc.ServiceRegistrar, srv ConsumerServiceServer) {
	s.RegisterService(&ConsumerService_ServiceDesc, srv)
}

func _ConsumerService_GetCapabilities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CapabilitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServiceServer).GetCapabilities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConsumerService_GetCapabilities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServiceServer).GetCapabilities(ctx, req.(*CapabilitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsumerService_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServiceServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConsumerService_Configure_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServiceServer).Configure(ctx, req.(*ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsumerService_Consume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServiceServer).Consume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConsumerService_Consume_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServiceServer).Consume(ctx, req.(*ConsumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsumerService_BatchConsume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchConsumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServiceServer).BatchConsume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConsumerService_BatchConsume_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServiceServer).BatchConsume(ctx, req.(*BatchConsumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsumerService_StreamConsume_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ConsumerServiceServer).StreamConsume(&consumerServiceStreamConsumeServer{stream})
}

type ConsumerService_StreamConsumeServer interface {
	Send(*ConsumeResponse) error
	Recv() (*ConsumeRequest, error)
	grpc.ServerStream
}

type consumerServiceStreamConsumeServer struct {
	grpc.ServerStream
}

func (x *consumerServiceStreamConsumeServer) Send(m *ConsumeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *consumerServiceStreamConsumeServer) Recv() (*ConsumeRequest, error) {
	m := new(ConsumeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ConsumerService_CheckHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServiceServer).CheckHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConsumerService_CheckHealth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServiceServer).CheckHealth(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsumerService_GetMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServiceServer).GetMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConsumerService_GetMetrics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServiceServer).GetMetrics(ctx, req.(*MetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsumerService_Control_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LifecycleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServiceServer).Control(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConsumerService_Control_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServiceServer).Control(ctx, req.(*LifecycleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConsumerService_ServiceDesc is the grpc.ServiceDesc for ConsumerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConsumerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "consumer.ConsumerService",
	HandlerType: (*ConsumerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCapabilities",
			Handler:    _ConsumerService_GetCapabilities_Handler,
		},
		{
			MethodName: "Configure",
			Handler:    _ConsumerService_Configure_Handler,
		},
		{
			MethodName: "Consume",
			Handler:    _ConsumerService_Consume_Handler,
		},
		{
			MethodName: "BatchConsume",
			Handler:    _ConsumerService_BatchConsume_Handler,
		},
		{
			MethodName: "CheckHealth",
			Handler:    _ConsumerService_CheckHealth_Handler,
		},
		{
			MethodName: "GetMetrics",
			Handler:    _ConsumerService_GetMetrics_Handler,
		},
		{
			MethodName: "Control",
			Handler:    _ConsumerService_Control_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamConsume",
			Handler:       _ConsumerService_StreamConsume_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "consumer/consumer.proto",
}
