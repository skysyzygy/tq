package tq

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"os"
)

type logHandler struct {
	fileLogger    slog.Handler
	consoleLogger slog.Handler
}

var console os.File

func init() {
	// For testing!
	console = *os.Stderr
}

func NewLogHandler(fileWriter io.Writer, level *slog.LevelVar) *logHandler {
	return &logHandler{
		fileLogger: slog.NewTextHandler(fileWriter, &slog.HandlerOptions{
			AddSource: true,
		}),
		consoleLogger: slog.NewTextHandler(&console, &slog.HandlerOptions{
			Level: level,
		}),
	}
}

func (h *logHandler) Enabled(context.Context, slog.Level) bool { return true }
func (h *logHandler) Handle(_ context.Context, record slog.Record) (err error) {
	if h.fileLogger.Enabled(context.TODO(), record.Level) {
		err = h.fileLogger.Handle(context.TODO(), record)
	}
	if h.consoleLogger.Enabled(context.TODO(), record.Level) {
		err = errors.Join(err, h.consoleLogger.Handle(context.TODO(), record))
	}
	return
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
