package tq

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"os"
	"regexp"
)

type logHandler struct {
	fileLogger    slog.Handler
	consoleLogger slog.Handler
}

func NewLogHandler(fileWriter io.Writer, level *slog.LevelVar) *logHandler {
	return &logHandler{
		fileLogger: slog.NewTextHandler(fileWriter, &slog.HandlerOptions{
			AddSource: true,
		}),
		consoleLogger: slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
			Level: level,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				if a.Key == "time" {
					return slog.Attr{}
				}
				return a
			},
		}),
	}
}

func (h *logHandler) Enabled(context.Context, slog.Level) bool { return true }
func (h *logHandler) Handle(_ context.Context, record slog.Record) (err error) {
	var fileErr, consoleErr error

	messages := regexp.MustCompile("\r?\n|(\\+r)?\\+n").Split(record.Message, -1)
	for _, message := range messages {

		if h.fileLogger.Enabled(context.TODO(), record.Level) {
			fileErr = h.fileLogger.Handle(context.TODO(), slog.NewRecord(record.Time, record.Level, message, record.PC))
		}
		if h.consoleLogger.Enabled(context.TODO(), record.Level) {
			consoleErr = h.consoleLogger.Handle(context.TODO(), slog.NewRecord(record.Time, record.Level, message, record.PC))
		}
	}
	return errors.Join(fileErr, consoleErr)
}
func (h *logHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &logHandler{
		fileLogger:    h.fileLogger.WithAttrs(attrs),
		consoleLogger: h.consoleLogger.WithAttrs(attrs)}
}
func (h *logHandler) WithGroup(name string) slog.Handler {
	return &logHandler{
		fileLogger:    h.fileLogger.WithGroup(name),
		consoleLogger: h.consoleLogger.WithGroup(name)}
}

// Capture output for testing
func CaptureOutput(f func()) ([]byte, []byte) {
	pipeOut, stdOut, _ := os.Pipe()
	pipeErr, stdErr, _ := os.Pipe()
	os_Stdout := os.Stdout
	os_Stderr := os.Stderr
	defer func() {
		os.Stdout = os_Stdout
		os.Stderr = os_Stderr
	}()
	os.Stdout = stdOut
	os.Stderr = stdErr

	f()

	stdOut.Close()
	stdErr.Close()
	out, e := io.ReadAll(pipeOut)
	if e != nil {
		panic(e)
	}
	err, e := io.ReadAll(pipeErr)
	if e != nil {
		panic(e)
	}

	return out, err
}
