// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package sample

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

// AddressServiceClient is the client API for AddressService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AddressServiceClient interface {
	Search(ctx context.Context, in *SampleRequest, opts ...grpc.CallOption) (*SampleResponse, error)
}

type addressServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAddressServiceClient(cc grpc.ClientConnInterface) AddressServiceClient {
	return &addressServiceClient{cc}
}

func (c *addressServiceClient) Search(ctx context.Context, in *SampleRequest, opts ...grpc.CallOption) (*SampleResponse, error) {
	out := new(SampleResponse)
	err := c.cc.Invoke(ctx, "/sample.AddressService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddressServiceServer is the server API for AddressService service.
// All implementations must embed UnimplementedAddressServiceServer
// for forward compatibility
type AddressServiceServer interface {
	Search(context.Context, *SampleRequest) (*SampleResponse, error)
	mustEmbedUnimplementedAddressServiceServer()
}

// UnimplementedAddressServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAddressServiceServer struct {
}

func (UnimplementedAddressServiceServer) Search(context.Context, *SampleRequest) (*SampleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedAddressServiceServer) mustEmbedUnimplementedAddressServiceServer() {}

// UnsafeAddressServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AddressServiceServer will
// result in compilation errors.
type UnsafeAddressServiceServer interface {
	mustEmbedUnimplementedAddressServiceServer()
}

func RegisterAddressServiceServer(s grpc.ServiceRegistrar, srv AddressServiceServer) {
	s.RegisterService(&AddressService_ServiceDesc, srv)
}

func _AddressService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SampleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sample.AddressService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).Search(ctx, req.(*SampleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AddressService_ServiceDesc is the grpc.ServiceDesc for AddressService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AddressService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sample.AddressService",
	HandlerType: (*AddressServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _AddressService_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sample/proto/sample.proto",
}
