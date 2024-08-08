// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.12
// source: medals/medals.proto

package medals

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
	MedalsService_GetCountryRanking_FullMethodName = "/MedalsService/GetCountryRanking"
	MedalsService_Getathlete_FullMethodName        = "/MedalsService/Getathlete"
	MedalsService_MedalCreate_FullMethodName       = "/MedalsService/MedalCreate"
	MedalsService_MedalUpdate_FullMethodName       = "/MedalsService/MedalUpdate"
	MedalsService_MedalDelete_FullMethodName       = "/MedalsService/MedalDelete"
)

// MedalsServiceClient is the client API for MedalsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MedalsServiceClient interface {
	GetCountryRanking(ctx context.Context, in *MedalRankRequest, opts ...grpc.CallOption) (*MedalRankResponse, error)
	Getathlete(ctx context.Context, in *AthleteGetRequest, opts ...grpc.CallOption) (*AthleteGetRespone, error)
	MedalCreate(ctx context.Context, in *MedalCreateRequest, opts ...grpc.CallOption) (*GeneralResponseMedals, error)
	MedalUpdate(ctx context.Context, in *MedalUpdateRequest, opts ...grpc.CallOption) (*GeneralResponseMedals, error)
	MedalDelete(ctx context.Context, in *MedalDeleteRequest, opts ...grpc.CallOption) (*GeneralResponseMedals, error)
}

type medalsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMedalsServiceClient(cc grpc.ClientConnInterface) MedalsServiceClient {
	return &medalsServiceClient{cc}
}

func (c *medalsServiceClient) GetCountryRanking(ctx context.Context, in *MedalRankRequest, opts ...grpc.CallOption) (*MedalRankResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MedalRankResponse)
	err := c.cc.Invoke(ctx, MedalsService_GetCountryRanking_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medalsServiceClient) Getathlete(ctx context.Context, in *AthleteGetRequest, opts ...grpc.CallOption) (*AthleteGetRespone, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AthleteGetRespone)
	err := c.cc.Invoke(ctx, MedalsService_Getathlete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medalsServiceClient) MedalCreate(ctx context.Context, in *MedalCreateRequest, opts ...grpc.CallOption) (*GeneralResponseMedals, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GeneralResponseMedals)
	err := c.cc.Invoke(ctx, MedalsService_MedalCreate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medalsServiceClient) MedalUpdate(ctx context.Context, in *MedalUpdateRequest, opts ...grpc.CallOption) (*GeneralResponseMedals, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GeneralResponseMedals)
	err := c.cc.Invoke(ctx, MedalsService_MedalUpdate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medalsServiceClient) MedalDelete(ctx context.Context, in *MedalDeleteRequest, opts ...grpc.CallOption) (*GeneralResponseMedals, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GeneralResponseMedals)
	err := c.cc.Invoke(ctx, MedalsService_MedalDelete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MedalsServiceServer is the server API for MedalsService service.
// All implementations must embed UnimplementedMedalsServiceServer
// for forward compatibility
type MedalsServiceServer interface {
	GetCountryRanking(context.Context, *MedalRankRequest) (*MedalRankResponse, error)
	Getathlete(context.Context, *AthleteGetRequest) (*AthleteGetRespone, error)
	MedalCreate(context.Context, *MedalCreateRequest) (*GeneralResponseMedals, error)
	MedalUpdate(context.Context, *MedalUpdateRequest) (*GeneralResponseMedals, error)
	MedalDelete(context.Context, *MedalDeleteRequest) (*GeneralResponseMedals, error)
	mustEmbedUnimplementedMedalsServiceServer()
}

// UnimplementedMedalsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMedalsServiceServer struct {
}

func (UnimplementedMedalsServiceServer) GetCountryRanking(context.Context, *MedalRankRequest) (*MedalRankResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCountryRanking not implemented")
}
func (UnimplementedMedalsServiceServer) Getathlete(context.Context, *AthleteGetRequest) (*AthleteGetRespone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Getathlete not implemented")
}
func (UnimplementedMedalsServiceServer) MedalCreate(context.Context, *MedalCreateRequest) (*GeneralResponseMedals, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MedalCreate not implemented")
}
func (UnimplementedMedalsServiceServer) MedalUpdate(context.Context, *MedalUpdateRequest) (*GeneralResponseMedals, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MedalUpdate not implemented")
}
func (UnimplementedMedalsServiceServer) MedalDelete(context.Context, *MedalDeleteRequest) (*GeneralResponseMedals, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MedalDelete not implemented")
}
func (UnimplementedMedalsServiceServer) mustEmbedUnimplementedMedalsServiceServer() {}

// UnsafeMedalsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MedalsServiceServer will
// result in compilation errors.
type UnsafeMedalsServiceServer interface {
	mustEmbedUnimplementedMedalsServiceServer()
}

func RegisterMedalsServiceServer(s grpc.ServiceRegistrar, srv MedalsServiceServer) {
	s.RegisterService(&MedalsService_ServiceDesc, srv)
}

func _MedalsService_GetCountryRanking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MedalRankRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedalsServiceServer).GetCountryRanking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MedalsService_GetCountryRanking_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedalsServiceServer).GetCountryRanking(ctx, req.(*MedalRankRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedalsService_Getathlete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AthleteGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedalsServiceServer).Getathlete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MedalsService_Getathlete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedalsServiceServer).Getathlete(ctx, req.(*AthleteGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedalsService_MedalCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MedalCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedalsServiceServer).MedalCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MedalsService_MedalCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedalsServiceServer).MedalCreate(ctx, req.(*MedalCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedalsService_MedalUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MedalUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedalsServiceServer).MedalUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MedalsService_MedalUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedalsServiceServer).MedalUpdate(ctx, req.(*MedalUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedalsService_MedalDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MedalDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedalsServiceServer).MedalDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MedalsService_MedalDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedalsServiceServer).MedalDelete(ctx, req.(*MedalDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MedalsService_ServiceDesc is the grpc.ServiceDesc for MedalsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MedalsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MedalsService",
	HandlerType: (*MedalsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCountryRanking",
			Handler:    _MedalsService_GetCountryRanking_Handler,
		},
		{
			MethodName: "Getathlete",
			Handler:    _MedalsService_Getathlete_Handler,
		},
		{
			MethodName: "MedalCreate",
			Handler:    _MedalsService_MedalCreate_Handler,
		},
		{
			MethodName: "MedalUpdate",
			Handler:    _MedalsService_MedalUpdate_Handler,
		},
		{
			MethodName: "MedalDelete",
			Handler:    _MedalsService_MedalDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "medals/medals.proto",
}
