package factory

import (
	"context"
	"fmt"
	"time"

	"github.com/TemaKut/messenger/internal/services/apigateway/internal/app/config"
	"github.com/TemaKut/messenger/internal/services/apigateway/internal/app/logger"
	"github.com/TemaKut/messenger/internal/services/apigateway/internal/transport/websocket/handler/public"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/websocket"
)

var HttpSet = wire.NewSet(
	ProvideHttpServer,
	public.NewHandler,
)

type HttpServerProvider struct{}

func ProvideHttpServer(
	cfg *config.Config,
	logger *logger.Logger,
	webSocketHandler *public.Handler,
) (HttpServerProvider, func(), error) {
	logger.Info("http server running..")

	e := echo.New()

	e.Any("/ws", func(c echo.Context) error {
		wsh := websocket.Handler(webSocketHandler.Handle)
		wsh.ServeHTTP(c.Response(), c.Request())

		return nil
	})

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	runServerErrCh := make(chan error, 1)

	go func() {
		if err := e.Start(cfg.Server.Websocket.Addr); err != nil {
			runServerErrCh <- fmt.Errorf("error start websocket server on %s. %w", cfg.Server.Websocket.Addr, err)
		}
	}()

	afterCh := time.After(1 * time.Second)

	select {
	case err := <-runServerErrCh:
		if err != nil {
			return HttpServerProvider{}, nil, err
		}
	case <-afterCh:
	}

	return HttpServerProvider{}, func() {
		defer close(runServerErrCh)
		logger.Info("shut down websocket server..")

		if err := e.Shutdown(context.TODO()); err != nil {
			logger.Error("error close webscoket server %s. %w", cfg.Server.Websocket.Addr, err)
		}
	}, nil
}
