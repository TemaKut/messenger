package app

import (
	"context"
	"fmt"

	"github.com/TemaKut/messenger/internal/services/auth/internal/transport/rpc"
	"golang.org/x/sync/errgroup"
)

type App struct {
	authServer *rpc.AuthServer
}

func NewApp(authServer *rpc.AuthServer) *App {
	return &App{authServer: authServer}
}

func (a *App) Run(ctx context.Context) error {
	eg, ctxErrGroup := errgroup.WithContext(ctx)

	eg.Go(
		func() error {
			if err := a.authServer.Run(); err != nil {
				return fmt.Errorf("error run auth server. %w", err)
			}

			return nil
		},
	)

	<-ctxErrGroup.Done() // TODO мутная логика по останову приложения, нужно доработать
	// TODO: log
	if err := a.Stop(); err != nil {
		return fmt.Errorf("error stop app. %w", err)
	}

	return nil
}

func (a *App) Stop() error {
	a.authServer.Stop()

	return nil
}
