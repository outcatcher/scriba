package menu

import (
	"context"
	"fmt"
	"strconv"

	"gopkg.in/telebot.v3"
)

func (u *userMenuState) userDetails(ctx context.Context, user userInfo) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		u.selectedUser = &user
		u.currentLabel = labelUserInfo

		menu := &telebot.ReplyMarkup{
			OneTimeKeyboard: true,
			RemoveKeyboard:  true,
		}

		score, err := u.app.GetPlayerCountByTelegramID(ctx, user.telegramID)
		if err != nil {
			return fmt.Errorf("error getting player score: %w", err)
		}

		scoreStr := strconv.FormatInt(int64(score), 10)

		if score < 0 {
			scoreStr = "\\" + scoreStr // escape -
		}

		rows := u.scoreButtonsToRows(ctx, menu, changeUserScoreButtons)

		zeroBtn := menu.Data("Обнулить очки", "0")
		u.handler.Handle(&zeroBtn, u.changeScore(ctx, -score), u.forbidSelf)

		rows = append(rows, telebot.Row{zeroBtn})

		exitBtn := menu.Data(textExit, btnExit)
		u.handler.Handle(&exitBtn, u.exit)

		backBtn := menu.Data(textBack, btnBack)
		u.handler.Handle(&backBtn, u.back)

		// some other data here
		rows = append(rows, telebot.Row{backBtn, exitBtn})

		menu.Inline(rows...)

		messageText := fmt.Sprintf("Количество баллов у игрока %s: %s", user.name, scoreStr)

		msg, err := c.Bot().Edit(u.baseMenuMsg, messageText, menu)
		if err != nil {
			return fmt.Errorf("error editing menu: %w", err)
		}

		u.baseMenuMsg = msg

		return nil
	}
}
