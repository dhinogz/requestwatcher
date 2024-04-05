package server

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/dhinogz/requestwatcher/parser"
	"github.com/dhinogz/requestwatcher/rand"
	"github.com/dhinogz/requestwatcher/views"

	"github.com/go-chi/chi/v5"
)

// GET /
// We show the home page with info of how it works and button to generate a watcher. Calls Post /watcher
func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	err := views.Index().Render(ctx, w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("something bad happened"))
	}
}

// POST /watcher
// Creates a watcher and sends HTML fragment with URL info and
// Server side event to call /watcher/{hash}
func (s *Server) handleWatcher(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	watcherID := rand.GenerateShortenedURL()
	_, err := s.Store.CreateWatcher(ctx, watcherID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	uri := fmt.Sprintf("/watcher/%s", watcherID)

	w.Header().Set("HX-Replace-URL", uri)
	err = views.WatcherHTMXFragment(r.Host, watcherID).Render(ctx, w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("something bad happened"))
	}
}

// GET /watcher/:watcherID
// Queries database with watcherID
func (s *Server) handleWatcherRequestsPage(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	watcherID := chi.URLParam(r, "watcherID")
	wRequests, err := s.Store.GetWatcherRequests(ctx, watcherID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = views.WatcherPage(r.Host, watcherID, wRequests).Render(ctx, w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("something bad happened"))
	}
}

// GET /:watcherID
// Stores request information in database, triggers server side event
func (s *Server) handleWatcherRequest(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	hashUri := chi.URLParam(r, "watcherID")
	watcherID, err := s.Store.GetWatcher(ctx, hashUri)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("watcher not found"))
		return
	}

	createRequestParams, err := parser.Request(r, watcherID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		s.Logger.Error("parse request", "err", err)
		return
	}

	wr, err := s.Store.CreateRequest(ctx, createRequestParams)

	resp := new(bytes.Buffer)
	views.WatcherCard(wr).Render(ctx, resp)
	s.Manager.SendMessage(hashUri, resp.String())

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Created"))
}

func headerValue(r *http.Request, key string) string {
	v, _ := r.Header[key]
	return fmt.Sprintf("%+v", v)
}

// GET /events/:watcherID
// Adds client in channel manager, which organizes server side events
// When connection is closed and done, we remove channel from manager
func (s *Server) handleEvent(w http.ResponseWriter, r *http.Request) {
	watcherID := chi.URLParam(r, "watcherID")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Transfer-Encoding", "chunked")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	ch := make(chan string)
	defer close(ch)

	s.Manager.AddClient(watcherID, ch)
	defer s.Manager.RemoveClient(watcherID)

	for {
		select {
		case <-r.Context().Done():
			return
		case data := <-ch:
			fmt.Fprintf(w, "data: %s\n\n", data)
			w.(http.Flusher).Flush()
		}
	}
}
