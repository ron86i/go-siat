package utils

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGzip(t *testing.T) {
	data := []byte("hello world")
	compressed, err := Gzip(data)
	assert.NoError(t, err)
	assert.NotNil(t, compressed)

	// Decompress to verify
	r, _ := gzip.NewReader(bytes.NewReader(compressed))
	decompressed, _ := io.ReadAll(r)
	assert.Equal(t, data, decompressed)
}

func TestCompressAndHash(t *testing.T) {
	data := []byte("test data")
	hash, encoded, err := CompressAndHash(data)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)
	assert.NotEmpty(t, encoded)
}

func TestCreateTarGz(t *testing.T) {
	files := map[string][]byte{
		"file1.txt": []byte("content1"),
		"file2.txt": []byte("content2"),
	}
	compressed, err := CreateTarGz(files)
	assert.NoError(t, err)
	assert.NotNil(t, compressed)

	// Decompress and verify TAR contents
	gr, _ := gzip.NewReader(bytes.NewReader(compressed))
	tr := tar.NewReader(gr)

	count := 0
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		assert.NoError(t, err)
		content, _ := io.ReadAll(tr)
		assert.Equal(t, files[header.Name], content)
		count++
	}
	assert.Equal(t, 2, count)
}
