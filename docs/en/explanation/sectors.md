# Understanding invoice sectors

<p align="right">
  <a href="../../es/explanation/sectores.md">🇪🇸 Español</a> · <a href="../README.md">Docs index</a>
</p>

SIAT does not have one invoice format. Its regulatory catalog defines **51 document sector codes**, each with its own XSD schema, its own mandatory fields, and its own validation rules. This page explains what a sector is, how the SDK models them, and how to find the one you need.

---

## What a sector is

A *documento sector* is SIAT's way of saying "this kind of business issues this shape of invoice". A hotel invoice needs the guest's stay dates. A mining export invoice needs mineral concentrations and international quotations. A telecom invoice needs a service period. Rather than one giant optional-everything document, SIAT publishes a separate schema per activity.

The sector code appears in two places, and they must agree:

| Where | Field | Who sets it |
| :--- | :--- | :--- |
| Inside the invoice document | `cabecera.codigoDocumentoSector` | The sector builder, with a correct default |
| In the submission request | `WithCodigoDocumentoSector(n)` | You, explicitly |

A mismatch is rejected with code `932 — CODIGO DOCUMENTO SECTOR NO CORRESPONDE AL SERVICIO`. Code `931` means the value itself is not a valid sector, and `940` means your NIT is not authorized for that sector.

> The SDK ships one builder per covered sector, each pre-filled with the right code. You should almost never call `WithCodigoDocumentoSector` on a *header* builder — the default is already correct. You do have to pass the code to the *request* builder.

---

## Every sector builder shares the same shape

All 51 builders in `pkg/models/invoices` (50 codes, one of which has two layouts) follow one pattern, so learning one teaches you all of them:

```go
cabecera := invoices.NewHotelCabeceraBuilder().
    WithNitEmisor(123456789).
    WithNumeroFactura(1).
    WithCuf(cuf).
    // ... sector-specific fields
    Build()

detalle := invoices.NewHotelDetalleBuilder().
    // ... sector-specific line fields
    Build()

factura := invoices.NewHotelBuilder().
    WithModalidad(siat.ModalidadElectronica).
    WithCabecera(cabecera).
    AddDetalle(detalle).
    Build()
```

Three constructors per sector: `New<Sector>CabeceraBuilder()` for the header, `New<Sector>DetalleBuilder()` for a line item, `New<Sector>Builder()` to assemble them. `AddDetalle` is additive — call it once per line.

`WithModalidad` on the root builder controls the XML namespace and whether a signature slot is emitted. It is separate from the `WithCodigoModalidad` you set on the *submission* request, and both need to be set.

What changes between sectors is only the field list on the header and detail builders. Your editor's autocomplete on the returned builder type is the fastest reference there is.

### Free-form JSON fields

Some sector fields accept free-form JSON, including the hotel
`WithDetalleHuespedes` field and cost, package, or other-data fields in export
headers. Pass either an already prepared JSON `string` or a Go value that
`encoding/json` can marshal. If the value cannot be serialized, the builder
keeps the error rather than panicking:

```go
detailBuilder := invoices.NewHotelDetalleBuilder().
    WithDetalleHuespedes(map[string]any{"name": "Ana"})
if err := detailBuilder.Err(); err != nil {
    return err
}
detail := detailBuilder.Build()
```

`Err()` is available on `HotelDetalle` and on commercial-export,
hydrocarbon-export, mining-export, and sale-price-export header builders.

---

## How a sector picks its service endpoint

SIAT exposes twelve SOAP endpoints. Which one accepts your invoice is not a free choice — it is determined by the sector.

**Six sectors have a dedicated endpoint.** They will not be accepted anywhere else:

| Facade | Sectors it serves |
| :--- | :--- |
| `s.CompraVenta()` | CompraVenta (1), CompraVentaBonificaciones (35), CompraVentaTasas (41) |
| `s.Telecomunicaciones()` | Telecomunicaciones (22), TelecomunicacionesZF (49) |
| `s.ServicioBasico()` | ServicioBasico (13), ServicioBasicoZF (40) |
| `s.EntidadFinanciera()` | EntidadFinanciera (15) |
| `s.BoletoAereo()` | BoletoAereo (30) |
| `s.DocumentoAjuste()` | NotaCreditoDebito (24), NotaFiscalCreditoDebito (24), NotaConciliacion (29), NotaCreditoDebitoDescuento (47), NotaCreditoDebitoIce (48) |

**The other 37 sectors route by modality**, not by activity. Pick the facade that matches how you are invoicing:

```go
s.Electronica().RecepcionFactura(ctx, req)     // ModalidadElectronica — signed
s.Computarizada().RecepcionFactura(ctx, req)   // ModalidadComputarizada — control code, no signature
```

Both expose the same nine-method interface, so switching modality is a one-word change at the call site. See [Architecture](architecture.md) for why the interfaces are shaped that way.

`DocumentoAjuste()` is the exception to the whole model: adjustment notes are not invoices, so its methods are named `RecepcionDocumentoAjuste`, `AnulacionDocumentoAjuste` and `ReversionAnulacionDocumentoAjuste` rather than the `*Factura` names.

---

## The full sector catalog

SIAT's regulatory catalog defines **51 document sector codes**. The SDK ships builders for **50 of them**; the single gap is flagged below.

Builder names drop the `New`/`Builder` wrapper — `CompraVenta` means `invoices.NewCompraVentaBuilder()`. The **Modalities** column is what `WithModalidad` on the root builder accepts for that sector; the **Facade** column is which service method the *submission* goes through — for sectors marked "`Electronica()` or `Computarizada()`" there, you pick the facade to match the modality you built.

| Code | Description | Tax credit | Modalities | Builder | Facade |
| ---: | :--- | :--- | :--- | :--- | :--- |
| 1 | Sales and purchases | Yes | Electrónica & Computarizada | `CompraVenta` | `CompraVenta()` |
| 2 | Real estate rental | Yes | Electrónica & Computarizada | `AlquilerBienInmueble` | `Electronica()` or `Computarizada()` |
| 3 | Commercial export | No | Electrónica & Computarizada | `ComercialExportacion` | `Electronica()` or `Computarizada()` |
| 4 | Commercial export on free consignment | No | Electrónica & Computarizada | `LibreConsignacion` | `Electronica()` or `Computarizada()` |
| 5 | Free trade zone sales | No | Electrónica & Computarizada | `ZonaFranca` | `Electronica()` or `Computarizada()` |
| 6 | Tourism and lodging services | No | Electrónica & Computarizada | `TurismoHospedaje` | `Electronica()` or `Computarizada()` |
| 7 | Food security and supply | No | Electrónica & Computarizada | `SeguridadAlimentaria` | `Electronica()` or `Computarizada()` |
| 8 | Zero rate — books and international road freight | No | Electrónica & Computarizada | `TasaCero` | `Electronica()` or `Computarizada()` |
| 9 | Foreign currency exchange | No | Electrónica & Computarizada | `MonedaExtranjera` | `Electronica()` or `Computarizada()` |
| 10 | Duty free | No | Electrónica & Computarizada | `DuttyFree` | `Electronica()` or `Computarizada()` |
| 11 | Education sector | Yes | Electrónica & Computarizada | `SectorEducativo` | `Electronica()` or `Computarizada()` |
| 12 | Hydrocarbon retail | Yes | Electrónica & Computarizada | `ComercializacionHidro` | `Electronica()` or `Computarizada()` |
| 13 | Basic services | Yes | Electrónica & Computarizada | `ServicioBasico` | `ServicioBasico()` |
| 14 | ICE-liable products | Yes | Electrónica & Computarizada | `AlcanzadaIce` | `Electronica()` or `Computarizada()` |
| 15 | Financial institutions | Yes | Electrónica & Computarizada | `EntidadFinanciera` | `EntidadFinanciera()` |
| 16 | Hotels | Yes | Electrónica & Computarizada | `Hotel` | `Electronica()` or `Computarizada()` |
| 17 | Hospitals / clinics | Yes | Electrónica & Computarizada | `HospitalClinica` | `Electronica()` or `Computarizada()` |
| 18 | Games of chance | Yes | Electrónica & Computarizada | `JuegoAzar` | `Electronica()` or `Computarizada()` |
| 19 | Hydrocarbons subject to IEHD | Yes | Electrónica & Computarizada | `HidrocarburoAlcanzadaIehd` | `Electronica()` or `Computarizada()` |
| 20 | Mineral export | No | Electrónica & Computarizada | `ComercialExportacionMinera` | `Electronica()` or `Computarizada()` |
| 21 | Domestic mineral sales | Yes | Electrónica & Computarizada | `VentaMineral` | `Electronica()` or `Computarizada()` |
| 22 | Telecommunications | Yes | Electrónica & Computarizada | `Telecomunicaciones` | `Telecomunicaciones()` |
| 23 | Prevalued | Yes | Electrónica & Computarizada | `Prevalorada` | `Electronica()` or `Computarizada()` |
| 24 | Credit/debit note | Adjustment | Electrónica & Computarizada | `NotaCreditoDebito` · `NotaFiscalCreditoDebito` | `DocumentoAjuste()` |
| 28 | Commercial export of services | No | Electrónica & Computarizada | `ComercialExportacionServicio` | `Electronica()` or `Computarizada()` |
| 29 | Conciliation note | Adjustment | Electrónica & Computarizada | `NotaConciliacion` | `DocumentoAjuste()` |
| 30 | Air ticket | Equivalent doc | Electrónica & Computarizada | `BoletoAereo` | `BoletoAereo()` |
| 31 | Energy supply | Yes | Electrónica & Computarizada | `SuministroEnergia` | `Electronica()` or `Computarizada()` |
| **33** | **Zero rate VAT, Law 1613** | No | — | **— no builder** | — |
| 34 | Insurance | Yes | Electrónica & Computarizada | `Seguros` | `Electronica()` or `Computarizada()` |
| 35 | Sales with bonuses | Yes | Electrónica & Computarizada | `CompraVentaBonificaciones` | `CompraVenta()` |
| 36 | Prevalued without tax credit | No | Electrónica & Computarizada | `PrevaloradaSinDerechoCreditoFiscal` | `Electronica()` or `Computarizada()` |
| 37 | CNG retail | Yes | Electrónica & Computarizada | `ComercializacionGnv` | `Electronica()` or `Computarizada()` |
| 38 | Hydrocarbons not subject to IEHD | Yes | Electrónica & Computarizada | `HidrocarburoNoAlcanzadaIehd` | `Electronica()` or `Computarizada()` |
| 39 | Natural gas and LPG retail | Yes | Electrónica & Computarizada | `ComercializacionGnGlp` | `Electronica()` or `Computarizada()` |
| 40 | Basic services, free trade zone | No | Electrónica & Computarizada | `ServicioBasicoZF` | `ServicioBasico()` |
| 41 | Sales with non-creditable fees | Yes | Electrónica & Computarizada | `CompraVentaTasas` | `CompraVenta()` |
| 42 | Rental, free trade zone | No | Electrónica & Computarizada | `AlquilerZF` | `Electronica()` or `Computarizada()` |
| 43 | Hydrocarbon export | No | Electrónica & Computarizada | `ComercialExportacionHidro` | `Electronica()` or `Computarizada()` |
| 44 | Lubricant import and retail | Yes | Electrónica & Computarizada | `ImportacionComercializacionLubricantes` | `Electronica()` or `Computarizada()` |
| 45 | Commercial export at sale price | No | Electrónica & Computarizada | `ComercialExportacionPVenta` | `Electronica()` or `Computarizada()` |
| 46 | Education, free trade zone | No | Electrónica & Computarizada | `SectorEducativoZF` | `Electronica()` or `Computarizada()` |
| 47 | Credit/debit note with discount | Adjustment | Electrónica & Computarizada | `NotaCreditoDebitoDescuento` | `DocumentoAjuste()` |
| 48 | ICE credit/debit note | Adjustment | Electrónica & Computarizada | `NotaCreditoDebitoIce` | `DocumentoAjuste()` |
| 49 | Telecommunications, free trade zone | No | Electrónica & Computarizada | `TelecomunicacionesZF` | `Telecomunicaciones()` |
| 50 | Hospitals / clinics, free trade zone | No | Electrónica & Computarizada | `HospitalClinicaZF` | `Electronica()` or `Computarizada()` |
| 51 | Gas bottling plants | Yes | Electrónica & Computarizada | `Engarrafadoras` | `Electronica()` or `Computarizada()` |
| 52 | Mineral sales to the Central Bank | No | **Electrónica only** | `VentaMineralBCB` | `Electronica()` only |
| 53 | Lubricant import and retail, IEHD | Yes | Electrónica & Computarizada | `LubricantesIehd` | `Electronica()` or `Computarizada()` |
| 54 | Biodiesel / ecological diesel feedstock | No | Electrónica & Computarizada | `Biodiesel` | `Electronica()` or `Computarizada()` |
| 55 | Fuel retail | Yes | Electrónica & Computarizada | `VentaCombustibleSinSubvencion` | `Electronica()` or `Computarizada()` |

Five things worth reading off this table:

**The codes are not contiguous.** 25, 26, 27 and 32 are absent from SIAT's own catalog. Do not invent them or use them as filler.

**24 appears twice deliberately.** `NotaCreditoDebito` and `NotaFiscalCreditoDebito` are two document layouts sharing one sector code. Pick the one whose field list matches the note you are issuing.

**Only 52 is restricted to one modality.** Every other sector's builder accepts both `ModalidadElectronica` and `ModalidadComputarizada` — the choice is yours based on how you are invoicing. Regulation restricts mineral sales to the Central Bank to electronic modality specifically, so `VentaMineralBCB` goes through `Electronica()` and always requires a signature; passing `ModalidadComputarizada` to its builder produces a document SIAT will reject.

**Modalities and Facade answer different questions.** Modalities is what the invoice *document* is built as. Facade is which *service method* you call to submit it. For the fourteen dedicated-endpoint sectors these are decoupled — a `CompraVenta()` submission can carry either an electronic or a computarized document, because the facade is chosen by activity, not modality. For the other sectors the two collapse into one choice: build electronic, call `Electronica()`; build computarized, call `Computarizada()`.

**`ZF` means *zona franca*.** It is the free-trade-zone variant of a sector, with a different code and extra fields — and almost always without tax credit even where the base sector grants it (compare 13 with 40, 22 with 49, 17 with 50).

### The four hydrocarbon and lubricant sectors

These are the easiest to mix up, because their names overlap and the difference is regulatory rather than structural:

| Code | Builder | Applies to |
| ---: | :--- | :--- |
| 12 | `ComercializacionHidro` | Fuel retail — diesel, gasoline, automotive |
| 19 | `HidrocarburoAlcanzadaIehd` | Activities **liable** for the IEHD |
| 38 | `HidrocarburoNoAlcanzadaIehd` | Activities **exempt** from the IEHD |
| 44 / 53 | `ImportacionComercializacionLubricantes` / `LubricantesIehd` | Lubricant import and retail, without and with IEHD |

Sectors 19 and 38 share an identical layout except that 19 carries `montoIehd` in the header and `porcentajeIehd` per detail line. If you are exempt, 38 rejects those fields; if you are liable, 19 requires them.

### The one uncovered sector

Code 33 exists in SIAT's catalog but has no builder in the SDK:

| Code | Description |
| ---: | :--- |
| 33 | Zero rate VAT, Law 1613 — capital goods and industrial plants |

If your activity falls under it, there is no shortcut inside the SDK: you must build the document struct yourself and feed it in with `WithArchivo` / `WithHashArchivo` instead of `WithFactura`. See [packaging by hand](../how-to/send-invoices.md#packaging-by-hand).

Do not substitute sector 8 (`TasaCero`) for 33. Both are zero-rate, but 8 covers books and international road freight under a different legal basis, and SIAT rejects the mismatch with code 932.

---

## Finding the fields for your sector

The SDK's own tests are the most reliable examples, because they are the code that actually round-trips against SIAT. Every sector has one:

```
pkg/models/invoices/<sector>_test.go
```

For example, `pkg/models/invoices/hotel_test.go` builds a complete hotel invoice with every required field populated. Copy that, change the values, and you have a working document.

Sectors whose codes you do not recognize are best confirmed against SIAT itself rather than guessed — `Sincronizacion()` exposes `SincronizarParametricaTipoDocumentoSector` for the full catalog and `SincronizarListaActividadesDocumentoSector` for which sectors your registered economic activities allow.

---

## Related

| | |
| :--- | :--- |
| Build and send your first invoice | [Tutorial](../tutorial/first-invoice.md) |
| Why the services are shaped this way | [Explanation: Architecture](architecture.md) |
| Signing electronic-modality invoices | [How-to: Sign invoices](../how-to/sign-invoices.md) |
| Sector rejection codes (931/932/940) | [How-to: Handle errors](../how-to/handle-errors.md) |
