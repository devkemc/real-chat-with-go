package repository

import "github.com/devkemc/real-chat-with-go/internal/chat/domain/entity"

type ChatRepository interface {
	SaveMessage(message *entity.Message) error
	GetMessages() ([]*entity.Message, error)
}
