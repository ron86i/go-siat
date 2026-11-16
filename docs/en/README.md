# Go SIAT v3 documentation

This module targets the Bolivian SIN SIAT 3.0 SOAP services under the /sfd endpoint. Start with the current v3 client and typed document and service APIs.

## Start here

- [First invoice tutorial](tutorial/first-invoice.md)
- [Send an invoice](how-to/send-invoices.md)
- [Sign electronic invoices](how-to/sign-invoices.md)
- [Send batches and contingency packages](how-to/send-batches.md)
- [Configuration](reference/configuration.md)
- [API reference](reference/api.md)
- [Fiscal document packages](reference/documents-v3.md)
- [Sector overview](explanation/sectors.md)
- [Architecture](explanation/architecture.md)
- [Runnable Compra-Venta example](../../examples/compraventa/main.go)

## Current API

- Create the client with siat.NewClient and siat.ClientConfig.
- Build fiscal documents with the public structs in documents/&lt;sector&gt;.
- Call SIAT services through typed clients and requests in protocols.
- Use pkg/utils for CUF, XML, signing, and compression helpers when a lower-level operation is needed.

The protocol clients prepare and serialize supported documents before sending SOAP requests. The SIN remains the authority for fiscal acceptance: inspect the returned transaction state, receipt code, and messages.

[Read this documentation in Spanish](../es/README.md).
