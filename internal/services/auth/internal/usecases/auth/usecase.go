package auth

import (
	"context"
	"fmt"

	authrepo "github.com/TemaKut/messenger/internal/services/auth/internal/repository/auth"
	authpb "github.com/TemaKut/messenger/pkg/service/models/auth"
)

type AuthUseCaseImpl struct {
	authRepository authrepo.AuthRepository
}

func NewAuthUseCase(authRepository authrepo.AuthRepository) AuthUseCase {

	return &AuthUseCaseImpl{authRepository: authRepository}
}

func (u *AuthUseCaseImpl) CreateUser(ctx context.Context, user *authpb.User) error {
	if err := u.authRepository.CreateUser(ctx, user); err != nil {
		return fmt.Errorf("error create user. %w", err)
	}

	return nil
}
