package common

import (
	"log/slog"

	"gopkg.in/telebot.v3"
)

// LogCommands returns middleware logging commands with slog.
func LogCommands() telebot.MiddlewareFunc {
	return func(next telebot.HandlerFunc) telebot.HandlerFunc {
		return func(c telebot.Context) error {
			slog.Info(
				"received command",
				"arguments", c.Args(),
			)

			err := next(c)
			if err != nil {
				slog.Error(
					"error processing command",
					"error", err,
				)
			}

			return err
		}
	}
}
