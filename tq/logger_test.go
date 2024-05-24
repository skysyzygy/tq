package tq

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
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
}

func Test_CaptureOutput(t *testing.T) {

	out, err := CaptureOutput(func() { fmt.Fprint(os.Stdout, "Hello world!") })
	assert.Empty(t, err)
	assert.Equal(t, string(out), "Hello world!")

	out, err = CaptureOutput(func() { fmt.Fprint(os.Stderr, "Oops!") })
	assert.Empty(t, out)
	assert.Contains(t, string(err), "Oops")

}

// Test logging to console and file severity levels and source for default setup
func Test_LoggerDefault(t *testing.T) {

	_, consoleOutput := CaptureOutput(func() {
		log = slog.New(NewLogHandler(pipeW, level))

		log.Error("Error")
		log.Warn("Warn")
		log.Info("Info")
		log.Debug("Debug")
	})

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
	level.Set(slog.LevelError)

	_, consoleOutput := CaptureOutput(func() {
		log = slog.New(NewLogHandler(pipeW, level))

		log.Error("Error")
		log.Warn("Warn")
		log.Info("Info")
		log.Debug("Debug")
	})

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

// test that logger splits up messages with newlines
func Test_LoggerNewlines(t *testing.T) {
	_, consoleOutput := CaptureOutput(func() {
		log = slog.New(NewLogHandler(pipeW, level))
		log.Error("Error\nAnother Error")
	})

	fileOutput := make([]byte, 1024)
	pipeR.Read(fileOutput)

	assert.Len(t, strings.Split(string(consoleOutput), "\n"), 3)
	assert.Len(t, strings.Split(string(fileOutput), "\n"), 3)
}
