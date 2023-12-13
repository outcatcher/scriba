package repo

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type cache[K comparable, V any] interface {
	Get(key K) (V, bool)
	Set(key K, value V)
	Delete(key K)
}

type queryExecutor interface {
	sqlx.ExtContext

	GetContext(ctx context.Context, dest any, query string, args ...any) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}
