package config

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func NewDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", newPostgresDBConnectionString(cfg))
	if err != nil {
		return nil, errors.Wrap(err, "connect to postgres db")
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "ping to postgres db")
	}

	return db, nil
}

func newPostgresDBConnectionString(cfg Config) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.PostgresHost,
		cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)
}
