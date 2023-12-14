package usecases

import (
	"context"
	"fmt"
	"time"

	"github.com/outcatcher/scriba/internal/entities"
)

var periodDelta = map[entities.HistoryPeriod]time.Duration{
	entities.HistoryPeriod3Days: -3 * 24 * time.Hour,
	entities.HistoryPeriodWeek:  -7 * 24 * time.Hour,
}

// UpdateCountByTelegramID find player by Telegram user ID and updates its count.
func (u *UseCases) UpdateCountByTelegramID(ctx context.Context, telegramID int64, delta int16) error {
	user, err := u.repo.FindUserByTelegramID(ctx, telegramID)
	if err != nil {
		return fmt.Errorf("error finding player: %w", err)
	}

	err = u.repo.InsertPlayerCountChange(ctx, user.ID, delta)
	if err != nil {
		return fmt.Errorf("error updating player count: %w", err)
	}

	return nil
}

// GetPlayerCountByTelegramID returns total count of player by TG ID.
func (u *UseCases) GetPlayerCountByTelegramID(ctx context.Context, telegramID int64) (int32, error) {
	user, err := u.repo.FindUserByTelegramID(ctx, telegramID)
	if err != nil {
		return 0, fmt.Errorf("error finding player: %w", err)
	}

	count, err := u.repo.GetPlayerCount(ctx, user.ID)
	if err != nil {
		return 0, fmt.Errorf("error finding player count: %w", err)
	}

	return count, nil
}

// GetPlayerHistoryByTelegramID returns player's event history for given period and total sum before.
func (u *UseCases) GetPlayerHistoryByTelegramID(
	ctx context.Context, telegramID int64, period entities.HistoryPeriod,
) ([]entities.CountHistoryEvent, error) {
	user, err := u.repo.FindUserByTelegramID(ctx, telegramID)
	if err != nil {
		return nil, fmt.Errorf("error finding player: %w", err)
	}

	endTime := time.Now().UTC()
	startTime := endTime.Add(periodDelta[period])

	history, err := u.repo.GetCountHistoryForPeriod(ctx, user.ID, startTime, endTime)
	if err != nil {
		return nil, fmt.Errorf("error getting event history for period: %w", err)
	}

	return history, nil
}
