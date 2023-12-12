package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/caarlos0/env/v10"
	"gopkg.in/yaml.v3"
)

// LoadConfig loads usecases configuration.
func LoadConfig(configPath string, useEnv bool) (*Configuration, error) {
	file, err := os.Open(filepath.Clean(configPath))
	if err != nil {
		return nil, fmt.Errorf("error loading config: failed to open file: %w", err)
	}

	defer func() {
		_ = file.Close() // we're just reading this file so can discard any errors on close
	}()

	config := new(Configuration)

	if err := yaml.NewDecoder(file).Decode(config); err != nil {
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
