package config

import "time"

// PostgresConfig - postgres connection configuration.
type PostgresConfig struct {
	Host     string `yaml:"host"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `env:"POSTGRES_PASSWORD" yaml:"-"`
	Port     int    `yaml:"port"`
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
	Messages *MessagesConfig `yaml:"messages,omitempty"`
	Poll     *PollConfig     `yaml:"poll,omitempty"`
	Token    string          `env:"BOT_TOKEN"           yaml:"-"`
	Verbose  bool            `yaml:"verbose"`
}

// Configuration - configuration of the app.
type Configuration struct {
	Storage StorageConfig `yaml:"storage"`
	Bot     BotConfig     `yaml:"bot"`
}
