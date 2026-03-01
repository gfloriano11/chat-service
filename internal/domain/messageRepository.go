package domain

type MessageRepository interface {
    Save(message *Message) error
    FindAllMessages() ([]Message, error)
}