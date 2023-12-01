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

// FindUserByTelegramID search for user by telegram ID.
func (r *Repo) FindUserByTelegramID(ctx context.Context, telegramID int64) (*User, error) {
	if value, ok := r.usersCacheByTGID.Get(telegramID); ok {
		return value, nil
	}

	result := new(User)

	err := r.db.GetContext(ctx, &result, `SELECT * FROM users WHERE tg_user_id = $1;`, telegramID)
	if err != nil {
		return nil, fmt.Errorf("error selecting current by Telegram ID: %w", err)
	}

	r.usersCacheByTGID.Set(result.TGUserID, result)

	return result, nil
}
