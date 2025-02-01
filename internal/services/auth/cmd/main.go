package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name: "Auth",
		Action: func(cliCtx *cli.Context) error {
			ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
			defer cancel()

			// TODO factory

			<-ctx.Done()
			log.Print("received stop signal")

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("[FATAL] run app. %s", err)
	}
}
