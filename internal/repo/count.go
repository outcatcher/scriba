package repo

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

// Count describes structure of count_history table.
type Count struct {
	Delta    int16     `db:"delta"`
	PlayerID uuid.UUID `db:"player_id"`
}

// UpdatePlayerCount updates player point count.
func (r *Repo) UpdatePlayerCount(ctx context.Context, playerID uuid.UUID, delta int16) error {
	_, err := r.db.ExecContext(
		ctx,
		`INSERT INTO count_history (delta, player_id) VALUES ($1, $2);`,
		delta, playerID,
	)
	if err != nil {
		return fmt.Errorf("error inserting into count_history: %w", err)
	}

	r.countSumCache.Delete(playerID)

	return nil
}

// GetPlayerCount returns player count.
func (r *Repo) GetPlayerCount(ctx context.Context, id uuid.UUID) (int32, error) {
	if sum, ok := r.countSumCache.Get(id); ok {
		return sum, nil
	}

	var count int32

	err := r.db.GetContext(
		ctx,
		&count,
		`SELECT COALESCE(SUM(delta), 0)
		FROM count_history
		WHERE player_id = $1;`,
		id,
	)
	if err != nil {
		return 0, fmt.Errorf("error selecting current count from users: %w", err)
	}

	r.countSumCache.Set(id, count)

	return count, nil
}
