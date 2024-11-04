// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: sanguisuga.proto

package sanguisuga

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TV_List_FullMethodName     = "/within.website.x.sanguisuga.TV/List"
	TV_Track_FullMethodName    = "/within.website.x.sanguisuga.TV/Track"
	TV_Untrack_FullMethodName  = "/within.website.x.sanguisuga.TV/Untrack"
	TV_Snatches_FullMethodName = "/within.website.x.sanguisuga.TV/Snatches"
)

// TVClient is the client API for TV service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TVClient interface {
	List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Shows, error)
	Track(ctx context.Context, in *Show, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Untrack(ctx context.Context, in *Show, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Snatches(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TVSnatches, error)
}

type tVClient struct {
	cc grpc.ClientConnInterface
}

func NewTVClient(cc grpc.ClientConnInterface) TVClient {
	return &tVClient{cc}
}

func (c *tVClient) List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Shows, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Shows)
	err := c.cc.Invoke(ctx, TV_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tVClient) Track(ctx context.Context, in *Show, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TV_Track_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tVClient) Untrack(ctx context.Context, in *Show, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TV_Untrack_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tVClient) Snatches(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TVSnatches, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TVSnatches)
	err := c.cc.Invoke(ctx, TV_Snatches_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TVServer is the server API for TV service.
// All implementations must embed UnimplementedTVServer
// for forward compatibility.
type TVServer interface {
	List(context.Context, *emptypb.Empty) (*Shows, error)
	Track(context.Context, *Show) (*emptypb.Empty, error)
	Untrack(context.Context, *Show) (*emptypb.Empty, error)
	Snatches(context.Context, *emptypb.Empty) (*TVSnatches, error)
	mustEmbedUnimplementedTVServer()
}

// UnimplementedTVServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTVServer struct{}

func (UnimplementedTVServer) List(context.Context, *emptypb.Empty) (*Shows, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedTVServer) Track(context.Context, *Show) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Track not implemented")
}
func (UnimplementedTVServer) Untrack(context.Context, *Show) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Untrack not implemented")
}
func (UnimplementedTVServer) Snatches(context.Context, *emptypb.Empty) (*TVSnatches, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Snatches not implemented")
}
func (UnimplementedTVServer) mustEmbedUnimplementedTVServer() {}
func (UnimplementedTVServer) testEmbeddedByValue()            {}

// UnsafeTVServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TVServer will
// result in compilation errors.
type UnsafeTVServer interface {
	mustEmbedUnimplementedTVServer()
}

func RegisterTVServer(s grpc.ServiceRegistrar, srv TVServer) {
	// If the following call pancis, it indicates UnimplementedTVServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TV_ServiceDesc, srv)
}

func _TV_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TVServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TV_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TVServer).List(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TV_Track_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Show)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TVServer).Track(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TV_Track_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TVServer).Track(ctx, req.(*Show))
	}
	return interceptor(ctx, in, info, handler)
}

func _TV_Untrack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Show)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TVServer).Untrack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TV_Untrack_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TVServer).Untrack(ctx, req.(*Show))
	}
	return interceptor(ctx, in, info, handler)
}

func _TV_Snatches_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TVServer).Snatches(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TV_Snatches_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TVServer).Snatches(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// TV_ServiceDesc is the grpc.ServiceDesc for TV service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TV_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "within.website.x.sanguisuga.TV",
	HandlerType: (*TVServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _TV_List_Handler,
		},
		{
			MethodName: "Track",
			Handler:    _TV_Track_Handler,
		},
		{
			MethodName: "Untrack",
			Handler:    _TV_Untrack_Handler,
		},
		{
			MethodName: "Snatches",
			Handler:    _TV_Snatches_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sanguisuga.proto",
}

const (
	Anime_List_FullMethodName     = "/within.website.x.sanguisuga.Anime/List"
	Anime_Track_FullMethodName    = "/within.website.x.sanguisuga.Anime/Track"
	Anime_Untrack_FullMethodName  = "/within.website.x.sanguisuga.Anime/Untrack"
	Anime_Snatches_FullMethodName = "/within.website.x.sanguisuga.Anime/Snatches"
)

// AnimeClient is the client API for Anime service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnimeClient interface {
	List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Shows, error)
	Track(ctx context.Context, in *Show, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Untrack(ctx context.Context, in *Show, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Snatches(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AnimeSnatches, error)
}

type animeClient struct {
	cc grpc.ClientConnInterface
}

func NewAnimeClient(cc grpc.ClientConnInterface) AnimeClient {
	return &animeClient{cc}
}

func (c *animeClient) List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Shows, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Shows)
	err := c.cc.Invoke(ctx, Anime_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *animeClient) Track(ctx context.Context, in *Show, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Anime_Track_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *animeClient) Untrack(ctx context.Context, in *Show, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Anime_Untrack_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *animeClient) Snatches(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AnimeSnatches, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AnimeSnatches)
	err := c.cc.Invoke(ctx, Anime_Snatches_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnimeServer is the server API for Anime service.
// All implementations must embed UnimplementedAnimeServer
// for forward compatibility.
type AnimeServer interface {
	List(context.Context, *emptypb.Empty) (*Shows, error)
	Track(context.Context, *Show) (*emptypb.Empty, error)
	Untrack(context.Context, *Show) (*emptypb.Empty, error)
	Snatches(context.Context, *emptypb.Empty) (*AnimeSnatches, error)
	mustEmbedUnimplementedAnimeServer()
}

// UnimplementedAnimeServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAnimeServer struct{}

func (UnimplementedAnimeServer) List(context.Context, *emptypb.Empty) (*Shows, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedAnimeServer) Track(context.Context, *Show) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Track not implemented")
}
func (UnimplementedAnimeServer) Untrack(context.Context, *Show) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Untrack not implemented")
}
func (UnimplementedAnimeServer) Snatches(context.Context, *emptypb.Empty) (*AnimeSnatches, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Snatches not implemented")
}
func (UnimplementedAnimeServer) mustEmbedUnimplementedAnimeServer() {}
func (UnimplementedAnimeServer) testEmbeddedByValue()               {}

// UnsafeAnimeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnimeServer will
// result in compilation errors.
type UnsafeAnimeServer interface {
	mustEmbedUnimplementedAnimeServer()
}

func RegisterAnimeServer(s grpc.ServiceRegistrar, srv AnimeServer) {
	// If the following call pancis, it indicates UnimplementedAnimeServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Anime_ServiceDesc, srv)
}

func _Anime_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnimeServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Anime_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnimeServer).List(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Anime_Track_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Show)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnimeServer).Track(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Anime_Track_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnimeServer).Track(ctx, req.(*Show))
	}
	return interceptor(ctx, in, info, handler)
}

func _Anime_Untrack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Show)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnimeServer).Untrack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Anime_Untrack_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnimeServer).Untrack(ctx, req.(*Show))
	}
	return interceptor(ctx, in, info, handler)
}

func _Anime_Snatches_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnimeServer).Snatches(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Anime_Snatches_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnimeServer).Snatches(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Anime_ServiceDesc is the grpc.ServiceDesc for Anime service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Anime_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "within.website.x.sanguisuga.Anime",
	HandlerType: (*AnimeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Anime_List_Handler,
		},
		{
			MethodName: "Track",
			Handler:    _Anime_Track_Handler,
		},
		{
			MethodName: "Untrack",
			Handler:    _Anime_Untrack_Handler,
		},
		{
			MethodName: "Snatches",
			Handler:    _Anime_Snatches_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sanguisuga.proto",
}
