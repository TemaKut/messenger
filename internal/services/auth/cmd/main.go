package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/TemaKut/messenger/internal/services/auth/cmd/factory"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	app, cleanup, err := factory.InitApp()
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
}
