// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package routeguide

import (
	context "context"
	grpc "github.com/fgiudici/grpc-go"
	codes "github.com/fgiudici/grpc-go/codes"
	status "github.com/fgiudici/grpc-go/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// RouteGuideClient is the client API for RouteGuide service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/github.com/fgiudici/grpc-go/?tab=doc#ClientConn.NewStream.
type RouteGuideClient interface {
	// A simple RPC.
	//
	// Obtains the feature at a given position.
	//
	// A feature with an empty name is returned if there's no feature at the given
	// position.
	GetFeature(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Feature, error)
	// A server-to-client streaming RPC.
	//
	// Obtains the Features available within the given Rectangle.  Results are
	// streamed rather than returned at once (e.g. in a response message with a
	// repeated field), as the rectangle may cover a large area and contain a
	// huge number of features.
	ListFeatures(ctx context.Context, in *Rectangle, opts ...grpc.CallOption) (RouteGuide_ListFeaturesClient, error)
	// A client-to-server streaming RPC.
	//
	// Accepts a stream of Points on a route being traversed, returning a
	// RouteSummary when traversal is completed.
	RecordRoute(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RecordRouteClient, error)
	// A Bidirectional streaming RPC.
	//
	// Accepts a stream of RouteNotes sent while a route is being traversed,
	// while receiving other RouteNotes (e.g. from other users).
	RouteChat(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RouteChatClient, error)
}

type routeGuideClient struct {
	cc grpc.ClientConnInterface
}

func NewRouteGuideClient(cc grpc.ClientConnInterface) RouteGuideClient {
	return &routeGuideClient{cc}
}

var routeGuideGetFeatureStreamDesc = &grpc.StreamDesc{
	StreamName: "GetFeature",
}

func (c *routeGuideClient) GetFeature(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Feature, error) {
	out := new(Feature)
	err := c.cc.Invoke(ctx, "/routeguide.RouteGuide/GetFeature", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var routeGuideListFeaturesStreamDesc = &grpc.StreamDesc{
	StreamName:    "ListFeatures",
	ServerStreams: true,
}

func (c *routeGuideClient) ListFeatures(ctx context.Context, in *Rectangle, opts ...grpc.CallOption) (RouteGuide_ListFeaturesClient, error) {
	stream, err := c.cc.NewStream(ctx, routeGuideListFeaturesStreamDesc, "/routeguide.RouteGuide/ListFeatures", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideListFeaturesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RouteGuide_ListFeaturesClient interface {
	Recv() (*Feature, error)
	grpc.ClientStream
}

type routeGuideListFeaturesClient struct {
	grpc.ClientStream
}

func (x *routeGuideListFeaturesClient) Recv() (*Feature, error) {
	m := new(Feature)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var routeGuideRecordRouteStreamDesc = &grpc.StreamDesc{
	StreamName:    "RecordRoute",
	ClientStreams: true,
}

func (c *routeGuideClient) RecordRoute(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RecordRouteClient, error) {
	stream, err := c.cc.NewStream(ctx, routeGuideRecordRouteStreamDesc, "/routeguide.RouteGuide/RecordRoute", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideRecordRouteClient{stream}
	return x, nil
}

type RouteGuide_RecordRouteClient interface {
	Send(*Point) error
	CloseAndRecv() (*RouteSummary, error)
	grpc.ClientStream
}

type routeGuideRecordRouteClient struct {
	grpc.ClientStream
}

func (x *routeGuideRecordRouteClient) Send(m *Point) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeGuideRecordRouteClient) CloseAndRecv() (*RouteSummary, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(RouteSummary)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var routeGuideRouteChatStreamDesc = &grpc.StreamDesc{
	StreamName:    "RouteChat",
	ServerStreams: true,
	ClientStreams: true,
}

func (c *routeGuideClient) RouteChat(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RouteChatClient, error) {
	stream, err := c.cc.NewStream(ctx, routeGuideRouteChatStreamDesc, "/routeguide.RouteGuide/RouteChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideRouteChatClient{stream}
	return x, nil
}

type RouteGuide_RouteChatClient interface {
	Send(*RouteNote) error
	Recv() (*RouteNote, error)
	grpc.ClientStream
}

type routeGuideRouteChatClient struct {
	grpc.ClientStream
}

func (x *routeGuideRouteChatClient) Send(m *RouteNote) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeGuideRouteChatClient) Recv() (*RouteNote, error) {
	m := new(RouteNote)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RouteGuideService is the service API for RouteGuide service.
// Fields should be assigned to their respective handler implementations only before
// RegisterRouteGuideService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type RouteGuideService struct {
	// A simple RPC.
	//
	// Obtains the feature at a given position.
	//
	// A feature with an empty name is returned if there's no feature at the given
	// position.
	GetFeature func(context.Context, *Point) (*Feature, error)
	// A server-to-client streaming RPC.
	//
	// Obtains the Features available within the given Rectangle.  Results are
	// streamed rather than returned at once (e.g. in a response message with a
	// repeated field), as the rectangle may cover a large area and contain a
	// huge number of features.
	ListFeatures func(*Rectangle, RouteGuide_ListFeaturesServer) error
	// A client-to-server streaming RPC.
	//
	// Accepts a stream of Points on a route being traversed, returning a
	// RouteSummary when traversal is completed.
	RecordRoute func(RouteGuide_RecordRouteServer) error
	// A Bidirectional streaming RPC.
	//
	// Accepts a stream of RouteNotes sent while a route is being traversed,
	// while receiving other RouteNotes (e.g. from other users).
	RouteChat func(RouteGuide_RouteChatServer) error
}

func (s *RouteGuideService) getFeature(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Point)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.GetFeature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/routeguide.RouteGuide/GetFeature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetFeature(ctx, req.(*Point))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *RouteGuideService) listFeatures(_ interface{}, stream grpc.ServerStream) error {
	m := new(Rectangle)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return s.ListFeatures(m, &routeGuideListFeaturesServer{stream})
}
func (s *RouteGuideService) recordRoute(_ interface{}, stream grpc.ServerStream) error {
	return s.RecordRoute(&routeGuideRecordRouteServer{stream})
}
func (s *RouteGuideService) routeChat(_ interface{}, stream grpc.ServerStream) error {
	return s.RouteChat(&routeGuideRouteChatServer{stream})
}

type RouteGuide_ListFeaturesServer interface {
	Send(*Feature) error
	grpc.ServerStream
}

type routeGuideListFeaturesServer struct {
	grpc.ServerStream
}

func (x *routeGuideListFeaturesServer) Send(m *Feature) error {
	return x.ServerStream.SendMsg(m)
}

type RouteGuide_RecordRouteServer interface {
	SendAndClose(*RouteSummary) error
	Recv() (*Point, error)
	grpc.ServerStream
}

type routeGuideRecordRouteServer struct {
	grpc.ServerStream
}

func (x *routeGuideRecordRouteServer) SendAndClose(m *RouteSummary) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeGuideRecordRouteServer) Recv() (*Point, error) {
	m := new(Point)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

type RouteGuide_RouteChatServer interface {
	Send(*RouteNote) error
	Recv() (*RouteNote, error)
	grpc.ServerStream
}

type routeGuideRouteChatServer struct {
	grpc.ServerStream
}

func (x *routeGuideRouteChatServer) Send(m *RouteNote) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeGuideRouteChatServer) Recv() (*RouteNote, error) {
	m := new(RouteNote)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RegisterRouteGuideService registers a service implementation with a gRPC server.
func RegisterRouteGuideService(s grpc.ServiceRegistrar, srv *RouteGuideService) {
	srvCopy := *srv
	if srvCopy.GetFeature == nil {
		srvCopy.GetFeature = func(context.Context, *Point) (*Feature, error) {
			return nil, status.Errorf(codes.Unimplemented, "method GetFeature not implemented")
		}
	}
	if srvCopy.ListFeatures == nil {
		srvCopy.ListFeatures = func(*Rectangle, RouteGuide_ListFeaturesServer) error {
			return status.Errorf(codes.Unimplemented, "method ListFeatures not implemented")
		}
	}
	if srvCopy.RecordRoute == nil {
		srvCopy.RecordRoute = func(RouteGuide_RecordRouteServer) error {
			return status.Errorf(codes.Unimplemented, "method RecordRoute not implemented")
		}
	}
	if srvCopy.RouteChat == nil {
		srvCopy.RouteChat = func(RouteGuide_RouteChatServer) error {
			return status.Errorf(codes.Unimplemented, "method RouteChat not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "routeguide.RouteGuide",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "GetFeature",
				Handler:    srvCopy.getFeature,
			},
		},
		Streams: []grpc.StreamDesc{
			{
				StreamName:    "ListFeatures",
				Handler:       srvCopy.listFeatures,
				ServerStreams: true,
			},
			{
				StreamName:    "RecordRoute",
				Handler:       srvCopy.recordRoute,
				ClientStreams: true,
			},
			{
				StreamName:    "RouteChat",
				Handler:       srvCopy.routeChat,
				ServerStreams: true,
				ClientStreams: true,
			},
		},
		Metadata: "examples/route_guide/routeguide/route_guide.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewRouteGuideService creates a new RouteGuideService containing the
// implemented methods of the RouteGuide service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewRouteGuideService(s interface{}) *RouteGuideService {
	ns := &RouteGuideService{}
	if h, ok := s.(interface {
		GetFeature(context.Context, *Point) (*Feature, error)
	}); ok {
		ns.GetFeature = h.GetFeature
	}
	if h, ok := s.(interface {
		ListFeatures(*Rectangle, RouteGuide_ListFeaturesServer) error
	}); ok {
		ns.ListFeatures = h.ListFeatures
	}
	if h, ok := s.(interface {
		RecordRoute(RouteGuide_RecordRouteServer) error
	}); ok {
		ns.RecordRoute = h.RecordRoute
	}
	if h, ok := s.(interface {
		RouteChat(RouteGuide_RouteChatServer) error
	}); ok {
		ns.RouteChat = h.RouteChat
	}
	return ns
}

// UnstableRouteGuideService is the service API for RouteGuide service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstableRouteGuideService interface {
	// A simple RPC.
	//
	// Obtains the feature at a given position.
	//
	// A feature with an empty name is returned if there's no feature at the given
	// position.
	GetFeature(context.Context, *Point) (*Feature, error)
	// A server-to-client streaming RPC.
	//
	// Obtains the Features available within the given Rectangle.  Results are
	// streamed rather than returned at once (e.g. in a response message with a
	// repeated field), as the rectangle may cover a large area and contain a
	// huge number of features.
	ListFeatures(*Rectangle, RouteGuide_ListFeaturesServer) error
	// A client-to-server streaming RPC.
	//
	// Accepts a stream of Points on a route being traversed, returning a
	// RouteSummary when traversal is completed.
	RecordRoute(RouteGuide_RecordRouteServer) error
	// A Bidirectional streaming RPC.
	//
	// Accepts a stream of RouteNotes sent while a route is being traversed,
	// while receiving other RouteNotes (e.g. from other users).
	RouteChat(RouteGuide_RouteChatServer) error
}
