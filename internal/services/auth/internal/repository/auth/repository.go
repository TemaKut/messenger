package auth

import (
	"context"
	"fmt"

	"github.com/TemaKut/messenger/internal/services/auth/internal/clients/db/postgres"
	authpb "github.com/TemaKut/messenger/pkg/service/models/proto/auth"
	"google.golang.org/protobuf/encoding/protojson"
)

type AuthRepositoryImpl struct {
	authDb postgres.AuthDB
}

func NewAuthRepository(authDb postgres.AuthDB) AuthRepository {
	return &AuthRepositoryImpl{authDb: authDb}
}

func (r *AuthRepositoryImpl) AddUser(ctx context.Context, user *authpb.User) error {
	propsBytes, err := protojson.Marshal(user.Props)
	if err != nil {
		return fmt.Errorf("error marshal user props. %w", err)
	}

	_, err = r.authDb.Exec(ctx,
		`INSERT INTO users(id, first_name, last_name, phone, props, created_at, updated_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7)
		`,
		user.Id, user.FirstName, user.LastName, user.Phone, propsBytes, user.CreatedAt.AsTime(), user.UpdatedAt.AsTime(),
	)
	if err != nil {
		return fmt.Errorf("error insert user. %w", err)
	}

	return nil
}
