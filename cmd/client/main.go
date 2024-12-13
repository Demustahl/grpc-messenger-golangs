package main

import (
	"context"
	"log"
	"time"

	"grpc-messenger-golang/api/generated"

	"google.golang.org/grpc"
)

func main() {
	const serverAddress = "localhost:50051"

	// Устанавливаем соединение с сервером
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := generated.NewMessengerServiceClient(conn)

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
	}

	resp, err := client.RegisterUser(ctx, user)
	if err != nil {
		log.Fatalf("Failed to register user: %v", err)
	}

	log.Printf("User registered successfully: %+v", resp)
}
