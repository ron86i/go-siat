package pipeline_test

import (
	"bytes"
	"compress/gzip"
	"context"
	"crypto/sha256"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/ron86i/go-siat/v3/internal/pipeline"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type dummyDoc struct {
	XMLName xml.Name `xml:"testDocument"`
	Valor   string   `xml:"valor"`
}

func (d dummyDoc) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type raw dummyDoc
	return e.EncodeElement(raw(d), start)
}

type dummySigner struct {
	signed bool
	err    error
	empty  bool
}

func (s *dummySigner) SignXML(in []byte) ([]byte, error) {
	if s.err != nil {
		return nil, s.err
	}
	if s.empty {
		return []byte{}, nil
	}
	s.signed = true
	return append(in, []byte("<!--signature-->")...), nil
}

func decompressGzip(t *testing.T, data []byte) []byte {
	t.Helper()
	r, err := gzip.NewReader(bytes.NewReader(data))
	require.NoError(t, err)
	defer r.Close()
	decompressed, err := io.ReadAll(r)
	require.NoError(t, err)
	return decompressed
}

func TestFactura_Computarizada_PipelineCompleto(t *testing.T) {
	doc := dummyDoc{Valor: "hola"}
	archivo, hash, err := pipeline.Factura(doc, pipeline.ModalidadComputarizada, nil)
	require.NoError(t, err)

	// 1. Verificación independiente de SHA-256
	expectedHash := fmt.Sprintf("%x", sha256.Sum256(archivo))
	assert.Equal(t, expectedHash, hash, "el hash debe ser sha256 en minúsculas del archivo comprimido")

	// 2. Verificación independiente de descompresión GZIP
	decompressed := decompressGzip(t, archivo)
	assert.Contains(t, string(decompressed), "<valor>hola</valor>")
	assert.NotContains(t, string(decompressed), "<!--signature-->", "computarizada no debe firmar")

	// 3. Si se intenta pasar un firmador a modalidad computarizada, debe fallar
	_, _, err = pipeline.Factura(doc, pipeline.ModalidadComputarizada, &dummySigner{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "no admite firma")
}

func TestFactura_Electronica_PipelineCompleto(t *testing.T) {
	doc := dummyDoc{Valor: "electronica"}

	// 1. Sin firmador debe fallar
	_, _, err := pipeline.Factura(doc, pipeline.ModalidadElectronica, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "requiere un firmador digital")

	// 2. Con error en el firmador debe propagar el error
	signerErr := &dummySigner{err: errors.New("falla criptográfica")}
	_, _, err = pipeline.Factura(doc, pipeline.ModalidadElectronica, signerErr)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "firmar factura")

	// 3. Con firmador que devuelve vacío debe fallar
	signerEmpty := &dummySigner{empty: true}
	_, _, err = pipeline.Factura(doc, pipeline.ModalidadElectronica, signerEmpty)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "firmador devolvió XML vacío")

	// 4. Con firmador exitoso: verifica firma, compresión y hash
	signer := &dummySigner{}
	archivo, hash, err := pipeline.Factura(doc, pipeline.ModalidadElectronica, signer)
	require.NoError(t, err)
	assert.True(t, signer.signed)

	expectedHash := fmt.Sprintf("%x", sha256.Sum256(archivo))
	assert.Equal(t, expectedHash, hash)

	decompressed := decompressGzip(t, archivo)
	assert.Contains(t, string(decompressed), "<valor>electronica</valor>")
	assert.Contains(t, string(decompressed), "<!--signature-->", "debe contener la firma digital")
}

func TestFactura_ModalidadInvalida(t *testing.T) {
	doc := dummyDoc{Valor: "test"}
	_, _, err := pipeline.Factura(doc, 99, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "modalidad inválida")
}

func TestPaquete(t *testing.T) {
	docs := []xml.Marshaler{
		dummyDoc{Valor: "doc1"},
		dummyDoc{Valor: "doc2"},
	}

	archivo, hash, err := pipeline.Paquete(context.Background(), docs, pipeline.ModalidadComputarizada, nil)
	require.NoError(t, err)
	assert.NotEmpty(t, archivo)

	expectedHash := fmt.Sprintf("%x", sha256.Sum256(archivo))
	assert.Equal(t, expectedHash, hash)
}
