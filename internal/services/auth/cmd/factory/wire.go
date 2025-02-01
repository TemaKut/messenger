//go:build wireinject

package factory

import "github.com/google/wire"

func InitService() (*App, func(), error) {
	panic(wire.Build(
		AppSet,
		GRPCSet,
	))
}
