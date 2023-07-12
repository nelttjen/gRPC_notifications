package main

import (
	"fmt"
	"log"
	cfg "notification_grpc/internal/config"
	"notification_grpc/pkg/app"
	env2 "notification_grpc/pkg/env"
)

func main() {
	newApp := app.NewApp()

	env := env2.NewEnv()
	err := env.LoadEnv(cfg.EnvRoot)
	if err != nil {
		log.Fatal(err)
	}
	val, err := env.GetEnvAsString("POSTGRES_HOST")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)
	if err := newApp.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
