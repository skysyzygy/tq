package tq

import (
	"io"
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var log *slog.Logger
var pipeR *os.File
var pipeW *os.File
var level *slog.LevelVar

func init() {
	r, w, _ := os.Pipe()
	pipeR = r
	pipeW = w
	level = new(slog.LevelVar)
	log = slog.New(NewLogHandler(w, level))
}

// Test logging to console and file severity levels and source for default setup
func Test_LoggerDefault(t *testing.T) {
	r, w, _ := os.Pipe()
	console = *w // capture console output

	log.Error("Error")
	log.Warn("Warn")
	log.Info("Info")
	log.Debug("Debug")

	w.Close()
	consoleOutput, _ := io.ReadAll(r)
	fileOutput := make([]byte, 1024)
	pipeR.Read(fileOutput)

	assert.Contains(t, string(fileOutput), "Error")
	assert.Contains(t, string(fileOutput), "Warn")
	assert.Contains(t, string(fileOutput), "Info")
	assert.NotContains(t, string(fileOutput), "Debug")
	assert.Contains(t, string(fileOutput), "source")

	// Default is info level
	assert.Contains(t, string(consoleOutput), "Error")
	assert.Contains(t, string(consoleOutput), "Warn")
	assert.Contains(t, string(consoleOutput), "Info")
	assert.NotContains(t, string(consoleOutput), "Debug")
	assert.NotContains(t, string(consoleOutput), "source")
}

// Test logging to console and file severity levels and source for errors-only setup
func Test_LoggerErrors(t *testing.T) {
	r, w, _ := os.Pipe()
	console = *w

	level.Set(slog.LevelError)

	log.Error("Error")
	log.Warn("Warn")
	log.Info("Info")
	log.Debug("Debug")

	w.Close()
	consoleOutput, _ := io.ReadAll(r)
	fileOutput := make([]byte, 1024)
	pipeR.Read(fileOutput)

	assert.Contains(t, string(fileOutput), "Error")
	assert.Contains(t, string(fileOutput), "Warn")
	assert.Contains(t, string(fileOutput), "Info")
	assert.NotContains(t, string(fileOutput), "Debug")
	assert.Contains(t, string(fileOutput), "source")

	// Errors only this time
	assert.Contains(t, string(consoleOutput), "Error")
	assert.NotContains(t, string(consoleOutput), "Warn")
	assert.NotContains(t, string(consoleOutput), "Info")
	assert.NotContains(t, string(consoleOutput), "Debug")
	assert.NotContains(t, string(consoleOutput), "source")
}
