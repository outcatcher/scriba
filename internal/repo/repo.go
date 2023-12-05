package repo

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/maypok86/otter"
)

const (
	cacheShardCount = 2
	cacheCapacity   = 10
)

// Repo - db communications.
type Repo struct {
	db *sqlx.DB

	usersCache       []Player
	usersCacheByTGID *otter.Cache[int64, *Player]
	countSumCache    *otter.Cache[uuid.UUID, int32]
}

// New creates new DB instance with initialized caches.
func New(db *sqlx.DB) *Repo {
	usersCacheByTGID, err := otter.MustBuilder[int64, *Player](cacheCapacity).ShardCount(cacheShardCount).Build()
	if err != nil {
		panic(err)
	}

	countCache, err := otter.MustBuilder[uuid.UUID, int32](cacheCapacity).ShardCount(cacheShardCount).Build()
	if err != nil {
		panic(err)
	}

	return &Repo{
		db:               db,
		usersCacheByTGID: usersCacheByTGID,
		countSumCache:    countCache,
	}
}
