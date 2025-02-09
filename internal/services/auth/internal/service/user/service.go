package user

import (
	"context"

	"github.com/TemaKut/messenger/internal/services/auth/internal/dto"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func CreateUser(ctx context.Context, u *dto.User) (*dto.User, error) {
	return u, nil // FIXME
}
