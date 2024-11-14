//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/Weeping-Willow/entain-task/internal/api"
	"github.com/Weeping-Willow/entain-task/internal/config"
	"github.com/Weeping-Willow/entain-task/internal/repository"
	"github.com/Weeping-Willow/entain-task/internal/service"
)

func InitializeApp() (*App, func(), error) {
	wire.Build(
		config.New,
		api.New,
		service.NewBalance,
		repository.NewUserStorage,
		config.NewDB,

		NewApp,
	)

	return &App{}, nil, nil
}
