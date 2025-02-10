package auth

import (
	"context"

	"github.com/TemaKut/messenger/internal/services/auth/internal/dto"
)

type UserService interface {
	CreateUser(ctx context.Context, u *dto.UnregisteredUser) (*dto.User, error)
}
