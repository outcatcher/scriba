package repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/maypok86/otter"
)

type Repo struct {
	db *sqlx.DB

	usersCacheByTGID *otter.Cache[int64, *User]
}

func New(db *sqlx.DB) *Repo {
	usersCacheByTGID := otter.NewCache(otter.Config[int64, *User]{
		Capacity:   10,
		ShardCount: 8,
	})

	return &Repo{
		db:               db,
		usersCacheByTGID: usersCacheByTGID,
	}
}
