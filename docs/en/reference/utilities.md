# Utilities reference

<p align="right">
  <a href="../../es/reference/utilidades.md">🇪🇸 Español</a> · <a href="../README.md">Docs index</a>
</p>

Complete surface of `github.com/ron86i/go-siat/v3/pkg/utils` — CUF generation, signing, compression, hashing, export and small conversion helpers.

```go
import "github.com/ron86i/go-siat/v3/pkg/utils"
```

Most of this runs automatically inside `WithFactura` / `WithFacturas` / `WithDocumento`. Reach for these directly when you need the intermediate artifacts or are working outside the normal send flow.

---

## CUF generation

The CUF is the per-invoice fiscal checksum. You compute it locally and put it in the invoice header before sending.

### Fluent builder

```go
cuf, err := utils.NewCUF().
    WithNit(123456789).
    WithFechaHora(fechaEmision).
    WithSucursal(0).
    WithModalidad(siat.ModalidadElectronica).
    WithTipoEmision(siat.EmisionOnline).
    WithTipoFactura(1).
    WithTipoDocumentoSector(1).
    WithNumeroFactura(1).
    WithPuntoVenta(0).
    WithCodigoControl(cufd.CodigoControl).
    Generate()
```

| Method | Type |
| :--- | :--- |
| `WithNit` | `int64` |
| `WithFechaHora` | `time.Time` |
| `WithSucursal` | `int` |
| `WithModalidad` | `int` |
| `WithTipoEmision` | `int` |
| `WithTipoFactura` | `int` |
| `WithTipoDocumentoSector` | `int` |
| `WithNumeroFactura` | `int64` |
| `WithPuntoVenta` | `int` |
| `WithCodigoControl` | `string` |
| `Generate` | → `(string, error)` |

**The `fechaHora` you pass here must be the exact same value you put in the header's `fechaEmision`.** They are hashed into the CUF, and any drift produces code 1002 (`CufInvalido`). Compute the timestamp once into a variable and reuse it.

`WithCodigoControl` takes the `CodigoControl` from the CUFD response, not from the CUFD code itself.

### Positional form

```go
func GenerarCUF(
    nit int64,
    fechaHora time.Time,
    sucursal, modalidad, tipoEmision, tipoFactura, tipoDocumentoSector int,
    numeroFactura int64,
    puntoVenta int,
    codigoControl string,
) (string, error)
```

Same result, ten positional arguments. The builder is safer — several of these are `int` and trivially transposable.

---

## XML signing

```go
type XMLSigner interface {
    SignXML(xmlBytes []byte) ([]byte, error)
}
```

`siat.Config` implements it, which is why `s.Config()` is what you pass wherever a signer is required.

| Function | Signature |
| :--- | :--- |
| `SignWithP12` | `(xmlBytes []byte, p12Path, password string) ([]byte, error)` |
| `SignWithP12Bytes` | `(xmlBytes, p12Data []byte, password string) ([]byte, error)` |
| `SignXML` | `(xmlBytes []byte, keyPath, certPath string) ([]byte, error)` |
| `SignXMLBytes` | `(xmlBytes, keyBytes, certBytes []byte) ([]byte, error)` |

Note the argument order on the PEM pair: `SignXML` takes **key first, then certificate**; `SignXMLBytes` likewise. Reversing them fails at parse time.

Signatures are RSA-SHA256 enveloped XMLDSig, the algorithm SIAT mandates.

### Certificate validation

```go
func VerifyP12Expiry(p12Data []byte, password string) error
func VerifyCertificateValidity(cert *x509.Certificate) error
```

Check at startup rather than discovering an expired certificate mid-send:

```go
p12, err := os.ReadFile("cert.p12")
if err != nil {
    log.Fatal(err)
}
if err := utils.VerifyP12Expiry(p12, password); err != nil {
    log.Fatal("certificate unusable: ", err)
}
```

`VerifyP12Expiry` takes bytes, not a path.

---

## Compression and hashing

```go
func Gzip(data []byte) ([]byte, error)
func CompressAndHash(data []byte) (hash, encoded string, err error)
func CreateTarGz(files map[string][]byte) ([]byte, error)
func SHA256Hex(data []byte) string
func SHA512Hex(data []byte) string
```

`CompressAndHash` is the one you want for a submission: it gzips, base64-encodes and hashes in one pass, returning exactly the two strings the request builders expect.

```go
hash, encoded, err := utils.CompressAndHash(signedXML)
builder.WithArchivo(encoded).WithHashArchivo(hash)
```

Mind the return order — `hash` first, then `encoded`. Both are strings, so swapping them compiles and fails only at SIAT with code 969.

`CreateTarGz` takes a filename-to-content map and returns the archive bytes, for building batch packages by hand.

---

## Export

```go
func ExportXML(factura any, path string) error
func ExportSignedXML(factura any, signer XMLSigner, path string) error
func ExportTarGz(facturas []any, signer XMLSigner, path string) error
```

Write documents to disk — for archiving, for offline submission, or for inspecting what actually gets sent.

```go
utils.ExportXML(factura, "invoice-001.xml")                        // unsigned
utils.ExportSignedXML(factura, s.Config(), "invoice-001.xml")      // signed
utils.ExportTarGz([]any{f1, f2, f3}, s.Config(), "batch.tar.gz")   // package
```

Pass `nil` as the signer to `ExportTarGz` for an unsigned package.

---

## Conversion helpers

```go
func ParseIntSafe(valStr string) (int, error)
func ParseInt64Safe(valStr string) (int64, error)
func Round(v float64, decimals int) float64
```

`Round` matters more than it looks: SIAT recomputes your totals and rejects mismatches with codes 1013, 1018 and 1024. Round each amount to two decimals consistently rather than letting float error accumulate across a long detail list.

### Pointer constructors

Many invoice header fields are optional and therefore pointers. These give you an inline address:

```go
func IntPtr(v int) *int
func Int64Ptr(v int64) *int64
func Float64Ptr(v float64) *float64
```

```go
cabecera := invoices.NewCompraVentaCabeceraBuilder().
    WithCodigoPuntoVenta(utils.IntPtr(0)).
    Build()
```

There is no `StringPtr` — take the address of a local for those:

```go
nombre := "CUSTOMER NAME"
cabecera.WithNombreRazonSocial(&nombre)
```

A nil pointer means the element is omitted from the XML. That is meaningful: for genuinely optional fields, nil and a zero value are not the same document.

---

## Related

| | |
| :--- | :--- |
| CUF in a working flow | [Tutorial → Step 4](../tutorial/first-invoice.md) |
| Certificates and signing | [How-to: Sign invoices](../how-to/sign-invoices.md) |
| Building packages by hand | [How-to: Send batches](../how-to/send-batches.md) |
| Codes 969, 1002, 1013 | [How-to: Handle errors](../how-to/handle-errors.md) |
