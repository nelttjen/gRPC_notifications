package app

import (
	"google.golang.org/grpc"
	"log"
	"net"
	v1 "notification_grpc/api"
	"notification_grpc/internal/config"
	notificationRPC "notification_grpc/internal/grpc/notifications"
	"notification_grpc/pkg/env"
)

var _ App = &app{}

type app struct {
	server *grpc.Server
	env    env.Env
}

type App interface {
	RegisterServices() error
	Run() error
	LoadEnv(path string) error
}

func NewApp() App {
	server := grpc.NewServer()
	newEnv := env.NewEnv()

	newApp := app{
		server: server,
		env:    newEnv,
	}
	return &newApp
}

func (a *app) RegisterServices() (err error) {
	v1.RegisterCreateNotificationsServer(a.server, &notificationRPC.NotificationService{})
	return
}

func (a *app) LoadEnv(path string) error {
	err := a.env.LoadEnv(path)
	return err
}

func (a *app) Run() (err error) {
	lis, err := net.Listen(config.PROTOCOL, config.PORT)
	if err != nil {
		return err
	}
	err = a.LoadEnv("../.env")
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
		return err
	}

	err = a.RegisterServices()

	if err != nil {
		return err
	}

	log.Println("Start serving...")

	if err := a.server.Serve(lis); err != nil {
		return err
	}
	return nil
}
