# Servicio de Códigos

El servicio de `Codigos` gestiona la obtención de los identificadores únicos necesarios para operar con el SIAT, como el CUIS y el CUFD, además de realizar la verificación de NITs.

## Funcionalidades

### 1. Solicitar CUIS (Código Único de Inicio de Sistemas)
Este código es necesario para que el SIAT reconozca tu sistema en un punto de venta.

```go
req := models.Codigos().NewCuisBuilder().
    WithCodigoAmbiente(1).
    WithCodigoPuntoVenta(0).
    WithNit(12345678).
    Build()

resp, err := client.Codigos().SolicitudCuis(ctx, config, req)
```

### 2. Solicitar CUFD (Código Único de Facturación por Día)
Es el código necesario para firmar y emitir facturas cada día.

```go
req := models.Codigos().NewCufdBuilder().
    WithCuis(cuis).
    WithNit(12345678).
    Build()

resp, err := client.Codigos().SolicitudCufd(ctx, config, req)
```

### 3. Verificar NIT
Consulta si un NIT es válido para la emisión de facturas ante Impuestos Nacionales.

```go
req := models.Codigos().NewVerificarNitBuilder().
    WithCuis(cuis).
    WithNitParaVerificacion(9876543).
    Build()

resp, err := client.Codigos().VerificarNit(ctx, config, req)
```
