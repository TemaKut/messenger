package app

import (
	"context"
	"fmt"

	"github.com/TemaKut/messenger/internal/services/apigateway/internal/clients/broker/kafka"
	"github.com/TemaKut/messenger/internal/services/apigateway/internal/logger"
)

type App struct {
	serviceConsumer *kafka.ServiceConsumer
	log             logger.Logger
}

func NewApp(serviceConsumer *kafka.ServiceConsumer, log logger.Logger) *App {
	return &App{serviceConsumer: serviceConsumer, log: log}
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

	select {
	case <-ctx.Done():
	case e := <-errCh:
		a.log.Error(e.Error())
	}

	return nil
}

func (a *App) Stop() error {
	if err := a.serviceConsumer.Close(); err != nil {
		return fmt.Errorf("error close service consumer. %w", err)
	}

	return nil
}
