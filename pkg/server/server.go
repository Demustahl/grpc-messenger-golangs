package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"grpc-messenger-golang/api/generated"

	"google.golang.org/grpc"
)

type Server struct {
	generated.UnimplementedMessengerServiceServer
	generated.UnimplementedAuthServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) RegisterUser(ctx context.Context, req *generated.User) (*generated.User, error) {
	// Пока просто возвращаем то, что получили
	log.Printf("Registering user: %+v", req)
	return req, nil
}

func (s *Server) Start(port int) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("failed to listen on port %d: %v", port, err)
	}

	grpcServer := grpc.NewServer()

	// Регистрируем сервисы
	generated.RegisterMessengerServiceServer(grpcServer, s)
	generated.RegisterAuthServiceServer(grpcServer, s)

	log.Printf("Server is running on port %d", port)
	return grpcServer.Serve(listener)
}
