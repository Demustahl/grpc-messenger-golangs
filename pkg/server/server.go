package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"grpc-messenger-golang/api/generated"
	"grpc-messenger-golang/pkg/db"
	"grpc-messenger-golang/pkg/utils"

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
	// Проверяем, существует ли email
	var existingUser generated.User
	err := s.db.UserCol.FindOne(ctx, bson.M{"email": req.Email}).Decode(&existingUser)
	if err != mongo.ErrNoDocuments {
		if err == nil {
			return nil, status.Errorf(codes.AlreadyExists, "user with this email already exists")
		}
		return nil, status.Errorf(codes.Internal, "error checking email existence: %v", err)
	}

	// Хэшируем пароль
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password")
	}

	// Заменяем пароль на его хэш
	req.Password = string(hashedPassword)

	// Сохраняем пользователя в базе
	_, err = s.db.UserCol.InsertOne(ctx, req)
	if err != nil {
		log.Printf("Failed to register user: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to register user")
	}

	log.Printf("User registered: %+v", req)
	return req, nil
}

func (s *Server) Login(ctx context.Context, req *generated.LoginRequest) (*generated.LoginResponse, error) {
	// Ищем пользователя в базе по email
	var user generated.User
	err := s.db.UserCol.FindOne(ctx, bson.M{"email": req.Email}).Decode(&user)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	// Проверяем пароль
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid credentials")
	}

	// Генерация JWT токена
	token, err := utils.GenerateJWT(user.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate token")
	}

	return &generated.LoginResponse{Token: token}, nil
}

func (s *Server) SendFriendRequest(ctx context.Context, req *generated.FriendRequest) (*emptypb.Empty, error) {
	// Проверяем токен
	userID, err := utils.VerifyToken(ctx)
	if err != nil {
		return nil, err
	}

	log.Printf("User %s sent a friend request to %s", userID, req.ToUserId)
	return &emptypb.Empty{}, nil
}
