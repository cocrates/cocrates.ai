package jsondb

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

// SetSchema stores a JSON Schema at the given schema path under schemaroot.
func (db *DB) SetSchema(schemaPath string, schema any) error {
	path := db.schemaFilePath(schemaPath)

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("jsondb: create schema directory: %w", err)
	}

	data, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		return fmt.Errorf("jsondb: marshal schema: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("jsondb: write schema file: %w", err)
	}

	return nil
}

// GetSchema retrieves a JSON Schema and unmarshals it into v.
func (db *DB) GetSchema(schemaPath string, v any) error {
	path := db.schemaFilePath(schemaPath)

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("jsondb: schema not found: %s", schemaPath)
		}
		return fmt.Errorf("jsondb: read schema file: %w", err)
	}

	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("jsondb: unmarshal schema: %w", err)
	}

	return nil
}

// DeleteSchema removes a schema file.
func (db *DB) DeleteSchema(schemaPath string) error {
	path := db.schemaFilePath(schemaPath)

	if err := os.Remove(path); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("jsondb: schema not found: %s", schemaPath)
		}
		return fmt.Errorf("jsondb: delete schema file: %w", err)
	}

	return nil
}

// validate validates data against the JSON Schema at schemaRef
// (relative to schemaroot, without extension).
func (db *DB) validate(schemaRef string, data []byte) error {
	if schemaRef == "" {
		return nil // schemaless — skip validation
	}

	schPath := db.schemaFilePath(schemaRef)

	schBytes, err := os.ReadFile(schPath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("jsondb: schema not found: %s (referenced by _schema)", schemaRef)
		}
		return fmt.Errorf("jsondb: read schema: %w", err)
	}

	var schemaDoc any
	if err := json.Unmarshal(schBytes, &schemaDoc); err != nil {
		return fmt.Errorf("jsondb: unmarshal schema: %w", err)
	}

	// Use a unique identifier for the schema resource.
	schemaURL := "jsondb://schema/" + schemaRef

	compiler := jsonschema.NewCompiler()
	if err := compiler.AddResource(schemaURL, schemaDoc); err != nil {
		return fmt.Errorf("jsondb: add schema resource: %w", err)
	}

	sch, err := compiler.Compile(schemaURL)
	if err != nil {
		return fmt.Errorf("jsondb: compile schema: %w", err)
	}

	var instance any
	if err := json.Unmarshal(data, &instance); err != nil {
		return fmt.Errorf("jsondb: unmarshal data for validation: %w", err)
	}

	// Strip internal _schema field before validation.
	if m, ok := instance.(map[string]any); ok {
		delete(m, "_schema")
		instance = m
	}

	if err := sch.Validate(instance); err != nil {
		return fmt.Errorf("jsondb: schema validation failed: %w", err)
	}

	return nil
}
