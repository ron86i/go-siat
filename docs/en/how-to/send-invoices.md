# Send an invoice

The current API uses a sector document model and a typed facturation request.

## 1. Create the document

Import the sector package, for example documents/compraventa. Construct its public Cabecera and Detalle values, then call compraventa.New. The constructor validates the document and returns an error for invalid data.

The [runnable example](../../../examples/compraventa/main.go) contains a complete Compra-Venta document, CUF calculation, and SIAT flow.

## 2. Prepare and submit

The CUIS, CUFD, CUF, invoice number, and fiscal amounts must belong to the document being issued. For electronic mode, also provide an XMLSigner.

    request := facturacion.SolicitudRecepcionFactura{
        CodigoModalidad: facturacion.ModalidadComputarizada,
        Cuis:            cuis,
    }
    request, err := client.Facturacion.PrepararFactura(request, invoice, nil)
    if err != nil {
        return err
    }

    response, err := client.Facturacion.RecepcionFactura(ctx, request)
    if err != nil {
        return err
    }
    if response.Transaccion == nil || !*response.Transaccion {
        return fmt.Errorf("SIN rejected invoice: %s", response.CodigoDescripcion)
    }

For electronic mode, set CodigoModalidad to ModalidadElectronica and pass the signer to PrepararFactura. Computarizada mode uses ModalidadComputarizada and no signer.

## 3. Read the result

A SOAP call can succeed while the SIN reports a rejected transaction. Inspect Transaccion, CodigoEstado, CodigoDescripcion, CodigoRecepcion, and Mensajes. Retain the receipt code and messages for the application's own audit trail.

## Other invoice operations

The facturation client also exposes VerificacionEstadoFactura, AnulacionFactura, ReversionAnulacionFactura, and the validation operations for package and massive submissions. Each takes a typed request from protocols/facturacion.

See [sign electronic invoices](sign-invoices.md), [configuration](../reference/configuration.md), and the [API reference](../reference/api.md).
