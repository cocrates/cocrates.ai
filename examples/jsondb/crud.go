package jsondb

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// dataDoc represents a JSON document with an optional _schema field.
// Used internally to extract the schema reference before storing.
type dataDoc struct {
	Schema string `json:"_schema,omitempty"`
	raw    []byte // full raw JSON
}

// Set stores data at the given path as a JSON file.
// If the data contains a _schema field, it validates against that schema.
func (db *DB) Set(path string, data any) error {
	// Marshal first so we can inspect _schema.
	raw, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("jsondb: marshal data: %w", err)
	}

	// Extract _schema reference for validation.
	var doc dataDoc
	if err := json.Unmarshal(raw, &doc); err != nil {
		return fmt.Errorf("jsondb: parse data for schema check: %w", err)
	}
	doc.raw = raw

	// Acquire write lock.
	_, fl := db.wLock(path)
	defer db.wUnlock(path, fl)

	// Validate if schema is referenced.
	if doc.Schema != "" {
		if err := db.validate(doc.Schema, doc.raw); err != nil {
			return err
		}
	}

	// Ensure directory exists.
	fp := db.filePath(path)
	dir := filepath.Dir(fp)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("jsondb: create directory: %w", err)
	}

	// Write file.
	if err := os.WriteFile(fp, doc.raw, 0644); err != nil {
		return fmt.Errorf("jsondb: write file: %w", err)
	}

	return nil
}

// Get retrieves the JSON document at path and unmarshals it into v.
func (db *DB) Get(path string, v any) error {
	_, fl := db.rLock(path)
	defer db.rUnlock(path, fl)

	fp := db.filePath(path)
	data, err := os.ReadFile(fp)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("jsondb: document not found: %s", path)
		}
		return fmt.Errorf("jsondb: read file: %w", err)
	}

	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("jsondb: unmarshal data: %w", err)
	}

	return nil
}

// Delete removes the JSON file at path.
func (db *DB) Delete(path string) error {
	_, fl := db.wLock(path)
	defer db.wUnlock(path, fl)

	fp := db.filePath(path)
	if err := os.Remove(fp); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("jsondb: document not found: %s", path)
		}
		return fmt.Errorf("jsondb: delete file: %w", err)
	}

	return nil
}

// Exists returns true if a document exists at the given path.
func (db *DB) Exists(path string) (bool, error) {
	_, fl := db.rLock(path)
	defer db.rUnlock(path, fl)

	fp := db.filePath(path)
	_, err := os.Stat(fp)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, fmt.Errorf("jsondb: check file: %w", err)
}
