package log_wrapper

import (
	"log/slog"
)

type dummyWriter struct{}

func (dummyWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func init() {

	logger := slog.New(slog.NewTextHandler(dummyWriter{}, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	LoggingService.logObj = logger
}
