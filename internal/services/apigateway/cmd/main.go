package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/TemaKut/messenger/internal/services/apigateway/cmd/factory"
	"github.com/TemaKut/messenger/internal/services/apigateway/internal/config"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
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
			// TODO consul
			consulDefCfg := api.DefaultConfig()
			consul, err := api.NewClient(consulDefCfg)
			if err != nil {
				panic(fmt.Sprintf("panic make consul client. %s", err))
			}

			serviceId := uuid.New().String()
			checkId := uuid.New().String()

			service := &api.AgentServiceRegistration{
				ID:      serviceId,
				Name:    AppName,
				Address: "localhost",
				Tags:    []string{"api", "apigateway"},
			}

			// service.Check = &api.AgentServiceCheck{
			// 	CheckID:                        checkId,
			// 	DeregisterCriticalServiceAfter: "20s",
			// }

			if err := consul.Agent().ServiceRegister(service); err != nil {
				panic(fmt.Sprintf("panic register service. %s", err))
			}

			defer consul.Agent().ServiceDeregister(serviceId)
			defer consul.Agent().CheckDeregister(checkId)

			srv, _, err := consul.Agent().Service(serviceId, nil)
			if err != nil {
				panic(fmt.Sprintf("panic fetch service %s. %s", serviceId, err))
			}

			fmt.Printf("%+v", srv)

			cfg := config.DefaultConfig() // TODO: Заменить на сервисный

			app, cleanup, err := factory.InitApp(ctx, cfg)
			if err != nil {
				log.Fatalf("error init app -> %s", err)
			}

			defer cleanup()

			if err := app.Run(ctx); err != nil {
				log.Fatalf("error run app -> %s", err)
			}

			defer app.Stop()
		},
	}

	cliApp.Run(os.Args)
}
