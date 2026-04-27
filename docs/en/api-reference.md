# API Reference

[← Back to Index](README.md)

> Complete reference for all 12 SIAT services exposed by the `go-siat` SDK. Each method includes its signature, builder, and response navigation.

---

## Table of Contents

1. [Client Initialization](#client-initialization)
2. [Codigos Service](#codigos-service)
3. [Sincronizacion Service](#sincronizacion-service)
4. [Operaciones Service](#operaciones-service)
5. [CompraVenta Service](#compraventa-service)
6. [Electronica Service](#electronica-service)
7. [Computarizada Service](#computarizada-service)
8. [DocumentoAjuste Service](#documentoajuste-service)
9. [Specialized Sector Services](#specialized-sector-services)
10. [Response Verification](#response-verification)

---

## Client Initialization

### `siat.New(baseUrl, httpClient)`

Creates the main SDK client.

```go
s, err := siat.New("https://pilotosiatservicios.impuestos.gob.bo/v2", nil)
```

| Parameter | Type | Required | Description |
|:----------|:-----|:---------|:------------|
| `baseUrl` | `string` | ✅ | SIAT service base URL |
| `httpClient` | `*http.Client` | ❌ | Custom HTTP client (nil = optimized defaults) |

**Returns**: `(*SiatServices, error)`

### `siat.NewWithMiddleware(baseUrl, httpClient, middlewares...)`

Creates the SDK client with HTTP middlewares for logging, metrics, retry, etc.

```go
s, err := siat.NewWithMiddleware(baseURL, nil, &LoggingMiddleware{})
```

| Parameter | Type | Required | Description |
|:----------|:-----|:---------|:------------|
| `baseUrl` | `string` | ✅ | SIAT service base URL |
| `httpClient` | `*http.Client` | ❌ | Custom HTTP client |
| `middlewares` | `...HTTPMiddleware` | ❌ | Middleware chain (first = outermost) |

### Service Accessors

```go
s.Codigos()            // → SiatCodigosService
s.Sincronizacion()     // → SiatSincronizacionService
s.Operaciones()        // → SiatOperacionesPort
s.CompraVenta()        // → SiatCompraVentaService
s.Electronica()        // → SiatElectronicaService
s.Computarizada()      // → SiatComputarizadaService
s.DocumentoAjuste()    // → SiatDocumentoAjusteService
s.Telecomunicaciones() // → SiatTelecomunicacionesService
s.ServicioBasico()     // → SiatServicioBasicoService
s.EntidadFinanciera()  // → SiatEntidadFinancieraService
s.BoletoAereo()        // → SiatBoletoAereoService
s.RecepcionCompras()   // → SiatRecepcionComprasService
```

### Helper Methods

| Method | Description |
|:-------|:------------|
| `s.WithTraceID(id)` | Sets a trace ID for distributed tracing (`X-Trace-ID` header). Returns `*SiatServices` for chaining. |
| `s.WithConfig(token)` | Creates a `Config` with the token and current trace ID. |

---

## Codigos Service

Manages SIAT billing codes (CUIS, CUFD), NIT validation, and certificate communication.

**Access**: `s.Codigos()`

### `SolicitudCuis` - Request CUIS

Obtains the System Identification Unique Code, required to operate with SIAT.

```go
req := models.Codigos().NewCuisBuilder().
    WithCodigoAmbiente(2).
    WithCodigoModalidad(1).
    WithCodigoPuntoVenta(0).
    WithCodigoSucursal(0).
    WithCodigoSistema("ABC123").
    WithNit(123456789).
    Build()

resp, err := s.Codigos().SolicitudCuis(ctx, cfg, req)
// Navigate: resp.Body.Content.RespuestaCuis.Codigo
```

| Builder Method | Type | Description |
|:---------------|:-----|:------------|
| `WithCodigoAmbiente` | `int` | 1=Production, 2=Testing |
| `WithCodigoModalidad` | `int` | 1=Electronic, 2=Computerized |
| `WithCodigoPuntoVenta` | `int` | Point of sale (0=main) |
| `WithCodigoSucursal` | `int` | Branch number (0=main) |
| `WithCodigoSistema` | `string` | SIAT system code |
| `WithNit` | `int64` | Taxpayer NIT |

### `SolicitudCuisMasivo` - Bulk CUIS Request

Request multiple CUIS codes in a single operation.

```go
req := models.Codigos().NewCuisMasivoBuilder().
    // ... same builders as CUIS
    Build()

resp, err := s.Codigos().SolicitudCuisMasivo(ctx, cfg, req)
```

### `SolicitudCufd` - Request CUFD

Obtains the Daily Invoicing Unique Code. Required daily for invoice emission.

```go
req := models.Codigos().NewCufdBuilder().
    WithCodigoAmbiente(2).
    WithCodigoModalidad(1).
    WithCodigoPuntoVenta(0).
    WithCodigoSucursal(0).
    WithCodigoSistema("ABC123").
    WithNit(123456789).
    WithCuis("C2FC682C").
    Build()

resp, err := s.Codigos().SolicitudCufd(ctx, cfg, req)
// Navigate: resp.Body.Content.RespuestaCufd.Codigo
// Navigate: resp.Body.Content.RespuestaCufd.CodigoControl
// Navigate: resp.Body.Content.RespuestaCufd.FechaVigCufd
```

### `SolicitudCufdMasivo` - Bulk CUFD Request

```go
req := models.Codigos().NewCufdMasivoBuilder().
    // ... same builders + WithCuis()
    Build()

resp, err := s.Codigos().SolicitudCufdMasivo(ctx, cfg, req)
```

### `VerificarNit` - Validate NIT

Checks if a tax identification number is active and valid.

```go
req := models.Codigos().NewVerificarNitBuilder().
    WithNit(123456789).
    Build()

resp, err := s.Codigos().VerificarNit(ctx, cfg, req)
// Navigate: resp.Body.Content.RespuestaVerificarNit.Transaccion
```

### `VerificarComunicacion` - Connectivity Test

Tests connectivity to the Codigos service.

```go
req := models.Codigos().NewVerificarComunicacionCodigosBuilder().Build()
resp, err := s.Codigos().VerificarComunicacion(ctx, cfg, req)
```

### `NotificaCertificadoRevocado` - Certificate Revocation

Notifies SIAT that a digital certificate has been revoked.

```go
req := models.Codigos().NewNotificaCertificadoRevocadoBuilder().
    // ... builder methods
    Build()

resp, err := s.Codigos().NotificaCertificadoRevocado(ctx, cfg, req)
```

**Integration tests**: [`siat_codigos_service_test.go`](../../internal/adapter/services/siat_codigos_service_test.go)

---

## Sincronizacion Service

Synchronizes master catalogs: economic activities, parametric tables, products, sector documents, and more.

**Access**: `s.Sincronizacion()`

### Catalog Synchronization Methods

All synchronization methods share a similar builder pattern:

```go
req := models.Sincronizacion().NewSincronizarActividadesBuilder().
    WithNit(nit).
    WithCodigoAmbiente(2).
    WithCodigoSucursal(0).
    WithCodigoPuntoVenta(0).
    WithCodigoSistema("ABC123").
    WithCuis("C2FC682C").
    Build()

resp, err := s.Sincronizacion().SincronizarActividades(ctx, cfg, req)
```

| Method | Description |
|:-------|:------------|
| `SincronizarActividades` | Economic activities catalog |
| `SincronizarListaActividadesDocumentoSector` | Activity-to-sector document mapping |
| `SincronizarListaLeyendasFactura` | Official invoice legends |
| `SincronizarListaMensajesServicios` | Service messages catalog |
| `SincronizarListaProductosServicios` | Products and services catalog |
| `SincronizarParametricaEventosSignificativos` | Significant event types |
| `SincronizarParametricaMotivoAnulacion` | Invoice annulment reasons |
| `SincronizarParametricaPaisOrigen` | Countries of origin |
| `SincronizarParametricaTipoDocumentoIdentidad` | Identity document types |
| `SincronizarParametricaTipoDocumentoSector` | Sector document types |
| `SincronizarParametricaTipoEmision` | Emission modes |
| `SincronizarParametricaTipoHabitacion` | Hotel room types |
| `SincronizarParametricaTipoMetodoPago` | Payment methods |
| `SincronizarParametricaTipoMoneda` | Currency types |
| `SincronizarParametricaTipoPuntoVenta` | Point of sale types |
| `SincronizarParametricaTiposFactura` | Invoice classification types |
| `SincronizarParametricaUnidadMedida` | Units of measure |
| `VerificarComunicacion` | Connectivity test |

**Integration tests**: [`siat_sincronizacion_service_test.go`](../../internal/adapter/services/siat_sincronizacion_service_test.go)

---

## Operaciones Service

Manages point of sale (POS) registration, significant events, and system closings.

**Access**: `s.Operaciones()`

### `RegistroPuntoVenta` - Register POS

```go
req := models.Operaciones().NewRegistroPuntoVentaBuilder().
    // ... builder methods
    Build()

resp, err := s.Operaciones().RegistroPuntoVenta(ctx, cfg, req)
```

### `ConsultaPuntoVenta` - Query POS

```go
req := models.Operaciones().NewConsultaPuntoVentaBuilder().Build()
resp, err := s.Operaciones().ConsultaPuntoVenta(ctx, cfg, req)
```

### Full Method Reference

| Method | Description |
|:-------|:------------|
| `RegistroPuntoVenta` | Register a new POS |
| `ConsultaPuntoVenta` | Query registered POS details |
| `CierrePuntoVenta` | Close/disable a POS |
| `RegistroPuntoVentaComisionista` | Register a third-party commission POS |
| `RegistroEventosSignificativos` | Report contingency events that prevent online invoicing |
| `ConsultaEventosSignificativos` | Query reported contingency events |
| `CierreOperacionesSistema` | Close system operations for a period |
| `VerificarComunicacion` | Connectivity test |

**Integration tests**: [`siat_operaciones_service_test.go`](../../internal/adapter/services/siat_operaciones_service_test.go)

---

## CompraVenta Service

Handles standard sales invoicing - the most common sector for general commerce.

**Access**: `s.CompraVenta()`

### `RecepcionFactura` - Send Invoice

```go
req := models.CompraVenta().NewRecepcionFacturaBuilder().
    WithCodigoAmbiente(2).
    WithNit(nit).
    WithCufd(cufd).
    WithCuis(cuis).
    WithTipoFacturaDocumento(1).
    WithArchivo(archivoBase64).
    WithFechaEnvio(time.Now()).
    WithHashArchivo(hash).
    Build()

resp, err := s.CompraVenta().RecepcionFactura(ctx, cfg, req)
```

### Full Method Reference

| Method | Description |
|:-------|:------------|
| `RecepcionFactura` | Send a single invoice |
| `AnulacionFactura` | Annul a previously accepted invoice |
| `ReversionAnulacionFactura` | Reverse an annulment (one-time only) |
| `RecepcionPaqueteFactura` | Send batch of up to 500 invoices |
| `ValidacionRecepcionPaqueteFactura` | Validate batch reception status |
| `RecepcionMasivaFactura` | Massive send (501–1000 invoices) |
| `ValidacionRecepcionMasivaFactura` | Validate massive reception status |
| `VerificacionEstadoFactura` | Check the status of a specific invoice |
| `RecepcionAnexos` | Send annexes (attachments) |
| `VerificarComunicacion` | Connectivity test |

**Integration tests**: [`siat_compra_venta_service_test.go`](../../internal/adapter/services/siat_compra_venta_service_test.go)

---

## Electronica Service

Handles electronic invoicing (with digital signature) for all sectors.

**Access**: `s.Electronica()`

### Full Method Reference

| Method | Description |
|:-------|:------------|
| `RecepcionFactura` | Send a digitally signed invoice |
| `AnulacionFactura` | Annul an electronic invoice |
| `ReversionAnulacionFactura` | Reverse an annulment |
| `RecepcionPaqueteFactura` | Send batch (up to 500) |
| `ValidacionRecepcionPaqueteFactura` | Validate batch status |
| `RecepcionMasivaFactura` | Massive send (501–2000) |
| `ValidacionRecepcionMasivaFactura` | Validate massive status |
| `VerificacionEstadoFactura` | Check invoice status |
| `RecepcionAnexosSuministroEnergia` | Energy supply annexes (recharges, gift cards) |
| `VerificarComunicacion` | Connectivity test |

All methods use builders from `models.Electronica()`:

```go
req := models.Electronica().NewRecepcionFacturaBuilder().
    WithCodigoAmbiente(2).
    WithNit(nit).
    WithCufd(cufd).
    WithCuis(cuis).
    WithTipoFacturaDocumento(1).
    WithArchivo(archivoBase64).
    WithFechaEnvio(time.Now()).
    WithHashArchivo(hash).
    Build()

resp, err := s.Electronica().RecepcionFactura(ctx, cfg, req)
```

**Integration tests**: [`siat_electronica_service_test.go`](../../internal/adapter/services/siat_electronica_service_test.go)

---

## Computarizada Service

Handles computerized invoicing without digital signature, based on fiscal cash registers.

**Access**: `s.Computarizada()`

### Full Method Reference

| Method | Description |
|:-------|:------------|
| `RecepcionFactura` | Send a computerized invoice |
| `AnulacionFactura` | Annul a computerized invoice |
| `ReversionAnulacionFactura` | Reverse an annulment |
| `RecepcionPaqueteFactura` | Send batch (up to 500) |
| `ValidacionRecepcionPaqueteFactura` | Validate batch status |
| `RecepcionMasivaFactura` | Massive send (501–2000) |
| `ValidacionRecepcionMasivaFactura` | Validate massive status |
| `VerificacionEstadoFactura` | Check invoice status |
| `RecepcionAnexosSuministroEnergia` | Energy supply annexes |
| `VerificarComunicacion` | Connectivity test |

Builders from `models.Computarizada()`:

```go
req := models.Computarizada().NewRecepcionFacturaBuilder().
    // ... same pattern as Electronica
    Build()

resp, err := s.Computarizada().RecepcionFactura(ctx, cfg, req)
```

**Integration tests**: [`siat_computarizada_service_test.go`](../../internal/adapter/services/siat_computarizada_service_test.go)

---

## DocumentoAjuste Service

Manages adjustment documents: credit/debit notes, conciliation notes, and reversals.

**Access**: `s.DocumentoAjuste()`

### Full Method Reference

| Method | Description |
|:-------|:------------|
| `RecepcionDocumentoAjuste` | Send an adjustment document (credit/debit note) |
| `AnulacionDocumentoAjuste` | Annul a previously issued adjustment document |
| `ReversionAnulacionDocumentoAjuste` | Reverse an adjustment document annulment |
| `VerificacionEstadoDocumentoAjuste` | Check adjustment document status |
| `VerificarComunicacion` | Connectivity test |

```go
req := models.DocumentoAjuste().NewRecepcionDocumentoAjusteBuilder().
    // ... builder methods
    Build()

resp, err := s.DocumentoAjuste().RecepcionDocumentoAjuste(ctx, cfg, req)
```

**Integration tests**: [`siat_documento_ajuste_service_test.go`](../../internal/adapter/services/siat_documento_ajuste_service_test.go)

---

## Specialized Sector Services

Several SIAT sectors require dedicated SOAP endpoints instead of using the generic `Electronica` or `Computarizada` services. `go-siat` provides native clients for all of them.

**Access**: 
- `s.Telecomunicaciones()`
- `s.ServicioBasico()`
- `s.EntidadFinanciera()`
- `s.BoletoAereo()`
- `s.RecepcionCompras()`

### Common Operations

The specialized services follow the same method patterns as the general ones:

| Method | Description |
|:-------|:------------|
| `RecepcionFactura` | Send a single invoice to the specialized endpoint |
| `AnulacionFactura` | Annul an invoice |
| `ReversionAnulacionFactura` | Reverse an annulment |
| `RecepcionPaqueteFactura` | Send batch (up to 500) |
| `RecepcionMasivaFactura` | Massive send |
| `VerificacionEstadoFactura` | Check invoice status |
| `VerificarComunicacion` | Connectivity test |

> [!NOTE]
> Not all services support all operations. For example, `BoletoAereo` only supports `RecepcionMasivaFactura` for mass sending, as required by SIAT for airlines.

---

## Response Verification

### `siat.Verify(response)`

Analyzes a SIAT response and returns a `*SiatError` if the operation was rejected:

```go
resp, err := s.Codigos().SolicitudCuis(ctx, cfg, req)
if err != nil {
    return err // Network error
}

// Verify the SIAT business response
if err := siat.Verify(resp.Body.Content.RespuestaCuis); err != nil {
    var siatErr *siat.SiatError
    if errors.As(err, &siatErr) {
        fmt.Printf("SIAT code: %d\n", siatErr.SiatCode)
        fmt.Printf("Retryable: %v\n", siatErr.IsRetryable)
    }
    return err
}
```

`Verify` works with **any** SDK response type. It uses the `Result` interface first, then falls back to reflection for compatibility.

For full error handling details, see the [Error Handling](error-handling.md) guide.

---

[← Back to Index](README.md) | [Next: Invoicing Guide →](invoicing-guide.md)
