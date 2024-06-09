package memory

import (
	"github.com/devkemc/real-chat-with-go/internal/chat/domain/entity"
)

type ChatRepositoryInMemory struct {
	db *DbInMemory
}

func NewChatRepositoryInMemory(db *DbInMemory) *ChatRepositoryInMemory {
	return &ChatRepositoryInMemory{db: db}
}

func (r *ChatRepositoryInMemory) SaveMessage(message *entity.Message) error {
	mx.Lock()
	r.db.messages = append(r.db.messages, message)
	mx.Unlock()
	return nil
}

func (r *ChatRepositoryInMemory) GetMessages() ([]*entity.Message, error) {
	mx.Lock()
	temporaryMessages := make([]*entity.Message, len(r.db.messages))
	copy(temporaryMessages, r.db.messages)
	mx.Unlock()
	return temporaryMessages, nil
}
