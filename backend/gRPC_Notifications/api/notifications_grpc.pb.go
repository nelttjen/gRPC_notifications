// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// CreateNotificationsClient is the client API for CreateNotifications service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CreateNotificationsClient interface {
	CreateNotificationsAction(ctx context.Context, in *NotificationCreateRequest, opts ...grpc.CallOption) (*NotificationCreateResponse, error)
	Create(ctx context.Context, in *NotificationCreateManualRequest, opts ...grpc.CallOption) (*NotificationCreateResponse, error)
	GetNotifications(ctx context.Context, in *UserNotificationsRequest, opts ...grpc.CallOption) (*UserNotificationsResponse, error)
	GetMassNotification(ctx context.Context, in *UserMassNotificationRequest, opts ...grpc.CallOption) (*UserMassNotificationResponse, error)
	DeleteNotifications(ctx context.Context, in *NotificationManageRequest, opts ...grpc.CallOption) (*NotificationManageResponse, error)
	MarkAsReadNotifications(ctx context.Context, in *NotificationManageRequest, opts ...grpc.CallOption) (*NotificationManageResponse, error)
}

type createNotificationsClient struct {
	cc grpc.ClientConnInterface
}

func NewCreateNotificationsClient(cc grpc.ClientConnInterface) CreateNotificationsClient {
	return &createNotificationsClient{cc}
}

func (c *createNotificationsClient) CreateNotificationsAction(ctx context.Context, in *NotificationCreateRequest, opts ...grpc.CallOption) (*NotificationCreateResponse, error) {
	out := new(NotificationCreateResponse)
	err := c.cc.Invoke(ctx, "/api.CreateNotifications/CreateNotificationsAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *createNotificationsClient) Create(ctx context.Context, in *NotificationCreateManualRequest, opts ...grpc.CallOption) (*NotificationCreateResponse, error) {
	out := new(NotificationCreateResponse)
	err := c.cc.Invoke(ctx, "/api.CreateNotifications/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *createNotificationsClient) GetNotifications(ctx context.Context, in *UserNotificationsRequest, opts ...grpc.CallOption) (*UserNotificationsResponse, error) {
	out := new(UserNotificationsResponse)
	err := c.cc.Invoke(ctx, "/api.CreateNotifications/GetNotifications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *createNotificationsClient) GetMassNotification(ctx context.Context, in *UserMassNotificationRequest, opts ...grpc.CallOption) (*UserMassNotificationResponse, error) {
	out := new(UserMassNotificationResponse)
	err := c.cc.Invoke(ctx, "/api.CreateNotifications/GetMassNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *createNotificationsClient) DeleteNotifications(ctx context.Context, in *NotificationManageRequest, opts ...grpc.CallOption) (*NotificationManageResponse, error) {
	out := new(NotificationManageResponse)
	err := c.cc.Invoke(ctx, "/api.CreateNotifications/DeleteNotifications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *createNotificationsClient) MarkAsReadNotifications(ctx context.Context, in *NotificationManageRequest, opts ...grpc.CallOption) (*NotificationManageResponse, error) {
	out := new(NotificationManageResponse)
	err := c.cc.Invoke(ctx, "/api.CreateNotifications/MarkAsReadNotifications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreateNotificationsServer is the server API for CreateNotifications service.
// All implementations must embed UnimplementedCreateNotificationsServer
// for forward compatibility
type CreateNotificationsServer interface {
	CreateNotificationsAction(context.Context, *NotificationCreateRequest) (*NotificationCreateResponse, error)
	Create(context.Context, *NotificationCreateManualRequest) (*NotificationCreateResponse, error)
	GetNotifications(context.Context, *UserNotificationsRequest) (*UserNotificationsResponse, error)
	GetMassNotification(context.Context, *UserMassNotificationRequest) (*UserMassNotificationResponse, error)
	DeleteNotifications(context.Context, *NotificationManageRequest) (*NotificationManageResponse, error)
	MarkAsReadNotifications(context.Context, *NotificationManageRequest) (*NotificationManageResponse, error)
	mustEmbedUnimplementedCreateNotificationsServer()
}

// UnimplementedCreateNotificationsServer must be embedded to have forward compatible implementations.
type UnimplementedCreateNotificationsServer struct {
}

func (UnimplementedCreateNotificationsServer) CreateNotificationsAction(context.Context, *NotificationCreateRequest) (*NotificationCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNotificationsAction not implemented")
}
func (UnimplementedCreateNotificationsServer) Create(context.Context, *NotificationCreateManualRequest) (*NotificationCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCreateNotificationsServer) GetNotifications(context.Context, *UserNotificationsRequest) (*UserNotificationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotifications not implemented")
}
func (UnimplementedCreateNotificationsServer) GetMassNotification(context.Context, *UserMassNotificationRequest) (*UserMassNotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMassNotification not implemented")
}
func (UnimplementedCreateNotificationsServer) DeleteNotifications(context.Context, *NotificationManageRequest) (*NotificationManageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNotifications not implemented")
}
func (UnimplementedCreateNotificationsServer) MarkAsReadNotifications(context.Context, *NotificationManageRequest) (*NotificationManageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkAsReadNotifications not implemented")
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

func _CreateNotifications_CreateNotificationsAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreateNotificationsServer).CreateNotificationsAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CreateNotifications/CreateNotificationsAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreateNotificationsServer).CreateNotificationsAction(ctx, req.(*NotificationCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreateNotifications_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationCreateManualRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreateNotificationsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CreateNotifications/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreateNotificationsServer).Create(ctx, req.(*NotificationCreateManualRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreateNotifications_GetNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserNotificationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreateNotificationsServer).GetNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CreateNotifications/GetNotifications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreateNotificationsServer).GetNotifications(ctx, req.(*UserNotificationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreateNotifications_GetMassNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserMassNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreateNotificationsServer).GetMassNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CreateNotifications/GetMassNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreateNotificationsServer).GetMassNotification(ctx, req.(*UserMassNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreateNotifications_DeleteNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationManageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreateNotificationsServer).DeleteNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CreateNotifications/DeleteNotifications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreateNotificationsServer).DeleteNotifications(ctx, req.(*NotificationManageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreateNotifications_MarkAsReadNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationManageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreateNotificationsServer).MarkAsReadNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CreateNotifications/MarkAsReadNotifications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreateNotificationsServer).MarkAsReadNotifications(ctx, req.(*NotificationManageRequest))
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
			MethodName: "CreateNotificationsAction",
			Handler:    _CreateNotifications_CreateNotificationsAction_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _CreateNotifications_Create_Handler,
		},
		{
			MethodName: "GetNotifications",
			Handler:    _CreateNotifications_GetNotifications_Handler,
		},
		{
			MethodName: "GetMassNotification",
			Handler:    _CreateNotifications_GetMassNotification_Handler,
		},
		{
			MethodName: "DeleteNotifications",
			Handler:    _CreateNotifications_DeleteNotifications_Handler,
		},
		{
			MethodName: "MarkAsReadNotifications",
			Handler:    _CreateNotifications_MarkAsReadNotifications_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notifications.proto",
}
