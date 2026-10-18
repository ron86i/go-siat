# API reference

<p align="right">
  <a href="../../es/reference/api.md">🇪🇸 Español</a> · <a href="../README.md">Docs index</a>
</p>

Complete public surface of `github.com/ron86i/go-siat/v3`. For generated godoc see [pkg.go.dev](https://pkg.go.dev/github.com/ron86i/go-siat/v3).

```go
import (
    siat "github.com/ron86i/go-siat/v3"
    "github.com/ron86i/go-siat/v3/pkg/models"
    "github.com/ron86i/go-siat/v3/pkg/models/invoices"
    "github.com/ron86i/go-siat/v3/pkg/utils"
)
```

The `/v2` suffix is required — it is a semantic import version, not an optional path.

---

## Contents

- [Client construction](#client-construction)
- [Service facades](#service-facades)
- [Service interfaces](#service-interfaces)
- [Request builders](#request-builders)
- [Response handling](#response-handling)
- [Errors](#errors)
- [Constants](#constants)
- [Helper types](#helper-types)

---

## Client construction

```go
func New(config Config) (*SiatServices, error)
func NewWithMiddleware(config Config, middlewares ...HTTPMiddleware) (*SiatServices, error)
```

`New` validates the config and returns the client. It fails when `Token`, `CodigoSistema` or `BaseURL` are empty, when `Nit` is not positive, or when `CodigoAmbiente` is neither `AmbienteProduccion` nor `AmbientePruebas`.

```go
s, err := siat.New(siat.Config{
    Token:          "YOUR_SIAT_TOKEN",
    Nit:            123456789,
    CodigoSistema:  "YOUR_SYSTEM_CODE",
    CodigoAmbiente: siat.AmbientePruebas,
    BaseURL:        "https://pilotosiatservicios.impuestos.gob.bo/sfd",
})
```

### Config

`siat.Config` is an alias for `ports.Config`.

| Field | Type | Required | Purpose |
| :--- | :--- | :--- | :--- |
| `Token` | `string` | yes | SIAT API token, sent as `apiKey: TokenApi <token>` |
| `Nit` | `int64` | yes | Taxpayer NIT, injected into every request |
| `CodigoSistema` | `string` | yes | System code SIAT assigned to your software |
| `CodigoAmbiente` | `int` | yes | `AmbienteProduccion` or `AmbientePruebas` |
| `BaseURL` | `string` | yes | Environment root; service paths are appended |
| `UserAgent` | `string` | no | Defaults to `"go-siat"` |
| `TraceId` | `string` | no | Sent as `X-Trace-ID` when non-empty |
| `HTTPClient` | `*http.Client` | no | Defaults to `NewHTTPClient(DefaultHTTPConfig())` |
| `CredentialSign` | `CredentialSign` | no | Required only for electronic modality |

`Nit`, `CodigoSistema` and `CodigoAmbiente` are injected into outgoing requests only where the corresponding field is still zero — an explicit value on a builder always wins.

```go
func (c Config) SignXML(xmlBytes []byte) ([]byte, error)
```

`Config` itself satisfies the `XMLSigner` interface, which is why `s.Config()` is what you pass to `WithFactura`, `WithFacturas`, `WithDocumento` and the `utils.Export*` functions.

### Signing credentials

```go
func NewP12Credential(p12 any, password string) CredentialSign
func NewPEMCredential(cert, privateKey any) CredentialSign
func (cf CredentialSign) GetType() string   // "P12" | "PEM" | "UNKNOWN"
```

Both constructors accept a file path (`string`) or raw bytes (`[]byte`). They do not panic on a missing file — the error surfaces at signing time, and `GetType()` returns `"UNKNOWN"`.

### Per-request config override

```go
func WithDynamicConfig(ctx context.Context, config Config) context.Context
func GetConfigFromContext(ctx context.Context) (Config, bool)
```

Overrides `Token`, `Nit`, `CodigoSistema`, `CodigoAmbiente`, `TraceId` and `CredentialSign` for a single call — non-zero fields win. Built for multi-tenant services sharing one client.

```go
ctx = ports.WithDynamicConfig(ctx, siat.Config{Nit: tenantNit, Token: tenantToken})
resp, err := s.Electronica().RecepcionFactura(ctx, req)
```

### HTTP client

```go
type HTTPConfig = services.HTTPConfig
func DefaultHTTPConfig() HTTPConfig
func NewHTTPClient(cfg HTTPConfig) *http.Client
type HTTPMiddleware = middleware.HTTPMiddleware
```

See [Configuration](configuration.md) for the fields and tuning guidance.

---

## Service facades

Twelve accessors on `*SiatServices`, each mapping to one SOAP endpoint.

| Method | Returns | Endpoint |
| :--- | :--- | :--- |
| `Codigos()` | `SiatCodigosService` | `FacturacionCodigos` |
| `Operaciones()` | `SiatOperacionesService` | `FacturacionOperaciones` |
| `Sincronizacion()` | `SiatSincronizacionService` | `FacturacionSincronizacion` |
| `CompraVenta()` | `SiatCompraVentaService` | `ServicioFacturacionCompraVenta` |
| `Electronica()` | `SiatSuministroEnergiaService` | `ServicioFacturacionElectronica` |
| `Computarizada()` | `SiatSuministroEnergiaService` | `ServicioFacturacionComputarizada` |
| `DocumentoAjuste()` | `SiatDocumentoAjusteService` | `ServicioFacturacionDocumentoAjuste` |
| `Telecomunicaciones()` | `FacturacionService` | `ServicioFacturacionTelecomunicaciones` |
| `ServicioBasico()` | `FacturacionService` | `ServicioFacturacionServicioBasico` |
| `EntidadFinanciera()` | `FacturacionService` | `ServicioFacturacionEntidadFinanciera` |
| `BoletoAereo()` | `SiatBoletoAereoService` | `ServicioFacturacionBoletoAereo` |
| `RecepcionCompras()` | `SiatRecepcionComprasService` | `ServicioRecepcionCompras` |

Plus `s.Config() Config`, returning the active configuration.

Which facade a given sector requires is fixed — see [Sectors](../explanation/sectors.md#how-a-sector-picks-its-service-endpoint).

---

## Service interfaces

Every method has the same shape:

```go
Method(ctx context.Context, req models.<RequestType>) (*soap.EnvelopeResponse[<ResponseType>], error)
```

### FacturacionService

The nine-method base, implemented by every invoicing facade.

| Method | Request builder |
| :--- | :--- |
| `RecepcionFactura` | `NewRecepcionFacturaBuilder()` |
| `RecepcionPaqueteFactura` | `NewRecepcionPaqueteFacturaBuilder()` |
| `RecepcionMasivaFactura` | `NewRecepcionMasivaFacturaBuilder()` |
| `ValidacionRecepcionPaqueteFactura` | `NewValidacionRecepcionPaqueteFacturaBuilder()` |
| `ValidacionRecepcionMasivaFactura` | `NewValidacionRecepcionMasivaFacturaBuilder()` |
| `AnulacionFactura` | `NewAnulacionFacturaBuilder()` |
| `ReversionAnulacionFactura` | `NewReversionAnulacionFacturaBuilder()` |
| `VerificacionEstadoFactura` | `NewVerificacionEstadoFacturaBuilder()` |
| `VerificarComunicacion` | `NewVerificarComunicacionCodigosBuilder()` |

Two interfaces extend it by embedding:

```go
type SiatCompraVentaService interface {
    FacturacionService
    RecepcionAnexos(ctx, req models.RecepcionAnexosCompraVenta) (...)
}

type SiatSuministroEnergiaService interface {
    FacturacionService
    RecepcionAnexosSuministroEnergia(ctx, req models.RecepcionAnexosSuministroEnergia) (...)
}
```

`Electronica()` and `Computarizada()` both return `SiatSuministroEnergiaService`, so switching modality never changes the call site beyond the facade name.

### SiatCodigosService

| Method | Builder | Purpose |
| :--- | :--- | :--- |
| `SolicitudCuis` | `NewCuisBuilder()` | Request the system code |
| `SolicitudCuisMasivo` | `NewCuisMasivoBuilder()` | Bulk CUIS |
| `SolicitudCufd` | `NewCufdBuilder()` | Request the daily code |
| `SolicitudCufdMasivo` | `NewCufdMasivoBuilder()` | Bulk CUFD |
| `VerificarNit` | `NewVerificarNitBuilder()` | Check a customer NIT |
| `NotificaCertificadoRevocado` | `NewNotificaCertificadoRevocadoBuilder()` | Report a revoked certificate |
| `VerificarComunicacion` | `NewVerificarComunicacionCodigosBuilder()` | Health check |

`NewVerificarNitBuilder` takes the NIT to check via `WithNitParaVerificacion`, not `WithNit` — `WithNit` would target your own NIT.

### SiatOperacionesService

`RegistroPuntoVenta`, `ConsultaPuntoVenta`, `CierrePuntoVenta`, `RegistroPuntoVentaComisionista`, `RegistroEventosSignificativos`, `ConsultaEventosSignificativos`, `CierreOperacionesSistema`, `VerificarComunicacion`.

Significant events are how you register an outage before sending a contingency batch.

### SiatSincronizacionService

Read-only catalogs. `SincronizarFechaHora`, `SincronizarActividades`, `SincronizarListaActividadesDocumentoSector`, `SincronizarListaProductosServicios`, `SincronizarListaLeyendasFactura`, `SincronizarListaMensajesServicios`, `VerificarComunicacion`, plus eleven `SincronizarParametrica*` methods:

`EventosSignificativos`, `MotivoAnulacion`, `PaisOrigen`, `TipoDocumentoIdentidad`, `TipoDocumentoSector`, `TipoEmision`, `TipoHabitacion`, `TipoMetodoPago`, `TipoMoneda`, `TipoPuntoVenta`, `TiposFactura`, `UnidadMedida`.

Fetch these rather than hardcoding codes — SIAT changes them.

### SiatDocumentoAjusteService

`RecepcionDocumentoAjuste`, `AnulacionDocumentoAjuste`, `ReversionAnulacionDocumentoAjuste`, `VerificacionEstadoDocumentoAjuste`, `VerificarComunicacion`.

### SiatBoletoAereoService

`RecepcionMasivaFactura`, `ValidacionRecepcionMasivaFactura`, `AnulacionFactura`, `ReversionAnulacionFactura`, `VerificacionEstadoFactura`, `VerificarComunicacion`. Note there is no single-invoice `RecepcionFactura` — air tickets go out in bulk.

### SiatRecepcionComprasService

`RecepcionPaqueteCompras`, `ValidacionRecepcionPaqueteCompras`, `ConsultaCompras`, `ConfirmacionCompras`, `AnulacionCompra`, `VerificarComunicacion`.

---

## Request builders

All in `pkg/models`, all flat — `models.NewXBuilder()`, never `models.Codigos().NewXBuilder()`. Every builder ends with `.Build()`, which returns an opaque request type.

### Common setters

Most invoicing builders accept these:

| Setter | Type |
| :--- | :--- |
| `WithCuis` / `WithCufd` / `WithCuf` | `string` |
| `WithCodigoModalidad` | `int` |
| `WithCodigoEmision` | `int` |
| `WithCodigoSucursal` / `WithCodigoPuntoVenta` | `int` |
| `WithCodigoDocumentoSector` | `int` |
| `WithTipoFacturaDocumento` | `int` |
| `WithFechaEnvio` | `time.Time` |
| `WithNit` / `WithCodigoSistema` / `WithCodigoAmbiente` | override only — injected from `Config` |

### Document-packaging setters

These break the fluent chain: they perform fallible work (marshal, sign, compress, hash) and return `error` rather than the builder.

```go
func (b *RecepcionFacturaBuilder) WithFactura(factura any, signer XMLSigner) error
func (b *RecepcionPaqueteFacturaBuilder) WithFacturas(facturas []any, signer XMLSigner) error
func (b *RecepcionMasivaFacturaBuilder) WithFacturas(facturas []any, signer XMLSigner) error
func (b *recepcionDocumentoAjusteBuilder) WithDocumento(documento any, signer XMLSigner) error
```

Call `WithCodigoModalidad` **before** these. They read the modality to decide whether to sign; a zero modality produces an unsigned document with no local error.

The manual alternative is `WithArchivo(base64 string)` plus `WithHashArchivo(hash string)`. Use one approach or the other, never both.

### Batch-only setters

| Setter | Type | Builder |
| :--- | :--- | :--- |
| `WithCantidadFacturas` | `int` | package, massive |
| `WithCodigoEvento` | `int64` | package only |
| `WithCodigoRecepcion` | `string` | validation builders |

### The complete builder list

`AnulacionCompra`, `AnulacionDocumentoAjuste`, `AnulacionFactura`, `CierreOperacionesSistema`, `CierrePuntoVenta`, `ConfirmacionCompras`, `ConsultaCompras`, `ConsultaEventoSignificativo`, `ConsultaPuntoVenta`, `Cufd`, `CufdMasivo`, `Cuis`, `CuisMasivo`, `NotificaCertificadoRevocado`, `RecepcionAnexos`, `RecepcionAnexosSuministroEnergia`, `RecepcionDocumentoAjuste`, `RecepcionFactura`, `RecepcionMasivaFactura`, `RecepcionPaqueteCompras`, `RecepcionPaqueteFactura`, `RegistroEventoSignificativo`, `RegistroPuntoVenta`, `RegistroPuntoVentaComisionista`, `ReversionAnulacionDocumentoAjuste`, `ReversionAnulacionFactura`, `SolicitudListaCufdDto`, `SolicitudListaCuisDto`, `SuministroEnergiaAnexo`, `ValidacionRecepcionMasivaFactura`, `ValidacionRecepcionPaqueteCompras`, `ValidacionRecepcionPaqueteFactura`, `VentaAnexo`, `VerificacionEstadoDocumentoAjuste`, `VerificacionEstadoFactura`, `VerificarComunicacionCodigos`, `VerificarComunicacionDocumentoAjuste`, `VerificarComunicacionOperaciones`, `VerificarComunicacionRecepcionCompras`, `VerificarComunicacionSincronizacion`, `VerificarNit`, and the seventeen `SincronizarParametrica*` / `SincronizarLista*` builders.

Each is `models.New<Name>Builder()`.

### Invoice builders

Sector document builders live in `pkg/models/invoices`, three per sector:

```go
invoices.New<Sector>CabeceraBuilder()   // header
invoices.New<Sector>DetalleBuilder()    // one line item
invoices.New<Sector>Builder()           // WithModalidad, WithCabecera, AddDetalle, Build
```

All 50 with builders are listed in [Sectors](../explanation/sectors.md#the-full-sector-catalog).

---

## Response handling

Every method returns `*soap.EnvelopeResponse[T]`:

```go
resp.Body.Fault      // *Fault — non-nil means the envelope itself was refused
resp.Body.Content    // T — the typed payload
```

The payload field name varies by operation:

| Response | Field |
| :--- | :--- |
| `RecepcionFacturaResponse` and most facturacion types | `RespuestaServicioFacturacion` |
| `RecepcionDocumentoAjusteResponse` | `RespuestaRecepcionFactura` |
| `RecepcionCompras*Response` | `RespuestaRecepcion` |
| `CuisResponse` | `RespuestaCuis` |
| `CufdResponse` | `RespuestaCufd` |
| `VerificarNitResponse` | `RespuestaVerificarNit` |
| `VerificarComunicacionResponse` | `RespuestaComunicacion` |

`RespuestaRecepcion` — the shape behind most invoicing responses:

```go
Transaccion       bool
CodigoEstado      int
CodigoDescripcion string
CodigoRecepcion   string
MensajesList      []MensajeServicio
```

### Verify

```go
func Verify(resp interface{}) error
```

Returns an error when `Transaccion` is false, or when any message is not a warning code. Warnings alone pass. Always call it — a `nil` transport error does not mean acceptance.

---

## Errors

```go
type SiatError = errors.SiatError

type SiatError struct {
    Code           string   // "NETWORK_ERROR" | "SIAT_SERVER_ERROR" | "AUTH_FAILED" | "TIMEOUT"
    Message        string
    SiatCode       int
    StatusCode     int
    IsNetworkError bool
    IsRetryable    bool
    Details        map[string]interface{}
    Mensajes       []MensajeServicio
    WrappedErr     error
}

func (e *SiatError) Error() string
func (e *SiatError) Unwrap() error
func (e *SiatError) HasCode(code int) bool
func (e *SiatError) GetWarnings() []MensajeServicio
```

Constructors: `NewNetworkError(msg, err)`, `NewSiatError(code, msg)`, `NewAuthError(msg)`, `NewTimeoutError(msg)`.

Classifiers:

```go
func IsRetryable(err error) bool
func IsNetworkError(err error) bool
func IsRetryableCode(code int) bool    // 123, 967, 991, 995, 999
func IsValidationCode(code int) bool   // 910–966, 968–985, 996–998, 1000–1061
func IsWarningCode(code int) bool      // 2000–2019, 3008
func IsConfigCode(code int) bool       // 910, 911, 912, 917, 958, 959, 975, 989
func GetMensaje(code int) string
```

Named constants exist for every SIAT code as `siat.Code<Name>`. See [Handle errors](../how-to/handle-errors.md).

---

## Constants

```go
const (
    AmbienteProduccion = 1
    AmbientePruebas    = 2
)

const (
    ModalidadElectronica   = 1   // requires XMLDSig signature
    ModalidadComputarizada = 2   // control code, no signature
)

const (
    EmisionOnline  = 1
    EmisionOffline = 2   // contingency
    EmisionMasiva  = 3
)
```

---

## Helper types

```go
type MensajeServicio = common.MensajeServicio   // { Codigo int; Descripcion string }

type Map map[string]interface{}
func (m Map) ToJSON() (string, error)
func (m Map) ToStruct(v interface{}) error
func (m Map) Sum() float64
```

Utility functions — CUF generation, hashing, compression, signing, export — live in `pkg/utils`. See [Utilities](utilities.md).

---

## Related

| | |
| :--- | :--- |
| HTTP tuning and middleware | [Reference: Configuration](configuration.md) |
| CUF, hashing, signing helpers | [Reference: Utilities](utilities.md) |
| Working example end to end | [Tutorial](../tutorial/first-invoice.md) |
| Why the API is shaped this way | [Explanation: Architecture](../explanation/architecture.md) |
