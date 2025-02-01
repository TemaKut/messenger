package factory

import (
	"fmt"

	"github.com/TemaKut/messenger/internal/services/auth/internal/app/config"
	"github.com/TemaKut/messenger/pkg/logger"
	"github.com/google/wire"
)

var AppSet = wire.NewSet(
	ProvideApp,
	config.NewConfig,
	ProvideLogger,
)

type App struct{}

func ProvideApp(
	l *logger.Logger,
	_ GRPCProvider,
) (*App, func()) {
	l.Info("app initialized")
	return &App{}, func() {
		l.Info("app start cleanup..")
	}
}

func ProvideLogger(cfg *config.Config) (*logger.Logger, error) {
	var env logger.LoggerEnv

	switch cfg.Environment {
	case config.EnvironmentLocal:
		env = logger.LoggerEnvLocal
	case config.EnvironmentStage:
		env = logger.LoggerEnvStage
	}

	l, err := logger.NewLogger(env)
	if err != nil {
		return nil, fmt.Errorf("error create logger: %w", err)
	}

	return l, nil
}
