// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: flow/v1/processor_service.proto

package flowv1

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
	ProcessorService_Transform_FullMethodName = "/flow.v1.ProcessorService/Transform"
)

// ProcessorServiceClient is the client API for ProcessorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProcessorServiceClient interface {
	// Flow sends Envelopes downstream; processor sends back transformed ones.
	Transform(ctx context.Context, opts ...grpc.CallOption) (ProcessorService_TransformClient, error)
}

type processorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProcessorServiceClient(cc grpc.ClientConnInterface) ProcessorServiceClient {
	return &processorServiceClient{cc}
}

func (c *processorServiceClient) Transform(ctx context.Context, opts ...grpc.CallOption) (ProcessorService_TransformClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProcessorService_ServiceDesc.Streams[0], ProcessorService_Transform_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &processorServiceTransformClient{stream}
	return x, nil
}

type ProcessorService_TransformClient interface {
	Send(*Envelope) error
	Recv() (*Envelope, error)
	grpc.ClientStream
}

type processorServiceTransformClient struct {
	grpc.ClientStream
}

func (x *processorServiceTransformClient) Send(m *Envelope) error {
	return x.ClientStream.SendMsg(m)
}

func (x *processorServiceTransformClient) Recv() (*Envelope, error) {
	m := new(Envelope)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProcessorServiceServer is the server API for ProcessorService service.
// All implementations must embed UnimplementedProcessorServiceServer
// for forward compatibility
type ProcessorServiceServer interface {
	// Flow sends Envelopes downstream; processor sends back transformed ones.
	Transform(ProcessorService_TransformServer) error
	mustEmbedUnimplementedProcessorServiceServer()
}

// UnimplementedProcessorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProcessorServiceServer struct {
}

func (UnimplementedProcessorServiceServer) Transform(ProcessorService_TransformServer) error {
	return status.Errorf(codes.Unimplemented, "method Transform not implemented")
}
func (UnimplementedProcessorServiceServer) mustEmbedUnimplementedProcessorServiceServer() {}

// UnsafeProcessorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProcessorServiceServer will
// result in compilation errors.
type UnsafeProcessorServiceServer interface {
	mustEmbedUnimplementedProcessorServiceServer()
}

func RegisterProcessorServiceServer(s grpc.ServiceRegistrar, srv ProcessorServiceServer) {
	s.RegisterService(&ProcessorService_ServiceDesc, srv)
}

func _ProcessorService_Transform_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProcessorServiceServer).Transform(&processorServiceTransformServer{stream})
}

type ProcessorService_TransformServer interface {
	Send(*Envelope) error
	Recv() (*Envelope, error)
	grpc.ServerStream
}

type processorServiceTransformServer struct {
	grpc.ServerStream
}

func (x *processorServiceTransformServer) Send(m *Envelope) error {
	return x.ServerStream.SendMsg(m)
}

func (x *processorServiceTransformServer) Recv() (*Envelope, error) {
	m := new(Envelope)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProcessorService_ServiceDesc is the grpc.ServiceDesc for ProcessorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProcessorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flow.v1.ProcessorService",
	HandlerType: (*ProcessorServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Transform",
			Handler:       _ProcessorService_Transform_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "flow/v1/processor_service.proto",
}
