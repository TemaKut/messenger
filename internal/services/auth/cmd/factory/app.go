package factory

import (
	"github.com/TemaKut/messenger/internal/services/auth/internal/app"
	"github.com/TemaKut/messenger/internal/services/auth/internal/logger"
	"github.com/google/wire"
)

var AppSet = wire.NewSet(
	app.NewApp,
	logger.NewLogger,
)
