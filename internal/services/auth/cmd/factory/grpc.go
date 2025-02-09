package factory

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/TemaKut/messenger/internal/services/auth/internal/app/config"
	"github.com/TemaKut/messenger/internal/services/auth/internal/transport/grpc/auth"
	"github.com/TemaKut/messenger/pkg/logger"
	authv1 "github.com/TemaKut/messenger/pkg/proto/service/gen/auth"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var GRPCSet = wire.NewSet(
	ProvideGRPC,
	ProvideGRPCServer,
	auth.NewService,
)

type GRPCProvider struct{}

func ProvideGRPC(_ *grpc.Server) GRPCProvider {
	return GRPCProvider{}
}

func ProvideGRPCServer(
	l *logger.Logger,
	cfg *config.Config,
	authService *auth.Service,
) (*grpc.Server, func(), error) {
	l.Info(fmt.Sprintf("starting grpc server on %s", cfg.Server.GRPC.Addr))

	lis, err := net.Listen("tcp", cfg.Server.GRPC.Addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()

	authv1.RegisterAuthServiceServer(srv, authService)
	reflection.Register(srv)

	errCh := make(chan error, 1)

	go func() {
		if err := srv.Serve(lis); err != nil {
			errCh <- fmt.Errorf("failed to serve %s %w", cfg.Server.GRPC.Addr, err)
		}
	}()

	after := time.After(1 * time.Second)

	select {
	case err := <-errCh:
		return nil, nil, err
	case <-after:
	}

	return srv, func() {
		defer close(errCh)

		l.Info(fmt.Sprintf("stopping grpc server %s", cfg.Server.GRPC.Addr))
		srv.GracefulStop()
	}, nil
}
