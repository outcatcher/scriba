package repo

import (
	"context"
	"fmt"
)

// Count describes structure of count_history table.
type Count struct {
	Delta    int16  `db:"delta"`
	PlayerID string `db:"player_id"`
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
