package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// generateTestCertificate creates a valid RSA key pair and self-signed certificate for testing.
func generateTestCertificate() (keyPEM, certPEM []byte, privKey *rsa.PrivateKey, cert *x509.Certificate, err error) {
	privKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"Test Org"},
		},
		NotBefore: time.Now().Add(-1 * time.Hour),
		NotAfter:  time.Now().Add(1 * time.Hour),

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, privKey.Public(), privKey)
	if err != nil {
		return
	}

	cert, _ = x509.ParseCertificate(derBytes)

	keyPEM = pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privKey)})
	certPEM = pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})

	return
}

// ─── SignXMLBytes Tests ──────────────────────────────────────────────────────

func TestSignXMLBytes(t *testing.T) {
	keyPEM, certPEM, _, _, err := generateTestCertificate()
	require.NoError(t, err)

	xmlInput := []byte("<root><node>value</node></root>")

	signedXML, err := SignXMLBytes(xmlInput, keyPEM, certPEM)
	assert.NoError(t, err)
	assert.Greater(t, len(signedXML), len(xmlInput), "Signed XML should be larger than input")
}

func TestSignXMLBytes_InvalidKeyPEM(t *testing.T) {
	_, certPEM, _, _, err := generateTestCertificate()
	require.NoError(t, err)

	_, err = SignXMLBytes([]byte("<test/>"), []byte("not-a-pem"), certPEM)
	assert.Error(t, err, "Should fail with invalid key PEM")
}

func TestSignXMLBytes_InvalidCertPEM(t *testing.T) {
	keyPEM, _, _, _, err := generateTestCertificate()
	require.NoError(t, err)

	_, err = SignXMLBytes([]byte("<test/>"), keyPEM, []byte("not-a-cert-pem"))
	assert.Error(t, err, "Should fail with invalid cert PEM")
}

func TestSignXMLBytes_InvalidXML(t *testing.T) {
	keyPEM, certPEM, _, _, err := generateTestCertificate()
	require.NoError(t, err)

	_, err = SignXMLBytes([]byte("<<<not xml>>>"), keyPEM, certPEM)
	assert.Error(t, err, "Should fail with invalid XML")
}

func TestSignXMLBytes_ExpiredCert(t *testing.T) {
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	template := x509.Certificate{
		SerialNumber: big.NewInt(2),
		Subject:      pkix.Name{Organization: []string{"Expired Org"}},
		NotBefore:    time.Now().Add(-48 * time.Hour),
		NotAfter:     time.Now().Add(-24 * time.Hour),
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}
	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, privKey.Public(), privKey)
	require.NoError(t, err)

	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privKey)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})

	_, err = SignXMLBytes([]byte("<root/>"), keyPEM, certPEM)
	assert.Error(t, err, "Should fail with expired certificate")
	assert.Contains(t, err.Error(), "expired")
}

func TestSignXMLBytes_FutureCert(t *testing.T) {
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	template := x509.Certificate{
		SerialNumber: big.NewInt(3),
		Subject:      pkix.Name{Organization: []string{"Future Org"}},
		NotBefore:    time.Now().Add(24 * time.Hour),
		NotAfter:     time.Now().Add(48 * time.Hour),
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}
	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, privKey.Public(), privKey)
	require.NoError(t, err)

	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privKey)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})

	_, err = SignXMLBytes([]byte("<root/>"), keyPEM, certPEM)
	assert.Error(t, err, "Should fail with future certificate")
	assert.Contains(t, err.Error(), "not yet valid")
}

func TestSignXMLBytes_CorruptCertDER(t *testing.T) {
	keyPEM, _, _, _, err := generateTestCertificate()
	require.NoError(t, err)

	// Valid PEM structure but corrupt DER certificate content
	corruptCertPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: []byte("corrupt-der-data")})
	_, err = SignXMLBytes([]byte("<root/>"), keyPEM, corruptCertPEM)
	assert.Error(t, err, "Should fail with corrupt certificate DER")
}

// ─── SignXML (from files) Tests ──────────────────────────────────────────────

func TestSignXMLFromFile(t *testing.T) {
	keyPEM, certPEM, _, _, err := generateTestCertificate()
	require.NoError(t, err)

	keyFile := "test_key.pem"
	certFile := "test_cert.pem"
	require.NoError(t, os.WriteFile(keyFile, keyPEM, 0644))
	require.NoError(t, os.WriteFile(certFile, certPEM, 0644))
	defer os.Remove(keyFile)
	defer os.Remove(certFile)

	xmlInput := []byte("<root>test</root>")
	signed, err := SignXML(xmlInput, keyFile, certFile)
	assert.NoError(t, err)
	assert.NotEmpty(t, signed)
}

func TestSignXML_NonExistentKeyFile(t *testing.T) {
	_, err := SignXML([]byte("<test/>"), "non_existent_key.pem", "non_existent_cert.pem")
	assert.Error(t, err, "Should fail with non-existent key file")
}

func TestSignXML_NonExistentCertFile(t *testing.T) {
	keyPEM, _, _, _, err := generateTestCertificate()
	require.NoError(t, err)

	keyFile := "test_key_only.pem"
	require.NoError(t, os.WriteFile(keyFile, keyPEM, 0644))
	defer os.Remove(keyFile)

	_, err = SignXML([]byte("<test/>"), keyFile, "non_existent_cert.pem")
	assert.Error(t, err, "Should fail when cert file does not exist")
}

// ─── SignWithP12Bytes Error Tests ────────────────────────────────────────────

func TestSignWithP12Bytes_InvalidData(t *testing.T) {
	_, err := SignWithP12Bytes([]byte("<test/>"), []byte("not-p12-data"), "pass")
	assert.Error(t, err, "Should fail with invalid P12 data")
}

// ─── SignWithP12 Error Tests ─────────────────────────────────────────────────

func TestSignWithP12_NonExistentFile(t *testing.T) {
	_, err := SignWithP12([]byte("<test/>"), "non_existent.p12", "pass")
	assert.Error(t, err, "Should fail with non-existent P12 file")
}

func TestSignWithP12_InvalidFileContent(t *testing.T) {
	p12File := "test_invalid.p12"
	require.NoError(t, os.WriteFile(p12File, []byte("not-real-p12"), 0644))
	defer os.Remove(p12File)

	_, err := SignWithP12([]byte("<test/>"), p12File, "pass")
	assert.Error(t, err, "Should fail with invalid P12 file content")
}

// ─── encodeP12ToPEM Tests ────────────────────────────────────────────────────

func TestEncodeP12ToPEM_Valid(t *testing.T) {
	_, _, privKey, cert, err := generateTestCertificate()
	require.NoError(t, err)

	keyPEM, certPEM, err := encodeP12ToPEM(privKey, cert)
	assert.NoError(t, err)
	assert.NotEmpty(t, keyPEM)
	assert.NotEmpty(t, certPEM)

	// Verify the PEM is valid
	keyBlock, _ := pem.Decode(keyPEM)
	assert.NotNil(t, keyBlock, "Key PEM should decode correctly")
	assert.Equal(t, "RSA PRIVATE KEY", keyBlock.Type)

	certBlock, _ := pem.Decode(certPEM)
	assert.NotNil(t, certBlock, "Cert PEM should decode correctly")
	assert.Equal(t, "CERTIFICATE", certBlock.Type)
}

func TestEncodeP12ToPEM_NonRSAKey(t *testing.T) {
	// Generate an ECDSA key (non-RSA) to trigger the "not of type RSA" error
	ecPriv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err)

	_, _, _, cert, err := generateTestCertificate()
	require.NoError(t, err)

	_, _, err = encodeP12ToPEM(ecPriv, cert)
	assert.Error(t, err, "Should fail with non-RSA private key")
	assert.Contains(t, err.Error(), "not of type RSA")
}

// ─── parseRSAPrivateKey Tests ────────────────────────────────────────────────

func TestParseRSAPrivateKey_PKCS1(t *testing.T) {
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	pkcs1PEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privKey),
	})

	parsed, err := parseRSAPrivateKey(pkcs1PEM)
	assert.NoError(t, err)
	assert.NotNil(t, parsed)
	assert.True(t, privKey.Equal(parsed))
}

func TestParseRSAPrivateKey_PKCS8(t *testing.T) {
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	pkcs8Bytes, err := x509.MarshalPKCS8PrivateKey(privKey)
	require.NoError(t, err)

	pkcs8PEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: pkcs8Bytes,
	})

	parsed, err := parseRSAPrivateKey(pkcs8PEM)
	assert.NoError(t, err)
	assert.NotNil(t, parsed)
	assert.True(t, privKey.Equal(parsed))
}

func TestParseRSAPrivateKey_PKCS8_NonRSA(t *testing.T) {
	// PKCS#8 with an ECDSA key (non-RSA)
	ecPriv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err)

	pkcs8Bytes, err := x509.MarshalPKCS8PrivateKey(ecPriv)
	require.NoError(t, err)

	pkcs8PEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: pkcs8Bytes,
	})

	_, err = parseRSAPrivateKey(pkcs8PEM)
	assert.Error(t, err, "Should fail with non-RSA key in PKCS#8 format")
	assert.Contains(t, err.Error(), "not of type RSA")
}

func TestParseRSAPrivateKey_InvalidPEM(t *testing.T) {
	_, err := parseRSAPrivateKey([]byte("Not a PEM"))
	assert.Error(t, err, "Should fail with non-PEM data")
	assert.Contains(t, err.Error(), "invalid PEM format")
}

func TestParseRSAPrivateKey_CorruptDER(t *testing.T) {
	invalidKey := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: []byte("corrupt-der-data"),
	})
	_, err := parseRSAPrivateKey(invalidKey)
	assert.Error(t, err, "Should fail parsing corrupt DER data")
}

// ─── VerifyP12Expiry Tests ───────────────────────────────────────────────────

func TestVerifyP12Expiry_InvalidData(t *testing.T) {
	err := VerifyP12Expiry([]byte("not-p12-data"), "pass")
	assert.Error(t, err, "Should fail with invalid P12 data")
	assert.Contains(t, err.Error(), "error decoding")
}

// ─── VerifyCertificateValidity Tests ─────────────────────────────────────────

func TestVerifyCertificateValidity_Valid(t *testing.T) {
	_, _, _, cert, err := generateTestCertificate()
	require.NoError(t, err)

	err = VerifyCertificateValidity(cert)
	assert.NoError(t, err)
}

func TestVerifyCertificateValidity_NotYetValid(t *testing.T) {
	_, _, _, cert, err := generateTestCertificate()
	require.NoError(t, err)

	cert.NotBefore = time.Now().Add(1 * time.Hour)
	cert.NotAfter = time.Now().Add(2 * time.Hour)

	err = VerifyCertificateValidity(cert)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not yet valid")
}

func TestVerifyCertificateValidity_Expired(t *testing.T) {
	_, _, _, cert, err := generateTestCertificate()
	require.NoError(t, err)

	cert.NotBefore = time.Now().Add(-2 * time.Hour)
	cert.NotAfter = time.Now().Add(-1 * time.Hour)

	err = VerifyCertificateValidity(cert)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "expired")
}

// ─── GetKeyPair Tests ────────────────────────────────────────────────────────

func TestGetKeyPair(t *testing.T) {
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	certBytes := []byte("test-cert-bytes")
	ks := &pemKeyStore{PrivateKey: privKey, Cert: certBytes}

	key, cert, err := ks.GetKeyPair()
	assert.NoError(t, err)
	assert.Equal(t, privKey, key)
	assert.Equal(t, certBytes, cert)
}
