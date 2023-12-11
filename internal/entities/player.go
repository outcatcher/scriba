package entities

import "github.com/google/uuid"

// Player - registered player.
type Player struct {
	ID         uuid.UUID
	TelegramID int64
}
