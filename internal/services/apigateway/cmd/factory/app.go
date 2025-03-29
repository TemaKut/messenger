package factory

import "github.com/google/wire"

var AppSet = wire.NewSet(
	ProvideApp,
)

type App struct{}

func ProvideApp() (*App, func(), error) {
	return &App{}, func() {}, nil
}
