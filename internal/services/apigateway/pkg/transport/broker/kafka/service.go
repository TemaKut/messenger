package kafka

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/TemaKut/messenger/internal/services/apigateway/internal/logger"
	"github.com/TemaKut/messenger/pkg/service/models/proto/events"
	"google.golang.org/protobuf/proto"
)

type ServiceHandler struct {
	log logger.Logger
}

func NewServiceHandler(log logger.Logger) *ServiceHandler {
	return &ServiceHandler{log: log}
}

func (h *ServiceHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h *ServiceHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h *ServiceHandler) ConsumeClaim(s sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		s.MarkMessage(msg, "")
		var event events.Event

		if err := proto.Unmarshal(msg.Value, &event); err != nil {
			h.log.Error(fmt.Sprintf("error unmarshal service event. %s", err))
			continue
		}

		if err := h.handleEvent(&event); err != nil {
			h.log.Error(fmt.Sprintf("error handle event. %s", err))
		}
	}

	return nil
}

func (h *ServiceHandler) handleEvent(event *events.Event) error {
	switch {
	case event.GetUserCreatedV1() != nil:
		fmt.Printf("Recived user created event) uid -> %s\n", event.GetUserCreatedV1().GetUserId())
	}

	return nil
}
