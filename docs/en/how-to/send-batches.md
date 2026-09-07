# How to send invoices in batches

<p align="right">
  <a href="../../es/how-to/envio-lotes.md">🇪🇸 Español</a> · <a href="../README.md">Docs index</a>
</p>

Sending invoices one at a time works, but it costs one HTTP round trip per invoice. SIAT offers two batching modes, and this guide covers when to use each and how to check the result.

---

## Which mode do I need?

| | **Package** (`RecepcionPaquete`) | **Massive** (`RecepcionMasiva`) |
| :--- | :--- | :--- |
| Built for | Offline / contingency batches | High-volume online invoicing |
| Emission code | `siat.EmisionOffline` or `EmisionMasiva` | `siat.EmisionMasiva` |
| Supports a CAFC | Yes — `WithCafc` | No |
| Supports an event code | Yes — `WithCodigoEvento` | No |
| Result | Asynchronous, validated later | Asynchronous, validated later |

Both are **asynchronous**: SIAT accepts the package and gives you a reception code, then processes it. You must poll separately to learn whether the individual invoices were accepted.

> Batch size limits are set by SIAT regulation, not by the SDK, and they differ per modality and sector. Check your current SIAT technical specification before sizing batches.

---

## Send a package

`WithFacturas` takes a slice of invoice structs and does the whole packaging pipeline: serialize each one, sign them when the modality requires it, build the `.tar.gz`, base64-encode it and hash it.

```go
invoices := []any{invoice1, invoice2, invoice3}

builder := models.NewRecepcionPaqueteFacturaBuilder().
    WithCodigoModalidad(siat.ModalidadElectronica).   // ← before WithFacturas
    WithCodigoSucursal(0).
    WithCodigoPuntoVenta(0).
    WithCodigoDocumentoSector(1).
    WithCodigoEmision(siat.EmisionOffline).
    WithTipoFacturaDocumento(1).
    WithCuis(cuis).
    WithCufd(cufd).
    WithCantidadFacturas(len(invoices)).
    WithFechaEnvio(time.Now())

if err := builder.WithFacturas(invoices, s.Config()); err != nil {
    log.Fatal("could not package the batch:", err)
}

resp, err := s.CompraVenta().RecepcionPaqueteFactura(ctx, builder.Build())
if err != nil {
    log.Fatal(err)
}
if err := siat.Verify(resp.Body.Content.RespuestaServicioFacturacion); err != nil {
    log.Fatal("SIAT rejected the package:", err)
}

receptionCode := resp.Body.Content.RespuestaServicioFacturacion.CodigoRecepcion
```

**Keep that `CodigoRecepcion`.** It is the only handle you have to find out what happened to the invoices inside.

`WithCantidadFacturas` must match the real number of invoices in the slice. A mismatch is rejected.

### Contingency packages

When the batch is the result of an outage, first register the significant event. Its reception code links the batch to that contingency:

```go
start := time.Now().Add(-15 * time.Minute) // actual outage start
end := time.Now()                          // recovery or end of offline issuance

eventReq := models.NewRegistroEventoSignificativoBuilder().
    WithCodigoSucursal(0).
    WithCodigoPuntoVenta(0).
    WithCuis(cuis).
    WithCufd(cufd).
    WithCufdEvento(eventCufd).
    WithCodigoMotivoEvento(eventReason).
    WithDescripcion("Connection outage").
    WithFechaInicio(start).
    WithFechaFin(end).
    Build()

eventResp, err := s.Operaciones().RegistroEventosSignificativos(ctx, eventReq)
if err != nil {
    log.Fatal(err)
}
event, err := eventResp.GetContent()
if err != nil {
    log.Fatal(err)
}
if !event.Respuesta.Transaccion {
    log.Fatalf("SIAT rejected the contingency event: %+v", event.Respuesta.MensajesList)
}
eventCode := event.Respuesta.CodigoRecepcionEventoSignificativo
```

Then add the CAFC and that event code to the package. The batch must use `siat.EmisionOffline`:

```go
cafc := "YOUR-CAFC-CODE"
builder.
    WithCafc(&cafc).
    WithCodigoEvento(eventCode)   // int64, from RegistroEventosSignificativos
```

`WithCafc` takes a `*string` so it can be omitted — pass `nil` or simply don't call it when there is no CAFC.

---

## Send a massive batch

Same shape, different builder and no CAFC:

```go
builder := models.NewRecepcionMasivaFacturaBuilder().
    WithCodigoModalidad(siat.ModalidadElectronica).
    WithCodigoSucursal(0).
    WithCodigoPuntoVenta(0).
    WithCodigoDocumentoSector(1).
    WithCodigoEmision(siat.EmisionMasiva).
    WithTipoFacturaDocumento(1).
    WithCuis(cuis).
    WithCufd(cufd).
    WithCantidadFacturas(len(invoices)).
    WithFechaEnvio(time.Now())

if err := builder.WithFacturas(invoices, s.Config()); err != nil {
    log.Fatal(err)
}

resp, err := s.Electronica().RecepcionMasivaFactura(ctx, builder.Build())
```

---

## Check what happened to the batch

Acceptance of the package is not acceptance of its invoices. Poll with the reception code:

```go
req := models.NewValidacionRecepcionPaqueteFacturaBuilder().
    WithCodigoRecepcion(receptionCode).
    WithCodigoModalidad(siat.ModalidadElectronica).
    WithCodigoSucursal(0).
    WithCodigoPuntoVenta(0).
    WithCodigoDocumentoSector(1).
    WithCodigoEmision(siat.EmisionOffline).
    WithTipoFacturaDocumento(1).
    WithCuis(cuis).
    WithCufd(cufd).
    Build()

resp, err := s.CompraVenta().ValidacionRecepcionPaqueteFactura(ctx, req)
if err != nil {
    log.Fatal(err)
}
if err := siat.Verify(resp.Body.Content.RespuestaServicioFacturacion); err != nil {
    // Not necessarily fatal: it may still be processing, or
    // individual invoices may have been rejected. Inspect the messages.
    log.Println("validation reported:", err)
}
```

For massive batches use `NewValidacionRecepcionMasivaFacturaBuilder` and `ValidacionRecepcionMasivaFactura` — identical shape.

Do not poll in a tight loop. SIAT processes batches asynchronously; wait between attempts and treat "still processing" as a normal state rather than an error.

---

## Building the archive yourself

If you need the `.tar.gz` before sending — to store it, or to send it later from an offline system — build it with the utilities and feed it in manually:

```go
// produce the file
err := utils.ExportTarGz(invoices, s.Config(), "batch.tar.gz")

// or produce the exact strings the request expects
raw, _ := os.ReadFile("batch.tar.gz")
hash, encoded, err := utils.CompressAndHash(raw)

builder.
    WithArchivo(encoded).
    WithHashArchivo(hash)
```

Use either `WithFacturas` **or** the manual `WithArchivo`/`WithHashArchivo` pair — not both. The last one called wins.

---

## Related

| | |
| :--- | :--- |
| Signing the invoices in a batch | [How-to: Sign invoices](sign-invoices.md) |
| Interpreting rejection codes | [How-to: Handle errors](handle-errors.md) |
| Compression and hashing helpers | [Reference: Utilities](../reference/utilities.md) |
