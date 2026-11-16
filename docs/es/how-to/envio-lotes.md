# Enviar lotes y paquetes de contingencia

Usa solicitudes tipadas de facturación para la emisión masiva en línea y los paquetes de contingencia. Prepara los documentos antes de llamar al servicio SOAP.

## Emisión masiva en línea

    solicitud := facturacion.SolicitudRecepcionMasiva{
        SolicitudRecepcionFactura: facturacion.SolicitudRecepcionFactura{
            CodigoModalidad: facturacion.ModalidadComputarizada,
            Cuis:            cuis,
        },
    }
    solicitud, err := solicitud.PrepararMasiva(ctx, documentos, nil)
    if err != nil {
        return err
    }
    respuesta, err := cliente.Facturacion.RecepcionMasivaFactura(ctx, solicitud)

En modalidad electrónica pasa el firmador. Los documentos de un lote deben cumplir los requisitos de sector y modalidad del protocolo.

## Paquete de contingencia

El paquete usa emisión fuera de línea y requiere el código del evento significativo.

    solicitud := facturacion.SolicitudRecepcionPaquete{
        SolicitudRecepcionFactura: facturacion.SolicitudRecepcionFactura{
            CodigoModalidad: facturacion.ModalidadComputarizada,
            Cuis:            cuis,
        },
        CodigoEvento: codigoEvento,
    }
    solicitud, err := solicitud.PrepararPaquete(ctx, documentos, nil)
    if err != nil {
        return err
    }
    respuesta, err := cliente.Facturacion.RecepcionPaqueteFactura(ctx, solicitud)

La preparación crea el TAR.GZ, registra la cantidad de documentos y calcula el hash. En modalidad electrónica, pasa el XMLSigner.

## Validar el procesamiento

Un código de recepción confirma que el SIN recibió el paquete; no confirma por sí solo que aceptó todas las facturas. Usa ValidacionRecepcionPaqueteFactura o ValidacionRecepcionMasivaFactura con los metadatos de la operación y el CodigoRecepcion devuelto inicialmente. Revisa transacción y mensajes en la respuesta de validación.

No reenvíes el paquete solo porque la validación final siga pendiente. Conserva el código de recepción y consulta su estado.

Consulta la [referencia de API](../reference/api.md) y la [guía de firma](firmar-facturas.md).
