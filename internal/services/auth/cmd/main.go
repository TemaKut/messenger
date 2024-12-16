package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/TemaKut/messenger/internal/services/auth/cmd/commands/migrations"
	"github.com/TemaKut/messenger/internal/services/auth/cmd/factory"
	"github.com/TemaKut/messenger/internal/services/auth/internal/config"
	"github.com/urfave/cli"
)

func main() {
	cliApp := cli.App{
		Name:     "Hello",
		Commands: []cli.Command{migrations.NewMigrationCommand()},
		Action: func(c *cli.Context) {
			ctx, stop := signal.NotifyContext(context.TODO(), os.Interrupt, os.Kill)
			defer stop()

			cfg := config.DefaultConfig() // TODO: Заменить на Consul

			app, cleanup, err := factory.InitApp(ctx, cfg)
			if err != nil {
				log.Fatalf("error init app -> %s", err)
			}

			defer cleanup()
			defer func() {
				if err := app.Stop(); err != nil {
					log.Printf("error stop app -> %s", err)
				}
			}()

			if err := app.Run(ctx); err != nil {
				log.Fatalf("error run app -> %s", err)
			}
		},
	}

	cliApp.Run(os.Args)
}
