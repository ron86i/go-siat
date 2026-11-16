# Arquitectura

La API pública v3 separa los documentos fiscales de las llamadas a los servicios SIAT.

## Documentos fiscales

Los paquetes bajo documents/&lt;sector&gt; contienen el modelo Go de cada familia documental. Por ejemplo, documents/compraventa contiene Compra-Venta. Cada paquete documenta sus tipos, constructores, validación y reglas de serialización XML.

Los paquetes de documentos no calculan los impuestos. El llamador proporciona los valores fiscales según sus reglas de negocio.

## Clientes SIAT

El paquete raíz expone siat.NewClient y siat.ClientConfig. El cliente agrupa clientes de servicios:

- Codigos: solicitudes CUIS, CUFD y otras operaciones de códigos.
- Sincronizacion y Operaciones: catálogos, puntos de venta y eventos.
- Facturacion y clientes de facturación especializados: recepción y operaciones de facturas.
- Compras: recepción de documentos de compra.

Los subpaquetes protocols definen solicitudes y respuestas tipadas por servicio. Los sobres SOAP y el transporte HTTP permanecen dentro del SDK.

## Envío de documentos

Para los tipos admitidos, el cliente de facturación valida y serializa el documento, lo firma cuando la modalidad electrónica lo requiere y prepara los datos comprimidos. Después envía la solicitud tipada.

    solicitud := facturacion.SolicitudRecepcionFactura{
        CodigoModalidad: facturacion.ModalidadComputarizada,
        Cuis:            cuis,
    }
    solicitud, err := cliente.Facturacion.PrepararFactura(solicitud, factura, nil)
    if err != nil {
        return err
    }
    respuesta, err := cliente.Facturacion.RecepcionFactura(ctx, solicitud)

Consulta [enviar una factura](../how-to/envio-facturas.md) para ver el flujo completo.

## Límites

- documents y protocols son paquetes públicos.
- internal/pipeline e internal/soapclient son detalles internos y no se pueden importar desde otros módulos.
- El SDK transporta solicitudes; el llamador conserva la responsabilidad por sus datos comerciales y cálculos tributarios.
