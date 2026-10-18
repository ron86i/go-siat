# Fiscal documents in v3

All 47 document families in the supplied ZIP are covered: 93 XML roots and 48 public models. Both credit/debit note models target the same roots and follow the same updated schema.

Each name links to its implementation. Use the matching `New…CabeceraBuilder` and `New…DetalleBuilder` for individual sections. Boleto Aereo has only a header; Prevalorada uses `WithDetalle`; Nota Conciliacion retains `AddDetalleOriginal` and `AddDetalleConciliacion`.

| Sector | Public builder | Modalities |
| --- | --- | --- |
| 1 | [`NewCompraVentaBuilder`](../../../pkg/models/invoices/compra_venta.go) | Electronic and computerized |
| 2 | [`NewAlquilerBienInmuebleBuilder`](../../../pkg/models/invoices/alquiler_bien_inmueble.go) | Electronic and computerized |
| 3 | [`NewComercialExportacionBuilder`](../../../pkg/models/invoices/comercial_exportacion.go) | Electronic and computerized |
| 4 | [`NewLibreConsignacionBuilder`](../../../pkg/models/invoices/libre_consignacion.go) | Electronic and computerized |
| 5 | [`NewZonaFrancaBuilder`](../../../pkg/models/invoices/zona_franca.go) | Electronic and computerized |
| 6 | [`NewTurismoHospedajeBuilder`](../../../pkg/models/invoices/turismo_hospedaje.go) | Electronic and computerized |
| 7 | [`NewSeguridadAlimentariaBuilder`](../../../pkg/models/invoices/seguridad_alimentaria.go) | Electronic and computerized |
| 8 | [`NewTasaCeroBuilder`](../../../pkg/models/invoices/tasa_cero.go) | Electronic and computerized |
| 9 | [`NewMonedaExtranjeraBuilder`](../../../pkg/models/invoices/moneda_extranjera.go) | Electronic and computerized |
| 10 | [`NewDuttyFreeBuilder`](../../../pkg/models/invoices/dutty_free.go) | Electronic and computerized |
| 11 | [`NewSectorEducativoBuilder`](../../../pkg/models/invoices/sector_educativo.go) | Electronic and computerized |
| 13 | [`NewServicioBasicoBuilder`](../../../pkg/models/invoices/servicio_basico.go) | Electronic and computerized |
| 14 | [`NewAlcanzadaIceBuilder`](../../../pkg/models/invoices/alcanzada_ice.go) | Electronic and computerized |
| 15 | [`NewEntidadFinancieraBuilder`](../../../pkg/models/invoices/entidad_financiera.go) | Electronic and computerized |
| 16 | [`NewHotelBuilder`](../../../pkg/models/invoices/hotel.go) | Electronic and computerized |
| 17 | [`NewHospitalClinicaBuilder`](../../../pkg/models/invoices/hospital_clinica.go) | Electronic and computerized |
| 18 | [`NewJuegoAzarBuilder`](../../../pkg/models/invoices/juego_azar.go) | Electronic and computerized |
| 19 | [`NewHidrocarburoAlcanzadaIehdBuilder`](../../../pkg/models/invoices/hidrocarburo_alcanzada_iehd.go) | Electronic and computerized |
| 20 | [`NewComercialExportacionMineraBuilder`](../../../pkg/models/invoices/comercial_exportacion_minera.go) | Electronic and computerized |
| 21 | [`NewVentaMineralBuilder`](../../../pkg/models/invoices/venta_mineral.go) | Electronic and computerized |
| 22 | [`NewTelecomunicacionesBuilder`](../../../pkg/models/invoices/telecomunicaciones.go) | Electronic and computerized |
| 23 | [`NewPrevaloradaBuilder`](../../../pkg/models/invoices/prevalorada.go) | Electronic and computerized |
| 24 | [`NewNotaCreditoDebitoBuilder`](../../../pkg/models/invoices/nota_credito_debito.go) / [`NewNotaFiscalCreditoDebitoBuilder`](../../../pkg/models/invoices/nota_fiscal_credito_debito.go) | Electronic and computerized |
| 28 | [`NewComercialExportacionServicioBuilder`](../../../pkg/models/invoices/comercial_exportacion_servicio.go) | Electronic and computerized |
| 29 | [`NewNotaConciliacionBuilder`](../../../pkg/models/invoices/nota_conciliacion.go) | Electronic and computerized |
| 30 | [`NewBoletoAereoBuilder`](../../../pkg/models/invoices/boleto_aereo.go) | Electronic and computerized |
| 31 | [`NewSuministroEnergiaBuilder`](../../../pkg/models/invoices/suministro_energia.go) | Electronic and computerized |
| 34 | [`NewSegurosBuilder`](../../../pkg/models/invoices/seguros.go) | Electronic and computerized |
| 35 | [`NewCompraVentaBonificacionesBuilder`](../../../pkg/models/invoices/compra_venta_bonificaciones.go) | Electronic and computerized |
| 36 | [`NewPrevaloradaSinDerechoCreditoFiscalBuilder`](../../../pkg/models/invoices/prevalorada_sin_derecho_credito_fiscal.go) | Electronic and computerized |
| 37 | [`NewComercializacionGnvBuilder`](../../../pkg/models/invoices/comercializacion_gnv.go) | Electronic and computerized |
| 38 | [`NewHidrocarburoNoAlcanzadaIehdBuilder`](../../../pkg/models/invoices/hidrocarburo_no_alcanzada_iehd.go) | Electronic and computerized |
| 39 | [`NewComercializacionGnGlpBuilder`](../../../pkg/models/invoices/comercializacion_gn_glp.go) | Electronic and computerized |
| 40 | [`NewServicioBasicoZFBuilder`](../../../pkg/models/invoices/servicio_basico_zf.go) | Electronic and computerized |
| 41 | [`NewVueloCharterBuilder`](../../../pkg/models/invoices/vuelo_charter.go) | Electronic and computerized |
| 42 | [`NewAlquilerZFBuilder`](../../../pkg/models/invoices/alquiler_zona_franca.go) | Electronic and computerized |
| 43 | [`NewComercialExportacionHidroBuilder`](../../../pkg/models/invoices/comercial_exportacion_hidrocarburos.go) | Electronic and computerized |
| 44 | [`NewImportacionComercializacionLubricantesBuilder`](../../../pkg/models/invoices/importacion_comercializacion_lubricantes.go) | Electronic and computerized |
| 45 | [`NewComercialExportacionPVentaBuilder`](../../../pkg/models/invoices/comercial_exportacion_punto_venta.go) | Electronic and computerized |
| 46 | [`NewSectorEducativoZFBuilder`](../../../pkg/models/invoices/sector_educativo_zona_franca.go) | Electronic and computerized |
| 48 | [`NewNotaCreditoDebitoIceBuilder`](../../../pkg/models/invoices/nota_credito_debito_ice.go) | Electronic and computerized |
| 49 | [`NewTelecomunicacionesZFBuilder`](../../../pkg/models/invoices/telecomunicaciones_zf.go) | Electronic and computerized |
| 50 | [`NewHospitalClinicaZFBuilder`](../../../pkg/models/invoices/hospital_clinica_zona_franca.go) | Electronic and computerized |
| 52 | [`NewVentaMineralBCBBuilder`](../../../pkg/models/invoices/venta_mineral_bcb.go) | Electronic only |
| 53 | [`NewLubricantesIehdBuilder`](../../../pkg/models/invoices/lubricantes_iehd.go) | Electronic and computerized |
| 54 | [`NewBiodieselBuilder`](../../../pkg/models/invoices/biodiesel.go) | Electronic and computerized |
| 55 | [`NewVentaCombustibleSinSubvencionBuilder`](../../../pkg/models/invoices/venta_combustible_sin_subvencion.go) | Electronic and computerized |

## Verification

The expectations in `pkg/models/invoices/xsd_v3_catalog_test.go` cover all 93 variants without loading schemas, official XML samples, credentials or network access. They check complete sequences, names, setter types, values, nulls and decimal precision.

```sh
SIAT_RUN=0 go test ./pkg/models/invoices ./internal/core/domain/documents -count=1
```

Local validation against the original XSDs passed for all 46 computerized formats with values and with nulls: 92 documents. Full validation of the 47 electronic formats requires signed XML and `SignatureSchema.xsd`, which is missing from the ZIP. This refactor did not test real SIAT acceptance.

[← Migration to v3](../how-to/siat-v3.md)
