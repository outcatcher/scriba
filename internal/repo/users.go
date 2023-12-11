package repo

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/outcatcher/scriba/internal/entities"
)

// player describes structure of users table.
type player struct {
	ID       uuid.UUID `db:"id"`
	TGUserID int64     `db:"tg_user_id"`
}

func (p player) toEntity() *entities.Player {
	return &entities.Player{
		ID:         p.ID,
		TelegramID: p.TGUserID,
	}
}

// CreateUserFromTG creates new user from Telegram and returning its id.
func (r *Repo) CreateUserFromTG(ctx context.Context, telegramUserID int64) (uuid.UUID, error) {
	var createdID uuid.UUID

	err := r.db.GetContext(
		ctx,
		&createdID,
		`INSERT INTO players (tg_user_id)
		VALUES ($1)
		RETURNING id;`, telegramUserID)
	if err != nil {
		return createdID, fmt.Errorf("error inserting user: %w", err)
	}

	r.usersCache = nil

	return createdID, nil
}

// FindUserByTelegramID search for user by telegram ID.
func (r *Repo) FindUserByTelegramID(ctx context.Context, telegramID int64) (*entities.Player, error) {
	if value, ok := r.usersCacheByTGID.Get(telegramID); ok {
		return value, nil
	}

	result := new(player)

	err := r.db.GetContext(ctx, result, `SELECT * FROM players WHERE tg_user_id = $1;`, telegramID)
	if err != nil {
		return nil, fmt.Errorf("error selecting player by Telegram ID: %w", err)
	}

	r.usersCacheByTGID.Set(result.TGUserID, result.toEntity())

	return result.toEntity(), nil
}

// ListPlayers lists existing players.
func (r *Repo) ListPlayers(ctx context.Context) ([]entities.Player, error) {
	if len(r.usersCache) > 0 {
		return r.usersCache, nil
	}

	result := make([]player, 0)

	err := r.db.SelectContext(ctx, &result, `SELECT * FROM players;`)
	if err != nil {
		return nil, fmt.Errorf("error selecting players: %w", err)
	}

	eResult := make([]entities.Player, len(result))

	for i, item := range result {
		eResult[i] = *item.toEntity()
	}

	r.usersCache = eResult

	return eResult, nil
}
