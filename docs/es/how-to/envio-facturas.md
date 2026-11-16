# Enviar una factura

La API actual usa un modelo documental sectorial y una solicitud tipada de facturación.

## 1. Crear el documento

Importa el paquete del sector, por ejemplo documents/compraventa. Completa sus structs Cabecera y Detalle y llama a compraventa.New. El constructor valida el documento y devuelve un error si los datos no cumplen su estructura.

El [ejemplo ejecutable](../../../examples/compraventa/main.go) contiene una factura Compra-Venta completa, el cálculo de CUF y el flujo SIAT.

## 2. Preparar y enviar

El CUIS, CUFD, CUF, número de factura e importes fiscales deben corresponder al documento. En modalidad electrónica también se debe pasar un XMLSigner.

    solicitud := facturacion.SolicitudRecepcionFactura{
        CodigoModalidad: facturacion.ModalidadComputarizada,
        Cuis:            cuis,
    }
    solicitud, err := cliente.Facturacion.PrepararFactura(solicitud, factura, nil)
    if err != nil {
        return err
    }

    respuesta, err := cliente.Facturacion.RecepcionFactura(ctx, solicitud)
    if err != nil {
        return err
    }
    if respuesta.Transaccion == nil || !*respuesta.Transaccion {
        return fmt.Errorf("el SIN rechazó la factura: %s", respuesta.CodigoDescripcion)
    }

Para modalidad electrónica usa ModalidadElectronica y pasa el firmador a PrepararFactura. La modalidad computarizada usa ModalidadComputarizada y no admite firmador.

## 3. Leer el resultado

Una llamada SOAP puede finalizar correctamente aunque el SIN rechace la transacción. Revisa Transaccion, CodigoEstado, CodigoDescripcion, CodigoRecepcion y Mensajes. Conserva el código y los mensajes en la auditoría de tu aplicación.

## Otras operaciones

El cliente de facturación también expone VerificacionEstadoFactura, AnulacionFactura, ReversionAnulacionFactura y operaciones de validación de envíos masivos y paquetes. Cada método recibe una solicitud tipada de protocols/facturacion.

Consulta [firmar facturas](firmar-facturas.md), [configuración](../reference/configuracion.md) y la [referencia de API](../reference/api.md).
