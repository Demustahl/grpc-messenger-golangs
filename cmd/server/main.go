package main

import (
	"log"

	"grpc-messenger-golang/pkg/server"
)

func main() {
	const port = 50051

	srv := server.NewServer()
	if err := srv.Start(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
