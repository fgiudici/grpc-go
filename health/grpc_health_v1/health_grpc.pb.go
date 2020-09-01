// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc_health_v1

import (
	context "context"
	grpc "github.com/fgiudici/grpc-go"
	codes "github.com/fgiudici/grpc-go/codes"
	status "github.com/fgiudici/grpc-go/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// HealthService is the service API for Health service.
// Fields should be assigned to their respective handler implementations only before
// RegisterHealthService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type HealthService struct {
	// If the requested service is unknown, the call will fail with status
	// NOT_FOUND.
	Check func(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	// Performs a watch for the serving status of the requested service.
	// The server will immediately send back a message indicating the current
	// serving status.  It will then subsequently send a new message whenever
	// the service's serving status changes.
	//
	// If the requested service is unknown when the call is received, the
	// server will send a message setting the serving status to
	// SERVICE_UNKNOWN but will *not* terminate the call.  If at some
	// future point, the serving status of the service becomes known, the
	// server will send a new message with the service's serving status.
	//
	// If the call terminates with status UNIMPLEMENTED, then clients
	// should assume this method is not supported and should not retry the
	// call.  If the call terminates with any other status (including OK),
	// clients should retry the call with appropriate exponential backoff.
	Watch func(*HealthCheckRequest, Health_WatchServer) error
}

func (s *HealthService) check(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/grpc.health.v1.Health/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Check(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *HealthService) watch(_ interface{}, stream grpc.ServerStream) error {
	m := new(HealthCheckRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return s.Watch(m, &healthWatchServer{stream})
}

// RegisterHealthService registers a service implementation with a gRPC server.
func RegisterHealthService(s grpc.ServiceRegistrar, srv *HealthService) {
	srvCopy := *srv
	if srvCopy.Check == nil {
		srvCopy.Check = func(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
		}
	}
	if srvCopy.Watch == nil {
		srvCopy.Watch = func(*HealthCheckRequest, Health_WatchServer) error {
			return status.Errorf(codes.Unimplemented, "method Watch not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "grpc.health.v1.Health",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "Check",
				Handler:    srvCopy.check,
			},
		},
		Streams: []grpc.StreamDesc{
			{
				StreamName:    "Watch",
				Handler:       srvCopy.watch,
				ServerStreams: true,
			},
		},
		Metadata: "grpc/health/v1/health.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewHealthService creates a new HealthService containing the
// implemented methods of the Health service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewHealthService(s interface{}) *HealthService {
	ns := &HealthService{}
	if h, ok := s.(interface {
		Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	}); ok {
		ns.Check = h.Check
	}
	if h, ok := s.(interface {
		Watch(*HealthCheckRequest, Health_WatchServer) error
	}); ok {
		ns.Watch = h.Watch
	}
	return ns
}

// UnstableHealthService is the service API for Health service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstableHealthService interface {
	// If the requested service is unknown, the call will fail with status
	// NOT_FOUND.
	Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	// Performs a watch for the serving status of the requested service.
	// The server will immediately send back a message indicating the current
	// serving status.  It will then subsequently send a new message whenever
	// the service's serving status changes.
	//
	// If the requested service is unknown when the call is received, the
	// server will send a message setting the serving status to
	// SERVICE_UNKNOWN but will *not* terminate the call.  If at some
	// future point, the serving status of the service becomes known, the
	// server will send a new message with the service's serving status.
	//
	// If the call terminates with status UNIMPLEMENTED, then clients
	// should assume this method is not supported and should not retry the
	// call.  If the call terminates with any other status (including OK),
	// clients should retry the call with appropriate exponential backoff.
	Watch(*HealthCheckRequest, Health_WatchServer) error
}
