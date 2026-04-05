<h1 align="center">
  <img src="../../.github/logo.svg" alt="go-siat logo" width="250">
  <br>
  go-siat Documentation
</h1>

<p align="center">
  <a href="../es/README.md">
    <img src="https://img.shields.io/badge/lang-español-blue?style=flat-square" alt="Spanish Version">
  </a>
  <a href="https://go.dev/">
    <img src="https://img.shields.io/badge/go-1.25+-00ADD8?style=flat-square" alt="Go Version">
  </a>
  <a href="../../LICENSE">
    <img src="https://img.shields.io/badge/license-MIT-green?style=flat-square" alt="License">
  </a>
</p>

<p align="center">
  <em>Professional SDK for integrating with Bolivia's <b>SIAT (Integrated Tax Administration System)</b> SOAP web services.</em>
</p>

---

## 📚 Documentation Index

Welcome to the **go-siat** technical documentation. This guide is organized to help you go from zero to production with Bolivia's electronic invoicing system.

| # | Document | Description |
|:--|:---------|:------------|
| 1 | [**Architecture**](architecture.md) | Hexagonal design, layers, design patterns, and internal data flow. |
| 2 | [**Getting Started**](getting-started.md) | Installation, prerequisites, environment setup, and your first API call. |
| 3 | [**API Reference**](api-reference.md) | Complete reference for all 7 SIAT services and 67+ methods. |
| 4 | [**Invoicing Guide**](invoicing-guide.md) | Invoice lifecycle, 35 regulatory sectors, digital signatures, and batch processing. |
| 5 | [**Error Handling**](error-handling.md) | Error types, 150+ SIAT error codes, retry strategies, and response verification. |
| 6 | [**Utilities**](utilities.md) | CUF generation, XML signing, compression, hashing, and parsing helpers. |
| 7 | [**Configuration**](configuration.md) | HTTP client tuning, middleware system, distributed tracing, and constants. |

---

## 🗺️ Quick Navigation by Use Case

### "I want to..."

| Goal | Start Here |
|:-----|:-----------|
| Install and make my first call | [Getting Started](getting-started.md) |
| Understand the project architecture | [Architecture](architecture.md) |
| Send an electronic invoice | [Invoicing Guide](invoicing-guide.md) |
| Know which sectors are supported | [Invoicing Guide → Sectors](invoicing-guide.md#supported-sectors-reference) |
| Handle SIAT errors in production | [Error Handling](error-handling.md) |
| Sign invoices with my digital certificate | [Utilities → XML Signing](utilities.md#xml-digital-signing) |
| Customize HTTP timeouts or add logging | [Configuration](configuration.md) |
| Find a specific method signature | [API Reference](api-reference.md) |

---

## 🔗 Additional Resources

| Resource | Location |
|:---------|:---------|
| Root README | [`README.md`](../../README.md) |
| Contributing Guide | [`CONTRIBUTING.md`](../../.github/CONTRIBUTING.md) |
| Support & Consulting | [`SUPPORT.md`](../../.github/SUPPORT.md) |
| Code of Conduct | [`CODE_OF_CONDUCT.md`](../../.github/CODE_OF_CONDUCT.md) |
| Integration Tests | [`internal/adapter/services/*_test.go`](../../internal/adapter/services/) |
| Sector Invoice Tests | [`pkg/models/invoices/*_test.go`](../../pkg/models/invoices/) |
| License (MIT) | [`LICENSE`](../../LICENSE) |

---

<p align="center">
  <sub>Copyright © 2026 Ronaldo Rua — Licensed under MIT</sub>
</p>
