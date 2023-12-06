package menu

import (
	"context"
	"fmt"
	"math"

	"gopkg.in/telebot.v3"
)

func (u *userMenuState) changeScore(ctx context.Context, delta int32) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		if u.currentUser == nil {
			return errMissingUser
		}

		// internal type is int16
		for delta > math.MaxInt16 {
			if err := u.app.UpdateCountByTelegramID(ctx, u.currentUser.telegramID, math.MaxInt16); err != nil {
				return fmt.Errorf("error updating player's count: %w", err)
			}

			delta -= math.MaxInt16
		}

		if err := u.app.UpdateCountByTelegramID(ctx, u.currentUser.telegramID, int16(delta)); err != nil {
			return fmt.Errorf("error updating player's count: %w", err)
		}

		return u.userDetails(ctx, *u.currentUser)(c)
	}
}

func uniqueByNumber(src int32) string {
	var prefix string
	if src < 0 {
		prefix = "minus"
	} else {
		prefix = "plus"
	}

	return fmt.Sprintf("%s%d", prefix, int32(math.Abs(float64(src))))
}

func (u *userMenuState) scoreButtonsToRows(
	ctx context.Context, menu *telebot.ReplyMarkup, src [][]int32,
) []telebot.Row {
	rows := make([]telebot.Row, len(src))

	for i, numRow := range changeUserScoreButtons {
		row := make(telebot.Row, len(numRow))

		for j, num := range numRow {
			btnValue := fmt.Sprintf("%+d", num)
			btn := menu.Data(btnValue, uniqueByNumber(num))

			u.grp.Handle(&btn, u.changeScore(ctx, num), u.forbidSelf)

			row[j] = btn
		}

		rows[i] = row
	}

	return rows
}
