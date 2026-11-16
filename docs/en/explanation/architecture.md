# Architecture

The public v3 API separates fiscal documents from SIAT service calls.

## Fiscal documents

Packages under documents/&lt;sector&gt; contain the Go model for a fiscal document family. For example, documents/compraventa contains Compra-Venta. Each package documents its constructors, validation, and XML serialization rules.

Document packages do not calculate tax totals. The caller supplies fiscal values according to its business rules.

## SIAT clients

The root package exposes siat.NewClient and siat.ClientConfig. The returned client groups service clients:

- Codigos: CUIS, CUFD, and code-related operations.
- Sincronizacion and Operaciones: catalogs and point-of-sale or event operations.
- Facturacion and the specialized facturation clients: invoice reception and related operations.
- Compras: receipt of purchase documents.

The protocols subpackages define each service's typed requests and responses. SOAP envelopes and transport details stay inside the SDK.

## Document submission

For supported document types, the facturation client validates and serializes the document, signs it when electronic mode requires a signer, and prepares the compressed request data. The operation then sends the prepared typed request.

    request := facturacion.SolicitudRecepcionFactura{
        CodigoModalidad: facturacion.ModalidadComputarizada,
        Cuis:            cuis,
    }
    request, err := client.Facturacion.PrepararFactura(request, invoice, nil)
    if err != nil {
        return err
    }
    response, err := client.Facturacion.RecepcionFactura(ctx, request)

See [send an invoice](../how-to/send-invoices.md) for the complete flow.

## Boundaries

- documents and protocols are public packages.
- internal/pipeline and internal/soapclient implement SDK internals and are not importable by downstream modules.
- The SDK transports requests; the caller remains responsible for its commercial data and business-specific tax calculations.
