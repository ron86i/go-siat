# Primera factura SIAT 3.0

Este tutorial usa la API v3 actual. El ejemplo Compra-Venta completo y ejecutable está en [examples/compraventa/main.go](../../../examples/compraventa/main.go).

## 1. Configurar el cliente

Usa el endpoint oficial de piloto durante el desarrollo y mantén el token delegado fuera del repositorio.

    cliente, err := siat.NewClient(siat.ClientConfig{
        BaseURL:        "https://pilotosiatservicios.impuestos.gob.bo/sfd",
        Token:          os.Getenv("SIAT_TOKEN"),
        Nit:            nit,
        CodigoSistema:  os.Getenv("SIAT_CODIGO_SISTEMA"),
        CodigoAmbiente: 2,
    })
    if err != nil {
        return err
    }

## 2. Obtener los códigos vigentes

El servicio de códigos recibe solicitudes tipadas.

    cuis, err := cliente.Codigos.SolicitarCUIS(ctx, codigos.SolicitudCUIS{
        CodigoModalidad: facturacion.ModalidadComputarizada,
        CodigoSucursal:  sucursal,
        CodigoPuntoVenta: &puntoVenta,
    })
    if err != nil {
        return err
    }

Usa el CUIS devuelto para solicitar el CUFD. Con el CUFD y su código de control se calcula el CUF de la factura.

## 3. Crear el documento sectorial

Importa el paquete correcto, como documents/compraventa. Completa los structs fiscales de cabecera y detalle con los datos ya validados por tu aplicación y llama a compraventa.New(cabecera, detalles). El constructor valida el documento y devuelve un error si faltan datos requeridos o son inconsistentes.

El SDK no calcula tus impuestos ni decide los importes de la factura.

## 4. Preparar y enviar

    solicitud := facturacion.SolicitudRecepcionFactura{
        CodigoModalidad: facturacion.ModalidadComputarizada,
        Cuis:            cuis.Codigo,
    }
    solicitud, err = cliente.Facturacion.PrepararFactura(solicitud, factura, nil)
    if err != nil {
        return err
    }
    respuesta, err := cliente.Facturacion.RecepcionFactura(ctx, solicitud)
    if err != nil {
        return err
    }

Revisa el estado de transacción, el código de respuesta, el código de recepción y los mensajes. Que SOAP responda no demuestra por sí solo la aceptación fiscal.

En modalidad electrónica, crea un XMLSigner y pásalo a PrepararFactura. Consulta [firma](../how-to/firmar-facturas.md), [configuración](../reference/configuracion.md) y la [referencia documental](../reference/documentos-v3.md).
