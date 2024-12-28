// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: pkg/grpcgen/dhcp-manager.proto

package grpcgen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	DHCPManager_GetLeases_FullMethodName = "/dhcp_manager.DHCPManager/GetLeases"
)

// DHCPManagerClient is the client API for DHCPManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DHCPManagerClient interface {
	GetLeases(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetLeasesResponse, error)
}

type dHCPManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewDHCPManagerClient(cc grpc.ClientConnInterface) DHCPManagerClient {
	return &dHCPManagerClient{cc}
}

func (c *dHCPManagerClient) GetLeases(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetLeasesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetLeasesResponse)
	err := c.cc.Invoke(ctx, DHCPManager_GetLeases_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DHCPManagerServer is the server API for DHCPManager service.
// All implementations must embed UnimplementedDHCPManagerServer
// for forward compatibility.
type DHCPManagerServer interface {
	GetLeases(context.Context, *emptypb.Empty) (*GetLeasesResponse, error)
	mustEmbedUnimplementedDHCPManagerServer()
}

// UnimplementedDHCPManagerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDHCPManagerServer struct{}

func (UnimplementedDHCPManagerServer) GetLeases(context.Context, *emptypb.Empty) (*GetLeasesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLeases not implemented")
}
func (UnimplementedDHCPManagerServer) mustEmbedUnimplementedDHCPManagerServer() {}
func (UnimplementedDHCPManagerServer) testEmbeddedByValue()                     {}

// UnsafeDHCPManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DHCPManagerServer will
// result in compilation errors.
type UnsafeDHCPManagerServer interface {
	mustEmbedUnimplementedDHCPManagerServer()
}

func RegisterDHCPManagerServer(s grpc.ServiceRegistrar, srv DHCPManagerServer) {
	// If the following call pancis, it indicates UnimplementedDHCPManagerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DHCPManager_ServiceDesc, srv)
}

func _DHCPManager_GetLeases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DHCPManagerServer).GetLeases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DHCPManager_GetLeases_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DHCPManagerServer).GetLeases(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// DHCPManager_ServiceDesc is the grpc.ServiceDesc for DHCPManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DHCPManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dhcp_manager.DHCPManager",
	HandlerType: (*DHCPManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLeases",
			Handler:    _DHCPManager_GetLeases_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/grpcgen/dhcp-manager.proto",
}