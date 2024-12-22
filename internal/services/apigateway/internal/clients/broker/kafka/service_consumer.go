package kafka

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/TemaKut/messenger/internal/services/apigateway/internal/config"
	kafkahandler "github.com/TemaKut/messenger/internal/services/apigateway/pkg/transport/broker/kafka"
	"github.com/TemaKut/messenger/pkg/service/events/kafka"
)

type ServiceConsumer struct {
	cg             sarama.ConsumerGroup
	serviceHandler *kafkahandler.ServiceHandler
	topicBuilder   *kafka.TopicBuilder
}

func NewServiceConsumer(
	cfg *config.Config,
	handler *kafkahandler.ServiceHandler,
	topicBuilder *kafka.TopicBuilder,
) (*ServiceConsumer, error) {
	c := sarama.NewConfig()
	c.Consumer.Offsets.Initial = sarama.OffsetOldest

	cg, err := sarama.NewConsumerGroup(cfg.GetState().Transport.Broker.Addrs, "api-gateway-1", c)
	if err != nil {
		return nil, fmt.Errorf("error connect to consumer group. %w", err)
	}

	return &ServiceConsumer{cg: cg, serviceHandler: handler, topicBuilder: topicBuilder}, nil
}

func (sc *ServiceConsumer) Consume(ctx context.Context) error {
	if err := sc.cg.Consume(ctx, sc.listenTopics(), sc.serviceHandler); err != nil {
		return fmt.Errorf("error consume. %w", err)
	}

	return nil
}

func (sc *ServiceConsumer) Close() error {
	if err := sc.cg.Close(); err != nil {
		return fmt.Errorf("error close consumer group. %w", err)
	}

	return nil
}

func (sc *ServiceConsumer) listenTopics() []string {
	return []string{
		sc.topicBuilder.UserCreatedV1Topic(),
		"test.user_created.v1", // TODO delete
	}
}
