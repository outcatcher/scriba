package repo

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/outcatcher/scriba/internal/entities"
)

// player describes structure of users table.
type player[T entities.Player] struct {
	ID       uuid.UUID `db:"id"`
	TGUserID int64     `db:"tg_user_id"`
}

func (p player[T]) toEntity() *T {
	return &T{
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

	result := new(player[entities.Player])

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

	result := make([]player[entities.Player], 0)

	err := r.db.SelectContext(ctx, &result, `SELECT * FROM players;`)
	if err != nil {
		return nil, fmt.Errorf("error selecting players: %w", err)
	}

	r.usersCache = convertEntitySlice(result)

	return r.usersCache, nil
}
