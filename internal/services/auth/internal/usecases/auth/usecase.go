package auth

import (
	"context"
	"fmt"

	authrepo "github.com/TemaKut/messenger/internal/services/auth/internal/repository/auth"
	authpb "github.com/TemaKut/messenger/pkg/service/models/proto/auth"
	"github.com/google/uuid"
)

type AuthUseCaseImpl struct {
	authRepository authrepo.AuthRepository
}

func NewAuthUseCase(authRepository authrepo.AuthRepository) AuthUseCase {

	return &AuthUseCaseImpl{authRepository: authRepository}
}

func (u *AuthUseCaseImpl) CreateUser(ctx context.Context, user *authpb.User) error {
	if user.Id == "" {
		user.Id = uuid.NewString()
	}

	if err := u.authRepository.AddUser(ctx, user); err != nil {
		return fmt.Errorf("error create user. %w", err)
	}

	return nil
}
