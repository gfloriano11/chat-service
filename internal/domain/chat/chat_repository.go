package chat

type ChatRepository interface {
	Save(chat *Chat) error
}