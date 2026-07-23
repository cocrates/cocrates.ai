package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/drajin/jsondb/engine"
	"github.com/drajin/jsondb/storage"
)

// managedEngine wraps an engine instance with metadata.
type managedEngine struct {
	engine  engine.Engine
	modTime time.Time
}

// Server represents the HTTP server with Long-lived Engine management.
type Server struct {
	engines   map[string]*managedEngine
	fileLocks map[string]*sync.RWMutex
	mu        sync.RWMutex // protects maps
	addr      string
	root      string
}

// New creates a new Server.
func New(addr string) *Server {
	return &Server{
		engines:   make(map[string]*managedEngine),
		fileLocks: make(map[string]*sync.RWMutex),
		addr:      addr,
	}
}

// SetRoot sets the root directory for file operations.
func (s *Server) SetRoot(root string) {
	s.root = root
}

// getOrCreateEngine returns an existing engine for the file or creates a new one.
func (s *Server) getOrCreateEngine(filePath string) (*managedEngine, error) {
	fullPath := filepath.Join(s.root, filePath)

	// Fast path: read lock
	s.mu.RLock()
	me, ok := s.engines[fullPath]
	s.mu.RUnlock()
	if ok {
		return me, nil
	}

	// Slow path: write lock (double-check)
	s.mu.Lock()
	defer s.mu.Unlock()

	// Double-check after acquiring write lock
	if me, ok = s.engines[fullPath]; ok {
		return me, nil
	}

	// Create new engine
	backend := storage.NewDirectIOBackend()
	eng := engine.New(backend)

	if err := eng.Open(fullPath); err != nil {
		return nil, fmt.Errorf("failed to open engine for %s: %w", filePath, err)
	}

	// Get file modification time
	info, err := os.Stat(fullPath)
	if err != nil {
		eng.Close()
		return nil, fmt.Errorf("failed to stat file %s: %w", filePath, err)
	}

	me = &managedEngine{
		engine:  eng,
		modTime: info.ModTime(),
	}
	s.engines[fullPath] = me

	return me, nil
}

// getFileLock returns an existing lock for the file or creates a new one.
func (s *Server) getFileLock(filePath string) *sync.RWMutex {
	// Fast path: read lock
	s.mu.RLock()
	lock, ok := s.fileLocks[filePath]
	s.mu.RUnlock()
	if ok {
		return lock
	}

	// Slow path: write lock (double-check)
	s.mu.Lock()
	defer s.mu.Unlock()

	// Double-check after acquiring write lock
	if lock, ok = s.fileLocks[filePath]; ok {
		return lock
	}

	lock = &sync.RWMutex{}
	s.fileLocks[filePath] = lock
	return lock
}

// Start starts the HTTP server with graceful shutdown support.
func (s *Server) Start() error {
	mux := http.NewServeMux()

	// File management
	mux.HandleFunc("/api/v1/files", s.handleFiles)
	mux.HandleFunc("/api/v1/files/", s.handleFile)

	// Data operations
	mux.HandleFunc("/api/v1/data/", s.handleData)

	// Status
	mux.HandleFunc("/api/v1/status", s.handleStatus)

	// Create server with timeout
	srv := &http.Server{
		Addr:    s.addr,
		Handler: mux,
	}

	// Channel to receive errors from ListenAndServe
	errChan := make(chan error, 1)

	// Start server in goroutine
	go func() {
		fmt.Printf("Server started on %s\n", s.addr)
		errChan <- srv.ListenAndServe()
	}()

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errChan:
		if err != nil && err != http.ErrServerClosed {
			return fmt.Errorf("server error: %w", err)
		}
	case sig := <-sigChan:
		fmt.Printf("\nReceived signal %v, shutting down gracefully...\n", sig)
	}

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("shutdown error: %w", err)
	}

	// Cleanup engines
	s.shutdown()

	fmt.Println("Server stopped")
	return nil
}

// shutdown closes all open engines and releases resources.
func (s *Server) shutdown() {
	s.mu.Lock()
	defer s.mu.Unlock()

	for path, me := range s.engines {
		fmt.Printf("Closing engine for %s\n", path)
		if err := me.engine.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "Error closing engine for %s: %v\n", path, err)
		}
		delete(s.engines, path)
	}

	// Clear file locks (RWMutex has no explicit close)
	s.fileLocks = make(map[string]*sync.RWMutex)
}

// handleFiles handles file listing and creation.
func (s *Server) handleFiles(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleFileList(w, r)
	case http.MethodPost:
		s.handleFileCreate(w, r)
	default:
		s.writeError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "Method not allowed")
	}
}

// handleFile handles single file operations.
func (s *Server) handleFile(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/v1/files/")
	if path == "" {
		s.handleFiles(w, r)
		return
	}

	switch r.Method {
	case http.MethodDelete:
		s.handleFileDelete(w, r, path)
	case http.MethodHead:
		s.handleFileExists(w, r, path)
	default:
		s.writeError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "Method not allowed")
	}
}

// handleData handles data operations.
func (s *Server) handleData(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/v1/data/")
	if path == "" {
		s.writeError(w, http.StatusBadRequest, "INVALID_PATH", "Path required")
		return
	}

	// Split file path and field query
	parts := strings.SplitN(path, "?", 2)
	filePath := parts[0]

	switch r.Method {
	case http.MethodGet:
		s.handleDataGet(w, r, filePath)
	case http.MethodPut:
		s.handleDataSet(w, r, filePath)
	case http.MethodPatch:
		s.handleDataPatch(w, r, filePath)
	case http.MethodDelete:
		s.handleDataDelete(w, r, filePath)
	case http.MethodPost:
		s.handleDataAdd(w, r, filePath)
	default:
		s.writeError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "Method not allowed")
	}
}

// handleStatus handles server status.
func (s *Server) handleStatus(w http.ResponseWriter, r *http.Request) {
	s.mu.RLock()
	engineCount := len(s.engines)
	s.mu.RUnlock()

	s.writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data": map[string]interface{}{
			"status":  "running",
			"root":    s.root,
			"engines": engineCount,
		},
	})
}

// handleFileList handles file listing.
func (s *Server) handleFileList(w http.ResponseWriter, r *http.Request) {
	dir := r.URL.Query().Get("dir")
	if dir == "" {
		dir = "."
	}

	// Resolve path relative to root
	fullPath := filepath.Join(s.root, dir)

	// Read directory
	entries, err := os.ReadDir(fullPath)
	if err != nil {
		s.writeError(w, http.StatusInternalServerError, "STORAGE_ERROR", fmt.Sprintf("Failed to read directory: %v", err))
		return
	}

	files := []interface{}{}
	directories := []interface{}{}

	for _, entry := range entries {
		if entry.IsDir() {
			directories = append(directories, map[string]interface{}{
				"name": entry.Name(),
				"path": filepath.Join(dir, entry.Name()),
			})
		} else {
			info, err := entry.Info()
			if err != nil {
				continue
			}
			files = append(files, map[string]interface{}{
				"name":     entry.Name(),
				"path":     filepath.Join(dir, entry.Name()),
				"size":     info.Size(),
				"modified": info.ModTime(),
			})
		}
	}

	s.writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data": map[string]interface{}{
			"files":       files,
			"directories": directories,
		},
	})
}

// handleFileCreate handles file creation.
func (s *Server) handleFileCreate(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Path string      `json:"path"`
		Data interface{} `json:"data"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		s.writeError(w, http.StatusBadRequest, "INVALID_JSON", "Invalid JSON body")
		return
	}

	if body.Path == "" {
		s.writeError(w, http.StatusBadRequest, "INVALID_PATH", "Path required")
		return
	}

	// Resolve path relative to root
	fullPath := filepath.Join(s.root, body.Path)

	// Create file with data
	data := body.Data
	if data == nil {
		data = map[string]interface{}{}
	}

	// Ensure directory exists
	dir := filepath.Dir(fullPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		s.writeError(w, http.StatusInternalServerError, "STORAGE_ERROR", fmt.Sprintf("Failed to create directory: %v", err))
		return
	}

	// Write file
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		s.writeError(w, http.StatusInternalServerError, "STORAGE_ERROR", fmt.Sprintf("Failed to marshal data: %v", err))
		return
	}

	if err := os.WriteFile(fullPath, jsonData, 0644); err != nil {
		s.writeError(w, http.StatusInternalServerError, "STORAGE_ERROR", fmt.Sprintf("Failed to write file: %v", err))
		return
	}

	s.writeJSON(w, http.StatusCreated, map[string]interface{}{
		"success": true,
		"data": map[string]interface{}{
			"path": body.Path,
		},
	})
}

// handleFileDelete handles file deletion.
func (s *Server) handleFileDelete(w http.ResponseWriter, r *http.Request, path string) {
	// Resolve path relative to root
	fullPath := filepath.Join(s.root, path)

	// Check if file exists
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		s.writeError(w, http.StatusNotFound, "FILE_NOT_FOUND", fmt.Sprintf("File not found: %s", path))
		return
	}

	// Delete file
	if err := os.Remove(fullPath); err != nil {
		s.writeError(w, http.StatusInternalServerError, "STORAGE_ERROR", fmt.Sprintf("Failed to delete file: %v", err))
		return
	}

	// Remove engine if exists
	s.mu.Lock()
	if me, ok := s.engines[fullPath]; ok {
		me.engine.Close()
		delete(s.engines, fullPath)
	}
	delete(s.fileLocks, path)
	s.mu.Unlock()

	s.writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data": map[string]interface{}{
			"path": path,
		},
	})
}

// handleFileExists handles file existence check.
func (s *Server) handleFileExists(w http.ResponseWriter, r *http.Request, path string) {
	// Resolve path relative to root
	fullPath := filepath.Join(s.root, path)

	// Check if file exists
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// handleDataGet handles data retrieval.
func (s *Server) handleDataGet(w http.ResponseWriter, r *http.Request, filePath string) {
	// Get or create engine (Long-lived)
	me, err := s.getOrCreateEngine(filePath)
	if err != nil {
		s.writeError(w, http.StatusNotFound, "FILE_NOT_FOUND", err.Error())
		return
	}

	// Acquire read lock for file
	lock := s.getFileLock(filePath)
	lock.RLock()
	defer lock.RUnlock()

	// Get field from query params
	field := r.URL.Query().Get("field")
	query := r.URL.Query().Get("query")

	var result interface{}

	if query != "" {
		result, err = me.engine.Get(query)
	} else if field != "" {
		result, err = me.engine.Get(field)
	} else {
		// Get all data
		result, err = me.engine.Get("")
	}

	if err != nil {
		s.writeError(w, http.StatusNotFound, "FIELD_NOT_FOUND", err.Error())
		return
	}

	s.writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    result,
	})
}

// handleDataSet handles data setting.
func (s *Server) handleDataSet(w http.ResponseWriter, r *http.Request, filePath string) {
	field := r.URL.Query().Get("field")
	if field == "" {
		s.writeError(w, http.StatusBadRequest, "INVALID_PATH", "Field parameter required")
		return
	}

	// Parse request body
	var body struct {
		Value interface{} `json:"value"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		s.writeError(w, http.StatusBadRequest, "INVALID_JSON", "Invalid JSON body")
		return
	}

	// Get or create engine (Long-lived)
	me, err := s.getOrCreateEngine(filePath)
	if err != nil {
		s.writeError(w, http.StatusNotFound, "FILE_NOT_FOUND", err.Error())
		return
	}

	// Acquire write lock for file
	lock := s.getFileLock(filePath)
	lock.Lock()
	defer lock.Unlock()

	// Set value
	if err := me.engine.Set(field, body.Value); err != nil {
		s.writeError(w, http.StatusBadRequest, "SET_FAILED", err.Error())
		return
	}

	// Save (write-through)
	if err := me.engine.Save(); err != nil {
		s.writeError(w, http.StatusInternalServerError, "STORAGE_ERROR", "Failed to save")
		return
	}

	// Update modTime
	fullPath := filepath.Join(s.root, filePath)
	if info, err := os.Stat(fullPath); err == nil {
		me.modTime = info.ModTime()
	}

	s.writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data": map[string]interface{}{
			"path":  filePath,
			"field": field,
			"value": body.Value,
		},
	})
}

// handleDataPatch handles data patching (JSON Merge Patch).
func (s *Server) handleDataPatch(w http.ResponseWriter, r *http.Request, filePath string) {
	// Parse request body
	var patch map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&patch); err != nil {
		s.writeError(w, http.StatusBadRequest, "INVALID_JSON", "Invalid JSON body")
		return
	}

	// Get or create engine (Long-lived)
	me, err := s.getOrCreateEngine(filePath)
	if err != nil {
		s.writeError(w, http.StatusNotFound, "FILE_NOT_FOUND", err.Error())
		return
	}

	// Acquire write lock for file
	lock := s.getFileLock(filePath)
	lock.Lock()
	defer lock.Unlock()

	// Get current data
	current, err := me.engine.Get("")
	if err != nil {
		s.writeError(w, http.StatusInternalServerError, "GET_FAILED", err.Error())
		return
	}

	// Merge patch
	currentMap, ok := current.(map[string]interface{})
	if !ok {
		s.writeError(w, http.StatusInternalServerError, "INVALID_DATA", "Data is not an object")
		return
	}

	// Apply patch (JSON Merge Patch)
	for key, value := range patch {
		if value == nil {
			// null means delete
			delete(currentMap, key)
		} else {
			currentMap[key] = value
		}
	}

	// Save (write-through)
	if err := me.engine.Save(); err != nil {
		s.writeError(w, http.StatusInternalServerError, "STORAGE_ERROR", "Failed to save")
		return
	}

	s.writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data": map[string]interface{}{
			"path":    filePath,
			"patched": len(patch),
		},
	})
}

// handleDataDelete handles data deletion.
func (s *Server) handleDataDelete(w http.ResponseWriter, r *http.Request, filePath string) {
	field := r.URL.Query().Get("field")
	if field == "" {
		s.writeError(w, http.StatusBadRequest, "INVALID_PATH", "Field parameter required")
		return
	}

	// Get or create engine (Long-lived)
	me, err := s.getOrCreateEngine(filePath)
	if err != nil {
		s.writeError(w, http.StatusNotFound, "FILE_NOT_FOUND", err.Error())
		return
	}

	// Acquire write lock for file
	lock := s.getFileLock(filePath)
	lock.Lock()
	defer lock.Unlock()

	// Delete field
	if err := me.engine.Delete(field); err != nil {
		s.writeError(w, http.StatusBadRequest, "DELETE_FAILED", err.Error())
		return
	}

	// Save (write-through)
	if err := me.engine.Save(); err != nil {
		s.writeError(w, http.StatusInternalServerError, "STORAGE_ERROR", "Failed to save")
		return
	}

	s.writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data": map[string]interface{}{
			"path":  filePath,
			"field": field,
		},
	})
}

// handleDataAdd handles data addition (arrays).
func (s *Server) handleDataAdd(w http.ResponseWriter, r *http.Request, filePath string) {
	// Parse request body
	var body struct {
		Value interface{} `json:"value"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		s.writeError(w, http.StatusBadRequest, "INVALID_JSON", "Invalid JSON body")
		return
	}

	// TODO: Implement data addition to arrays
	s.writeError(w, http.StatusNotImplemented, "NOT_IMPLEMENTED", "Data addition not implemented yet")
}

// writeJSON writes a JSON response.
func (s *Server) writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// writeError writes an error response.
func (s *Server) writeError(w http.ResponseWriter, status int, code, message string) {
	s.writeJSON(w, status, map[string]interface{}{
		"success": false,
		"error": map[string]interface{}{
			"code":    code,
			"message": message,
		},
	})
}

// StartWithRoot starts the server with a root directory.
func StartWithRoot(addr, root string) error {
	srv := New(addr)
	srv.SetRoot(root)
	return srv.Start()
}
