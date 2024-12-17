package auth

import (
	"context"

	authpb "github.com/TemaKut/messenger/pkg/service/models/proto/auth"
)

type AuthRepository interface {
	AddUser(ctx context.Context, user *authpb.User) error
}
