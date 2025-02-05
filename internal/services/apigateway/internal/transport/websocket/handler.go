package websocket

import (
	"fmt"

	wsm "github.com/TemaKut/messenger/internal/services/apigateway/internal/sessions/websocket"
	"github.com/TemaKut/messenger/pkg/logger"
	ws "golang.org/x/net/websocket"
	"golang.org/x/sync/errgroup"
)

type Handler struct {
	logger         *logger.Logger
	sessionManager *wsm.SessionManager
}

func NewHandler(logger *logger.Logger, sessionManager *wsm.SessionManager) *Handler {
	return &Handler{
		logger:         logger,
		sessionManager: sessionManager,
	}
}

func (h *Handler) HandleConnection(conn *ws.Conn) {
	defer conn.Close()

	sessId := h.sessionManager.RegisterSessionFromConnection(conn)

	eg, ctx := errgroup.WithContext(conn.Request().Context())

	eg.Go(func() error {
		if err := h.sessionManager.HandleSessionRequests(ctx, sessId); err != nil {
			return fmt.Errorf("error handle session requests. %w", err)
		}

		return nil
	})

	if err := eg.Wait(); err != nil {
		h.logger.Error(fmt.Sprintf("error wait wait group. %s", err))
	}
}
