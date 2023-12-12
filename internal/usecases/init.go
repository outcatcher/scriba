package usecases

import (
	"fmt"

	"github.com/outcatcher/scriba/internal/core/config"
	"github.com/outcatcher/scriba/internal/core/storage"
	"github.com/outcatcher/scriba/internal/repo"
)

// InitApp initializes new app instance.
func InitApp(cfg *config.Configuration) (*UseCases, error) {
	db, err := storage.Connect(cfg.Storage.Postgres)
	if err != nil {
		return nil, fmt.Errorf("error creating use cases: %w", err)
	}

	cases := new(UseCases)
	cases.WithRepo(repo.New(db))

	return cases, nil
}
