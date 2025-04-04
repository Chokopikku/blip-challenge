package utils

import (
	"errors"
	"log"
	"os"
)

// Log levels
const (
	LevelInfo  = "INFO"
	LevelDebug = "DEBUG"
	LevelWarn  = "WARN"
	LevelError = "ERROR"
)

// CustomLogger is a log.Logger wrapper, with levels, that outputs logs to a file.
type CustomLogger struct {
	logger  *log.Logger
	logFile *os.File
}

func NewLogger(filePath string) (*CustomLogger, error) {
	if filePath == "" {
		return nil, errors.New("log file path must not be empty")
	}

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	logger := log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
	return &CustomLogger{logger: logger, logFile: file}, nil
}

func (l *CustomLogger) Info(message string) {
	l.log(LevelInfo, message)
}

func (l *CustomLogger) Debug(message string) {
	l.log(LevelDebug, message)
}

func (l *CustomLogger) Warn(message string) {
	l.log(LevelWarn, message)
}

func (l *CustomLogger) Error(message string) {
	l.log(LevelError, message)
}

func (l *CustomLogger) log(level string, message string) {
	l.logger.Printf("[%s] %s", level, message)
}

func (l *CustomLogger) Close() {
	if l.logFile != nil {
		_ = l.logFile.Close()
	}
}
