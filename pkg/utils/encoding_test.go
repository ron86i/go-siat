package utils

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGzip(t *testing.T) {
	t.Run("Valid Data", func(t *testing.T) {
		data := []byte("hello world")
		compressed, err := Gzip(data)
		assert.NoError(t, err)
		assert.NotNil(t, compressed)

		// Decompress to verify round-trip
		r, err := gzip.NewReader(bytes.NewReader(compressed))
		require.NoError(t, err)
		decompressed, err := io.ReadAll(r)
		require.NoError(t, err)
		assert.Equal(t, data, decompressed)
	})

	t.Run("Empty Data", func(t *testing.T) {
		compressed, err := Gzip([]byte{})
		assert.NoError(t, err)
		assert.NotNil(t, compressed)

		r, err := gzip.NewReader(bytes.NewReader(compressed))
		require.NoError(t, err)
		decompressed, err := io.ReadAll(r)
		require.NoError(t, err)
		assert.Equal(t, []byte{}, decompressed)
	})

	t.Run("Large Data", func(t *testing.T) {
		data := bytes.Repeat([]byte("A"), 10000)
		compressed, err := Gzip(data)
		assert.NoError(t, err)
		assert.NotNil(t, compressed)
		// Compressed should be smaller than original for repetitive data
		assert.Less(t, len(compressed), len(data))
	})
}

func TestCompressAndHash(t *testing.T) {
	t.Run("Valid Data", func(t *testing.T) {
		data := []byte("test data")
		hash, encoded, err := CompressAndHash(data)
		assert.NoError(t, err)
		assert.NotEmpty(t, hash)
		assert.NotEmpty(t, encoded)
		// SHA256 hex hash should be 64 chars
		assert.Equal(t, 64, len(hash))
	})

	t.Run("Empty Data", func(t *testing.T) {
		hash, encoded, err := CompressAndHash([]byte{})
		assert.NoError(t, err)
		assert.NotEmpty(t, hash)
		assert.NotEmpty(t, encoded)
	})

	t.Run("Large Data", func(t *testing.T) {
		data := bytes.Repeat([]byte("SIAT"), 5000)
		hash, encoded, err := CompressAndHash(data)
		assert.NoError(t, err)
		assert.Equal(t, 64, len(hash))
		assert.NotEmpty(t, encoded)
	})
}

func TestCreateTarGz(t *testing.T) {
	t.Run("Multiple Files", func(t *testing.T) {
		files := map[string][]byte{
			"file1.txt": []byte("content1"),
			"file2.txt": []byte("content2"),
		}
		compressed, err := CreateTarGz(files)
		assert.NoError(t, err)
		assert.NotNil(t, compressed)

		// Decompress and verify TAR contents
		gr, err := gzip.NewReader(bytes.NewReader(compressed))
		require.NoError(t, err)
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
	})

	t.Run("Single File", func(t *testing.T) {
		files := map[string][]byte{"test.txt": []byte("test")}
		compressed, err := CreateTarGz(files)
		assert.NoError(t, err)
		assert.NotNil(t, compressed)
	})

	t.Run("Empty Map", func(t *testing.T) {
		compressed, err := CreateTarGz(map[string][]byte{})
		assert.NoError(t, err)
		assert.NotNil(t, compressed)
	})

	t.Run("Binary Content", func(t *testing.T) {
		files := map[string][]byte{
			"binary.bin": {0x00, 0x01, 0xFF, 0xFE, 0x80},
		}
		compressed, err := CreateTarGz(files)
		assert.NoError(t, err)
		assert.NotNil(t, compressed)

		// Verify binary round-trip
		gr, err := gzip.NewReader(bytes.NewReader(compressed))
		require.NoError(t, err)
		tr := tar.NewReader(gr)
		_, err = tr.Next()
		require.NoError(t, err)
		content, _ := io.ReadAll(tr)
		assert.Equal(t, files["binary.bin"], content)
	})
}
