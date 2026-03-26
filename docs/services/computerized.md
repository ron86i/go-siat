# Computerized Invoicing Service

The `Computerized` service is for modalities that DO NOT require XML signatures (e.g., Online Computerized / Hash-based).

## Sending Workflow

Computarizada uses a standard XML but requires a **Hash (SHA-256)** and **Base64** encoding for transmission.

```go
req := models.Computarizada().NewRecepcionFacturaBuilder().
    WithCodigoAmbiente(1).
    WithNit(nit).
    WithArchivo(base64Invoice).
    WithHashArchivo(hashSHA256).
    Build()

resp, err := client.Computarizada().RecepcionFactura(ctx, config, req)
```

## Difference from Electronic

- No `<ds:Signature>` node.
- No requirement for a digital certificate (`.p12`).
- Faster processing time due to simpler XML verification.
