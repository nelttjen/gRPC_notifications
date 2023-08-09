package logger

import (
	"github.com/sirupsen/logrus"
	"notification_grpc/internal/config"
)

func GetLogger(name string) (*Logger, error) {
	if !initialized {
		panic(&notInitializedError{})
	}

	if config.LoggerEnabled == config.LogDisabled || config.Environment == config.Testing {
		name = "disabled"
	}

	for _, logger := range loggers {
		if logger.IsThisName(name) {
			return logger, nil
		}
	}

	return nil, &noSuchLogger{}
}

func LogIfExists(loggerName string, message string, level logrus.Level, requiredLevel uint8) {
	logger, err := GetLogger(loggerName)
	if err != nil {
		return
	}

	logger.Log(message, level, requiredLevel)
}

func LogfIfExists(loggerName string, message string, level logrus.Level, requiredLevel uint8, args ...interface{}) {
	logger, err := GetLogger(loggerName)
	if err != nil {
		return
	}

	logger.Logf(message, level, requiredLevel, args...)
}

func LoglnIfExists(loggerName string, message string, level logrus.Level, requiredLevel uint8) {
	logger, err := GetLogger(loggerName)
	if err != nil {
		return
	}

	logger.Logln(message, level, requiredLevel)
}
