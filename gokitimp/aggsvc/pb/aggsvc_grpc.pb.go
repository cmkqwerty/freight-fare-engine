// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: gokitimp/aggsvc/pb/aggsvc.proto

package pb

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
	Aggregator_Aggregate_FullMethodName  = "/pb.Aggregator/Aggregate"
	Aggregator_GetInvoice_FullMethodName = "/pb.Aggregator/GetInvoice"
)

// AggregatorClient is the client API for Aggregator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AggregatorClient interface {
	Aggregate(ctx context.Context, in *AggregateRequest, opts ...grpc.CallOption) (*AggregateReply, error)
	GetInvoice(ctx context.Context, in *GetInvoiceRequest, opts ...grpc.CallOption) (*GetInvoiceReply, error)
}

type aggregatorClient struct {
	cc grpc.ClientConnInterface
}

func NewAggregatorClient(cc grpc.ClientConnInterface) AggregatorClient {
	return &aggregatorClient{cc}
}

func (c *aggregatorClient) Aggregate(ctx context.Context, in *AggregateRequest, opts ...grpc.CallOption) (*AggregateReply, error) {
	out := new(AggregateReply)
	err := c.cc.Invoke(ctx, Aggregator_Aggregate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorClient) GetInvoice(ctx context.Context, in *GetInvoiceRequest, opts ...grpc.CallOption) (*GetInvoiceReply, error) {
	out := new(GetInvoiceReply)
	err := c.cc.Invoke(ctx, Aggregator_GetInvoice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AggregatorServer is the server API for Aggregator service.
// All implementations must embed UnimplementedAggregatorServer
// for forward compatibility
type AggregatorServer interface {
	Aggregate(context.Context, *AggregateRequest) (*AggregateReply, error)
	GetInvoice(context.Context, *GetInvoiceRequest) (*GetInvoiceReply, error)
	mustEmbedUnimplementedAggregatorServer()
}

// UnimplementedAggregatorServer must be embedded to have forward compatible implementations.
type UnimplementedAggregatorServer struct {
}

func (UnimplementedAggregatorServer) Aggregate(context.Context, *AggregateRequest) (*AggregateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Aggregate not implemented")
}
func (UnimplementedAggregatorServer) GetInvoice(context.Context, *GetInvoiceRequest) (*GetInvoiceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInvoice not implemented")
}
func (UnimplementedAggregatorServer) mustEmbedUnimplementedAggregatorServer() {}

// UnsafeAggregatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AggregatorServer will
// result in compilation errors.
type UnsafeAggregatorServer interface {
	mustEmbedUnimplementedAggregatorServer()
}

func RegisterAggregatorServer(s grpc.ServiceRegistrar, srv AggregatorServer) {
	s.RegisterService(&Aggregator_ServiceDesc, srv)
}

func _Aggregator_Aggregate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AggregateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServer).Aggregate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Aggregator_Aggregate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServer).Aggregate(ctx, req.(*AggregateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Aggregator_GetInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInvoiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServer).GetInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Aggregator_GetInvoice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServer).GetInvoice(ctx, req.(*GetInvoiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Aggregator_ServiceDesc is the grpc.ServiceDesc for Aggregator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Aggregator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Aggregator",
	HandlerType: (*AggregatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Aggregate",
			Handler:    _Aggregator_Aggregate_Handler,
		},
		{
			MethodName: "GetInvoice",
			Handler:    _Aggregator_GetInvoice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gokitimp/aggsvc/pb/aggsvc.proto",
}
