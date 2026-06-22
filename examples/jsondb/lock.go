package jsondb

import (
	"path/filepath"
	"sync"
	"sync/atomic"
)

// fileLock wraps sync.RWMutex with a reference counter.
type fileLock struct {
	rwmu     sync.RWMutex
	refCount atomic.Int32
}

// lockMap manages per-file locks.
type lockMap struct {
	muMap sync.Map // map[string]*fileLock
}

// getOrCreate returns the lock for the given key, creating one if needed.
// The caller must call release when done with the lock.
func (lm *lockMap) getOrCreate(key string) *fileLock {
	// Check if lock already exists.
	if existing, ok := lm.muMap.Load(key); ok {
		fl := existing.(*fileLock)
		fl.refCount.Add(1)
		return fl
	}

	// Create new lock and try to store it.
	fl := &fileLock{}
	stored, loaded := lm.muMap.LoadOrStore(key, fl)
	if loaded {
		// Another goroutine beat us to it; use theirs.
		stored.(*fileLock).refCount.Add(1)
		return stored.(*fileLock)
	}

	// We stored our new lock.
	fl.refCount.Add(1)
	return fl
}

// release decrements the reference counter.
// In the future, if refCount reaches 0, the entry could be deleted from the map.
func (lm *lockMap) release(key string, fl *fileLock) {
	fl.refCount.Add(-1)
	// Cleanup is deferred: sync.RWMutex is small (~56 bytes),
	// and stale entries cause no correctness issues.
	// If memory becomes a concern, periodically delete entries with refCount == 0.
}

// lockKey normalizes a data path into a canonical map key.
// It joins with db.root, cleans the path, and converts to forward slashes.
func (db *DB) lockKey(dataPath string) string {
	clean := filepath.Clean(filepath.Join(db.root, dataPath))
	return filepath.ToSlash(clean)
}

// rLock acquires a read lock for the given data path.
// Returns the canonical key and the fileLock. Caller must call rUnlock.
func (db *DB) rLock(dataPath string) (string, *fileLock) {
	key := db.lockKey(dataPath)
	fl := db.locks.getOrCreate(key)
	fl.rwmu.RLock()
	return key, fl
}

// rUnlock releases a read lock.
func (db *DB) rUnlock(key string, fl *fileLock) {
	fl.rwmu.RUnlock()
	db.locks.release(key, fl)
}

// wLock acquires a write lock for the given data path.
// Returns the canonical key and the fileLock. Caller must call wUnlock.
func (db *DB) wLock(dataPath string) (string, *fileLock) {
	key := db.lockKey(dataPath)
	fl := db.locks.getOrCreate(key)
	fl.rwmu.Lock()
	return key, fl
}

// wUnlock releases a write lock.
func (db *DB) wUnlock(key string, fl *fileLock) {
	fl.rwmu.Unlock()
	db.locks.release(key, fl)
}
