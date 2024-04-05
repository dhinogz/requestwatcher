package server

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/dhinogz/requestwatcher/db"
	"github.com/dhinogz/requestwatcher/rand"
	"github.com/dhinogz/requestwatcher/views"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

// GET / We show the home page with info of how it works and button to generate a watcher. Calls Post /watcher
// Calls CreateWatcher query
func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	err := views.Index().Render(ctx, w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("something bad happened"))
	}
}

// POST /watcher -> Creates a watcher and sends HTML fragment with URL info and
// server side event to call /watcher/{hash}
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

// GET /watcher/:watcherID -> queries database with watcherID
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

// GET /:watcherID -> stores request information in database, triggers server side event
func (s *Server) handleWatcherRequest(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	hashUri := chi.URLParam(r, "watcherID")
	watcherID, err := s.Store.GetWatcher(ctx, hashUri)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("watcher not found"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	wr, err := s.Store.CreateRequest(ctx, db.CreateRequestParams{
		Method: r.Method,
		Ip:     r.RemoteAddr,
		Url:    r.RequestURI,
		Timestamp: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
		Accept:                  headerValue(r, "Accept"),
		AcceptEncoding:          headerValue(r, "Accept-Encoding"),
		AcceptLanguage:          headerValue(r, "Accept-Value"),
		CfConnectingIp:          headerValue(r, "Cf-Connecting-Ip"),
		CfIpcountry:             headerValue(r, "Cf-Ipcountry"),
		CfRay:                   headerValue(r, "Cf-Ray"),
		CfVisitor:               headerValue(r, "Cf-Visitor"),
		Connection:              headerValue(r, "Connection"),
		Host:                    headerValue(r, "Host"),
		Priority:                headerValue(r, "Priority"),
		SecChUa:                 headerValue(r, "Sec-Ch-Ua"),
		SecChUaMobile:           headerValue(r, "Sec-Ch-Ua-Mobile"),
		SecChUaPlatform:         headerValue(r, "Sec-Ch-Ua-Platform"),
		SecFetchDest:            headerValue(r, "Sec-Fetch-Dest"),
		SecFetchMode:            headerValue(r, "Sec-Fetch-Mode"),
		SecFetchSite:            headerValue(r, "Sec-Fetch-Site"),
		SecFetchUser:            headerValue(r, "Sec-Fetch-User"),
		UpgradeInsecureRequests: headerValue(r, "Upgrade-Insecure-Requests"),
		UserAgent:               headerValue(r, "User-Agent"),
		XForwardedProto:         headerValue(r, "X-Forwarded-Proto"),
		XRealIp:                 headerValue(r, "X-Real-Ip"),
		Body:                    string(body),
		WatcherID:               watcherID,
	})
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
