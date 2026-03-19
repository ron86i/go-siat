package utils

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
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

	hash = fmt.Sprintf("%x", sha256.Sum256(compressed))
	encoded = base64.StdEncoding.EncodeToString(compressed)
	return hash, encoded, nil
}

// CreateTarGz creates a TAR.GZ archive from a map of filenames and contents.
func CreateTarGz(files map[string][]byte) ([]byte, error) {
	var buf bytes.Buffer
	gw := gzip.NewWriter(&buf)
	tw := tar.NewWriter(gw)

	for name, content := range files {
		header := &tar.Header{
			Name: name,
			Mode: 0600,
			Size: int64(len(content)),
		}
		if err := tw.WriteHeader(header); err != nil {
			return nil, err
		}
		if _, err := tw.Write(content); err != nil {
			return nil, err
		}
	}

	if err := tw.Close(); err != nil {
		return nil, err
	}
	if err := gw.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
