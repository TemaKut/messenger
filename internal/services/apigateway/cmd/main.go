package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/TemaKut/messenger/internal/services/apigateway/cmd/factory"
	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name: "Apigateway",
		Action: func(cliCtx *cli.Context) error {
			ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
			defer cancel()

			_, cleanup, err := factory.InitService()
			if err != nil {
				log.Fatalf("[FATAL] init service. %s", err)
			}

			defer cleanup()

			<-ctx.Done()
			log.Print("received stop signal")

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("[FATAL] run app. %s", err)
	}
}
