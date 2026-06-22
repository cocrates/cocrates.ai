package main

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/cocrates/jsondb/server"
)

func main() {
	portDefault := envInt("JSONDB_PORT", 8080)
	dbrootDefault := envString("JSONDB_ROOT", "./jsondb")

	port := flag.Int("port", portDefault, "port to listen on")
	dbroot := flag.String("dbroot", dbrootDefault, "database root directory")
	schemaroot := flag.String("schemaroot", "", "schema root directory (default: {dbroot}/schema)")
	flag.Parse()

	if err := server.Run(server.Config{
		Port:       *port,
		DBRoot:     *dbroot,
		SchemaRoot: *schemaroot,
	}); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func envString(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func envInt(key string, fallback int) int {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return fallback
	}
	return n
}
