# Synchronization Service

The `Synchronization` service provides access to master catalogs essential for validating invoice fields.

## Functionalities

- **Economic Activities**: List of authorized activities.
- **Parametrics**: Standard units of measurement, document types, countries.
- **Sectors**: List of authorized sectoral documents.
- **Legends**: Mandatory texts for invoice footers.

## Example Sync Catalog

```go
req := models.Sincronizacion().NewSincronizarListaActividadesBuilder().
    WithCodigoAmbiente(1).
    WithCodigoPuntoVenta(0).
    WithCuis(cuis).
    WithNit(nit).
    Build()

resp, err := client.Sincronizacion().SincronizarListaActividades(ctx, config, req)
```
