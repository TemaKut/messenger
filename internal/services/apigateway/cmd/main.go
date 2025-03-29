package main

import (
	"fmt"
	"log"
	"os"

	"github.com/TemaKut/messenger/internal/services/apigateway/cmd/factory"
	"github.com/urfave/cli"
)

func main() {
	app := cli.App{
		Action: func(c *cli.Context) error {
			_, cleanup, err := factory.InitApp()
			if err != nil {
				return fmt.Errorf("error init app. %w", err)
			}

			defer cleanup()

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("fatal run app. %s", err)
	}
}
