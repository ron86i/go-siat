# Cryptography & Utils

One of the most powerful features of `go-siat` is its built-in handling of BOLIVIAN specific security requirements (XMLDSig, CUF generation, and Control Codes).

## 🔏 XML Digital Signature (XMLDSig)

The SDK uses `xml-crypto` (embedded) to sign and verify SIAT invoices using a standard PKCS#12 (`.p12` or `.pfx`) certificate.

```go
signedXML, err := utils.SignXML(rawXML, certData, certPassword)
if err != nil {
    // Handle signing error
}
```

## 🆔 CUF Generation

The **Unique Invoice Code (CUF)** generation follows SIAT's Base11 encoding and algorithmic requirements.

```go
cuf, err := utils.GenerateCUF(nit, fechaCuf, sucursal, modalidad, tipoEmision, tipoFactura, puntoVenta, codigoControl)
```

## 🔐 Control Code (Computarizada)

For computerized invoices, the SDK generates the mandatory hash (SHA-256) used for reception without digital signatures.

```go
hash := utils.ComputeHash(xmlContent)
```

## 📅 Date Parsing

Utility to format Go `time.Time` objects to the specific SIAT ISO 8601 format with milliseconds: `2026-03-26T14:30:15.000`.

```go
formated := utils.FormatSiatTime(time.Now())
```
