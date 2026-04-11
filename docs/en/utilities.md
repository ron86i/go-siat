# Utilities

[← Back to Index](README.md)

> Reference for the `pkg/utils` package, which provides cryptographic signing, CUF generation, compression, hashing, and parsing helpers.

---

## Table of Contents

1. [XML Digital Signing](#xml-digital-signing)
2. [CUF Generation](#cuf-generation)
3. [Compression and Encoding](#compression-and-encoding)
4. [Cryptographic Hashing](#cryptographic-hashing)
5. [Parsing Helpers](#parsing-helpers)

---

## XML Digital Signing

**Package**: `github.com/ron86i/go-siat/pkg/utils`

The SDK provides multiple ways to digitally sign invoice XML documents. All methods use **RSA-SHA256** with **Enveloped Signature** (C14N 1.0 with Comments), as mandated by the SIAT.

### `SignXML` - Sign from PEM Files

Reads certificate and key from the filesystem.

```go
func SignXML(xmlBytes []byte, keyPath, certPath string) ([]byte, error)
```

| Parameter | Type | Description |
|:----------|:-----|:------------|
| `xmlBytes` | `[]byte` | Serialized XML invoice (from `xml.Marshal()`) |
| `keyPath` | `string` | Path to the PEM private key file |
| `certPath` | `string` | Path to the PEM certificate file |

**Returns**: Signed XML bytes, or error if certificate is expired/invalid.

```go
signedXML, err := utils.SignXML(xmlData, "key.pem", "cert.crt")
```

### `SignXMLBytes` - Sign from PEM Bytes

Receives certificate and key directly as byte slices. Ideal for loading from databases or secret vaults.

```go
func SignXMLBytes(xmlBytes, keyBytes, certBytes []byte) ([]byte, error)
```

```go
keyBytes, _ := os.ReadFile("key.pem")
certBytes, _ := os.ReadFile("cert.crt")
signedXML, err := utils.SignXMLBytes(xmlData, keyBytes, certBytes)
```

### `SignWithP12` - Sign from P12/PFX File

Uses a PKCS#12 container file (combines both key and certificate).

```go
func SignWithP12(xmlBytes []byte, p12Path, password string) ([]byte, error)
```

```go
signedXML, err := utils.SignWithP12(xmlData, "cert.p12", "my_password")
```

### `SignWithP12Bytes` - Sign from P12 Bytes

Receives the P12 data directly in memory. Best for when certificates are stored as BLOBs in databases or fetched from a vault.

```go
func SignWithP12Bytes(xmlBytes, p12Data []byte, password string) ([]byte, error)
```

```go
p12Data, _ := fetchFromVault("siat-cert")
signedXML, err := utils.SignWithP12Bytes(xmlData, p12Data, "my_password")
```

### `VerifyP12Expiry` - Check Certificate Validity

Validates that the certificate inside a P12 file hasn't expired.

```go
func VerifyP12Expiry(p12Data []byte, password string) error
```

```go
if err := utils.VerifyP12Expiry(p12Data, "password"); err != nil {
    log.Printf("Certificate issue: %v", err)
    // "certificate expired on: 2026-01-15" or
    // "certificate is not yet valid (starts: 2026-06-01)"
}
```

### `VerifyCertificateValidity` - Check x509 Certificate

Validates an already-parsed x509 certificate.

```go
func VerifyCertificateValidity(cert *x509.Certificate) error
```

> [!IMPORTANT]
> All signing functions automatically validate the certificate before signing. If the certificate is expired or not yet valid, the signing will fail with a descriptive error.

### Supported Key Formats

| Format | Support |
|:-------|:--------|
| PKCS#1 (RSA PRIVATE KEY) | ✅ |
| PKCS#8 (PRIVATE KEY) | ✅ |
| PKCS#12 (.p12 / .pfx) | ✅ |
| EC Keys | ❌ (RSA only) |

---

## CUF Generation

The CUF (Código Único de Factura) is a unique identifier for each invoice, calculated using a specific algorithm mandated by the SIAT.

### `GenerarCUF`

```go
func GenerarCUF(
    nit int64,
    fechaHora time.Time,
    sucursal, modalidad, tipoEmision, tipoFactura,
    tipoDocumentoSector, numeroFactura, puntoVenta int,
    codigoControl string,
) (string, error)
```

| Parameter | Type | Digits | Description |
|:----------|:-----|:-------|:------------|
| `nit` | `int64` | 13 | Taxpayer NIT (zero-padded) |
| `fechaHora` | `time.Time` | 17 | Emission date/time (yyyyMMddHHmmssSSS) |
| `sucursal` | `int` | 4 | Branch number (zero-padded) |
| `modalidad` | `int` | 1 | 1=Electronic, 2=Computerized |
| `tipoEmision` | `int` | 1 | 1=Online, 2=Offline, 3=Massive |
| `tipoFactura` | `int` | 1 | Invoice type |
| `tipoDocumentoSector` | `int` | 2 | Sector document type |
| `numeroFactura` | `int` | 10 | Invoice number (zero-padded) |
| `puntoVenta` | `int` | 4 | Point of sale (zero-padded) |
| `codigoControl` | `string` | Variable | Control code from CUFD response |

**Returns**: The complete CUF string (hex + control code).

### Algorithm

1. **Format** each field to its fixed length with zero-padding.
2. **Concatenate** all fields (53 digits total).
3. **Calculate** Modulo 11 verification digit (weights 2-9).
4. **Convert** the numeric string to hexadecimal (Base16).
5. **Append** the CUFD control code.

```go
cuf, err := utils.GenerarCUF(
    123456789,              // NIT
    time.Now(),             // Emission date/time
    0,                      // Branch
    siat.ModalidadElectronica, // Modality (1)
    siat.EmisionOnline,     // Emission type (1)
    1,                      // Invoice type
    1,                      // Sector document type
    42,                     // Invoice number
    0,                      // Point of sale
    cufdControl,            // Control code from CUFD
)
```

---

## Compression and Encoding

### `Gzip` - Compress Data

```go
func Gzip(data []byte) ([]byte, error)
```

Standard Gzip compression.

```go
compressed, err := utils.Gzip(xmlBytes)
```

### `CompressAndHash` - Compress + SHA256 + Base64

The all-in-one function for preparing invoice data for transmission. This combines the three steps SIAT requires:

```go
func CompressAndHash(data []byte) (hash, encoded string, err error)
```

| Return | Type | Description |
|:-------|:-----|:------------|
| `hash` | `string` | SHA-256 hex hash of the **compressed** data |
| `encoded` | `string` | Base64 encoding of the compressed data |
| `err` | `error` | Any compression error |

```go
hash, archivoBase64, err := utils.CompressAndHash(signedXML)
// hash → used in WithHashArchivo()
// archivoBase64 → used in WithArchivo()
```

### `CreateTarGz` - Create TAR.GZ Archive

Creates a TAR.GZ archive from a map of filenames to contents. Useful for packaging multiple invoices.

```go
func CreateTarGz(files map[string][]byte) ([]byte, error)
```

```go
files := map[string][]byte{
    "factura_001.xml": xmlBytes1,
    "factura_002.xml": xmlBytes2,
}
archive, err := utils.CreateTarGz(files)
```

---

## Cryptographic Hashing

### `SHA256Hex`

```go
func SHA256Hex(data []byte) string
```

Returns the SHA-256 hash as a lowercase hex string.

```go
hash := utils.SHA256Hex(xmlBytes) // "a1b2c3d4..."
```

### `SHA512Hex`

```go
func SHA512Hex(data []byte) string
```

Returns the SHA-512 hash as a lowercase hex string.

```go
hash := utils.SHA512Hex(xmlBytes)
```

---

## Parsing Helpers

### `ParseIntSafe`

Safely converts a string to `int`, trimming whitespace and providing descriptive errors.

```go
func ParseIntSafe(valStr string) (int, error)
```

```go
ambiente, err := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
// Returns error for: "", "  ", "abc", "12.5"
```

### `ParseInt64Safe`

Safely converts a string to `int64`. Ideal for NIT values.

```go
func ParseInt64Safe(valStr string) (int64, error)
```

```go
nit, err := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
```

### Pointer Helpers

Functions that return pointers to values. Useful for optional fields in invoice builders:

```go
func Float64Ptr(v float64) *float64
func Int64Ptr(v int64) *int64
func IntPtr(v int) *int
```

```go
cabecera := invoices.NewCompraVentaCabeceraBuilder().
    WithMontoDescuento(utils.Float64Ptr(10.50)).   // Optional field
    WithNombreRazonSocial(nil).                     // Explicitly nil
    Build()
```

---

[← Back to Index](README.md) | [Next: Configuration →](configuration.md)
