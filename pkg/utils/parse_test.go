package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseIntSafe(t *testing.T) {
	t.Run("Valid Integer", func(t *testing.T) {
		val, err := ParseIntSafe(" 123 ")
		assert.NoError(t, err)
		assert.Equal(t, 123, val)
	})

	t.Run("Empty Value", func(t *testing.T) {
		_, err := ParseIntSafe("")
		assert.Error(t, err)
		assert.Equal(t, "value is empty, expected a number", err.Error())
	})

	t.Run("Invalid Integer", func(t *testing.T) {
		_, err := ParseIntSafe("abc")
		assert.Error(t, err)
	})
}

func TestParseInt64Safe(t *testing.T) {
	t.Run("Valid Int64", func(t *testing.T) {
		val, err := ParseInt64Safe(" 1234567890123 ")
		assert.NoError(t, err)
		assert.Equal(t, int64(1234567890123), val)
	})

	t.Run("Empty Value", func(t *testing.T) {
		_, err := ParseInt64Safe("")
		assert.Error(t, err)
		assert.Equal(t, "value is empty, expected a number", err.Error())
	})
}

func TestPointers(t *testing.T) {
	t.Run("Float64Ptr", func(t *testing.T) {
		v := 10.5
		ptr := Float64Ptr(v)
		assert.Equal(t, v, *ptr)
	})

	t.Run("Int64Ptr", func(t *testing.T) {
		v := int64(100)
		ptr := Int64Ptr(v)
		assert.Equal(t, v, *ptr)
	})

	t.Run("IntPtr", func(t *testing.T) {
		v := int(100)
		ptr := IntPtr(v)
		assert.Equal(t, v, *ptr)
	})
}
