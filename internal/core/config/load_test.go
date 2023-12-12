package config

import (
	"math/rand"
	"os"
	"path/filepath"
	"testing"
	"text/template"
	"time"

	"github.com/stretchr/testify/require"
)

const configTemplate = `
storage:
  postgres:
    host: {{ .Storage.Postgres.Host }}
    port: {{ .Storage.Postgres.Port }}
    database: {{ .Storage.Postgres.Database }}
    username: {{ .Storage.Postgres.Username }}
    password: {{ .Storage.Postgres.Password }}

bot:
  poll:
    interval: {{ .Bot.Poll.Interval }}
  webhook:
  messages:
    lifetime: {{ .Bot.Messages.Lifetime }}
  verbose: {{ .Bot.Verbose }}
`

func prepareTestConfig(t *testing.T, data *Configuration) string {
	t.Helper()

	tmpl, err := template.New("config").Parse(configTemplate)
	require.NoError(t, err)

	tmpDirPath := t.TempDir()
	tmpFilePath := filepath.Join(tmpDirPath, "config.yaml")

	file, err := os.Create(tmpFilePath)
	require.NoError(t, err)

	require.NoError(t, tmpl.Execute(file, data))

	return tmpFilePath
}

func TestLoadConfig_ok(t *testing.T) {
	t.Parallel()

	expectedConfig := &Configuration{
		Storage: StorageConfig{
			Postgres: PostgresConfig{
				Host:     "http://test-host.com",
				Port:     rand.Intn(0xffff),
				Database: "griuhlgn4l",
				Username: "t43q0jvfpkml;d",
				Password: "", // defined via env vars
			},
		},
		Bot: BotConfig{
			Token: "", // defined via env vars
			Messages: &MessagesConfig{
				Lifetime: time.Minute,
			},
			Poll: &PollConfig{
				Interval: time.Second,
			},
			Verbose: true,
		},
	}

	filePath := prepareTestConfig(t, expectedConfig)

	config, err := LoadConfig(filePath, false)
	require.NoError(t, err)
	require.Equal(t, expectedConfig, config)
}

func TestLoadConfig_tokenIgnored(t *testing.T) {
	t.Parallel()

	expectedConfig := &Configuration{
		Bot: BotConfig{
			Token:    "i34ojrtp342jt", // defined via env vars
			Messages: &MessagesConfig{},
			Poll:     &PollConfig{},
		},
	}

	filePath := prepareTestConfig(t, expectedConfig)

	config, err := LoadConfig(filePath, false)
	require.NoError(t, err)
	require.NotEqual(t, expectedConfig, config)
}

func TestLoadConfig_dbPasswordIgnored(t *testing.T) {
	t.Parallel()

	expectedConfig := &Configuration{
		Storage: StorageConfig{
			Postgres: PostgresConfig{
				Password: "fads234", // defined via env vars
			},
		},
		Bot: BotConfig{
			Messages: &MessagesConfig{},
			Poll:     &PollConfig{},
		},
	}

	filePath := prepareTestConfig(t, expectedConfig)

	config, err := LoadConfig(filePath, false)
	require.NoError(t, err)
	require.NotEqual(t, expectedConfig, config)
}

func TestLoadConfig_envVars(t *testing.T) {
	expectedPassword := "fads234"
	expectedToken := "tokekekf[ps"

	t.Setenv("POSTGRES_PASSWORD", expectedPassword)
	t.Setenv("BOT_TOKEN", expectedToken)

	expectedConfig := &Configuration{
		Storage: StorageConfig{
			Postgres: PostgresConfig{
				Password: expectedPassword, // defined via env vars
			},
		},
		Bot: BotConfig{
			Token:    expectedToken, // defined via env vars
			Messages: &MessagesConfig{},
			Poll:     &PollConfig{},
		},
	}

	filePath := prepareTestConfig(t, expectedConfig)

	t.Run("not using env", func(t *testing.T) {
		t.Parallel()

		config, err := LoadConfig(filePath, false)
		require.NoError(t, err)
		require.NotEqual(t, expectedConfig, config)
	})

	t.Run("using env", func(t *testing.T) {
		t.Parallel()

		config, err := LoadConfig(filePath, true)
		require.NoError(t, err)
		require.Equal(t, expectedConfig, config)
	})
}
