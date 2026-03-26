# Servicio de Ventas (Sectoriales)

El servicio de `Ventas` es el componente más utilizado del SDK `go-siat`, permitiendo emitir facturas para más de 35 sectores regulatorios.

## Modalidades Soportadas

1. **Electrónica en Línea**: Requiere firma digital (XMLDSig).
2. **Computarizada en Línea**: Utiliza Hash/Código de Control.

## Builders Sectoriales

Cada factura sectorial tiene dos partes: **Cabecera** y **Detalle**.

### Ejemplo: Venta Compra-Venta (Sector 1)

```go
cabecera := invoices.NewCompraVentaCabeceraBuilder().
    WithCodigoPuntoVenta(0).
    WithCodigoSucursal(0).
    WithCodigoMetodoPago(1).
    WithNitEmisor(12345).
    // ... otros campos
    Build()

detalle := invoices.NewCompraVentaDetalleBuilder().
    WithActividadEconomica("551010").
    WithCodigoProducto("1234").
    WithDescripcion("Descripción del item").
    Build()

factura := invoices.NewCompraVentaBuilder().
    WithCabecera(cabecera).
    AddDetalle(detalle).
    Build()
```

## Flujo de Recepción

Una vez construido el objeto `factura`, se envía mediante el servicio correspondiente:

```go
// Para Electrónica
resp, err := client.Electronica().RecepcionFactura(ctx, config, req)

// Para Computarizada
resp, err := client.Computarizada().RecepcionFactura(ctx, config, req)
```
