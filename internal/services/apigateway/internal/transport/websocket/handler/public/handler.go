package public

import (
	"fmt"

	"github.com/TemaKut/messenger/internal/services/apigateway/internal/app/logger"
	"golang.org/x/net/websocket"
)

type Handler struct {
	logger *logger.Logger
}

func NewHandler(logger *logger.Logger) *Handler {
	return &Handler{logger: logger}
}

func (h *Handler) Handle(ws *websocket.Conn) {
	defer h.closeWebsocketConnection(ws)

}

func (h *Handler) closeWebsocketConnection(ws *websocket.Conn) {
	if err := ws.Close(); err != nil {
		h.logger.Error(fmt.Sprintf("error close ws connection: %s", err))
	}
}
