package app

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	v1 "notification_grpc/api"
	cfg "notification_grpc/internal/config"
	notificationRPC "notification_grpc/internal/grpc/notifications"
	"notification_grpc/pkg/env"
	"notification_grpc/pkg/logger"
)

var _ App = &app{}

type app struct {
	server *grpc.Server
	env    env.Env
}

type App interface {
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

func (a *app) LoadEnv(path string) error {
	err := a.env.LoadEnv(path)
	return err
}

func (a *app) Run() (err error) {
	lis, err := net.Listen(cfg.PROTOCOL, cfg.PORT)
	if err != nil {
		return err
	}

	err = a.LoadEnv(cfg.EnvRoot)
	if err != nil {
		logger.LogflnIfExists("error", "Failed to load .env file: %v", logrus.FatalLevel, cfg.LoggerLevelImportant, err)
		return err
	}

	v1.RegisterCreateNotificationsServer(a.server, &notificationRPC.NotificationService{})
	logger.LoglnIfExists("info", "Initialization done, Start serving...", logrus.InfoLevel, cfg.LoggerLevelImportant)

	if err := a.server.Serve(lis); err != nil {
		return err
	}
	return nil
}
