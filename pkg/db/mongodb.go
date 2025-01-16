package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client     *mongo.Client
	UserCol    *mongo.Collection
	ChatCol    *mongo.Collection
	MessageCol *mongo.Collection
}

func Connect(uri string, dbName string) (*MongoDB, error) {
	clientOptions := options.Client().ApplyURI(uri).SetMaxPoolSize(20) // Оптимизация пула соединений
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB")

	db := client.Database(dbName)
	return &MongoDB{
		Client:     client,
		UserCol:    db.Collection("users"),
		ChatCol:    db.Collection("chats"),
		MessageCol: db.Collection("messages"),
	}, nil
}
