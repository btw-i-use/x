// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: uplodr.proto

package uplodr

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
	Image_Ping_FullMethodName   = "/within.website.x.xedn.uplodr.Image/Ping"
	Image_Upload_FullMethodName = "/within.website.x.xedn.uplodr.Image/Upload"
)

// ImageClient is the client API for Image service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImageClient interface {
	Ping(ctx context.Context, in *Echo, opts ...grpc.CallOption) (*Echo, error)
	Upload(ctx context.Context, in *UploadReq, opts ...grpc.CallOption) (*UploadResp, error)
}

type imageClient struct {
	cc grpc.ClientConnInterface
}

func NewImageClient(cc grpc.ClientConnInterface) ImageClient {
	return &imageClient{cc}
}

func (c *imageClient) Ping(ctx context.Context, in *Echo, opts ...grpc.CallOption) (*Echo, error) {
	out := new(Echo)
	err := c.cc.Invoke(ctx, Image_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageClient) Upload(ctx context.Context, in *UploadReq, opts ...grpc.CallOption) (*UploadResp, error) {
	out := new(UploadResp)
	err := c.cc.Invoke(ctx, Image_Upload_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageServer is the server API for Image service.
// All implementations must embed UnimplementedImageServer
// for forward compatibility
type ImageServer interface {
	Ping(context.Context, *Echo) (*Echo, error)
	Upload(context.Context, *UploadReq) (*UploadResp, error)
	mustEmbedUnimplementedImageServer()
}

// UnimplementedImageServer must be embedded to have forward compatible implementations.
type UnimplementedImageServer struct {
}

func (UnimplementedImageServer) Ping(context.Context, *Echo) (*Echo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedImageServer) Upload(context.Context, *UploadReq) (*UploadResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedImageServer) mustEmbedUnimplementedImageServer() {}

// UnsafeImageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImageServer will
// result in compilation errors.
type UnsafeImageServer interface {
	mustEmbedUnimplementedImageServer()
}

func RegisterImageServer(s grpc.ServiceRegistrar, srv ImageServer) {
	s.RegisterService(&Image_ServiceDesc, srv)
}

func _Image_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Echo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Image_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServer).Ping(ctx, req.(*Echo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Image_Upload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServer).Upload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Image_Upload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServer).Upload(ctx, req.(*UploadReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Image_ServiceDesc is the grpc.ServiceDesc for Image service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Image_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "within.website.x.xedn.uplodr.Image",
	HandlerType: (*ImageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Image_Ping_Handler,
		},
		{
			MethodName: "Upload",
			Handler:    _Image_Upload_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "uplodr.proto",
}
