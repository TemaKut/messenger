package factory

import (
	"github.com/TemaKut/messenger/internal/services/auth/pkg/transport/rpc"
	"github.com/TemaKut/messenger/internal/services/auth/pkg/transport/rpc/auth"
	"github.com/google/wire"
)

var TransportSet = wire.NewSet(
	rpc.NewAuthServer,
	auth.NewAuthService,
)
