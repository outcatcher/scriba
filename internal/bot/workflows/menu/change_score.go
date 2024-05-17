package menu

import (
	"errors"
	"fmt"
	"log/slog"
	"math"

	"github.com/outcatcher/scriba/internal/bot/common"
	"gopkg.in/telebot.v3"
)

var (
	errUpdatingCount      = errors.New("error updating player's count")
	errLoggingCountChange = errors.New("error logging players count change")
)

func (u *userMenuState) changeScore(delta int32) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		ctx := common.GetContextFromContext(c)

		if u.selectedUser == nil {
			return errMissingUser
		}

		// internal type is int16
		for delta > math.MaxInt16 {
			if err := u.app.UpdateCountByTelegramID(ctx, u.selectedUser.telegramID, math.MaxInt16); err != nil {
				return fmt.Errorf("%w: %w", errUpdatingCount, err)
			}

			delta -= math.MaxInt16
		}

		if err := u.app.UpdateCountByTelegramID(ctx, u.selectedUser.telegramID, int16(delta)); err != nil {
			return fmt.Errorf("%w: %w", errUpdatingCount, err)
		}

		if err := u.logScoreChange(c, u.selectedUser.name, delta); err != nil {
			slog.Error("log score change", "error", err)
		}

		return u.userDetails(*u.selectedUser)(c)
	}
}

func (u *userMenuState) logScoreChange(c telebot.Context, username string, delta int32) error {
	deltaStr := fmt.Sprintf("\\%+d", delta)

	messageText := fmt.Sprintf("Количество баллов для %s изменено: %s", username, deltaStr)

	_, err := u.api.Send(c.Recipient(), messageText)
	if err != nil {
		return fmt.Errorf("%w: %w", errLoggingCountChange, err)
	}

	return nil
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

func (u *userMenuState) scoreButtonsToRows(menu *telebot.ReplyMarkup, src [][]int32,
) []telebot.Row {
	rows := make([]telebot.Row, len(src))

	for i, numRow := range changeUserScoreButtons {
		row := make(telebot.Row, len(numRow))

		for j, num := range numRow {
			btnValue := fmt.Sprintf("%+d", num)
			btn := menu.Data(btnValue, uniqueByNumber(num))

			u.handler.Handle(&btn, u.changeScore(num), u.forbidSelf)

			row[j] = btn
		}

		rows[i] = row
	}

	return rows
}
