// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc_reflection_v1alpha

import (
	grpc "github.com/fgiudici/grpc-go"
	codes "github.com/fgiudici/grpc-go/codes"
	status "github.com/fgiudici/grpc-go/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ServerReflectionService is the service API for ServerReflection service.
// Fields should be assigned to their respective handler implementations only before
// RegisterServerReflectionService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type ServerReflectionService struct {
	// The reflection service is structured as a bidirectional stream, ensuring
	// all related requests go to a single server.
	ServerReflectionInfo func(ServerReflection_ServerReflectionInfoServer) error
}

func (s *ServerReflectionService) serverReflectionInfo(_ interface{}, stream grpc.ServerStream) error {
	return s.ServerReflectionInfo(&serverReflectionServerReflectionInfoServer{stream})
}

// RegisterServerReflectionService registers a service implementation with a gRPC server.
func RegisterServerReflectionService(s grpc.ServiceRegistrar, srv *ServerReflectionService) {
	srvCopy := *srv
	if srvCopy.ServerReflectionInfo == nil {
		srvCopy.ServerReflectionInfo = func(ServerReflection_ServerReflectionInfoServer) error {
			return status.Errorf(codes.Unimplemented, "method ServerReflectionInfo not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "grpc.reflection.v1alpha.ServerReflection",
		Methods:     []grpc.MethodDesc{},
		Streams: []grpc.StreamDesc{
			{
				StreamName:    "ServerReflectionInfo",
				Handler:       srvCopy.serverReflectionInfo,
				ServerStreams: true,
				ClientStreams: true,
			},
		},
		Metadata: "reflection/grpc_reflection_v1alpha/reflection.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewServerReflectionService creates a new ServerReflectionService containing the
// implemented methods of the ServerReflection service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewServerReflectionService(s interface{}) *ServerReflectionService {
	ns := &ServerReflectionService{}
	if h, ok := s.(interface {
		ServerReflectionInfo(ServerReflection_ServerReflectionInfoServer) error
	}); ok {
		ns.ServerReflectionInfo = h.ServerReflectionInfo
	}
	return ns
}

// UnstableServerReflectionService is the service API for ServerReflection service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstableServerReflectionService interface {
	// The reflection service is structured as a bidirectional stream, ensuring
	// all related requests go to a single server.
	ServerReflectionInfo(ServerReflection_ServerReflectionInfoServer) error
}
