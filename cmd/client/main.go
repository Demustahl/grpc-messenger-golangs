package main

import (
	"context"
	"log"
	"time"

	"grpc-messenger-golang/api/generated"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	const serverAddress = "localhost:50051"

	// Устанавливаем соединение с сервером
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Создаём клиента MessengerService
	messengerClient := generated.NewMessengerServiceClient(conn)

	// Создаём нового пользователя
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	user := &generated.User{
		Id:        "1",
		FirstName: "John",
		LastName:  "Doe",
		City:      "New York",
		Country:   "USA",
		Email:     "john.doe@example.com",
		Password:  "securepassword123",
	}

	resp, err := messengerClient.RegisterUser(ctx, user)
	if err != nil {
		log.Fatalf("Failed to register user: %v", err)
	}

	log.Printf("User registered successfully: %+v", resp)

	// Создаём клиента AuthService
	authClient := generated.NewAuthServiceClient(conn)

	// Выполняем вход и получаем токен
	token := testLogin(authClient)

	// Отправляем запрос на добавление друга
	testSendFriendRequest(messengerClient, token)
}

func testLogin(client generated.AuthServiceClient) string {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := &generated.LoginRequest{
		Email:    "john.doe@example.com",
		Password: "securepassword123",
	}

	resp, err := client.Login(ctx, req)
	if err != nil {
		log.Fatalf("Failed to login: %v", err)
	}

	log.Printf("Login successful, token: %s", resp.Token)
	return resp.Token // Возвращаем токен для использования в других запросах
}

func testSendFriendRequest(client generated.MessengerServiceClient, token string) {
	// Добавляем токен в заголовки
	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", token)

	req := &generated.FriendRequest{
		FromUserId: "1",
		ToUserId:   "2",
	}

	_, err := client.SendFriendRequest(ctx, req)
	if err != nil {
		log.Fatalf("Failed to send friend request: %v", err)
	}

	log.Println("Friend request sent successfully!")
}
