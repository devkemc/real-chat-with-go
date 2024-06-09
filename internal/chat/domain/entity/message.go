package entity

import (
	"github.com/google/uuid"
	"time"
)

type Message struct {
	ID        string `json:"id"`
	From      string `json:"from"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

func NewMessage(from, content string) *Message {
	return &Message{
		ID:        uuid.New().String(),
		From:      from,
		Content:   content,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
}
