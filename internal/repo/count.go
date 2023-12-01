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

// UpdateCount updates user point count.
func (r *Repo) UpdateCount(ctx context.Context, count Count) error {
	_, err := r.db.NamedExecContext(
		ctx,
		`INSERT INTO count_history (delta, player_id) VALUES (:delta, :player_id);`,
		count,
	)
	if err != nil {
		return fmt.Errorf("error inserting into count_history: %w", err)
	}

	return nil
}

// GetUserCount returns user count.
func (r *Repo) GetUserCount(ctx context.Context, id uuid.UUID) (int32, error) {
	var count int32

	err := r.db.GetContext(
		ctx,
		&count,
		`SELECT SUM(delta)
		FROM count_history
		WHERE user_id = $1;`,
		id,
	)
	if err != nil {
		return 0, fmt.Errorf("error selecting current count from users: %w", err)
	}

	return count, err
}
