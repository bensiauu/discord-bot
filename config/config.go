package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	cfgPath = "CONFIG_PATH"
)

type Config struct {
	Authorization Auth `yaml:"authorization"`
}

type Auth struct {
	Token string `yaml:"token"`
}

func ParseConfig() (*Config, error) {
	buf, err := os.ReadFile(os.Getenv(cfgPath))
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	cfg := &Config{}

	err = yaml.Unmarshal(buf, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return cfg, nil
}
