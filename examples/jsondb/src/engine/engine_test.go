package engine_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/drajin/jsondb/engine"
	"github.com/drajin/jsondb/storage"
)

func TestEngineBasicOperations(t *testing.T) {
	// Create temp directory
	tmpDir, err := os.MkdirTemp("", "jsondb-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create test JSON file
	testFile := filepath.Join(tmpDir, "test.json")
	testData := `{
  "title": "Test Document",
  "content": "Hello World",
  "tags": ["important", "draft"],
  "metadata": {
    "author": "Test User",
    "version": 1
  }
}`
	if err := os.WriteFile(testFile, []byte(testData), 0644); err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}

	// Create engine
	backend := storage.NewDirectIOBackend()
	eng := engine.New(backend)

	// Open file
	if err := eng.Open(testFile); err != nil {
		t.Fatalf("failed to open file: %v", err)
	}
	defer eng.Close()

	// Test Get
	title, err := eng.Get("title")
	if err != nil {
		t.Fatalf("failed to get title: %v", err)
	}
	if title != "Test Document" {
		t.Errorf("expected 'Test Document', got '%v'", title)
	}

	// Test nested Get
	author, err := eng.Get("metadata.author")
	if err != nil {
		t.Fatalf("failed to get metadata.author: %v", err)
	}
	if author != "Test User" {
		t.Errorf("expected 'Test User', got '%v'", author)
	}

	// Test array Get
	tags, err := eng.Get("tags")
	if err != nil {
		t.Fatalf("failed to get tags: %v", err)
	}
	tagArr, ok := tags.([]interface{})
	if !ok {
		t.Fatalf("tags is not an array")
	}
	if len(tagArr) != 2 {
		t.Errorf("expected 2 tags, got %d", len(tagArr))
	}

	// Test Set
	if err := eng.Set("title", "Updated Title"); err != nil {
		t.Fatalf("failed to set title: %v", err)
	}

	title, err = eng.Get("title")
	if err != nil {
		t.Fatalf("failed to get updated title: %v", err)
	}
	if title != "Updated Title" {
		t.Errorf("expected 'Updated Title', got '%v'", title)
	}

	// Test Save
	if err := eng.Save(); err != nil {
		t.Fatalf("failed to save: %v", err)
	}

	// Verify file was updated
	data, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("failed to read test file: %v", err)
	}
	if !contains(string(data), "Updated Title") {
		t.Errorf("file was not updated")
	}

	// Test Delete
	if err := eng.Delete("content"); err != nil {
		t.Fatalf("failed to delete content: %v", err)
	}

	_, err = eng.Get("content")
	if err == nil {
		t.Errorf("expected error for deleted field")
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && (s[0:1] == substr[0:1] || contains(s[1:], substr)))
}
