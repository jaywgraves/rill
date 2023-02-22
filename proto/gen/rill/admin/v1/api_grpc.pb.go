// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: rill/admin/v1/api.proto

package adminv1

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

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminServiceClient interface {
	// Ping returns information about the admin
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	// findOrganizations lists all the organizations currently managed by the admin
	FindOrganizations(ctx context.Context, in *FindOrganizationsRequest, opts ...grpc.CallOption) (*FindOrganizationsResponse, error)
	// FindOrganization returns information about a specific organization
	FindOrganization(ctx context.Context, in *FindOrganizationRequest, opts ...grpc.CallOption) (*FindOrganizationResponse, error)
	// CreateOrganization creates a new organization
	CreateOrganization(ctx context.Context, in *CreateOrganizationRequest, opts ...grpc.CallOption) (*CreateOrganizationResponse, error)
	// DeleteOrganization deletes an organizations
	DeleteOrganization(ctx context.Context, in *DeleteOrganizationRequest, opts ...grpc.CallOption) (*DeleteOrganizationResponse, error)
	// UpdateOrganization deletes an organizations
	UpdateOrganization(ctx context.Context, in *UpdateOrganizationRequest, opts ...grpc.CallOption) (*UpdateOrganizationResponse, error)
	// FindProjects lists all the projects currently available for given organizations
	FindProjects(ctx context.Context, in *FindProjectsRequest, opts ...grpc.CallOption) (*FindProjectsResponse, error)
	// FindProject returns information about a specific project
	FindProject(ctx context.Context, in *FindProjectRequest, opts ...grpc.CallOption) (*FindProjectResponse, error)
	// CreateProject creates a new project
	CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*CreateProjectResponse, error)
	// DeleteProject deletes an project
	DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*DeleteProjectResponse, error)
	// UpdateProject update a project
	UpdateProject(ctx context.Context, in *UpdateProjectRequest, opts ...grpc.CallOption) (*UpdateProjectResponse, error)
}

type adminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminServiceClient(cc grpc.ClientConnInterface) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/rill.admin.v1.AdminService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) FindOrganizations(ctx context.Context, in *FindOrganizationsRequest, opts ...grpc.CallOption) (*FindOrganizationsResponse, error) {
	out := new(FindOrganizationsResponse)
	err := c.cc.Invoke(ctx, "/rill.admin.v1.AdminService/FindOrganizations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) FindOrganization(ctx context.Context, in *FindOrganizationRequest, opts ...grpc.CallOption) (*FindOrganizationResponse, error) {
	out := new(FindOrganizationResponse)
	err := c.cc.Invoke(ctx, "/rill.admin.v1.AdminService/FindOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) CreateOrganization(ctx context.Context, in *CreateOrganizationRequest, opts ...grpc.CallOption) (*CreateOrganizationResponse, error) {
	out := new(CreateOrganizationResponse)
	err := c.cc.Invoke(ctx, "/rill.admin.v1.AdminService/CreateOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) DeleteOrganization(ctx context.Context, in *DeleteOrganizationRequest, opts ...grpc.CallOption) (*DeleteOrganizationResponse, error) {
	out := new(DeleteOrganizationResponse)
	err := c.cc.Invoke(ctx, "/rill.admin.v1.AdminService/DeleteOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) UpdateOrganization(ctx context.Context, in *UpdateOrganizationRequest, opts ...grpc.CallOption) (*UpdateOrganizationResponse, error) {
	out := new(UpdateOrganizationResponse)
	err := c.cc.Invoke(ctx, "/rill.admin.v1.AdminService/UpdateOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) FindProjects(ctx context.Context, in *FindProjectsRequest, opts ...grpc.CallOption) (*FindProjectsResponse, error) {
	out := new(FindProjectsResponse)
	err := c.cc.Invoke(ctx, "/rill.admin.v1.AdminService/FindProjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) FindProject(ctx context.Context, in *FindProjectRequest, opts ...grpc.CallOption) (*FindProjectResponse, error) {
	out := new(FindProjectResponse)
	err := c.cc.Invoke(ctx, "/rill.admin.v1.AdminService/FindProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*CreateProjectResponse, error) {
	out := new(CreateProjectResponse)
	err := c.cc.Invoke(ctx, "/rill.admin.v1.AdminService/CreateProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*DeleteProjectResponse, error) {
	out := new(DeleteProjectResponse)
	err := c.cc.Invoke(ctx, "/rill.admin.v1.AdminService/DeleteProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) UpdateProject(ctx context.Context, in *UpdateProjectRequest, opts ...grpc.CallOption) (*UpdateProjectResponse, error) {
	out := new(UpdateProjectResponse)
	err := c.cc.Invoke(ctx, "/rill.admin.v1.AdminService/UpdateProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
// All implementations must embed UnimplementedAdminServiceServer
// for forward compatibility
type AdminServiceServer interface {
	// Ping returns information about the admin
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	// findOrganizations lists all the organizations currently managed by the admin
	FindOrganizations(context.Context, *FindOrganizationsRequest) (*FindOrganizationsResponse, error)
	// FindOrganization returns information about a specific organization
	FindOrganization(context.Context, *FindOrganizationRequest) (*FindOrganizationResponse, error)
	// CreateOrganization creates a new organization
	CreateOrganization(context.Context, *CreateOrganizationRequest) (*CreateOrganizationResponse, error)
	// DeleteOrganization deletes an organizations
	DeleteOrganization(context.Context, *DeleteOrganizationRequest) (*DeleteOrganizationResponse, error)
	// UpdateOrganization deletes an organizations
	UpdateOrganization(context.Context, *UpdateOrganizationRequest) (*UpdateOrganizationResponse, error)
	// FindProjects lists all the projects currently available for given organizations
	FindProjects(context.Context, *FindProjectsRequest) (*FindProjectsResponse, error)
	// FindProject returns information about a specific project
	FindProject(context.Context, *FindProjectRequest) (*FindProjectResponse, error)
	// CreateProject creates a new project
	CreateProject(context.Context, *CreateProjectRequest) (*CreateProjectResponse, error)
	// DeleteProject deletes an project
	DeleteProject(context.Context, *DeleteProjectRequest) (*DeleteProjectResponse, error)
	// UpdateProject update a project
	UpdateProject(context.Context, *UpdateProjectRequest) (*UpdateProjectResponse, error)
	mustEmbedUnimplementedAdminServiceServer()
}

// UnimplementedAdminServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServiceServer struct {
}

func (UnimplementedAdminServiceServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedAdminServiceServer) FindOrganizations(context.Context, *FindOrganizationsRequest) (*FindOrganizationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOrganizations not implemented")
}
func (UnimplementedAdminServiceServer) FindOrganization(context.Context, *FindOrganizationRequest) (*FindOrganizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOrganization not implemented")
}
func (UnimplementedAdminServiceServer) CreateOrganization(context.Context, *CreateOrganizationRequest) (*CreateOrganizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrganization not implemented")
}
func (UnimplementedAdminServiceServer) DeleteOrganization(context.Context, *DeleteOrganizationRequest) (*DeleteOrganizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrganization not implemented")
}
func (UnimplementedAdminServiceServer) UpdateOrganization(context.Context, *UpdateOrganizationRequest) (*UpdateOrganizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrganization not implemented")
}
func (UnimplementedAdminServiceServer) FindProjects(context.Context, *FindProjectsRequest) (*FindProjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindProjects not implemented")
}
func (UnimplementedAdminServiceServer) FindProject(context.Context, *FindProjectRequest) (*FindProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindProject not implemented")
}
func (UnimplementedAdminServiceServer) CreateProject(context.Context, *CreateProjectRequest) (*CreateProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProject not implemented")
}
func (UnimplementedAdminServiceServer) DeleteProject(context.Context, *DeleteProjectRequest) (*DeleteProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProject not implemented")
}
func (UnimplementedAdminServiceServer) UpdateProject(context.Context, *UpdateProjectRequest) (*UpdateProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProject not implemented")
}
func (UnimplementedAdminServiceServer) mustEmbedUnimplementedAdminServiceServer() {}

// UnsafeAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServiceServer will
// result in compilation errors.
type UnsafeAdminServiceServer interface {
	mustEmbedUnimplementedAdminServiceServer()
}

func RegisterAdminServiceServer(s grpc.ServiceRegistrar, srv AdminServiceServer) {
	s.RegisterService(&AdminService_ServiceDesc, srv)
}

func _AdminService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rill.admin.v1.AdminService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_FindOrganizations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOrganizationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).FindOrganizations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rill.admin.v1.AdminService/FindOrganizations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).FindOrganizations(ctx, req.(*FindOrganizationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_FindOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).FindOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rill.admin.v1.AdminService/FindOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).FindOrganization(ctx, req.(*FindOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_CreateOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).CreateOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rill.admin.v1.AdminService/CreateOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).CreateOrganization(ctx, req.(*CreateOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_DeleteOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).DeleteOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rill.admin.v1.AdminService/DeleteOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).DeleteOrganization(ctx, req.(*DeleteOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_UpdateOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).UpdateOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rill.admin.v1.AdminService/UpdateOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).UpdateOrganization(ctx, req.(*UpdateOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_FindProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindProjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).FindProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rill.admin.v1.AdminService/FindProjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).FindProjects(ctx, req.(*FindProjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_FindProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).FindProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rill.admin.v1.AdminService/FindProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).FindProject(ctx, req.(*FindProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rill.admin.v1.AdminService/CreateProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).CreateProject(ctx, req.(*CreateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_DeleteProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).DeleteProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rill.admin.v1.AdminService/DeleteProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).DeleteProject(ctx, req.(*DeleteProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_UpdateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).UpdateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rill.admin.v1.AdminService/UpdateProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).UpdateProject(ctx, req.(*UpdateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminService_ServiceDesc is the grpc.ServiceDesc for AdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rill.admin.v1.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _AdminService_Ping_Handler,
		},
		{
			MethodName: "FindOrganizations",
			Handler:    _AdminService_FindOrganizations_Handler,
		},
		{
			MethodName: "FindOrganization",
			Handler:    _AdminService_FindOrganization_Handler,
		},
		{
			MethodName: "CreateOrganization",
			Handler:    _AdminService_CreateOrganization_Handler,
		},
		{
			MethodName: "DeleteOrganization",
			Handler:    _AdminService_DeleteOrganization_Handler,
		},
		{
			MethodName: "UpdateOrganization",
			Handler:    _AdminService_UpdateOrganization_Handler,
		},
		{
			MethodName: "FindProjects",
			Handler:    _AdminService_FindProjects_Handler,
		},
		{
			MethodName: "FindProject",
			Handler:    _AdminService_FindProject_Handler,
		},
		{
			MethodName: "CreateProject",
			Handler:    _AdminService_CreateProject_Handler,
		},
		{
			MethodName: "DeleteProject",
			Handler:    _AdminService_DeleteProject_Handler,
		},
		{
			MethodName: "UpdateProject",
			Handler:    _AdminService_UpdateProject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rill/admin/v1/api.proto",
}
