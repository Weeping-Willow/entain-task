//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/Weeping-Willow/entain-task/internal/api"
	"github.com/Weeping-Willow/entain-task/internal/config"
)

func InitializeApp() (*App, func(), error) {
	wire.Build(
		config.New,
		api.New,

		NewApp,
	)

	return &App{}, nil, nil
}
