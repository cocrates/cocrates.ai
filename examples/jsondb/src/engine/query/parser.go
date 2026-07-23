package query

import (
	"fmt"
	"strconv"
	"strings"
)

// QueryType represents the type of query.
type QueryType int

const (
	// QueryTypeBasic is a simple dot-separated path.
	QueryTypeBasic QueryType = iota
	// QueryTypeJSONPath is a JSONPath expression.
	QueryTypeJSONPath
)

// ParsedQuery represents a parsed query.
type ParsedQuery struct {
	Type     QueryType
	Path     string
	Segments []PathSegment
	JSONPath string // Raw JSONPath expression
}

// PathSegment represents a segment in a basic path.
type PathSegment struct {
	Field string
	Index *int
}

// Parse parses a query path and returns a ParsedQuery.
func Parse(path string) (*ParsedQuery, error) {
	if path == "" {
		return nil, fmt.Errorf("empty path")
	}

	// Detect JSONPath (starts with $)
	if strings.HasPrefix(path, "$") {
		return &ParsedQuery{
			Type:     QueryTypeJSONPath,
			Path:     path,
			JSONPath: path,
		}, nil
	}

	// Parse basic path (dot-separated)
	segments, err := parseBasicPath(path)
	if err != nil {
		return nil, err
	}

	return &ParsedQuery{
		Type:     QueryTypeBasic,
		Path:     path,
		Segments: segments,
	}, nil
}

// parseBasicPath parses a dot-separated path into segments.
func parseBasicPath(path string) ([]PathSegment, error) {
	if path == "" {
		return nil, fmt.Errorf("empty path")
	}

	var segments []PathSegment
	parts := strings.Split(path, ".")

	for _, part := range parts {
		segment, err := parseSegment(part)
		if err != nil {
			return nil, err
		}
		segments = append(segments, segment)
	}

	return segments, nil
}

// parseSegment parses a single path segment (e.g., "field[0]").
func parseSegment(part string) (PathSegment, error) {
	segment := PathSegment{}

	// Check for array index
	bracketStart := strings.Index(part, "[")
	if bracketStart != -1 {
		segment.Field = part[:bracketStart]
		bracketEnd := strings.Index(part, "]")
		if bracketEnd == -1 {
			return segment, fmt.Errorf("unclosed bracket in segment: %s", part)
		}

		indexStr := part[bracketStart+1 : bracketEnd]
		if indexStr == "" {
			return segment, fmt.Errorf("empty index in segment: %s", part)
		}

		index, err := strconv.Atoi(indexStr)
		if err != nil {
			return segment, fmt.Errorf("invalid index in segment: %s", part)
		}
		segment.Index = &index
	} else {
		segment.Field = part
	}

	if segment.Field == "" {
		return segment, fmt.Errorf("empty field name in segment: %s", part)
	}

	return segment, nil
}
