package app

import (
	"context"
	"fmt"

	"github.com/TemaKut/messenger/internal/services/auth/internal/logger"
	"github.com/TemaKut/messenger/internal/services/auth/internal/transport/rpc"
)

type App struct {
	authServer *rpc.AuthServer
	log        logger.Logger
}

func NewApp(authServer *rpc.AuthServer, log logger.Logger) *App {
	return &App{authServer: authServer, log: log}
}

func (a *App) Run(ctx context.Context) error {
	a.log.Info("app running..")

	errCh := make(chan error, 1)
	// TODO доработать работу с запуском
	go func() {
		if err := a.authServer.Run(); err != nil {
			errCh <- fmt.Errorf("error run auth server. %w", err)
		}
	}()

	select {
	case <-ctx.Done():
	case e := <-errCh:
		a.log.Error(e.Error())
	}

	if err := a.Stop(); err != nil {
		return fmt.Errorf("error stop app. %w", err)
	}

	return nil
}

func (a *App) Stop() error {
	a.log.Info("stopping app..")

	a.authServer.Stop()

	return nil
}
