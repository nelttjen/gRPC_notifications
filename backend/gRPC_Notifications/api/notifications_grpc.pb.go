// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.14.0
// source: notifications.proto

package api

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

const (
	CreateNotifications_CreateModels_FullMethodName = "/api.CreateNotifications/CreateModels"
)

// CreateNotificationsClient is the client API for CreateNotifications service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CreateNotificationsClient interface {
	CreateModels(ctx context.Context, in *NotificationCreateRequest, opts ...grpc.CallOption) (*NotificationCreateResponse, error)
}

type createNotificationsClient struct {
	cc grpc.ClientConnInterface
}

func NewCreateNotificationsClient(cc grpc.ClientConnInterface) CreateNotificationsClient {
	return &createNotificationsClient{cc}
}

func (c *createNotificationsClient) CreateModels(ctx context.Context, in *NotificationCreateRequest, opts ...grpc.CallOption) (*NotificationCreateResponse, error) {
	out := new(NotificationCreateResponse)
	err := c.cc.Invoke(ctx, CreateNotifications_CreateModels_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreateNotificationsServer is the server API for CreateNotifications service.
// All implementations must embed UnimplementedCreateNotificationsServer
// for forward compatibility
type CreateNotificationsServer interface {
	CreateModels(context.Context, *NotificationCreateRequest) (*NotificationCreateResponse, error)
	mustEmbedUnimplementedCreateNotificationsServer()
}

// UnimplementedCreateNotificationsServer must be embedded to have forward compatible implementations.
type UnimplementedCreateNotificationsServer struct {
}

func (UnimplementedCreateNotificationsServer) CreateModels(context.Context, *NotificationCreateRequest) (*NotificationCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateModels not implemented")
}
func (UnimplementedCreateNotificationsServer) mustEmbedUnimplementedCreateNotificationsServer() {}

// UnsafeCreateNotificationsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CreateNotificationsServer will
// result in compilation errors.
type UnsafeCreateNotificationsServer interface {
	mustEmbedUnimplementedCreateNotificationsServer()
}

func RegisterCreateNotificationsServer(s grpc.ServiceRegistrar, srv CreateNotificationsServer) {
	s.RegisterService(&CreateNotifications_ServiceDesc, srv)
}

func _CreateNotifications_CreateModels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreateNotificationsServer).CreateModels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CreateNotifications_CreateModels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreateNotificationsServer).CreateModels(ctx, req.(*NotificationCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CreateNotifications_ServiceDesc is the grpc.ServiceDesc for CreateNotifications service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CreateNotifications_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.CreateNotifications",
	HandlerType: (*CreateNotificationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateModels",
			Handler:    _CreateNotifications_CreateModels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notifications.proto",
}
