package auth

import (
	"context"

	authpb "github.com/TemaKut/messenger/pkg/service/models/auth"
)

type AuthRepository interface {
	CreateUser(ctx context.Context, user *authpb.User) error
}
