package migrations

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"

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
			// TODO в factory
			// TODO в миграциях сделать разбиение на файлы по DB
			cfg := config.DefaultConfig()

			db, err := sql.Open("postgres", cfg.GetState().Databases.AuthDb.DSN)
			if err != nil {
				log.Fatal(err.Error())
			}

			provider, err := goose.NewProvider(goose.DialectPostgres, db, migrations.EmbedMigrations)
			if err != nil {
				log.Fatal(err.Error())
			}

			switch c.Args().Get(0) {
			case "up":
				mr, err := provider.Up(context.TODO())
				if err != nil {
					log.Fatal(err.Error())
				}

				for _, result := range mr {
					fmt.Println(result.String())
				}
			case "down-to":
				version, err := strconv.Atoi(c.Args().Get(1))
				if err != nil {
					log.Fatalf("error parse down-to version. %s", err)
				}

				mr, err := provider.DownTo(context.TODO(), int64(version))
				if err != nil {
					log.Fatal(err.Error())
				}

				for _, result := range mr {
					fmt.Println(result.String())
				}
			}
		},
	}

	return cmd
}
