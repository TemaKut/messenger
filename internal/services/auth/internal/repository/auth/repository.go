package auth

import (
	"context"
	"fmt"

	"github.com/TemaKut/messenger/internal/services/auth/internal/clients/db/postgres"
	authpb "github.com/TemaKut/messenger/pkg/service/models/auth"
)

type AuthRepositoryImpl struct {
	authDb postgres.AuthDB
}

func NewAuthRepository(authDb postgres.AuthDB) AuthRepository {
	return &AuthRepositoryImpl{authDb: authDb}
}

func (r *AuthRepositoryImpl) CreateUser(ctx context.Context, user *authpb.User) error {
	_, err := r.authDb.Exec(ctx, "SELECT count(1) FROM not_existing_table")
	if err != nil {
		return fmt.Errorf("error exec. %w", err)
	}

	return nil
}
