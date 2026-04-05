<h1 align="center">
  <img src="./.github/logo.svg" alt="go-siat logo" width="300">
  <br>
  <a href="https://github.com/ron86i/go-siat">
    <img src="https://img.shields.io/badge/status-active-success?style=flat-square" alt="Status">
  </a>
  <a href="https://go.dev/">
    <img src="https://img.shields.io/badge/go-1.25+-00ADD8?style=flat-square" alt="Go Version">
  </a>
  <a href="LICENSE">
    <img src="https://img.shields.io/badge/license-MIT-green?style=flat-square" alt="License">
  </a>
  <a href="i18n/es/README.md">
    <img src="https://img.shields.io/badge/lang-español-blue?style=flat-square" alt="Spanish Version">
  </a>
</h1>

<p align="center">
  <em><b>go-siat</b> is a professional SDK developed in Go, designed to simplify integration with <b>SIAT (Integrated Tax Administration System)</b> SOAP web services.</em>
</p>

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
- 🏢 **Multi-Sector**: Native and verified support for **35 different sectors** (Sales, Hotels, Mining, Hospitals, Hydrocarbons, etc.).

---

## 📖 Table of Contents

1. [Why go-siat?](#-why-go-siat)
2. [Features](#-features)
3. [Quick Start Guide](#quick-start-guide)
4. [Advanced Examples](#-advanced-examples)
5. [Usage Reference (Tests)](#-usage-reference-tests)
6. [Contribution and Support](#-contribution-and-support)
7. [License](#-license)

---

## Implemented Capabilities

The SDK covers the critical services of the SIAT ecosystem:

| Services | Key Functionalities |
| :--- | :--- |
| **Codes** | CUIS/CUFD Request (Individual and Massive), NIT Validation, Communication. |
| **Synchronization** | Catalogs for activities, parametric, products, services, and sector documents. |
| **Operations** | POS Registration/Closing, Significant Event Management. |
| **Sales** | Specific service for sales, bonuses, and fees. |
| **Online Electronic** | Full support for invoicing with digital signature. |
| **Online Computerized** | Support for modalities without digital signature. |
| **Special Sectors** | Verified support for the **35 regulatory sectors** of SIAT. |

---

## Supported Sectors

`go-siat` includes domain models, builders, and **integration tests** for the **35 regulatory sectors** of SIAT (located in `pkg/models/invoices/`):

### 🏢 Standard and Services
- **Sales (Sector 1)**: The standard sector for most businesses.
- **Rental of Real Estate**: For the real estate and leasing sector.
- **Insurance**: Issuance of policies and insurance services.
- **Basic Services**: Electricity, water, gas, and telecommunications.
- **Tourism and Lodging / Hotels**: For the hotel sector and tour operators.
- **Hospitals and Clinics**: Health services (National and Free Trade Zone).
- **Food Security**: Commercialization of basic food basket products.

### 🏺 Export and Free Trade Zone
- **Export of Goods and Services**: Commercial Export, Services, and Free Consignment.
- **Free Trade Zone**: ZF Invoices, ZF Rental, and ZF Hospital Services.
- **Duty Free**: Invoicing for duty-free shops in airports.

### ⛽ Hydrocarbons and Energy
- **Commercialization of Hydrocarbons**: Fuels, Lubricants (with and without IEHD).
- **Bottling Plants**: LPG distribution sector.
- **CNG and GNV**: Commercialization of Natural Gas for Vehicles.
- **Unsubsidized Fuel**: For sale at international price.

### ⛰️ Mining and Metals
- **Sale of Minerals**: Internal Sale and Export of Minerals.
- **Sale to BCB**: Sale of gold and minerals to the **Central Bank of Bolivia**.

### 🎓 Education
- **Educational Sectors**: Schools, Universities, and Institutes (National and Free Trade Zone).

### 🔄 Adjustment Documents
- **Credit / Debit Notes**: Standard, ICE, and Fiscal credit/debit notes.
- **Conciliation Notes**: Conciliation notes for billing adjustments.

### 🎲 Other Special Sectors
- **Games of Chance**: Casinos and entertainment venues.
- **Zero Tax (Tasa Cero)**: Books and international cargo transportation.
- **ICE Products**: Items covered by the Specific Consumption Tax.
- **Prepayments and Shared Invoice**: Complex invoicing flows.
- **Prevalued**: Invoices with fixed price and recurring tax service.
- **Foreign Currency Exchange**: Exchange houses and financial entities.

---

## 🚀 Quick Start Guide

### 1. Requirements

- Go 1.25 or higher.
- Valid digital certificate (p12/pfx) and private key (for Electronic modality).

> [!TIP]
> **Context Best Practices**: Always provide a context with a timeout (e.g., 30s) to all SDK calls. Avoid using `context.Background()` directly to prevent hanging requests if the SIAT server is slow.

### 2. Installation

```bash
go get github.com/ron86i/go-siat
```

### 3. Usage Example: Verifying NIT

```go
package main

import (
	"context"
	"fmt"
	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/pkg/models"
)

func main() {
	// 1. Initialize the client
	s, _ := siat.New("YOUR_SIAT_URL", nil)

	// 2. Prepare the request using builders
	req := models.Codigos().NewVerificarNitBuilder().
		WithNit(123456789).
		Build()

	// 3. Execute call
	resp, err := s.Codigos().VerificarNit(context.Background(), req)
	if err != nil {
		panic(err)
	}

	fmt.Printf("NIT Transaction status: %v\n", resp.Body.Content.RespuestaVerificarNit.Transaccion)
}
```

---

---

The best way to learn how to use each service is by reviewing the integration tests:

| Category | Test File |
| :--- | :--- |
| **Codes** | [`siat_codigos_service_test.go`](internal/adapter/services/siat_codigos_service_test.go) |
| **Synchronization** | [`siat_sincronizacion_service_test.go`](internal/adapter/services/siat_sincronizacion_service_test.go) |
| **Operations** | [`siat_operaciones_service_test.go`](internal/adapter/services/siat_operaciones_service_test.go) |
| **Sales** | [`siat_compra_venta_service_test.go`](internal/adapter/services/siat_compra_venta_service_test.go) |
| **Electronic** | [`siat_electronica_service_test.go`](internal/adapter/services/siat_electronica_service_test.go) |
| **Computerized** | [`siat_computarizada_service_test.go`](internal/adapter/services/siat_computarizada_service_test.go) |
| **Adjustment Documents** | [`siat_documento_ajuste_service_test.go`](internal/adapter/services/siat_documento_ajuste_service_test.go) |
| **Invoicing (Sectors)** | [`pkg/models/invoices/`](pkg/models/invoices/) |
| **End-to-End** | [`siat_test.go`](siat_test.go) |
s
---

Contributions are welcome! If you find a bug or have a suggestion, please open an **Issue** or a **Pull Request** (please review the [`CONTRIBUTING.md`](.github/CONTRIBUTING.md)).

If you need technical help or commercial support for integrating electronic invoicing in your company, please check our [`SUPPORT.md`](.github/SUPPORT.md).

---

## 📄 License

This project is licensed under the **MIT License**. See the [`LICENSE`](LICENSE) file for details.
