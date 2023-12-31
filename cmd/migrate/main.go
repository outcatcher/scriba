/*
Package main is a script to run migrations.
*/
package main

import (
	"context"
	"flag"
	"log"

	"github.com/outcatcher/scriba/internal/core/config"
	"github.com/outcatcher/scriba/internal/core/storage"
)

var (
	configPath     = flag.String("config", "./config/config.yaml", "Configuration file to be used")
	migrationsPath = flag.String("migrations", "./internal/migrations", "Path to migrations")
)

func main() {
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath, true)
	if err != nil {
		log.Fatalln("failed to load configuration", err)
	}

	ctx := context.Background()

	if err := storage.ApplyMigrations(ctx, cfg.Storage.Postgres, *migrationsPath); err != nil {
		log.Fatalln("failed to apply migrations", err)
	}
}
