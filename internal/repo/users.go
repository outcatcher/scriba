package repo

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

// User describes structure of users table.
type User struct {
	ID           uuid.UUID `db:"id"`
	Username     string    `db:"username"`
	TGUserID     int64     `db:"tg_user_id"`
	CurrentCount int32     `db:"current_count"`
}

// CreateUserFromTG creates new user from Telegram and returning its id.
func (r *Repo) CreateUserFromTG(ctx context.Context, user *User) (uuid.UUID, error) {
	var createdID uuid.UUID

	query, args, err := r.db.BindNamed(
		`INSERT INTO users (username, tg_user_id) VALUES (:username, :tg_user_id) RETURNING id;`,
		user,
	)
	if err != nil {
		return createdID, fmt.Errorf("error binding query args: %w", err)
	}

	if err := r.db.GetContext(ctx, &createdID, query, args); err != nil {
		return createdID, fmt.Errorf("error inserting user: %w", err)
	}

	return createdID, nil
}

// GetUserCountByTGID returns user count by Telegram ID.
func (r *Repo) GetUserCountByTGID(ctx context.Context, telegramID int64) (int32, error) {
	var count int32

	err := r.db.GetContext(ctx, &count, `SELECT current_count FROM users WHERE tg_user_id = $1;`, telegramID)
	if err != nil {
		return 0, fmt.Errorf("error selecting current count from users: %w", err)
	}

	return count, err
}
