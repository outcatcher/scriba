package menu

import (
	"context"
	"fmt"
	"strconv"

	"gopkg.in/telebot.v3"
)

func (u *userMenuState) userDetails(ctx context.Context, user selectedUserState) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		u.currentUser = &user
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
		u.grp.Handle(&zeroBtn, u.changeScore(ctx, -score))

		rows = append(rows, telebot.Row{zeroBtn})

		exitBtn := menu.Data(textExit, btnExit)
		u.grp.Handle(&exitBtn, u.exit)

		backBtn := menu.Data(textBack, btnBack)
		u.grp.Handle(&backBtn, u.back)

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
