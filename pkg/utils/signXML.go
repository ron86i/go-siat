package utils

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"time"

	"github.com/beevik/etree"
	dsig "github.com/russellhaering/goxmldsig"
	"golang.org/x/crypto/pkcs12"
)

// pemKeyStore implements dsig.X509KeyStore to load keys from PEM files
type pemKeyStore struct {
	PrivateKey *rsa.PrivateKey
	Cert       []byte
}

// XMLDocumentSigner keeps already parsed signing material in memory. It is safe
// to use concurrently: each signature creates its own XMLDSig context.
type XMLDocumentSigner struct {
	privateKey  *rsa.PrivateKey
	certificate []byte
}

// ConcurrentXMLSigning reports that SignXML can run concurrently. Each call
// creates an isolated XMLDSig context while reusing immutable key material.
func (*XMLDocumentSigner) ConcurrentXMLSigning() bool {
	return true
}

// NewXMLDocumentSigner creates a reusable signer from PEM-encoded key and
// certificate bytes. The certificate is validated once at construction time.
func NewXMLDocumentSigner(keyBytes, certBytes []byte) (*XMLDocumentSigner, error) {
	privateKey, err := parseRSAPrivateKey(keyBytes)
	if err != nil {
		return nil, err
	}
	privateKey.Precompute()

	blockCert, _ := pem.Decode(certBytes)
	if blockCert == nil {
		return nil, fmt.Errorf("error decoding PEM certificate")
	}
	cert, err := x509.ParseCertificate(blockCert.Bytes)
	if err != nil {
		return nil, err
	}
	if err := VerifyCertificateValidity(cert); err != nil {
		return nil, err
	}

	return &XMLDocumentSigner{
		privateKey:  privateKey,
		certificate: append([]byte(nil), blockCert.Bytes...),
	}, nil
}

// NewXMLDocumentSignerFromP12 creates a reusable signer from PKCS#12 bytes.
func NewXMLDocumentSignerFromP12(p12Data []byte, password string) (*XMLDocumentSigner, error) {
	privateKey, cert, err := pkcs12.Decode(p12Data, password)
	if err != nil {
		return nil, err
	}
	if err := VerifyCertificateValidity(cert); err != nil {
		return nil, err
	}
	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("private key is not of type RSA")
	}
	rsaPrivateKey.Precompute()
	return &XMLDocumentSigner{
		privateKey:  rsaPrivateKey,
		certificate: append([]byte(nil), cert.Raw...),
	}, nil
}

// SignXML signs one XML document using the reusable signing material.
func (s *XMLDocumentSigner) SignXML(xmlBytes []byte) ([]byte, error) {
	ks := &pemKeyStore{PrivateKey: s.privateKey, Cert: s.certificate}
	ctx := dsig.NewDefaultSigningContext(ks)
	ctx.Canonicalizer = dsig.MakeC14N10WithCommentsCanonicalizer()
	ctx.SetSignatureMethod(dsig.RSASHA256SignatureMethod)

	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(xmlBytes); err != nil {
		return nil, err
	}
	signedElement, err := ctx.SignEnveloped(doc.Root())
	if err != nil {
		return nil, err
	}
	signedDoc := etree.NewDocument()
	signedDoc.SetRoot(signedElement)
	var buf bytes.Buffer
	if _, err := signedDoc.WriteTo(&buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (ks *pemKeyStore) GetKeyPair() (*rsa.PrivateKey, []byte, error) {
	return ks.PrivateKey, ks.Cert, nil
}

// SignXMLBytes signs an XML document receiving certificates and key directly in bytes
func SignXMLBytes(xmlBytes, keyBytes, certBytes []byte) ([]byte, error) {
	signer, err := NewXMLDocumentSigner(keyBytes, certBytes)
	if err != nil {
		return nil, err
	}
	return signer.SignXML(xmlBytes)
}

// SignXML signs an XML document receiving certificates and key from files
func SignXML(xmlBytes []byte, keyPath, certPath string) ([]byte, error) {
	// Read key from file
	keyData, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}
	// Read certificate from file
	certData, err := os.ReadFile(certPath)
	if err != nil {
		return nil, err
	}

	return SignXMLBytes(xmlBytes, keyData, certData)
}

// SignWithP12Bytes signs an XML using the bytes of the p12 file
// Ideal for when the certificate comes from a DB or a Vault
func SignWithP12Bytes(xmlBytes, p12Data []byte, password string) ([]byte, error) {
	signer, err := NewXMLDocumentSignerFromP12(p12Data, password)
	if err != nil {
		return nil, err
	}
	return signer.SignXML(xmlBytes)
}

// SignWithP12 acts as a bridge between the p12 file and the existing signing function
func SignWithP12(xmlBytes []byte, p12Path, password string) ([]byte, error) {
	p12Data, err := os.ReadFile(p12Path)
	if err != nil {
		return nil, err
	}

	return SignWithP12Bytes(xmlBytes, p12Data, password)
}

// encodeP12ToPEM is an internal helper to avoid code duplication
func encodeP12ToPEM(priv any, cert *x509.Certificate) ([]byte, []byte, error) {
	rsaPriv, ok := priv.(*rsa.PrivateKey)
	if !ok {
		return nil, nil, fmt.Errorf("private key is not of type RSA")
	}

	keyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(rsaPriv),
	})

	certPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert.Raw,
	})

	return keyPEM, certPEM, nil
}

// parseRSAPrivateKey processes PEM key bytes (PKCS#1 or PKCS#8)
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

// VerifyP12Expiry checks the validity of the certificate within the P12 bytes
func VerifyP12Expiry(p12Data []byte, password string) error {
	_, cert, err := pkcs12.Decode(p12Data, password)
	if err != nil {
		return fmt.Errorf("error decoding for validation: %w", err)
	}

	now := time.Now()
	if now.Before(cert.NotBefore) {
		return fmt.Errorf("certificate is not yet valid (starts: %s)", cert.NotBefore.Format("2006-01-02"))
	}
	if now.After(cert.NotAfter) {
		return fmt.Errorf("certificate expired on: %s", cert.NotAfter.Format("2006-01-02"))
	}
	return nil
}

// VerifyCertificateValidity checks if the certificate is valid today
func VerifyCertificateValidity(cert *x509.Certificate) error {
	now := time.Now()
	if now.Before(cert.NotBefore) {
		return fmt.Errorf("certificate is not yet valid (starts: %s)", cert.NotBefore.Format("2006-01-02"))
	}
	if now.After(cert.NotAfter) {
		return fmt.Errorf("certificate expired on: %s", cert.NotAfter.Format("2006-01-02"))
	}
	return nil
}
