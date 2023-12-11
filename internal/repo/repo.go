package repo

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/maypok86/otter"
	"github.com/outcatcher/scriba/internal/entities"
)

const (
	cacheShardCount = 2
	cacheCapacity   = 10
)

// Repo - db communications.
type Repo struct {
	db *sqlx.DB

	usersCache       []entities.Player
	usersCacheByTGID *otter.Cache[int64, *entities.Player]
	countSumCache    *otter.Cache[uuid.UUID, int32]
}

// New creates new DB instance with initialized caches.
func New(db *sqlx.DB) *Repo {
	usersCacheByTGID, err := otter.MustBuilder[int64, *entities.Player](cacheCapacity).
		ShardCount(cacheShardCount).Build()
	if err != nil {
		panic(err)
	}

	countCache, err := otter.MustBuilder[uuid.UUID, int32](cacheCapacity).
		ShardCount(cacheShardCount).Build()
	if err != nil {
		panic(err)
	}

	return &Repo{
		db:               db,
		usersCacheByTGID: usersCacheByTGID,
		countSumCache:    countCache,
	}
}
