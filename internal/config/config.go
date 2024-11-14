package config

import (
	"log/slog"

	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type Config struct {
	HTTPPort         uint   `envconfig:"HTTP_PORT" default:"8080"`
	PostgresHost     string `envconfig:"POSTGRES_HOST"`
	PostgresPort     uint   `envconfig:"POSTGRES_PORT"`
	PostgresUser     string `envconfig:"POSTGRES_USER"`
	PostgresPassword string `envconfig:"POSTGRES_PASSWORD"`
	PostgresDB       string `envconfig:"POSTGRES_DB"`
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
