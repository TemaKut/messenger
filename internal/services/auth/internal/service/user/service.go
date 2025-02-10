package user

import (
	"context"

	"github.com/TemaKut/messenger/internal/services/auth/internal/dto"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

type CreateUserParams struct {
}

func CreateUser(ctx context.Context, u *dto.UnregisteredUser) (*dto.User, error) {
	return nil, nil // FIXME
}
