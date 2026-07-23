package query

import (
	"fmt"

	"github.com/ohler55/ojg/jp"
)

// Executor executes parsed queries against JSON data.
type Executor struct{}

// NewExecutor creates a new Executor.
func NewExecutor() *Executor {
	return &Executor{}
}

// Get retrieves a value from data using the parsed query.
func (e *Executor) Get(data interface{}, query *ParsedQuery) (interface{}, error) {
	switch query.Type {
	case QueryTypeBasic:
		return e.getBasic(data, query)
	case QueryTypeJSONPath:
		return e.getJSONPath(data, query)
	default:
		return nil, fmt.Errorf("unknown query type: %d", query.Type)
	}
}

// Set sets a value in data using the parsed query.
func (e *Executor) Set(data interface{}, query *ParsedQuery, value interface{}) error {
	switch query.Type {
	case QueryTypeBasic:
		return e.setBasic(data, query, value)
	case QueryTypeJSONPath:
		return e.setJSONPath(data, query, value)
	default:
		return fmt.Errorf("unknown query type: %d", query.Type)
	}
}

// Delete removes a value from data using the parsed query.
func (e *Executor) Delete(data interface{}, query *ParsedQuery) error {
	switch query.Type {
	case QueryTypeBasic:
		return e.deleteBasic(data, query)
	case QueryTypeJSONPath:
		return e.deleteJSONPath(data, query)
	default:
		return fmt.Errorf("unknown query type: %d", query.Type)
	}
}

// getBasic retrieves a value using a basic path.
func (e *Executor) getBasic(data interface{}, query *ParsedQuery) (interface{}, error) {
	current := data

	for i, segment := range query.Segments {
		obj, ok := current.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("cannot access field '%s' on non-object at segment %d", segment.Field, i)
		}

		val, exists := obj[segment.Field]
		if !exists {
			return nil, fmt.Errorf("field '%s' not found", segment.Field)
		}

		if segment.Index != nil {
			arr, ok := val.([]interface{})
			if !ok {
				return nil, fmt.Errorf("field '%s' is not an array", segment.Field)
			}
			if *segment.Index < 0 || *segment.Index >= len(arr) {
				return nil, fmt.Errorf("index %d out of bounds for field '%s'", *segment.Index, segment.Field)
			}
			val = arr[*segment.Index]
		}

		current = val
	}

	return current, nil
}

// setBasic sets a value using a basic path.
func (e *Executor) setBasic(data interface{}, query *ParsedQuery, value interface{}) error {
	if len(query.Segments) == 0 {
		return fmt.Errorf("empty path")
	}

	obj, ok := data.(map[string]interface{})
	if !ok {
		return fmt.Errorf("data is not an object")
	}

	current := obj

	for i := 0; i < len(query.Segments)-1; i++ {
		segment := query.Segments[i]
		val, exists := current[segment.Field]
		if !exists {
			// Create intermediate object
			newObj := make(map[string]interface{})
			current[segment.Field] = newObj
			current = newObj
			continue
		}

		if segment.Index != nil {
			arr, ok := val.([]interface{})
			if !ok {
				return fmt.Errorf("field '%s' is not an array", segment.Field)
			}
			if *segment.Index < 0 || *segment.Index >= len(arr) {
				return fmt.Errorf("index %d out of bounds for field '%s'", *segment.Index, segment.Field)
			}
			innerObj, ok := arr[*segment.Index].(map[string]interface{})
			if !ok {
				return fmt.Errorf("element at index %d is not an object", *segment.Index)
			}
			current = innerObj
		} else {
			innerObj, ok := val.(map[string]interface{})
			if !ok {
				return fmt.Errorf("field '%s' is not an object", segment.Field)
			}
			current = innerObj
		}
	}

	// Set the final value
	lastSegment := query.Segments[len(query.Segments)-1]
	if lastSegment.Index != nil {
		val, exists := current[lastSegment.Field]
		if !exists {
			return fmt.Errorf("field '%s' not found", lastSegment.Field)
		}
		arr, ok := val.([]interface{})
		if !ok {
			return fmt.Errorf("field '%s' is not an array", lastSegment.Field)
		}
		if *lastSegment.Index < 0 || *lastSegment.Index >= len(arr) {
			return fmt.Errorf("index %d out of bounds for field '%s'", *lastSegment.Index, lastSegment.Field)
		}
		arr[*lastSegment.Index] = value
	} else {
		current[lastSegment.Field] = value
	}

	return nil
}

// deleteBasic deletes a value using a basic path.
func (e *Executor) deleteBasic(data interface{}, query *ParsedQuery) error {
	if len(query.Segments) == 0 {
		return fmt.Errorf("empty path")
	}

	obj, ok := data.(map[string]interface{})
	if !ok {
		return fmt.Errorf("data is not an object")
	}

	current := obj

	for i := 0; i < len(query.Segments)-1; i++ {
		segment := query.Segments[i]
		val, exists := current[segment.Field]
		if !exists {
			return fmt.Errorf("field '%s' not found", segment.Field)
		}

		if segment.Index != nil {
			arr, ok := val.([]interface{})
			if !ok {
				return fmt.Errorf("field '%s' is not an array", segment.Field)
			}
			if *segment.Index < 0 || *segment.Index >= len(arr) {
				return fmt.Errorf("index %d out of bounds for field '%s'", *segment.Index, segment.Field)
			}
			innerObj, ok := arr[*segment.Index].(map[string]interface{})
			if !ok {
				return fmt.Errorf("element at index %d is not an object", *segment.Index)
			}
			current = innerObj
		} else {
			innerObj, ok := val.(map[string]interface{})
			if !ok {
				return fmt.Errorf("field '%s' is not an object", segment.Field)
			}
			current = innerObj
		}
	}

	// Delete the final value
	lastSegment := query.Segments[len(query.Segments)-1]
	if lastSegment.Index != nil {
		val, exists := current[lastSegment.Field]
		if !exists {
			return fmt.Errorf("field '%s' not found", lastSegment.Field)
		}
		arr, ok := val.([]interface{})
		if !ok {
			return fmt.Errorf("field '%s' is not an array", lastSegment.Field)
		}
		if *lastSegment.Index < 0 || *lastSegment.Index >= len(arr) {
			return fmt.Errorf("index %d out of bounds for field '%s'", *lastSegment.Index, lastSegment.Field)
		}
		// Remove element from array
		newArr := append(arr[:*lastSegment.Index], arr[*lastSegment.Index+1:]...)
		current[lastSegment.Field] = newArr
	} else {
		delete(current, lastSegment.Field)
	}

	return nil
}

// getJSONPath retrieves a value using JSONPath.
func (e *Executor) getJSONPath(data interface{}, query *ParsedQuery) (interface{}, error) {
	expr, err := jp.ParseString(query.JSONPath)
	if err != nil {
		return nil, fmt.Errorf("invalid JSONPath expression: %w", err)
	}

	results := expr.Get(data)
	if len(results) == 0 {
		return nil, fmt.Errorf("no results for JSONPath: %s", query.JSONPath)
	}

	if len(results) == 1 {
		return results[0], nil
	}

	return results, nil
}

// setJSONPath sets a value using JSONPath.
func (e *Executor) setJSONPath(data interface{}, query *ParsedQuery, value interface{}) error {
	expr, err := jp.ParseString(query.JSONPath)
	if err != nil {
		return fmt.Errorf("invalid JSONPath expression: %w", err)
	}

	// For JSONPath set, we need to handle it differently
	// This is a simplified implementation
	_ = expr
	_ = value

	return fmt.Errorf("JSONPath set not fully implemented yet")
}

// deleteJSONPath deletes a value using JSONPath.
func (e *Executor) deleteJSONPath(data interface{}, query *ParsedQuery) error {
	expr, err := jp.ParseString(query.JSONPath)
	if err != nil {
		return fmt.Errorf("invalid JSONPath expression: %w", err)
	}

	// For JSONPath delete, we need to handle it differently
	// This is a simplified implementation
	_ = expr

	return fmt.Errorf("JSONPath delete not fully implemented yet")
}
