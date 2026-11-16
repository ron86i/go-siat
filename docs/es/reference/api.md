# Referencia de API

El punto de entrada actual para SIAT 3.0 es el paquete raíz.

    cliente, err := siat.NewClient(siat.ClientConfig{
        BaseURL:        "https://pilotosiatservicios.impuestos.gob.bo/sfd",
        Token:          token,
        Nit:            123456789,
        CodigoSistema:  "MI-SISTEMA",
        CodigoAmbiente: 2,
    })

## Paquetes

| Paquete | Responsabilidad |
| --- | --- |
| github.com/ron86i/go-siat/v3 | NewClient y ClientConfig |
| github.com/ron86i/go-siat/v3/documents/&lt;sector&gt; | Modelos fiscales tipados |
| github.com/ron86i/go-siat/v3/protocols/codigos | Solicitudes CUIS, CUFD y códigos |
| github.com/ron86i/go-siat/v3/protocols/facturacion | Recepción y operaciones de facturas |
| github.com/ron86i/go-siat/v3/protocols/compras | Servicio de recepción de compras |
| github.com/ron86i/go-siat/v3/protocols/sincronizacion | Sincronización de catálogos |
| github.com/ron86i/go-siat/v3/protocols/operaciones | Eventos y puntos de venta |
| github.com/ron86i/go-siat/v3/pkg/utils | CUF, firma, XML y compresión |

## Servicios del cliente

El protocols.Client devuelto expone Codigos, Compras, Facturacion, Electronica, Computarizada, ServicioBasico, DocumentoAjuste, Telecomunicaciones, EntidadFinanciera, BoletoAereo, Sincronizacion y Operaciones.

Cada servicio recibe structs de solicitud tipados y devuelve respuestas tipadas. Los paquetes de servicio documentan los campos y validaciones de cada operación.

## Operaciones de facturación

protocols/facturacion.Client ofrece:

- RecepcionFactura
- VerificacionEstadoFactura
- AnulacionFactura y ReversionAnulacionFactura
- RecepcionMasivaFactura y ValidacionRecepcionMasivaFactura
- RecepcionPaqueteFactura y ValidacionRecepcionPaqueteFactura
- PrepararFactura para modelos fiscales compatibles

Usa SolicitudRecepcionFactura.PrepararFactura o Client.PrepararFactura para serializar un documento individual. Para lotes usa SolicitudRecepcionMasiva.PrepararMasiva o SolicitudRecepcionPaquete.PrepararPaquete.

## Respuestas y errores

Los errores de transporte, decodificación y SOAP se devuelven como error. Una respuesta HTTP satisfactoria todavía puede contener un resultado fiscal negativo. Revisa RespuestaRecepcion.Transaccion, CodigoEstado, CodigoDescripcion, CodigoRecepcion y Mensajes.

Consulta la [configuración](configuracion.md), la [referencia de documentos](documentos-v3.md) y las [guías de operaciones](../how-to/envio-facturas.md).
