// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: event/event.proto

package event1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	EventService_CreateEvent_FullMethodName         = "/EventService/CreateEvent"
	EventService_GetEvent_FullMethodName            = "/EventService/GetEvent"
	EventService_GetEventsTour_FullMethodName       = "/EventService/GetEventsTour"
	EventService_GetEventsToursMatch_FullMethodName = "/EventService/GetEventsToursMatch"
	EventService_UpdateMatch_FullMethodName         = "/EventService/UpdateMatch"
	EventService_GetAllEvent_FullMethodName         = "/EventService/GetAllEvent"
	EventService_GetAllTour_FullMethodName          = "/EventService/GetAllTour"
)

// EventServiceClient is the client API for EventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventServiceClient interface {
	CreateEvent(ctx context.Context, in *CreatEventReq, opts ...grpc.CallOption) (*EventStatus, error)
	GetEvent(ctx context.Context, in *GetEventReq, opts ...grpc.CallOption) (*Event, error)
	GetEventsTour(ctx context.Context, in *GetEventsTourReq, opts ...grpc.CallOption) (*Tour, error)
	GetEventsToursMatch(ctx context.Context, in *GetEventsToursMatchReq, opts ...grpc.CallOption) (*Match, error)
	UpdateMatch(ctx context.Context, in *UpdateMatchReq, opts ...grpc.CallOption) (*EventStatus, error)
	GetAllEvent(ctx context.Context, in *GetAllEventReq, opts ...grpc.CallOption) (*GetAllEventRes, error)
	GetAllTour(ctx context.Context, in *GetAllTourReq, opts ...grpc.CallOption) (*GetAllTourRes, error)
}

type eventServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventServiceClient(cc grpc.ClientConnInterface) EventServiceClient {
	return &eventServiceClient{cc}
}

func (c *eventServiceClient) CreateEvent(ctx context.Context, in *CreatEventReq, opts ...grpc.CallOption) (*EventStatus, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EventStatus)
	err := c.cc.Invoke(ctx, EventService_CreateEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) GetEvent(ctx context.Context, in *GetEventReq, opts ...grpc.CallOption) (*Event, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Event)
	err := c.cc.Invoke(ctx, EventService_GetEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) GetEventsTour(ctx context.Context, in *GetEventsTourReq, opts ...grpc.CallOption) (*Tour, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Tour)
	err := c.cc.Invoke(ctx, EventService_GetEventsTour_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) GetEventsToursMatch(ctx context.Context, in *GetEventsToursMatchReq, opts ...grpc.CallOption) (*Match, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Match)
	err := c.cc.Invoke(ctx, EventService_GetEventsToursMatch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) UpdateMatch(ctx context.Context, in *UpdateMatchReq, opts ...grpc.CallOption) (*EventStatus, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EventStatus)
	err := c.cc.Invoke(ctx, EventService_UpdateMatch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) GetAllEvent(ctx context.Context, in *GetAllEventReq, opts ...grpc.CallOption) (*GetAllEventRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllEventRes)
	err := c.cc.Invoke(ctx, EventService_GetAllEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) GetAllTour(ctx context.Context, in *GetAllTourReq, opts ...grpc.CallOption) (*GetAllTourRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllTourRes)
	err := c.cc.Invoke(ctx, EventService_GetAllTour_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventServiceServer is the server API for EventService service.
// All implementations must embed UnimplementedEventServiceServer
// for forward compatibility
type EventServiceServer interface {
	CreateEvent(context.Context, *CreatEventReq) (*EventStatus, error)
	GetEvent(context.Context, *GetEventReq) (*Event, error)
	GetEventsTour(context.Context, *GetEventsTourReq) (*Tour, error)
	GetEventsToursMatch(context.Context, *GetEventsToursMatchReq) (*Match, error)
	UpdateMatch(context.Context, *UpdateMatchReq) (*EventStatus, error)
	GetAllEvent(context.Context, *GetAllEventReq) (*GetAllEventRes, error)
	GetAllTour(context.Context, *GetAllTourReq) (*GetAllTourRes, error)
	mustEmbedUnimplementedEventServiceServer()
}

// UnimplementedEventServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEventServiceServer struct {
}

func (UnimplementedEventServiceServer) CreateEvent(context.Context, *CreatEventReq) (*EventStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}
func (UnimplementedEventServiceServer) GetEvent(context.Context, *GetEventReq) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvent not implemented")
}
func (UnimplementedEventServiceServer) GetEventsTour(context.Context, *GetEventsTourReq) (*Tour, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventsTour not implemented")
}
func (UnimplementedEventServiceServer) GetEventsToursMatch(context.Context, *GetEventsToursMatchReq) (*Match, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventsToursMatch not implemented")
}
func (UnimplementedEventServiceServer) UpdateMatch(context.Context, *UpdateMatchReq) (*EventStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMatch not implemented")
}
func (UnimplementedEventServiceServer) GetAllEvent(context.Context, *GetAllEventReq) (*GetAllEventRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllEvent not implemented")
}
func (UnimplementedEventServiceServer) GetAllTour(context.Context, *GetAllTourReq) (*GetAllTourRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTour not implemented")
}
func (UnimplementedEventServiceServer) mustEmbedUnimplementedEventServiceServer() {}

// UnsafeEventServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventServiceServer will
// result in compilation errors.
type UnsafeEventServiceServer interface {
	mustEmbedUnimplementedEventServiceServer()
}

func RegisterEventServiceServer(s grpc.ServiceRegistrar, srv EventServiceServer) {
	s.RegisterService(&EventService_ServiceDesc, srv)
}

func _EventService_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatEventReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_CreateEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).CreateEvent(ctx, req.(*CreatEventReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_GetEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).GetEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_GetEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).GetEvent(ctx, req.(*GetEventReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_GetEventsTour_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventsTourReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).GetEventsTour(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_GetEventsTour_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).GetEventsTour(ctx, req.(*GetEventsTourReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_GetEventsToursMatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventsToursMatchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).GetEventsToursMatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_GetEventsToursMatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).GetEventsToursMatch(ctx, req.(*GetEventsToursMatchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_UpdateMatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMatchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).UpdateMatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_UpdateMatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).UpdateMatch(ctx, req.(*UpdateMatchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_GetAllEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllEventReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).GetAllEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_GetAllEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).GetAllEvent(ctx, req.(*GetAllEventReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_GetAllTour_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllTourReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).GetAllTour(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_GetAllTour_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).GetAllTour(ctx, req.(*GetAllTourReq))
	}
	return interceptor(ctx, in, info, handler)
}

// EventService_ServiceDesc is the grpc.ServiceDesc for EventService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "EventService",
	HandlerType: (*EventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEvent",
			Handler:    _EventService_CreateEvent_Handler,
		},
		{
			MethodName: "GetEvent",
			Handler:    _EventService_GetEvent_Handler,
		},
		{
			MethodName: "GetEventsTour",
			Handler:    _EventService_GetEventsTour_Handler,
		},
		{
			MethodName: "GetEventsToursMatch",
			Handler:    _EventService_GetEventsToursMatch_Handler,
		},
		{
			MethodName: "UpdateMatch",
			Handler:    _EventService_UpdateMatch_Handler,
		},
		{
			MethodName: "GetAllEvent",
			Handler:    _EventService_GetAllEvent_Handler,
		},
		{
			MethodName: "GetAllTour",
			Handler:    _EventService_GetAllTour_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "event/event.proto",
}
