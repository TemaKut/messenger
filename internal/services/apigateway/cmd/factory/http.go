package factory

import (
	"github.com/TemaKut/messenger/internal/services/apigateway/pkg/transport/websocket"
	"github.com/google/wire"
)

var HttpSet = wire.NewSet(
	websocket.NewWebsocketServer,
)
