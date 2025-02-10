package logging

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"os"
)

func createSlogLogger(w io.Writer) *slog.Logger {
	return slog.New(slog.NewTextHandler(w, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
}

func MockLogSerivce(buffer *bytes.Buffer) logger {

	return logger{createSlogLogger(buffer)}
}

func init() {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	LogSerive = logger{
		createSlogLogger(file),
	}
}
