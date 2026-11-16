# First SIAT 3.0 invoice

This tutorial follows the current v3 API. The full runnable Compra-Venta sample is in [examples/compraventa/main.go](../../../examples/compraventa/main.go).

## 1. Configure the client

Use the official pilot endpoint while developing and keep the delegated token outside source control.

    client, err := siat.NewClient(siat.ClientConfig{
        BaseURL:        "https://pilotosiatservicios.impuestos.gob.bo/sfd",
        Token:          os.Getenv("SIAT_TOKEN"),
        Nit:            nit,
        CodigoSistema:  os.Getenv("SIAT_CODIGO_SISTEMA"),
        CodigoAmbiente: 2,
    })
    if err != nil {
        return err
    }

## 2. Obtain the active codes

The codes service receives typed request structs.

    cuis, err := client.Codigos.SolicitarCUIS(ctx, codigos.SolicitudCUIS{
        CodigoModalidad: facturacion.ModalidadComputarizada,
        CodigoSucursal:  sucursal,
        CodigoPuntoVenta: &puntoVenta,
    })
    if err != nil {
        return err
    }

Use the returned CUIS to request a CUFD, then use the CUFD and its control code to create the invoice's CUF.

## 3. Create a sector document

Import the correct package, such as documents/compraventa. Fill the fiscal header and detail structs from your validated business data, then call compraventa.New(header, details). It validates the document and returns an error if required values are missing or inconsistent.

The SDK does not calculate your taxes or decide your invoice amounts.

## 4. Prepare and submit

    request := facturacion.SolicitudRecepcionFactura{
        CodigoModalidad: facturacion.ModalidadComputarizada,
        Cuis:            cuis.Codigo,
    }
    request, err = client.Facturacion.PrepararFactura(request, invoice, nil)
    if err != nil {
        return err
    }
    response, err := client.Facturacion.RecepcionFactura(ctx, request)
    if err != nil {
        return err
    }

Check the transaction state, response code, receipt code, and messages. A successful SOAP exchange is not by itself proof of fiscal acceptance.

For electronic modality, create an XMLSigner and pass it to PrepararFactura. See [signing](../how-to/sign-invoices.md), [configuration](../reference/configuration.md), and the [document reference](../reference/documents-v3.md).
