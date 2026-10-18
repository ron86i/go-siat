# Configuration reference

<p align="right">
  <a href="../../es/reference/configuracion.md">🇪🇸 Español</a> · <a href="../README.md">Docs index</a>
</p>

Every knob the SDK exposes: identity, environments, HTTP tuning, middleware, tracing and multi-tenancy.

---

## Config

`siat.Config` carries your identity. It is passed once to `siat.New` and reused for every call.

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

| Field | Required | Default | Notes |
| :--- | :--- | :--- | :--- |
| `Token` | yes | — | Sent as `apiKey: TokenApi <token>` |
| `Nit` | yes | — | Must be > 0 |
| `CodigoSistema` | yes | — | Assigned by SIAT per system |
| `CodigoAmbiente` | yes | — | `AmbienteProduccion` or `AmbientePruebas` only |
| `BaseURL` | yes | — | No trailing slash; service paths are appended |
| `UserAgent` | no | `"go-siat"` | Sent on every request |
| `TraceId` | no | — | Sent as `X-Trace-ID` when non-empty |
| `HTTPClient` | no | `NewHTTPClient(DefaultHTTPConfig())` | Cloned, never mutated |
| `CredentialSign` | no | zero value | Only needed for electronic modality |

`New` rejects a config that fails any "required" rule, so a bad setup fails at construction rather than on the first call.

### Environments

| Environment | `CodigoAmbiente` | `BaseURL` |
| :--- | :--- | :--- |
| Pilot / testing | `siat.AmbientePruebas` (2) | `https://pilotosiatservicios.impuestos.gob.bo/sfd` |
| Production | `siat.AmbienteProduccion` (1) | SIN-assigned v3 endpoint; confirm before deployment |

These two must agree. A production `CodigoAmbiente` against the pilot URL is rejected with code 910. Confirm the current production host against SIAT's own documentation before deploying — endpoints do change.

### Loading from the environment

Keep credentials out of source. The SDK has no opinion about how you load them:

```go
cfg := siat.Config{
    Token:          os.Getenv("SIAT_TOKEN"),
    Nit:            mustParseInt64(os.Getenv("SIAT_NIT")),
    CodigoSistema:  os.Getenv("SIAT_CODIGO_SISTEMA"),
    CodigoAmbiente: siat.AmbientePruebas,
    BaseURL:        os.Getenv("SIAT_BASE_URL"),
}
```

`utils.ParseInt64Safe` handles the NIT conversion without a panic path.

### Signing credentials

```go
CredentialSign: siat.NewP12Credential("cert.p12", os.Getenv("P12_PASSWORD"))
CredentialSign: siat.NewPEMCredential("cert.crt", "key.pem")
```

Both accept a path (`string`) or bytes (`[]byte`). See [Sign invoices](../how-to/sign-invoices.md).

---

## HTTP client

If you leave `HTTPClient` nil, the SDK builds one from `DefaultHTTPConfig()`.

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

| Field | Default | Meaning |
| :--- | :--- | :--- |
| `Timeout` | 45s | Total per-request budget, including body read |
| `MaxIdleConns` | 100 | Global idle pool |
| `MaxConnsPerHost` | 0 | Concurrent connections per host — 0 means unlimited |
| `MaxIdleConnsPerHost` | 100 | Idle connections kept per host |
| `TLSMinVersion` | TLS 1.2 | Minimum accepted TLS version |
| `DisableKeepAlives` | false | Leave false; SIAT calls come in bursts |
| `IdleConnTimeout` | 90s | When an idle connection is closed |
| `DialTimeout` | 10s | TCP connect budget |
| `DialKeepAlive` | 30s | TCP keep-alive probe interval |
| `TLSHandshakeTimeout` | 10s | TLS handshake budget |

The transport also honours `HTTP_PROXY` / `HTTPS_PROXY` / `NO_PROXY` via `http.ProxyFromEnvironment`.

### Customising

Start from the defaults and change what you need:

```go
httpCfg := siat.DefaultHTTPConfig()
httpCfg.Timeout = 90 * time.Second        // large batches
httpCfg.MaxConnsPerHost = 20              // cap concurrency

s, err := siat.New(siat.Config{
    // ... identity fields
    HTTPClient: siat.NewHTTPClient(httpCfg),
})
```

Two practical notes:

**`Timeout` is a ceiling, not a strategy.** A context deadline is what actually lets you cancel — set both, and keep the context deadline shorter than the client timeout for individual calls.

**Batch sends need a bigger budget.** A package of hundreds of invoices is a large upload; 45 seconds can be tight. Raise `Timeout` for the client you use for batching, or give batch calls their own client.

You can also supply a fully custom `*http.Client`. The SDK clones it before attaching middleware, so your instance is never mutated.

---

## Middleware

```go
type HTTPMiddleware interface {
    WrapTransport(base http.RoundTripper) http.RoundTripper
}

func NewWithMiddleware(config Config, middlewares ...HTTPMiddleware) (*SiatServices, error)
```

Middleware wraps the transport, so it sees every request the SDK makes. The first middleware in the list is the outermost.

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

Useful for logging, metrics, circuit breaking and transport-level retry.

**Never log the raw request body at this layer** — it contains your API token in the `apiKey` header and taxpayer data in the payload. Log method, path, status and duration instead.

Retry middleware operating on `http.Response` cannot see SIAT-level rejections, which arrive as HTTP 200 with a rejection inside. Retry transport failures here; handle rejections with `siat.Verify` in your application code. See [Handle errors](../how-to/handle-errors.md).

---

## Tracing

Set `TraceId` and the SDK forwards it as `X-Trace-ID` on every request:

```go
cfg.TraceId = uuid.NewString()
```

For per-request trace IDs, use dynamic config rather than mutating shared state:

```go
ctx = ports.WithDynamicConfig(ctx, siat.Config{TraceId: requestID})
```

---

## Multi-tenancy

One client can serve many taxpayers. `WithDynamicConfig` overrides selected fields for a single call:

```go
import "github.com/ron86i/go-siat/v3/internal/core/ports"

ctx = ports.WithDynamicConfig(ctx, siat.Config{
    Nit:            tenant.Nit,
    Token:          tenant.Token,
    CodigoSistema:  tenant.CodigoSistema,
    CredentialSign: tenant.Credential,
})

resp, err := s.Electronica().RecepcionFactura(ctx, req)
```

Overridable: `Token`, `Nit`, `CodigoSistema`, `CodigoAmbiente`, `TraceId`, `CredentialSign`. Only non-zero values override; everything else falls back to the client's config.

`BaseURL` is **not** overridable — it is bound at construction. Tenants in different environments need separate clients.

A `*SiatServices` is safe for concurrent use, and this mechanism carries per-tenant state on the context rather than the client, so no locking is required.

---

## Constants

```go
siat.AmbienteProduccion      // 1
siat.AmbientePruebas         // 2

siat.ModalidadElectronica    // 1 — XMLDSig signature required
siat.ModalidadComputarizada  // 2 — control code, no signature

siat.EmisionOnline           // 1
siat.EmisionOffline          // 2 — contingency
siat.EmisionMasiva           // 3
```

Everything else — document sectors, invoice types, payment methods, units, annulment reasons — is a SIAT parametric, not an SDK constant. Fetch them through `s.Sincronizacion()` and cache them; do not hardcode.

---

## Related

| | |
| :--- | :--- |
| Full type signatures | [Reference: API](api.md) |
| Certificates and signing | [How-to: Sign invoices](../how-to/sign-invoices.md) |
| Timeout and retry behaviour | [How-to: Handle errors](../how-to/handle-errors.md) |
| Why config is passed once | [Explanation: Architecture](../explanation/architecture.md) |
