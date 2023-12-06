package menu

import (
	"fmt"
	"log/slog"

	"gopkg.in/telebot.v3"
)

func (u *userMenuState) forbidSelf(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		if c.Sender() != nil &&
			u.currentUser != nil &&
			c.Sender().ID == u.currentUser.telegramID {
			slog.Error("users tried to rate themselves", "sender", u.currentUser.telegramID)

			err := c.Reply("Вы пытаетесь изменить балы для самого себя")
			if err != nil {
				return fmt.Errorf("error replying at forbidSelf: %w", err)
			}

			return nil
		}

		return next(c)
	}
}
