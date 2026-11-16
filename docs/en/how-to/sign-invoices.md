# Sign electronic invoices

Electronic documents must be signed before submission. The facturation preparation method accepts an XMLSigner and includes the signature in the prepared request. Computerized invoices must not be signed.

## Load a PKCS#12 signer

    p12, err := os.ReadFile(os.Getenv("SIAT_P12_PATH"))
    if err != nil {
        return err
    }
    signer, err := utils.NewXMLDocumentSignerFromP12(
        p12,
        os.Getenv("SIAT_P12_PASSWORD"),
    )
    if err != nil {
        return err
    }

Keep the PKCS#12 file and password in a secret store. Do not log either value.

## Prepare an electronic invoice

The invoice itself is a public sector model, for example documents/compraventa. Set the request modality before preparing it.

    request := facturacion.SolicitudRecepcionFactura{
        CodigoModalidad: facturacion.ModalidadElectronica,
        Cuis:            cuis,
    }
    request, err = client.Facturacion.PrepararFactura(request, invoice, signer)
    if err != nil {
        return err
    }
    response, err := client.Facturacion.RecepcionFactura(ctx, request)

The preparation step validates the document, serializes the sector XML, signs it, compresses it, and calculates the request hash. Submit the returned request, not the original value.

## Verify before submission

Check that the certificate is valid for the intended signing period, the document fields match the CUF, and the request uses electronic modality. See [configuration](../reference/configuration.md) and the [Compra-Venta example](../../../examples/compraventa/main.go).
