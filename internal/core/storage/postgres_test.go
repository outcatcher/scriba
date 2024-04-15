package storage

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/outcatcher/scriba/internal/core/config"
	"github.com/stretchr/testify/require"
)

const maxPort = 0xffff

func TestPgConnectionString(t *testing.T) {
	t.Parallel()

	dbConfig := config.PostgresConfig{
		Host:     "postgre-test",
		Port:     rand.Intn(maxPort),
		Database: "griuhlgn4l",
		Username: "t43q0jvfpkml;d",
		Password: "124431234dfs",
	}

	expectedStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		dbConfig.Username, dbConfig.Password,
		dbConfig.Host, dbConfig.Port,
		dbConfig.Database,
	)

	require.Equal(t, expectedStr, pgConnString(dbConfig))
}
