package main

import (
	"log"
	"notification_grpc/pkg/app"
)

func main() {
	app := app.NewApp()

	if err := app.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
