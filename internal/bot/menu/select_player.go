package menu

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/outcatcher/scriba/internal/entities"
	"gopkg.in/telebot.v3"
)

func playersList(c telebot.Context, players []entities.Player) []userInfo {
	result := make([]userInfo, 0, len(players))

	switch c.Chat().Type { //nolint:exhaustive
	case telebot.ChatPrivate:
		for _, player := range players {
			chat, err := c.Bot().ChatByID(player.TelegramID)
			if err != nil {
				slog.Error("failed to get chat membership", "chat_id", c.Chat().ID, "user_id", player.TelegramID)

				continue
			}

			result = append(result, userInfo{
				name:       chat.FirstName,
				telegramID: player.TelegramID,
			})
		}
	case telebot.ChatGroup, telebot.ChatSuperGroup:
		for _, player := range players {
			member, err := c.Bot().ChatMemberOf(c.Chat(), &telebot.User{ID: player.TelegramID})
			if err != nil {
				slog.Error("failed to get chat membership", "chat_id", c.Chat().ID, "user_id", player.TelegramID)

				continue
			}

			result = append(result, userInfo{
				name:       member.User.FirstName,
				telegramID: player.TelegramID,
			})
		}
	default:
		return nil
	}

	return result
}

func (u *userMenuState) selectPlayer(c telebot.Context) error {
	ctx := context.Background()
	msgText := "Выберите игрока"

	if u.baseMsg == nil {
		u.baseMsg = c.Message()
	}

	menu := &telebot.ReplyMarkup{
		OneTimeKeyboard: true,
		RemoveKeyboard:  true,
	}

	players, err := u.app.ListPlayers(ctx)
	if err != nil {
		return fmt.Errorf("error listing players: %w", err)
	}

	rows := make([]telebot.Row, 0, len(players)+1)

	for _, player := range playersList(c, players) {
		btn := menu.Data(player.name, strconv.FormatInt(player.telegramID, 10))

		selectedUser := userInfo{player.name, player.telegramID}

		u.handler.Handle(&btn, u.userDetails(ctx, selectedUser))

		rows = append(rows, menu.Row(btn))
	}

	exitBtn := menu.Data(textExit, btnExit)
	u.handler.Handle(&exitBtn, u.exit)

	rows = append(rows, telebot.Row{exitBtn})

	menu.Inline(rows...)

	if u.baseMenuMsg != nil {
		msg, err := c.Bot().Edit(u.baseMenuMsg, msgText, menu)
		if err != nil {
			return fmt.Errorf("error editing menu: %w", err)
		}

		u.baseMenuMsg = msg

		return nil
	}

	baseMenuMsg, err := c.Bot().Reply(c.Message(), msgText, menu)
	if err != nil {
		return fmt.Errorf("error sending reply: %w", err)
	}

	u.baseMenuMsg = baseMenuMsg

	return nil
}
