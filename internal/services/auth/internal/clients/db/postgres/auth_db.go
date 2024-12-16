package postgres

import (
	"context"
	"fmt"

	"github.com/TemaKut/messenger/internal/services/auth/internal/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

type AuthDB = *pgxpool.Pool

func NewAuthDB(ctx context.Context, cfg *config.Config) (AuthDB, func(), error) {
	pool, err := pgxpool.Connect(ctx, cfg.GetState().Databases.AuthDb.DSN)
	if err != nil {
		return nil, nil, fmt.Errorf("error connect to db. %w", err)
	}

	return pool, pool.Close, nil
}
