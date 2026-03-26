<h1 align="center">
  <img src="./.github/logo.svg" alt="go-siat logo" width="300">
  <br>
  <a href="https://github.com/ron86i/go-siat">
    <img src="https://img.shields.io/badge/status-active-success?style=flat-square" alt="Status">
  </a>
  <a href="https://go.dev/">
    <img src="https://img.shields.io/badge/go-1.18+-00ADD8?style=flat-square" alt="Go Version">
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
5. [Project Architecture](#-project-architecture)
6. [Documentation Reference](#-documentation-reference)
7. [Usage Reference (Tests)](#-usage-reference-tests)
8. [Contribution and Support](#-contribution-and-support)
9. [License](#-license)

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

---

## 🚀 Quick Start Guide

### 1. Requirements

- Go 1.25 or higher.
- Valid digital certificate (p12/pfx) and private key (for Electronic modality).

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

## 🏗️ Project Architecture

The project follows a modular architecture based on **Ports and Adapters (Hexagonal)** to ensure maintainability and testability:

- **`internal/core/domain/`**: Pure business logic and SIAT data structures.
- **`internal/core/port/`**: Definition of interfaces (contracts).
- **`internal/adapter/services/`**: Implementation of SOAP clients and communication.
- **`pkg/builders/`**: Fluent entities for building complex XML/JSON (Invoices).
- **`pkg/utils/`**: Utilities for signatures, compression, and formatting.

---

## 📂 Documentation Reference

For detailed guides and implementation details, see:
- [Architecture Overview](ARCHITECTURE.md)
- [Best Practices](CONTEXT_BEST_PRACTICES.md)

---

## 🛠️ Usage Reference (Tests)

The best way to learn how to use each service is by reviewing the integration tests:
- `siat_test.go`: End-to-end flows.
- `pkg/models/invoices/`: Specific examples for each of the 35 sectors.

---

## 🤝 Contribution and Support

Contributions are welcome! If you find a bug or have a suggestion, please open an **Issue** or a **Pull Request**.

---

## 📄 License

This project is licensed under the **MIT License**. See the `LICENSE` file for details.
