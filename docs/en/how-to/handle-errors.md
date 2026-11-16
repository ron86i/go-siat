# How to handle errors

<p align="right">
  <a href="../../es/how-to/manejo-errores.md">🇪🇸 Español</a> · <a href="../README.md">Docs index</a>
</p>

Every SIAT call can fail in two independent ways, and the SDK reports them separately. This guide covers how to tell them apart, how to classify a SIAT rejection, and when it makes sense to retry.

---

## The rule: check twice

```go
resp, err := s.Electronica().RecepcionFactura(ctx, req)
if err != nil {
    // 1. Transport: no answer, or the answer was unparseable.
    return err
}
if err := siat.Verify(resp.Body.Content.RespuestaServicioFacturacion); err != nil {
    // 2. SIAT answered — and said no.
    return err
}
```

A `nil` error from the call means the round trip succeeded. It says nothing about whether SIAT accepted your document. Skipping the second check is the most common integration bug against this API: invoices appear to send fine and never exist on SIAT's side.

`siat.Verify` accepts any response payload that carries a `Transaccion` flag and a message list. It returns an error when `Transaccion` is false **or** when any message is not a warning — so an accepted-with-warnings response passes.

---

## What you get back

Both checks return `*siat.SiatError`, so one type handles everything:

```go
var siatErr *siat.SiatError
if errors.As(err, &siatErr) {
    siatErr.Code           // string  — "NETWORK_ERROR", "SIAT_SERVER_ERROR", "AUTH_FAILED", "TIMEOUT"
    siatErr.SiatCode       // int     — the SIAT code, 0 for transport errors
    siatErr.Message        // string  — joined "[code] description" messages
    siatErr.IsNetworkError // bool
    siatErr.IsRetryable    // bool
    siatErr.Mensajes       // []MensajeServicio — every message SIAT returned
    siatErr.StatusCode     // int     — HTTP status, when applicable
}
```

`SiatError` implements `Unwrap`, so `errors.Is` reaches the underlying transport error:

```go
if errors.Is(err, context.DeadlineExceeded) {
    // your ctx timeout fired
}
```

Two package-level helpers cover the common branches without unwrapping:

```go
siat.IsRetryable(err)     // worth trying again
siat.IsNetworkError(err)  // connectivity, not rejection
```

---

## Classify a SIAT code

`SiatCode` is the actionable field. The SDK ships the full code table plus four classifiers, so you can branch on category instead of memorizing numbers:

| Helper | True for | What it means for you |
| :--- | :--- | :--- |
| `siat.IsWarningCode(c)` | 2000–2019, 3008 | Accepted. Log it, keep going. |
| `siat.IsConfigCode(c)` | 910, 911, 912, 917, 958, 959, 975, 989 | Your credentials or setup are wrong. Never retry. |
| `siat.IsValidationCode(c)` | 910–966, 968–985, 996–998, 1000–1061 | The document is wrong. Fix data, then resend. |
| `siat.IsRetryableCode(c)` | 123, 967, 991, 995, 999 | SIAT-side transient. Back off and retry. |

`siat.GetMensaje(code)` returns the official Spanish description for any code, including ones not listed above.

```go
var siatErr *siat.SiatError
if errors.As(err, &siatErr) {
    switch {
    case siat.IsConfigCode(siatErr.SiatCode):
        log.Fatal("setup problem, retrying will not help: ", siatErr)
    case siat.IsRetryableCode(siatErr.SiatCode):
        return retryLater(req)
    case siat.IsValidationCode(siatErr.SiatCode):
        return fmt.Errorf("invoice rejected: %w", err)
    }
}
```

### Inspecting individual messages

SIAT often returns several messages at once. `SiatCode` holds only the first non-warning one, so read the list when you need the full picture:

```go
if siatErr.HasCode(siat.CodeCufdNoVigente) {
    cufd = refreshCufd()
}

for _, w := range siatErr.GetWarnings() {
    log.Printf("warning %d: %s", w.Codigo, w.Descripcion)
}
```

`GetWarnings()` filters to codes 2000–2019 and 3008. Note the asymmetry: if a response contains *only* warnings, `Verify` returns `nil` and there is no error to inspect — read `MensajesList` off the response itself in that case.

---

## Codes you will actually hit

Named constants exist for every code (`siat.CodeCuisNoVigente` and so on). These are the ones worth handling explicitly:

### Setup and credentials

| Code | Constant | Fix |
| ---: | :--- | :--- |
| 989 | `CodeTokenInvalido` | `Config.Token` is wrong or expired |
| 912 | `CodeSistemaNoAsociadoAlContribuyente` | `CodigoSistema` not registered for this NIT |
| 911 | `CodeCodigoSistemaInvalido` | Malformed `CodigoSistema` |
| 910 | `CodeAmbienteInvalido` | `CodigoAmbiente` does not match your `BaseURL` |
| 958 | `CodeUsuarioNoAutorizado` | Not authorized for that service |

Code 910 usually means pointing a production `CodigoAmbiente` at the pilot URL or vice versa. The two must agree.

### Expiring codes

| Code | Constant | Fix |
| ---: | :--- | :--- |
| 3008 | `CodeWarnCuisExpira` | **Warning.** Request a new CUIS soon |
| 973 / 929 | `CodeCuisNoSeEncuentraVigente` / `CodeCuisNoVigente` | Request a new CUIS |
| 953 | `CodeCufdNoVigente` | CUFD expired — request today's |
| 123 | `CodeCufdFueraDeTolerancia` | Clock skew, or a stale CUFD. Retryable |

CUFD is valid for one day. A long-running process that caches one at startup will start failing with 953 after midnight — refresh on a schedule, not once.

### Signature

| Code | Constant | Fix |
| ---: | :--- | :--- |
| 921 | `CodeFirmadoXmlIncorrecto` | Document unsigned or signature malformed |
| 922 | `CodeFirmaXmlNoCorrespondeAlContribuyente` | Certificate issued for a different NIT |
| 927 | `CodeCertificadoFirmaInvalido` | Certificate invalid or expired |
| 928 | `CodeCertificadoRevocado` | Certificate revoked |
| 965 | `CodeContribuyenteSinFirmaVigente` | No valid signature registered with SIAT |

A 921 indicates the SIN rejected the signed XML. Set `CodigoModalidad` on the typed request before calling `PrepararFactura`, and pass an XMLSigner for electronic mode — see [Sign invoices](sign-invoices.md).

### Document content

| Code | Constant | Fix |
| ---: | :--- | :--- |
| 939 | `CodeFacturaNoCumpleXsd` | Missing or malformed required field |
| 1002 | `CodeCufInvalido` | CUF wrong — usually a `fechaEmision` mismatch |
| 1000 / 952 | `CodeCufYaExisteEnSin` / `CodeCufYaRegistradoEnSin` | This CUF was already sent |
| 969 | `CodeHashInvalido` | Hash does not match the file |
| 1013 / 1024 | `CodeCalculoMontoTotalErroneo` / `CodeSumatoriaDetallesErronea` | Totals do not add up |
| 931 / 932 / 940 | Sector codes | See [Sectors](../explanation/sectors.md) |

For 1002, the usual cause is generating the CUF with one timestamp and putting a different one in `fechaEmision`. Compute the time once and reuse the variable.

### Transient

| Code | Constant |
| ---: | :--- |
| 967 | `CodeTiempoEsperaAgotadoDB` |
| 991 | `CodeErrorBaseDeDatos` |
| 995 | `CodeServicioNoDisponible` |
| 999 | `CodeErrorEjecucionServicio` |

---

## Retry sensibly

Retry only what `IsRetryable` marks. Network errors and timeouts qualify; a rejected document does not — resending identical bad data produces an identical rejection.

```go
func withRetry(ctx context.Context, attempts int, fn func() error) error {
    var err error
    for i := range attempts {
        if err = fn(); err == nil {
            return nil
        }
        if !siat.IsRetryable(err) {
            return err
        }
        select {
        case <-time.After(time.Second << i):   // exponential backoff
        case <-ctx.Done():
            return ctx.Err()
        }
    }
    return err
}
```

Two cautions specific to this API:

**Retrying a send is not idempotent.** If the first attempt reached SIAT and only the response was lost, resending the same CUF returns 1000 (`CufYaExisteEnSin`). Treat that code as "it actually went through" rather than as a failure.

**Always bound the call with a context deadline.** SIAT can be slow under load, and without a deadline a stalled request blocks indefinitely:

```go
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()
```

---

## SOAP faults

A malformed envelope comes back as a SOAP fault rather than a normal rejection. It surfaces on the response, not in `err`:

```go
if resp.Body.Fault != nil {
    log.Fatal("SOAP fault: ", resp.Body.Fault.String())
}
```

A fault means the request never reached SIAT's business logic — usually an endpoint mismatch or a corrupted payload. Check it before `Verify` when you are debugging a new integration.

---

## Related

| | |
| :--- | :--- |
| Signature failures in depth | [How-to: Sign invoices](sign-invoices.md) |
| Sector rejection codes | [Explanation: Sectors](../explanation/sectors.md) |
| Timeouts and HTTP tuning | [Reference: Configuration](../reference/configuration.md) |
| Complete error API | [Reference: API](../reference/api.md) |
