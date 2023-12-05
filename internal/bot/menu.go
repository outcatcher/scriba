package bot

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"gopkg.in/telebot.v3"
)

const btnExit = "exit"

var errMenuNoChat = errors.New("select user menu without chat")

func (h *handlers) selectUserMenu(c telebot.Context) (*telebot.ReplyMarkup, error) {
	ctx := context.Background()

	players, err := h.app.ListPlayers(ctx)
	if err != nil {
		return nil, fmt.Errorf("error listing players: %w", err)
	}

	response := &telebot.ReplyMarkup{
		OneTimeKeyboard: true,
		RemoveKeyboard:  true,
	}

	chat, ok := c.Recipient().(*telebot.Chat)
	if !ok {
		return nil, errMenuNoChat
	}

	rows := make([]telebot.Row, 0, len(players)+1)

	for _, player := range players {
		member, err := c.Bot().ChatMemberOf(chat, &telebot.User{ID: player.TelegramID})
		if err != nil {
			slog.Error("failed to get chat membership", "chat_id", chat.ID, "user_id", player.TelegramID)

			continue
		}

		btn := response.Data(member.User.FirstName, member.User.Recipient())

		rows = append(rows, response.Row(btn))
	}

	rows = append(rows, response.Row(response.Data("Выход", btnExit)))

	response.Inline()

	return response, err
}
