# API SIAT 3.0

El módulo v3 se conecta a los servicios SIAT 3.0 del SIN bajo /sfd. La API separa los documentos fiscales de las solicitudes de servicio.

## Documentos

Importa un paquete sectorial desde documents/&lt;sector&gt;. El modelo usa tipos públicos de Go y el constructor y los métodos XML documentados por ese paquete. Compra-Venta está en documents/compraventa; consulta la [referencia de documentos](../reference/documentos-v3.md).

## Servicios

Crea el cliente con siat.NewClient. El cliente agrupa servicios tipados para códigos, compras, facturación, sincronización y operaciones.

    cliente, err := siat.NewClient(siat.ClientConfig{
        BaseURL:        "https://pilotosiatservicios.impuestos.gob.bo/sfd",
        Token:          token,
        Nit:            nit,
        CodigoSistema:  codigoSistema,
        CodigoAmbiente: 2,
    })

Las solicitudes y respuestas están en protocols/&lt;servicio&gt;. Por ejemplo, CUIS y CUFD usan protocols/codigos, y la recepción de facturas usa protocols/facturacion.

## Preparación

Para documentos fiscales compatibles, el cliente de facturación valida y serializa el documento y prepara el XML, la firma, la compresión y el hash que requiere cada operación. Usa el método adecuado: PrepararFactura, PrepararMasiva o PrepararPaquete.

Consulta [configuración](../reference/configuracion.md), [envío de facturas](envio-facturas.md) y [envío por lotes](envio-lotes.md).
