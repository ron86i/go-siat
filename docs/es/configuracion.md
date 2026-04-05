# Configuración

[← Volver al Índice](README.md)

> Guía para configurar el cliente HTTP, implementar middlewares personalizados, habilitar trazabilidad distribuida y entender las constantes del SDK.

---

## Tabla de Contenidos

1. [Autenticación (Config)](#autenticación-config)
2. [Configuración del Cliente HTTP](#configuración-del-cliente-http)
3. [Sistema de Middleware](#sistema-de-middleware)
4. [Trazabilidad Distribuida](#trazabilidad-distribuida)
5. [Constantes](#constantes)
6. [Tipo Utilitario Map](#tipo-utilitario-map)

---

## Autenticación (Config)

Cada llamada a un servicio requiere un `siat.Config` con tu token de autenticación:

```go
cfg := siat.Config{
    Token:     "tu_token_api_siat",      // Requerido
    UserAgent: "MiApp/1.0 (Bolivia)",    // Opcional (por defecto "go-siat")
    TraceId:   "req-12345",              // Opcional (para trazabilidad distribuida)
}
```

| Campo | Tipo | Requerido | Descripción |
|:------|:-----|:----------|:------------|
| `Token` | `string` | ✅ | Token de autenticación de la API del SIAT |
| `UserAgent` | `string` | ❌ | Header User-Agent personalizado (por defecto: `"go-siat"`) |
| `TraceId` | `string` | ❌ | ID de traza inyectado como header `X-Trace-ID` |

### Usando el Helper `WithConfig`

Si has establecido un trace ID en el servicio, puedes usar `WithConfig` para inyectarlo automáticamente:

```go
s, _ := siat.New(baseURL, nil)
s.WithTraceID("trace-abc-123")

cfg := s.WithConfig("tu_token")
// cfg.TraceId es automáticamente "trace-abc-123"
```

---

## Configuración del Cliente HTTP

### Configuración por Defecto

Por defecto, `siat.New()` crea un cliente HTTP optimizado para producción:

| Configuración | Valor por Defecto | Descripción |
|:--------------|:-----------------|:------------|
| Timeout | 45 segundos | Timeout total de la solicitud |
| MaxIdleConns | 100 | Tamaño global del pool de conexiones |
| MaxConnsPerHost | 10 | Conexiones concurrentes por host |
| MaxIdleConnsPerHost | 5 | Conexiones ociosas por host en cache |
| Versión TLS Mínima | TLS 1.2 | Versión mínima del protocolo TLS |
| TLS Handshake Timeout | 15 segundos | Timeout del handshake TLS |
| Proxy | Desde variables de entorno | Respeta `HTTP_PROXY`/`HTTPS_PROXY` |

### Configuración Personalizada

#### Opción 1: Usando `HTTPConfig`

```go
import "time"

cfg := siat.DefaultHTTPConfig()
cfg.Timeout = 60 * time.Second       // Timeout más largo para redes lentas
cfg.MaxConnsPerHost = 5              // Limitar conexiones concurrentes

client := siat.NewHTTPClient(cfg)
s, err := siat.New(baseURL, client)
```

#### Opción 2: Traer Tu Propio Cliente

```go
client := &http.Client{
    Timeout: 30 * time.Second,
    Transport: &http.Transport{
        Proxy:               http.ProxyFromEnvironment,
        MaxIdleConns:        50,
        MaxConnsPerHost:     3,
        TLSHandshakeTimeout: 10 * time.Second,
    },
}

s, err := siat.New(baseURL, client)
```

> [!NOTE]
> Cuando pasas tu propio `http.Client`, el SDK lo **clona** para evitar mutaciones no intencionadas. Tu cliente original permanece sin cambios.

### Campos de `HTTPConfig`

```go
type HTTPConfig struct {
    Timeout              time.Duration // Timeout total de la solicitud
    MaxIdleConns         int           // Pool global de conexiones ociosas
    MaxConnsPerHost      int           // Máx conexiones concurrentes por host
    MaxIdleConnsPerHost  int           // Máx conexiones ociosas por host
    TLSHandshakeTimeout  time.Duration // Timeout del handshake TLS
    TLSMinVersion        uint16        // Versión TLS mínima (tls.VersionTLS12)
}
```

---

## Sistema de Middleware

El sistema de middleware permite interceptar, modificar y observar todas las solicitudes HTTP realizadas por el SDK. Es útil para:

- **Logging** de todas las solicitudes y respuestas al SIAT
- **Métricas** (latencia, tasas de error)
- **Retry** con backoff exponencial
- **Circuit breaker**
- **Mutación** de solicitudes/respuestas

### Interfaz `HTTPMiddleware`

```go
type HTTPMiddleware interface {
    WrapTransport(base http.RoundTripper) http.RoundTripper
}
```

### Ejemplo: Middleware de Logging

```go
type LoggingMiddleware struct {
    logger *log.Logger
}

func (m *LoggingMiddleware) WrapTransport(base http.RoundTripper) http.RoundTripper {
    return &loggingRoundTripper{base: base, logger: m.logger}
}

type loggingRoundTripper struct {
    base   http.RoundTripper
    logger *log.Logger
}

func (rt *loggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
    start := time.Now()
    rt.logger.Printf("→ %s %s", req.Method, req.URL.Path)

    resp, err := rt.base.RoundTrip(req)

    if err != nil {
        rt.logger.Printf("← ERROR después de %v: %v", time.Since(start), err)
    } else {
        rt.logger.Printf("← %d después de %v", resp.StatusCode, time.Since(start))
    }

    return resp, err
}
```

### Usando Middlewares

```go
s, err := siat.NewWithMiddleware(
    baseURL,
    nil,                                     // Usar config HTTP por defecto
    &LoggingMiddleware{logger: miLogger},    // Primero = más externo
    &MetricsMiddleware{...},                 // Middleware interno
)
```

### Orden de la Cadena de Middleware

Los middlewares se aplican en **orden inverso**, por lo que el primer middleware de la lista es el envoltorio más externo:

```
Flujo de solicitud:  LoggingMiddleware → MetricsMiddleware → Transport por Defecto → SIAT
Flujo de respuesta:  SIAT → Transport por Defecto → MetricsMiddleware → LoggingMiddleware
```

> [!TIP]
> El sistema de middleware envuelve `http.RoundTripper`, lo que significa que intercepta la solicitud/respuesta HTTP cruda. La construcción del sobre SOAP ocurre **antes** de que el middleware vea la solicitud.

---

## Trazabilidad Distribuida

### `WithTraceID`

Establece un identificador de traza que se inyecta en todas las solicitudes subsiguientes vía el header HTTP `X-Trace-ID`:

```go
s, _ := siat.New(baseURL, nil)

// Establecer una vez, se aplica a todas las solicitudes
s.WithTraceID("trace-abc-123-def")

// Todas las llamadas subsiguientes incluyen X-Trace-ID: trace-abc-123-def
resp, err := s.Codigos().SolicitudCuis(ctx, cfg, req)
```

### Integración con Sistemas de Tracing

```go
// Ejemplo de integración con OpenTelemetry
func manejarFactura(ctx context.Context, s *siat.SiatServices) {
    span := trace.SpanFromContext(ctx)
    traceID := span.SpanContext().TraceID().String()

    s.WithTraceID(traceID)

    // Todas las llamadas al SIAT ahora llevan el ID de traza
    // Correlacionar con tu plataforma de observabilidad
}
```

---

## Constantes

### Constantes de Ambiente

```go
const (
    siat.AmbienteProduccion  = 1  // Producción: facturas reales con validez tributaria
    siat.AmbientePruebas     = 2  // Pruebas: ambiente sandbox para desarrollo
)
```

### Constantes de Modalidad

```go
const (
    siat.ModalidadElectronica    = 1  // Requiere firma digital (XMLDSig)
    siat.ModalidadComputarizada  = 2  // No requiere firma digital
)
```

### Constantes de Tipo de Emisión

```go
const (
    siat.EmisionOnline   = 1  // En línea: conexión en tiempo real con el SIAT
    siat.EmisionOffline  = 2  // Fuera de línea: modo contingencia (requiere evento significativo)
    siat.EmisionMasiva   = 3  // Masiva: procesamiento por lotes de alto volumen
)
```

### Usando las Constantes

```go
req := models.Codigos().NewCuisBuilder().
    WithCodigoAmbiente(siat.AmbientePruebas).
    WithCodigoModalidad(siat.ModalidadElectronica).
    Build()

cuf, _ := utils.GenerarCUF(nit, time.Now(), 0,
    siat.ModalidadElectronica,
    siat.EmisionOnline,
    1, 1, 1, 0, control)
```

---

## Tipo Utilitario Map

`siat.Map` es un alias para `map[string]interface{}` con métodos utilitarios para JSON y operaciones numéricas:

### `ToJSON`

```go
m := siat.Map{"clave": "valor", "cantidad": 42}
jsonStr, err := m.ToJSON()
// `{"cantidad":42,"clave":"valor"}`
```

### `Sum`

Suma todos los valores numéricos en el mapa (float64, float32, int, int64, int32):

```go
m := siat.Map{
    "subtotal":   100.50,
    "impuesto":   13.065,
    "descuento":  10,
}
total := m.Sum() // 123.565
```

### `ToStruct`

Convierte el mapa a un struct de Go usando marshaling JSON:

```go
type Factura struct {
    Numero int     `json:"numero"`
    Total  float64 `json:"total"`
}

m := siat.Map{"numero": 42, "total": 100.50}
var fac Factura
err := m.ToStruct(&fac)
// fac.Numero = 42, fac.Total = 100.50
```

---

[← Volver al Índice](README.md)
