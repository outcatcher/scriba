package usecases

import (
	"context"

	"github.com/google/uuid"
	"github.com/outcatcher/scriba/internal/entities"
)

type repository interface {
	GetPlayerCount(ctx context.Context, id uuid.UUID) (int32, error)
	InsertPlayerCountChange(ctx context.Context, playerID uuid.UUID, delta int16) error
	CreateUserFromTG(ctx context.Context, telegramID int64) (uuid.UUID, error)
	FindUserByTelegramID(ctx context.Context, telegramID int64) (*entities.Player, error)
	ListPlayers(ctx context.Context) ([]entities.Player, error)
}

// UseCases holds main app logic.
type UseCases struct {
	repo repository
}

// WithRepo make use cases to use given repository.
func (u *UseCases) WithRepo(repo repository) {
	u.repo = repo
}
