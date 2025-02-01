package websocket

import (
	"fmt"

	"github.com/TemaKut/messenger/pkg/logger"
	authv1 "github.com/TemaKut/messenger/pkg/proto/service/gen/auth"
	ws "golang.org/x/net/websocket"
)

type Handler struct {
	logger      *logger.Logger
	authService authv1.AuthServiceClient
}

func NewHandler(logger *logger.Logger, authService authv1.AuthServiceClient) *Handler {
	return &Handler{logger: logger, authService: authService}
}

func (h *Handler) Handle(conn *ws.Conn) {
	defer func() {
		if err := conn.Close(); err != nil {
			h.logger.Error(fmt.Sprintf("error close ws connection: %s", err))
		}

		h.logger.Info(fmt.Sprintf("ws connection %s closed", conn.RemoteAddr().String()))
	}()

	// TODO продумать как организовать систему коннектов и реквестов с респонсами

	for {
		var msg string

		if err := ws.Message.Receive(conn, &msg); err != nil {
			h.logger.Error(fmt.Sprintf("error receive msg. %s", err))
		}

		fmt.Println(h.authService.InitSession(conn.Request().Context(), &authv1.InitSessionRequest{}))
	}
}
