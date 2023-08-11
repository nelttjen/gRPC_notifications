package logger

import (
	"fmt"
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

func LogfIfExists(loggerName string, format string, level logrus.Level, requiredLevel uint8, args ...interface{}) {
	logger, err := GetLogger(loggerName)
	if err != nil {
		return
	}

	message := fmt.Sprintf(format, args...)

	logger.Log(message, level, requiredLevel)
}

func LoglnIfExists(loggerName string, message string, level logrus.Level, requiredLevel uint8) {
	logger, err := GetLogger(loggerName)
	if err != nil {
		return
	}

	message = fmt.Sprintf("%s\n", message)

	logger.Log(message, level, requiredLevel)
}

func LogflnIfExists(loggerName string, format string, level logrus.Level, requiredLevel uint8, args ...interface{}) {
	format = fmt.Sprintf("%s\n", format)
	LogfIfExists(loggerName, format, level, requiredLevel, args...)
}
