package auth

import (
	"context"
	"fmt"

	authuc "github.com/TemaKut/messenger/internal/services/auth/internal/usecases/auth"
	"github.com/TemaKut/messenger/pkg/service/models/auth"
)

type AuthService struct {
	auth.UnimplementedAuthServiceServer

	usersUseCase authuc.AuthUseCase
}

func NewAuthService(usersUseCase authuc.AuthUseCase) *AuthService {
	return &AuthService{usersUseCase: usersUseCase}
}

func (a *AuthService) CreateUser(ctx context.Context, req *auth.CreateUserRequest) (*auth.CreateUserResponse, error) {
	if err := a.usersUseCase.CreateUser(ctx, req.GetUser()); err != nil {
		return nil, fmt.Errorf("error create user. %w", err)
	}

	return &auth.CreateUserResponse{}, nil
}
