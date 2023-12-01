package usecases

import (
	"context"
	"fmt"

	"github.com/outcatcher/scriba/internal/repo"
)

func (u *UseCases) UpdateCountByTelegramID(ctx context.Context, telegramID int64, delta int16) error {
	user, err := u.repo.FindUserByTelegramID(ctx, telegramID)
	if err != nil {
		return fmt.Errorf("error finding user: %w", err)
	}

	err = u.repo.UpdateCount(ctx, repo.Count{
		Delta:    delta,
		PlayerID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("error updating user count: %w", err)
	}

	return nil
}
