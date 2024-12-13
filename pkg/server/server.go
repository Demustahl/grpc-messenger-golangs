package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"grpc-messenger-golang/api/generated"
	"grpc-messenger-golang/pkg/db"

	"google.golang.org/grpc"
)

type Server struct {
	db *db.MongoDB
	generated.UnimplementedMessengerServiceServer
	generated.UnimplementedAuthServiceServer
}

func NewServer(db *db.MongoDB) *Server {
	return &Server{db: db}
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

func (s *Server) RegisterUser(ctx context.Context, req *generated.User) (*generated.User, error) {
	_, err := s.db.UserCol.InsertOne(ctx, req)
	if err != nil {
		log.Printf("Failed to register user: %v", err)
		return nil, err
	}
	log.Printf("User registered: %+v", req)
	return req, nil
}
