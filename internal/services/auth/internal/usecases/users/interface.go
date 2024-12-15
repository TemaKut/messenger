package users

import (
	"context"

	"github.com/TemaKut/messenger/pkg/service/models/auth"
)

type UsersUseCase interface {
	CreateUser(ctx context.Context, req *auth.CreateUserRequest) error
}
