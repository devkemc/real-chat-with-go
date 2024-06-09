package repository

import "github.com/devkemc/real-chat-with-go/internal/chat/domain/entity"

type ActiveUserRepository interface {
	SaveActiveUser(activeUser *entity.ActiveUser) error
	GetActiveUsers() ([]*entity.ActiveUser, error)
}
