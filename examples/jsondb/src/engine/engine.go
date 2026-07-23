package engine

import (
	"fmt"
	"sync"

	"github.com/drajin/jsondb/engine/query"
	"github.com/drajin/jsondb/storage"
)

// Engine defines the interface for jsondb operations.
type Engine interface {
	// Open opens a JSON file.
	Open(path string) error
	// Get retrieves a field value.
	Get(fieldPath string) (interface{}, error)
	// Set sets a field value.
	Set(fieldPath string, value interface{}) error
	// Delete removes a field.
	Delete(fieldPath string) error
	// Save saves the current state to disk.
	Save() error
	// Close closes the engine and releases resources.
	Close() error
}

// jsonEngine implements the Engine interface.
type jsonEngine struct {
	storage  storage.StorageBackend
	executor *query.Executor
	filePath string
	data     map[string]interface{}
	mu       sync.RWMutex
	open     bool
}

// New creates a new jsonEngine with the given storage backend.
func New(backend storage.StorageBackend) Engine {
	return &jsonEngine{
		storage:  backend,
		executor: query.NewExecutor(),
	}
}

// Open opens a JSON file.
func (e *jsonEngine) Open(path string) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	if err := e.storage.Open(path); err != nil {
		return fmt.Errorf("failed to open storage: %w", err)
	}

	// Read initial data
	data, err := e.storage.Read()
	if err != nil {
		e.storage.Close()
		return fmt.Errorf("failed to read data: %w", err)
	}

	e.filePath = path
	e.data = data
	e.open = true
	return nil
}

// Get retrieves a field value.
func (e *jsonEngine) Get(fieldPath string) (interface{}, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()

	if !e.open {
		return nil, ErrNotOpen
	}

	// Empty path returns all data
	if fieldPath == "" {
		return e.data, nil
	}

	parsed, err := query.Parse(fieldPath)
	if err != nil {
		return nil, fmt.Errorf("failed to parse path: %w", err)
	}

	return e.executor.Get(e.data, parsed)
}

// Set sets a field value.
func (e *jsonEngine) Set(fieldPath string, value interface{}) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	if !e.open {
		return ErrNotOpen
	}

	parsed, err := query.Parse(fieldPath)
	if err != nil {
		return fmt.Errorf("failed to parse path: %w", err)
	}

	return e.executor.Set(e.data, parsed, value)
}

// Delete removes a field.
func (e *jsonEngine) Delete(fieldPath string) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	if !e.open {
		return ErrNotOpen
	}

	parsed, err := query.Parse(fieldPath)
	if err != nil {
		return fmt.Errorf("failed to parse path: %w", err)
	}

	return e.executor.Delete(e.data, parsed)
}

// Save saves the current state to disk.
func (e *jsonEngine) Save() error {
	e.mu.Lock()
	defer e.mu.Unlock()

	if !e.open {
		return ErrNotOpen
	}

	return e.storage.Write(e.data)
}

// Close closes the engine and releases resources.
func (e *jsonEngine) Close() error {
	e.mu.Lock()
	defer e.mu.Unlock()

	if !e.open {
		return nil
	}

	if err := e.storage.Close(); err != nil {
		return fmt.Errorf("failed to close storage: %w", err)
	}

	e.open = false
	e.data = nil
	return nil
}
