// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "github.com/fgiudici/grpc-go"
	codes "github.com/fgiudici/grpc-go/codes"
	status "github.com/fgiudici/grpc-go/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ProfilingService is the service API for Profiling service.
// Fields should be assigned to their respective handler implementations only before
// RegisterProfilingService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type ProfilingService struct {
	// Enable allows users to toggle profiling on and off remotely.
	Enable func(context.Context, *EnableRequest) (*EnableResponse, error)
	// GetStreamStats is used to retrieve an array of stream-level stats from a
	// gRPC client/server.
	GetStreamStats func(context.Context, *GetStreamStatsRequest) (*GetStreamStatsResponse, error)
}

func (s *ProfilingService) enable(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Enable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/grpc.go.profiling.v1alpha.Profiling/Enable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Enable(ctx, req.(*EnableRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *ProfilingService) getStreamStats(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStreamStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.GetStreamStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/grpc.go.profiling.v1alpha.Profiling/GetStreamStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetStreamStats(ctx, req.(*GetStreamStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterProfilingService registers a service implementation with a gRPC server.
func RegisterProfilingService(s grpc.ServiceRegistrar, srv *ProfilingService) {
	srvCopy := *srv
	if srvCopy.Enable == nil {
		srvCopy.Enable = func(context.Context, *EnableRequest) (*EnableResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method Enable not implemented")
		}
	}
	if srvCopy.GetStreamStats == nil {
		srvCopy.GetStreamStats = func(context.Context, *GetStreamStatsRequest) (*GetStreamStatsResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method GetStreamStats not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "grpc.go.profiling.v1alpha.Profiling",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "Enable",
				Handler:    srvCopy.enable,
			},
			{
				MethodName: "GetStreamStats",
				Handler:    srvCopy.getStreamStats,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "profiling/proto/service.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewProfilingService creates a new ProfilingService containing the
// implemented methods of the Profiling service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewProfilingService(s interface{}) *ProfilingService {
	ns := &ProfilingService{}
	if h, ok := s.(interface {
		Enable(context.Context, *EnableRequest) (*EnableResponse, error)
	}); ok {
		ns.Enable = h.Enable
	}
	if h, ok := s.(interface {
		GetStreamStats(context.Context, *GetStreamStatsRequest) (*GetStreamStatsResponse, error)
	}); ok {
		ns.GetStreamStats = h.GetStreamStats
	}
	return ns
}

// UnstableProfilingService is the service API for Profiling service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstableProfilingService interface {
	// Enable allows users to toggle profiling on and off remotely.
	Enable(context.Context, *EnableRequest) (*EnableResponse, error)
	// GetStreamStats is used to retrieve an array of stream-level stats from a
	// gRPC client/server.
	GetStreamStats(context.Context, *GetStreamStatsRequest) (*GetStreamStatsResponse, error)
}
