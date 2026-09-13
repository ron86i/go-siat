package utils

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
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

// TarGzBase64Writer writes a TAR archive directly into GZIP, while calculating
// the SHA-256 hash and producing its Base64 representation. It avoids retaining
// both the uncompressed TAR and the compressed archive in memory.
//
// Call Close after writing all TAR entries and before reading Hash or Encoded.
type TarGzBase64Writer struct {
	encoded bytes.Buffer
	hash    hash.Hash
	size    countingWriter
	base64  io.WriteCloser
	gzip    *gzip.Writer
	tar     *tar.Writer
	closed  bool
}

// NewTarGzBase64Writer creates a streaming TAR.GZ encoder for SIAT archives.
func NewTarGzBase64Writer() *TarGzBase64Writer {
	return NewTarGzBase64WriterWithOutput(nil)
}

// NewTarGzBase64WriterWithOutput creates a streaming TAR.GZ encoder and, when
// output is not nil, copies the compressed archive to it while it is built.
// The caller owns output and is responsible for closing it when applicable.
func NewTarGzBase64WriterWithOutput(output io.Writer) *TarGzBase64Writer {
	w := &TarGzBase64Writer{hash: sha256.New()}
	w.base64 = base64.NewEncoder(base64.StdEncoding, &w.encoded)
	writers := []io.Writer{w.hash, &w.size, w.base64}
	if output != nil {
		writers = append(writers, output)
	}
	w.gzip = gzip.NewWriter(io.MultiWriter(writers...))
	w.tar = tar.NewWriter(w.gzip)
	return w
}

// Add writes one XML document to the TAR archive.
func (w *TarGzBase64Writer) Add(name string, content []byte) error {
	if w.closed {
		return fmt.Errorf("no se puede escribir un archivo TAR.GZ cerrado")
	}
	header := &tar.Header{Name: name, Mode: 0600, Size: int64(len(content))}
	if err := w.tar.WriteHeader(header); err != nil {
		return err
	}
	_, err := w.tar.Write(content)
	return err
}

// Close finalizes the TAR, GZIP and Base64 streams. It is safe to call once.
func (w *TarGzBase64Writer) Close() error {
	if w.closed {
		return nil
	}
	if err := w.tar.Close(); err != nil {
		return err
	}
	if err := w.gzip.Close(); err != nil {
		return err
	}
	if err := w.base64.Close(); err != nil {
		return err
	}
	w.closed = true
	return nil
}

// Hash returns the lowercase hexadecimal SHA-256 of the compressed archive.
// It must be called after Close.
func (w *TarGzBase64Writer) Hash() string {
	return hex.EncodeToString(w.hash.Sum(nil))
}

// Encoded returns the Base64-encoded compressed archive. It must be called
// after Close.
func (w *TarGzBase64Writer) Encoded() string {
	return w.encoded.String()
}

// Size returns the compressed TAR.GZ size in bytes. It is final after Close.
func (w *TarGzBase64Writer) Size() int64 {
	return w.size.n
}

type countingWriter struct {
	n int64
}

func (w *countingWriter) Write(data []byte) (int, error) {
	w.n += int64(len(data))
	return len(data), nil
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
