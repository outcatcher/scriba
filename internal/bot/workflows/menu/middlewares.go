package menu

import (
	"log/slog"

	"github.com/outcatcher/scriba/internal/bot/common"
	"gopkg.in/telebot.v3"
)

func (u *userMenuState) forbidSelf(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		if c.Sender() != nil &&
			u.selectedUser != nil &&
			c.Sender().ID == u.selectedUser.telegramID {
			slog.Error("users tried to rate themselves", "sender", u.selectedUser.telegramID)

			common.ErrorReply(u.api, c.Message(), "Вы пытаетесь изменить балы для самого себя")

			return nil
		}

		return next(c)
	}
}
