package storage

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/outcatcher/scriba/internal/core/config"
	"github.com/pressly/goose/v3"
)

// ApplyMigrations applies all available migrations.
func ApplyMigrations(ctx context.Context, cfg config.PostgresConfig, migrationsDir string) error {
	if err := goose.SetDialect(dbDriver); err != nil {
		return fmt.Errorf("error selecting dialect for migrations: %w", err)
	}

	db, err := Connect(cfg)
	if err != nil {
		return err
	}

	absPath, err := filepath.Abs(migrationsDir)
	if err != nil {
		return fmt.Errorf("error getting abs path for %s: %w", migrationsDir, err)
	}

	if err := goose.UpContext(ctx, db.DB, absPath, goose.WithAllowMissing()); err != nil {
		return fmt.Errorf("error applying migrations: %w", err)
	}

	return nil
}
