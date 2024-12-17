package kafka

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/TemaKut/messenger/internal/services/auth/internal/config"
)

type KafkaSyncProducer = sarama.SyncProducer

func NewProducer(cfg *config.Config) (KafkaSyncProducer, func(), error) {
	producer, err := sarama.NewSyncProducer(cfg.GetState().Broker.Addrs, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("error make sync producer. %w", err)
	}

	return producer, func() { _ = producer.Close() }, nil
}
