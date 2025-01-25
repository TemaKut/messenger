//go:build wireinject

package factory

import "github.com/google/wire"

func InitService() (*Service, func(), error) {
	panic(wire.Build(
		ProvideService,
	))
}

type Service struct{}

func ProvideService() (*Service, func()) {
	return &Service{}, func() {

	}
}
