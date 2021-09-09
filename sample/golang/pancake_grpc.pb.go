// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package golang

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

// BakePancakeServiceClient is the client API for BakePancakeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BakePancakeServiceClient interface {
	// bake pancake service on the specified menu
	Bake(ctx context.Context, in *BakeRequest, opts ...grpc.CallOption) (*BakeResponse, error)
	// Get the total number of panques for each menu
	Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportResponse, error)
}

type bakePancakeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBakePancakeServiceClient(cc grpc.ClientConnInterface) BakePancakeServiceClient {
	return &bakePancakeServiceClient{cc}
}

func (c *bakePancakeServiceClient) Bake(ctx context.Context, in *BakeRequest, opts ...grpc.CallOption) (*BakeResponse, error) {
	out := new(BakeResponse)
	err := c.cc.Invoke(ctx, "/pancake.baker.BakePancakeService/Bake", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bakePancakeServiceClient) Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportResponse, error) {
	out := new(ReportResponse)
	err := c.cc.Invoke(ctx, "/pancake.baker.BakePancakeService/Report", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BakePancakeServiceServer is the server API for BakePancakeService service.
// All implementations should embed UnimplementedBakePancakeServiceServer
// for forward compatibility
type BakePancakeServiceServer interface {
	// bake pancake service on the specified menu
	Bake(context.Context, *BakeRequest) (*BakeResponse, error)
	// Get the total number of panques for each menu
	Report(context.Context, *ReportRequest) (*ReportResponse, error)
}

// UnimplementedBakePancakeServiceServer should be embedded to have forward compatible implementations.
type UnimplementedBakePancakeServiceServer struct {
}

func (UnimplementedBakePancakeServiceServer) Bake(context.Context, *BakeRequest) (*BakeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bake not implemented")
}
func (UnimplementedBakePancakeServiceServer) Report(context.Context, *ReportRequest) (*ReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Report not implemented")
}

// UnsafeBakePancakeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BakePancakeServiceServer will
// result in compilation errors.
type UnsafeBakePancakeServiceServer interface {
	mustEmbedUnimplementedBakePancakeServiceServer()
}

func RegisterBakePancakeServiceServer(s grpc.ServiceRegistrar, srv BakePancakeServiceServer) {
	s.RegisterService(&BakePancakeService_ServiceDesc, srv)
}

func _BakePancakeService_Bake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BakeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BakePancakeServiceServer).Bake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pancake.baker.BakePancakeService/Bake",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BakePancakeServiceServer).Bake(ctx, req.(*BakeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BakePancakeService_Report_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BakePancakeServiceServer).Report(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pancake.baker.BakePancakeService/Report",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BakePancakeServiceServer).Report(ctx, req.(*ReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BakePancakeService_ServiceDesc is the grpc.ServiceDesc for BakePancakeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BakePancakeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pancake.baker.BakePancakeService",
	HandlerType: (*BakePancakeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Bake",
			Handler:    _BakePancakeService_Bake_Handler,
		},
		{
			MethodName: "Report",
			Handler:    _BakePancakeService_Report_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sample/proto/pancake.proto",
}
