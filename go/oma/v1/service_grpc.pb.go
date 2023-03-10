// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: oma/v1/service.proto

package omav1

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

// OMAServiceClient is the client API for OMAService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OMAServiceClient interface {
	// Returns information about a single `Organization`.
	//
	// If the organization is not accessible by the authenticated principal or
	// does not exist, an empty `Organization` message is returned within the
	// `OrganizationGetResponse` message's `organization` field and the
	// `status` field will contain a 404 code.
	GetOrganization(ctx context.Context, in *GetOrganizationRequest, opts ...grpc.CallOption) (*GetOrganizationResponse, error)
	// Returns information about `Organizations` in the authenticated
	// principal's `any.cloud` `Account` that the principal has access to.
	ListOrganizations(ctx context.Context, in *ListOrganizationsRequest, opts ...grpc.CallOption) (OMAService_ListOrganizationsClient, error)
	// Creates a new `Organization`
	CreateOrganization(ctx context.Context, in *CreateOrganizationRequest, opts ...grpc.CallOption) (*CreateOrganizationResponse, error)
	// Deletes an `Organization`
	DeleteOrganization(ctx context.Context, in *DeleteOrganizationRequest, opts ...grpc.CallOption) (*DeleteOrganizationResponse, error)
}

type oMAServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOMAServiceClient(cc grpc.ClientConnInterface) OMAServiceClient {
	return &oMAServiceClient{cc}
}

func (c *oMAServiceClient) GetOrganization(ctx context.Context, in *GetOrganizationRequest, opts ...grpc.CallOption) (*GetOrganizationResponse, error) {
	out := new(GetOrganizationResponse)
	err := c.cc.Invoke(ctx, "/oma.v1.OMAService/GetOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oMAServiceClient) ListOrganizations(ctx context.Context, in *ListOrganizationsRequest, opts ...grpc.CallOption) (OMAService_ListOrganizationsClient, error) {
	stream, err := c.cc.NewStream(ctx, &OMAService_ServiceDesc.Streams[0], "/oma.v1.OMAService/ListOrganizations", opts...)
	if err != nil {
		return nil, err
	}
	x := &oMAServiceListOrganizationsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OMAService_ListOrganizationsClient interface {
	Recv() (*ListOrganizationsResponse, error)
	grpc.ClientStream
}

type oMAServiceListOrganizationsClient struct {
	grpc.ClientStream
}

func (x *oMAServiceListOrganizationsClient) Recv() (*ListOrganizationsResponse, error) {
	m := new(ListOrganizationsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *oMAServiceClient) CreateOrganization(ctx context.Context, in *CreateOrganizationRequest, opts ...grpc.CallOption) (*CreateOrganizationResponse, error) {
	out := new(CreateOrganizationResponse)
	err := c.cc.Invoke(ctx, "/oma.v1.OMAService/CreateOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oMAServiceClient) DeleteOrganization(ctx context.Context, in *DeleteOrganizationRequest, opts ...grpc.CallOption) (*DeleteOrganizationResponse, error) {
	out := new(DeleteOrganizationResponse)
	err := c.cc.Invoke(ctx, "/oma.v1.OMAService/DeleteOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OMAServiceServer is the server API for OMAService service.
// All implementations must embed UnimplementedOMAServiceServer
// for forward compatibility
type OMAServiceServer interface {
	// Returns information about a single `Organization`.
	//
	// If the organization is not accessible by the authenticated principal or
	// does not exist, an empty `Organization` message is returned within the
	// `OrganizationGetResponse` message's `organization` field and the
	// `status` field will contain a 404 code.
	GetOrganization(context.Context, *GetOrganizationRequest) (*GetOrganizationResponse, error)
	// Returns information about `Organizations` in the authenticated
	// principal's `any.cloud` `Account` that the principal has access to.
	ListOrganizations(*ListOrganizationsRequest, OMAService_ListOrganizationsServer) error
	// Creates a new `Organization`
	CreateOrganization(context.Context, *CreateOrganizationRequest) (*CreateOrganizationResponse, error)
	// Deletes an `Organization`
	DeleteOrganization(context.Context, *DeleteOrganizationRequest) (*DeleteOrganizationResponse, error)
	mustEmbedUnimplementedOMAServiceServer()
}

// UnimplementedOMAServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOMAServiceServer struct {
}

func (UnimplementedOMAServiceServer) GetOrganization(context.Context, *GetOrganizationRequest) (*GetOrganizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrganization not implemented")
}
func (UnimplementedOMAServiceServer) ListOrganizations(*ListOrganizationsRequest, OMAService_ListOrganizationsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListOrganizations not implemented")
}
func (UnimplementedOMAServiceServer) CreateOrganization(context.Context, *CreateOrganizationRequest) (*CreateOrganizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrganization not implemented")
}
func (UnimplementedOMAServiceServer) DeleteOrganization(context.Context, *DeleteOrganizationRequest) (*DeleteOrganizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrganization not implemented")
}
func (UnimplementedOMAServiceServer) mustEmbedUnimplementedOMAServiceServer() {}

// UnsafeOMAServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OMAServiceServer will
// result in compilation errors.
type UnsafeOMAServiceServer interface {
	mustEmbedUnimplementedOMAServiceServer()
}

func RegisterOMAServiceServer(s grpc.ServiceRegistrar, srv OMAServiceServer) {
	s.RegisterService(&OMAService_ServiceDesc, srv)
}

func _OMAService_GetOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OMAServiceServer).GetOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oma.v1.OMAService/GetOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OMAServiceServer).GetOrganization(ctx, req.(*GetOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OMAService_ListOrganizations_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListOrganizationsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OMAServiceServer).ListOrganizations(m, &oMAServiceListOrganizationsServer{stream})
}

type OMAService_ListOrganizationsServer interface {
	Send(*ListOrganizationsResponse) error
	grpc.ServerStream
}

type oMAServiceListOrganizationsServer struct {
	grpc.ServerStream
}

func (x *oMAServiceListOrganizationsServer) Send(m *ListOrganizationsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _OMAService_CreateOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OMAServiceServer).CreateOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oma.v1.OMAService/CreateOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OMAServiceServer).CreateOrganization(ctx, req.(*CreateOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OMAService_DeleteOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OMAServiceServer).DeleteOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oma.v1.OMAService/DeleteOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OMAServiceServer).DeleteOrganization(ctx, req.(*DeleteOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OMAService_ServiceDesc is the grpc.ServiceDesc for OMAService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OMAService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "oma.v1.OMAService",
	HandlerType: (*OMAServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrganization",
			Handler:    _OMAService_GetOrganization_Handler,
		},
		{
			MethodName: "CreateOrganization",
			Handler:    _OMAService_CreateOrganization_Handler,
		},
		{
			MethodName: "DeleteOrganization",
			Handler:    _OMAService_DeleteOrganization_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListOrganizations",
			Handler:       _OMAService_ListOrganizations_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "oma/v1/service.proto",
}
