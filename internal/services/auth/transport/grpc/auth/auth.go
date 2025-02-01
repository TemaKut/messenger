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

func (s *Service) InitSession(ctx context.Context, req *authv1.InitSessionRequest) (*authv1.InitSessionResponse, error) {
	return &authv1.InitSessionResponse{}, nil
}
