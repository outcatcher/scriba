package usecases

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

// Player - registered player.
type Player struct {
	ID         uuid.UUID
	TelegramID int64
}

// RegisterWithTelegram creates new user from Telegram.
func (u *UseCases) RegisterWithTelegram(ctx context.Context, telegramID int64) error {
	_, err := u.repo.CreateUserFromTG(ctx, telegramID)
	if err != nil {
		return fmt.Errorf("failed to register new player: %w", err)
	}

	return nil
}

// ListPlayers lists registered users.
func (u *UseCases) ListPlayers(ctx context.Context) ([]Player, error) {
	players, err := u.repo.ListPlayers(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list players: %w", err)
	}

	result := make([]Player, len(players))

	for i, player := range players {
		result[i] = Player{
			ID:         player.ID,
			TelegramID: player.TGUserID,
		}
	}

	return result, nil
}
