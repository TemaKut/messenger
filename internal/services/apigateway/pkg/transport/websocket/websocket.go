package websocket

import (
	"context"
	"fmt"
	"net/http"

	"github.com/TemaKut/messenger/internal/services/apigateway/internal/config"
	"github.com/TemaKut/messenger/internal/services/apigateway/internal/logger"
	gws "github.com/gorilla/websocket"
)

type WebsocketServer struct {
	cfg      *config.Config
	upgrader gws.Upgrader
	log      logger.Logger
}

func NewWebsocketServer(cfg *config.Config, log logger.Logger) *WebsocketServer {
	upgrader := gws.Upgrader{
		ReadBufferSize:  4096,
		WriteBufferSize: 4096,
	}

	return &WebsocketServer{upgrader: upgrader, cfg: cfg, log: log}
}

func (w *WebsocketServer) Start(ctx context.Context) error {
	w.log.Info(fmt.Sprintf("start websocket server on <%s>", w.cfg.GetState().Transport.Websocket.Addr))

	server := &http.Server{
		Addr:    w.cfg.GetState().Transport.Websocket.Addr,
		Handler: http.HandlerFunc(w.handleWebsocket),
	}
	errCh := make(chan error)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			errCh <- fmt.Errorf("error listen and serve. %w", err)
		}
	}()

	defer func() {
		if err := server.Shutdown(context.Background()); err != nil {
			w.log.Error("error shutdown websocket server")
		}
	}()

	select {
	case <-errCh:
	case <-ctx.Done():
	}

	return nil
}
func (w *WebsocketServer) handleWebsocket(wr http.ResponseWriter, req *http.Request) {
	ws, err := w.upgrader.Upgrade(wr, req, nil)
	if err != nil {
		fmt.Printf("error ws -> %s\n", err)
		return
	}

	defer ws.Close()

	for {
		_, b, err := ws.ReadMessage()
		if err != nil {
			w.log.Error(fmt.Sprintf("error read message %s", err))
			break
		}

		fmt.Printf("%s -> %s\n", ws.RemoteAddr(), string(b))

		if err := ws.WriteMessage(gws.BinaryMessage, b); err != nil {
			w.log.Error(fmt.Sprintf("error write message %s", err))
			break
		}
	}
}
