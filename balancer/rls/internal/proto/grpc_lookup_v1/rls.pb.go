// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/lookup/v1/rls.proto

package grpc_lookup_v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "github.com/fgiudici/grpc-go"
	codes "github.com/fgiudici/grpc-go/codes"
	status "github.com/fgiudici/grpc-go/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RouteLookupRequest struct {
	// Full host name of the target server, e.g. firestore.googleapis.com.
	// Only set for gRPC requests; HTTP requests must use key_map explicitly.
	Server string `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	// Full path of the request, i.e. "/service/method".
	// Only set for gRPC requests; HTTP requests must use key_map explicitly.
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	// Target type allows the client to specify what kind of target format it
	// would like from RLS to allow it to find the regional server, e.g. "grpc".
	TargetType string `protobuf:"bytes,3,opt,name=target_type,json=targetType,proto3" json:"target_type,omitempty"`
	// Map of key values extracted via key builders for the gRPC or HTTP request.
	KeyMap               map[string]string `protobuf:"bytes,4,rep,name=key_map,json=keyMap,proto3" json:"key_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RouteLookupRequest) Reset()         { *m = RouteLookupRequest{} }
func (m *RouteLookupRequest) String() string { return proto.CompactTextString(m) }
func (*RouteLookupRequest) ProtoMessage()    {}
func (*RouteLookupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bab962d3362f3ca, []int{0}
}

func (m *RouteLookupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteLookupRequest.Unmarshal(m, b)
}
func (m *RouteLookupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteLookupRequest.Marshal(b, m, deterministic)
}
func (m *RouteLookupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteLookupRequest.Merge(m, src)
}
func (m *RouteLookupRequest) XXX_Size() int {
	return xxx_messageInfo_RouteLookupRequest.Size(m)
}
func (m *RouteLookupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteLookupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RouteLookupRequest proto.InternalMessageInfo

func (m *RouteLookupRequest) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

func (m *RouteLookupRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *RouteLookupRequest) GetTargetType() string {
	if m != nil {
		return m.TargetType
	}
	return ""
}

func (m *RouteLookupRequest) GetKeyMap() map[string]string {
	if m != nil {
		return m.KeyMap
	}
	return nil
}

type RouteLookupResponse struct {
	// Prioritized list (best one first) of addressable entities to use
	// for routing, using syntax requested by the request target_type.
	// The targets will be tried in order until a healthy one is found.
	Targets []string `protobuf:"bytes,3,rep,name=targets,proto3" json:"targets,omitempty"`
	// Optional header value to pass along to AFE in the X-Google-RLS-Data header.
	// Cached with "target" and sent with all requests that match the request key.
	// Allows the RLS to pass its work product to the eventual target.
	HeaderData           string   `protobuf:"bytes,2,opt,name=header_data,json=headerData,proto3" json:"header_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteLookupResponse) Reset()         { *m = RouteLookupResponse{} }
func (m *RouteLookupResponse) String() string { return proto.CompactTextString(m) }
func (*RouteLookupResponse) ProtoMessage()    {}
func (*RouteLookupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bab962d3362f3ca, []int{1}
}

func (m *RouteLookupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteLookupResponse.Unmarshal(m, b)
}
func (m *RouteLookupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteLookupResponse.Marshal(b, m, deterministic)
}
func (m *RouteLookupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteLookupResponse.Merge(m, src)
}
func (m *RouteLookupResponse) XXX_Size() int {
	return xxx_messageInfo_RouteLookupResponse.Size(m)
}
func (m *RouteLookupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteLookupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RouteLookupResponse proto.InternalMessageInfo

func (m *RouteLookupResponse) GetTargets() []string {
	if m != nil {
		return m.Targets
	}
	return nil
}

func (m *RouteLookupResponse) GetHeaderData() string {
	if m != nil {
		return m.HeaderData
	}
	return ""
}

func init() {
	proto.RegisterType((*RouteLookupRequest)(nil), "grpc.lookup.v1.RouteLookupRequest")
	proto.RegisterMapType((map[string]string)(nil), "grpc.lookup.v1.RouteLookupRequest.KeyMapEntry")
	proto.RegisterType((*RouteLookupResponse)(nil), "grpc.lookup.v1.RouteLookupResponse")
}

func init() { proto.RegisterFile("grpc/lookup/v1/rls.proto", fileDescriptor_3bab962d3362f3ca) }

var fileDescriptor_3bab962d3362f3ca = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xd1, 0x4b, 0xeb, 0x30,
	0x14, 0xc6, 0x6f, 0xd7, 0xdd, 0x6e, 0x3b, 0x85, 0xcb, 0x6e, 0xee, 0x45, 0xca, 0x5e, 0x1c, 0xf5,
	0x65, 0x0f, 0x92, 0xb2, 0xf9, 0xa2, 0x3e, 0x0e, 0x45, 0x50, 0x07, 0xa3, 0xfa, 0x20, 0x3e, 0x58,
	0xe2, 0x76, 0xc8, 0x46, 0x6b, 0x13, 0x93, 0xb4, 0xd0, 0x3f, 0xd8, 0xff, 0x43, 0xda, 0x54, 0xd8,
	0x14, 0xf4, 0xed, 0xfb, 0xbe, 0x73, 0x48, 0x7e, 0x27, 0x39, 0x10, 0x70, 0x25, 0x57, 0x51, 0x26,
	0x44, 0x5a, 0xc8, 0xa8, 0x9c, 0x46, 0x2a, 0xd3, 0x54, 0x2a, 0x61, 0x04, 0xf9, 0x53, 0x57, 0xa8,
	0xad, 0xd0, 0x72, 0x1a, 0xbe, 0x39, 0x40, 0x62, 0x51, 0x18, 0xbc, 0x6d, 0xa2, 0x18, 0x5f, 0x0b,
	0xd4, 0x86, 0x1c, 0x80, 0xa7, 0x51, 0x95, 0xa8, 0x02, 0x67, 0xec, 0x4c, 0x06, 0x71, 0xeb, 0x08,
	0x81, 0xae, 0x64, 0x66, 0x13, 0x74, 0x9a, 0xb4, 0xd1, 0xe4, 0x10, 0x7c, 0xc3, 0x14, 0x47, 0x93,
	0x98, 0x4a, 0x62, 0xe0, 0x36, 0x25, 0xb0, 0xd1, 0x7d, 0x25, 0x91, 0x5c, 0x41, 0x2f, 0xc5, 0x2a,
	0x79, 0x61, 0x32, 0xe8, 0x8e, 0xdd, 0x89, 0x3f, 0xa3, 0x74, 0x9f, 0x82, 0x7e, 0x25, 0xa0, 0x37,
	0x58, 0x2d, 0x98, 0xbc, 0xcc, 0x8d, 0xaa, 0x62, 0x2f, 0x6d, 0xcc, 0xe8, 0x0c, 0xfc, 0x9d, 0x98,
	0x0c, 0xc1, 0x4d, 0xb1, 0x6a, 0x09, 0x6b, 0x49, 0xfe, 0xc3, 0xef, 0x92, 0x65, 0x05, 0xb6, 0x7c,
	0xd6, 0x9c, 0x77, 0x4e, 0x9d, 0xf0, 0x09, 0xfe, 0xed, 0x5d, 0xa2, 0xa5, 0xc8, 0x35, 0x92, 0x00,
	0x7a, 0x16, 0x54, 0x07, 0xee, 0xd8, 0x9d, 0x0c, 0xe2, 0x0f, 0x5b, 0x4f, 0xb5, 0x41, 0xb6, 0x46,
	0x95, 0xac, 0x99, 0x61, 0xed, 0x81, 0x60, 0xa3, 0x0b, 0x66, 0xd8, 0x75, 0xb7, 0xef, 0x0c, 0x3b,
	0xb1, 0x67, 0xfb, 0x67, 0xf9, 0xde, 0x33, 0xde, 0xa1, 0x2a, 0xb7, 0x2b, 0x24, 0x0f, 0xe0, 0xef,
	0xa4, 0x24, 0xfc, 0x79, 0xee, 0xd1, 0xd1, 0xb7, 0x3d, 0x16, 0x3b, 0xfc, 0x35, 0x5f, 0xc0, 0xdf,
	0xad, 0xf8, 0xd4, 0x3a, 0xef, 0xc7, 0x99, 0x5e, 0xd6, 0xdf, 0xbc, 0x74, 0x1e, 0x8f, 0xb9, 0x10,
	0x3c, 0x43, 0xca, 0x45, 0xc6, 0x72, 0x4e, 0x85, 0xe2, 0xd1, 0xee, 0x52, 0xd4, 0x3a, 0xb1, 0x3a,
	0x29, 0xa7, 0xcf, 0x5e, 0xb3, 0x1d, 0x27, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0xca, 0x8d, 0x5c,
	0xc7, 0x39, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RouteLookupServiceClient is the client API for RouteLookupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/github.com/fgiudici/grpc-go#ClientConn.NewStream.
type RouteLookupServiceClient interface {
	// Lookup returns a target for a single key.
	RouteLookup(ctx context.Context, in *RouteLookupRequest, opts ...grpc.CallOption) (*RouteLookupResponse, error)
}

type routeLookupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRouteLookupServiceClient(cc grpc.ClientConnInterface) RouteLookupServiceClient {
	return &routeLookupServiceClient{cc}
}

func (c *routeLookupServiceClient) RouteLookup(ctx context.Context, in *RouteLookupRequest, opts ...grpc.CallOption) (*RouteLookupResponse, error) {
	out := new(RouteLookupResponse)
	err := c.cc.Invoke(ctx, "/grpc.lookup.v1.RouteLookupService/RouteLookup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouteLookupServiceServer is the server API for RouteLookupService service.
type RouteLookupServiceServer interface {
	// Lookup returns a target for a single key.
	RouteLookup(context.Context, *RouteLookupRequest) (*RouteLookupResponse, error)
}

// UnimplementedRouteLookupServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRouteLookupServiceServer struct {
}

func (*UnimplementedRouteLookupServiceServer) RouteLookup(ctx context.Context, req *RouteLookupRequest) (*RouteLookupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RouteLookup not implemented")
}

func RegisterRouteLookupServiceServer(s *grpc.Server, srv RouteLookupServiceServer) {
	s.RegisterService(&_RouteLookupService_serviceDesc, srv)
}

func _RouteLookupService_RouteLookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteLookupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteLookupServiceServer).RouteLookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.lookup.v1.RouteLookupService/RouteLookup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteLookupServiceServer).RouteLookup(ctx, req.(*RouteLookupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RouteLookupService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.lookup.v1.RouteLookupService",
	HandlerType: (*RouteLookupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RouteLookup",
			Handler:    _RouteLookupService_RouteLookup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/lookup/v1/rls.proto",
}
