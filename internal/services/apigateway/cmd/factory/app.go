package factory

import (
	"github.com/TemaKut/messenger/internal/services/apigateway/internal/app/config"
	"github.com/TemaKut/messenger/internal/services/apigateway/internal/app/logger"
	"github.com/google/wire"
)

var AppSet = wire.NewSet(
	ProvideApp,
	config.NewConfig,
	logger.NewLogger,
)

type App struct{}

func ProvideApp(
	l *logger.Logger,
	_ HttpServerProvider,
) (*App, func()) {
	l.Info("app initialized")
	return &App{}, func() {
		l.Info("app start cleanup..")
	}
}
