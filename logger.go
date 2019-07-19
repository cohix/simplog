package simplog

import (
	"fmt"
)

//LogLevelError and others represent the log levels
const (
	LevelError     = 1
	LevelWarn      = 2
	LevelInfo      = 3
	LevelDebug     = 4
	LevelTrace     = 5
	LevelSensitive = 6
)

//SLogger is a simplogger
type SLogger struct {
	Level int
}

// New returns a new SLogger
func New(level int) *SLogger {
	sl := &SLogger{
		Level: level,
	}

	return sl
}

// Error logs an error
func (l *SLogger) Error(err error) {
	fmt.Printf("(E) %s\n", err.Error())
}

// ErrorString logs an error
func (l *SLogger) ErrorString(msg string) {
	fmt.Printf("(E) %s\n", msg)
}

// Warn logs a warning
func (l *SLogger) Warn(msg string) {
	if l.Level >= LevelWarn {
		fmt.Printf("(W) %s\n", msg)
	}
}

// Info logs an information message
func (l *SLogger) Info(msg string) {
	if l.Level >= LevelInfo {
		fmt.Printf("(I) %s\n", msg)
	}
}

// Debug logs an information message
func (l *SLogger) Debug(msg string) {
	if l.Level > LevelDebug {
		fmt.Printf("(D) %s\n", msg)
	}
}

// Trace logs a function call and returns a function to be deferred marking the end of the function
func (l *SLogger) Trace(name string) func() {
	if l.Level > LevelTrace {
		fmt.Printf("[trace] %s began", name)
	}
	return func() {
		if l.Level > LevelTrace {
			fmt.Printf("[trace] %s completed", name)
		}
	}
}

// Sensitive logs a sensitive log line
func (l *SLogger) Sensitive(msg string) {
	if l.Level >= LevelSensitive {
		fmt.Printf("(S) %s\n", msg)
	}
}
