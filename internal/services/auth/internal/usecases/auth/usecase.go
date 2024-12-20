package auth

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/TemaKut/messenger/internal/services/auth/internal/clients/broker/kafka"
	authrepo "github.com/TemaKut/messenger/internal/services/auth/internal/repository/auth"
	authpb "github.com/TemaKut/messenger/pkg/service/models/proto/auth"
	"github.com/TemaKut/messenger/pkg/service/models/proto/events"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthUseCaseImpl struct {
	authRepository authrepo.AuthRepository
	kafkaProducer  kafka.KafkaSyncProducer
}

func NewAuthUseCase(authRepository authrepo.AuthRepository, kafkaProducer kafka.KafkaSyncProducer) AuthUseCase {
	return &AuthUseCaseImpl{authRepository: authRepository, kafkaProducer: kafkaProducer}
}

func (u *AuthUseCaseImpl) CreateUser(ctx context.Context, user *authpb.User) error {
	if user.Id == "" {
		user.Id = uuid.NewString()
	}

	if err := u.authRepository.AddUser(ctx, user); err != nil {
		return fmt.Errorf("error create user. %w", err)
	}

	event := events.Event{
		Event: &events.Event_UserCreatedV1{
			UserCreatedV1: &events.UserCreatedV1{
				UserId:    user.Id,
				CreatedAt: timestamppb.New(user.CreatedAt.AsTime()),
			},
		},
	}

	b, err := proto.Marshal(&event)
	if err != nil {
		return fmt.Errorf("error marshal event. %w", err)
	}

	_, _, err = u.kafkaProducer.SendMessage(&sarama.ProducerMessage{
		Topic: "test.user_created.v1",
		Key:   sarama.StringEncoder(user.Id),
		Value: sarama.ByteEncoder(b),
	})
	if err != nil {
		return fmt.Errorf("error fire event. %w", err)
	}

	return nil
}
