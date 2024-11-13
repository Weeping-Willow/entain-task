package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/Weeping-Willow/entain-task/internal/api"
)

type App struct {
	api *api.Server
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo, AddSource: true}))

	slog.SetDefault(logger)

	app, cleanup, err := InitializeApp()
	if err != nil {
		panic(err)
	}

	defer cleanup()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-c
		app.Stop()
	}()

	app.Start()
}

func NewApp(apiServer *api.Server) *App {
	return &App{
		api: apiServer,
	}
}

func (app *App) Start() {
	err := app.api.Start()
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
}

func (app *App) Stop() {
	err := app.api.Stop()
	if err != nil {
		slog.Error(err.Error())
	}
}
