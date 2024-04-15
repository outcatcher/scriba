package menu

import (
	"fmt"
	"log/slog"
	"strconv"

	"github.com/outcatcher/scriba/internal/bot/common"
	"github.com/outcatcher/scriba/internal/entities"
	"gopkg.in/telebot.v3"
)

func (u *userMenuState) playersList(chat *telebot.Chat, players []entities.Player) []telegramUserInfo {
	result := make([]telegramUserInfo, 0, len(players))

	switch chat.Type { //nolint:exhaustive
	case telebot.ChatPrivate:
		for _, player := range players {
			playerChat, err := u.api.ChatByID(player.TelegramID)
			if err != nil {
				slog.Error("failed to get chat membership", "chat_id", chat.ID, "user_id", player.TelegramID)

				continue
			}

			result = append(result, telegramUserInfo{
				name:       playerChat.FirstName,
				telegramID: player.TelegramID,
			})
		}
	case telebot.ChatGroup, telebot.ChatSuperGroup:
		for _, player := range players {
			member, err := u.api.ChatMemberOf(chat, &telebot.User{ID: player.TelegramID})
			if err != nil {
				slog.Error("failed to get chat membership", "chat_id", chat.ID, "user_id", player.TelegramID)

				continue
			}

			result = append(result, telegramUserInfo{
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
	ctx := common.GetContextFromContext(c)
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

	for _, player := range u.playersList(c.Chat(), players) {
		btn := menu.Data(player.name, strconv.FormatInt(player.telegramID, 10))

		selectedUser := telegramUserInfo{player.name, player.telegramID}

		u.handler.Handle(&btn, u.userDetails(selectedUser))

		rows = append(rows, menu.Row(btn))
	}

	exitBtn := menu.Data(textExit, btnExit)
	u.handler.Handle(&exitBtn, u.exit)

	rows = append(rows, telebot.Row{exitBtn})

	menu.Inline(rows...)

	if u.baseMenuMsg != nil {
		msg, editErr := u.api.Edit(u.baseMenuMsg, msgText, menu)
		if editErr != nil {
			return fmt.Errorf("error editing menu: %w", editErr)
		}

		u.baseMenuMsg = msg

		return nil
	}

	baseMenuMsg, err := u.api.Reply(c.Message(), msgText, menu)
	if err != nil {
		return fmt.Errorf("error sending reply: %w", err)
	}

	u.baseMenuMsg = baseMenuMsg

	return nil
}
