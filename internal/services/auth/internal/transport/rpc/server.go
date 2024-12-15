package rpc

import (
	"fmt"
	"net"

	"github.com/TemaKut/messenger/internal/services/auth/internal/config"
	authsrv "github.com/TemaKut/messenger/internal/services/auth/internal/transport/rpc/auth"
	"github.com/TemaKut/messenger/pkg/service/models/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type AuthServer struct {
	srv *grpc.Server
	cfg *config.Config
}

func NewAuthServer(s *authsrv.AuthService, cfg *config.Config) *AuthServer {
	srv := grpc.NewServer()
	auth.RegisterAuthServiceServer(srv, s)
	reflection.Register(srv)

	return &AuthServer{srv: srv, cfg: cfg}
}

func (s *AuthServer) Run() error {
	l, err := net.Listen("tcp", s.cfg.Transport.Rpc.Addres)
	if err != nil {
		return fmt.Errorf("error listen addres (%s). %w", s.cfg.Transport.Rpc.Addres, err)
	}

	if err := s.srv.Serve(l); err != nil {
		return fmt.Errorf("error serve grpc. %w", err)
	}

	return nil
}

func (s *AuthServer) Stop() {
	s.srv.GracefulStop()
}
