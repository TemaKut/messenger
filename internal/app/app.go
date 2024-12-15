package app

import (
	"fmt"
	"time"
)

type App struct {
}

func NewApp() (*App, func()) {
	app := &App{}
	return app, app.Stop
}

func (a *App) Run() error {
	fmt.Println("Running")
	time.Sleep(10 * time.Second)

	return nil
}

func (a *App) Stop() {
	fmt.Println("Stopping")
}
