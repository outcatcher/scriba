package schema

import "gopkg.in/telebot.v3"

// TelegramAPI implements telegram interactions.
type TelegramAPI interface {
	// message handling

	Send(to telebot.Recipient, what any, opts ...any) (*telebot.Message, error)
	Reply(to *telebot.Message, what interface{}, opts ...interface{}) (*telebot.Message, error)
	Edit(msg telebot.Editable, what interface{}, opts ...interface{}) (*telebot.Message, error)
	Delete(msg telebot.Editable) error

	// users handling

	ChatMemberOf(chat, user telebot.Recipient) (*telebot.ChatMember, error)
	ChatByID(id int64) (*telebot.Chat, error)
}
