package app

import (
	"google.golang.org/grpc"
	"log"
	"net"
	v1 "notification_grpc/api"
	configConsts "notification_grpc/internal/config"
	myGrpc "notification_grpc/pkg/grpc"
)

var _ App = &app{}

type app struct {
}

type App interface {
	RegisterServices(server *grpc.Server) error
	Run() error
}

func NewApp() App {
	newApp := app{}
	return &newApp
}

func (a *app) RegisterServices(server *grpc.Server) (err error) {
	v1.RegisterCreateNotificationsServer(server, &myGrpc.NotificationService{})
	return
}

func (a *app) Run() (err error) {
	lis, err := net.Listen(configConsts.PROTOCOL, configConsts.PORT)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	err = a.RegisterServices(grpcServer)

	if err != nil {
		return err
	}

	log.Printf("Start serving...")

	if err := grpcServer.Serve(lis); err != nil {
		return err
	}
	return nil
}
