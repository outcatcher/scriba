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

type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"user"`
	Password string `env:"POSTGRES_PASSWORD"`
}

type PollConfig struct {
	Interval time.Duration `yaml:"interval"`
}

type MessagesConfig struct {
	Lifetime time.Duration `yaml:"lifetime"`
}

type StorageConfig struct {
	Postgres PostgresConfig `yaml:"postgres"`
}

type BotMethod string

type BotConfig struct {
	Token    string          `env:"BOT_TOKEN"`
	Method   BotMethod       `yaml:"method"`
	Messages *MessagesConfig `yaml:"messages"`
	Poll     *PollConfig     `yaml:"poll,omitempty"`
	Verbose  bool            `yaml:"verbose"`
}

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
