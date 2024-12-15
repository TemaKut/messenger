package factory

import (
	"my_project/internal/app"

	"github.com/google/wire"
)

var AppSet = wire.NewSet(
	app.NewApp,
)
