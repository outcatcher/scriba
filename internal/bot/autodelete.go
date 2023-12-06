package bot

import (
	"github.com/outcatcher/scriba/internal/core/config"
	"gopkg.in/telebot.v3"
)

func autodelete(cfg config.BotConfig) telebot.MiddlewareFunc {
	return func(next telebot.HandlerFunc) telebot.HandlerFunc {
		lifetime := cfg.Messages.Lifetime

		if lifetime == 0 {
			return next
		}

		return func(c telebot.Context) error {
			c.DeleteAfter(lifetime)

			return next(c)
		}
	}
}
