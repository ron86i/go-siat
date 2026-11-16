# API reference

The current entry point for SIAT 3.0 is the root package.

    client, err := siat.NewClient(siat.ClientConfig{
        BaseURL:        "https://pilotosiatservicios.impuestos.gob.bo/sfd",
        Token:          token,
        Nit:            123456789,
        CodigoSistema:  "MY-SYSTEM",
        CodigoAmbiente: 2,
    })

## Package map

| Package | Purpose |
| --- | --- |
| github.com/ron86i/go-siat/v3 | NewClient and ClientConfig |
| github.com/ron86i/go-siat/v3/documents/&lt;sector&gt; | Typed fiscal document models |
| github.com/ron86i/go-siat/v3/protocols/codigos | CUIS, CUFD, and code requests |
| github.com/ron86i/go-siat/v3/protocols/facturacion | Invoice reception and invoice operations |
| github.com/ron86i/go-siat/v3/protocols/compras | Purchase-receipt service |
| github.com/ron86i/go-siat/v3/protocols/sincronizacion | Catalog synchronization |
| github.com/ron86i/go-siat/v3/protocols/operaciones | Events and point-of-sale operations |
| github.com/ron86i/go-siat/v3/pkg/utils | CUF, signing, XML, and compression helpers |

## Client services

The returned protocols.Client exposes Codigos, Compras, Facturacion, Electronica, Computarizada, ServicioBasico, DocumentoAjuste, Telecomunicaciones, EntidadFinanciera, BoletoAereo, Sincronizacion, and Operaciones.

Each service accepts typed request structs and returns typed responses. The service packages document operation-specific fields and validation.

## Facturation operations

protocols/facturacion.Client provides:

- RecepcionFactura
- VerificacionEstadoFactura
- AnulacionFactura and ReversionAnulacionFactura
- RecepcionMasivaFactura and ValidacionRecepcionMasivaFactura
- RecepcionPaqueteFactura and ValidacionRecepcionPaqueteFactura
- PrepararFactura for supported fiscal document models

Use SolicitudRecepcionFactura.PrepararFactura or Client.PrepararFactura to serialize a single document before submission. Use SolicitudRecepcionMasiva.PrepararMasiva or SolicitudRecepcionPaquete.PrepararPaquete for batches.

## Responses and errors

A transport, decoding, or SOAP fault is returned as an error. A successful HTTP response can still contain a negative fiscal result. Check RespuestaRecepcion.Transaccion, CodigoEstado, CodigoDescripcion, CodigoRecepcion, and Mensajes.

See the [configuration reference](configuration.md), [document reference](documents-v3.md), and [operation guides](../how-to/send-invoices.md).
