package storage

import (
	"fmt"

	"scriba/internal/config"
)

const storageInMemory = "in-memory"

type Storage interface {
}

type InMemoryStorage struct {
}

func NewStorage(cfg config.StorageConfig) (Storage, error) {
	switch cfg.Type {
	case config.StorageTypeInMemory:
		return new(InMemoryStorage), nil
	}

	return nil, fmt.Errorf("unsupported storage type %s", cfg.Type)
}
