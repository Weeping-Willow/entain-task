package config

import (
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

	return cfg, nil
}
