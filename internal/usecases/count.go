package usecases

import (
	"context"
	"fmt"
)

// UpdateCountByTelegramID find player by Telegram user ID and updates its count.
func (u *UseCases) UpdateCountByTelegramID(ctx context.Context, telegramID int64, delta int16) error {
	user, err := u.repo.FindUserByTelegramID(ctx, telegramID)
	if err != nil {
		return fmt.Errorf("error finding player: %w", err)
	}

	err = u.repo.UpdatePlayerCount(ctx, user.ID, delta)
	if err != nil {
		return fmt.Errorf("error updating player count: %w", err)
	}

	return nil
}

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
