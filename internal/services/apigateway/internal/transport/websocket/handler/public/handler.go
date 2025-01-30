package public

import (
	"fmt"
	"time"

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

	var handleErr error

	for {
		// TODO chans for receive and send
		if err := websocket.Message.Send(ws, "Hello))"); err != nil {
			handleErr = fmt.Errorf("error send message. %w", err)

			break
		}
		time.Sleep(time.Second)
	}

	if handleErr != nil {
		h.logger.Error(fmt.Sprintf("error handle connection. %s", handleErr))
	}
}

func (h *Handler) closeWebsocketConnection(ws *websocket.Conn) {
	if err := ws.Close(); err != nil {
		h.logger.Error(fmt.Sprintf("error close ws connection: %s", err))
	}
}
