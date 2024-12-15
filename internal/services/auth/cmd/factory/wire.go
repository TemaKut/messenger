//go:build wireinject
// +build wireinject

package factory

import (
	"github.com/TemaKut/messenger/internal/services/auth/internal/app"
	"github.com/google/wire"
)

func InitApp() (*app.App, func(), error) {
	panic(
		wire.Build(
			AppSet,
			TransportSet,
			UseCasesSet,
		),
	)
}
