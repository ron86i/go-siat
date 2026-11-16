# Send batches and contingency packages

Use typed facturation requests for both massive online issuance and contingency packages. Prepare the documents before calling the SOAP operation.

## Massive online issuance

    request := facturacion.SolicitudRecepcionMasiva{
        SolicitudRecepcionFactura: facturacion.SolicitudRecepcionFactura{
            CodigoModalidad: facturacion.ModalidadComputarizada,
            Cuis:            cuis,
        },
    }
    request, err := request.PrepararMasiva(ctx, documents, nil)
    if err != nil {
        return err
    }
    response, err := client.Facturacion.RecepcionMasivaFactura(ctx, request)

Use a signer for electronic mode. Documents in one batch must satisfy the protocol's sector and modality requirements.

## Contingency package

A package uses offline emission and must include the significant-event code.

    request := facturacion.SolicitudRecepcionPaquete{
        SolicitudRecepcionFactura: facturacion.SolicitudRecepcionFactura{
            CodigoModalidad: facturacion.ModalidadComputarizada,
            Cuis:            cuis,
        },
        CodigoEvento: eventCode,
    }
    request, err := request.PrepararPaquete(ctx, documents, nil)
    if err != nil {
        return err
    }
    response, err := client.Facturacion.RecepcionPaqueteFactura(ctx, request)

The preparation method creates the TAR.GZ, records the document count, and calculates the hash. For electronic mode, pass the XMLSigner.

## Validate processing

A receipt code means the SIN received the package; it does not by itself mean all invoices were accepted. Use ValidacionRecepcionPaqueteFactura or ValidacionRecepcionMasivaFactura with the operation metadata and CodigoRecepcion returned by the initial response. Check the transaction and messages in the validation response.

Do not resend a package just because its final validation is pending. Keep the response receipt code and query its processing state.

See the [API reference](../reference/api.md) and [signing guide](sign-invoices.md).
