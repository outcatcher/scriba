package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/outcatcher/scriba/internal/entities"
)

type countHistoryItem[T entities.CountHistoryEvent] struct {
	Timestamp time.Time `db:"timestamp"`
	Delta     int16     `db:"delta"`
}

func (c countHistoryItem[T]) toEntity() *T {
	return &T{
		Timestamp: c.Timestamp,
		Delta:     c.Delta,
	}
}

// InsertPlayerCountChange updates player point count.
func (r *Repo) InsertPlayerCountChange(ctx context.Context, playerID uuid.UUID, delta int16) error {
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

// GetCountHistoryForPeriod shows all score change events for given time.
func (r *Repo) GetCountHistoryForPeriod(
	ctx context.Context, id uuid.UUID, startDate, endDate time.Time,
) ([]entities.CountHistoryEvent, error) {
	result := make([]countHistoryItem[entities.CountHistoryEvent], 0)

	err := r.db.SelectContext(
		ctx,
		&result,
		`SELECT delta, timestamp
		FROM count_history
		WHERE player_id = $1
		  AND "timestamp" BETWEEN $2 AND $3
		ORDER BY "timestamp" DESC`,
		id, startDate, endDate,
	)
	if err != nil {
		return nil, fmt.Errorf("error getting count history on %s - %s: %w", startDate.String(), endDate.String(), err)
	}

	return convertEntitySlice(result), nil
}
