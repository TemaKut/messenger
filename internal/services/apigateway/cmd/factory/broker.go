package factory

import (
	kafkacli "github.com/TemaKut/messenger/internal/services/apigateway/internal/clients/broker/kafka"
	kafkahandler "github.com/TemaKut/messenger/internal/services/apigateway/pkg/transport/broker/kafka"
	"github.com/TemaKut/messenger/pkg/service/events/kafka"
	"github.com/google/wire"
)

var BrokerSet = wire.NewSet(
	kafkahandler.NewServiceHandler,
	kafkacli.NewServiceConsumer,
	ProvideTopicBuilder,
)

func ProvideTopicBuilder() *kafka.TopicBuilder {
	return &kafka.TopicBuilder{}
}
