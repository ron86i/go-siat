# Referencia de configuración

<p align="right">
  <a href="../../en/reference/configuration.md">🇬🇧 English</a> · <a href="../README.md">Índice de documentación</a>
</p>

Todas las perillas que expone el SDK: identidad, ambientes, ajuste de HTTP, middleware, trazabilidad y multi-inquilino.

---

## Config

`siat.Config` lleva tu identidad. Se pasa una sola vez a `siat.New` y se reutiliza en cada llamada.

```go
type Config struct {
    Token          string
    Nit            int64
    CodigoSistema  string
    CodigoAmbiente int
    BaseURL        string
    UserAgent      string
    TraceId        string
    HTTPClient     *http.Client
    CredentialSign CredentialSign
}
```

| Campo | Obligatorio | Por defecto | Notas |
| :--- | :--- | :--- | :--- |
| `Token` | sí | — | Se envía como `apiKey: TokenApi <token>` |
| `Nit` | sí | — | Tiene que ser > 0 |
| `CodigoSistema` | sí | — | Lo asigna el SIAT por sistema |
| `CodigoAmbiente` | sí | — | Solo `AmbienteProduccion` o `AmbientePruebas` |
| `BaseURL` | sí | — | Sin barra final; las rutas de servicio se concatenan |
| `UserAgent` | no | `"go-siat"` | Se envía en cada solicitud |
| `TraceId` | no | — | Se envía como `X-Trace-ID` cuando no está vacío |
| `HTTPClient` | no | `NewHTTPClient(DefaultHTTPConfig())` | Se clona, nunca se muta |
| `CredentialSign` | no | valor cero | Solo hace falta en modalidad electrónica |

`New` rechaza una configuración que incumpla cualquier regla de "obligatorio", así que un setup mal armado falla al construir y no en la primera llamada.

### Ambientes

| Ambiente | `CodigoAmbiente` | `BaseURL` |
| :--- | :--- | :--- |
| Piloto / pruebas | `siat.AmbientePruebas` (2) | `https://pilotosiatservicios.impuestos.gob.bo/sfd` |
| Producción | `siat.AmbienteProduccion` (1) | Dirección v3 asignada por el SIN; confirmar antes del despliegue |

Los dos tienen que coincidir. Un `CodigoAmbiente` de producción contra la URL del piloto se rechaza con el código 910. Confirmá el host de producción vigente contra la documentación del propio SIAT antes de desplegar — los endpoints cambian.

### Cargar desde el entorno

Mantené las credenciales fuera del código fuente. Al SDK le da igual cómo las cargues:

```go
cfg := siat.Config{
    Token:          os.Getenv("SIAT_TOKEN"),
    Nit:            mustParseInt64(os.Getenv("SIAT_NIT")),
    CodigoSistema:  os.Getenv("SIAT_CODIGO_SISTEMA"),
    CodigoAmbiente: siat.AmbientePruebas,
    BaseURL:        os.Getenv("SIAT_BASE_URL"),
}
```

`utils.ParseInt64Safe` resuelve la conversión del NIT sin camino de pánico.

### Credenciales de firma

```go
CredentialSign: siat.NewP12Credential("cert.p12", os.Getenv("P12_PASSWORD"))
CredentialSign: siat.NewPEMCredential("cert.crt", "key.pem")
```

Los dos aceptan una ruta (`string`) o bytes (`[]byte`). Mirá [Firmar facturas](../how-to/firmar-facturas.md).

---

## Cliente HTTP

Si dejás `HTTPClient` en nil, el SDK construye uno a partir de `DefaultHTTPConfig()`.

```go
type HTTPConfig struct {
    Timeout             time.Duration
    MaxIdleConns        int
    MaxConnsPerHost     int
    MaxIdleConnsPerHost int
    TLSMinVersion       uint16
    DisableKeepAlives   bool
    IdleConnTimeout     time.Duration
    DialTimeout         time.Duration
    DialKeepAlive       time.Duration
    TLSHandshakeTimeout time.Duration
}
```

| Campo | Por defecto | Qué significa |
| :--- | :--- | :--- |
| `Timeout` | 45s | Presupuesto total por solicitud, incluida la lectura del cuerpo |
| `MaxIdleConns` | 100 | Pool global de conexiones ociosas |
| `MaxConnsPerHost` | 0 | Conexiones concurrentes por host — 0 significa sin límite |
| `MaxIdleConnsPerHost` | 100 | Conexiones ociosas que se guardan por host |
| `TLSMinVersion` | TLS 1.2 | Versión mínima de TLS aceptada |
| `DisableKeepAlives` | false | Dejalo en false; las llamadas al SIAT vienen en ráfagas |
| `IdleConnTimeout` | 90s | Cuándo se cierra una conexión ociosa |
| `DialTimeout` | 10s | Presupuesto de conexión TCP |
| `DialKeepAlive` | 30s | Intervalo de sondeo keep-alive TCP |
| `TLSHandshakeTimeout` | 10s | Presupuesto del handshake TLS |

El transporte también respeta `HTTP_PROXY` / `HTTPS_PROXY` / `NO_PROXY` vía `http.ProxyFromEnvironment`.

### Personalizar

Partí de los defaults y cambiá lo que necesites:

```go
httpCfg := siat.DefaultHTTPConfig()
httpCfg.Timeout = 90 * time.Second        // lotes grandes
httpCfg.MaxConnsPerHost = 20              // limitar concurrencia

s, err := siat.New(siat.Config{
    // ... campos de identidad
    HTTPClient: siat.NewHTTPClient(httpCfg),
})
```

Dos notas prácticas:

**`Timeout` es un techo, no una estrategia.** Lo que realmente te permite cancelar es un deadline de contexto — poné los dos, y mantené el deadline del contexto más corto que el timeout del cliente para las llamadas individuales.

**Los envíos por lotes necesitan más presupuesto.** Un paquete de cientos de facturas es una subida grande; 45 segundos pueden quedar justos. Subí `Timeout` para el cliente que uses en lotes, o dale a esas llamadas su propio cliente.

También podés pasar un `*http.Client` totalmente propio. El SDK lo clona antes de enchufarle middleware, así que tu instancia nunca se muta.

---

## Middleware

```go
type HTTPMiddleware interface {
    WrapTransport(base http.RoundTripper) http.RoundTripper
}

func NewWithMiddleware(config Config, middlewares ...HTTPMiddleware) (*SiatServices, error)
```

El middleware envuelve el transporte, así que ve todas las solicitudes que hace el SDK. El primer middleware de la lista es el más externo.

```go
type LoggingMiddleware struct{}

func (m *LoggingMiddleware) WrapTransport(base http.RoundTripper) http.RoundTripper {
    return roundTripFunc(func(req *http.Request) (*http.Response, error) {
        start := time.Now()
        resp, err := base.RoundTrip(req)
        log.Printf("%s %s -> %v (%s)", req.Method, req.URL.Path, err, time.Since(start))
        return resp, err
    })
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

s, err := siat.NewWithMiddleware(cfg, &LoggingMiddleware{}, &MetricsMiddleware{})
```

Sirve para logging, métricas, circuit breaking y reintentos a nivel de transporte.

**Nunca registres el cuerpo crudo de la solicitud en esta capa** — contiene tu token de API en la cabecera `apiKey` y datos del contribuyente en el payload. Registrá método, ruta, estado y duración en su lugar.

Un middleware de reintentos que opera sobre `http.Response` no puede ver los rechazos a nivel SIAT, que llegan como HTTP 200 con el rechazo adentro. Reintentá fallas de transporte acá; manejá los rechazos con `siat.Verify` en el código de tu aplicación. Mirá [Manejo de errores](../how-to/manejo-errores.md).

---

## Trazabilidad

Poné `TraceId` y el SDK lo reenvía como `X-Trace-ID` en cada solicitud:

```go
cfg.TraceId = uuid.NewString()
```

Para IDs de traza por solicitud, usá configuración dinámica en vez de mutar estado compartido:

```go
ctx = ports.WithDynamicConfig(ctx, siat.Config{TraceId: requestID})
```

---

## Multi-inquilino

Un solo cliente puede atender a muchos contribuyentes. `WithDynamicConfig` sobrescribe campos elegidos para una sola llamada:

```go
import "github.com/ron86i/go-siat/v3/internal/core/ports"

ctx = ports.WithDynamicConfig(ctx, siat.Config{
    Nit:            inquilino.Nit,
    Token:          inquilino.Token,
    CodigoSistema:  inquilino.CodigoSistema,
    CredentialSign: inquilino.Credential,
})

resp, err := s.Electronica().RecepcionFactura(ctx, req)
```

Se pueden sobrescribir: `Token`, `Nit`, `CodigoSistema`, `CodigoAmbiente`, `TraceId`, `CredentialSign`. Solo los valores distintos de cero sobrescriben; el resto cae al config del cliente.

`BaseURL` **no** se puede sobrescribir — queda fijado en la construcción. Los inquilinos en ambientes distintos necesitan clientes separados.

Un `*SiatServices` es seguro para uso concurrente, y este mecanismo lleva el estado por inquilino en el contexto y no en el cliente, así que no hace falta ningún lock.

---

## Constantes

```go
siat.AmbienteProduccion      // 1
siat.AmbientePruebas         // 2

siat.ModalidadElectronica    // 1 — requiere firma XMLDSig
siat.ModalidadComputarizada  // 2 — código de control, sin firma

siat.EmisionOnline           // 1
siat.EmisionOffline          // 2 — contingencia
siat.EmisionMasiva           // 3
```

Todo lo demás — documentos sector, tipos de factura, métodos de pago, unidades, motivos de anulación — es una paramétrica del SIAT, no una constante del SDK. Traelas con `s.Sincronizacion()` y cacheálas; no las hardcodees.

---

## Relacionado

| | |
| :--- | :--- |
| Firmas de tipos completas | [Referencia: API](api.md) |
| Certificados y firma | [Guía: Firmar facturas](../how-to/firmar-facturas.md) |
| Comportamiento de timeouts y reintentos | [Guía: Manejo de errores](../how-to/manejo-errores.md) |
| Por qué el config se pasa una sola vez | [Explicación: Arquitectura](../explanation/arquitectura.md) |
