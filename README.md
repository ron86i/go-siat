<p align="center">
  <img src="./.github/logo.svg" alt="go-siat logo" width="300">
</p>

<p align="center">
  <a href="https://masterminds.github.io/stability/active.html"><img src="https://masterminds.github.io/stability/active.svg" alt="Stability: Active"></a>
  <br>
  <a href="https://pkg.go.dev/github.com/ron86i/go-siat/v2"><img src="https://pkg.go.dev/badge/github.com/ron86i/go-siat/v2.svg" alt="Go Reference"></a>
  <a href="https://go.dev/"><img src="https://img.shields.io/github/go-mod/go-version/ron86i/go-siat?style=flat" alt="Go Version"></a>
  <a href="LICENSE"><img src="https://img.shields.io/github/license/ron86i/go-siat?style=flat" alt="License"></a>
  <a href="https://github.com/ron86i/go-siat/releases"><img src="https://img.shields.io/github/v/release/ron86i/go-siat?style=flat&label=release" alt="Latest Release"></a>
  <a href="https://deepwiki.com/ron86i/go-siat"><img src="https://deepwiki.com/badge.svg" alt="Ask DeepWiki"></a>
  <a href="i18n/es/README.md"><img src="https://img.shields.io/badge/lang-español-blue?style=flat" alt="Spanish Version"></a>
</p>

<p align="center">
  <em><b>go-siat</b> is a professional SDK developed in Go, designed to simplify integration with <b>SIAT (Integrated Tax Administration System)</b> SOAP web services.</em>
</p>

> [!IMPORTANT]
> **Full Documentation**: [English](docs/en/README.md) | [Español](docs/es/README.md)
>
> Architecture, API Reference, Invoicing Guide, Error Codes, Utilities, and Configuration.

## 💡 Why go-siat?

Integrating with SIAT's SOAP web services for electronic invoicing in Bolivia is often a complex process involving manual XML handling, digital signatures (XMLDSig), and error-prone nested data structures.

**go-siat** abstracts all this complexity behind a modern, idiomatic, and type-safe SDK. Our goal is to allow Bolivian developers to focus on the business logic of their applications (POS, ERPs), while the SDK handles:

- Building perfect SOAP envelopes.
- Digital signatures required by the tax authority.
- Compression and encoding of invoice packages.
- Structured management of catalogs and operations.

---

## 🎯 Features

- 🛡️ **Type-Safe**: Rigorous data structures for ALL requests and responses (goodbye to generic maps and hardcoded strings).
- 🏗️ **Builder Pattern**: Intuitive construction of complex requests (such as invoices and cancellations) through fluid interfaces.
- 📦 **Total SOAP Abstraction**: Transparent management of the SOAP layer. The developer interacts with structs, not XML.
- ✍️ **Integrated Digital Signature (XMLDSig)**: Utilities to automatically sign invoices with your digital certificate.
- 🚀 **High Performance**: Zero unnecessary dependencies, leveraging Go's native speed for byte manipulation and compression.
- 🧩 **Modular**: Multiple services (`Codes`, `Synchronization`, `Operations`, `Sales`, `Electronic`, `Computerized`) clearly separated.
- 🏢 **Multi-Sector**: Native and verified support for **48 different sectors** (Sales, Hotels, Mining, Hospitals, Hydrocarbons, etc.).

---

## Implemented Capabilities

The SDK covers the critical services of the SIAT ecosystem:

| Services                 | Key Functionalities                                                                       |
| :----------------------- | :---------------------------------------------------------------------------------------- |
| **Codes**                | CUIS/CUFD Request (Individual and Massive), NIT Validation, Communication.                |
| **Synchronization**      | Catalogs for activities, parametric, products, services, and sector documents.            |
| **Operations**           | POS Registration/Closing, Significant Event Management.                                   |
| **Sales**                | Specific service for sales, bonuses, and fees.                                            |
| **Adjustment Documents** | Management of Credit/Debit notes, Conciliation, and Reversions.                           |
| **Online Electronic**    | Full support for invoicing with digital signature.                                        |
| **Online Computerized**  | Support for modalities without digital signature.                                         |
| **Specialized Sectors**  | Services for Telecommunications, Basic Services, Financial Entities, and Airline Tickets. |
| **Special Sectors**      | Builders for **50 of SIAT's 51 regulatory sectors**, each verified against the pilot server.                               |

---

## 🚀 Quick Start Guide

### Installation

```bash
go get github.com/ron86i/go-siat/v2
```

### Requirements

- Go 1.26 or higher.
- Valid digital certificate (p12/pfx) and private key (for Electronic modality).

> [!TIP]
> **Context Best Practices**: Always provide a context with a timeout (e.g., 30s) to all SDK calls. Avoid using `context.Background()` directly to prevent hanging requests if the SIAT server is slow.

### Example: Requesting a CUIS

Every SIAT flow starts by requesting a **CUIS** (System Start Code), so this is the shortest end-to-end call you can make.

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/ron86i/go-siat/v2"
    "github.com/ron86i/go-siat/v2/pkg/models"
)

func main() {
    // 1. Configure the client once. Taxpayer identity lives here,
    //    so you never repeat it on every call.
    s, err := siat.New(siat.Config{
        Token:          "YOUR_SIAT_TOKEN",
        Nit:            123456789,
        CodigoSistema:  "YOUR_SYSTEM_CODE",
        CodigoAmbiente: siat.AmbientePruebas,
        BaseURL:        "https://pilotosiatservicios.impuestos.gob.bo/v2",
    })
    if err != nil {
        log.Fatal(err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    // 2. Build the request. Only per-request fields are needed.
    req := models.NewCuisBuilder().
        WithCodigoSucursal(0).
        WithCodigoPuntoVenta(0).
        WithCodigoModalidad(siat.ModalidadElectronica).
        Build()

    // 3. Call the service, then verify SIAT accepted it.
    resp, err := s.Codigos().SolicitudCuis(ctx, req)
    if err != nil {
        log.Fatal(err) // network / transport failure
    }
    if err := siat.Verify(resp.Body.Content.RespuestaCuis); err != nil {
        log.Fatal(err) // SIAT rejected the request
    }

    fmt.Println("CUIS:", resp.Body.Content.RespuestaCuis.Codigo)
}
```

> For advanced guides (full invoicing flow, digital signatures, error handling, middleware), see the [**Full Documentation**](docs/en/README.md).

---

## 🛠️ Usage Reference (Tests)

The **Integration Tests** serve as living technical documentation for the SDK:

| Category                 | Test File                                                                                                      |
| :----------------------- | :------------------------------------------------------------------------------------------------------------- |
| **Codes**                | [`siat_codigos_service_test.go`](internal/adapter/services/siat_codigos_service_test.go)                       |
| **Synchronization**      | [`siat_sincronizacion_service_test.go`](internal/adapter/services/siat_sincronizacion_service_test.go)         |
| **Operations**           | [`siat_operaciones_service_test.go`](internal/adapter/services/siat_operaciones_service_test.go)               |
| **Invoicing (Unified)**  | [`siat_facturacion_service_test.go`](internal/adapter/services/siat_facturacion_service_test.go)               |
| **Adjustment Documents** | [`siat_documento_ajuste_service_test.go`](internal/adapter/services/siat_documento_ajuste_service_test.go)     |
| **Airline Tickets**      | [`siat_boleto_aereo_service_test.go`](internal/adapter/services/siat_boleto_aereo_service_test.go)             |
| **Invoicing (Sectors)**  | [`pkg/models/invoices/`](pkg/models/invoices/)                                                                 |
| **End-to-End**           | [`siat_test.go`](siat_test.go)                                                                                 |

---

## 👍 Contribution and Support

Contributions are welcome! If you find a bug or have a suggestion, please open an **Issue** or a **Pull Request** (please review the [`CONTRIBUTING.md`](.github/CONTRIBUTING.md)).

If you need technical help or commercial support for integrating electronic invoicing in your company, please check our [`SUPPORT.md`](.github/SUPPORT.md).

---

## 📄 License

This project is licensed under the **MIT License**. See the [`LICENSE`](LICENSE) file for details.
