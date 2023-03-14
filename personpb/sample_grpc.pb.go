// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/sample.proto

package personpb

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

// PersonServiceClient is the client API for PersonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PersonServiceClient interface {
	CreatePerson(ctx context.Context, in *CreatePersonRequest, opts ...grpc.CallOption) (*CreatePersonResponse, error)
	ReadPerson(ctx context.Context, in *ReadPersonRequest, opts ...grpc.CallOption) (*ReadPersonResponse, error)
	ListPerson(ctx context.Context, in *ListPersonRequest, opts ...grpc.CallOption) (PersonService_ListPersonClient, error)
}

type personServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPersonServiceClient(cc grpc.ClientConnInterface) PersonServiceClient {
	return &personServiceClient{cc}
}

func (c *personServiceClient) CreatePerson(ctx context.Context, in *CreatePersonRequest, opts ...grpc.CallOption) (*CreatePersonResponse, error) {
	out := new(CreatePersonResponse)
	err := c.cc.Invoke(ctx, "/PersonService/CreatePerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personServiceClient) ReadPerson(ctx context.Context, in *ReadPersonRequest, opts ...grpc.CallOption) (*ReadPersonResponse, error) {
	out := new(ReadPersonResponse)
	err := c.cc.Invoke(ctx, "/PersonService/ReadPerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personServiceClient) ListPerson(ctx context.Context, in *ListPersonRequest, opts ...grpc.CallOption) (PersonService_ListPersonClient, error) {
	stream, err := c.cc.NewStream(ctx, &PersonService_ServiceDesc.Streams[0], "/PersonService/ListPerson", opts...)
	if err != nil {
		return nil, err
	}
	x := &personServiceListPersonClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PersonService_ListPersonClient interface {
	Recv() (*ListPersonResponse, error)
	grpc.ClientStream
}

type personServiceListPersonClient struct {
	grpc.ClientStream
}

func (x *personServiceListPersonClient) Recv() (*ListPersonResponse, error) {
	m := new(ListPersonResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PersonServiceServer is the server API for PersonService service.
// All implementations must embed UnimplementedPersonServiceServer
// for forward compatibility
type PersonServiceServer interface {
	CreatePerson(context.Context, *CreatePersonRequest) (*CreatePersonResponse, error)
	ReadPerson(context.Context, *ReadPersonRequest) (*ReadPersonResponse, error)
	ListPerson(*ListPersonRequest, PersonService_ListPersonServer) error
	mustEmbedUnimplementedPersonServiceServer()
}

// UnimplementedPersonServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPersonServiceServer struct {
}

func (UnimplementedPersonServiceServer) CreatePerson(context.Context, *CreatePersonRequest) (*CreatePersonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePerson not implemented")
}
func (UnimplementedPersonServiceServer) ReadPerson(context.Context, *ReadPersonRequest) (*ReadPersonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadPerson not implemented")
}
func (UnimplementedPersonServiceServer) ListPerson(*ListPersonRequest, PersonService_ListPersonServer) error {
	return status.Errorf(codes.Unimplemented, "method ListPerson not implemented")
}
func (UnimplementedPersonServiceServer) mustEmbedUnimplementedPersonServiceServer() {}

// UnsafePersonServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PersonServiceServer will
// result in compilation errors.
type UnsafePersonServiceServer interface {
	mustEmbedUnimplementedPersonServiceServer()
}

func RegisterPersonServiceServer(s grpc.ServiceRegistrar, srv PersonServiceServer) {
	s.RegisterService(&PersonService_ServiceDesc, srv)
}

func _PersonService_CreatePerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).CreatePerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PersonService/CreatePerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).CreatePerson(ctx, req.(*CreatePersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonService_ReadPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadPersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).ReadPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PersonService/ReadPerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).ReadPerson(ctx, req.(*ReadPersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonService_ListPerson_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListPersonRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PersonServiceServer).ListPerson(m, &personServiceListPersonServer{stream})
}

type PersonService_ListPersonServer interface {
	Send(*ListPersonResponse) error
	grpc.ServerStream
}

type personServiceListPersonServer struct {
	grpc.ServerStream
}

func (x *personServiceListPersonServer) Send(m *ListPersonResponse) error {
	return x.ServerStream.SendMsg(m)
}

// PersonService_ServiceDesc is the grpc.ServiceDesc for PersonService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PersonService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "PersonService",
	HandlerType: (*PersonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePerson",
			Handler:    _PersonService_CreatePerson_Handler,
		},
		{
			MethodName: "ReadPerson",
			Handler:    _PersonService_ReadPerson_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListPerson",
			Handler:       _PersonService_ListPerson_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/sample.proto",
}
