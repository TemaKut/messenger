package rpc

import (
	"context"
	"fmt"

	"github.com/TemaKut/messenger/internal/services/auth/internal/usecases/users"
	"github.com/TemaKut/messenger/pkg/service/models/auth"
)

type AuthService struct {
	auth.UnimplementedAuthServiceServer

	usersUseCase users.UsersUseCase
}

func NewAuthService(usersUseCase users.UsersUseCase) *AuthService {
	return &AuthService{usersUseCase: usersUseCase}
}

func (a *AuthService) CreateUser(ctx context.Context, req *auth.CreateUserRequest) (*auth.CreateUserResponse, error) {
	fmt.Println("!!!")
	return nil, nil
}
