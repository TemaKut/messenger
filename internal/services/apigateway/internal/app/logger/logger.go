package logger

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/TemaKut/messenger/internal/services/apigateway/internal/app/config"
)

type Logger struct {
	*slog.Logger
}

func NewLogger(cfg *config.Config) (*Logger, error) {
	var h slog.Handler

	switch cfg.Environment {
	case config.EnvironmentLocal:
		h = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
		})
	case config.EnvironmentStage:
		h = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
		})
	default:
		return nil, fmt.Errorf("error unknown environment < %s >", cfg.Environment)
	}

	return &Logger{Logger: slog.New(h)}, nil
}
