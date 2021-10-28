package jsonlog

import (
	"encoding/json"
	"io"
	"os"
	"runtime/debug"
	"sync"
	"time"
)

// Logger holds the output destination that the log entries
// will be written to, the minimum severity level that log
// entries will be written for, and a mutex for concurrent writes.
type Logger struct {
	out      io.Writer
	minLevel Level
	mu       sync.Mutex
}

// New create a new Logger
func New(out io.Writer, minLevel Level) *Logger {
	return &Logger{
		out:      out,
		minLevel: minLevel,
	}
}

// Default write to stdout with the minimum level set to LevelDebug.
func Default() *Logger {
	return &Logger{
		out:      os.Stdout,
		minLevel: LevelDebug,
	}
}

// Debug writes a log entry at LevelDebug to the output destination.
func (l *Logger) Debug(msg string, props map[string]string) {
	l.print(LevelInfo, msg, props)
}

// Info writes a log entry at LevelInfo to the output destination.
func (l *Logger) Info(msg string, props map[string]string) {
	l.print(LevelInfo, msg, props)
}

// Err writes a log entry at LevelError to the output destination.
func (l *Logger) Err(err error, props map[string]string) {
	l.print(LevelError, err.Error(), props)
}

// Fatal writes a log entry at LevelFatal to the output destination
// and exit 1.
func (l *Logger) Fatal(err error, props map[string]string) {
	l.print(LevelFatal, err.Error(), props)
	os.Exit(1)
}

func (l *Logger) print(level Level, msg string, props map[string]string) (int, error) {
	if level < l.minLevel {
		return 0, nil
	}
	aux := struct {
		Level      string            `json:"level"`
		Time       string            `json:"time"`
		Message    string            `json:"message"`
		Properties map[string]string `json:"properties,omitempty"`
		Trace      string            `json:"trace,omitempty"`
	}{
		Level:      level.String(),
		Time:       time.Now().UTC().Format(time.RFC3339),
		Message:    msg,
		Properties: props,
	}

	if level >= LevelError {
		aux.Trace = string(debug.Stack())
	}
	var line []byte
	line, err := json.Marshal(aux)
	if err != nil {
		line = []byte(LevelError.String() + ": unable to marshal log message:" + err.Error())
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.out.Write(append(line, '\n'))
}

// WriteRaw writes a raw log entry at LevelError to the output destination.
func (l *Logger) WriteRaw(msg []byte) (int, error) {
	return l.print(LevelError, string(msg), nil)
}
