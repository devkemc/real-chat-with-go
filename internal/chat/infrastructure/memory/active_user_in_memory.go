package memory

import (
	"github.com/devkemc/real-chat-with-go/internal/chat/domain/entity"
)

type ActiveUserInMemory struct {
	db *DbInMemory
}

func NewActiveUserInMemory(db *DbInMemory) *ActiveUserInMemory {
	return &ActiveUserInMemory{
		db: db,
	}
}

func (r *ActiveUserInMemory) SaveActiveUser(activeUser *entity.ActiveUser) error {
	mx.Lock()
	r.db.activeUsers = append(r.db.activeUsers, activeUser)
	mx.Unlock()
	return nil
}

func (r *ActiveUserInMemory) GetActiveUsers() ([]*entity.ActiveUser, error) {
	mx.Lock()
	temporaryActiveUsers := make([]*entity.ActiveUser, len(r.db.activeUsers))
	copy(temporaryActiveUsers, r.db.activeUsers)
	mx.Unlock()
	return temporaryActiveUsers, nil
}
