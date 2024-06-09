package memory

import (
	"github.com/devkemc/real-chat-with-go/internal/chat/domain/entity"
	"strconv"
	"sync"
	"testing"
)

func TestActiveUserInMemory_SaveActiveUser(t *testing.T) {
	activeUser := &entity.ActiveUser{Name: "TestUser"}
	mockDb.resetDb()
	inMemory := NewActiveUserInMemory(mockDb)

	err := inMemory.SaveActiveUser(activeUser)
	if err != nil {
		t.Errorf("SaveActiveUser() error = %v, want nil", err)
	}

	if len(mockDb.activeUsers) != 1 {
		t.Errorf("Expected 1 active user, got %d", len(mockDb.activeUsers))
	}
	if mockDb.activeUsers[0].Name != "TestUser" {
		t.Errorf("Expected active user name to be 'TestUser', got %s", mockDb.activeUsers[0].Name)
	}
}

func TestActiveUserInMemory_GetActiveUsers(t *testing.T) {
	activeUser := &entity.ActiveUser{Name: "TestUser"}
	mockDb.resetDb()
	mockDb.activeUsers = append(mockDb.activeUsers, activeUser)
	inMemory := NewActiveUserInMemory(mockDb)
	activeUsers, err := inMemory.GetActiveUsers()
	if err != nil {
		t.Errorf("GetActiveUsers() error = %v, want nil", err)
	}

	if len(activeUsers) != 1 {
		t.Errorf("Expected 1 active user, got %d", len(activeUsers))
	}
	if activeUsers[0].Name != "TestUser" {
		t.Errorf("Expected active user name to be 'TestUser', got %s", activeUsers[0].Name)
	}
}

func TestActiveUserInMemory_ConcurrentAccess(t *testing.T) {
	mockDb.resetDb()
	inMemory := NewActiveUserInMemory(mockDb)
	var wg sync.WaitGroup
	numRoutines := 10
	wg.Add(numRoutines)

	for i := 0; i < numRoutines; i++ {
		go func(i int) {
			defer wg.Done()
			user := &entity.ActiveUser{Name: "User" + strconv.Itoa(i)}
			err := inMemory.SaveActiveUser(user)
			if err != nil {
				t.Errorf("SaveActiveUser() error = %v, want nil", err)
			}
		}(i)
	}

	wg.Wait()

	activeUsers, _ := inMemory.GetActiveUsers()
	if len(activeUsers) != numRoutines {
		t.Errorf("Expected %d active users, got %d", numRoutines, len(activeUsers))
	}
}
