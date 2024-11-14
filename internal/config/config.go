package config

import (
	"log/slog"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type Config struct {
	HTTPPort uint `envconfig:"HTTP_PORT" default:"8080"`
}

func New() (Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return Config{}, errors.Wrap(err, "load config")
	}

	slog.Info("Config loaded", "config", cfg)
	return cfg, nil
}
