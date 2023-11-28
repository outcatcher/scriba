package bot

import (
	"gopkg.in/telebot.v3"
	"scriba/internal/config"
)

func autodelete(cfg config.BotConfig) telebot.MiddlewareFunc {
	return func(next telebot.HandlerFunc) telebot.HandlerFunc {
		lifetime := cfg.Messages.Lifetime

		if lifetime == 0 {
			return next
		}

		return func(ctx telebot.Context) error {
			ctx.DeleteAfter(lifetime)

			return next(ctx)
		}
	}
}
