package router

import (
	"foveo/internal/handler"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func NewRouter(handler *handler.Handler) http.Handler {
	r := chi.NewRouter()

	// middleware по необходимости
	// r.Use(middleware.Logger)

	// user-related routes
	r.Route("/api/users", func(r chi.Router) {
		r.Post("/", handler.User.Create)
		r.Get("/{id}", handler.User.GetByID)
		r.Put("/{id}", handler.User.Update)
		r.Delete("/{id}", handler.User.Delete)
	})

	// healthcheck
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	return r
}
