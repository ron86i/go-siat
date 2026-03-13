package utils

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
)

// Gzip compress the provided data using Gzip.
func Gzip(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	if _, err := zw.Write(data); err != nil {
		return nil, err
	}
	if err := zw.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// CompressAndHash compresses the data with Gzip, calculates the SHA256 hex hash
// of the compressed bytes, and then encodes the compressed bytes to Base64.
// This is a common requirement for SIAT file uploads.
func CompressAndHash(data []byte) (hash, encoded string, err error) {
	compressed, err := Gzip(data)
	if err != nil {
		return "", "", err
	}

	hash = SHA256Hex(compressed)
	encoded = base64.StdEncoding.EncodeToString(compressed)
	return hash, encoded, nil
}
