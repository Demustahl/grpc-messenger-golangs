package server

import (
	"context"
	"fmt"
	"grpc-messenger-golang/api/generated/auth"
	"grpc-messenger-golang/api/generated/messenger"
	"grpc-messenger-golang/pkg/db"
	"grpc-messenger-golang/pkg/utils"
	"log"
	"net"

	"google.golang.org/grpc/reflection"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	db *db.MongoDB
	messenger.UnimplementedMessengerServiceServer
	auth.UnimplementedAuthServiceServer
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

	// Register services
	messenger.RegisterMessengerServiceServer(grpcServer, s)
	auth.RegisterAuthServiceServer(grpcServer, s)

	// Включаем Reflection
	reflection.Register(grpcServer)

	log.Printf("Server is running on port %d", port)
	return grpcServer.Serve(listener)
}

func (s *Server) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.AuthResponse, error) {
	// Check if email exists
	var existingUser auth.RegisterRequest
	err := s.db.UserCol.FindOne(ctx, bson.M{"email": req.Email}).Decode(&existingUser)
	if err != mongo.ErrNoDocuments {
		if err == nil {
			return nil, status.Errorf(codes.AlreadyExists, "user with this email already exists")
		}
		return nil, status.Errorf(codes.Internal, "error checking email existence: %v", err)
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password")
	}

	// Save user to database
	req.Password = string(hashedPassword)
	_, err = s.db.UserCol.InsertOne(ctx, req)
	if err != nil {
		log.Printf("Failed to register user: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to register user")
	}

	log.Printf("User registered: %+v", req)
	return &auth.AuthResponse{Token: "User registered successfully"}, nil
}

func (s *Server) Login(ctx context.Context, req *auth.LoginRequest) (*auth.AuthResponse, error) {
	// Find user by email
	var user auth.RegisterRequest
	err := s.db.UserCol.FindOne(ctx, bson.M{"email": req.Email}).Decode(&user)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid credentials")
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.Email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate token")
	}

	return &auth.AuthResponse{Token: token}, nil
}

func (s *Server) SendFriendRequest(ctx context.Context, req *messenger.FriendRequest) (*emptypb.Empty, error) {
	// Verify token
	userID, err := utils.VerifyToken(ctx)
	if err != nil {
		return nil, err
	}

	log.Printf("User %s sent a friend request to %s", userID, req.FriendId)
	return &emptypb.Empty{}, nil
}
