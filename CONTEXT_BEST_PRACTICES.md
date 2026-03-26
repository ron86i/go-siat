# Guía de Mejores Prácticas para Context en go-siat

## Introducción

El SDK go-siat utiliza `context.Context` para propagar plazos, señales de cancelación e información de solicitud a través de las llamadas a la API. Esta guía explica cómo usar correctamente los contextos para maximizar la confiabilidad de tu aplicación.

## Principios Fundamentales

### 1. **Nunca uses `context.Background()` directamente en métodos del SDK**

❌ **Incorrecto:**
```go
// NUNCA hagas esto
config := siat.Config{Token: "myToken"}
resp, err := client.Codigos().VerificarNit(context.Background(), config, nitRequest)
```

✅ **Correcto:**
```go
// Los llamadores proporcionan el contexto
ctx := context.Background()
if deadline, ok := time.ParseDuration("30s"); ok {
    ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
    defer cancel()
}
config := siat.WithConfig("myToken")
resp, err := client.Codigos().VerificarNit(ctx, config, nitRequest)
```

**Por qué:** `context.Background()` no tiene ningún plazo o señal de cancelación. El servidor SIAT puede tardar más que tolerado, y tu aplicación no podrá detener la solicitud.

---

### 2. **Establece siempre un plazo con `context.WithTimeout()` o `context.WithDeadline()`**

#### Opción A: Timeout para operaciones individuales
```go
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel() // IMPORTANTÍSIMO: siempre hace defer del cancel

config := client.WithConfig("myToken")
resp, err := client.Codigos().VerificarNit(ctx, config, nitRequest)
if err != nil {
    if errors.Is(err, context.DeadlineExceeded) {
        log.Println("Timeout: SIAT tardó más de 30 segundos")
    }
}
```

#### Opción B: Deadline absoluto
```go
deadline := time.Now().Add(45 * time.Second)
ctx, cancel := context.WithDeadline(context.Background(), deadline)
defer cancel()

// Usar ctx en las llamadas al SDK
```

#### Opción C: Heredar deadline de contexto padre
```go
func ProcessSiatRequest(ctx context.Context) error {
    // Si ctx tiene deadline, lo hereda automáticamente
    config := client.WithConfig("myToken")
    resp, err := client.Codigos().VerificarNit(ctx, config, nitRequest)
    return err
}
```

**Importante:** Siempre hacer `defer cancel()` después de `WithTimeout()` o `WithDeadline()` para liberar recursos.

---

### 3. **Propagar contexos desde handlers HTTP**

En una aplicación web, siempre propaga el contexto del handler HTTP:

```go
// En tu servidor HTTP
func handleFacturaRequest(w http.ResponseWriter, r *http.Request) {
    // r.Context() contiene deadline, cancelación y valores del cliente HTTP
    ctx := r.Context()
    
    config := siatClient.WithConfig(getUserToken(r))
    resp, err := siatClient.CompraVenta().RecepcionFactura(ctx, config, facturaData)
    
    if err != nil {
        if errors.Is(err, context.Canceled) {
            // Cliente desconectó
            return
        }
        if errors.Is(err, context.DeadlineExceeded) {
            http.Error(w, "Timeout en SIAT", http.StatusGatewayTimeout)
            return
        }
        http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/xml")
    // Usar GetContent() para extraer la respuesta XML de dentro del sobre SOAP
    if content, err := resp.GetContent(); err == nil {
        xml.NewEncoder(w).Encode(content)
    } else {
        xml.NewEncoder(w).Encode(resp.Body.Fault)
    }
}
```

**Ventaja:** Si el cliente desconecta o el navegador cancela la solicitud, automáticamente se cancela la solicitud al SIAT también.

---

### 4. **Usar `context.WithValue()` para Trace ID (Recomendado)**

Además del `WithTraceID()` del SDK, puedes propagar información de seguimiento:

```go
// Generar Trace ID único por solicitud
traceID := fmt.Sprintf("trace-%d-%s", time.Now().UnixMicro(), uuid.New().String()[:8])

// Crear contexto con Trace ID
ctx := context.WithValue(context.Background(), "traceID", traceID)

// Usar Trace ID en cliente SIAT
config := siatClient.WithTraceID(traceID).WithConfig("myToken")
resp, err := siatClient.Codigos().VerificarNit(ctx, config, nitRequest)

// Ahora todas las solicitudes SIAT incluyen X-Trace-ID: trace-1234567-abcd1234
// Puedes loguear con el mismo traceID en múltiples servicios
```

---

### 5. **Patrones para diferentes escenarios**

#### Escenario 1: Operación rápida esperada (≤5s)
```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
resp, err := client.Sincronizacion().SincronizarParametricaTipoMoneda(ctx, cfg, request)
```

#### Escenario 2: Operación lenta (factura con muchos items)
```go
ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
defer cancel()
resp, err := client.Electronica().RecepcionFactura(ctx, cfg, facturaGrande)
```

#### Escenario 3: Bucle de reintentos con timeout global
```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
defer cancel()

var resp *soap.EnvelopeResponse
var err error
for attempt := 1; attempt <= 3; attempt++ {
    resp, err = client.Operaciones().ConsultaPuntoVenta(ctx, cfg, request)
    if err == nil {
        break
    }
    if errors.Is(err, context.DeadlineExceeded) {
        break // No reintentar si ya pasó el deadline global
    }
    time.Sleep(time.Second * time.Duration(attempt))
}
```

#### Escenario 4: Middleware HTTP con timeout heredado
```go
func timeoutMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ctx, cancel := context.WithTimeout(r.Context(), 30*time.Second)
        defer cancel()
        
        // Pasar contexto con timeout al siguiente handler
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
```

---

## Errores Comunes

### ❌ Error 1: Olvidar `defer cancel()`
```go
// MAL: semáforo no liberado, goroutine leak
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// Falta: defer cancel()
resp, err := client.Codigos().VerificarNit(ctx, cfg, req)
```

### ❌ Error 2: Crear contexto en cada iteración
```go
// MAL: Si loop requiere 3 minutos totales, timeout independiente es insuficiente
for i := 0; i < 100; i++ {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    client.Sincronizacion().SincronizarActividades(ctx, cfg, req)
    cancel()
}

// BUENO: Timeout global para todo el loop
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
defer cancel()
for i := 0; i < 100; i++ {
    if err := client.Sincronizacion().SincronizarActividades(ctx, cfg, req); err != nil {
        if errors.Is(err, context.DeadlineExceeded) {
            break
        }
    }
}
```

### ❌ Error 3: No capturar `context.DeadlineExceeded`
```go
// MAL: No distingues timeout de otros errores
resp, err := client.Codigos().VerificarNit(ctx, cfg, req)
if err != nil {
    log.Fatal("Error:", err)
}

// BUENO: Maneja timeout diferente
resp, err := client.Codigos().VerificarNit(ctx, cfg, req)
if err != nil {
    if errors.Is(err, context.DeadlineExceeded) {
        log.Println("Timeout en SIAT")
    } else if errors.Is(err, context.Canceled) {
        log.Println("Solicitud cancelada")
    } else {
        log.Println("Error SIAT:", err)
    }
}
```

---

## Mejores Prácticas Resumidas

| Práctica | Razón |
|----------|-------|
| Siempre proporcionar `context.Context` | Permite cancelación y plazos |
| Usar `WithTimeout()` o `WithDeadline()` | Evita bloqueos indefinidos |
| Hacer `defer cancel()` | Libera recursos correctamente |
| Propagar contexto desde HTTP handlers | Cancela solicitudes SIAT si cliente desconecta |
| Usar Trace ID con distributed tracing | Correlaciona logs entre servicios |
| Capturar `DeadlineExceeded` específicamente | Diferencia timeouts de otros errores |
| Heredar timeouts en loops | Mantiene límite de tiempo global |

---

## Referencias

- [Go Context Documentation](https://golang.org/pkg/context/)
- [Using Context in Go](https://go.dev/blog/context)
- [Go Concurrency Patterns](https://go.dev/blog/pipelines)
- [HTTP Handler with Context](https://pkg.go.dev/net/http#Request.Context)

