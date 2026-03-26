# go-siat Architecture

This document describes the architecture, folder structure, and design patterns of the **go-siat** SDK.

## 🏗️ Overview

**go-siat** implements a **modular architecture inspired by Hexagonal Architecture (Ports and Adapters)** with a clear separation of concerns:

```
┌─────────────────────────────────────────────────────────────┐
│                    User Code                                │
│          (Uses the SDK through SiatServices)                │
└─────────────────────┬───────────────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────────────┐
│              SiatServices (Entry Point)                     │
│   - Operations()                                            │
│   - Synchronization()                                       │
│   - Codes()                                                 │
│   - PurchaseSale()                                          │
│   - Computerized()                                          │
│   - Electronic()                                            │
└─────────────────────┬───────────────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────────────┐
│              Ports (Interfaces)                             │
│   - SiatCodigosService                                      │
│   - SiatSincronizacionService                               │
│   - SiatOperacionesPort                                     │
│   - SiatCompraVentaService                                  │
│   - SiatComputarizadaService                                │
│   - SiatElectronicaService                                  │
└─────────────────────┬───────────────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────────────┐
│            Adapters (Implementations)                       │
│   - SiatCodigosService (HTTP/SOAP)                          │
│   - SiatSincronizacionService (HTTP/SOAP)                   │
│   - SiatOperacionesService (HTTP/SOAP)                      │
│   - ... (other services)                                    │
└─────────────────────┬───────────────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────────────┐
│              Domain Core                                    │
│   - SOAP Models                                             │
│   - Response Structures                                     │
│   - SIAT Data Types                                         │
└─────────────────────┬───────────────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────────────┐
│             Utilities (Utility Layer)                       │
│   - XML Signing (XMLDSig)                                   │
│   - Compression and Hash                                    │
│   - Base64 Encoding                                         │
│   - CUF Generation                                          │
└─────────────────────────────────────────────────────────────┘
```

## 📂 Folder Structure

```
go-siat/
├── README.md                      # Main documentation
├── ARCHITECTURE.md               # This file
├── LICENSE                       # MIT License
├── go.mod                        # Go module
│
├── pkg/                          # SDK Public Code
│   ├── config/                   # Configuration
│   ├── models/                   # User Data Models
│   │   ├── codigos.go           # Builders for codes
│   │   ├── electronica.go       # Builders for electronic invoices
│   │   ├── computarizada.go     # Builders for computerized invoices
│   │   ├── compra_venta.go      # Builders for purchase-sale
│   │   ├── operaciones.go       # Builders for operations
│   │   ├── sincronizacion.go    # Builders for synchronization
│   │   └── invoices/             # 35 specific sectors
│   │       ├── compra_venta.go
│   │       ├── hoteles.go
│   │       ├── mineria.go
│   │       ├── ... (other sectors)
│   │
│   └── utils/                    # Public Utilities
│       ├── crypto.go             # Cryptographic functions
│       ├── cuf.go                # CUF generation
│       ├── encoding.go           # Encoding/Decoding
│       ├── parse.go              # Structure parsing
│       └── signXML.go            # XMLDSig digital signature
│
├── internal/                     # SDK Private Code
│   │
│   ├── adapter/                  # Concrete Implementations (Adapters)
│   │   └── service/              # HTTP/SOAP Services
│   │       ├── siat_codigos_service.go
│   │       ├── siat_electronica_service.go
│   │       ├── siat_operaciones_service.go
│   │       ├── siat_sincronizacion_service.go
│   │       ├── siat_compra_venta_service.go
│   │       └── siat_computarizada_service.go
│   │
│   ├── core/                     # Domain Core
│   │   ├── domain/               # Domain Models (Internal)
│   │   │   ├── datatype/         # Custom SIAT SIAT types
│   │   │   │   └── soap/         # SOAP Envelopes/Faults
│   │   │   └── siat/             # SOAP Request/Response definitions
│   │   │       ├── codigos/
│   │   │       ├── facturacion/
│   │   │       ├── operaciones/
│   │   │       └── sincronizacion/
│   │   │
│   │   └── port/                 # Port Definitions (Interfaces)
│   │       ├── siat_codigos_port.go
│   │       └── ...
│   │
│   └── errors/                   # Internal Error management
```

## 🧩 Key Design Patterns

### 1. **Hexagonal Architecture (Ports and Adapters)**
The SDK defines ports (interfaces) in `internal/core/port` and implements them in `internal/adapter/service`. This allows the communication layer (HTTP/SOAP) to be separated from the business logic, facilitating mock testing and future updates to the underlying protocol.

### 2. **Builder Pattern**
Given the complexity of SIAT data structures (especially in invoices), we use the **Builder Pattern** to provide a fluent and safe API for the user.
- **Location**: `pkg/models/`
- **Example**: `models.Codigos().NewCuisBuilder().WithNit(...).Build()`

### 3. **Request/Response Wrapper (Envelope)**
All SOAP web service calls are encapsulated in a generic `Envelope` that manages the standard SOAP headers and handles `Fault` (SOAP errors) automatically.
- **Location**: `internal/core/domain/datatype/soap/envelope.go`

### 4. **Safe Parsing (Nilable Types)**
SIAT often requires specific treatments for null or omitted values in XML. We use a custom `Nilable[T]` type or pointers to ensure correct serialization to XML.

---

## 🛡️ Principles and Best Practices

1.  **Immutability**: Once an `EnvelopeResponse` is received, it is treated as immutable.
2.  **Context-Aware**: All public methods of the SDK accept a `context.Context` to allow cancellation, timeouts, and tracing.
3.  **No Global State**: The SDK does not use global variables. Each `SiatServices` client is isolated.
4.  **Fail Fast**: We validate critical fields in builders before executing the request to save bandwidth and improve the developer experience.

## 🤝 Next Steps and Contributions

If you want to contribute, please review the [CONTRIBUTING.md](CONTRIBUTING.md) guide. For usage examples, see the `*_test.go` files in `internal/adapter/service/`.
