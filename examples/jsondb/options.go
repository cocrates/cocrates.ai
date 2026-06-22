package jsondb

// Option configures a DB instance.
type Option func(*DB)

// WithSchemaRoot sets a custom schema root directory.
// If not set, defaults to {dbroot}/schema.
func WithSchemaRoot(root string) Option {
	return func(db *DB) {
		db.schemaRoot = root
	}
}

// WithPort sets the default server port used by jsondbd.
// This is stored on the DB for the server to use.
func WithPort(port int) Option {
	return func(db *DB) {
		db.port = port
	}
}
