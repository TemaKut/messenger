package auth

import (
	"context"

	authv1 "github.com/TemaKut/messenger/pkg/proto/service/gen/auth"
)

type Service struct {
	authv1.UnimplementedAuthServiceServer
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
	return &authv1.RegisterUserResponse{}, nil
}
