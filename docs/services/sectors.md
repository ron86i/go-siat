# Sectoral Invoices (Builders)

The `go-siat` SDK provides **Builders** for more than 35 regulatory sectors defined by the SIAT. A builder allows you to construct a type-safe, validated XML structure without manually handling XML namespaces or SOAP structures.

## Complete Sector List

| Sector ID | Sector Name | Main Builder |
| :--- | :--- | :--- |
| **1** | Compra-Venta | `NewCompraVentaBuilder` |
| **2** | Alquiler de Bienes Inmuebles | `NewAlquilerBienInmuebleBuilder` |
| **8** | Comercial de Exportación | `NewComercialExportacionBuilder` |
| **12** | Comercialización de Hidrocarburos | **`NewComercializacionHidroBuilder`** |
| **16** | Hoteles | **`NewHotelBuilder`** |
| **20** | Comercial de Exportación Minera | `NewComercialExportacionMineraBuilder` |
| **28** | Comercial de Exportación de Servicios | `NewComercialExportacionServicioBuilder` |
| **31** | Suministro de Energía Eléctrica | **`NewSuministroEnergiaBuilder`** |
| **34** | Hospitales y Clínicas | `NewHospitalClinicaBuilder` |
| **37** | Comercialización de GNV | `NewComercializacionGnvBuilder` |
| **43** | Comercial de Exportación de Hidrocarburos | **`NewComercialExportacionHidrocarburosBuilder`** |
| **44** | Importación y Comercialización de Lubricantes | `NewImportacionComercializacionLubricantesBuilder` |
| **50** | Hospital Clínico Zona Franca | **`NewHospitalClinicaZFBuilder`** |
| **53** | Comercialización de Lubricantes IEHD | `NewLubricantesIEHDBuilder` |

## Anatomy of a Sectoral Invoice

Each sectoral invoice is composed of:
1.  **Header (Cabecera)**: Common fields across all modalities.
2.  **Details (Detalle)**: List of items/products.
3.  **Specific Attributes**: Fields like `placaVehiculo`, `fechaIngresoHospedaje`, etc.

```go
// Example Sector 16: Hotel
inv := invoices.NewHotelBuilder().
    WithCabecera(hotelHeader).
    AddDetalle(hotelDetail).
    Build()
```
