package app

import (
	"context"
	"fmt"

	"github.com/TemaKut/messenger/internal/services/apigateway/internal/clients/broker/kafka"
	"github.com/TemaKut/messenger/internal/services/apigateway/internal/logger"
	"github.com/TemaKut/messenger/internal/services/apigateway/pkg/transport/websocket"
)

type App struct {
	serviceConsumer *kafka.ServiceConsumer
	websocketServer *websocket.WebsocketServer
	log             logger.Logger
}

func NewApp(
	websocketServer *websocket.WebsocketServer,
	serviceConsumer *kafka.ServiceConsumer,
	log logger.Logger,
) *App {
	return &App{
		serviceConsumer: serviceConsumer,
		websocketServer: websocketServer,
		log:             log,
	}
}

func (a *App) Run(ctx context.Context) error {
	a.log.Info("starting app..")

	errCh := make(chan error, 1)

	// TODO доработать работу с запуском
	go func() {
		if err := a.serviceConsumer.Consume(ctx); err != nil {
			errCh <- fmt.Errorf("error consume. %w", err)
		}
	}()

	go func() {
		if err := a.websocketServer.Start(ctx); err != nil {
			errCh <- fmt.Errorf("error start websocket service. %w", err)
		}
	}()

	select {
	case <-ctx.Done():
	case e := <-errCh:
		a.log.Error(e.Error())
	}

	a.stop()

	return nil
}

func (a *App) stop() {
	if err := a.serviceConsumer.Close(); err != nil {
		a.log.Error(fmt.Errorf("error close service consumer. %w", err).Error())
	}
}
