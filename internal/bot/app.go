package bot

import (
	"fmt"

	"github.com/outcatcher/scriba/internal/core/config"
	"github.com/outcatcher/scriba/internal/core/storage"
	"github.com/outcatcher/scriba/internal/repo"
	"github.com/outcatcher/scriba/internal/usecases"
)

func initApp(cfg *config.Configuration) (*usecases.UseCases, error) {
	db, err := storage.Connect(cfg.Storage.Postgres)
	if err != nil {
		return nil, fmt.Errorf("error creating use cases: %w", err)
	}

	cases := new(usecases.UseCases)
	cases.WithRepo(repo.New(db))

	return cases, nil
}
