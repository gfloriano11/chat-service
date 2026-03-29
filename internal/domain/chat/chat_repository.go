package chat

type ChatRepository interface {
	Save(chat *Chat) error
	IsUserParticipant(id, userId int) bool
	FindChatsByUserId(id int) (*[]ChatListItem, error)
}