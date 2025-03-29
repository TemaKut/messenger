//go:build wireinject
// +build wireinject

package factory

import "github.com/google/wire"

func InitApp() (*App, func(), error) {
	panic(
		wire.Build(
			AppSet,
		),
	)
}
