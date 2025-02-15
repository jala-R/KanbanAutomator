package log_wrapper_test

import (
	"bufio"
	"bytes"
	"io"
	"log/slog"
	"strings"
	"testing"
)

func createSlogLogger(w io.Writer) *slog.Logger {
	return slog.New(slog.NewTextHandler(w, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
}

func TestDebug(t *testing.T) {

	//setup
	var (
		buffer         = bytes.NewBuffer([]byte{})
		logMsg         = "test1 from test"
		logMockService = createSlogLogger(buffer)
	)

	logMockService.Debug(logMsg)

	//verification
	result, err := getLogDetails(buffer)
	if err != nil {
		t.Fatalf("Got getting log details\n Want :3 Got : %d \n", len(result))
	}
	if result[1] != "DEBUG" {
		t.Fatalf("Got getting log details\n Want : DEBUG, Got :%s", result[1])
	}

	if result[2] != logMsg {
		t.Fatalf("Got getting log details\n Want : %s, Got :%s", logMsg, result[2])
	}
}

func TestWarn(t *testing.T) {
	//setup
	var (
		buffer         = bytes.NewBuffer([]byte{})
		logMsg         = "test1 from test"
		logMockService = createSlogLogger(buffer)
	)

	logMockService.Warn(logMsg)

	//verification
	result, err := getLogDetails(buffer)
	if err != nil {
		t.Fatalf("Got getting log details\n Want :3 Got : %d \n", len(result))
	}
	if result[1] != "WARN" {
		t.Fatalf("Got getting log details\n Want : DEBUG, Got :%s", result[1])
	}

	if result[2] != logMsg {
		t.Fatalf("Got getting log details\n Want : %s, Got :%s", logMsg, result[2])
	}
}

func TestError(t *testing.T) {
	//setup
	var (
		buffer         = bytes.NewBuffer([]byte{})
		logMsg         = "test1 from test"
		logMockService = createSlogLogger(buffer)
	)

	logMockService.Error(logMsg)

	//verification
	result, err := getLogDetails(buffer)
	if err != nil {
		t.Fatalf("Got getting log details\n Want :3 Got : %d \n", len(result))
	}
	if result[1] != "ERROR" {
		t.Fatalf("Got getting log details\n Want : DEBUG, Got :%s", result[1])
	}

	if result[2] != logMsg {
		t.Fatalf("Got getting log details\n Want : %s, Got :%s", logMsg, result[2])
	}
}

func TestInfo(t *testing.T) {
	//setup
	var (
		buffer         = bytes.NewBuffer([]byte{})
		logMsg         = "test1 from test"
		logMockService = createSlogLogger(buffer)
	)

	logMockService.Info(logMsg)

	//verification
	result, err := getLogDetails(buffer)
	if err != nil {
		t.Fatalf("Got getting log details\n Want :3 Got : %d \n", len(result))
	}
	if result[1] != "INFO" {
		t.Fatalf("Got getting log details\n Want : DEBUG, Got :%s", result[1])
	}

	if result[2] != logMsg {
		t.Fatalf("Got getting log details\n Want : %s, Got :%s", logMsg, result[2])
	}

}

func getLogDetails(buffer *bytes.Buffer) ([]string, error) {
	var ans = []string{}

	reader := bufio.NewReader(buffer)

	//extract time
	time, err := reader.ReadBytes(' ')
	if err != nil {
		return nil, err
	}
	ans = append(ans, string(time))

	//extract level
	level, err := reader.ReadBytes(' ')
	if err != nil {
		return nil, err
	}
	if level[len(level)-1] == ' ' {
		level = level[:len(level)-1]
	}
	levelSplit := strings.Split(string(level), "=")
	ans = append(ans, levelSplit[1])

	//extract message
	msg, err := reader.ReadBytes('\n')
	if err != nil {
		return nil, err
	}
	if msg[len(msg)-1] == '\n' {
		msg = msg[:len(msg)-1]
	}
	index := strings.Index(string(msg), "=")
	logMsg := string(msg[index+1:])
	ans = append(ans, logMsg[1:len(logMsg)-1])

	return ans, nil

}
