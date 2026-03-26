# Sales Service (Sectoral)

The `Sales` service is the most used component of the SIAT SDK, allowing the issuance of invoices for more than 35 different sectors.

## Modalidades

1. **Electronic Online**: Uses XML signatures (XMLDSig).
2. **Computerized Online**: Uses Hash/Control Code.

## Sectoral Builders

Every sectoral invoice has two parts: **Cabecera** (Header) and **Detalle** (Details).

### Example: Basic Sale (Sector 1)

```go
cabecera := invoices.NewCompraVentaCabeceraBuilder().
    WithCodigoPuntoVenta(0).
    WithCodigoSucursal(0).
    WithCodigoMetodoPago(1).
    WithNitEmisor(12345).
    // ... other fields
    Build()

detalle := invoices.NewCompraVentaDetalleBuilder().
    WithActividadEconomica("551010").
    WithCodigoProducto("1234").
    WithDescripcion("Item description").
    WithCantidad(1.0).
    WithPrecioUnitario(100.0).
    WithSubTotal(100.0).
    Build()

factura := invoices.NewCompraVentaBuilder().
    WithCabecera(cabecera).
    AddDetalle(detalle).
    Build()
```

## Reception Workflow

Once the `factura` object is built, send it via the corresponding service:

```go
// For Electronic
resp, err := client.Electronica().RecepcionFactura(ctx, config, req)

// For Computerized
resp, err := client.Computarizada().RecepcionFactura(ctx, config, req)
```
