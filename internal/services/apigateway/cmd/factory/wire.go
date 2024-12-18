//go:build wireinject
// +build wireinject

package factory

import (
	"context"

	"github.com/TemaKut/messenger/internal/services/apigateway/internal/app"
	"github.com/TemaKut/messenger/internal/services/apigateway/internal/config"
	"github.com/google/wire"
)

func InitApp(ctx context.Context, cfg *config.Config) (*app.App, func(), error) {
	panic(
		wire.Build(
			AppSet,
		),
	)
}
