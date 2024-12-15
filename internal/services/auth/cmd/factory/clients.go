package factory

import (
	"github.com/TemaKut/messenger/internal/services/auth/internal/clients/db/postgres"
	"github.com/google/wire"
)

var ClientsSet = wire.NewSet(
	postgres.NewAuthDB,
)
