# Migrar a go-siat v3

El módulo `github.com/ron86i/go-siat/v3` usa los contratos SOAP SIAT v3
de la ruta `/sfd/`. Tiene una sola entrada: `siat.New(config)`.
La preparación de esta rama no publica automáticamente `v3.0.0`.

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
