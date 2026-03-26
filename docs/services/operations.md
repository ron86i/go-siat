# Operations Service

The `Operations` service allows the management of business locations (Point of Sale - PV) and significant events within the SIAT ecosystem.

## Point of Sale (PV) Management

- **Registration**: Create a new PV for a branch.
- **Closure**: Close an active PV.
- **Consultation**: Retrieve all registered and active PVs for a branch.

```go
req := models.Operaciones().NewConsultaPuntoVentaBuilder().
    WithCodigoAmbiente(1).
    WithCodigoSistema("ABC").
    WithCodigoSucursal(0).
    WithCuis(cuis).
    WithNit(nit).
    Build()

resp, err := client.Operaciones().ConsultaPuntoVenta(ctx, config, req)
```

## Significant Events

Registers events where online invoicing is not possible (e.g., loss of connectivity, power outage) to justify delayed transmission.

```go
event := models.Operaciones().NewRegistroEventoSignificativoBuilder().
    WithCodigoAmbiente(1).
    WithCodigoEvento(1). // 1: Power failure
    WithCodigoPuntoVenta(0).
    WithCodigoSucursal(0).
    WithCuis(cuis).
    WithNit(nit).
    WithDescripcion("Total blackout at branch").
    WithFechaInicioEvento("2026-03-26T10:00:00.000").
    WithFechaFinEvento("2026-03-26T12:00:00.000").
    Build()

resp, err := client.Operaciones().RegistroEventoSignificativo(ctx, config, event)
```
