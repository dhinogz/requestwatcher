package server

import (
	"github.com/go-chi/chi/v5"
	slogchi "github.com/samber/slog-chi"
)

func (s *Server) routes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(slogchi.New(s.Logger))

	r.Get("/", s.handleIndex)

	r.Get("/{watcherID}", s.handleWatcherRequest)

	r.Post("/watcher", s.handleWatcher)
	r.Get("/watcher/{watcherID}", s.handleWatcherRequestsPage)

	r.Get("/events/{watcherID}", s.handleEvent)

	return r
}
