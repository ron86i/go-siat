# Referencia de API

<p align="right">
  <a href="../../en/reference/api.md">🇬🇧 English</a> · <a href="../README.md">Índice de documentación</a>
</p>

Superficie pública completa de `github.com/ron86i/go-siat/v3`. Para el godoc generado mirá [pkg.go.dev](https://pkg.go.dev/github.com/ron86i/go-siat/v3).

```go
import (
    siat "github.com/ron86i/go-siat/v3"
    "github.com/ron86i/go-siat/v3/pkg/models"
    "github.com/ron86i/go-siat/v3/pkg/models/invoices"
    "github.com/ron86i/go-siat/v3/pkg/utils"
)
```

El sufijo `/v2` es obligatorio — es una versión semántica de import, no una ruta opcional.

---

## Contenido

- [Construcción del cliente](#construcción-del-cliente)
- [Fachadas de servicio](#fachadas-de-servicio)
- [Interfaces de servicio](#interfaces-de-servicio)
- [Builders de solicitud](#builders-de-solicitud)
- [Manejo de respuestas](#manejo-de-respuestas)
- [Errores](#errores)
- [Constantes](#constantes)
- [Tipos auxiliares](#tipos-auxiliares)

---

## Construcción del cliente

```go
func New(config Config) (*SiatServices, error)
func NewWithMiddleware(config Config, middlewares ...HTTPMiddleware) (*SiatServices, error)
```

`New` valida la configuración y devuelve el cliente. Falla cuando `Token`, `CodigoSistema` o `BaseURL` están vacíos, cuando `Nit` no es positivo, o cuando `CodigoAmbiente` no es ni `AmbienteProduccion` ni `AmbientePruebas`.

```go
s, err := siat.New(siat.Config{
    Token:          "TU_TOKEN_SIAT",
    Nit:            123456789,
    CodigoSistema:  "TU_CODIGO_SISTEMA",
    CodigoAmbiente: siat.AmbientePruebas,
    BaseURL:        "https://pilotosiatservicios.impuestos.gob.bo/sfd",
})
```

### Config

`siat.Config` es un alias de `ports.Config`.

| Campo | Tipo | Obligatorio | Para qué sirve |
| :--- | :--- | :--- | :--- |
| `Token` | `string` | sí | Token de API del SIAT, se envía como `apiKey: TokenApi <token>` |
| `Nit` | `int64` | sí | NIT del contribuyente, se inyecta en cada solicitud |
| `CodigoSistema` | `string` | sí | Código de sistema que el SIAT asignó a tu software |
| `CodigoAmbiente` | `int` | sí | `AmbienteProduccion` o `AmbientePruebas` |
| `BaseURL` | `string` | sí | Raíz del ambiente; las rutas de servicio se concatenan |
| `UserAgent` | `string` | no | Por defecto `"go-siat"` |
| `TraceId` | `string` | no | Se envía como `X-Trace-ID` cuando no está vacío |
| `HTTPClient` | `*http.Client` | no | Por defecto `NewHTTPClient(DefaultHTTPConfig())` |
| `CredentialSign` | `CredentialSign` | no | Solo hace falta en modalidad electrónica |

`Nit`, `CodigoSistema` y `CodigoAmbiente` se inyectan en las solicitudes salientes únicamente donde el campo correspondiente sigue en cero — un valor explícito puesto en el builder siempre gana.

```go
func (c Config) SignXML(xmlBytes []byte) ([]byte, error)
```

`Config` cumple por sí mismo la interfaz `XMLSigner`, y por eso `s.Config()` es lo que le pasás a `WithFactura`, `WithFacturas`, `WithDocumento` y a las funciones `utils.Export*`.

### Credenciales de firma

```go
func NewP12Credential(p12 any, password string) CredentialSign
func NewPEMCredential(cert, privateKey any) CredentialSign
func (cf CredentialSign) GetType() string   // "P12" | "PEM" | "UNKNOWN"
```

Los dos constructores aceptan una ruta de archivo (`string`) o bytes crudos (`[]byte`). No entran en pánico si falta el archivo — el error aparece al momento de firmar, y `GetType()` devuelve `"UNKNOWN"`.

### Override de configuración por solicitud

```go
func WithDynamicConfig(ctx context.Context, config Config) context.Context
func GetConfigFromContext(ctx context.Context) (Config, bool)
```

Sobrescribe `Token`, `Nit`, `CodigoSistema`, `CodigoAmbiente`, `TraceId` y `CredentialSign` para una sola llamada — ganan los campos que no estén en cero. Pensado para servicios multi-inquilino que comparten un cliente.

```go
ctx = ports.WithDynamicConfig(ctx, siat.Config{Nit: nitInquilino, Token: tokenInquilino})
resp, err := s.Electronica().RecepcionFactura(ctx, req)
```

### Cliente HTTP

```go
type HTTPConfig = services.HTTPConfig
func DefaultHTTPConfig() HTTPConfig
func NewHTTPClient(cfg HTTPConfig) *http.Client
type HTTPMiddleware = middleware.HTTPMiddleware
```

Mirá [Configuración](configuracion.md) para los campos y las recomendaciones de ajuste.

---

## Fachadas de servicio

Doce accesores sobre `*SiatServices`, cada uno mapeado a un endpoint SOAP.

| Método | Devuelve | Endpoint |
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

Más `s.Config() Config`, que devuelve la configuración activa.

Qué fachada exige cada sector es fijo — mirá [Sectores](../explanation/sectores.md#cómo-elige-un-sector-su-endpoint).

---

## Interfaces de servicio

Todos los métodos tienen la misma forma:

```go
Metodo(ctx context.Context, req models.<TipoSolicitud>) (*soap.EnvelopeResponse[<TipoRespuesta>], error)
```

### FacturacionService

La base de nueve métodos, implementada por todas las fachadas de facturación.

| Método | Builder de solicitud |
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

Dos interfaces la extienden por embebido:

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

`Electronica()` y `Computarizada()` devuelven las dos `SiatSuministroEnergiaService`, así que cambiar de modalidad nunca cambia el punto de llamada más allá del nombre de la fachada.

### SiatCodigosService

| Método | Builder | Para qué |
| :--- | :--- | :--- |
| `SolicitudCuis` | `NewCuisBuilder()` | Pedir el código de sistema |
| `SolicitudCuisMasivo` | `NewCuisMasivoBuilder()` | CUIS masivo |
| `SolicitudCufd` | `NewCufdBuilder()` | Pedir el código diario |
| `SolicitudCufdMasivo` | `NewCufdMasivoBuilder()` | CUFD masivo |
| `VerificarNit` | `NewVerificarNitBuilder()` | Verificar el NIT de un cliente |
| `NotificaCertificadoRevocado` | `NewNotificaCertificadoRevocadoBuilder()` | Informar un certificado revocado |
| `VerificarComunicacion` | `NewVerificarComunicacionCodigosBuilder()` | Chequeo de salud |

`NewVerificarNitBuilder` recibe el NIT a verificar con `WithNitParaVerificacion`, no con `WithNit` — `WithNit` apuntaría a tu propio NIT.

### SiatOperacionesService

`RegistroPuntoVenta`, `ConsultaPuntoVenta`, `CierrePuntoVenta`, `RegistroPuntoVentaComisionista`, `RegistroEventosSignificativos`, `ConsultaEventosSignificativos`, `CierreOperacionesSistema`, `VerificarComunicacion`.

Los eventos significativos son la forma de registrar una caída antes de enviar un paquete de contingencia.

### SiatSincronizacionService

Catálogos de solo lectura. `SincronizarFechaHora`, `SincronizarActividades`, `SincronizarListaActividadesDocumentoSector`, `SincronizarListaProductosServicios`, `SincronizarListaLeyendasFactura`, `SincronizarListaMensajesServicios`, `VerificarComunicacion`, más once métodos `SincronizarParametrica*`:

`EventosSignificativos`, `MotivoAnulacion`, `PaisOrigen`, `TipoDocumentoIdentidad`, `TipoDocumentoSector`, `TipoEmision`, `TipoHabitacion`, `TipoMetodoPago`, `TipoMoneda`, `TipoPuntoVenta`, `TiposFactura`, `UnidadMedida`.

Traé estos catálogos en vez de hardcodear códigos — el SIAT los cambia.

### SiatDocumentoAjusteService

`RecepcionDocumentoAjuste`, `AnulacionDocumentoAjuste`, `ReversionAnulacionDocumentoAjuste`, `VerificacionEstadoDocumentoAjuste`, `VerificarComunicacion`.

### SiatBoletoAereoService

`RecepcionMasivaFactura`, `ValidacionRecepcionMasivaFactura`, `AnulacionFactura`, `ReversionAnulacionFactura`, `VerificacionEstadoFactura`, `VerificarComunicacion`. Notá que no hay un `RecepcionFactura` individual — los boletos aéreos salen de forma masiva.

### SiatRecepcionComprasService

`RecepcionPaqueteCompras`, `ValidacionRecepcionPaqueteCompras`, `ConsultaCompras`, `ConfirmacionCompras`, `AnulacionCompra`, `VerificarComunicacion`.

---

## Builders de solicitud

Todos en `pkg/models`, todos planos — `models.NewXBuilder()`, nunca `models.Codigos().NewXBuilder()`. Todos terminan en `.Build()`, que devuelve un tipo de solicitud opaco.

### Setters comunes

La mayoría de los builders de facturación aceptan estos:

| Setter | Tipo |
| :--- | :--- |
| `WithCuis` / `WithCufd` / `WithCuf` | `string` |
| `WithCodigoModalidad` | `int` |
| `WithCodigoEmision` | `int` |
| `WithCodigoSucursal` / `WithCodigoPuntoVenta` | `int` |
| `WithCodigoDocumentoSector` | `int` |
| `WithTipoFacturaDocumento` | `int` |
| `WithFechaEnvio` | `time.Time` |
| `WithNit` / `WithCodigoSistema` / `WithCodigoAmbiente` | solo override — se inyectan desde `Config` |

### Setters de empaquetado de documentos

Estos rompen la cadena fluida: hacen trabajo que puede fallar (marshal, firma, compresión, hash) y devuelven `error` en vez del builder.

```go
func (b *RecepcionFacturaBuilder) WithFactura(factura any, signer XMLSigner) error
func (b *RecepcionPaqueteFacturaBuilder) WithFacturas(facturas []any, signer XMLSigner) error
func (b *RecepcionMasivaFacturaBuilder) WithFacturas(facturas []any, signer XMLSigner) error
func (b *recepcionDocumentoAjusteBuilder) WithDocumento(documento any, signer XMLSigner) error
```

Llamá a `WithCodigoModalidad` **antes** que a estos. Leen la modalidad para decidir si firman; una modalidad en cero produce un documento sin firmar y sin error local.

La alternativa manual es `WithArchivo(base64 string)` más `WithHashArchivo(hash string)`. Usá un enfoque o el otro, nunca los dos.

### Setters exclusivos de lotes

| Setter | Tipo | Builder |
| :--- | :--- | :--- |
| `WithCantidadFacturas` | `int` | paquete, masivo |
| `WithCodigoEvento` | `int64` | solo paquete |
| `WithCodigoRecepcion` | `string` | builders de validación |

### Lista completa de builders

`AnulacionCompra`, `AnulacionDocumentoAjuste`, `AnulacionFactura`, `CierreOperacionesSistema`, `CierrePuntoVenta`, `ConfirmacionCompras`, `ConsultaCompras`, `ConsultaEventoSignificativo`, `ConsultaPuntoVenta`, `Cufd`, `CufdMasivo`, `Cuis`, `CuisMasivo`, `NotificaCertificadoRevocado`, `RecepcionAnexos`, `RecepcionAnexosSuministroEnergia`, `RecepcionDocumentoAjuste`, `RecepcionFactura`, `RecepcionMasivaFactura`, `RecepcionPaqueteCompras`, `RecepcionPaqueteFactura`, `RegistroEventoSignificativo`, `RegistroPuntoVenta`, `RegistroPuntoVentaComisionista`, `ReversionAnulacionDocumentoAjuste`, `ReversionAnulacionFactura`, `SolicitudListaCufdDto`, `SolicitudListaCuisDto`, `SuministroEnergiaAnexo`, `ValidacionRecepcionMasivaFactura`, `ValidacionRecepcionPaqueteCompras`, `ValidacionRecepcionPaqueteFactura`, `VentaAnexo`, `VerificacionEstadoDocumentoAjuste`, `VerificacionEstadoFactura`, `VerificarComunicacionCodigos`, `VerificarComunicacionDocumentoAjuste`, `VerificarComunicacionOperaciones`, `VerificarComunicacionRecepcionCompras`, `VerificarComunicacionSincronizacion`, `VerificarNit`, y los diecisiete builders `SincronizarParametrica*` / `SincronizarLista*`.

Cada uno es `models.New<Nombre>Builder()`.

### Builders de factura

Los builders de documento por sector viven en `pkg/models/invoices`, tres por sector:

```go
invoices.New<Sector>CabeceraBuilder()   // cabecera
invoices.New<Sector>DetalleBuilder()    // una línea de detalle
invoices.New<Sector>Builder()           // WithModalidad, WithCabecera, AddDetalle, Build
```

Los 50 con builder están listados en [Sectores](../explanation/sectores.md#catálogo-completo-de-sectores).

---

## Manejo de respuestas

Todos los métodos devuelven `*soap.EnvelopeResponse[T]`:

```go
resp.Body.Fault      // *Fault — si no es nil, el sobre mismo fue rechazado
resp.Body.Content    // T — el payload tipado
```

El nombre del campo del payload varía según la operación:

| Respuesta | Campo |
| :--- | :--- |
| `RecepcionFacturaResponse` y la mayoría de tipos de facturación | `RespuestaServicioFacturacion` |
| `RecepcionDocumentoAjusteResponse` | `RespuestaRecepcionFactura` |
| `RecepcionCompras*Response` | `RespuestaRecepcion` |
| `CuisResponse` | `RespuestaCuis` |
| `CufdResponse` | `RespuestaCufd` |
| `VerificarNitResponse` | `RespuestaVerificarNit` |
| `VerificarComunicacionResponse` | `RespuestaComunicacion` |

`RespuestaRecepcion` — la forma detrás de la mayoría de las respuestas de facturación:

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

Devuelve error cuando `Transaccion` es false, o cuando algún mensaje no es un código de advertencia. Las advertencias solas pasan. Llamalo siempre — un error de transporte `nil` no significa aceptación.

---

## Errores

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

Constructores: `NewNetworkError(msg, err)`, `NewSiatError(code, msg)`, `NewAuthError(msg)`, `NewTimeoutError(msg)`.

Clasificadores:

```go
func IsRetryable(err error) bool
func IsNetworkError(err error) bool
func IsRetryableCode(code int) bool    // 123, 967, 991, 995, 999
func IsValidationCode(code int) bool   // 910–966, 968–985, 996–998, 1000–1061
func IsWarningCode(code int) bool      // 2000–2019, 3008
func IsConfigCode(code int) bool       // 910, 911, 912, 917, 958, 959, 975, 989
func GetMensaje(code int) string
```

Hay constantes con nombre para todos los códigos del SIAT como `siat.Code<Nombre>`. Mirá [Manejo de errores](../how-to/manejo-errores.md).

---

## Constantes

```go
const (
    AmbienteProduccion = 1
    AmbientePruebas    = 2
)

const (
    ModalidadElectronica   = 1   // requiere firma XMLDSig
    ModalidadComputarizada = 2   // código de control, sin firma
)

const (
    EmisionOnline  = 1
    EmisionOffline = 2   // contingencia
    EmisionMasiva  = 3
)
```

---

## Tipos auxiliares

```go
type MensajeServicio = common.MensajeServicio   // { Codigo int; Descripcion string }

type Map map[string]interface{}
func (m Map) ToJSON() (string, error)
func (m Map) ToStruct(v interface{}) error
func (m Map) Sum() float64
```

Las funciones utilitarias — generación de CUF, hashing, compresión, firma, exportación — viven en `pkg/utils`. Mirá [Utilidades](utilidades.md).

---

## Relacionado

| | |
| :--- | :--- |
| Ajuste de HTTP y middleware | [Referencia: Configuración](configuracion.md) |
| Helpers de CUF, hash y firma | [Referencia: Utilidades](utilidades.md) |
| Ejemplo funcionando de punta a punta | [Tutorial](../tutorial/primera-factura.md) |
| Por qué la API tiene esta forma | [Explicación: Arquitectura](../explanation/arquitectura.md) |
