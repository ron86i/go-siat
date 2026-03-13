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

// pemKeyStore implements dsig.X509KeyStore to load keys from PEM files.
type pemKeyStore struct {
	PrivateKey *rsa.PrivateKey
	Cert       []byte
}

func (ks *pemKeyStore) GetKeyPair() (*rsa.PrivateKey, []byte, error) {
	return ks.PrivateKey, ks.Cert, nil
}

// SignXMLBytes signs an XML document receiving certificates and key directly in bytes.
func SignXMLBytes(xmlBytes, keyBytes, certBytes []byte) ([]byte, error) {
	// 1. Parse private key from provided bytes
	privKey, err := parseRSAPrivateKey(keyBytes)
	if err != nil {
		return nil, err
	}

	// 2. Decode PEM certificate
	blockCert, _ := pem.Decode(certBytes)
	if blockCert == nil {
		return nil, fmt.Errorf("error decoding PEM certificate")
	}

	// 3. Configure KeyStore
	ks := &pemKeyStore{
		PrivateKey: privKey,
		Cert:       blockCert.Bytes,
	}

	// 4. Configure signing context
	ctx := dsig.NewDefaultSigningContext(ks)
	ctx.Canonicalizer = dsig.MakeC14N10WithCommentsCanonicalizer()
	ctx.SetSignatureMethod(dsig.RSASHA256SignatureMethod)

	// 5. Parse XML
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(xmlBytes); err != nil {
		return nil, err
	}

	// 6. Sign XML (Enveloped Signature)
	signedElement, err := ctx.SignEnveloped(doc.Root())
	if err != nil {
		return nil, err
	}

	signedDoc := etree.NewDocument()
	signedDoc.SetRoot(signedElement)

	// 7. Render to bytes
	var buf bytes.Buffer
	if _, err := signedDoc.WriteTo(&buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// SignXML signs an XML document receiving certificates and key from files.
func SignXML(xmlBytes []byte, keyPath, certPath string) ([]byte, error) {
	keyData, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}

	certData, err := os.ReadFile(certPath)
	if err != nil {
		return nil, err
	}

	return SignXMLBytes(xmlBytes, keyData, certData)
}

// parseRSAPrivateKey processes PEM key bytes (PKCS#1 or PKCS#8).
func parseRSAPrivateKey(keyData []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(keyData)
	if block == nil {
		return nil, fmt.Errorf("invalid PEM format in private key")
	}

	// Try parsing as PKCS#1
	if key, err := x509.ParsePKCS1PrivateKey(block.Bytes); err == nil {
		return key, nil
	}

	// Try parsing as PKCS#8
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("loaded private key is not of type RSA")
	}

	return rsaKey, nil
}
