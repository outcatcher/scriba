package usecases

import (
	"context"
	"fmt"
)

// RegisterWithTelegram creates new user from Telegram.
func (u *UseCases) RegisterWithTelegram(ctx context.Context, telegramID int64) error {
	_, err := u.repo.CreateUserFromTG(ctx, telegramID)
	if err != nil {
		return fmt.Errorf("failed to register new player: %w", err)
	}

	return nil
}
