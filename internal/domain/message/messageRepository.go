package message

type MessageRepository interface {
    Save(message *Message) error
    FindMessagesByChatId(chatId int) ([]Message, error)
}