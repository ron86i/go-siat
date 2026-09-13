# Entender los sectores de facturación

<p align="right">
  <a href="../../en/explanation/sectors.md">🇬🇧 English</a> · <a href="../README.md">Índice de documentación</a>
</p>

El SIAT no tiene un solo formato de factura. Su catálogo normativo define **51 códigos de documento sector**, cada uno con su propio esquema XSD, sus campos obligatorios y sus reglas de validación. Esta página explica qué es un sector, cómo los modela el SDK y cómo encontrar el que necesitás.

---

## Qué es un sector

Un *documento sector* es la forma que tiene el SIAT de decir "este tipo de negocio emite esta forma de factura". Una factura de hotel necesita las fechas de la estadía. Una de exportación minera necesita concentraciones de mineral y cotizaciones internacionales. Una de telecomunicaciones necesita un período de servicio. En vez de un documento gigante con todo opcional, el SIAT publica un esquema separado por actividad.

El código de sector aparece en dos lugares, y tienen que coincidir:

| Dónde | Campo | Quién lo pone |
| :--- | :--- | :--- |
| Dentro del documento de la factura | `cabecera.codigoDocumentoSector` | El builder del sector, con el default correcto |
| En la solicitud de envío | `WithCodigoDocumentoSector(n)` | Vos, de forma explícita |

Si no coinciden, se rechaza con el código `932 — CODIGO DOCUMENTO SECTOR NO CORRESPONDE AL SERVICIO`. El código `931` significa que el valor no es un sector válido, y el `940` que tu NIT no está habilitado para ese sector.

> El SDK trae un builder por sector cubierto, cada uno precargado con el código correcto. Casi nunca deberías llamar a `WithCodigoDocumentoSector` en un builder de *cabecera* — el default ya está bien. Lo que sí tenés que hacer es pasarle el código al builder de la *solicitud*.

---

## Todos los builders de sector tienen la misma forma

Los 51 builders de `pkg/models/invoices` (50 códigos, uno de ellos con dos diseños) siguen un solo patrón, así que aprender uno te enseña todos:

```go
cabecera := invoices.NewHotelCabeceraBuilder().
    WithNitEmisor(123456789).
    WithNumeroFactura(1).
    WithCuf(cuf).
    // ... campos propios del sector
    Build()

detalle := invoices.NewHotelDetalleBuilder().
    // ... campos de línea propios del sector
    Build()

factura := invoices.NewHotelBuilder().
    WithModalidad(siat.ModalidadElectronica).
    WithCabecera(cabecera).
    AddDetalle(detalle).
    Build()
```

Tres constructores por sector: `New<Sector>CabeceraBuilder()` para la cabecera, `New<Sector>DetalleBuilder()` para una línea de detalle y `New<Sector>Builder()` para armarlos. `AddDetalle` es acumulativo — llamalo una vez por línea.

`WithModalidad` en el builder raíz controla el namespace XML y si se emite el espacio para la firma. Es distinto del `WithCodigoModalidad` que ponés en la solicitud de *envío*, y hay que setear los dos.

Entre sectores lo único que cambia es la lista de campos de la cabecera y el detalle. El autocompletado de tu editor sobre el tipo de builder devuelto es la referencia más rápida que existe.

### Campos JSON libres

Algunos campos sectoriales aceptan JSON libre, por ejemplo
`WithDetalleHuespedes` de hotel y los campos de costos, bultos u otros datos de
las cabeceras de exportación. Pasales un `string` JSON ya preparado o un valor
Go que `encoding/json` pueda serializar. Si el valor no se puede serializar, el
builder conserva el error y no entra en pánico:

```go
detalleBuilder := invoices.NewHotelDetalleBuilder().
    WithDetalleHuespedes(map[string]any{"nombre": "Ana"})
if err := detalleBuilder.Err(); err != nil {
    return err
}
detalle := detalleBuilder.Build()
```

`Err()` está disponible en `HotelDetalle` y en las cabeceras de exportación
comercial, exportación de hidrocarburos, exportación minera y exportación a
precio de venta.

---

## Cómo elige un sector su endpoint

El SIAT expone doce endpoints SOAP. Cuál acepta tu factura no es una elección libre — lo determina el sector.

**Seis sectores tienen endpoint propio.** No los aceptan en ningún otro lado:

| Fachada | Sectores que atiende |
| :--- | :--- |
| `s.CompraVenta()` | CompraVenta (1), CompraVentaBonificaciones (35), CompraVentaTasas (41) |
| `s.Telecomunicaciones()` | Telecomunicaciones (22), TelecomunicacionesZF (49) |
| `s.ServicioBasico()` | ServicioBasico (13), ServicioBasicoZF (40) |
| `s.EntidadFinanciera()` | EntidadFinanciera (15) |
| `s.BoletoAereo()` | BoletoAereo (30) |
| `s.DocumentoAjuste()` | NotaCreditoDebito (24), NotaFiscalCreditoDebito (24), NotaConciliacion (29), NotaCreditoDebitoDescuento (47), NotaCreditoDebitoIce (48) |

**Los otros 37 sectores se enrutan por modalidad**, no por actividad. Elegí la fachada que corresponde a cómo estás facturando:

```go
s.Electronica().RecepcionFactura(ctx, req)     // ModalidadElectronica — firmada
s.Computarizada().RecepcionFactura(ctx, req)   // ModalidadComputarizada — código de control, sin firma
```

Las dos exponen la misma interfaz de nueve métodos, así que cambiar de modalidad es cambiar una palabra en el punto de llamada. Mirá [Arquitectura](arquitectura.md) para entender por qué las interfaces tienen esa forma.

`DocumentoAjuste()` es la excepción a todo el modelo: las notas de ajuste no son facturas, así que sus métodos se llaman `RecepcionDocumentoAjuste`, `AnulacionDocumentoAjuste` y `ReversionAnulacionDocumentoAjuste` en vez de los nombres `*Factura`.

---

## Catálogo completo de sectores

El catálogo normativo del SIAT define **51 códigos de documento sector**. El SDK trae builders para **50 de ellos**; el único que falta está marcado abajo.

Los nombres de builder omiten el envoltorio `New`/`Builder` — `CompraVenta` significa `invoices.NewCompraVentaBuilder()`. La columna **Modalidades** es lo que acepta `WithModalidad` en el builder raíz para ese sector; la columna **Fachada** es qué método de servicio usás para el *envío* — en los sectores marcados "`Electronica()` o `Computarizada()`", elegís la fachada según la modalidad que construiste.

| Cód. | Descripción | Crédito fiscal | Modalidades | Builder | Fachada |
| ---: | :--- | :--- | :--- | :--- | :--- |
| 1 | Compra y Venta | Con | Electrónica y Computarizada | `CompraVenta` | `CompraVenta()` |
| 2 | Alquiler de Bienes Inmuebles | Con | Electrónica y Computarizada | `AlquilerBienInmueble` | `Electronica()` o `Computarizada()` |
| 3 | Comercial de Exportación | Sin | Electrónica y Computarizada | `ComercialExportacion` | `Electronica()` o `Computarizada()` |
| 4 | Comercial de Exportación en Libre Consignación | Sin | Electrónica y Computarizada | `LibreConsignacion` | `Electronica()` o `Computarizada()` |
| 5 | Venta en Zona Franca | Sin | Electrónica y Computarizada | `ZonaFranca` | `Electronica()` o `Computarizada()` |
| 6 | Servicio Turístico y Hospedaje | Sin | Electrónica y Computarizada | `TurismoHospedaje` | `Electronica()` o `Computarizada()` |
| 7 | Seguridad Alimentaria y Abastecimiento | Sin | Electrónica y Computarizada | `SeguridadAlimentaria` | `Electronica()` o `Computarizada()` |
| 8 | Tasa Cero — libros y transporte internacional de carga | Sin | Electrónica y Computarizada | `TasaCero` | `Electronica()` o `Computarizada()` |
| 9 | Compra y Venta de Moneda Extranjera | Sin | Electrónica y Computarizada | `MonedaExtranjera` | `Electronica()` o `Computarizada()` |
| 10 | Dutty Free | Sin | Electrónica y Computarizada | `DuttyFree` | `Electronica()` o `Computarizada()` |
| 11 | Sectores Educativos | Con | Electrónica y Computarizada | `SectorEducativo` | `Electronica()` o `Computarizada()` |
| 12 | Comercialización de Hidrocarburos | Con | Electrónica y Computarizada | `ComercializacionHidro` | `Electronica()` o `Computarizada()` |
| 13 | Servicios Básicos | Con | Electrónica y Computarizada | `ServicioBasico` | `ServicioBasico()` |
| 14 | Productos Alcanzados por el ICE | Con | Electrónica y Computarizada | `AlcanzadaIce` | `Electronica()` o `Computarizada()` |
| 15 | Entidades Financieras | Con | Electrónica y Computarizada | `EntidadFinanciera` | `EntidadFinanciera()` |
| 16 | Hoteles | Con | Electrónica y Computarizada | `Hotel` | `Electronica()` o `Computarizada()` |
| 17 | Hospitales / Clínicas | Con | Electrónica y Computarizada | `HospitalClinica` | `Electronica()` o `Computarizada()` |
| 18 | Juegos de Azar | Con | Electrónica y Computarizada | `JuegoAzar` | `Electronica()` o `Computarizada()` |
| 19 | Hidrocarburos Alcanzada IEHD | Con | Electrónica y Computarizada | `HidrocarburoAlcanzadaIehd` | `Electronica()` o `Computarizada()` |
| 20 | Comercial de Exportación de Minerales | Sin | Electrónica y Computarizada | `ComercialExportacionMinera` | `Electronica()` o `Computarizada()` |
| 21 | Venta de Minerales | Con | Electrónica y Computarizada | `VentaMineral` | `Electronica()` o `Computarizada()` |
| 22 | Telecomunicaciones | Con | Electrónica y Computarizada | `Telecomunicaciones` | `Telecomunicaciones()` |
| 23 | Prevalorada | Con | Electrónica y Computarizada | `Prevalorada` | `Electronica()` o `Computarizada()` |
| 24 | Nota de Crédito-Débito | Ajuste | Electrónica y Computarizada | `NotaCreditoDebito` · `NotaFiscalCreditoDebito` | `DocumentoAjuste()` |
| 28 | Comercial de Exportación de Servicios | Sin | Electrónica y Computarizada | `ComercialExportacionServicio` | `Electronica()` o `Computarizada()` |
| 29 | Nota de Conciliación | Ajuste | Electrónica y Computarizada | `NotaConciliacion` | `DocumentoAjuste()` |
| 30 | Boleto Aéreo | Equivalente | Electrónica y Computarizada | `BoletoAereo` | `BoletoAereo()` |
| 31 | Suministro de Energía | Con | Electrónica y Computarizada | `SuministroEnergia` | `Electronica()` o `Computarizada()` |
| **33** | **Tasa Cero IVA Ley N° 1613** | Sin | — | **— sin builder** | — |
| 34 | Seguros | Con | Electrónica y Computarizada | `Seguros` | `Electronica()` o `Computarizada()` |
| 35 | Compra Venta Bonificaciones | Con | Electrónica y Computarizada | `CompraVentaBonificaciones` | `CompraVenta()` |
| 36 | Prevalorada Sin Derecho a Crédito Fiscal | Sin | Electrónica y Computarizada | `PrevaloradaSinDerechoCreditoFiscal` | `Electronica()` o `Computarizada()` |
| 37 | Comercialización de GNV | Con | Electrónica y Computarizada | `ComercializacionGnv` | `Electronica()` o `Computarizada()` |
| 38 | Hidrocarburos No Alcanzada IEHD | Con | Electrónica y Computarizada | `HidrocarburoNoAlcanzadaIehd` | `Electronica()` o `Computarizada()` |
| 39 | Comercialización de GN y GLP | Con | Electrónica y Computarizada | `ComercializacionGnGlp` | `Electronica()` o `Computarizada()` |
| 40 | Servicios Básicos Zona Franca | Sin | Electrónica y Computarizada | `ServicioBasicoZF` | `ServicioBasico()` |
| 41 | Compra Venta Tasas | Con | Electrónica y Computarizada | `CompraVentaTasas` | `CompraVenta()` |
| 42 | Alquiler Zona Franca | Sin | Electrónica y Computarizada | `AlquilerZF` | `Electronica()` o `Computarizada()` |
| 43 | Comercial de Exportación Hidrocarburos | Sin | Electrónica y Computarizada | `ComercialExportacionHidro` | `Electronica()` o `Computarizada()` |
| 44 | Importación y Comercialización de Lubricantes | Con | Electrónica y Computarizada | `ImportacionComercializacionLubricantes` | `Electronica()` o `Computarizada()` |
| 45 | Comercial de Exportación Precio Venta | Sin | Electrónica y Computarizada | `ComercialExportacionPVenta` | `Electronica()` o `Computarizada()` |
| 46 | Sector Educativo Zona Franca | Sin | Electrónica y Computarizada | `SectorEducativoZF` | `Electronica()` o `Computarizada()` |
| 47 | Nota Crédito Débito Descuentos | Ajuste | Electrónica y Computarizada | `NotaCreditoDebitoDescuento` | `DocumentoAjuste()` |
| 48 | Nota Crédito Débito ICE | Ajuste | Electrónica y Computarizada | `NotaCreditoDebitoIce` | `DocumentoAjuste()` |
| 49 | Telecomunicaciones Zona Franca | Sin | Electrónica y Computarizada | `TelecomunicacionesZF` | `Telecomunicaciones()` |
| 50 | Hospitales / Clínicas Zona Franca | Sin | Electrónica y Computarizada | `HospitalClinicaZF` | `Electronica()` o `Computarizada()` |
| 51 | Engarrafadoras | Con | Electrónica y Computarizada | `Engarrafadoras` | `Electronica()` o `Computarizada()` |
| 52 | Venta de Minerales al Banco Central | Sin | **Solo Electrónica** | `VentaMineralBCB` | `Electronica()` únicamente |
| 53 | Importación y Comercialización de Lubricantes IEHD | Con | Electrónica y Computarizada | `LubricantesIehd` | `Electronica()` o `Computarizada()` |
| 54 | Compra-Venta de Insumos para Biodiésel / Diésel Ecológico | Sin | Electrónica y Computarizada | `Biodiesel` | `Electronica()` o `Computarizada()` |
| 55 | Comercialización de Combustible | Con | Electrónica y Computarizada | `VentaCombustibleSinSubvencion` | `Electronica()` o `Computarizada()` |

Cinco cosas que vale la pena leer de esta tabla:

**Los códigos no son contiguos.** Faltan el 25, 26, 27 y 32 en el propio catálogo del SIAT. No los inventes ni los uses como relleno.

**El 24 aparece dos veces a propósito.** `NotaCreditoDebito` y `NotaFiscalCreditoDebito` son dos diseños de documento que comparten el mismo código de sector. Elegí el que tenga la lista de campos que corresponde a la nota que estás emitiendo.

**Solo el 52 está restringido a una modalidad.** El builder de todos los demás sectores acepta tanto `ModalidadElectronica` como `ModalidadComputarizada` — la elección es tuya según cómo estés facturando. La normativa restringe la venta de minerales al Banco Central específicamente a la modalidad electrónica, así que `VentaMineralBCB` va por `Electronica()` y siempre exige firma; pasarle `ModalidadComputarizada` a su builder produce un documento que el SIAT va a rechazar.

**Modalidades y Fachada responden preguntas distintas.** Modalidades es cómo se construye el *documento* de la factura. Fachada es qué *método de servicio* llamás para enviarlo. En los catorce sectores con endpoint propio estas dos cosas están desacopladas — un envío por `CompraVenta()` puede llevar un documento electrónico o computarizado, porque la fachada se elige por actividad, no por modalidad. En el resto de los sectores las dos cosas se reducen a una sola elección: si construiste electrónica, llamá a `Electronica()`; si construiste computarizada, llamá a `Computarizada()`.

**`ZF` significa zona franca.** Es la variante del sector para zona franca, con otro código y campos adicionales — y casi siempre sin derecho a crédito fiscal aunque el sector base sí lo tenga (compará 13 con 40, 22 con 49, 17 con 50).

### Los cuatro sectores de hidrocarburos y lubricantes

Son los más fáciles de confundir, porque los nombres se superponen y la diferencia es normativa y no estructural:

| Cód. | Builder | Se aplica a |
| ---: | :--- | :--- |
| 12 | `ComercializacionHidro` | Venta de combustible — diésel, gasolina, automotores |
| 19 | `HidrocarburoAlcanzadaIehd` | Actividades **alcanzadas** por el IEHD |
| 38 | `HidrocarburoNoAlcanzadaIehd` | Actividades **exentas** del IEHD |
| 44 / 53 | `ImportacionComercializacionLubricantes` / `LubricantesIehd` | Importación y venta de lubricantes, sin y con IEHD |

Los sectores 19 y 38 comparten un diseño idéntico salvo que el 19 lleva `montoIehd` en la cabecera y `porcentajeIehd` por línea de detalle. Si estás exento, el 38 rechaza esos campos; si estás alcanzado, el 19 los exige.

### El único sector sin cobertura

El código 33 está en el catálogo del SIAT pero no tiene builder en el SDK:

| Cód. | Descripción |
| ---: | :--- |
| 33 | Factura Tasa Cero IVA Ley N° 1613 — bienes de capital y plantas industriales |

Si tu actividad cae ahí, no hay atajo dentro del SDK: tenés que armar el struct del documento vos mismo y cargarlo con `WithArchivo` / `WithHashArchivo` en lugar de `WithFactura`. Mirá [armar el archivo a mano](../how-to/envio-facturas.md#empaquetar-a-mano).

No reemplaces el 33 por el sector 8 (`TasaCero`). Los dos son tasa cero, pero el 8 cubre libros y transporte internacional de carga con otra base legal, y el SIAT rechaza la confusión con el código 932.

---

## Encontrar los campos de tu sector

Los tests del propio SDK son los ejemplos más confiables, porque son el código que realmente va y vuelve contra el SIAT. Cada sector tiene el suyo:

```
pkg/models/invoices/<sector>_test.go
```

Por ejemplo, `pkg/models/invoices/hotel_test.go` arma una factura de hotel completa con todos los campos obligatorios cargados. Copiá eso, cambiá los valores y ya tenés un documento que funciona.

Los sectores cuyos códigos no reconocés conviene confirmarlos contra el SIAT en vez de adivinarlos — `Sincronizacion()` expone `SincronizarParametricaTipoDocumentoSector` para el catálogo completo y `SincronizarListaActividadesDocumentoSector` para saber qué sectores habilitan tus actividades económicas registradas.

---

## Relacionado

| | |
| :--- | :--- |
| Armar y enviar tu primera factura | [Tutorial](../tutorial/primera-factura.md) |
| Por qué los servicios tienen esta forma | [Explicación: Arquitectura](arquitectura.md) |
| Firmar facturas en modalidad electrónica | [Guía: Firmar facturas](../how-to/firmar-facturas.md) |
| Códigos de rechazo por sector (931/932/940) | [Guía: Manejo de errores](../how-to/manejo-errores.md) |
