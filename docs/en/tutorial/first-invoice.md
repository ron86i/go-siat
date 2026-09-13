# Tutorial: From Zero to Your First Invoice

<p align="right">
  <a href="../../es/tutorial/primera-factura.md">🇪🇸 Español</a>
</p>

In this tutorial you will send a real invoice to SIAT's pilot environment and get back an acceptance code.

You will do it in seven steps, and by the end you will understand the chain that every Bolivian electronic invoice must follow: **CUIS → CUFD → CUF → invoice → submission**.

> This is a *learning* exercise. We use the pilot environment, so nothing you do here has tax validity. Once you finish, the [how-to guides](../how-to/) will show you how to handle the same steps in production.

---

## What you need before starting

| Requirement | Notes |
| :---------- | :---- |
| Go 1.26+ | `go version` to check |
| A SIAT API token | Obtained from your SIAT taxpayer portal |
| Your NIT | The taxpayer number the token belongs to |
| An authorized system code | `CodigoSistema`, issued by SIAT when you register your system |
| A digital certificate | Only for the *electronic* modality. We start with **computerized**, which needs none |

We deliberately start with the **computerized** modality so you can complete the whole flow without a certificate. Adding the signature is a one-line change we make at the end.

---

## Step 1: Install and configure the client

```bash
go get github.com/ron86i/go-siat/v2
```

Create `main.go`. The most important idea in this SDK: **your identity is configured once**, not repeated on every call.

```go
package main

import (
    "context"
    "log"
    "time"

    "github.com/ron86i/go-siat/v2"
)

const pilotURL = "https://pilotosiatservicios.impuestos.gob.bo/v2"

func main() {
    s, err := siat.New(siat.Config{
        Token:          "YOUR_SIAT_TOKEN",
        Nit:            123456789,
        CodigoSistema:  "YOUR_SYSTEM_CODE",
        CodigoAmbiente: siat.AmbientePruebas,
        BaseURL:        pilotURL,
    })
    if err != nil {
        log.Fatal(err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    _ = s // we start using it in the next step
}
```

`siat.New` validates your config immediately. If you forget the token or pass an invalid environment code, you get an error here rather than a confusing failure later.

> **Always pass a context with a timeout.** SIAT's servers can be slow, and a bare `context.Background()` will let a request hang indefinitely.

---

## Step 2: Request the CUIS

The **CUIS** (*Código Único de Inicio de Sistemas*) identifies your system to SIAT. Every later call needs it.

```go
cuisReq := models.NewCuisBuilder().
    WithCodigoSucursal(0).      // 0 = head office
    WithCodigoPuntoVenta(0).    // 0 = default point of sale
    WithCodigoModalidad(siat.ModalidadComputarizada).
    Build()

cuisResp, err := s.Codigos().SolicitudCuis(ctx, cuisReq)
if err != nil {
    log.Fatal("transport error:", err)
}
if err := siat.Verify(cuisResp.Body.Content.RespuestaCuis); err != nil {
    log.Fatal("SIAT rejected the CUIS request:", err)
}

cuis := cuisResp.Body.Content.RespuestaCuis.Codigo
log.Println("CUIS:", cuis)
```

Add `"github.com/ron86i/go-siat/v2/pkg/models"` to your imports.

Notice the **two-stage error check**. This pattern repeats in every single call you will make:

1. `err != nil` means the request never completed — network failure, timeout, bad URL.
2. `siat.Verify(...)` means the request completed, but SIAT *rejected* it — invalid NIT, expired code, unauthorized system.

Skipping the second check is the most common mistake when integrating with SIAT: your program appears to succeed while SIAT actually refused the operation.

---

## Step 3: Request the CUFD

The **CUFD** (*Código Único de Facturación Diaria*) is a daily code. It expires, so production systems request a fresh one each day and cache it.

```go
cufdReq := models.NewCufdBuilder().
    WithCuis(cuis).
    WithCodigoSucursal(0).
    WithCodigoPuntoVenta(0).
    WithCodigoModalidad(siat.ModalidadComputarizada).
    Build()

cufdResp, err := s.Codigos().SolicitudCufd(ctx, cufdReq)
if err != nil {
    log.Fatal(err)
}
if err := siat.Verify(cufdResp.Body.Content.RespuestaCufd); err != nil {
    log.Fatal(err)
}

cufd := cufdResp.Body.Content.RespuestaCufd.Codigo
codigoControl := cufdResp.Body.Content.RespuestaCufd.CodigoControl
```

This call returns **two** values you need: the code itself and a `CodigoControl`. Keep both — the control code feeds directly into the next step.

---

## Step 4: Generate the CUF

The **CUF** (*Código Único de Facturación*) is a checksum that uniquely identifies one invoice. Unlike CUIS and CUFD, you compute it locally — no network call.

```go
import "github.com/ron86i/go-siat/v2/pkg/utils"

fechaEmision := time.Now()

cuf, err := utils.NewCUF().
    WithNit(123456789).
    WithFechaHora(fechaEmision).
    WithSucursal(0).
    WithModalidad(siat.ModalidadComputarizada).
    WithTipoEmision(siat.EmisionOnline).
    WithTipoFactura(1).            // 1 = standard invoice with tax credit
    WithTipoDocumentoSector(1).    // 1 = sales sector
    WithNumeroFactura(1).          // your own sequential counter
    WithPuntoVenta(0).
    WithCodigoControl(codigoControl).
    Generate()
if err != nil {
    log.Fatal(err)
}
```

`WithNumeroFactura` is **your** sequential invoice number. You own this counter — SIAT does not assign it. It must never repeat for the same point of sale.

> The exact `fechaEmision` you pass here must be the same timestamp you put in the invoice header in the next step. A mismatch invalidates the CUF.

---

## Step 5: Build the invoice document

Now the invoice itself. It has a **header** (who, when, totals) and one or more **detail lines** (what was sold).

```go
import "github.com/ron86i/go-siat/v2/pkg/models/invoices"

customerName := "JUAN PEREZ"
posCode := 0

header := invoices.NewCompraVentaCabeceraBuilder().
    WithNitEmisor(123456789).
    WithRazonSocialEmisor("YOUR COMPANY NAME").
    WithMunicipio("Tarija").
    WithNumeroFactura(1).                 // same number used in the CUF
    WithCuf(cuf).
    WithCufd(cufd).
    WithCodigoSucursal(0).
    WithDireccion("AVENIDA LA PAZ #123").
    WithCodigoPuntoVenta(&posCode).
    WithFechaEmision(fechaEmision).       // same timestamp used in the CUF
    WithNombreRazonSocial(&customerName).
    WithCodigoTipoDocumentoIdentidad(1).  // 1 = CI
    WithNumeroDocumento("5115889").
    WithCodigoCliente("1").
    WithCodigoMetodoPago(1).              // 1 = cash
    WithMontoTotal(100).
    WithMontoTotalSujetoIva(100).
    WithCodigoMoneda(1).                  // 1 = BOB
    WithTipoCambio(1).
    WithMontoTotalMoneda(100).
    WithLeyenda("Ley N° 453: Tienes derecho a recibir información...").
    WithUsuario("cashier01").
    WithCodigoDocumentoSector(1).
    Build()

line := invoices.NewCompraVentaDetalleBuilder().
    WithActividadEconomica("477300").
    WithCodigoProductoSin(622539).
    WithCodigoProducto("abc123").
    WithDescripcion("GAUZE").
    WithCantidad(1).
    WithUnidadMedida(1).
    WithPrecioUnitario(100).
    WithSubTotal(100).
    Build()

invoice := invoices.NewCompraVentaBuilder().
    WithModalidad(siat.ModalidadComputarizada).
    WithCabecera(header).
    AddDetalle(line).
    Build()
```

Several fields here are **catalog codes**, not free text: `WithCodigoMetodoPago`, `WithUnidadMedida`, `WithActividadEconomica`, `WithCodigoProductoSin`. SIAT publishes the valid values, and you fetch them with the `Sincronizacion()` service. Using a value that is not in the catalog gets the invoice rejected.

---

## Step 6: Package and send

This is where the SDK saves you the most work. `WithFactura` serializes the invoice to XML, signs it when required, gzips it, base64-encodes it, and computes the SHA-256 hash — all in one call.

```go
builder := models.NewRecepcionFacturaBuilder().
    WithCodigoModalidad(siat.ModalidadComputarizada).
    WithCodigoSucursal(0).
    WithCodigoPuntoVenta(0).
    WithCodigoDocumentoSector(1).
    WithCodigoEmision(siat.EmisionOnline).
    WithTipoFacturaDocumento(1).
    WithCuis(cuis).
    WithCufd(cufd).
    WithFechaEnvio(fechaEmision)

// Attach the invoice. s.Config() acts as the signer.
if err := builder.WithFactura(invoice, s.Config()); err != nil {
    log.Fatal("could not package the invoice:", err)
}

req := builder.Build()

resp, err := s.CompraVenta().RecepcionFactura(ctx, req)
if err != nil {
    log.Fatal(err)
}
if err := siat.Verify(resp.Body.Content.RespuestaServicioFacturacion); err != nil {
    log.Fatal("SIAT rejected the invoice:", err)
}

log.Println("Accepted. Reception code:", resp.Body.Content.RespuestaServicioFacturacion.CodigoRecepcion)
```

Two details worth pausing on:

**`WithFactura` must come after `WithCodigoModalidad`.** It reads the modality you set to decide whether the document needs a digital signature. Calling it first means it never signs.

**`s.Config()` is the signer.** `Config` carries your `CredentialSign` and implements the signing interface, so passing it hands the credentials to the packaging step. In computerized modality no signature happens, but passing it keeps the code identical when you switch to electronic.

---

## Step 7: Switch to electronic (with digital signature)

You now have a working flow. Turning it into a legally-signed electronic invoice takes two changes.

First, add your certificate to the config:

```go
s, err := siat.New(siat.Config{
    Token:          "YOUR_SIAT_TOKEN",
    Nit:            123456789,
    CodigoSistema:  "YOUR_SYSTEM_CODE",
    CodigoAmbiente: siat.AmbientePruebas,
    BaseURL:        pilotURL,

    // .p12 / .pfx file plus its password
    CredentialSign: siat.NewP12Credential("cert.p12", "p12-password"),
})
```

Second, replace every `siat.ModalidadComputarizada` with `siat.ModalidadElectronica`, and call `s.Electronica()` instead of `s.CompraVenta()`.

That is the whole difference. Because `WithFactura` already receives `s.Config()`, it now finds a credential, sees the electronic modality, and signs the XML with XMLDSig before compressing it.

> If your certificate is a separate certificate/key pair rather than a `.p12`, use `siat.NewPEMCredential("cert.crt", "key.pem")` instead.

---

## What you learned

You built the full chain that every Bolivian electronic invoice requires:

```
Config ──► CUIS ──► CUFD ──► CUF ──► Invoice XML ──► Submission
 once      once     daily    per-invoice  per-invoice
```

You also met the three habits that make SIAT integration reliable:

- **Check errors twice** — transport error, then `siat.Verify`.
- **Reuse the timestamp** — the same `fechaEmision` in the CUF and the header.
- **Order matters** — `WithCodigoModalidad` before `WithFactura`.

## Where to go next

| You want to… | Read |
| :----------- | :--- |
| Sign with a certificate, in depth | [How-to: Sign invoices](../how-to/sign-invoices.md) |
| Send hundreds of invoices at once | [How-to: Send batches](../how-to/send-batches.md) |
| Handle rejections and retries properly | [How-to: Handle errors](../how-to/handle-errors.md) |
| Invoice for a sector other than sales | [Explanation: Sectors](../explanation/sectors.md) |
| Look up an exact method signature | [Reference: API](../reference/api.md) |
| Understand how the SDK is put together | [Explanation: Architecture](../explanation/architecture.md) |
