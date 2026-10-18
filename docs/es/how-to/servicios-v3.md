# Migrar a go-siat v3

El módulo `github.com/ron86i/go-siat/v3` usa los contratos SOAP SIAT v3
de la ruta `/sfd/`. Tiene una sola entrada: `siat.New(config)`.
La preparación de esta rama no publica automáticamente `v3.0.0`.

## Documentos fiscales: esquemas de importes netos

Los modelos cubren las 47 familias y las 93 raíces XSD del paquete de documentos
proporcionado para esta migración. La [cobertura por documento](../reference/documentos-v3.md)
detalla sus builders y modalidades. Cada familia conserva su propia secuencia,
tipos, campos sectoriales y precisión decimal.

- La cabecera usa `WithSubTotalNeto`, `WithTotalPrecioNeto`, `WithIva` y
  `WithMontoAPagar`, en lugar de `WithMontoTotal` y `WithMontoTotalSujetoIva`.
- `WithCodigoInterno` configura el campo obligatorio de 1 a 10 caracteres;
  su valor debe corresponder al que requiera el servicio, no a un valor ficticio.
- El detalle usa `WithPrecioNetoUnitario` en lugar de `WithPrecioUnitario`.
- Se elimina `WithCafc` de la cabecera de Compra/Venta.
- `WithComplemento("")` envía una cadena vacía cuando no aplica; este campo no
  admite `xsi:nil`. `WithCodigoInterno` recibe `string`.
- La precisión de cada setter decimal sigue `fractionDigits` de su XSD:
  Compra/Venta usa dos decimales; el detalle de Vuelo Chárter admite cinco.
- Los campos nillables conservan parámetros puntero y emiten `xsi:nil="true"`
  cuando reciben `nil`. Un puntero a cero emite el valor `0`.
- Las notas usan importes separados de la factura original y de la devolución,
  como `WithSubTotalNetoOriginal`, `WithIvaDevuelto` y `WithMontoAPagarDevuelto`.
- Se añade `NewVueloCharterBuilder` (sector 41), con `WithMontoTasa`, series e IMEI.
  Venta Mineral BCB sólo tiene esquema electrónico en el paquete suministrado.

Los setters no calculan impuestos ni deducen un importe de otro: el consumidor
debe suministrar los valores correspondientes. Los nombres XML distinguen
mayúsculas: `subTotalNeto`, `totalPrecioNeto`, `iva`, `montoaPagar`.

Los nombres anteriores se retiran en las familias migradas. Por ejemplo, el
detalle usa `WithPrecioNetoUnitario` y los campos de JSON libre conservan sus
setters y los errores originales de `encoding/json`.

Para verificar toda la cobertura sin credenciales ni conexión al SIAT:

```sh
SIAT_RUN=0 go test ./pkg/models/invoices ./internal/core/domain/documents -count=1
```

`TestFacturasV3ContratosXSD` cubre raíces, secuencias, valores, precisión y nulos
para las 93 variantes. La validación local adicional contra los XSD pasó para
los 46 formatos computarizados con valores y nulos. La validación XSD integral
de los 47 electrónicos requiere el `SignatureSchema.xsd` ausente del ZIP y XML
firmado. Estas comprobaciones no acreditan aceptación del servicio SIAT.

Los fixtures de Dutty Free, Boleto Aéreo y Nota de Conciliación tienen una
regresión adicional que construye sus documentos sin conexión:

```sh
go test ./internal/adapter/services -run '^TestFacturasV3FixturesSinRed$' -count=1
```

Para comprobar el XML computarizado contra un XSD local, sin copiar contratos
al repositorio ni llamar al SIAT (requiere `xmllint`):

```sh
SIAT_COMPRA_VENTA_XSD='/ruta/facturaComputarizadaCompraVenta.xsd' \
go test ./pkg/models/invoices -run '^TestCompraVentaContratoNeto$' -count=1
```

La prueba comprueba campos y secuencias en ambas modalidades; la validación
XSD externa se aplica a la computarizada. Validar la electrónica completa
requiere además su firma y el archivo `SignatureSchema.xsd` importado.

## Cambiar los imports y la dirección base

```go
import (
    "github.com/ron86i/go-siat/v3"
    "github.com/ron86i/go-siat/v3/pkg/models"
)

cliente, err := siat.New(siat.Config{
    BaseURL:        "https://pilotosiatservicios.impuestos.gob.bo/sfd",
    Token:          token,
    Nit:            nit,
    CodigoSistema:  codigoSistema,
    CodigoAmbiente: siat.AmbientePruebas,
})
if err != nil {
    return err
}

solicitud := models.NewCufdBuilder().
    WithCodigoModalidad(siat.ModalidadElectronica).
    WithCodigoSucursal(0).
    WithCodigoPuntoVenta(1).
    WithCuis(cuis).
    Build()
respuesta, err := cliente.Codigos().SolicitudCufd(ctx, solicitud)
if err != nil {
    return err
}
cufd := respuesta.Body.Content.RespuestaCufd
if !cufd.Transaccion {
    return fmt.Errorf("CUFD rechazado: %+v", cufd.MensajesList)
}
// cufd.CodigoInterno conserva el campo adicional de la respuesta v3.
```

La dirección es explícita: el SDK no convierte una configuración de producción
en piloto ni sustituye su host. El constructor admite una barra final y la
normaliza. Confirmar la dirección v3 de producción con el SIN; el endpoint
`siatrest.impuestos.gob.bo/sfd/FacturacionCodigos?wsdl` devolvió HTTP 404 en la
verificación del 18/10/2026.

Se mantienen los accesos `Codigos`, `Sincronizacion`, `Operaciones`,
`CompraVenta`, `Electronica`, `Computarizada`, `DocumentoAjuste`,
`ServicioBasico`, `Telecomunicaciones`, `EntidadFinanciera`, `BoletoAereo`
y `RecepcionCompras`. Los builders siguen en `pkg/models`; el transporte
conserva el cliente HTTP configurable y `siat.WithDynamicConfig`.

## Cambios de contrato y API

- `RecepcionPaqueteFacturaBuilder.WithCafc` se retira: los paquetes SOAP v3 no
  contienen `cafc`. Los campos del XML de cada factura siguen su XSD sectorial.
- `NewConsultaEventoSignificativoBuilder().WithCufd(cufd)` corrige una omisión
  del SDK: la consulta de eventos ya requería el CUFD en v2; no es un requisito
  nuevo de v3.
- `RespuestaCufd.CodigoInterno` conserva el nuevo dato de Códigos.
- `ListaCodigosProductos.CodigoActividad` pasa a `string`, como exige el WSDL;
  `CodigoProducto` es `int64` y `Nandina` conserva la lista adicional.
- `RespuestaRecepcion.MensajesList` usa `MensajeRecepcion`, que incluye
  `Codigo`, `Descripcion`, `Advertencia`, `NumeroArchivo` y `NumeroDetalle`.
- La anulación serializa `codigoMotivo` antes de `cuf`.
- La identidad dinámica se inyecta en una copia: reutilizar un builder no
  conserva las credenciales de una petición anterior.

## Verificaciones

Los tests de servicios existentes cargan la dirección desde `SIAT_URL`.
Actualizá ese valor en tu configuración local antes de ejecutar pruebas reales:

```dotenv
SIAT_URL=https://pilotosiatservicios.impuestos.gob.bo/sfd
```

El archivo `.env` local no se modifica durante esta migración.

Las pruebas HTTP locales se ejecutan con:

```sh
go test . -run '^TestSIATV3_' -count=1
```

`ServicioRecepcionCompras` queda pendiente de verificación externa por una
respuesta HTTP 503 del servidor.

Las pruebas HTTP locales verifican rutas, namespace, bodies y configuración dinámica. No
validan valores fiscales ni prueban aceptación oficial del SIN.

### Pruebas reales de códigos

Las pruebas de CUIS/CUFD requieren `SIAT_RUN=1`, ambiente 2 y una `SIAT_URL`
terminada en `/sfd`. Usan `SIAT_PV` (1 por defecto) y `SIAT_SUCURSAL` (0 por
defecto); también admiten `SIAT_CODIGO_PUNTO_VENTA` y `SIAT_CODIGO_SUCURSAL`.
Las variables cortas tienen prioridad sobre las largas.

```sh
SIAT_RUN=1 SIAT_PV=1 go test -v ./internal/adapter/services \
  -timeout 30s -run '^TestSolicitudCuis$' -count=1
```

Cada prueba guarda solicitudes y respuestas completas en `.siat-traces/<nombre>`.
`SIAT_CUIS` permite reutilizar un código vigente para las pruebas de CUFD y
verificación de NIT. `TestSolicitudCuis` siempre solicita uno para comprobar
esa operación. `TestSolicitudCufdAll` obtiene CUIS una sola vez y solicita
`SIAT_VECES` CUFD (10 por defecto), deteniéndose al primer rechazo.

`TestSolicitudCufdMasivo` llama a la operación masiva de CUFD y comprueba su
respuesta y cada entrada. La notificación de revocación sólo se configura
cuando `SIAT_CERT_REVOCADO` apunta a un certificado real de pruebas.

## Fechas con zona horaria

`TimeSiat` es un alias de `time.Time`: XML y JSON usan el formato estándar
RFC 3339, con `Z` para UTC o un desplazamiento como `-04:00`. Los builders
conservan la zona recibida; no convierten automáticamente a UTC. Usa
`fecha.UTC()` para realizar esa conversión explícitamente.

Este cambio también afecta las fechas de los documentos fiscales. Las fechas
de entrada sin zona ya no son aceptadas por el decodificador estándar. La
compatibilidad del formato con cada XSD de facturación debe verificarse antes
de enviar documentos; las pruebas locales no demuestran aceptación del SIAT.
