package factory

import (
	"fmt"

	"github.com/TemaKut/messenger/internal/services/apigateway/internal/app/config"
	"github.com/TemaKut/messenger/pkg/logger"
	authv1 "github.com/TemaKut/messenger/pkg/proto/service/gen/auth"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var GRPCSet = wire.NewSet(
	ProvideAuthServiceGRPCClient,
)

func ProvideAuthServiceGRPCClient(l *logger.Logger, cfg *config.Config) (authv1.AuthServiceClient, func(), error) {
	l.Info(fmt.Sprintf("connecting to auth service %s", cfg.Clients.AuthService.Addr))

	conn, err := grpc.NewClient(
		cfg.Clients.AuthService.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("error make grpc client. %w", err)
	}

	return authv1.NewAuthServiceClient(conn), func() {
		if err := conn.Close(); err != nil {
			l.Error(fmt.Sprintf("error close connection %s. %s", cfg.Clients.AuthService.Addr, err))

			return
		}

		l.Info(fmt.Sprintf("connection to %s closed", cfg.Clients.AuthService.Addr))
	}, nil
}
