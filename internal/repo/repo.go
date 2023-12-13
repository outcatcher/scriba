package repo

import (
	"github.com/google/uuid"
	"github.com/maypok86/otter"
	"github.com/outcatcher/scriba/internal/entities"
)

const (
	cacheShardCount = 2
	cacheCapacity   = 1000 // breaks then capacity is too low
)

// Repo - db communications.
type Repo struct {
	db queryExecutor

	usersCache       []entities.Player
	usersCacheByTGID cache[int64, *entities.Player]
	countSumCache    cache[uuid.UUID, int32]
}

// New creates new DB instance with initialized caches.
func New(db queryExecutor) *Repo {
	usersCacheByTGID, err := otter.MustBuilder[int64, *entities.Player](cacheCapacity).
		ShardCount(cacheShardCount).Build()
	if err != nil {
		panic(err) // todo: return error
	}

	countCache, err := otter.MustBuilder[uuid.UUID, int32](cacheCapacity).
		ShardCount(cacheShardCount).Build()
	if err != nil {
		panic(err) // todo: return error
	}

	return &Repo{
		db:               db,
		usersCacheByTGID: usersCacheByTGID,
		countSumCache:    countCache,
	}
}
