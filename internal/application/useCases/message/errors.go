package message

import "errors"

var ErrUserNotInChat = errors.New("User can't send messages to a chat that user isn't added.")