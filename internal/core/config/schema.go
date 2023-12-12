package config

import "time"

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
