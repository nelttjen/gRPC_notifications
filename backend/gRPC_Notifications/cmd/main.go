package main

import (
	"log"
	"notification_grpc/pkg/app"
)

func main() {
	newApp := app.NewApp()

	if err := newApp.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
