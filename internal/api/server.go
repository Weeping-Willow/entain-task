package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"

	"github.com/Weeping-Willow/entain-task/internal/config"
)

const (
	defaultReadTimeout  = 5 * time.Second
	defaultWriteTimeout = 30 * time.Second
)

type Server struct {
	httpServer *http.Server
	validator  *validator.Validate
}

func New(cfg config.Config) *Server {
	api := &Server{
		validator: validator.New(validator.WithRequiredStructEnabled()),
	}
	api.httpServer = newHTTPServer(cfg.HTTPPort, api.NewRouter())

	return api
}

func newHTTPServer(port uint, handler http.Handler) *http.Server {
	addr := fmt.Sprintf(":%d", port)

	srv := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  defaultReadTimeout,
		WriteTimeout: defaultWriteTimeout,
	}

	return srv
}

func (a *Server) Start() error {
	return errors.Wrap(a.httpServer.ListenAndServe(), "http server start")
}

func (a *Server) Stop() error {
	return errors.Wrap(a.httpServer.Shutdown(context.Background()), "http server shutdown")
}
