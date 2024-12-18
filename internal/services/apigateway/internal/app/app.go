package app

import "context"

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run(ctx context.Context) error {

	return nil
}

func (a *App) Stop() error {

	return nil
}
