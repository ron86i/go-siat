package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSHA256Hex(t *testing.T) {
	data := []byte("hello")
	result := SHA256Hex(data)
	assert.NotEmpty(t, result)
	assert.Equal(t, "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824", result)
}

func TestSHA512Hex(t *testing.T) {
	data := []byte("hello")
	result := SHA512Hex(data)
	assert.NotEmpty(t, result)
}
