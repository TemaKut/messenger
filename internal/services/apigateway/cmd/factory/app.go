package factory

import (
	"github.com/TemaKut/messenger/internal/services/apigateway/internal/app"
	"github.com/TemaKut/messenger/internal/services/apigateway/internal/logger"
	"github.com/google/wire"
)

var AppSet = wire.NewSet(
	app.NewApp,
	logger.NewLogger,

// logger.NewLogger,
)
