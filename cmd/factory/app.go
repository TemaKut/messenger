package factory

import (
	"github.com/TemaKut/messenger/internal/app"

	"github.com/google/wire"
)

var AppSet = wire.NewSet(
	app.NewApp,
)
