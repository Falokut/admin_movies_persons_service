// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: admin_movies_persons_service_v1.proto

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MoviesPersonsServiceV1Client is the client API for MoviesPersonsServiceV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MoviesPersonsServiceV1Client interface {
	GetPersons(ctx context.Context, in *GetPersonsRequest, opts ...grpc.CallOption) (*Persons, error)
	SearchPerson(ctx context.Context, in *SearchPersonRequest, opts ...grpc.CallOption) (*Persons, error)
	IsPersonWithIDExists(ctx context.Context, in *IsPersonWithIDExistsRequest, opts ...grpc.CallOption) (*IsPersonWithIDExistsResponse, error)
	IsPersonExists(ctx context.Context, in *IsPersonExistsRequest, opts ...grpc.CallOption) (*IsPersonExistsResponse, error)
	UpdatePersonFields(ctx context.Context, in *UpdatePersonFieldsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdatePerson(ctx context.Context, in *UpdatePersonRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreatePerson(ctx context.Context, in *CreatePersonRequest, opts ...grpc.CallOption) (*CreatePersonResponce, error)
	DeletePersons(ctx context.Context, in *DeletePersonsRequest, opts ...grpc.CallOption) (*DeletePersonsResponce, error)
}

type moviesPersonsServiceV1Client struct {
	cc grpc.ClientConnInterface
}

func NewMoviesPersonsServiceV1Client(cc grpc.ClientConnInterface) MoviesPersonsServiceV1Client {
	return &moviesPersonsServiceV1Client{cc}
}

func (c *moviesPersonsServiceV1Client) GetPersons(ctx context.Context, in *GetPersonsRequest, opts ...grpc.CallOption) (*Persons, error) {
	out := new(Persons)
	err := c.cc.Invoke(ctx, "/admin_movies_persons_service.moviesPersonsServiceV1/GetPersons", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moviesPersonsServiceV1Client) SearchPerson(ctx context.Context, in *SearchPersonRequest, opts ...grpc.CallOption) (*Persons, error) {
	out := new(Persons)
	err := c.cc.Invoke(ctx, "/admin_movies_persons_service.moviesPersonsServiceV1/SearchPerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moviesPersonsServiceV1Client) IsPersonWithIDExists(ctx context.Context, in *IsPersonWithIDExistsRequest, opts ...grpc.CallOption) (*IsPersonWithIDExistsResponse, error) {
	out := new(IsPersonWithIDExistsResponse)
	err := c.cc.Invoke(ctx, "/admin_movies_persons_service.moviesPersonsServiceV1/IsPersonWithIDExists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moviesPersonsServiceV1Client) IsPersonExists(ctx context.Context, in *IsPersonExistsRequest, opts ...grpc.CallOption) (*IsPersonExistsResponse, error) {
	out := new(IsPersonExistsResponse)
	err := c.cc.Invoke(ctx, "/admin_movies_persons_service.moviesPersonsServiceV1/IsPersonExists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moviesPersonsServiceV1Client) UpdatePersonFields(ctx context.Context, in *UpdatePersonFieldsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/admin_movies_persons_service.moviesPersonsServiceV1/UpdatePersonFields", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moviesPersonsServiceV1Client) UpdatePerson(ctx context.Context, in *UpdatePersonRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/admin_movies_persons_service.moviesPersonsServiceV1/UpdatePerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moviesPersonsServiceV1Client) CreatePerson(ctx context.Context, in *CreatePersonRequest, opts ...grpc.CallOption) (*CreatePersonResponce, error) {
	out := new(CreatePersonResponce)
	err := c.cc.Invoke(ctx, "/admin_movies_persons_service.moviesPersonsServiceV1/CreatePerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moviesPersonsServiceV1Client) DeletePersons(ctx context.Context, in *DeletePersonsRequest, opts ...grpc.CallOption) (*DeletePersonsResponce, error) {
	out := new(DeletePersonsResponce)
	err := c.cc.Invoke(ctx, "/admin_movies_persons_service.moviesPersonsServiceV1/DeletePersons", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MoviesPersonsServiceV1Server is the server API for MoviesPersonsServiceV1 service.
// All implementations must embed UnimplementedMoviesPersonsServiceV1Server
// for forward compatibility
type MoviesPersonsServiceV1Server interface {
	GetPersons(context.Context, *GetPersonsRequest) (*Persons, error)
	SearchPerson(context.Context, *SearchPersonRequest) (*Persons, error)
	IsPersonWithIDExists(context.Context, *IsPersonWithIDExistsRequest) (*IsPersonWithIDExistsResponse, error)
	IsPersonExists(context.Context, *IsPersonExistsRequest) (*IsPersonExistsResponse, error)
	UpdatePersonFields(context.Context, *UpdatePersonFieldsRequest) (*emptypb.Empty, error)
	UpdatePerson(context.Context, *UpdatePersonRequest) (*emptypb.Empty, error)
	CreatePerson(context.Context, *CreatePersonRequest) (*CreatePersonResponce, error)
	DeletePersons(context.Context, *DeletePersonsRequest) (*DeletePersonsResponce, error)
	mustEmbedUnimplementedMoviesPersonsServiceV1Server()
}

// UnimplementedMoviesPersonsServiceV1Server must be embedded to have forward compatible implementations.
type UnimplementedMoviesPersonsServiceV1Server struct {
}

func (UnimplementedMoviesPersonsServiceV1Server) GetPersons(context.Context, *GetPersonsRequest) (*Persons, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPersons not implemented")
}
func (UnimplementedMoviesPersonsServiceV1Server) SearchPerson(context.Context, *SearchPersonRequest) (*Persons, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchPerson not implemented")
}
func (UnimplementedMoviesPersonsServiceV1Server) IsPersonWithIDExists(context.Context, *IsPersonWithIDExistsRequest) (*IsPersonWithIDExistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsPersonWithIDExists not implemented")
}
func (UnimplementedMoviesPersonsServiceV1Server) IsPersonExists(context.Context, *IsPersonExistsRequest) (*IsPersonExistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsPersonExists not implemented")
}
func (UnimplementedMoviesPersonsServiceV1Server) UpdatePersonFields(context.Context, *UpdatePersonFieldsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePersonFields not implemented")
}
func (UnimplementedMoviesPersonsServiceV1Server) UpdatePerson(context.Context, *UpdatePersonRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePerson not implemented")
}
func (UnimplementedMoviesPersonsServiceV1Server) CreatePerson(context.Context, *CreatePersonRequest) (*CreatePersonResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePerson not implemented")
}
func (UnimplementedMoviesPersonsServiceV1Server) DeletePersons(context.Context, *DeletePersonsRequest) (*DeletePersonsResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePersons not implemented")
}
func (UnimplementedMoviesPersonsServiceV1Server) mustEmbedUnimplementedMoviesPersonsServiceV1Server() {
}

// UnsafeMoviesPersonsServiceV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MoviesPersonsServiceV1Server will
// result in compilation errors.
type UnsafeMoviesPersonsServiceV1Server interface {
	mustEmbedUnimplementedMoviesPersonsServiceV1Server()
}

func RegisterMoviesPersonsServiceV1Server(s grpc.ServiceRegistrar, srv MoviesPersonsServiceV1Server) {
	s.RegisterService(&MoviesPersonsServiceV1_ServiceDesc, srv)
}

func _MoviesPersonsServiceV1_GetPersons_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPersonsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoviesPersonsServiceV1Server).GetPersons(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_movies_persons_service.moviesPersonsServiceV1/GetPersons",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoviesPersonsServiceV1Server).GetPersons(ctx, req.(*GetPersonsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MoviesPersonsServiceV1_SearchPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchPersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoviesPersonsServiceV1Server).SearchPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_movies_persons_service.moviesPersonsServiceV1/SearchPerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoviesPersonsServiceV1Server).SearchPerson(ctx, req.(*SearchPersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MoviesPersonsServiceV1_IsPersonWithIDExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsPersonWithIDExistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoviesPersonsServiceV1Server).IsPersonWithIDExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_movies_persons_service.moviesPersonsServiceV1/IsPersonWithIDExists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoviesPersonsServiceV1Server).IsPersonWithIDExists(ctx, req.(*IsPersonWithIDExistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MoviesPersonsServiceV1_IsPersonExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsPersonExistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoviesPersonsServiceV1Server).IsPersonExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_movies_persons_service.moviesPersonsServiceV1/IsPersonExists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoviesPersonsServiceV1Server).IsPersonExists(ctx, req.(*IsPersonExistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MoviesPersonsServiceV1_UpdatePersonFields_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePersonFieldsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoviesPersonsServiceV1Server).UpdatePersonFields(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_movies_persons_service.moviesPersonsServiceV1/UpdatePersonFields",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoviesPersonsServiceV1Server).UpdatePersonFields(ctx, req.(*UpdatePersonFieldsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MoviesPersonsServiceV1_UpdatePerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoviesPersonsServiceV1Server).UpdatePerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_movies_persons_service.moviesPersonsServiceV1/UpdatePerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoviesPersonsServiceV1Server).UpdatePerson(ctx, req.(*UpdatePersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MoviesPersonsServiceV1_CreatePerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoviesPersonsServiceV1Server).CreatePerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_movies_persons_service.moviesPersonsServiceV1/CreatePerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoviesPersonsServiceV1Server).CreatePerson(ctx, req.(*CreatePersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MoviesPersonsServiceV1_DeletePersons_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePersonsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoviesPersonsServiceV1Server).DeletePersons(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_movies_persons_service.moviesPersonsServiceV1/DeletePersons",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoviesPersonsServiceV1Server).DeletePersons(ctx, req.(*DeletePersonsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MoviesPersonsServiceV1_ServiceDesc is the grpc.ServiceDesc for MoviesPersonsServiceV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MoviesPersonsServiceV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin_movies_persons_service.moviesPersonsServiceV1",
	HandlerType: (*MoviesPersonsServiceV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPersons",
			Handler:    _MoviesPersonsServiceV1_GetPersons_Handler,
		},
		{
			MethodName: "SearchPerson",
			Handler:    _MoviesPersonsServiceV1_SearchPerson_Handler,
		},
		{
			MethodName: "IsPersonWithIDExists",
			Handler:    _MoviesPersonsServiceV1_IsPersonWithIDExists_Handler,
		},
		{
			MethodName: "IsPersonExists",
			Handler:    _MoviesPersonsServiceV1_IsPersonExists_Handler,
		},
		{
			MethodName: "UpdatePersonFields",
			Handler:    _MoviesPersonsServiceV1_UpdatePersonFields_Handler,
		},
		{
			MethodName: "UpdatePerson",
			Handler:    _MoviesPersonsServiceV1_UpdatePerson_Handler,
		},
		{
			MethodName: "CreatePerson",
			Handler:    _MoviesPersonsServiceV1_CreatePerson_Handler,
		},
		{
			MethodName: "DeletePersons",
			Handler:    _MoviesPersonsServiceV1_DeletePersons_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin_movies_persons_service_v1.proto",
}
