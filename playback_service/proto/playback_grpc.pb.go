// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	common "playback_service/proto/common"
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

// PlaybackServiceClient is the client API for PlaybackService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlaybackServiceClient interface {
	CreatePlaylist(ctx context.Context, in *CreatePlaylistRequest, opts ...grpc.CallOption) (*PlaylistResponse, error)
	RemovePlaylist(ctx context.Context, in *RemovePlaylistRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AddTracksToPlaylist(ctx context.Context, in *AddTracksToPlaylistRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RemoveTracksFromPlaylist(ctx context.Context, in *RemoveTracksFromPlaylistRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetPlaylistById(ctx context.Context, in *PlaylistIdRequest, opts ...grpc.CallOption) (*PlaylistResponse, error)
	PlayPlaylist(ctx context.Context, in *PlaylistIdRequest, opts ...grpc.CallOption) (*PlayPlaylistResponse, error)
	Status(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*common.StatusResponse, error)
}

type playbackServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPlaybackServiceClient(cc grpc.ClientConnInterface) PlaybackServiceClient {
	return &playbackServiceClient{cc}
}

func (c *playbackServiceClient) CreatePlaylist(ctx context.Context, in *CreatePlaylistRequest, opts ...grpc.CallOption) (*PlaylistResponse, error) {
	out := new(PlaylistResponse)
	err := c.cc.Invoke(ctx, "/PlaybackService/CreatePlaylist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playbackServiceClient) RemovePlaylist(ctx context.Context, in *RemovePlaylistRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/PlaybackService/RemovePlaylist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playbackServiceClient) AddTracksToPlaylist(ctx context.Context, in *AddTracksToPlaylistRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/PlaybackService/AddTracksToPlaylist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playbackServiceClient) RemoveTracksFromPlaylist(ctx context.Context, in *RemoveTracksFromPlaylistRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/PlaybackService/RemoveTracksFromPlaylist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playbackServiceClient) GetPlaylistById(ctx context.Context, in *PlaylistIdRequest, opts ...grpc.CallOption) (*PlaylistResponse, error) {
	out := new(PlaylistResponse)
	err := c.cc.Invoke(ctx, "/PlaybackService/GetPlaylistById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playbackServiceClient) PlayPlaylist(ctx context.Context, in *PlaylistIdRequest, opts ...grpc.CallOption) (*PlayPlaylistResponse, error) {
	out := new(PlayPlaylistResponse)
	err := c.cc.Invoke(ctx, "/PlaybackService/PlayPlaylist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playbackServiceClient) Status(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*common.StatusResponse, error) {
	out := new(common.StatusResponse)
	err := c.cc.Invoke(ctx, "/PlaybackService/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlaybackServiceServer is the server API for PlaybackService service.
// All implementations must embed UnimplementedPlaybackServiceServer
// for forward compatibility
type PlaybackServiceServer interface {
	CreatePlaylist(context.Context, *CreatePlaylistRequest) (*PlaylistResponse, error)
	RemovePlaylist(context.Context, *RemovePlaylistRequest) (*emptypb.Empty, error)
	AddTracksToPlaylist(context.Context, *AddTracksToPlaylistRequest) (*emptypb.Empty, error)
	RemoveTracksFromPlaylist(context.Context, *RemoveTracksFromPlaylistRequest) (*emptypb.Empty, error)
	GetPlaylistById(context.Context, *PlaylistIdRequest) (*PlaylistResponse, error)
	PlayPlaylist(context.Context, *PlaylistIdRequest) (*PlayPlaylistResponse, error)
	Status(context.Context, *emptypb.Empty) (*common.StatusResponse, error)
	mustEmbedUnimplementedPlaybackServiceServer()
}

// UnimplementedPlaybackServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPlaybackServiceServer struct {
}

func (UnimplementedPlaybackServiceServer) CreatePlaylist(context.Context, *CreatePlaylistRequest) (*PlaylistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePlaylist not implemented")
}
func (UnimplementedPlaybackServiceServer) RemovePlaylist(context.Context, *RemovePlaylistRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePlaylist not implemented")
}
func (UnimplementedPlaybackServiceServer) AddTracksToPlaylist(context.Context, *AddTracksToPlaylistRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTracksToPlaylist not implemented")
}
func (UnimplementedPlaybackServiceServer) RemoveTracksFromPlaylist(context.Context, *RemoveTracksFromPlaylistRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveTracksFromPlaylist not implemented")
}
func (UnimplementedPlaybackServiceServer) GetPlaylistById(context.Context, *PlaylistIdRequest) (*PlaylistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlaylistById not implemented")
}
func (UnimplementedPlaybackServiceServer) PlayPlaylist(context.Context, *PlaylistIdRequest) (*PlayPlaylistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayPlaylist not implemented")
}
func (UnimplementedPlaybackServiceServer) Status(context.Context, *emptypb.Empty) (*common.StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedPlaybackServiceServer) mustEmbedUnimplementedPlaybackServiceServer() {}

// UnsafePlaybackServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlaybackServiceServer will
// result in compilation errors.
type UnsafePlaybackServiceServer interface {
	mustEmbedUnimplementedPlaybackServiceServer()
}

func RegisterPlaybackServiceServer(s grpc.ServiceRegistrar, srv PlaybackServiceServer) {
	s.RegisterService(&PlaybackService_ServiceDesc, srv)
}

func _PlaybackService_CreatePlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePlaylistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaybackServiceServer).CreatePlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PlaybackService/CreatePlaylist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaybackServiceServer).CreatePlaylist(ctx, req.(*CreatePlaylistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaybackService_RemovePlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemovePlaylistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaybackServiceServer).RemovePlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PlaybackService/RemovePlaylist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaybackServiceServer).RemovePlaylist(ctx, req.(*RemovePlaylistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaybackService_AddTracksToPlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTracksToPlaylistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaybackServiceServer).AddTracksToPlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PlaybackService/AddTracksToPlaylist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaybackServiceServer).AddTracksToPlaylist(ctx, req.(*AddTracksToPlaylistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaybackService_RemoveTracksFromPlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveTracksFromPlaylistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaybackServiceServer).RemoveTracksFromPlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PlaybackService/RemoveTracksFromPlaylist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaybackServiceServer).RemoveTracksFromPlaylist(ctx, req.(*RemoveTracksFromPlaylistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaybackService_GetPlaylistById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaylistIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaybackServiceServer).GetPlaylistById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PlaybackService/GetPlaylistById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaybackServiceServer).GetPlaylistById(ctx, req.(*PlaylistIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaybackService_PlayPlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaylistIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaybackServiceServer).PlayPlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PlaybackService/PlayPlaylist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaybackServiceServer).PlayPlaylist(ctx, req.(*PlaylistIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaybackService_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaybackServiceServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PlaybackService/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaybackServiceServer).Status(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// PlaybackService_ServiceDesc is the grpc.ServiceDesc for PlaybackService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PlaybackService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "PlaybackService",
	HandlerType: (*PlaybackServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePlaylist",
			Handler:    _PlaybackService_CreatePlaylist_Handler,
		},
		{
			MethodName: "RemovePlaylist",
			Handler:    _PlaybackService_RemovePlaylist_Handler,
		},
		{
			MethodName: "AddTracksToPlaylist",
			Handler:    _PlaybackService_AddTracksToPlaylist_Handler,
		},
		{
			MethodName: "RemoveTracksFromPlaylist",
			Handler:    _PlaybackService_RemoveTracksFromPlaylist_Handler,
		},
		{
			MethodName: "GetPlaylistById",
			Handler:    _PlaybackService_GetPlaylistById_Handler,
		},
		{
			MethodName: "PlayPlaylist",
			Handler:    _PlaybackService_PlayPlaylist_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _PlaybackService_Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "playback.proto",
}
