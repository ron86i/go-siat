# Utilities

Most callers should use the typed document and protocol APIs. These helpers are available when the application needs to perform a lower-level operation itself.

## CUF

Use CUFParams with named fields to avoid mixing positional numeric values.

    cuf, err := (utils.CUFParams{
        Nit:                 nit,
        FechaHora:           fechaEmision,
        Sucursal:            sucursal,
        Modalidad:           modalidad,
        TipoEmision:         tipoEmision,
        TipoFactura:         tipoFactura,
        TipoDocumentoSector: sector,
        NumeroFactura:       numeroFactura,
        PuntoVenta:          puntoVenta,
        CodigoControl:       codigoControl,
    }).Generate()

The CUF inputs must match the document and the current CUFD. See the [invoice guide](../how-to/send-invoices.md).

## XML and compression

| Function | Purpose |
| --- | --- |
| MarshalIndentedXML | Marshal a Go document to readable XML where the document supports ordinary XML marshaling. |
| Gzip | Compress one byte slice with GZIP. |
| CreateTarGz | Create a TAR.GZ archive from named byte slices. |
| CompressAndHash | Return a SHA-256 hash and Base64-encoded GZIP data. |
| SHA256Hex | Calculate a SHA-256 digest as hexadecimal. |

The facturation preparation methods perform the document-specific serialization and request preparation for supported operations. Prefer those methods over manually rebuilding their payload.

## Signing

NewXMLDocumentSignerFromP12 creates a reusable signer from PKCS#12 bytes. SignXMLBytes and SignWithP12Bytes are lower-level helpers for applications that manage XML themselves.

See the [signing guide](../how-to/sign-invoices.md) and the package documentation for the full exported API.
