# Migrate to go-siat v3

Use `github.com/ron86i/go-siat/v3` and create the client with `siat.New(config)`.
Set the pilot `BaseURL` to `https://pilotosiatservicios.impuestos.gob.bo/sfd`.
This development branch does not publish `v3.0.0` automatically.

All existing service accessors remain available. There is no parallel `NewV3`
facade. Supply the production v3 endpoint assigned by SIN explicitly; the
previous production host with `/sfd/FacturacionCodigos?wsdl` returned HTTP 404
during the 2026-10-18 check.

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
