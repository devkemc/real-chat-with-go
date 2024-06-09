package memory

import (
	"github.com/devkemc/real-chat-with-go/internal/chat/domain/entity"
	"sync"
)

var (
	once               sync.Once
	dbInMemoryInstance *DbInMemory
	mx                 sync.Mutex
)

type DbInMemory struct {
	activeUsers []*entity.ActiveUser
	messages    []*entity.Message
}

func NewDbInMemory() *DbInMemory {
	once.Do(func() {
		dbInMemoryInstance = &DbInMemory{}
	})
	return dbInMemoryInstance
}

func (db *DbInMemory) resetDb() {
	db.activeUsers = make([]*entity.ActiveUser, 0)
	db.messages = make([]*entity.Message, 0)
}
