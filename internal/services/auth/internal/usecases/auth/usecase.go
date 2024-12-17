package auth

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/TemaKut/messenger/internal/services/auth/internal/clients/broker/kafka"
	authrepo "github.com/TemaKut/messenger/internal/services/auth/internal/repository/auth"
	authpb "github.com/TemaKut/messenger/pkg/service/models/proto/auth"
	"github.com/google/uuid"
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

	_, _, err := u.kafkaProducer.SendMessage(&sarama.ProducerMessage{
		Topic: "test.user_created.v1",
		Key:   sarama.StringEncoder(user.Id),
		Value: sarama.StringEncoder("i am created)"),
	})
	if err != nil {
		return fmt.Errorf("error fire event. %w", err)
	}

	return nil
}
