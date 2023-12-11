package storage

import (
	"fmt"
	"net"
	"net/url"
	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // use `postgres` driver
	"github.com/outcatcher/scriba/internal/core/config"
)

const dbDriver = "postgres"

func pgConnString(dbConfig config.PostgresConfig) string {
	pgURL := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(dbConfig.Username, dbConfig.Password),
		Host:     net.JoinHostPort(dbConfig.Host, strconv.Itoa(dbConfig.Port)),
		Path:     dbConfig.Database,
		RawQuery: "sslmode=disable",
	}

	return pgURL.String()
}

// Connect connects to the database with given configuration.
func Connect(cfg config.PostgresConfig) (*sqlx.DB, error) {
	db, err := sqlx.Connect(dbDriver, pgConnString(cfg))
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return db, nil
}
