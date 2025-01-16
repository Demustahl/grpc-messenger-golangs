package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db := client.Database("grpc_messenger")

	// Users collection
	_, err = db.Collection("users").Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		log.Fatalf("Failed to create index on users: %v", err)
	}

	// Chats collection
	_, err = db.Collection("chats").Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{{Key: "participants", Value: 1}},
	})
	if err != nil {
		log.Fatalf("Failed to create index on chats: %v", err)
	}

	// Messages collection
	_, err = db.Collection("messages").Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{{Key: "chat_id", Value: 1}},
	})
	if err != nil {
		log.Fatalf("Failed to create index on messages: %v", err)
	}

	log.Println("Database and indexes initialized successfully!")
}
