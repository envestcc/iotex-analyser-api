// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

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

// HermesServiceClient is the client API for HermesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HermesServiceClient interface {
	// Hermes gives delegates who register the service of automatic reward distribution an overview of the reward distributions to their voters within a range of epochs
	Hermes(ctx context.Context, in *HermesRequest, opts ...grpc.CallOption) (*HermesResponse, error)
	//HermesByVoter returns Hermes voters' receiving history
	HermesByVoter(ctx context.Context, in *HermesByVoterRequest, opts ...grpc.CallOption) (*HermesByVoterResponse, error)
	// HermesByDelegate returns Hermes delegates' distribution history
	HermesByDelegate(ctx context.Context, in *HermesByDelegateRequest, opts ...grpc.CallOption) (*HermesByDelegateResponse, error)
	// HermesMeta provides Hermes platform metadata
	HermesMeta(ctx context.Context, in *HermesMetaRequest, opts ...grpc.CallOption) (*HermesMetaResponse, error)
	// HermesAverageStats returns the Hermes average statistics
	HermesAverageStats(ctx context.Context, in *HermesAverageStatsRequest, opts ...grpc.CallOption) (*HermesAverageStatsResponse, error)
}

type hermesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHermesServiceClient(cc grpc.ClientConnInterface) HermesServiceClient {
	return &hermesServiceClient{cc}
}

func (c *hermesServiceClient) Hermes(ctx context.Context, in *HermesRequest, opts ...grpc.CallOption) (*HermesResponse, error) {
	out := new(HermesResponse)
	err := c.cc.Invoke(ctx, "/api.HermesService/Hermes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hermesServiceClient) HermesByVoter(ctx context.Context, in *HermesByVoterRequest, opts ...grpc.CallOption) (*HermesByVoterResponse, error) {
	out := new(HermesByVoterResponse)
	err := c.cc.Invoke(ctx, "/api.HermesService/HermesByVoter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hermesServiceClient) HermesByDelegate(ctx context.Context, in *HermesByDelegateRequest, opts ...grpc.CallOption) (*HermesByDelegateResponse, error) {
	out := new(HermesByDelegateResponse)
	err := c.cc.Invoke(ctx, "/api.HermesService/HermesByDelegate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hermesServiceClient) HermesMeta(ctx context.Context, in *HermesMetaRequest, opts ...grpc.CallOption) (*HermesMetaResponse, error) {
	out := new(HermesMetaResponse)
	err := c.cc.Invoke(ctx, "/api.HermesService/HermesMeta", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hermesServiceClient) HermesAverageStats(ctx context.Context, in *HermesAverageStatsRequest, opts ...grpc.CallOption) (*HermesAverageStatsResponse, error) {
	out := new(HermesAverageStatsResponse)
	err := c.cc.Invoke(ctx, "/api.HermesService/HermesAverageStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HermesServiceServer is the server API for HermesService service.
// All implementations must embed UnimplementedHermesServiceServer
// for forward compatibility
type HermesServiceServer interface {
	// Hermes gives delegates who register the service of automatic reward distribution an overview of the reward distributions to their voters within a range of epochs
	Hermes(context.Context, *HermesRequest) (*HermesResponse, error)
	//HermesByVoter returns Hermes voters' receiving history
	HermesByVoter(context.Context, *HermesByVoterRequest) (*HermesByVoterResponse, error)
	// HermesByDelegate returns Hermes delegates' distribution history
	HermesByDelegate(context.Context, *HermesByDelegateRequest) (*HermesByDelegateResponse, error)
	// HermesMeta provides Hermes platform metadata
	HermesMeta(context.Context, *HermesMetaRequest) (*HermesMetaResponse, error)
	// HermesAverageStats returns the Hermes average statistics
	HermesAverageStats(context.Context, *HermesAverageStatsRequest) (*HermesAverageStatsResponse, error)
	mustEmbedUnimplementedHermesServiceServer()
}

// UnimplementedHermesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHermesServiceServer struct {
}

func (UnimplementedHermesServiceServer) Hermes(context.Context, *HermesRequest) (*HermesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hermes not implemented")
}
func (UnimplementedHermesServiceServer) HermesByVoter(context.Context, *HermesByVoterRequest) (*HermesByVoterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HermesByVoter not implemented")
}
func (UnimplementedHermesServiceServer) HermesByDelegate(context.Context, *HermesByDelegateRequest) (*HermesByDelegateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HermesByDelegate not implemented")
}
func (UnimplementedHermesServiceServer) HermesMeta(context.Context, *HermesMetaRequest) (*HermesMetaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HermesMeta not implemented")
}
func (UnimplementedHermesServiceServer) HermesAverageStats(context.Context, *HermesAverageStatsRequest) (*HermesAverageStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HermesAverageStats not implemented")
}
func (UnimplementedHermesServiceServer) mustEmbedUnimplementedHermesServiceServer() {}

// UnsafeHermesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HermesServiceServer will
// result in compilation errors.
type UnsafeHermesServiceServer interface {
	mustEmbedUnimplementedHermesServiceServer()
}

func RegisterHermesServiceServer(s grpc.ServiceRegistrar, srv HermesServiceServer) {
	s.RegisterService(&HermesService_ServiceDesc, srv)
}

func _HermesService_Hermes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HermesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HermesServiceServer).Hermes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.HermesService/Hermes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HermesServiceServer).Hermes(ctx, req.(*HermesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HermesService_HermesByVoter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HermesByVoterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HermesServiceServer).HermesByVoter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.HermesService/HermesByVoter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HermesServiceServer).HermesByVoter(ctx, req.(*HermesByVoterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HermesService_HermesByDelegate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HermesByDelegateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HermesServiceServer).HermesByDelegate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.HermesService/HermesByDelegate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HermesServiceServer).HermesByDelegate(ctx, req.(*HermesByDelegateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HermesService_HermesMeta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HermesMetaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HermesServiceServer).HermesMeta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.HermesService/HermesMeta",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HermesServiceServer).HermesMeta(ctx, req.(*HermesMetaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HermesService_HermesAverageStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HermesAverageStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HermesServiceServer).HermesAverageStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.HermesService/HermesAverageStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HermesServiceServer).HermesAverageStats(ctx, req.(*HermesAverageStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HermesService_ServiceDesc is the grpc.ServiceDesc for HermesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HermesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.HermesService",
	HandlerType: (*HermesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hermes",
			Handler:    _HermesService_Hermes_Handler,
		},
		{
			MethodName: "HermesByVoter",
			Handler:    _HermesService_HermesByVoter_Handler,
		},
		{
			MethodName: "HermesByDelegate",
			Handler:    _HermesService_HermesByDelegate_Handler,
		},
		{
			MethodName: "HermesMeta",
			Handler:    _HermesService_HermesMeta_Handler,
		},
		{
			MethodName: "HermesAverageStats",
			Handler:    _HermesService_HermesAverageStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api_hermes.proto",
}
