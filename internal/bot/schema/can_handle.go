package schema

import "gopkg.in/telebot.v3"

// Handler provides Handle method.
type Handler interface {
	Handle(endpoint interface{}, h telebot.HandlerFunc, m ...telebot.MiddlewareFunc)
}
