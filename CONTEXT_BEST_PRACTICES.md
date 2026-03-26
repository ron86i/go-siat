# Context Best Practices Guide for go-siat

## Introduction

The **go-siat** SDK uses `context.Context` to propagate deadlines, cancellation signals, and request information across API calls. This guide explains how to correctly use contexts to maximize the reliability of your application.

## Core Principles

### 1. **Never use `context.Background()` directly in SDK methods**

❌ **Incorrect:**
```go
// NEVER do this
config := siat.Config{Token: "myToken"}
resp, err := client.Codes().VerificarNit(context.Background(), config, nitRequest)
```

✅ **Correct:**
```go
// Callers provide the context
ctx := context.Background()
if deadline, ok := time.ParseDuration("30s"); ok {
    ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
    defer cancel()
}
config := siat.WithConfig("myToken")
resp, err := client.Codes().VerificarNit(ctx, config, nitRequest)
```

**Why:** `context.Background()` does not have any deadline or cancellation signal. The SIAT server may take longer than tolerated, and your application will not be able to stop the request.

---

### 2. **Always set a deadline with `context.WithTimeout()` or `context.WithDeadline()`**

#### Option A: Timeout for individual operations
```go
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel() // CRITICAL: always defer the cancel

config := client.WithConfig("myToken")
resp, err := client.Codes().VerificarNit(ctx, config, nitRequest)
if err != nil {
    if errors.Is(err, context.DeadlineExceeded) {
        log.Println("Timeout: SIAT took more than 30 seconds")
    }
}
```

#### Option B: Absolute Deadline
```go
deadline := time.Now().Add(45 * time.Second)
ctx, cancel := context.WithDeadline(context.Background(), deadline)
defer cancel()

// Use ctx in SDK calls
```

#### Option C: Inherit deadline from parent context
```go
func ProcessSiatRequest(ctx context.Context) error {
    // If ctx has a deadline, it is automatically inherited
    config := client.WithConfig("myToken")
    resp, err := client.Codes().VerificarNit(ctx, config, nitRequest)
    return err
}
```

**Important**: Always call `defer cancel()` after `WithTimeout()` or `WithDeadline()` to release resources.

---

### 3. **Propagate contexts from HTTP handlers**

In a web application, always propagate the context from the HTTP handler:

```go
// In your HTTP server
func handleFacturaRequest(w http.ResponseWriter, r *http.Request) {
    // r.Context() contains deadline, cancellation, and HTTP client values
    ctx := r.Context()
    
    config := siatClient.WithConfig(getUserToken(r))
    resp, err := siatClient.PurchaseSale().RecepcionFactura(ctx, config, facturaData)
    
    if err != nil {
        if errors.Is(err, context.Canceled) {
            // Client disconnected
            return
        }
        if errors.Is(err, context.DeadlineExceeded) {
            http.Error(w, "Timeout in SIAT", http.StatusGatewayTimeout)
            return
        }
        http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/xml")
    // Use GetContent() to extract the XML response from inside the SOAP envelope
}
```

---

### 4. **Handling Cancellation**

The SDK respects context cancellation. If you cancel the context, the network connection will be closed immediately.

```go
ctx, cancel := context.WithCancel(context.Background())

go func() {
    time.Sleep(2 * time.Second)
    cancel() // Stops the request after 2 seconds
}()

resp, err := client.Synchronization().SincronizarActividades(ctx, config, req)
if errors.Is(err, context.Canceled) {
    log.Println("Request canceled by the user")
}
```

---

### 5. **Avoid long timeouts for real-time operations**

Bolivian taxation rules (SIAT) often require "Online" (En Línea) responses. 
- For **NIT verification**, we recommend a timeout of **5-10 seconds**.
- For **Invoice Reception**, we recommend **15-30 seconds**.
- For **Bulk Validation**, a longer timeout (**60-120 seconds**) may be required.

---

## 🛠️ Tracing and Monitoring

If you use tracing tools like **OpenTelemetry** or **Elastic APM**, the context passed to the SDK will automatically propagate the tracing headers if your `http.Client` is instrumented.

```go
// Example with instrumented client
httpClient := &http.Client{
    Transport: otelhttp.NewTransport(http.DefaultTransport),
}
s, _ := siat.New(siatURL, httpClient)

// The trace context will be sent to SIAT (if they support it) 
// and recorded in your monitoring system.
resp, err := s.Codes().VerificarNit(ctx, config, req)
```

## Summary Checklist

- [ ] Does every call to the SDK have a `context.Context`?
- [ ] Did you use `WithTimeout` or `WithDeadline`?
- [ ] Did you call `defer cancel()`?
- [ ] Do you handle `context.DeadlineExceeded` and `context.Canceled` errors separately for better logging?
- [ ] Is the timeout appropriate for the type of operation (Codes vs. Sync vs. Bulk)?
