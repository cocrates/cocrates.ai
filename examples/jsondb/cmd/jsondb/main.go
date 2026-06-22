package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/cocrates/jsondb/server"
)

func main() {
	// Global flags.
	serverAddr := flag.String("server", "http://localhost:8080", "jsondb server address")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		printUsage()
		os.Exit(1)
	}

	cmd := args[0]

	switch cmd {
	case "serve":
		runServer(args[1:])
	case "get":
		runGet(*serverAddr, args[1:])
	case "set":
		runSet(*serverAddr, args[1:])
	case "delete":
		runDelete(*serverAddr, args[1:])
	case "list":
		runList(*serverAddr, args[1:])
	case "exists":
		runExists(*serverAddr, args[1:])
	case "schema":
		runSchema(*serverAddr, args[1:])
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", cmd)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Fprintf(os.Stderr, `jsondb — Local JSON Document Store

Usage:
  jsondb [--server <addr>] <command> [arguments]

Commands:
  serve          Start the REST API server
  get <path> [<jsonpath>]     Get document or property
  set <path> [<jsonpath>] <value>  Set document or property
  delete <path> [<jsonpath>]  Delete document or property
  list [<prefix>]             List documents
  exists <path>               Check if document exists
  schema get <schemapath>     Get schema
  schema set <schemapath> <file>  Set schema from file
  schema delete <schemapath>  Delete schema

Examples:
  jsondb get overview
  jsondb get overview title
  jsondb set overview '{"title":"Hello"}'
  jsondb set overview title '"Hello"'
  jsondb list episode/
`)
}

// --- Client helpers ---

func doRequest(client *http.Client, method, url string, body []byte) ([]byte, error) {
	var reqBody io.Reader
	if body != nil {
		reqBody = bytes.NewReader(body)
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	// Parse response envelope.
	var env struct {
		OK    bool   `json:"ok"`
		Data  any    `json:"data,omitempty"`
		Error string `json:"error,omitempty"`
	}
	if err := json.Unmarshal(respBody, &env); err != nil {
		return nil, fmt.Errorf("parse response: %s", string(respBody))
	}

	if !env.OK {
		return nil, fmt.Errorf("%s", env.Error)
	}

	return respBody, nil
}

func extractData(respBody []byte) (any, error) {
	var env struct {
		Data any `json:"data"`
	}
	if err := json.Unmarshal(respBody, &env); err != nil {
		return nil, err
	}
	return env.Data, nil
}

// --- Serve command ---

func runServer(args []string) {
	fs := flag.NewFlagSet("serve", flag.ExitOnError)
	portDefault := envInt("JSONDB_PORT", 8080)
	dbrootDefault := envString("JSONDB_ROOT", "./jsondb")
	port := fs.Int("port", portDefault, "port to listen on")
	dbroot := fs.String("dbroot", dbrootDefault, "database root directory")
	schemaroot := fs.String("schemaroot", "", "schema root directory")
	fs.Parse(args)

	if err := server.Run(server.Config{
		Port:       *port,
		DBRoot:     *dbroot,
		SchemaRoot: *schemaroot,
	}); err != nil {
		log.Fatalf("server error: %v", err)
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

// --- Client commands ---

func runGet(serverAddr string, args []string) {
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "usage: jsondb get <path> [<jsonpath>]")
		os.Exit(1)
	}

	docPath := args[0]
	client := &http.Client{}

	if len(args) >= 2 {
		// Property-level get.
		url := fmt.Sprintf("%s/data/%s?path=%s", serverAddr, docPath, args[1])
		respBody, err := doRequest(client, "GET", url, nil)
		if err != nil {
			log.Fatal(err)
		}
		data, _ := extractData(respBody)
		printValue(data)
		return
	}

	// Document-level get.
	url := fmt.Sprintf("%s/data/%s", serverAddr, docPath)
	respBody, err := doRequest(client, "GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	data, _ := extractData(respBody)
	printJSON(data)
}

func runSet(serverAddr string, args []string) {
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: jsondb set <path> [<jsonpath>] <value>")
		os.Exit(1)
	}

	docPath := args[0]
	client := &http.Client{}

	if len(args) >= 3 {
		// Property-level set.
		url := fmt.Sprintf("%s/data/%s?path=%s", serverAddr, docPath, args[1])
		value := args[2]
		respBody, err := doRequest(client, "PUT", url, []byte(value))
		if err != nil {
			log.Fatal(err)
		}
		printResponse(respBody)
		return
	}

	// Document-level set.
	url := fmt.Sprintf("%s/data/%s", serverAddr, docPath)
	value := args[1]
	respBody, err := doRequest(client, "PUT", url, []byte(value))
	if err != nil {
		log.Fatal(err)
	}
	printResponse(respBody)
}

func runDelete(serverAddr string, args []string) {
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "usage: jsondb delete <path> [<jsonpath>]")
		os.Exit(1)
	}

	docPath := args[0]
	client := &http.Client{}

	if len(args) >= 2 {
		url := fmt.Sprintf("%s/data/%s?path=%s", serverAddr, docPath, args[1])
		respBody, err := doRequest(client, "DELETE", url, nil)
		if err != nil {
			log.Fatal(err)
		}
		printResponse(respBody)
		return
	}

	url := fmt.Sprintf("%s/data/%s", serverAddr, docPath)
	respBody, err := doRequest(client, "DELETE", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	printResponse(respBody)
}

func runList(serverAddr string, args []string) {
	prefix := ""
	if len(args) >= 1 {
		prefix = args[0]
	}

	url := fmt.Sprintf("%s/data?prefix=%s", serverAddr, prefix)
	client := &http.Client{}
	respBody, err := doRequest(client, "GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	data, _ := extractData(respBody)
	printJSON(data)
}

func runExists(serverAddr string, args []string) {
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "usage: jsondb exists <path>")
		os.Exit(1)
	}

	url := fmt.Sprintf("%s/data/%s", serverAddr, args[0])
	client := &http.Client{}
	_, err := doRequest(client, "GET", url, nil)
	if err != nil {
		// Check if it's a "not found" error.
		if strings.Contains(err.Error(), "not found") {
			printJSON(false)
			return
		}
		log.Fatal(err)
	}
	printJSON(true)
}

func runSchema(serverAddr string, args []string) {
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: jsondb schema get|set|delete <schemapath> [<file>]")
		os.Exit(1)
	}

	subcmd := args[0]
	schemaPath := args[1]

	baseURL := fmt.Sprintf("%s/schema/%s", serverAddr, schemaPath)
	client := &http.Client{}

	switch subcmd {
	case "get":
		respBody, err := doRequest(client, "GET", baseURL, nil)
		if err != nil {
			log.Fatal(err)
		}
		data, _ := extractData(respBody)
		printJSON(data)

	case "set":
		if len(args) < 3 {
			fmt.Fprintln(os.Stderr, "usage: jsondb schema set <schemapath> <file>")
			os.Exit(1)
		}
		schemaBytes, err := os.ReadFile(args[2])
		if err != nil {
			log.Fatalf("read schema file: %v", err)
		}
		respBody, err := doRequest(client, "PUT", baseURL, schemaBytes)
		if err != nil {
			log.Fatal(err)
		}
		printResponse(respBody)

	case "delete":
		respBody, err := doRequest(client, "DELETE", baseURL, nil)
		if err != nil {
			log.Fatal(err)
		}
		printResponse(respBody)

	default:
		fmt.Fprintf(os.Stderr, "unknown schema subcommand: %s\n", subcmd)
		os.Exit(1)
	}
}

// --- Output helpers ---

func printJSON(v any) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(v)
}

func printValue(v any) {
	// Print scalar values without JSON encoding.
	switch val := v.(type) {
	case string:
		fmt.Println(val)
	case float64:
		fmt.Println(val)
	case bool:
		fmt.Println(val)
	case nil:
		fmt.Println("null")
	default:
		printJSON(val)
	}
}

func printResponse(respBody []byte) {
	var env struct {
		OK    bool   `json:"ok"`
		Data  any    `json:"data,omitempty"`
		Error string `json:"error,omitempty"`
	}
	if err := json.Unmarshal(respBody, &env); err != nil {
		fmt.Println(string(respBody))
		return
	}
	if env.OK {
		if env.Data != nil {
			printJSON(env.Data)
		}
	} else {
		fmt.Fprintf(os.Stderr, "error: %s\n", env.Error)
	}
}
