package schema

import (
	"context"

	"github.com/outcatcher/scriba/internal/entities"
)

// UseCases - requirements for the use cases.
type UseCases interface {
	UpdateCountByTelegramID(ctx context.Context, telegramID int64, delta int16) error
	GetPlayerCountByTelegramID(ctx context.Context, telegramID int64) (int32, error)
	RegisterWithTelegram(ctx context.Context, telegramID int64) error
	ListPlayers(ctx context.Context) ([]entities.Player, error)
}
