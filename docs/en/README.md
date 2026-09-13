<h1 align="center">
  <img src="../../.github/logo.svg" alt="go-siat logo" width="250">
  <br>
  go-siat Documentation
</h1>

<p align="center">
  <a href="../es/README.md"><img src="https://img.shields.io/badge/lang-español-blue?style=flat" alt="Spanish Version"></a>
  <a href="https://masterminds.github.io/stability/active.html"><img src="https://masterminds.github.io/stability/active.svg" alt="Stability: Active"></a>
  <br>
  <a href="https://pkg.go.dev/github.com/ron86i/go-siat/v2"><img src="https://pkg.go.dev/badge/github.com/ron86i/go-siat/v2.svg" alt="Go Reference"></a>
  <a href="https://go.dev/"><img src="https://img.shields.io/github/go-mod/go-version/ron86i/go-siat?style=flat" alt="Go Version"></a>
  <a href="../../LICENSE"><img src="https://img.shields.io/github/license/ron86i/go-siat?style=flat" alt="License"></a>
  <a href="https://github.com/ron86i/go-siat/releases"><img src="https://img.shields.io/github/v/release/ron86i/go-siat?style=flat&label=release" alt="Latest Release"></a>
  <a href="https://deepwiki.com/ron86i/go-siat"><img src="https://deepwiki.com/badge.svg" alt="Ask DeepWiki"></a>
</p>

<p align="center">
  <em>Professional SDK for integrating with Bolivia's <b>SIAT (Integrated Tax Administration System)</b> SOAP web services.</em>
</p>

---

## Start here

**New to the SDK?** Go straight to the [**tutorial**](tutorial/first-invoice.md). It takes you from `go get` to a SIAT-accepted invoice in seven steps, and everything else will make more sense afterwards.

---

## How this documentation is organized

These docs follow [Diátaxis](https://diataxis.fr/): four kinds of material, each answering a different question. Pick the column that matches what you need right now.

|  | **Practical steps** | **Theoretical knowledge** |
| :--- | :--- | :--- |
| **Learning**<br><sub>when you're new</sub> | 📚 **[Tutorial](tutorial/)**<br>Follow along and build something.<br><br>· [Your first invoice](tutorial/first-invoice.md) | 💡 **[Explanation](explanation/)**<br>Understand why it works this way.<br><br>· [Architecture](explanation/architecture.md)<br>· [Sectors](explanation/sectors.md) |
| **Working**<br><sub>when you have a goal</sub> | 🔧 **[How-to guides](how-to/)**<br>Solve one specific problem.<br><br>· [Send invoices](how-to/send-invoices.md)<br>· [Send batches](how-to/send-batches.md)<br>· [Sign invoices](how-to/sign-invoices.md)<br>· [Handle errors](how-to/handle-errors.md) | 📖 **[Reference](reference/)**<br>Look up exact details.<br><br>· [API](reference/api.md)<br>· [Configuration](reference/configuration.md)<br>· [Utilities](reference/utilities.md) |

---

## Find it by what you want to do

| Goal | Go to |
| :--- | :--- |
| Install the SDK and make my first call | [Tutorial](tutorial/first-invoice.md) |
| Send a single electronic invoice | [Tutorial → Step 6](tutorial/first-invoice.md#step-6-package-and-send) |
| Sign invoices with my `.p12` certificate | [How-to: Sign invoices](how-to/sign-invoices.md) |
| Send hundreds of invoices in one request | [How-to: Send batches](how-to/send-batches.md) |
| Annul an invoice or issue a credit note | [How-to: Send invoices](how-to/send-invoices.md) |
| Find the builder and facade for my industry | [Explanation: Sectors](explanation/sectors.md) |
| Tell a network failure from a SIAT rejection | [How-to: Handle errors](how-to/handle-errors.md) |
| Look up a method signature or a `Config` field | [Reference: API](reference/api.md) |
| Tune HTTP timeouts, pooling or add middleware | [Reference: Configuration](reference/configuration.md) |
| Generate a CUF, compress, hash or sign XML | [Reference: Utilities](reference/utilities.md) |
| Understand the hexagonal design | [Explanation: Architecture](explanation/architecture.md) |

---

## The essentials in 60 seconds

Three things carry most of the SDK's design. If you only remember these, you can read the rest as you go.

**1. Identity is configured once.** Your `Token`, `Nit`, `CodigoSistema`, `CodigoAmbiente` and `BaseURL` live in `siat.Config` and are passed to `siat.New` a single time. Service methods only take `(ctx, req)`.

```go
s, err := siat.New(siat.Config{
    Token:          "...",
    Nit:            123456789,
    CodigoSistema:  "...",
    CodigoAmbiente: siat.AmbientePruebas,
    BaseURL:        "https://pilotosiatservicios.impuestos.gob.bo/v2",
})
```

**2. Every call needs two error checks.** A `nil` error does not mean SIAT accepted anything.

```go
resp, err := s.Codigos().SolicitudCuis(ctx, req)
if err != nil { /* network / transport failed */ }
if err := siat.Verify(resp.Body.Content.RespuestaCuis); err != nil { /* SIAT rejected it */ }
```

**3. Requests are built with builders.** Every request type has a flat constructor in `pkg/models` and is finished with `.Build()`.

```go
req := models.NewCuisBuilder().
    WithCodigoSucursal(0).
    WithCodigoModalidad(siat.ModalidadElectronica).
    Build()
```

---

## Additional resources

| Resource | Location |
| :--- | :--- |
| Root README | [`README.md`](../../README.md) |
| Contributing Guide | [`CONTRIBUTING.md`](../../.github/CONTRIBUTING.md) |
| Support & Consulting | [`SUPPORT.md`](../../.github/SUPPORT.md) |
| Code of Conduct | [`CODE_OF_CONDUCT.md`](../../.github/CODE_OF_CONDUCT.md) |
| Changelog | [`CHANGELOG.md`](../../CHANGELOG.md) |
| Integration tests (living examples) | [`internal/adapter/services/`](../../internal/adapter/services/) |
| Sector invoice tests | [`pkg/models/invoices/`](../../pkg/models/invoices/) |
| License (MIT) | [`LICENSE`](../../LICENSE) |

---

<p align="center">
  <sub>Copyright © 2026 Ronaldo Rua — Licensed under MIT</sub>
</p>
