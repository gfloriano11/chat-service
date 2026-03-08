package module

import "gorm.io/gorm"

type Modules struct {
	ChatModule    ChatModule
	MessageModule MessageModule
}

func NewModules(chat ChatModule, message MessageModule) *Modules {
	return &Modules{
		ChatModule:    chat,
		MessageModule: message,
	}
}

func CreateModules(db *gorm.DB) *Modules {
	return NewModules(
		NewChatModule(db),
		NewMessageModule(db),
	)
}