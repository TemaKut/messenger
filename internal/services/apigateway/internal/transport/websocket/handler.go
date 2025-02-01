package websocket

import (
	"fmt"

	"github.com/TemaKut/messenger/pkg/logger"
	ws "golang.org/x/net/websocket"
)

type Handler struct {
	logger *logger.Logger
}

func NewHandler(logger *logger.Logger) *Handler {
	return &Handler{logger: logger}
}

func (h *Handler) Handle(conn *ws.Conn) {
	defer func() {
		if err := conn.Close(); err != nil {
			h.logger.Error(fmt.Sprintf("error close ws connection: %s", err))
		}

		h.logger.Info(fmt.Sprintf("ws connection %s closed", conn.RemoteAddr().String()))
	}()

	dispatcher := NewDispatcher(conn)
	defer dispatcher.Close()
}
