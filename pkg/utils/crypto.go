package utils

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

// SHA256Hex calculates the SHA-256 hash of the provided data
// and returns it as a hexadecimal encoded string.
func SHA256Hex(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}

// SHA512Hex calculates the SHA-512 hash of the provided data
// and returns it as a hexadecimal encoded string.
func SHA512Hex(data []byte) string {
	hash := sha512.Sum512(data)
	return hex.EncodeToString(hash[:])
}
