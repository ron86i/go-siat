# Documentos fiscales v3

Los 47 contratos de familia del ZIP suministrado están cubiertos: 93 raíces XML, 48 modelos públicos y ningún documento del paquete sin modelo. Los dos modelos de nota fiscal crédito/débito comparten las mismas raíces y ambos están alineados.

Cada nombre enlaza con su implementación. Para construir la cabecera y el detalle se usan los correspondientes `New…CabeceraBuilder` y `New…DetalleBuilder`. Boleto Aéreo sólo tiene cabecera; Prevalorada usa `WithDetalle`; Nota de Conciliación conserva `AddDetalleOriginal` y `AddDetalleConciliacion`.

| Sector | Builder público | Modalidades |
| --- | --- | --- |
| 1 | [`NewCompraVentaBuilder`](../../../pkg/models/invoices/compra_venta.go) | Electrónica y computarizada |
| 2 | [`NewAlquilerBienInmuebleBuilder`](../../../pkg/models/invoices/alquiler_bien_inmueble.go) | Electrónica y computarizada |
| 3 | [`NewComercialExportacionBuilder`](../../../pkg/models/invoices/comercial_exportacion.go) | Electrónica y computarizada |
| 4 | [`NewLibreConsignacionBuilder`](../../../pkg/models/invoices/libre_consignacion.go) | Electrónica y computarizada |
| 5 | [`NewZonaFrancaBuilder`](../../../pkg/models/invoices/zona_franca.go) | Electrónica y computarizada |
| 6 | [`NewTurismoHospedajeBuilder`](../../../pkg/models/invoices/turismo_hospedaje.go) | Electrónica y computarizada |
| 7 | [`NewSeguridadAlimentariaBuilder`](../../../pkg/models/invoices/seguridad_alimentaria.go) | Electrónica y computarizada |
| 8 | [`NewTasaCeroBuilder`](../../../pkg/models/invoices/tasa_cero.go) | Electrónica y computarizada |
| 9 | [`NewMonedaExtranjeraBuilder`](../../../pkg/models/invoices/moneda_extranjera.go) | Electrónica y computarizada |
| 10 | [`NewDuttyFreeBuilder`](../../../pkg/models/invoices/dutty_free.go) | Electrónica y computarizada |
| 11 | [`NewSectorEducativoBuilder`](../../../pkg/models/invoices/sector_educativo.go) | Electrónica y computarizada |
| 13 | [`NewServicioBasicoBuilder`](../../../pkg/models/invoices/servicio_basico.go) | Electrónica y computarizada |
| 14 | [`NewAlcanzadaIceBuilder`](../../../pkg/models/invoices/alcanzada_ice.go) | Electrónica y computarizada |
| 15 | [`NewEntidadFinancieraBuilder`](../../../pkg/models/invoices/entidad_financiera.go) | Electrónica y computarizada |
| 16 | [`NewHotelBuilder`](../../../pkg/models/invoices/hotel.go) | Electrónica y computarizada |
| 17 | [`NewHospitalClinicaBuilder`](../../../pkg/models/invoices/hospital_clinica.go) | Electrónica y computarizada |
| 18 | [`NewJuegoAzarBuilder`](../../../pkg/models/invoices/juego_azar.go) | Electrónica y computarizada |
| 19 | [`NewHidrocarburoAlcanzadaIehdBuilder`](../../../pkg/models/invoices/hidrocarburo_alcanzada_iehd.go) | Electrónica y computarizada |
| 20 | [`NewComercialExportacionMineraBuilder`](../../../pkg/models/invoices/comercial_exportacion_minera.go) | Electrónica y computarizada |
| 21 | [`NewVentaMineralBuilder`](../../../pkg/models/invoices/venta_mineral.go) | Electrónica y computarizada |
| 22 | [`NewTelecomunicacionesBuilder`](../../../pkg/models/invoices/telecomunicaciones.go) | Electrónica y computarizada |
| 23 | [`NewPrevaloradaBuilder`](../../../pkg/models/invoices/prevalorada.go) | Electrónica y computarizada |
| 24 | [`NewNotaCreditoDebitoBuilder`](../../../pkg/models/invoices/nota_credito_debito.go) / [`NewNotaFiscalCreditoDebitoBuilder`](../../../pkg/models/invoices/nota_fiscal_credito_debito.go) | Electrónica y computarizada |
| 28 | [`NewComercialExportacionServicioBuilder`](../../../pkg/models/invoices/comercial_exportacion_servicio.go) | Electrónica y computarizada |
| 29 | [`NewNotaConciliacionBuilder`](../../../pkg/models/invoices/nota_conciliacion.go) | Electrónica y computarizada |
| 30 | [`NewBoletoAereoBuilder`](../../../pkg/models/invoices/boleto_aereo.go) | Electrónica y computarizada |
| 31 | [`NewSuministroEnergiaBuilder`](../../../pkg/models/invoices/suministro_energia.go) | Electrónica y computarizada |
| 34 | [`NewSegurosBuilder`](../../../pkg/models/invoices/seguros.go) | Electrónica y computarizada |
| 35 | [`NewCompraVentaBonificacionesBuilder`](../../../pkg/models/invoices/compra_venta_bonificaciones.go) | Electrónica y computarizada |
| 36 | [`NewPrevaloradaSinDerechoCreditoFiscalBuilder`](../../../pkg/models/invoices/prevalorada_sin_derecho_credito_fiscal.go) | Electrónica y computarizada |
| 37 | [`NewComercializacionGnvBuilder`](../../../pkg/models/invoices/comercializacion_gnv.go) | Electrónica y computarizada |
| 38 | [`NewHidrocarburoNoAlcanzadaIehdBuilder`](../../../pkg/models/invoices/hidrocarburo_no_alcanzada_iehd.go) | Electrónica y computarizada |
| 39 | [`NewComercializacionGnGlpBuilder`](../../../pkg/models/invoices/comercializacion_gn_glp.go) | Electrónica y computarizada |
| 40 | [`NewServicioBasicoZFBuilder`](../../../pkg/models/invoices/servicio_basico_zf.go) | Electrónica y computarizada |
| 41 | [`NewVueloCharterBuilder`](../../../pkg/models/invoices/vuelo_charter.go) | Electrónica y computarizada |
| 42 | [`NewAlquilerZFBuilder`](../../../pkg/models/invoices/alquiler_zona_franca.go) | Electrónica y computarizada |
| 43 | [`NewComercialExportacionHidroBuilder`](../../../pkg/models/invoices/comercial_exportacion_hidrocarburos.go) | Electrónica y computarizada |
| 44 | [`NewImportacionComercializacionLubricantesBuilder`](../../../pkg/models/invoices/importacion_comercializacion_lubricantes.go) | Electrónica y computarizada |
| 45 | [`NewComercialExportacionPVentaBuilder`](../../../pkg/models/invoices/comercial_exportacion_punto_venta.go) | Electrónica y computarizada |
| 46 | [`NewSectorEducativoZFBuilder`](../../../pkg/models/invoices/sector_educativo_zona_franca.go) | Electrónica y computarizada |
| 48 | [`NewNotaCreditoDebitoIceBuilder`](../../../pkg/models/invoices/nota_credito_debito_ice.go) | Electrónica y computarizada |
| 49 | [`NewTelecomunicacionesZFBuilder`](../../../pkg/models/invoices/telecomunicaciones_zf.go) | Electrónica y computarizada |
| 50 | [`NewHospitalClinicaZFBuilder`](../../../pkg/models/invoices/hospital_clinica_zona_franca.go) | Electrónica y computarizada |
| 52 | [`NewVentaMineralBCBBuilder`](../../../pkg/models/invoices/venta_mineral_bcb.go) | Sólo electrónica |
| 53 | [`NewLubricantesIehdBuilder`](../../../pkg/models/invoices/lubricantes_iehd.go) | Electrónica y computarizada |
| 54 | [`NewBiodieselBuilder`](../../../pkg/models/invoices/biodiesel.go) | Electrónica y computarizada |
| 55 | [`NewVentaCombustibleSinSubvencionBuilder`](../../../pkg/models/invoices/venta_combustible_sin_subvencion.go) | Electrónica y computarizada |

## Verificación

Las pruebas de contrato de `pkg/models/invoices/xsd_v3_catalog_test.go` describen las 93 variantes sin cargar XSD, XML oficiales, credenciales ni red. Comprueban las secuencias completas, nombres, tipos de los setters, valores, nulos y precisión decimal.

```sh
SIAT_RUN=0 go test ./pkg/models/invoices ./internal/core/domain/documents -count=1
```

En la comprobación local contra los XSD originales pasaron los 46 formatos computarizados con valores y con nulos: 92 documentos. La validación integral de los 47 electrónicos requiere XML firmado y `SignatureSchema.xsd`, ausente del ZIP. La aceptación real de SIAT no se ha probado en esta refactorización.

[← Migración a v3](../how-to/servicios-v3.md)
