package jsondb

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// List returns all document paths under the given prefix.
// The paths are relative to the db root, without the .json extension.
func (db *DB) List(prefix string) ([]string, error) {
	// Walk is not locked — filesystem provides best-effort consistency.
	searchDir := filepath.Join(db.root, filepath.FromSlash(prefix))

	var results []string
	err := filepath.Walk(searchDir, func(walkPath string, info os.FileInfo, err error) error {
		if err != nil {
			// If the directory doesn't exist yet, return empty list.
			if os.IsNotExist(err) {
				return nil
			}
			return err
		}

		if info.IsDir() {
			return nil
		}

		if !strings.HasSuffix(info.Name(), ".json") {
			return nil
		}

		// Compute path relative to db root.
		rel, err := filepath.Rel(db.root, walkPath)
		if err != nil {
			return err
		}

		// Strip .json extension.
		rel = strings.TrimSuffix(rel, ".json")
		rel = filepath.ToSlash(rel)

		results = append(results, rel)
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("jsondb: list: %w", err)
	}

	return results, nil
}
