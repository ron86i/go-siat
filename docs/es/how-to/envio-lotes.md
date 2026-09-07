# Cómo enviar facturas por lotes

<p align="right">
  <a href="../../en/how-to/send-batches.md">🇬🇧 English</a> · <a href="../README.md">Índice de documentación</a>
</p>

Enviar facturas de a una funciona, pero cuesta un viaje HTTP por factura. El SIAT ofrece dos modos de lote, y esta guía cubre cuándo usar cada uno y cómo verificar el resultado.

---

## ¿Cuál necesito?

| | **Paquete** (`RecepcionPaquete`) | **Masivo** (`RecepcionMasiva`) |
| :--- | :--- | :--- |
| Pensado para | Lotes offline / contingencia | Facturación online de alto volumen |
| Código de emisión | `siat.EmisionOffline` o `EmisionMasiva` | `siat.EmisionMasiva` |
| Admite CAFC | Sí — `WithCafc` | No |
| Admite código de evento | Sí — `WithCodigoEvento` | No |
| Resultado | Asíncrono, se valida después | Asíncrono, se valida después |

Los dos son **asíncronos**: el SIAT acepta el paquete y te da un código de recepción, después lo procesa. Tenés que consultar por separado para saber si las facturas individuales fueron aceptadas.

> Los límites de tamaño de lote los fija la normativa del SIAT, no el SDK, y varían según modalidad y sector. Revisá tu especificación técnica vigente del SIAT antes de dimensionar lotes.

---

## Enviar un paquete

`WithFacturas` recibe un slice de structs de factura y hace todo el empaquetado: serializa cada una, las firma cuando la modalidad lo exige, arma el `.tar.gz`, lo codifica en base64 y lo hashea.

```go
facturas := []any{factura1, factura2, factura3}

builder := models.NewRecepcionPaqueteFacturaBuilder().
    WithCodigoModalidad(siat.ModalidadElectronica).   // ← antes de WithFacturas
    WithCodigoSucursal(0).
    WithCodigoPuntoVenta(0).
    WithCodigoDocumentoSector(1).
    WithCodigoEmision(siat.EmisionOffline).
    WithTipoFacturaDocumento(1).
    WithCuis(cuis).
    WithCufd(cufd).
    WithCantidadFacturas(len(facturas)).
    WithFechaEnvio(time.Now())

if err := builder.WithFacturas(facturas, s.Config()); err != nil {
    log.Fatal("no se pudo empaquetar el lote:", err)
}

resp, err := s.CompraVenta().RecepcionPaqueteFactura(ctx, builder.Build())
if err != nil {
    log.Fatal(err)
}
if err := siat.Verify(resp.Body.Content.RespuestaServicioFacturacion); err != nil {
    log.Fatal("el SIAT rechazó el paquete:", err)
}

codigoRecepcion := resp.Body.Content.RespuestaServicioFacturacion.CodigoRecepcion
```

**Guardá ese `CodigoRecepcion`.** Es la única manija que tenés para averiguar qué pasó con las facturas de adentro.

`WithCantidadFacturas` tiene que coincidir con la cantidad real de facturas del slice. Si no coincide, se rechaza.

### Paquetes de contingencia

Cuando el lote es resultado de una caída, registrá primero el evento significativo. El código de recepción del evento vincula el lote con esa contingencia:

```go
inicio := time.Now().Add(-15 * time.Minute) // inicio real de la caída
fin := time.Now()                           // recuperación o fin de emisión offline

eventoReq := models.NewRegistroEventoSignificativoBuilder().
    WithCodigoSucursal(0).
    WithCodigoPuntoVenta(0).
    WithCuis(cuis).
    WithCufd(cufd).
    WithCufdEvento(cufdEvento).
    WithCodigoMotivoEvento(motivoEvento).
    WithDescripcion("Corte de conexión").
    WithFechaInicio(inicio).
    WithFechaFin(fin).
    Build()

eventoResp, err := s.Operaciones().RegistroEventosSignificativos(ctx, eventoReq)
if err != nil {
    log.Fatal(err)
}
evento, err := eventoResp.GetContent()
if err != nil {
    log.Fatal(err)
}
if !evento.Respuesta.Transaccion {
    log.Fatalf("el SIAT rechazó el evento de contingencia: %+v", evento.Respuesta.MensajesList)
}
codigoEvento := evento.Respuesta.CodigoRecepcionEventoSignificativo
```

Después agregá el CAFC y ese código al paquete. El lote debe usar `siat.EmisionOffline`:

```go
cafc := "TU-CODIGO-CAFC"
builder.
    WithCafc(&cafc).
    WithCodigoEvento(codigoEvento)   // int64, de RegistroEventosSignificativos
```

`WithCafc` recibe un `*string` para poder omitirse — pasá `nil` o simplemente no lo llames cuando no hay CAFC.

---

## Enviar un lote masivo

Misma forma, otro builder y sin CAFC:

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
    WithCantidadFacturas(len(facturas)).
    WithFechaEnvio(time.Now())

if err := builder.WithFacturas(facturas, s.Config()); err != nil {
    log.Fatal(err)
}

resp, err := s.Electronica().RecepcionMasivaFactura(ctx, builder.Build())
```

---

## Verificar qué pasó con el lote

Que se acepte el paquete no significa que se acepten sus facturas. Consultá con el código de recepción:

```go
req := models.NewValidacionRecepcionPaqueteFacturaBuilder().
    WithCodigoRecepcion(codigoRecepcion).
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
    // No necesariamente fatal: puede seguir procesándose, o
    // pueden haberse rechazado facturas individuales. Revisá los mensajes.
    log.Println("la validación informó:", err)
}
```

Para lotes masivos usá `NewValidacionRecepcionMasivaFacturaBuilder` y `ValidacionRecepcionMasivaFactura` — forma idéntica.

No consultes en un bucle cerrado. El SIAT procesa los lotes de forma asíncrona; esperá entre intentos y tratá "todavía procesando" como un estado normal, no como un error.

---

## Armar el archivo vos mismo

Si necesitás el `.tar.gz` antes de enviarlo — para guardarlo, o para mandarlo después desde un sistema offline — armalo con las utilidades y cargalo manualmente:

```go
// generar el archivo
err := utils.ExportTarGz(facturas, s.Config(), "lote.tar.gz")

// o generar exactamente las cadenas que espera la solicitud
raw, _ := os.ReadFile("lote.tar.gz")
hash, encoded, err := utils.CompressAndHash(raw)

builder.
    WithArchivo(encoded).
    WithHashArchivo(hash)
```

Usá `WithFacturas` **o** el par manual `WithArchivo`/`WithHashArchivo` — no ambos. Gana el último que se llame.

---

## Relacionado

| | |
| :--- | :--- |
| Firmar las facturas de un lote | [Guía: Firmar facturas](firmar-facturas.md) |
| Interpretar códigos de rechazo | [Guía: Manejo de errores](manejo-errores.md) |
| Helpers de compresión y hashing | [Referencia: Utilidades](../reference/utilidades.md) |
