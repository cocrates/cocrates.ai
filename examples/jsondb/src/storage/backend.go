package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"syscall"
)

// StorageBackend defines the interface for storage operations.
type StorageBackend interface {
	// Open opens a JSON file for reading/writing.
	Open(path string) error
	// Read reads and parses the JSON file.
	Read() (map[string]interface{}, error)
	// Write writes data to the JSON file.
	Write(data map[string]interface{}) error
	// Close closes the file and releases locks.
	Close() error
	// Lock acquires an exclusive lock on the file.
	Lock() error
	// Unlock releases the lock on the file.
	Unlock() error
	// RLock acquires a shared lock on the file.
	RLock() error
	// RUnlock releases the shared lock.
	RUnlock() error
}

// DirectIOBackend implements StorageBackend using direct file I/O.
type DirectIOBackend struct {
	file    *os.File
	path    string
	data    map[string]interface{}
	dirty   bool
	mu      sync.RWMutex
	locked  bool
	shared  bool
}

// NewDirectIOBackend creates a new DirectIOBackend.
func NewDirectIOBackend() *DirectIOBackend {
	return &DirectIOBackend{}
}

// Open opens a JSON file.
func (b *DirectIOBackend) Open(path string) error {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("failed to resolve path: %w", err)
	}

	// Check if file exists
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		// Create empty JSON file
		b.file, err = os.Create(absPath)
		if err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}
		if _, err := b.file.Write([]byte("{}")); err != nil {
			b.file.Close()
			return fmt.Errorf("failed to write empty JSON: %w", err)
		}
		if _, err := b.file.Seek(0, 0); err != nil {
			b.file.Close()
			return fmt.Errorf("failed to seek: %w", err)
		}
	} else if err != nil {
		return fmt.Errorf("failed to stat file: %w", err)
	} else {
		b.file, err = os.OpenFile(absPath, os.O_RDWR, 0644)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
	}

	b.path = absPath
	b.data = make(map[string]interface{})
	b.dirty = false
	return nil
}

// Read reads and parses the JSON file.
func (b *DirectIOBackend) Read() (map[string]interface{}, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if b.file == nil {
		return nil, ErrNotOpen
	}

	// Seek to beginning
	if _, err := b.file.Seek(0, 0); err != nil {
		return nil, fmt.Errorf("failed to seek: %w", err)
	}

	// Read file content
	info, err := b.file.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to stat file: %w", err)
	}

	if info.Size() == 0 {
		return make(map[string]interface{}), nil
	}

	data := make([]byte, info.Size())
	if _, err := b.file.Read(data); err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	// Parse JSON
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	b.data = result
	return result, nil
}

// Write writes data to the JSON file.
func (b *DirectIOBackend) Write(data map[string]interface{}) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.file == nil {
		return ErrNotOpen
	}

	// Marshal JSON with indentation
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	// Truncate and write
	if err := b.file.Truncate(0); err != nil {
		return fmt.Errorf("failed to truncate file: %w", err)
	}

	if _, err := b.file.Seek(0, 0); err != nil {
		return fmt.Errorf("failed to seek: %w", err)
	}

	if _, err := b.file.Write(jsonData); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	b.data = data
	b.dirty = false
	return nil
}

// Close closes the file.
func (b *DirectIOBackend) Close() error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.file != nil {
		if err := b.file.Close(); err != nil {
			return fmt.Errorf("failed to close file: %w", err)
		}
		b.file = nil
	}
	return nil
}

// Lock acquires an exclusive lock on the file.
func (b *DirectIOBackend) Lock() error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.file == nil {
		return ErrNotOpen
	}

	if err := syscall.Flock(int(b.file.Fd()), syscall.LOCK_EX); err != nil {
		return fmt.Errorf("failed to acquire lock: %w", err)
	}
	b.locked = true
	return nil
}

// Unlock releases the lock on the file.
func (b *DirectIOBackend) Unlock() error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.file == nil {
		return ErrNotOpen
	}

	if err := syscall.Flock(int(b.file.Fd()), syscall.LOCK_UN); err != nil {
		return fmt.Errorf("failed to release lock: %w", err)
	}
	b.locked = false
	return nil
}

// RLock acquires a shared lock on the file.
func (b *DirectIOBackend) RLock() error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.file == nil {
		return ErrNotOpen
	}

	if err := syscall.Flock(int(b.file.Fd()), syscall.LOCK_SH); err != nil {
		return fmt.Errorf("failed to acquire shared lock: %w", err)
	}
	b.shared = true
	return nil
}

// RUnlock releases the shared lock.
func (b *DirectIOBackend) RUnlock() error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.file == nil {
		return ErrNotOpen
	}

	if err := syscall.Flock(int(b.file.Fd()), syscall.LOCK_UN); err != nil {
		return fmt.Errorf("failed to release shared lock: %w", err)
	}
	b.shared = false
	return nil
}

// Errors
var ErrNotOpen = fmt.Errorf("storage backend not open")
