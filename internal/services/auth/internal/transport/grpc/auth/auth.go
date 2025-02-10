package auth

import (
	"context"
	"fmt"

	"github.com/TemaKut/messenger/internal/services/auth/internal/dto/mappers"
	authv1 "github.com/TemaKut/messenger/pkg/proto/service/gen/auth"
)

type Service struct {
	authv1.UnimplementedAuthServiceServer

	userService UserService // TODO прокинуть сервис
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) InitSession(
	ctx context.Context,
	req *authv1.InitSessionRequest,
) (*authv1.InitSessionResponse, error) {
	// TODO implement the method
	return &authv1.InitSessionResponse{}, nil
}

func (s *Service) RegisterUser(
	ctx context.Context,
	req *authv1.RegisterUserRequest,
) (*authv1.RegisterUserResponse, error) {
	user, err := s.userService.CreateUser(ctx, mappers.UnregisteredUserFromRegisterUserRequest(req))
	if err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}

	return &authv1.RegisterUserResponse{User: mappers.UserToProto(user)}, nil
}
