package usecases

import (
	"context"
	"fmt"

	"github.com/outcatcher/scriba/internal/entities"
)

// RegisterWithTelegram creates new user from Telegram.
func (u *UseCases) RegisterWithTelegram(ctx context.Context, telegramID int64) error {
	_, err := u.repo.CreateUserFromTG(ctx, telegramID)
	if err != nil {
		return fmt.Errorf("failed to register new player: %w", err)
	}

	return nil
}

// ListPlayers lists registered users.
func (u *UseCases) ListPlayers(ctx context.Context) ([]entities.Player, error) {
	players, err := u.repo.ListPlayers(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list players: %w", err)
	}

	return players, nil
}
