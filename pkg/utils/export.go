package utils

import (
	"archive/tar"
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"os"
)

// XMLSigner defines the interface for signing XML documents
type XMLSigner interface {
	SignXML(xmlBytes []byte) ([]byte, error)
}

// ExportXML serializes a given struct to XML and writes it to the specified file path.
func ExportXML(factura any, path string) error {
	xmlData, err := xml.Marshal(factura)
	if err != nil {
		return fmt.Errorf("error marshaling XML: %w", err)
	}

	// Add XML header
	xmlWithHeader := []byte(xml.Header + string(xmlData))

	return os.WriteFile(path, xmlWithHeader, 0644)
}

// ExportSignedXML serializes, signs the XML using the provided signer, and writes it to the file path.
func ExportSignedXML(factura any, signer XMLSigner, path string) error {
	xmlData, err := xml.Marshal(factura)
	if err != nil {
		return fmt.Errorf("error marshaling XML: %w", err)
	}

	signedXML, err := signer.SignXML(xmlData)
	if err != nil {
		return fmt.Errorf("error signing XML: %w", err)
	}

	// Add XML header
	xmlWithHeader := []byte(xml.Header + string(signedXML))

	return os.WriteFile(path, xmlWithHeader, 0644)
}

// ExportTarGz serializes and packages multiple invoices into a single tar.gz file.
// If signer is not nil, it will also sign each XML before packaging.
func ExportTarGz(facturas []any, signer XMLSigner, path string) error {
	var tarBuf bytes.Buffer
	tw := tar.NewWriter(&tarBuf)

	for i, f := range facturas {
		xmlData, err := xml.Marshal(f)
		if err != nil {
			return fmt.Errorf("error marshaling invoice %d: %w", i+1, err)
		}

		var xmlToSend = xmlData
		if signer != nil {
			xmlToSend, err = signer.SignXML(xmlData)
			if err != nil {
				return fmt.Errorf("error signing invoice %d: %w", i+1, err)
			}
		}

		// Add XML header
		xmlWithHeader := []byte(xml.Header + string(xmlToSend))

		hdr := &tar.Header{
			Name: fmt.Sprintf("factura_%d.xml", i+1),
			Mode: 0600,
			Size: int64(len(xmlWithHeader)),
		}

		if err := tw.WriteHeader(hdr); err != nil {
			return fmt.Errorf("error writing tar header for invoice %d: %w", i+1, err)
		}

		if _, err := tw.Write(xmlWithHeader); err != nil {
			return fmt.Errorf("error writing tar data for invoice %d: %w", i+1, err)
		}
	}

	if err := tw.Close(); err != nil {
		return fmt.Errorf("error closing tar writer: %w", err)
	}

	// Compress using the existing compress utility which returns base64 string
	// But since we want to export to a raw .tar.gz file, we should write the raw gzip bytes.
	// Wait, CompressAndHash returns base64 string. 
	// To save to disk as a binary .tar.gz, we should decode the base64 back to binary,
	// or use a raw compression method. Let's decode the base64.
	_, b64EncodedArchivo, err := CompressAndHash(tarBuf.Bytes())
	if err != nil {
		return fmt.Errorf("error compressing to gzip: %w", err)
	}

	rawGzBytes, err := base64.StdEncoding.DecodeString(b64EncodedArchivo)
	if err != nil {
		return fmt.Errorf("error decoding base64 gzip: %w", err)
	}

	return os.WriteFile(path, rawGzBytes, 0644)
}
