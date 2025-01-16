package main

import (
	"context"
	"grpc-messenger-golang/pkg/db"
	"grpc-messenger-golang/pkg/server"
	"log"
)

func main() {
	const (
		port   = 50051
		dbURI  = "mongodb://localhost:27017"
		dbName = "messenger"
	)

	// Подключение к MongoDB
	mongoDB, err := db.Connect(dbURI, dbName)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer mongoDB.Client.Disconnect(context.Background())

	// Запуск сервера
	srv := server.NewServer(mongoDB)
	if err := srv.Start(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
