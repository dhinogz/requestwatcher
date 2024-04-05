package server

import (
	"log/slog"
	"net/http"

	"github.com/dhinogz/requestwatcher/db"
	"github.com/dhinogz/requestwatcher/manager"
)

const (
	DefaultPort = ":4000"
)

type Server struct {
	Store   *db.Queries
	Port    string
	Logger  *slog.Logger
	Manager *manager.Manager
}

func New(options ...func(*Server)) *Server {
	svr := &Server{}
	svr.Port = DefaultPort
	svr.Logger = slog.Default()
	for _, o := range options {
		o(svr)
	}
	return svr
}

func WithPort(port string) func(*Server) {
	if port == "" {
		port = DefaultPort
	}
	return func(s *Server) {
		s.Port = port
	}
}

func WithLogger(logger *slog.Logger) func(*Server) {
	if logger == nil {
		logger = slog.Default()
	}
	return func(s *Server) {
		s.Logger = logger
	}
}

func WithStore(store *db.Queries) func(*Server) {
	return func(s *Server) {
		s.Store = store
	}
}

func WithManager(manager *manager.Manager) func(*Server) {
	return func(s *Server) {
		s.Manager = manager
	}
}

func (s *Server) ListenAndServe() error {
	r := s.routes()

	return http.ListenAndServe(s.Port, r)
}
