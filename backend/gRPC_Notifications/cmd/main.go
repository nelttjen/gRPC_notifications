package main

import (
	"log"
	app "notification_grpc/pkg/app"
)

func main() {
	err := app.NewApp().Run()
	panic("PANIIIICA")
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
