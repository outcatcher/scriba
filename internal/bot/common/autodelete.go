package common

import (
	"time"

	"gopkg.in/telebot.v3"
)

// AutodeleteMiddleware - middleware to support automatic deletion of messages.
func AutodeleteMiddleware(lifetime time.Duration) telebot.MiddlewareFunc {
	return func(next telebot.HandlerFunc) telebot.HandlerFunc {
		if lifetime == 0 {
			return next
		}

		return func(c telebot.Context) error {
			c.DeleteAfter(lifetime)

			return next(c)
		}
	}
}
