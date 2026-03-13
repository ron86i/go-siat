package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseIntSafe converts a string to int, removing whitespace.
// Returns a descriptive error if the conversion fails.
func ParseIntSafe(valStr string) (int, error) {
	cleanVal := strings.TrimSpace(valStr)
	if cleanVal == "" {
		return 0, fmt.Errorf("value is empty, expected a number")
	}

	val, err := strconv.Atoi(cleanVal)
	if err != nil {
		return 0, err
	}
	return val, nil
}

// ParseInt64Safe converts a string to int64 (ideal for NIT).
// Returns a descriptive error if the conversion fails.
func ParseInt64Safe(valStr string) (int64, error) {
	cleanVal := strings.TrimSpace(valStr)
	if cleanVal == "" {
		return 0, fmt.Errorf("value is empty, expected a number")
	}

	val, err := strconv.ParseInt(cleanVal, 10, 64)
	if err != nil {
		return 0, err
	}
	return val, nil
}
