package factory

import (
	"github.com/TemaKut/messenger/internal/services/auth/internal/transport/rpc"
	"github.com/google/wire"
)

var TransportSet = wire.NewSet(
	rpc.NewAuthServer,
	rpc.NewAuthService,
)
