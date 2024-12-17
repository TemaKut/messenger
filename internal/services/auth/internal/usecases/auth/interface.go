package auth

import (
	"context"

	"github.com/TemaKut/messenger/pkg/service/models/proto/auth"
)

type AuthUseCase interface {
	CreateUser(ctx context.Context, req *auth.User) error
}
