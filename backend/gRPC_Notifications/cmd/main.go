package main

import (
	"github.com/sirupsen/logrus"
	"notification_grpc/internal/config"
	"notification_grpc/pkg/app"
	"notification_grpc/pkg/logger"
)

func main() {
	logger.InitLoggers()
	app := app.NewApp()

	if err := app.Run(); err != nil {
		logger.LogflnIfExists("error", "Failed to run server: %v", logrus.FatalLevel, config.LoggerLevelNone, err)
	}
}
