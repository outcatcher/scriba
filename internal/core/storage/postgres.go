package storage

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // use `postgres` driver
	"github.com/outcatcher/scriba/internal/core/config"
)

const dbDriver = "postgres"

func pgConnString(dbConfig config.PostgresConfig) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		dbConfig.Username, dbConfig.Password,
		dbConfig.Host, dbConfig.Port,
		dbConfig.Database,
	)
}

// Connect connects to the database with given configuration.
func Connect(cfg config.PostgresConfig) (*sqlx.DB, error) {
	db, err := sqlx.Connect(dbDriver, pgConnString(cfg))
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return db, nil
}
