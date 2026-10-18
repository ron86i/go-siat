# Migrate to go-siat v3

Use `github.com/ron86i/go-siat/v3` and create the client with `siat.New(config)`.
Set the pilot `BaseURL` to `https://pilotosiatservicios.impuestos.gob.bo/sfd`.
This development branch does not publish `v3.0.0` automatically.

All existing service accessors remain available. There is no parallel `NewV3`
facade. Supply the production v3 endpoint assigned by SIN explicitly; the
previous production host with `/sfd/FacturacionCodigos?wsdl` returned HTTP 404
during the 2026-10-18 check.

## Fiscal-document net-amount schemas

The models cover the supplied 47 document families and 93 XSD roots. See the
[document inventory](../reference/documents-v3.md) for builders and modalities.
Each family follows its own field sequence, types, nullability and precision.
Header builders now use `WithSubTotalNeto`,
`WithTotalPrecioNeto`, `WithIva`, `WithMontoAPagar` and `WithCodigoInterno`.
`WithMontoTotal`, `WithMontoTotalSujetoIva` and header `WithCafc` are removed.
Details use `WithPrecioNetoUnitario` instead of `WithPrecioUnitario`.
`WithComplemento("")` sends an empty string; this field is not nillable.
`WithCodigoInterno` takes a string. Decimal precision follows each XSD's
`fractionDigits`: Compra/Venta uses two places, while Vuelo Charter details use five.
Nillable setters take pointers: `nil` emits `xsi:nil="true"`, while a pointer
to zero emits `0`. Credit/debit notes use separate original and returned amounts,
including `WithSubTotalNetoOriginal`, `WithIvaDevuelto` and `WithMontoAPagarDevuelto`.
The SDK
does not calculate taxes or derive totals; callers supply these values.
The required `codigoInterno` must contain the appropriate service value.

`NewVueloCharterBuilder` adds sector 41 with `WithMontoTasa`, serial numbers
and IMEI. Venta Mineral BCB has only an electronic schema in the supplied ZIP.
Free-form JSON setters retain the original `encoding/json` errors.

Run the offline invoice suite with:

```sh
SIAT_RUN=0 go test ./pkg/models/invoices ./internal/core/domain/documents -count=1
```

`TestFacturasV3ContratosXSD` covers all 93 roots, field sequences, values,
precision and nulls. Additional local XSD validation passed for all 46 computerized
formats with populated and null values. Full validation of the 47 electronic
formats requires signed XML and the `SignatureSchema.xsd` dependency missing
from the ZIP. These checks do not establish acceptance by SIAT.

The Dutty Free, Boleto Aereo and Nota Conciliacion integration fixtures also
have an offline construction regression:

```sh
go test ./internal/adapter/services -run '^TestFacturasV3FixturesSinRed$' -count=1
```

Set `SIAT_COMPRA_VENTA_XSD` to a local computerized XSD path and run
`go test ./pkg/models/invoices -run '^TestCompraVentaContratoNeto$' -count=1`
to validate the generated XML using `xmllint`, without contacting SIAT or
copying contracts into the repository. Field/order tests cover both modalities;
full electronic XSD validation additionally needs a signature and the imported
`SignatureSchema.xsd` file.

## Contract and API changes

- Remove `RecepcionPaqueteFacturaBuilder.WithCafc` calls. The v3 SOAP package
  request has no `cafc` element. Individual invoice XML follows its sector XSD.
- `NewConsultaEventoSignificativoBuilder().WithCufd(cufd)` fixes an SDK
  omission: event queries already required CUFD in v2. This is not a new
  requirement introduced by v3.
- CUFD responses expose `CodigoInterno`.
- Product synchronization returns `CodigoActividad` as `string`,
  `CodigoProducto` as `int64`, and the `Nandina` list.
- Facturation reception messages use `MensajeRecepcion`, preserving
  `Advertencia`, `NumeroArchivo` and `NumeroDetalle` in addition to code/text.
- Invoice cancellation serializes `codigoMotivo` before `cuf`.
- Dynamic identity is injected into a copy of the request, preserving builder
  reuse across requests.

## Verification

The existing service integration tests read `SIAT_URL`. Set it to
`https://pilotosiatservicios.impuestos.gob.bo/sfd` before opting into a real
test with `SIAT_RUN=1`. This migration does not modify your local `.env` file.

Run the local HTTP tests with:

```sh
go test . -run '^TestSIATV3_' -count=1
```

External verification of `ServicioRecepcionCompras` remains pending because
the server returned HTTP 503.

Offline tests check local HTTP routing, namespaces, XML bodies and dynamic configuration. They do not
validate fiscal values or demonstrate official SIN acceptance.

### Real code-service tests

CUIS/CUFD tests require `SIAT_RUN=1`, environment 2 and a `/sfd` base URL.
`SIAT_PV` defaults to 1 and `SIAT_SUCURSAL` to 0. Their longer aliases
`SIAT_CODIGO_PUNTO_VENTA` and `SIAT_CODIGO_SUCURSAL` are also supported;
the short names take precedence.

```sh
SIAT_RUN=1 SIAT_PV=1 go test -v ./internal/adapter/services \
  -timeout 30s -run '^TestSolicitudCuis$' -count=1
```

Full requests and responses are stored under `.siat-traces/<test name>`.
Dependent tests can reuse `SIAT_CUIS`. `TestSolicitudCuis` always requests a
code, while `TestSolicitudCufdAll` obtains CUIS once for `SIAT_VECES` CUFD
requests (10 by default) and stops on the first rejection. The CUFD batch test
calls the actual batch operation. Revocation notification requires a test
certificate file configured through `SIAT_CERT_REVOCADO`.
