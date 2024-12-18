package logger

import (
	"errors"
	"log/slog"
	"os"

	"github.com/TemaKut/messenger/internal/services/apigateway/internal/config"
)

type Logger = *slog.Logger

func NewLogger(cfg *config.Config) (Logger, error) {
	env := cfg.GetState().Environment

	var h slog.Handler

	switch env {
	case config.EnvironmentLocal:
		h = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelInfo,
			AddSource: true,
		})
	// case config.EnvironmentProd:
	// case config.EnvironmentStage:
	default:
		return nil, errors.New("error unsupported env for logger handler")
	}

	return slog.New(h), nil
}
