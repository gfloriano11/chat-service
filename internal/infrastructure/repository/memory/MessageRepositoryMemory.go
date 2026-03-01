package memory

import (
	"chat-service/internal/domain"
	"sync"
)

type MessageRepositoryMemory struct {
	messages []domain.Message
	mutex sync.Mutex
	autoId int
}

func NewMessageRepositoryMemory() *MessageRepositoryMemory {
	return &MessageRepositoryMemory{
		messages: []domain.Message{},
		autoId: 0,
	}
}

func (repository *MessageRepositoryMemory) Save(message *domain.Message) error {
	repository.mutex.Lock()
	defer repository.mutex.Unlock()

	repository.autoId++
	message.Id = repository.autoId

	repository.messages = append(repository.messages, *message)

	return nil
}

func (repository *MessageRepositoryMemory) FindAllMessages() ([]domain.Message, error) {
	repository.mutex.Lock()
	defer repository.mutex.Unlock()

	return repository.messages, nil
}