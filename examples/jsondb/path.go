package jsondb

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// GetPath retrieves a specific value from a JSON document using a GJSON path.
func (db *DB) GetPath(path, jsonPath string) (*gjson.Result, error) {
	_, fl := db.rLock(path)
	defer db.rUnlock(path, fl)

	fp := db.filePath(path)
	data, err := os.ReadFile(fp)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("jsondb: document not found: %s", path)
		}
		return nil, fmt.Errorf("jsondb: read file: %w", err)
	}

	result := gjson.GetBytes(data, jsonPath)
	if !result.Exists() {
		return nil, fmt.Errorf("jsondb: path %q not found in document %s", jsonPath, path)
	}

	return &result, nil
}

// SetPath sets a specific value within a JSON document using an SJSON path.
func (db *DB) SetPath(path, jsonPath string, value any) error {
	_, fl := db.wLock(path)
	defer db.wUnlock(path, fl)

	fp := db.filePath(path)
	data, err := os.ReadFile(fp)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("jsondb: document not found: %s", path)
		}
		return fmt.Errorf("jsondb: read file: %w", err)
	}

	// Convert value to a JSON-safe form.
	var jsonVal any
	switch v := value.(type) {
	case string:
		// Try to parse as JSON first (for objects, arrays, numbers, booleans).
		if err := json.Unmarshal([]byte(v), &jsonVal); err != nil {
			// Not valid JSON — treat as raw string.
			jsonVal = v
		}
	default:
		jsonVal = v
	}

	updated, err := sjson.SetBytes(data, jsonPath, jsonVal)
	if err != nil {
		return fmt.Errorf("jsondb: set path: %w", err)
	}

	if err := os.WriteFile(fp, updated, 0644); err != nil {
		return fmt.Errorf("jsondb: write file: %w", err)
	}

	return nil
}

// DeletePath removes a specific value from a JSON document using an SJSON path.
func (db *DB) DeletePath(path, jsonPath string) error {
	_, fl := db.wLock(path)
	defer db.wUnlock(path, fl)

	fp := db.filePath(path)
	data, err := os.ReadFile(fp)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("jsondb: document not found: %s", path)
		}
		return fmt.Errorf("jsondb: read file: %w", err)
	}

	updated, err := sjson.DeleteBytes(data, jsonPath)
	if err != nil {
		return fmt.Errorf("jsondb: delete path: %w", err)
	}

	if err := os.WriteFile(fp, updated, 0644); err != nil {
		return fmt.Errorf("jsondb: write file: %w", err)
	}

	return nil
}
