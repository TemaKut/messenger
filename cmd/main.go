package main

import (
	"log"
	"my_project/cmd/factory"
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
