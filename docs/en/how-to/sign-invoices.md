# How to sign invoices with a digital certificate

<p align="right">
  <a href="../../es/how-to/firmar-facturas.md">🇪🇸 Español</a> · <a href="../README.md">Docs index</a>
</p>

Electronic-modality invoices must carry an **XMLDSig** signature made with your SIAT-issued certificate. This guide shows how to load that certificate, how signing gets applied, and how to handle the cases that are easy to get wrong.

> Computerized modality does **not** need a signature. If that is your case, you can skip this guide entirely.

---

## Load your certificate

The SDK accepts the two formats SIAT issues. Both produce a `CredentialSign` that you put in your `Config`.

### PKCS#12 (`.p12` / `.pfx`)

```go
s, err := siat.New(siat.Config{
    Token:          "...",
    Nit:            123456789,
    CodigoSistema:  "...",
    CodigoAmbiente: siat.AmbientePruebas,
    BaseURL:        "https://pilotosiatservicios.impuestos.gob.bo/sfd",

    CredentialSign: siat.NewP12Credential("cert.p12", "your-password"),
})
```

### Separate certificate and key (PEM)

```go
CredentialSign: siat.NewPEMCredential("cert.crt", "key.pem"),
```

### Loading from memory instead of disk

Both constructors accept either a **file path** (`string`) or **raw bytes** (`[]byte`). This matters when your certificate comes from a secret manager, an environment variable, or a database rather than the filesystem:

```go
p12Bytes := mustFetchFromVault("siat/cert")           // []byte
cred := siat.NewP12Credential(p12Bytes, os.Getenv("P12_PASSWORD"))

certBytes, keyBytes := loadFromSecrets()               // []byte, []byte
cred := siat.NewPEMCredential(certBytes, keyBytes)
```

---

## How signing actually happens

You rarely call the signer yourself. `WithFactura` does it for you:

```go
builder := models.NewRecepcionFacturaBuilder().
    WithCodigoModalidad(siat.ModalidadElectronica).   // ← must come first
    WithCodigoSucursal(0).
    WithCufd(cufd).
    WithCuis(cuis)

// serialize → sign → gzip → base64 → hash, in one call
if err := builder.WithFactura(invoice, s.Config()); err != nil {
    log.Fatal(err)
}
```

Two rules govern this, and both are silent when broken:

**`WithCodigoModalidad` must be called before `WithFactura`.** The signing step reads the modality from the request being built. If the modality is still zero when `WithFactura` runs, it concludes no signature is needed and sends an unsigned document — which SIAT rejects.

**`s.Config()` is what carries the credentials.** `Config` implements the signer interface by delegating to its `CredentialSign`. Passing anything else, or passing a `Config` with no credential, means no signature.

---

## Verify a certificate before you rely on it

An expired certificate fails at submission time, which is the worst moment to find out. Check it at startup instead:

```go
import "github.com/ron86i/go-siat/v3/pkg/utils"

p12Bytes, err := os.ReadFile("cert.p12")
if err != nil {
    log.Fatal(err)
}
if err := utils.VerifyP12Expiry(p12Bytes, "your-password"); err != nil {
    log.Fatal("certificate unusable:", err)
}
```

You can also confirm which format a credential ended up as — useful when the source is dynamic:

```go
cred := siat.NewP12Credential(data, password)
switch cred.GetType() {
case "P12": // PKCS#12 loaded
case "PEM": // certificate + key loaded
case "UNKNOWN":
    log.Fatal("no usable credential was loaded")
}
```

`GetType()` returning `"UNKNOWN"` is the signal that a path was wrong or the bytes were empty. The constructors do not panic on a missing file — they record the error and surface it when you first try to sign.

---

## Sign an XML document directly

When you need the signature outside the invoice flow — to archive a signed copy, or to inspect what gets sent — call the signer yourself:

```go
signed, err := s.Config().SignXML(xmlBytes)
```

Or use the `utils` helpers, which take the certificate explicitly:

| Function | Use when |
| :------- | :------- |
| `utils.SignWithP12(xmlBytes, p12Path, password)` | `.p12` on disk |
| `utils.SignWithP12Bytes(xmlBytes, p12Data, password)` | `.p12` already in memory |
| `utils.SignXML(xmlBytes, keyPath, certPath)` | PEM pair on disk |
| `utils.SignXMLBytes(xmlBytes, keyBytes, certBytes)` | PEM pair in memory |

---

## Export a signed invoice to a file

To keep a signed copy on disk — for archiving or for offline submission:

```go
// single invoice, signed
err := utils.ExportSignedXML(invoice, s.Config(), "invoice-001.xml")

// unsigned
err := utils.ExportXML(invoice, "invoice-001.xml")

// many invoices as a .tar.gz package
err := utils.ExportTarGz([]any{inv1, inv2, inv3}, s.Config(), "batch.tar.gz")
```

Pass `nil` as the signer to `ExportTarGz` when you want the package unsigned.

---

## Troubleshooting

| Symptom | Cause |
| :------ | :---- |
| SIAT rejects with a signature error, no local error | `WithFactura` ran before `WithCodigoModalidad` |
| `no se configuraron credenciales válidas de firma digital` | `Config.CredentialSign` is the zero value, or both loads failed |
| `error al leer archivo P12` | Wrong path — the error surfaces at signing time, not at construction |
| Signature succeeds locally, SIAT still refuses | Certificate expired, or issued for a different NIT than `Config.Nit` |

---

## Related

| | |
| :--- | :--- |
| Full flow including signing | [Tutorial → Step 7](../tutorial/first-invoice.md#step-7-switch-to-electronic-with-digital-signature) |
| Signing and hashing helpers | [Reference: Utilities](../reference/utilities.md) |
| Interpreting SIAT rejections | [How-to: Handle errors](handle-errors.md) |
