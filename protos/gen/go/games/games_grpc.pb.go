// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: games/games.proto

package games

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Games_GetUserGames_FullMethodName      = "/games.Games/GetUserGames"
	Games_AddUserGame_FullMethodName       = "/games.Games/AddUserGame"
	Games_DeleteUserGame_FullMethodName    = "/games.Games/DeleteUserGame"
	Games_SearchGamesByName_FullMethodName = "/games.Games/SearchGamesByName"
)

// GamesClient is the client API for Games service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GamesClient interface {
	GetUserGames(ctx context.Context, in *GetUserGamesRequest, opts ...grpc.CallOption) (*GetUserGamesResponse, error)
	AddUserGame(ctx context.Context, in *AddUserGameRequest, opts ...grpc.CallOption) (*AddUserGameResponse, error)
	DeleteUserGame(ctx context.Context, in *DeleteUserGameRequest, opts ...grpc.CallOption) (*DeleteUserGameResponse, error)
	SearchGamesByName(ctx context.Context, in *SearchGamesByNameRequest, opts ...grpc.CallOption) (*SearchGamesByNameResponse, error)
}

type gamesClient struct {
	cc grpc.ClientConnInterface
}

func NewGamesClient(cc grpc.ClientConnInterface) GamesClient {
	return &gamesClient{cc}
}

func (c *gamesClient) GetUserGames(ctx context.Context, in *GetUserGamesRequest, opts ...grpc.CallOption) (*GetUserGamesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserGamesResponse)
	err := c.cc.Invoke(ctx, Games_GetUserGames_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gamesClient) AddUserGame(ctx context.Context, in *AddUserGameRequest, opts ...grpc.CallOption) (*AddUserGameResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddUserGameResponse)
	err := c.cc.Invoke(ctx, Games_AddUserGame_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gamesClient) DeleteUserGame(ctx context.Context, in *DeleteUserGameRequest, opts ...grpc.CallOption) (*DeleteUserGameResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteUserGameResponse)
	err := c.cc.Invoke(ctx, Games_DeleteUserGame_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gamesClient) SearchGamesByName(ctx context.Context, in *SearchGamesByNameRequest, opts ...grpc.CallOption) (*SearchGamesByNameResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchGamesByNameResponse)
	err := c.cc.Invoke(ctx, Games_SearchGamesByName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GamesServer is the server API for Games service.
// All implementations must embed UnimplementedGamesServer
// for forward compatibility.
type GamesServer interface {
	GetUserGames(context.Context, *GetUserGamesRequest) (*GetUserGamesResponse, error)
	AddUserGame(context.Context, *AddUserGameRequest) (*AddUserGameResponse, error)
	DeleteUserGame(context.Context, *DeleteUserGameRequest) (*DeleteUserGameResponse, error)
	SearchGamesByName(context.Context, *SearchGamesByNameRequest) (*SearchGamesByNameResponse, error)
	mustEmbedUnimplementedGamesServer()
}

// UnimplementedGamesServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGamesServer struct{}

func (UnimplementedGamesServer) GetUserGames(context.Context, *GetUserGamesRequest) (*GetUserGamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserGames not implemented")
}
func (UnimplementedGamesServer) AddUserGame(context.Context, *AddUserGameRequest) (*AddUserGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserGame not implemented")
}
func (UnimplementedGamesServer) DeleteUserGame(context.Context, *DeleteUserGameRequest) (*DeleteUserGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserGame not implemented")
}
func (UnimplementedGamesServer) SearchGamesByName(context.Context, *SearchGamesByNameRequest) (*SearchGamesByNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchGamesByName not implemented")
}
func (UnimplementedGamesServer) mustEmbedUnimplementedGamesServer() {}
func (UnimplementedGamesServer) testEmbeddedByValue()               {}

// UnsafeGamesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GamesServer will
// result in compilation errors.
type UnsafeGamesServer interface {
	mustEmbedUnimplementedGamesServer()
}

func RegisterGamesServer(s grpc.ServiceRegistrar, srv GamesServer) {
	// If the following call pancis, it indicates UnimplementedGamesServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Games_ServiceDesc, srv)
}

func _Games_GetUserGames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserGamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GamesServer).GetUserGames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Games_GetUserGames_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GamesServer).GetUserGames(ctx, req.(*GetUserGamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Games_AddUserGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GamesServer).AddUserGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Games_AddUserGame_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GamesServer).AddUserGame(ctx, req.(*AddUserGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Games_DeleteUserGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GamesServer).DeleteUserGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Games_DeleteUserGame_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GamesServer).DeleteUserGame(ctx, req.(*DeleteUserGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Games_SearchGamesByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchGamesByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GamesServer).SearchGamesByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Games_SearchGamesByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GamesServer).SearchGamesByName(ctx, req.(*SearchGamesByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Games_ServiceDesc is the grpc.ServiceDesc for Games service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Games_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "games.Games",
	HandlerType: (*GamesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserGames",
			Handler:    _Games_GetUserGames_Handler,
		},
		{
			MethodName: "AddUserGame",
			Handler:    _Games_AddUserGame_Handler,
		},
		{
			MethodName: "DeleteUserGame",
			Handler:    _Games_DeleteUserGame_Handler,
		},
		{
			MethodName: "SearchGamesByName",
			Handler:    _Games_SearchGamesByName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "games/games.proto",
}
