package domain

type MessageRepository interface {
    Save(message *Message) error
    FindMessagesByChatId(chatId int) ([]Message, error)
}