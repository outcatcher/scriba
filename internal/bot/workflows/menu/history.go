package menu

import (
	"fmt"
	"strings"
	"time"
	_ "time/tzdata" // we need support of MSK timezone

	"github.com/outcatcher/scriba/internal/bot/common"
	"github.com/outcatcher/scriba/internal/entities"
	"gopkg.in/telebot.v3"
)

const timestampLayout = "02\\.01\\.2006 15:04"

var locationMoscow, _ = time.LoadLocation("Europe/Moscow")

var historyOptions = map[entities.HistoryPeriod]string{
	entities.HistoryPeriod3Days: "История за 3 дня",
	entities.HistoryPeriodWeek:  "История за неделю",
}

func (u *userMenuState) replyScoreHistory(period entities.HistoryPeriod) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		ctx := common.GetContextFromContext(c)

		// get sum before start of the period
		// get list of deltas for the period
		events, err := u.app.GetPlayerHistoryByTelegramID(ctx, u.selectedUser.telegramID, period)
		if err != nil {
			return fmt.Errorf("error gettig event history: %w", err)
		}

		lines := make([]string, len(events))

		for i, event := range events {
			timestamp := event.Timestamp.In(locationMoscow).Format(timestampLayout)

			lines[i] = fmt.Sprintf("%s:           \\%+d", timestamp, event.Delta)
		}

		resultMsgText := historyOptions[period] + "\n\n" + strings.Join(lines, "\n")

		_, err = u.api.Send(c.Recipient(), resultMsgText)
		if err != nil {
			return fmt.Errorf("error replying to the message: %w", err)
		}

		return nil
	}
}
