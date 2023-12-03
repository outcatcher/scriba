package config

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/caarlos0/env/v10"
	"gopkg.in/yaml.v3"
)

// PostgresConfig - postgres connection configuration.
type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `env:"POSTGRES_PASSWORD" yaml:"-"`
}

// PollConfig - configuration of long polling for telegram.
type PollConfig struct {
	Interval time.Duration `yaml:"interval"`
}

// MessagesConfig - configuration of sent messages.
type MessagesConfig struct {
	Lifetime time.Duration `yaml:"lifetime"`
}

// StorageConfig - storage configuration.
type StorageConfig struct {
	Postgres PostgresConfig `yaml:"postgres"`
}

// BotConfig - configuration of telegram bot.
type BotConfig struct {
	Token    string          `env:"BOT_TOKEN"           yaml:"-"`
	Messages *MessagesConfig `yaml:"messages,omitempty"`
	Poll     *PollConfig     `yaml:"poll,omitempty"`
	Verbose  bool            `yaml:"verbose"`
}

// Configuration - configuration of the app.
type Configuration struct {
	Storage StorageConfig `yaml:"storage"`
	Bot     BotConfig     `yaml:"bot"`
}

// LoadConfig loads usecases configuration.
func LoadConfig(configPath string, useEnv bool) (*Configuration, error) {
	file, err := os.Open(filepath.Clean(configPath))
	if err != nil {
		return nil, fmt.Errorf("error loading config: failed to open file: %w", err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			slog.Error("failed to close file", "error", err)
		}
	}()

	decoder := yaml.NewDecoder(file)

	config := new(Configuration)

	if err := decoder.Decode(config); err != nil {
		return nil, fmt.Errorf("error loading config: error decoding yaml: %w", err)
	}

	if !useEnv {
		return config, nil
	}

	if err := env.Parse(config); err != nil {
		return nil, fmt.Errorf("error loading config: error loading env vars: %w", err)
	}

	return config, nil
}
