# SIAT 3.0 API

The v3 module targets the SIN SIAT 3.0 services under /sfd. Its current API separates fiscal documents from service requests.

## Documents

Import a sector package from documents/&lt;sector&gt;. The model uses public Go types and the package's documented constructor and XML methods. Compra-Venta is available at documents/compraventa; see the [document reference](../reference/documents-v3.md).

## Services

Create the client with siat.NewClient. The returned client groups typed service clients for codes, purchases, facturation, synchronization, and operations.

    client, err := siat.NewClient(siat.ClientConfig{
        BaseURL:        "https://pilotosiatservicios.impuestos.gob.bo/sfd",
        Token:          token,
        Nit:            nit,
        CodigoSistema:  systemCode,
        CodigoAmbiente: 2,
    })

Requests and responses are defined in protocols/&lt;service&gt;. For example, CUIS and CUFD use protocols/codigos; invoice reception uses protocols/facturacion.

## Preparation

For supported fiscal documents, the facturation client validates and serializes the document and prepares the XML, signature, compression, and hash required by the operation. Use the preparation method that matches the operation: PrepararFactura, PrepararMasiva, or PrepararPaquete.

See [configuration](../reference/configuration.md), [sending invoices](send-invoices.md), and [batch submission](send-batches.md).
