package factory

import (
	"github.com/TemaKut/messenger/internal/services/auth/internal/repository/auth"
	"github.com/google/wire"
)

var RepositorySet = wire.NewSet(
	auth.NewAuthRepository,
)
