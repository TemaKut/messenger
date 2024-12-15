package factory

import (
	"github.com/TemaKut/messenger/internal/services/auth/internal/usecases/users"
	"github.com/google/wire"
)

var UseCasesSet = wire.NewSet(
	users.NewUsersUseCase,
)
