package rpc

import (
	"fmt"
	"net"

	"github.com/TemaKut/messenger/pkg/service/models/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type AuthServer struct {
	srv *grpc.Server
}

func NewAuthServer(s *AuthService) *AuthServer {
	srv := grpc.NewServer()
	auth.RegisterAuthServiceServer(srv, s)
	reflection.Register(srv)

	return &AuthServer{srv: srv}
}

func (s *AuthServer) Run() error {
	addres := ":8001"

	l, err := net.Listen("tcp", addres)
	if err != nil {
		return fmt.Errorf("error listen addres (%s). %w", addres, err)
	}

	if err := s.srv.Serve(l); err != nil {
		return fmt.Errorf("error serve grpc. %w", err)
	}

	return nil
}

func (s *AuthServer) Stop() {
	s.srv.GracefulStop()
}
