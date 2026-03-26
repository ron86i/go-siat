# Electronic Invoicing Service

The `Electronic` service is for modalities that require a **Digital Signature (XMLDSig)** using a `.p12` certificate.

## Reception

The reception involves wrapping the signed XML in a TAR package (optional Base64 encoding).

### Sending Electronic Invoice

```go
resp, err := client.Electronica().RecepcionFactura(ctx, config, req)

if resp.RespuestaServicioFacturacion.CodigoEstado == 908 {
    // Validated by SIAT as Correct
}
```

## Communication Verification

Verify the connection to the electronic invoicing service.

```go
resp, err := client.Electronica().VerificarComunicacion(ctx, cfg, req)
```
