package memory

import "chat-service/internal/domain"

type MessageRepositoryMemory struct {
	messages []*domain.Message
}

func NewMessageRepositoryMemory() *MessageRepositoryMemory {
	return &MessageRepositoryMemory{
		messages: []*domain.Message{},
	}
}

func (r *MessageRepositoryMemory) Save(message *domain.Message) error {
	r.messages = append(r.messages, message)
	return nil
}