package utils

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/beevik/etree"
	dsig "github.com/russellhaering/goxmldsig"
)

// pemKeyStore implementa dsig.X509KeyStore para cargar llaves desde archivos PEM.
type pemKeyStore struct {
	PrivateKey *rsa.PrivateKey
	Cert       []byte
}

func (ks *pemKeyStore) GetKeyPair() (*rsa.PrivateKey, []byte, error) {
	return ks.PrivateKey, ks.Cert, nil
}

// SignXML firma un documento XML utilizando firma envuelta (enveloped) siguiendo el estándar XMLDSig.
// Utiliza las librerías goxmldsig y etree para garantizar una canonicalización (C14N) correcta.
func SignXML(xmlBytes []byte, keyPath, certPath string) ([]byte, error) {
	// 1. Cargar clave privada
	privKey, err := loadRSAPrivateKey(keyPath)
	if err != nil {
		return nil, err
	}

	// 2. Cargar certificado
	certData, err := os.ReadFile(certPath)
	if err != nil {
		return nil, fmt.Errorf("error leyendo certificado: %w", err)
	}
	blockCert, _ := pem.Decode(certData)
	if blockCert == nil {
		return nil, fmt.Errorf("error decodificando certificado PEM")
	}

	// 3. Configurar KeyStore
	ks := &pemKeyStore{
		PrivateKey: privKey,
		Cert:       blockCert.Bytes,
	}

	// 4. Configurar contexto de firma
	ctx := dsig.NewDefaultSigningContext(ks)
	ctx.Canonicalizer = dsig.MakeC14N10WithCommentsCanonicalizer()
	ctx.SetSignatureMethod(dsig.RSASHA256SignatureMethod)

	// 5. Parsear XML
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(xmlBytes); err != nil {
		return nil, fmt.Errorf("error leyendo XML: %w", err)
	}

	// 6. Firmar XML (Enveloped Signature)
	signedElement, err := ctx.SignEnveloped(doc.Root())
	if err != nil {
		return nil, fmt.Errorf("error firmando XML: %w", err)
	}

	signedDoc := etree.NewDocument()
	signedDoc.SetRoot(signedElement)

	// 7. Renderizar a bytes
	var buf bytes.Buffer
	if _, err := signedDoc.WriteTo(&buf); err != nil {
		return nil, fmt.Errorf("error convirtiendo XML firmado a bytes: %w", err)
	}

	return buf.Bytes(), nil
}

// loadRSAPrivateKey carga una clave privada RSA desde un archivo PEM, soportando PKCS#1 y PKCS#8.
func loadRSAPrivateKey(path string) (*rsa.PrivateKey, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error leyendo clave privada: %w", err)
	}
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, fmt.Errorf("formato PEM inválido en llave privada")
	}

	// Intentar parsear como PKCS#1 (RSA Private Key)
	if key, err := x509.ParsePKCS1PrivateKey(block.Bytes); err == nil {
		return key, nil
	}

	// Intentar parsear como PKCS#8 (un-encrypted)
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("error parseando clave privada (identificada como PKCS#8): %w", err)
	}

	rsaKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("la clave privada cargada no es de tipo RSA")
	}

	return rsaKey, nil
}
