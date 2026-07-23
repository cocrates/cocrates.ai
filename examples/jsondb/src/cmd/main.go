package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/drajin/jsondb/server"
)

const (
	defaultServerURL = "http://localhost:8080"
	version          = "0.1.0"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "get":
		handleGet()
	case "set":
		handleSet()
	case "delete":
		handleDelete()
	case "patch":
		handlePatch()
	case "ls":
		handleLs()
	case "create":
		handleCreate()
	case "rm":
		handleRm()
	case "serve":
		handleServe()
	case "status":
		handleStatus()
	case "stop":
		handleStop()
	case "--help", "-h":
		printUsage()
	case "--version", "-v":
		fmt.Printf("jsondb version %s\n", version)
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println(`jsondb - JSON Database for Agent Collaboration

Usage:
  jsondb <command> [options]

Commands:
  get <file> [field]           Get data from file
  set <file> <field> <value>   Set field value
  delete <file> <field>        Delete field
  patch <file> <patch>         Apply patch (JSON Merge Patch)
  ls [directory]               List files
  create <file>                Create new file
  rm <file>                    Delete file
  serve                        Start HTTP server
  status                       Check server status
  stop                         Stop server

Options:
  --server <url>               Server URL (default: http://localhost:8080)
  --pretty                     Pretty print output
  --verbose                    Verbose output

Examples:
  jsondb get overview.json title
  jsondb set overview.json title "New Title"
  jsondb delete overview.json description
  jsondb serve --port 8080`)
}

// getServerURL returns the server URL from args or environment.
func getServerURL() string {
	for i, arg := range os.Args {
		if arg == "--server" && i+1 < len(os.Args) {
			return os.Args[i+1]
		}
	}
	if url := os.Getenv("JSONDB_SERVER"); url != "" {
		return url
	}
	return defaultServerURL
}

// isPretty checks if --pretty flag is set.
func isPretty() bool {
	for _, arg := range os.Args {
		if arg == "--pretty" {
			return true
		}
	}
	return false
}

// isVerbose checks if --verbose flag is set.
func isVerbose() bool {
	for _, arg := range os.Args {
		if arg == "--verbose" {
			return true
		}
	}
	return false
}

// httpRequest makes an HTTP request to the server.
func httpRequest(method, path string, body interface{}) (map[string]interface{}, error) {
	serverURL := getServerURL()
	url := serverURL + path

	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewReader(jsonData)
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to server: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return result, nil
}

// checkError checks if response indicates an error.
func checkError(resp map[string]interface{}) {
	if success, ok := resp["success"].(bool); ok && !success {
		if errObj, ok := resp["error"].(map[string]interface{}); ok {
			code := errObj["code"].(string)
			message := errObj["message"].(string)
			fmt.Fprintf(os.Stderr, "Error [%s]: %s\n", code, message)
			os.Exit(1)
		}
	}
}

// formatOutput formats the output based on flags.
func formatOutput(data interface{}) string {
	if isPretty() {
		bytes, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			return fmt.Sprintf("%v", data)
		}
		return string(bytes)
	}

	if str, ok := data.(string); ok {
		return fmt.Sprintf("%q", str)
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		return fmt.Sprintf("%v", data)
	}
	return string(bytes)
}

func handleGet() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: jsondb get <file> [field]\n")
		os.Exit(1)
	}

	file := os.Args[2]
	field := ""
	if len(os.Args) > 3 {
		field = os.Args[3]
	}

	// Build query path
	path := fmt.Sprintf("/api/v1/data/%s", file)
	if field != "" {
		// Check if it's a JSONPath (starts with $)
		if strings.HasPrefix(field, "$") {
			path += fmt.Sprintf("?query=%s", field)
		} else {
			path += fmt.Sprintf("?field=%s", field)
		}
	}

	resp, err := httpRequest("GET", path, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	checkError(resp)

	if data, ok := resp["data"]; ok {
		fmt.Println(formatOutput(data))
	}
}

func handleSet() {
	if len(os.Args) < 5 {
		fmt.Fprintf(os.Stderr, "Usage: jsondb set <file> <field> <value>\n")
		os.Exit(1)
	}

	file := os.Args[2]
	field := os.Args[3]
	valueStr := os.Args[4]

	// Parse value
	var value interface{}
	if err := json.Unmarshal([]byte(valueStr), &value); err != nil {
		// Treat as string
		value = valueStr
	}

	path := fmt.Sprintf("/api/v1/data/%s?field=%s", file, field)
	body := map[string]interface{}{"value": value}

	resp, err := httpRequest("PUT", path, body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	checkError(resp)

	if isVerbose() {
		fmt.Printf("OK: %s.%s updated\n", file, field)
	} else {
		fmt.Println("OK")
	}
}

func handleDelete() {
	if len(os.Args) < 4 {
		fmt.Fprintf(os.Stderr, "Usage: jsondb delete <file> <field>\n")
		os.Exit(1)
	}

	file := os.Args[2]
	field := os.Args[3]

	path := fmt.Sprintf("/api/v1/data/%s?field=%s", file, field)

	resp, err := httpRequest("DELETE", path, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	checkError(resp)

	if isVerbose() {
		fmt.Printf("OK: %s.%s deleted\n", file, field)
	} else {
		fmt.Println("OK")
	}
}

func handlePatch() {
	if len(os.Args) < 4 {
		fmt.Fprintf(os.Stderr, "Usage: jsondb patch <file> <patch>\n")
		os.Exit(1)
	}

	file := os.Args[2]
	patchStr := os.Args[3]

	// Parse patch
	var patch interface{}
	if err := json.Unmarshal([]byte(patchStr), &patch); err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid JSON patch: %v\n", err)
		os.Exit(1)
	}

	path := fmt.Sprintf("/api/v1/data/%s", file)
	resp, err := httpRequest("PATCH", path, patch)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	checkError(resp)

	if isVerbose() {
		fmt.Printf("OK: %s patched\n", file)
	} else {
		fmt.Println("OK")
	}
}

func handleLs() {
	dir := "."
	if len(os.Args) > 2 {
		dir = os.Args[2]
	}

	path := fmt.Sprintf("/api/v1/files?dir=%s", dir)

	resp, err := httpRequest("GET", path, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	checkError(resp)

	if data, ok := resp["data"].(map[string]interface{}); ok {
		if files, ok := data["files"].([]interface{}); ok {
			for _, f := range files {
				if file, ok := f.(map[string]interface{}); ok {
					fmt.Println(file["name"])
				}
			}
		}
	}
}

func handleCreate() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: jsondb create <file>\n")
		os.Exit(1)
	}

	file := os.Args[2]

	path := "/api/v1/files"
	body := map[string]interface{}{
		"path": file,
		"data": map[string]interface{}{},
	}

	resp, err := httpRequest("POST", path, body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	checkError(resp)

	fmt.Println("OK")
}

func handleRm() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: jsondb rm <file>\n")
		os.Exit(1)
	}

	file := os.Args[2]

	path := fmt.Sprintf("/api/v1/files/%s", file)

	resp, err := httpRequest("DELETE", path, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	checkError(resp)

	fmt.Println("OK")
}

func handleServe() {
	port := "8080"
	root := "."

	// Parse flags
	for i := 2; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "--port":
			if i+1 < len(os.Args) {
				port = os.Args[i+1]
				i++
			}
		case "--root":
			if i+1 < len(os.Args) {
				root = os.Args[i+1]
				i++
			}
		}
	}

	// Create and start server (Long-lived Engine management)
	srv := server.New(":" + port)
	srv.SetRoot(root)
	if err := srv.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
		os.Exit(1)
	}
}

func handleStatus() {
	serverURL := getServerURL()
	resp, err := http.Get(serverURL + "/api/v1/status")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Server is not running\n")
		os.Exit(1)
	}
	defer resp.Body.Close()

	fmt.Printf("Server is running at %s\n", serverURL)
}

func handleStop() {
	// TODO: Implement server stop
	fmt.Fprintf(os.Stderr, "stop not implemented yet\n")
	os.Exit(1)
}
