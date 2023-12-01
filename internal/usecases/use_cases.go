package usecases

import (
	"fmt"

	"github.com/outcatcher/scriba/internal/core/config"
	"github.com/outcatcher/scriba/internal/core/storage"
	"github.com/outcatcher/scriba/internal/repo"
)

type UseCases struct {
	repo *repo.Repo
}

func New(cfg config.Configuration) (*UseCases, error) {
	db, err := storage.Connect(cfg.Storage.Postgres)
	if err != nil {
		return nil, fmt.Errorf("error creating use cases: %w", err)
	}

	return &UseCases{repo: repo.New(db)}, nil
}
