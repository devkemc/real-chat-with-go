package memory

import (
	"github.com/devkemc/real-chat-with-go/internal/chat/domain/entity"
	"strconv"
	"sync"
	"testing"
)

func TestChatRepositoryInMemory_SaveMessage(t *testing.T) {
	activeUser := &entity.Message{Content: "This is a test message"}
	mockDb.resetDb()
	inMemory := NewChatRepositoryInMemory(mockDb)

	err := inMemory.SaveMessage(activeUser)
	if err != nil {
		t.Errorf("SaveMessage() error = %v, want nil", err)
	}

	if len(mockDb.messages) != 1 {
		t.Errorf("Expected 1 message, got %d", len(mockDb.messages))
	}
	if mockDb.messages[0].Content != "This is a test message" {
		t.Errorf("Expected message content to be 'This is a test message', got %s", mockDb.messages[0].Content)
	}
}

func TestChatRepositoryInMemory_GetMessages(t *testing.T) {
	mockDb.resetDb()
	mockDb.messages = append(mockDb.messages, &entity.Message{Content: "This is a test message"})
	inMemory := NewChatRepositoryInMemory(mockDb)

	messages, err := inMemory.GetMessages()
	if err != nil {
		t.Errorf("GetMessages() error = %v, want nil", err)
	}

	if len(messages) != 1 {
		t.Errorf("Expected 1 message, got %d", len(messages))
	}
	if messages[0].Content != "This is a test message" {
		t.Errorf("Expected message content to be 'This is a test message', got %s", messages[0].Content)
	}
}

func TestChatRepositoryInMemory_ConcurrentAccess(t *testing.T) {
	mockDb.resetDb()
	inMemory := NewChatRepositoryInMemory(mockDb)
	var wg sync.WaitGroup
	numRoutines := 10
	wg.Add(numRoutines)

	for i := 0; i < numRoutines; i++ {
		go func(i int) {
			defer wg.Done()
			message := &entity.Message{Content: "This is a test message" + strconv.Itoa(i)}
			err := inMemory.SaveMessage(message)
			if err != nil {
				t.Errorf("SaveMessage() error = %v, want nil", err)
			}
		}(i)
	}

	wg.Wait()
	if len(mockDb.messages) != numRoutines {
		t.Errorf("Expected %d messages, got %d", numRoutines, len(mockDb.messages))
	}
}
