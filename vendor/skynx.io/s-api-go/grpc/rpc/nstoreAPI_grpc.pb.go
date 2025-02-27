// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0--rc3
// source: skynx/protobuf/rpc/v1/nstoreAPI.proto

package rpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	hsecdb "skynx.io/s-api-go/grpc/resources/nstore/hsecdb"
	metricsdb "skynx.io/s-api-go/grpc/resources/nstore/metricsdb"
	netdb "skynx.io/s-api-go/grpc/resources/nstore/netdb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	NStoreAPI_GetNodeHostMetrics_FullMethodName        = "/api.NStoreAPI/GetNodeHostMetrics"
	NStoreAPI_GetNodeNetCtState_FullMethodName         = "/api.NStoreAPI/GetNodeNetCtState"
	NStoreAPI_GetNodeNetCtLog_FullMethodName           = "/api.NStoreAPI/GetNodeNetCtLog"
	NStoreAPI_GetNodeNetTrafficMetrics_FullMethodName  = "/api.NStoreAPI/GetNodeNetTrafficMetrics"
	NStoreAPI_GetNodeHostSecurityReport_FullMethodName = "/api.NStoreAPI/GetNodeHostSecurityReport"
)

// NStoreAPIClient is the client API for NStoreAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// [nstore-api] NStoreAPI Definition: nstore services
type NStoreAPIClient interface {
	// Gets node hostMetrics.
	GetNodeHostMetrics(ctx context.Context, in *metricsdb.HostMetricsRequest, opts ...grpc.CallOption) (*metricsdb.HostMetricsResponse, error)
	// Gets node netConntrackState.
	GetNodeNetCtState(ctx context.Context, in *netdb.ConntrackTableRequest, opts ...grpc.CallOption) (*netdb.ConntrackTableResponse, error)
	// Gets node netConntrackLog.
	GetNodeNetCtLog(ctx context.Context, in *netdb.ConntrackLogRequest, opts ...grpc.CallOption) (*netdb.ConntrackLogResponse, error)
	// Gets node netTrafficMetrics.
	GetNodeNetTrafficMetrics(ctx context.Context, in *netdb.TrafficMetricsRequest, opts ...grpc.CallOption) (*netdb.TrafficMetricsResponse, error)
	// Gets node hostSecurityReport.
	GetNodeHostSecurityReport(ctx context.Context, in *hsecdb.HostSecurityReportRequest, opts ...grpc.CallOption) (*hsecdb.HostSecurityReportResponse, error)
}

type nStoreAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewNStoreAPIClient(cc grpc.ClientConnInterface) NStoreAPIClient {
	return &nStoreAPIClient{cc}
}

func (c *nStoreAPIClient) GetNodeHostMetrics(ctx context.Context, in *metricsdb.HostMetricsRequest, opts ...grpc.CallOption) (*metricsdb.HostMetricsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(metricsdb.HostMetricsResponse)
	err := c.cc.Invoke(ctx, NStoreAPI_GetNodeHostMetrics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nStoreAPIClient) GetNodeNetCtState(ctx context.Context, in *netdb.ConntrackTableRequest, opts ...grpc.CallOption) (*netdb.ConntrackTableResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(netdb.ConntrackTableResponse)
	err := c.cc.Invoke(ctx, NStoreAPI_GetNodeNetCtState_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nStoreAPIClient) GetNodeNetCtLog(ctx context.Context, in *netdb.ConntrackLogRequest, opts ...grpc.CallOption) (*netdb.ConntrackLogResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(netdb.ConntrackLogResponse)
	err := c.cc.Invoke(ctx, NStoreAPI_GetNodeNetCtLog_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nStoreAPIClient) GetNodeNetTrafficMetrics(ctx context.Context, in *netdb.TrafficMetricsRequest, opts ...grpc.CallOption) (*netdb.TrafficMetricsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(netdb.TrafficMetricsResponse)
	err := c.cc.Invoke(ctx, NStoreAPI_GetNodeNetTrafficMetrics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nStoreAPIClient) GetNodeHostSecurityReport(ctx context.Context, in *hsecdb.HostSecurityReportRequest, opts ...grpc.CallOption) (*hsecdb.HostSecurityReportResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(hsecdb.HostSecurityReportResponse)
	err := c.cc.Invoke(ctx, NStoreAPI_GetNodeHostSecurityReport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NStoreAPIServer is the server API for NStoreAPI service.
// All implementations must embed UnimplementedNStoreAPIServer
// for forward compatibility
//
// [nstore-api] NStoreAPI Definition: nstore services
type NStoreAPIServer interface {
	// Gets node hostMetrics.
	GetNodeHostMetrics(context.Context, *metricsdb.HostMetricsRequest) (*metricsdb.HostMetricsResponse, error)
	// Gets node netConntrackState.
	GetNodeNetCtState(context.Context, *netdb.ConntrackTableRequest) (*netdb.ConntrackTableResponse, error)
	// Gets node netConntrackLog.
	GetNodeNetCtLog(context.Context, *netdb.ConntrackLogRequest) (*netdb.ConntrackLogResponse, error)
	// Gets node netTrafficMetrics.
	GetNodeNetTrafficMetrics(context.Context, *netdb.TrafficMetricsRequest) (*netdb.TrafficMetricsResponse, error)
	// Gets node hostSecurityReport.
	GetNodeHostSecurityReport(context.Context, *hsecdb.HostSecurityReportRequest) (*hsecdb.HostSecurityReportResponse, error)
	mustEmbedUnimplementedNStoreAPIServer()
}

// UnimplementedNStoreAPIServer must be embedded to have forward compatible implementations.
type UnimplementedNStoreAPIServer struct {
}

func (UnimplementedNStoreAPIServer) GetNodeHostMetrics(context.Context, *metricsdb.HostMetricsRequest) (*metricsdb.HostMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeHostMetrics not implemented")
}
func (UnimplementedNStoreAPIServer) GetNodeNetCtState(context.Context, *netdb.ConntrackTableRequest) (*netdb.ConntrackTableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeNetCtState not implemented")
}
func (UnimplementedNStoreAPIServer) GetNodeNetCtLog(context.Context, *netdb.ConntrackLogRequest) (*netdb.ConntrackLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeNetCtLog not implemented")
}
func (UnimplementedNStoreAPIServer) GetNodeNetTrafficMetrics(context.Context, *netdb.TrafficMetricsRequest) (*netdb.TrafficMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeNetTrafficMetrics not implemented")
}
func (UnimplementedNStoreAPIServer) GetNodeHostSecurityReport(context.Context, *hsecdb.HostSecurityReportRequest) (*hsecdb.HostSecurityReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeHostSecurityReport not implemented")
}
func (UnimplementedNStoreAPIServer) mustEmbedUnimplementedNStoreAPIServer() {}

// UnsafeNStoreAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NStoreAPIServer will
// result in compilation errors.
type UnsafeNStoreAPIServer interface {
	mustEmbedUnimplementedNStoreAPIServer()
}

func RegisterNStoreAPIServer(s grpc.ServiceRegistrar, srv NStoreAPIServer) {
	s.RegisterService(&NStoreAPI_ServiceDesc, srv)
}

func _NStoreAPI_GetNodeHostMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(metricsdb.HostMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NStoreAPIServer).GetNodeHostMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NStoreAPI_GetNodeHostMetrics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NStoreAPIServer).GetNodeHostMetrics(ctx, req.(*metricsdb.HostMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NStoreAPI_GetNodeNetCtState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(netdb.ConntrackTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NStoreAPIServer).GetNodeNetCtState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NStoreAPI_GetNodeNetCtState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NStoreAPIServer).GetNodeNetCtState(ctx, req.(*netdb.ConntrackTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NStoreAPI_GetNodeNetCtLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(netdb.ConntrackLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NStoreAPIServer).GetNodeNetCtLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NStoreAPI_GetNodeNetCtLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NStoreAPIServer).GetNodeNetCtLog(ctx, req.(*netdb.ConntrackLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NStoreAPI_GetNodeNetTrafficMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(netdb.TrafficMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NStoreAPIServer).GetNodeNetTrafficMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NStoreAPI_GetNodeNetTrafficMetrics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NStoreAPIServer).GetNodeNetTrafficMetrics(ctx, req.(*netdb.TrafficMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NStoreAPI_GetNodeHostSecurityReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(hsecdb.HostSecurityReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NStoreAPIServer).GetNodeHostSecurityReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NStoreAPI_GetNodeHostSecurityReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NStoreAPIServer).GetNodeHostSecurityReport(ctx, req.(*hsecdb.HostSecurityReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NStoreAPI_ServiceDesc is the grpc.ServiceDesc for NStoreAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NStoreAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.NStoreAPI",
	HandlerType: (*NStoreAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNodeHostMetrics",
			Handler:    _NStoreAPI_GetNodeHostMetrics_Handler,
		},
		{
			MethodName: "GetNodeNetCtState",
			Handler:    _NStoreAPI_GetNodeNetCtState_Handler,
		},
		{
			MethodName: "GetNodeNetCtLog",
			Handler:    _NStoreAPI_GetNodeNetCtLog_Handler,
		},
		{
			MethodName: "GetNodeNetTrafficMetrics",
			Handler:    _NStoreAPI_GetNodeNetTrafficMetrics_Handler,
		},
		{
			MethodName: "GetNodeHostSecurityReport",
			Handler:    _NStoreAPI_GetNodeHostSecurityReport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "skynx/protobuf/rpc/v1/nstoreAPI.proto",
}
