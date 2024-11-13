package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	spec "github.com/Weeping-Willow/entain-task/pkg/oapi/api"
)

func (a *Server) NewRouter() chi.Router {
	r := chi.NewRouter()

	corsMiddleware := cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"*"},
	})

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(corsMiddleware)

	strictHandler := spec.NewStrictHandlerWithOptions(a, nil, spec.StrictHTTPServerOptions{
		RequestErrorHandlerFunc:  responseErrorHandler,
		ResponseErrorHandlerFunc: responseErrorHandler,
	})

	spec.HandlerFromMux(strictHandler, r)

	return r
}

func responseErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	if err := json.NewEncoder(w).Encode(spec.Error{
		Message: err.Error(),
	}); err != nil {
		slog.Error("error while encoding error response", "error", err)
	}
}
