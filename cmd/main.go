package main

import (
	"log"

	"github.com/TemaKut/messenger/cmd/factory"
)

func main() {
	app, cleanup, err := factory.InitApp()
	if err != nil {
		log.Fatal(err)
	}

	defer cleanup()

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
