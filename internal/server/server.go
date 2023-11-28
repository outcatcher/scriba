package server

import "context"

type Server interface {
	RegisterUser(ctx context.Context, id, chatID int64, role string) error
}

func Start() error {
	return nil
}
