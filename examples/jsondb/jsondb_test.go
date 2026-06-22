package jsondb_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cocrates/jsondb"
)

// Smoke test for basic CRUD operations.
func TestSmoke(t *testing.T) {
	dir := t.TempDir()

	db, err := jsondb.Open(filepath.Join(dir, "db"))
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// Set a document.
	err = db.Set("overview", map[string]any{
		"title": "My Project",
		"version": 1,
	})
	if err != nil {
		t.Fatal(err)
	}

	// Verify file exists.
	if _, err := os.Stat(filepath.Join(dir, "db", "overview.json")); err != nil {
		t.Fatal("file not created:", err)
	}

	// Get the document.
	var result map[string]any
	err = db.Get("overview", &result)
	if err != nil {
		t.Fatal(err)
	}
	if result["title"] != "My Project" {
		t.Fatalf("expected 'My Project', got %v", result["title"])
	}

	// Exists.
	ok, err := db.Exists("overview")
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("exists should return true")
	}

	// GetPath.
	val, err := db.GetPath("overview", "title")
	if err != nil {
		t.Fatal(err)
	}
	if val.String() != "My Project" {
		t.Fatalf("expected 'My Project', got %s", val.String())
	}

	// SetPath.
	err = db.SetPath("overview", "version", 2)
	if err != nil {
		t.Fatal(err)
	}
	val, err = db.GetPath("overview", "version")
	if err != nil {
		t.Fatal(err)
	}
	if val.Int() != 2 {
		t.Fatalf("expected 2, got %d", val.Int())
	}

	// DeletePath.
	err = db.DeletePath("overview", "version")
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.GetPath("overview", "version")
	if err == nil {
		t.Fatal("expected error for deleted path")
	}

	// List.
	db.Set("episode/e1", map[string]any{"title": "E1"})
	db.Set("episode/e2", map[string]any{"title": "E2"})
	paths, err := db.List("episode/")
	if err != nil {
		t.Fatal(err)
	}
	if len(paths) != 2 {
		t.Fatalf("expected 2 paths, got %d: %v", len(paths), paths)
	}

	// Delete.
	err = db.Delete("overview")
	if err != nil {
		t.Fatal(err)
	}
	ok, err = db.Exists("overview")
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatal("exists should return false after delete")
	}
}

// Test schema validation.
func TestSchema(t *testing.T) {
	dir := t.TempDir()

	schemaRoot := filepath.Join(dir, "schemas")
	db, err := jsondb.Open(filepath.Join(dir, "db"), jsondb.WithSchemaRoot(schemaRoot))
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// Set schema.
	schema := map[string]any{
		"type": "object",
		"properties": map[string]any{
			"title": map[string]any{
				"type": "string",
			},
		},
		"required":             []string{"title"},
		"additionalProperties": false,
	}
	err = db.SetSchema("/blog/episode", schema)
	if err != nil {
		t.Fatal(err)
	}

	// Valid document.
	err = db.Set("episode/e1", map[string]any{
		"_schema": "/blog/episode",
		"title":   "Valid Episode",
	})
	if err != nil {
		t.Fatal("expected valid:", err)
	}

	// Invalid document (missing required field, extra field).
	err = db.Set("episode/e2", map[string]any{
		"_schema": "/blog/episode",
		"name":    "No Title",
	})
	if err == nil {
		t.Fatal("expected validation error for missing title")
	}

	// Schemaless document (no _schema field).
	err = db.Set("freeform/data", map[string]any{
		"anything": "goes",
	})
	if err != nil {
		t.Fatal("expected schemaless write to succeed:", err)
	}
}

// Test concurrency safety.
func TestConcurrency(t *testing.T) {
	dir := t.TempDir()

	db, err := jsondb.Open(filepath.Join(dir, "db"))
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	done := make(chan bool, 20)

	// Concurrent writes to different files.
	for i := 0; i < 10; i++ {
		go func(n int) {
			path := "concurrent/doc"
			if n%2 == 0 {
				path = "concurrent/other"
			}
			err := db.Set(path, map[string]any{
				"index": n,
			})
			if err != nil {
				t.Errorf("concurrent write %d: %v", n, err)
			}
			done <- true
		}(i)
	}

	// Concurrent reads.
	for i := 0; i < 10; i++ {
		go func() {
			var v any
			err := db.Get("concurrent/doc", &v)
			if err != nil && err.Error() != "jsondb: document not found: concurrent/doc" {
				t.Errorf("concurrent read: %v", err)
			}
			done <- true
		}()
	}

	// Wait for all goroutines.
	for i := 0; i < 20; i++ {
		<-done
	}
}
