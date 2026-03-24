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

// Float64Ptr returns a pointer to the given float64 value.
func Float64Ptr(v float64) *float64 {
	return &v
}

// Int64Ptr returns a pointer to the given int64 value.
func Int64Ptr(v int64) *int64 {
	return &v
}

// IntPtr returns a pointer to the given int value.
func IntPtr(v int) *int {
	return &v
}
