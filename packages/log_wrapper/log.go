package log_wrapper

import (
	"fmt"
	"log/slog"
	"os"
)

var (
	LoggingService loggingServiceImpl
)

type ILogging interface {
	Debug(string)
	Info(string)
	Warn(string)
	Error(string)
}

type loggingServiceImpl struct {
	logObj *slog.Logger
}

func (l loggingServiceImpl) Debug(msg string) {
	l.logObj.Debug(msg)
}

func (l loggingServiceImpl) Warn(msg string) {
	l.logObj.Warn(msg)
}

func (l loggingServiceImpl) Info(msg string) {
	l.logObj.Info(msg)
}

func (l loggingServiceImpl) Error(msg string) {
	l.logObj.Error(msg)
}

func init() {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	logger := slog.New(slog.NewTextHandler(file, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	LoggingService = loggingServiceImpl{logger}
}
