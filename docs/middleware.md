# Middleware & Advanced Tracing

The `go-siat` SDK can be easily integrated into modern microservice architectures through its tracing capabilities and custom HTTP client support.

## 🆔 Request Tracing (Trace ID)

Correlate every SIAT SOAP request with your own internal log IDs (Correlation ID or Trace ID) by adding the `X-Trace-ID` header automatically.

```go
client, _ := siat.New(baseUrl, nil)

// Injects "X-Trace-ID: invoice-12345" in every subsequent SOAP call.
client.WithTraceID("invoice-12345").Codigos().SolicitudCuis(ctx, cfg, req)
```

## 🛡️ Customizable HTTP Client

By default, `siat.New` uses an optimized, secure HTTP client (`http.Client`) with:
- **TLS 1.2+** for SIAT compliance.
- **Handshake Timeout (15s)** to prevent hanging connections.
- **Max Idle Connections** configured for high-concurrency environments.

However, if your system uses a **REST proxy**, a custom **CA certificate**, or requires specific **timeouts**, you can pass your own client:

```go
customClient := &http.Client{
    Transport: &http.Transport{
        Proxy: http.ProxyURL(proxyURL),
    },
    Timeout: time.Minute,
}

client, _ := siat.New(baseUrl, customClient)
```

## 📜 Logging Middleware

Coming soon: Support for custom function injection for pre/post request logging and error metrics.
