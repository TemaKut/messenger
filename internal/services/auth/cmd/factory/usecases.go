package factory

import (
	"github.com/TemaKut/messenger/internal/services/auth/internal/usecases/auth"
	"github.com/google/wire"
)

var UseCasesSet = wire.NewSet(
	auth.NewAuthUseCase,
)
