package logging

import (
	"log"
)

const (
	// Debug ...
	Debug LogLevel = iota
	// Info ...
	Info
)

// LogLevel ...
type LogLevel int

// Logger ...
type Logger struct {
	Level LogLevel
}

// Info ...
func (l *Logger) Info(v ...interface{}) {
	log.Println("[INFO]", v)
}

// Warning ...
func (l *Logger) Warning(v ...interface{}) {
	log.Println("[WARNING]", v)
}

// Debug ...
func (l *Logger) Debug(v ...interface{}) {
	if l.Level > Debug {
		return
	}

	log.Println("[DEBUG]", v)
}

// Fatalf ...
func (l *Logger) Fatalf(v ...interface{}) {
	log.Fatalf("[FATAL] %s", v)
}
