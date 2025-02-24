package logger

import (
	"log/slog"
	"os"
)

type Service struct {
	logger *slog.Logger
}

// InitLogger initializes the logger.
func NewLogger(debug bool) *Service {
	var logHandler slog.Handler
	if debug {
		logHandler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource:   true,
			Level:       slog.LevelInfo,
			ReplaceAttr: nil,
		})
	} else {
		logHandler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			AddSource:   true,
			Level:       slog.LevelInfo,
			ReplaceAttr: nil,
		})
	}

	return &Service{
		logger: slog.New(logHandler),
	}
}

// Debug logs a debug message.
func (l *Service) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

// Info logs an informational message.
func (l *Service) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

// Warn logs a warning message.
func (l *Service) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}

// Error logs an error message.
func (l *Service) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}
