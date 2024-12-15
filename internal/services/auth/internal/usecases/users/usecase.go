package users

import (
	"context"

	"github.com/TemaKut/messenger/pkg/service/models/auth"
)

type UsersUseCaseImpl struct {
}

func NewUsersUseCase() UsersUseCase {
	return &UsersUseCaseImpl{}
}

func (u *UsersUseCaseImpl) CreateUser(ctx context.Context, req *auth.CreateUserRequest) error {

	return nil
}
