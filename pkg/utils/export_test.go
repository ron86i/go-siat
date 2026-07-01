package utils

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"
)

// dummyFactura is a simple struct for testing XML marshaling
type dummyFactura struct {
	XMLName xml.Name `xml:"Factura"`
	ID      int      `xml:"id"`
	Total   float64  `xml:"total"`
}

// dummySigner is a mock implementation of XMLSigner
type dummySigner struct{}

func (d dummySigner) SignXML(xmlBytes []byte) ([]byte, error) {
	// Just append a dummy signature for testing purposes
	return append(xmlBytes, []byte("<!-- signed -->")...), nil
}

func TestExportXML(t *testing.T) {
	factura := dummyFactura{ID: 1, Total: 100.50}

	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "factura.xml")

	err := ExportXML(factura, path)
	if err != nil {
		t.Fatalf("ExportXML failed: %v", err)
	}

	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Failed to read exported file: %v", err)
	}

	expectedXML := xml.Header + `<Factura><id>1</id><total>100.5</total></Factura>`
	if string(content) != expectedXML {
		t.Errorf("Unexpected XML content.\nGot: %s\nWant: %s", string(content), expectedXML)
	}
}

func TestExportSignedXML(t *testing.T) {
	factura := dummyFactura{ID: 2, Total: 250.75}
	signer := dummySigner{}

	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "factura_signed.xml")

	err := ExportSignedXML(factura, signer, path)
	if err != nil {
		t.Fatalf("ExportSignedXML failed: %v", err)
	}

	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Failed to read exported signed file: %v", err)
	}

	expectedXML := xml.Header + `<Factura><id>2</id><total>250.75</total></Factura><!-- signed -->`
	if string(content) != expectedXML {
		t.Errorf("Unexpected signed XML content.\nGot: %s\nWant: %s", string(content), expectedXML)
	}
}

func TestExportTarGz(t *testing.T) {
	facturas := []any{
		dummyFactura{ID: 3, Total: 300.00},
		dummyFactura{ID: 4, Total: 400.00},
	}
	signer := dummySigner{}

	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "facturas.tar.gz")

	err := ExportTarGz(facturas, signer, path)
	if err != nil {
		t.Fatalf("ExportTarGz failed: %v", err)
	}

	// Verify the file was created and is a valid tar.gz
	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("Failed to open exported tar.gz: %v", err)
	}
	defer file.Close()

	gzr, err := gzip.NewReader(file)
	if err != nil {
		t.Fatalf("Failed to create gzip reader: %v", err)
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	fileCount := 0
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("Failed to read tar header: %v", err)
		}

		fileCount++

		// Expected file name: factura_1.xml, factura_2.xml
		expectedName := fmt.Sprintf("factura_%d.xml", fileCount)
		if header.Name != expectedName {
			t.Errorf("Expected file name %s, got %s", expectedName, header.Name)
		}

		var buf bytes.Buffer
		_, err = io.Copy(&buf, tr)
		if err != nil {
			t.Fatalf("Failed to read file content from tar: %v", err)
		}

		contentStr := buf.String()
		if !bytes.Contains([]byte(contentStr), []byte("<!-- signed -->")) {
			t.Errorf("Expected file %s to be signed", header.Name)
		}
	}

	if fileCount != 2 {
		t.Errorf("Expected 2 files in tar.gz, got %d", fileCount)
	}
}
