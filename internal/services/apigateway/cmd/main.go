package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/TemaKut/messenger/internal/services/apigateway/cmd/factory"
	"github.com/TemaKut/messenger/internal/services/apigateway/internal/config"
	"github.com/urfave/cli"
)

const (
	AppName = "ApiGateway"
)

func main() {
	cliApp := cli.App{
		Name: AppName,
		Action: func(c *cli.Context) {
			ctx, stop := signal.NotifyContext(context.TODO(), os.Interrupt, os.Kill)
			defer stop()

			cfg := config.DefaultConfig() // TODO: Заменить на сервисный

			app, cleanup, err := factory.InitApp(ctx, cfg)
			if err != nil {
				log.Fatalf("error init app -> %s", err)
			}

			defer cleanup()

			if err := app.Run(ctx); err != nil {
				log.Fatalf("error run app -> %s", err)
			}
		},
	}

	cliApp.Run(os.Args)
}
