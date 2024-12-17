package factory

import (
	"github.com/TemaKut/messenger/internal/services/auth/internal/clients/broker/kafka"
	"github.com/google/wire"
)

var BrokerSet = wire.NewSet(
	kafka.NewProducer,
)
