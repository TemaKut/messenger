package logger

import (
	"fmt"
	"log/slog"
	"os"
)

type Logger struct {
	*slog.Logger
}

type LoggerEnv int

const (
	LoggerEnvLocal LoggerEnv = iota
	LoggerEnvStage
)

// TODO вынести в pkg прокета
func NewLogger(env LoggerEnv) (*Logger, error) {
	var h slog.Handler

	switch env {
	case LoggerEnvLocal:
		h = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
		})
	case LoggerEnvStage:
		h = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
		})
	default:
		return nil, fmt.Errorf("error unknown environment")
	}

	return &Logger{Logger: slog.New(h)}, nil
}
