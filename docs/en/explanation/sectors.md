# Fiscal sectors

Fiscal document models live in documents/&lt;sector&gt;. The package name identifies the document family; use that package's public Go types and constructor.

## Choosing a document

1. Find the sector and document family in the [document reference](../reference/documents-v3.md).
2. Import the matching documents package.
3. Construct the public document model with the types and constructor documented by that package.
4. Submit it through the SIAT service client that supports that document.

For example, Compra-Venta documents use documents/compraventa and the facturation service client. The runnable [Compra-Venta example](../../../examples/compraventa/main.go) shows the full code, CUIS/CUFD and submission flow.

## Service selection

siat.NewClient returns a client with service fields such as Facturacion, Electronica, Computarizada, ServicioBasico, DocumentoAjuste, Telecomunicaciones, EntidadFinanciera, and BoletoAereo. The sector reference and each package's documentation identify the applicable client.

Do not infer that two document packages share the same XML contract because they have similar fields. Each sector package owns its own model and serialization rules.
