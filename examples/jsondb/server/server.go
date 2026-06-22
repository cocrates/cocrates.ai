package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/cocrates/jsondb"
)

// Config holds REST server settings.
type Config struct {
	Port       int
	DBRoot     string
	SchemaRoot string
}

// Run starts the REST API server and blocks until SIGINT or SIGTERM.
func Run(cfg Config) error {
	opts := []jsondb.Option{jsondb.WithPort(cfg.Port)}
	if cfg.SchemaRoot != "" {
		opts = append(opts, jsondb.WithSchemaRoot(cfg.SchemaRoot))
	}

	db, err := jsondb.Open(cfg.DBRoot, opts...)
	if err != nil {
		return fmt.Errorf("open database: %w", err)
	}
	defer db.Close()

	srv := &handler{db: db}
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", srv.handleHealth)
	mux.HandleFunc("GET /schema/", srv.handleGetSchema)
	mux.HandleFunc("PUT /schema/", srv.handleSetSchema)
	mux.HandleFunc("DELETE /schema/", srv.handleDeleteSchema)
	mux.HandleFunc("GET /data", srv.handleList)
	mux.HandleFunc("GET /data/", srv.handleGet)
	mux.HandleFunc("PUT /data/", srv.handleSet)
	mux.HandleFunc("DELETE /data/", srv.handleDelete)

	addr := fmt.Sprintf(":%d", cfg.Port)
	httpServer := &http.Server{
		Addr:    addr,
		Handler: withLogging(mux),
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Printf("jsondb listening on %s (dbroot: %s)", addr, db.Root())
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	<-ctx.Done()
	log.Print("shutting down gracefully...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return httpServer.Shutdown(shutdownCtx)
}

type handler struct {
	db *jsondb.DB
}

type response struct {
	OK    bool   `json:"ok"`
	Data  any    `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

func writeJSON(w http.ResponseWriter, status int, resp response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}

func ok(w http.ResponseWriter, data any) {
	writeJSON(w, http.StatusOK, response{OK: true, Data: data})
}

func errResp(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, response{OK: false, Error: msg})
}

func withLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}

func (s *handler) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func (s *handler) handleGet(w http.ResponseWriter, r *http.Request) {
	docPath := strings.TrimPrefix(r.URL.Path, "/data/")
	jsonPath := r.URL.Query().Get("path")

	if jsonPath != "" {
		result, err := s.db.GetPath(docPath, jsonPath)
		if err != nil {
			errResp(w, http.StatusNotFound, err.Error())
			return
		}
		ok(w, result.Value())
		return
	}

	var data any
	if err := s.db.Get(docPath, &data); err != nil {
		errResp(w, http.StatusNotFound, err.Error())
		return
	}
	ok(w, data)
}

func (s *handler) handleSet(w http.ResponseWriter, r *http.Request) {
	docPath := strings.TrimPrefix(r.URL.Path, "/data/")
	jsonPath := r.URL.Query().Get("path")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		errResp(w, http.StatusBadRequest, fmt.Sprintf("read body: %v", err))
		return
	}

	if jsonPath != "" {
		var value any
		if err := json.Unmarshal(body, &value); err != nil {
			value = string(body)
		}
		if err := s.db.SetPath(docPath, jsonPath, value); err != nil {
			errResp(w, http.StatusInternalServerError, err.Error())
			return
		}
		ok(w, nil)
		return
	}

	var data any
	if err := json.Unmarshal(body, &data); err != nil {
		errResp(w, http.StatusBadRequest, fmt.Sprintf("invalid JSON: %v", err))
		return
	}
	if err := s.db.Set(docPath, data); err != nil {
		errResp(w, http.StatusInternalServerError, err.Error())
		return
	}
	ok(w, nil)
}

func (s *handler) handleDelete(w http.ResponseWriter, r *http.Request) {
	docPath := strings.TrimPrefix(r.URL.Path, "/data/")
	jsonPath := r.URL.Query().Get("path")

	if jsonPath != "" {
		if err := s.db.DeletePath(docPath, jsonPath); err != nil {
			errResp(w, http.StatusNotFound, err.Error())
			return
		}
		ok(w, nil)
		return
	}

	if err := s.db.Delete(docPath); err != nil {
		errResp(w, http.StatusNotFound, err.Error())
		return
	}
	ok(w, nil)
}

func (s *handler) handleList(w http.ResponseWriter, r *http.Request) {
	prefix := r.URL.Query().Get("prefix")
	paths, err := s.db.List(prefix)
	if err != nil {
		errResp(w, http.StatusInternalServerError, err.Error())
		return
	}
	if paths == nil {
		paths = []string{}
	}
	ok(w, paths)
}

func (s *handler) handleGetSchema(w http.ResponseWriter, r *http.Request) {
	schemaPath := strings.TrimPrefix(r.URL.Path, "/schema/")
	var data any
	if err := s.db.GetSchema(schemaPath, &data); err != nil {
		errResp(w, http.StatusNotFound, err.Error())
		return
	}
	ok(w, data)
}

func (s *handler) handleSetSchema(w http.ResponseWriter, r *http.Request) {
	schemaPath := strings.TrimPrefix(r.URL.Path, "/schema/")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		errResp(w, http.StatusBadRequest, fmt.Sprintf("read body: %v", err))
		return
	}

	var schema any
	if err := json.Unmarshal(body, &schema); err != nil {
		errResp(w, http.StatusBadRequest, fmt.Sprintf("invalid JSON Schema: %v", err))
		return
	}

	if err := s.db.SetSchema(schemaPath, schema); err != nil {
		errResp(w, http.StatusInternalServerError, err.Error())
		return
	}
	ok(w, nil)
}

func (s *handler) handleDeleteSchema(w http.ResponseWriter, r *http.Request) {
	schemaPath := strings.TrimPrefix(r.URL.Path, "/schema/")
	if err := s.db.DeleteSchema(schemaPath); err != nil {
		errResp(w, http.StatusNotFound, err.Error())
		return
	}
	ok(w, nil)
}
