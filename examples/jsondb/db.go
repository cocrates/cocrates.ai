package jsondb

import (
	"fmt"
	"os"
	"path/filepath"
)

// DB is a local JSON document store.
// It provides path-addressable CRUD operations with per-file locking.
type DB struct {
	root       string
	schemaRoot string
	port       int

	locks lockMap
}

// Open opens or creates a jsondb at the given root directory.
// If the directory does not exist, it is created.
func Open(root string, opts ...Option) (*DB, error) {
	abs, err := filepath.Abs(root)
	if err != nil {
		return nil, fmt.Errorf("jsondb: resolve root path: %w", err)
	}

	if err := os.MkdirAll(abs, 0755); err != nil {
		return nil, fmt.Errorf("jsondb: create root directory: %w", err)
	}

	db := &DB{
		root: abs,
	}

	// Apply options.
	for _, opt := range opts {
		opt(db)
	}

	// Default schema root.
	if db.schemaRoot == "" {
		db.schemaRoot = filepath.Join(abs, "schema")
	}
	if err := os.MkdirAll(db.schemaRoot, 0755); err != nil {
		return nil, fmt.Errorf("jsondb: create schema directory: %w", err)
	}

	return db, nil
}

// Close closes the database and releases resources.
func (db *DB) Close() error {
	// Currently no long-lived resources to release.
	// In the future this may flush buffers or clean up watchers.
	return nil
}

// Root returns the absolute path to the database root directory.
func (db *DB) Root() string {
	return db.root
}

// SchemaRoot returns the absolute path to the schema root directory.
func (db *DB) SchemaRoot() string {
	return db.schemaRoot
}

// filePath returns the absolute filesystem path for a data path.
func (db *DB) filePath(dataPath string) string {
	return filepath.Join(db.root, filepath.FromSlash(dataPath)) + ".json"
}

// schemaFilePath returns the absolute filesystem path for a schema path.
func (db *DB) schemaFilePath(schemaPath string) string {
	return filepath.Join(db.schemaRoot, filepath.FromSlash(schemaPath)) + ".json"
}
