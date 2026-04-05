# Configuration

[← Back to Index](README.md)

> Guide to configuring the HTTP client, implementing custom middlewares, enabling distributed tracing, and understanding SDK constants.

---

## Table of Contents

1. [Authentication (Config)](#authentication-config)
2. [HTTP Client Configuration](#http-client-configuration)
3. [Middleware System](#middleware-system)
4. [Distributed Tracing](#distributed-tracing)
5. [Constants](#constants)
6. [Map Utility Type](#map-utility-type)

---

## Authentication (Config)

Every service call requires a `siat.Config` with your authentication token:

```go
cfg := siat.Config{
    Token:     "your_siat_api_token",   // Required
    UserAgent: "MyApp/1.0 (Bolivia)",   // Optional (defaults to "go-siat")
    TraceId:   "req-12345",             // Optional (for distributed tracing)
}
```

| Field | Type | Required | Description |
|:------|:-----|:---------|:------------|
| `Token` | `string` | ✅ | SIAT API authentication token |
| `UserAgent` | `string` | ❌ | Custom User-Agent header (default: `"go-siat"`) |
| `TraceId` | `string` | ❌ | Trace ID injected as `X-Trace-ID` header |

### Using `WithConfig` Helper

If you've set a trace ID on the service, you can use `WithConfig` to auto-inject it:

```go
s, _ := siat.New(baseURL, nil)
s.WithTraceID("trace-abc-123")

cfg := s.WithConfig("your_token")
// cfg.TraceId is automatically "trace-abc-123"
```

---

## HTTP Client Configuration

### Default Configuration

By default, `siat.New()` creates an HTTP client optimized for production:

| Setting | Default Value | Description |
|:--------|:-------------|:------------|
| Timeout | 45 seconds | Total request timeout |
| MaxIdleConns | 100 | Global connection pool size |
| MaxConnsPerHost | 10 | Concurrent connections per host |
| MaxIdleConnsPerHost | 5 | Idle connections cached per host |
| TLS Min Version | TLS 1.2 | Minimum TLS protocol version |
| TLS Handshake Timeout | 15 seconds | Handshake timeout |
| Proxy | From environment | Respects `HTTP_PROXY`/`HTTPS_PROXY` |

### Custom Configuration

#### Option 1: Using `HTTPConfig`

```go
import "time"

cfg := siat.DefaultHTTPConfig()
cfg.Timeout = 60 * time.Second       // Longer timeout for slow networks
cfg.MaxConnsPerHost = 5              // Limit concurrent connections

client := siat.NewHTTPClient(cfg)
s, err := siat.New(baseURL, client)
```

#### Option 2: Bring Your Own Client

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
> When you pass your own `http.Client`, the SDK **clones** it to avoid unintended mutations. Your original client remains unchanged.

### `HTTPConfig` Fields

```go
type HTTPConfig struct {
    Timeout              time.Duration // Total request timeout
    MaxIdleConns         int           // Global idle connection pool
    MaxConnsPerHost      int           // Max concurrent connections per host
    MaxIdleConnsPerHost  int           // Max idle connections per host
    TLSHandshakeTimeout  time.Duration // TLS handshake timeout
    TLSMinVersion        uint16        // Min TLS version (tls.VersionTLS12)
}
```

---

## Middleware System

The middleware system allows you to intercept, modify, and observe all HTTP requests made by the SDK. This is useful for:

- **Logging** all SIAT requests and responses
- **Metrics** collection (latency, error rates)
- **Retry** with exponential backoff
- **Circuit breaker** patterns
- **Request/Response** mutation

### `HTTPMiddleware` Interface

```go
type HTTPMiddleware interface {
    WrapTransport(base http.RoundTripper) http.RoundTripper
}
```

### Example: Logging Middleware

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
        rt.logger.Printf("← ERROR after %v: %v", time.Since(start), err)
    } else {
        rt.logger.Printf("← %d after %v", resp.StatusCode, time.Since(start))
    }

    return resp, err
}
```

### Example: Metrics Middleware

```go
type MetricsMiddleware struct {
    requestCounter  prometheus.Counter
    latencyHistogram prometheus.Histogram
}

func (m *MetricsMiddleware) WrapTransport(base http.RoundTripper) http.RoundTripper {
    return &metricsRoundTripper{base: base, metrics: m}
}

type metricsRoundTripper struct {
    base    http.RoundTripper
    metrics *MetricsMiddleware
}

func (rt *metricsRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
    rt.metrics.requestCounter.Inc()
    start := time.Now()

    resp, err := rt.base.RoundTrip(req)

    rt.metrics.latencyHistogram.Observe(time.Since(start).Seconds())
    return resp, err
}
```

### Using Middlewares

```go
s, err := siat.NewWithMiddleware(
    baseURL,
    nil,                                    // Use default HTTP config
    &LoggingMiddleware{logger: myLogger},   // First = outermost
    &MetricsMiddleware{...},                // Inner middleware
)
```

### Middleware Chain Order

Middlewares are applied in **reverse order**, so the first middleware in the list is the outermost wrapper:

```
Request flow:  LoggingMiddleware → MetricsMiddleware → Default Transport → SIAT
Response flow: SIAT → Default Transport → MetricsMiddleware → LoggingMiddleware
```

> [!TIP]
> The middleware system wraps `http.RoundTripper`, which means it intercepts the raw HTTP request/response. SOAP envelope construction happens **before** the middleware sees the request.

---

## Distributed Tracing

### `WithTraceID`

Sets a trace identifier that is injected into all subsequent requests via the `X-Trace-ID` HTTP header:

```go
s, _ := siat.New(baseURL, nil)

// Set once, applied to all requests
s.WithTraceID("trace-abc-123-def")

// All subsequent calls include X-Trace-ID: trace-abc-123-def
resp, err := s.Codigos().SolicitudCuis(ctx, cfg, req)
```

### Integration with Tracing Systems

```go
// OpenTelemetry integration example
func handleInvoice(ctx context.Context, s *siat.SiatServices) {
    span := trace.SpanFromContext(ctx)
    traceID := span.SpanContext().TraceID().String()

    s.WithTraceID(traceID)

    // All SIAT calls now carry the trace ID
    // Correlate with your observability platform
}
```

### Using `WithConfig` with Trace ID

```go
s.WithTraceID("my-trace-id")
cfg := s.WithConfig("my-token")
// cfg.Token = "my-token"
// cfg.TraceId = "my-trace-id" (auto-injected)
```

---

## Constants

### Environment Constants

```go
const (
    siat.AmbienteProduccion  = 1  // Production: real invoices with tax validity
    siat.AmbientePruebas     = 2  // Testing: sandbox environment for development
)
```

### Modality Constants

```go
const (
    siat.ModalidadElectronica    = 1  // Requires digital signature (XMLDSig)
    siat.ModalidadComputarizada  = 2  // No digital signature required
)
```

### Emission Type Constants

```go
const (
    siat.EmisionOnline   = 1  // Online: real-time connection to SIAT
    siat.EmisionOffline  = 2  // Offline: contingency mode (significant event required)
    siat.EmisionMasiva   = 3  // Massive: high-volume batch processing
)
```

### Using Constants

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

## Map Utility Type

`siat.Map` is an alias for `map[string]interface{}` with utility methods for JSON and numeric operations:

### `ToJSON`

```go
m := siat.Map{"key": "value", "count": 42}
jsonStr, err := m.ToJSON()
// `{"count":42,"key":"value"}`
```

### `Sum`

Sums all numeric values in the map (float64, float32, int, int64, int32):

```go
m := siat.Map{
    "subtotal":  100.50,
    "tax":       13.065,
    "discount":  10,
}
total := m.Sum() // 123.565
```

### `ToStruct`

Converts the map to a Go struct using JSON marshaling:

```go
type Invoice struct {
    Number int    `json:"number"`
    Total  float64 `json:"total"`
}

m := siat.Map{"number": 42, "total": 100.50}
var inv Invoice
err := m.ToStruct(&inv)
// inv.Number = 42, inv.Total = 100.50
```

---

[← Back to Index](README.md)
