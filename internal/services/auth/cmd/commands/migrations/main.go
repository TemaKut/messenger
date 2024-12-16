package migrations

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/TemaKut/messenger/internal/services/auth/internal/config"
	"github.com/TemaKut/messenger/internal/services/auth/internal/migrations"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/urfave/cli"
)

func NewMigrationCommand() cli.Command {
	cmd := cli.Command{
		Name: "migrations",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name: "how",
			},
		},
		Action: func(c *cli.Context) {
			cfg := config.DefaultConfig()

			db, err := sql.Open("postgres", cfg.GetState().Databases.AuthDb.DSN)
			if err != nil {
				log.Fatal(err.Error())
			}

			provider, err := goose.NewProvider(goose.DialectPostgres, db, migrations.EmbedMigrations)
			if err != nil {
				log.Fatal(err.Error())
			}

			mr, err := provider.DownTo(context.TODO(), 00)
			if err != nil {
				log.Fatal(err.Error())
			}

			for _, result := range mr {
				fmt.Println(result.String())
			}
		},
	}

	return cmd
}
