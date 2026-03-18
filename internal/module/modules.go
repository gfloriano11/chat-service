package module

import "gorm.io/gorm"

type Modules struct {
	ChatModule    ChatModule
	MessageModule MessageModule
	UserModule		UserModule
}

func NewModules(chat ChatModule, message MessageModule, user UserModule) *Modules {
	return &Modules{
		ChatModule:    chat,
		MessageModule: message,
		UserModule:  	 user,
	}
}

func CreateModules(db *gorm.DB) *Modules {
	return NewModules(
		NewChatModule(db),
		NewMessageModule(db),
		NewUserModule(db),
	)
}