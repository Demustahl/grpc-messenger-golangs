// messenger.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: messenger.proto

package messenger

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
	MessengerService_SearchUsers_FullMethodName = "/messenger.MessengerService/SearchUsers"
	MessengerService_AddFriend_FullMethodName   = "/messenger.MessengerService/AddFriend"
	MessengerService_GetFriends_FullMethodName  = "/messenger.MessengerService/GetFriends"
	MessengerService_CreateChat_FullMethodName  = "/messenger.MessengerService/CreateChat"
	MessengerService_GetChats_FullMethodName    = "/messenger.MessengerService/GetChats"
	MessengerService_SendMessage_FullMethodName = "/messenger.MessengerService/SendMessage"
)

// MessengerServiceClient is the client API for MessengerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessengerServiceClient interface {
	// Поиск пользователей
	SearchUsers(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	// Добавление друга
	AddFriend(ctx context.Context, in *FriendRequest, opts ...grpc.CallOption) (*FriendResponse, error)
	// Получение списка друзей
	GetFriends(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*FriendListResponse, error)
	// Создание чата
	CreateChat(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (*ChatResponse, error)
	// Получение списка чатов
	GetChats(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*ChatListResponse, error)
	// Отправка сообщения
	SendMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponse, error)
}

type messengerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessengerServiceClient(cc grpc.ClientConnInterface) MessengerServiceClient {
	return &messengerServiceClient{cc}
}

func (c *messengerServiceClient) SearchUsers(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, MessengerService_SearchUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messengerServiceClient) AddFriend(ctx context.Context, in *FriendRequest, opts ...grpc.CallOption) (*FriendResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FriendResponse)
	err := c.cc.Invoke(ctx, MessengerService_AddFriend_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messengerServiceClient) GetFriends(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*FriendListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FriendListResponse)
	err := c.cc.Invoke(ctx, MessengerService_GetFriends_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messengerServiceClient) CreateChat(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (*ChatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChatResponse)
	err := c.cc.Invoke(ctx, MessengerService_CreateChat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messengerServiceClient) GetChats(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*ChatListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChatListResponse)
	err := c.cc.Invoke(ctx, MessengerService_GetChats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messengerServiceClient) SendMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MessageResponse)
	err := c.cc.Invoke(ctx, MessengerService_SendMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessengerServiceServer is the server API for MessengerService service.
// All implementations must embed UnimplementedMessengerServiceServer
// for forward compatibility.
type MessengerServiceServer interface {
	// Поиск пользователей
	SearchUsers(context.Context, *SearchRequest) (*SearchResponse, error)
	// Добавление друга
	AddFriend(context.Context, *FriendRequest) (*FriendResponse, error)
	// Получение списка друзей
	GetFriends(context.Context, *UserRequest) (*FriendListResponse, error)
	// Создание чата
	CreateChat(context.Context, *ChatRequest) (*ChatResponse, error)
	// Получение списка чатов
	GetChats(context.Context, *UserRequest) (*ChatListResponse, error)
	// Отправка сообщения
	SendMessage(context.Context, *MessageRequest) (*MessageResponse, error)
	mustEmbedUnimplementedMessengerServiceServer()
}

// UnimplementedMessengerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMessengerServiceServer struct{}

func (UnimplementedMessengerServiceServer) SearchUsers(context.Context, *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchUsers not implemented")
}
func (UnimplementedMessengerServiceServer) AddFriend(context.Context, *FriendRequest) (*FriendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFriend not implemented")
}
func (UnimplementedMessengerServiceServer) GetFriends(context.Context, *UserRequest) (*FriendListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriends not implemented")
}
func (UnimplementedMessengerServiceServer) CreateChat(context.Context, *ChatRequest) (*ChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChat not implemented")
}
func (UnimplementedMessengerServiceServer) GetChats(context.Context, *UserRequest) (*ChatListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChats not implemented")
}
func (UnimplementedMessengerServiceServer) SendMessage(context.Context, *MessageRequest) (*MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedMessengerServiceServer) mustEmbedUnimplementedMessengerServiceServer() {}
func (UnimplementedMessengerServiceServer) testEmbeddedByValue()                          {}

// UnsafeMessengerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessengerServiceServer will
// result in compilation errors.
type UnsafeMessengerServiceServer interface {
	mustEmbedUnimplementedMessengerServiceServer()
}

func RegisterMessengerServiceServer(s grpc.ServiceRegistrar, srv MessengerServiceServer) {
	// If the following call pancis, it indicates UnimplementedMessengerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MessengerService_ServiceDesc, srv)
}

func _MessengerService_SearchUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServiceServer).SearchUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessengerService_SearchUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServiceServer).SearchUsers(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessengerService_AddFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServiceServer).AddFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessengerService_AddFriend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServiceServer).AddFriend(ctx, req.(*FriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessengerService_GetFriends_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServiceServer).GetFriends(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessengerService_GetFriends_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServiceServer).GetFriends(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessengerService_CreateChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServiceServer).CreateChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessengerService_CreateChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServiceServer).CreateChat(ctx, req.(*ChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessengerService_GetChats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServiceServer).GetChats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessengerService_GetChats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServiceServer).GetChats(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessengerService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessengerService_SendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServiceServer).SendMessage(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MessengerService_ServiceDesc is the grpc.ServiceDesc for MessengerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessengerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "messenger.MessengerService",
	HandlerType: (*MessengerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchUsers",
			Handler:    _MessengerService_SearchUsers_Handler,
		},
		{
			MethodName: "AddFriend",
			Handler:    _MessengerService_AddFriend_Handler,
		},
		{
			MethodName: "GetFriends",
			Handler:    _MessengerService_GetFriends_Handler,
		},
		{
			MethodName: "CreateChat",
			Handler:    _MessengerService_CreateChat_Handler,
		},
		{
			MethodName: "GetChats",
			Handler:    _MessengerService_GetChats_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _MessengerService_SendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "messenger.proto",
}
