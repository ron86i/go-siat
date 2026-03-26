# Codes Service

The `Codes` service handles the acquisition of unique IDs required to operate with SIAT, such as CUIS and CUFD, as well as NIT verification.

## Functions

### 1. Request CUIS
```go
req := models.Codigos().NewCuisBuilder().
    WithCodigoAmbiente(1).
    WithCodigoPuntoVenta(0).
    WithNit(12345678).
    Build()

resp, err := client.Codigos().SolicitudCuis(ctx, config, req)
```

### 2. Request CUFD
```go
req := models.Codigos().NewCufdBuilder().
    WithCuis(cuis).
    WithNit(12345678).
    Build()

resp, err := client.Codigos().SolicitudCufd(ctx, config, req)
```
