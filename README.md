<p align="center">
  <a href="https://masterminds.github.io/stability/active.html"><img src="https://masterminds.github.io/stability/active.svg" alt="Stability: Active"></a>
  <br>
  <a href="https://go.dev/"><img src="https://img.shields.io/badge/Go-1.26.0-00ADD8?logo=go&logoColor=white" alt="Go 1.26.0"></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/license-MIT-555555?style=flat" alt="MIT license"></a>
  <a href="i18n/es/README.md"><img src="https://img.shields.io/badge/lang-español-blue?style=flat" alt="Spanish Version"></a>
</p>

<p align="center">
  <em><b>go-siat</b> is a professional, high-performance Go SDK designed to simplify integration with Bolivia's <b>SIAT (Sistema de Facturación Virtual / Servicio de Impuestos Nacionales)</b> SOAP web services.</em>
</p>

> [!IMPORTANT]
> **Complete Documentation**: [Docs Index (ES)](docs/es/README.md) | [Docs Index (EN)](docs/en/README.md)
>
> Executable examples are available in [`examples/`](examples/).

---

## 💡 What is go-siat?

**go-siat** represents a redesign of the public SDK API for idiomatic Go, with typed document models and SIAT 3.0 services:

- **Typed v3 API**: Fiscal documents are declared with sector-specific Go structs, and SIAT operations use typed request and response structs.
- **SIAT 3.0**: Uses the `/sfd/` services and net-amount document schemas.
- **Sector Coverage**: 47 fiscal-document families under `documents/<sector>`.
- **End-to-End Emission Pipeline**: Automated document validation, XMLDSig digital signature, GZIP and TAR.GZ compression, SHA-256 hashing, and Base64 SOAP envelope transport.
- **Thread-Safe & Multitenant**: Zero global mutable state. Taxpayer configuration is immutable per client instance, making it fully safe for high-concurrency and multitenant POS/ERP environments (e.g. FALMUS).
- **Clean Codebase**: Zero testdata fixtures or bloat committed to Git. All tests validate in-memory.

---

## 🎯 Architecture

```text
go-siat/
├── siat.go                   # Root client constructor: NewClient
├── documents/                # 47 SIAT 3.0 sector document packages
├── protocols/                # SIAT 3.0 SOAP clients (/sfd/)
│   ├── codigos/
│   ├── facturacion/
│   ├── operaciones/
├── internal/
│   ├── pipeline/             # Validate -> SignXML -> Gzip/TarGz -> SHA-256 hash
│   └── soapclient/           # Isolated HTTP/SOAP client with context, timeouts & faults
├── pkg/utils/                # CUF calculation, XMLDSig digital signer, encoding helpers
└── examples/
    └── compraventa/          # Working SIAT 3.0 emission example
```

---

## 🚀 Quick Start (SIAT 3.0)

### 1. Installation

```bash
go get github.com/ron86i/go-siat/v3
```


> **Requirements**: Go 1.26 or higher.

### 2. Complete Emission Flow

The following example demonstrates obtaining codes, generating the official CUF, building a declarative SIAT 3.0 invoice, and submitting it to the SIN:

```go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ron86i/go-siat/v3"
	"github.com/ron86i/go-siat/v3/documents/compraventa"
	"github.com/ron86i/go-siat/v3/pkg/utils"
	"github.com/ron86i/go-siat/v3/protocols/codigos"
	"github.com/ron86i/go-siat/v3/protocols/facturacion"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// 1. Initialize the SIAT 3.0 client (/sfd/)
	client, err := siat.NewClient(siat.ClientConfig{
		BaseURL:        "https://pilotosiatservicios.impuestos.gob.bo/sfd",
		Token:          "YOUR_SIAT_DELEGATED_TOKEN",
		Nit:            123456789,
		CodigoSistema:  "YOUR_SYSTEM_CODE",
		CodigoAmbiente: 2, // 1: Production, 2: Pilot / Testing
	})
	if err != nil {
		log.Fatalf("Client init error: %v", err)
	}

	sucursal := 0
	puntoVenta := 0
	modalidad := 2 // 1: Online Electronic (with digital signature), 2: Online Computerized
	numeroFactura := int64(1001)
	fechaEmision := time.Now()

	// 2. Request CUIS & CUFD
	cuisResp, err := client.Codigos.SolicitarCUIS(ctx, codigos.SolicitudCUIS{
		CodigoSucursal:   sucursal,
		CodigoPuntoVenta: &puntoVenta,
	})
	if err != nil {
		log.Fatalf("CUIS error: %v", err)
	}

	cufdResp, err := client.Codigos.SolicitarCUFD(ctx, codigos.SolicitudCUFD{
		CodigoSucursal:   sucursal,
		CodigoPuntoVenta: &puntoVenta,
		Cuis:             cuisResp.Codigo,
	})
	if err != nil {
		log.Fatalf("CUFD error: %v", err)
	}

	// 3. Generate Official CUF (Código Único de Factura)
	cuf, err := utils.CUFParams{
		Nit:                 123456789,
		FechaHora:           fechaEmision,
		Sucursal:            sucursal,
		Modalidad:           modalidad,
		TipoEmision:         1, // Online
		TipoFactura:         1, // With tax credit right
		TipoDocumentoSector: 1, // Sector 1: Compra-Venta
		NumeroFactura:       numeroFactura,
		PuntoVenta:          puntoVenta,
		CodigoControl:       cufdResp.CodigoControl,
	}.Generate()
	if err != nil {
		log.Fatalf("CUF generation error: %v", err)
	}

	// 4. Create and validate the invoice with compraventa.New()
	cliente := "JUAN PEREZ"
	cabecera := compraventa.Cabecera{
		NitEmisor:                    123456789,
		RazonSocialEmisor:            "MI EMPRESA S.R.L.",
		Municipio:                    "La Paz",
		NumeroFactura:                numeroFactura,
		Cuf:                          cuf,
		Cufd:                         cufdResp.Codigo,
		CodigoSucursal:               sucursal,
		Direccion:                    "Av. 6 de Agosto Nro 1234",
		CodigoPuntoVenta:             &puntoVenta,
		FechaEmision:                 fechaEmision,
		NombreRazonSocial:            &cliente,
		CodigoTipoDocumentoIdentidad: 1, // C.I.
		NumeroDocumento:              "1234567",
		CodigoCliente:                "CLI-001",
		CodigoMetodoPago:             1, // Cash
		// SIAT 3.0 uses net pricing fields:
		SubTotalNeto:                 100.00,
		TotalPrecioNeto:              100.00,
		Iva:                          13.00,
		MontoAPagar:                  100.00,
		CodigoMoneda:                 1, // BOB
		TipoCambio:                   1.0,
		MontoTotalMoneda:             100.00,
		Leyenda:                      "Ley N° 453: Tienes derecho a un trato equitativo.",
		Usuario:                      "operador",
		CodigoInterno:                "FAC-1001",
		CodigoDocumentoSector:        1,
	}
	detalles := []compraventa.Detalle{
		{
			ActividadEconomica: "620100",
			CodigoProductoSin:  83141,
			CodigoProducto:     "SERV-01",
			Descripcion:        "Servicios de Consultoría de Software",
			Cantidad:           1.0,
			UnidadMedida:       58, // Services unit
			PrecioNetoUnitario: 100.00,
			SubTotal:           100.00,
		},
	}

	factura, err := compraventa.New(cabecera, detalles)
	if err != nil {
		log.Fatalf("Invoice creation error: %v", err)
	}

	// 5. Automatic Pipeline: Validate, Compress (GZIP) & Hash (SHA-256)
	solicitudBase := facturacion.SolicitudRecepcionFactura{
		CodigoModalidad: modalidad, // 1: Electronic, 2: Computerized
		Cuis:            cuisResp.Codigo,
	}
	solicitud, err := solicitudBase.PrepararFactura(factura, nil) // Pass an XMLSigner for electronic mode
	if err != nil {
		log.Fatalf("Invoice preparation error: %v", err)
	}

	// 6. Send to the SIAT service
	resp, err := client.Facturacion.RecepcionFactura(ctx, solicitud)
	if err != nil {
		log.Fatalf("SOAP error: %v", err)
	}

	fmt.Printf("Status: %s (Code: %d, Reception: %s)\n",
		resp.CodigoDescripcion, resp.CodigoEstado, resp.CodigoRecepcion)
}
```

---

For SIAT v2 (`/v2/`), use the separate [`github.com/ron86i/go-siat/v2`](https://pkg.go.dev/github.com/ron86i/go-siat/v2) module.

---

## ✍️ Digital Signature (XMLDSig)

For **Online Electronic Modality** (`Modalidad = 1`), invoices must be signed digitally using a valid PKCS#12 (`.p12` / `.pfx`) or PEM certificate before transmission.

The SDK includes a high-performance, thread-safe XMLDSig signer in `pkg/utils`:

```go
// From PEM files or byte arrays:
signer, err := utils.NewXMLDocumentSigner(privateKeyPEM, certPEM)

// Or directly from PKCS#12 (.p12):
signer, err := utils.NewXMLDocumentSignerFromP12(p12Bytes, "certificate_password")
if err != nil {
    log.Fatal(err)
}

// Pass the signer while preparing an electronic invoice:
solicitud, err := solicitudBase.PrepararFactura(factura, signer)
```

---

## 📦 Supported Sectors (47 Document Families)

The 47 fiscal-document families supported by this module are implemented under `documents/`:

| Sector | Package | Description |
| :---: | :--- | :--- |
| **1** | `documents/compraventa` | Factura de Compra y Venta |
| **2** | `documents/alquileres` | Factura de Alquiler de Bienes Inmuebles |
| **3** | `documents/comercialexportacion` | Factura Comercial de Exportación |
| **4** | `documents/libreconsignacion` | Factura Comercial de Exportación en Libre Consignación |
| **5** | `documents/zonafranca` | Factura de Zona Franca |
| **6** | `documents/turismohospedaje` | Factura de Servicio Turístico y Hospedaje |
| **7** | `documents/seguridadalimentaria` | Factura de Seguridad Alimentaria y Abastecimiento |
| **8** | `documents/tasacero` | Factura Tasa Cero (Libros y Transporte Internacional) |
| **9** | `documents/monedaextranjera` | Factura de Compra y Venta de Moneda Extranjera |
| **10** | `documents/duttyfree` | Factura Dutty Free |
| **11** | `documents/sectoreducativo` | Factura de Sectores Educativos |
| **13** | `documents/serviciosbasicos` | Factura de Servicios Básicos |
| **14** | `documents/alcanzadaice` | Factura Productos Alcanzados por ICE |
| **15** | `documents/entidadesfinancieras` | Factura de Entidades Financieras |
| **16** | `documents/hoteles` | Factura de Hoteles |
| **17** | `documents/hospitales` | Factura de Hospitales y Clínicas |
| **18** | `documents/juegosazar` | Factura de Juegos de Azar |
| **19** | `documents/hidrocarburosiehd` | Factura Hidrocarburos Alcanzada IEHD |
| **20** | `documents/exportacionminerales` | Factura Comercial de Exportación de Minerales |
| **21** | `documents/ventaminerales` | Factura Venta de Minerales |
| **22** | `documents/telecomunicaciones` | Factura de Telecomunicaciones |
| **23** | `documents/prevalorada` | Factura Prevalorada |
| **24** | `documents/notacreditodebito` | Nota de Crédito - Débito |
| **28** | `documents/exportacionservicios` | Factura Comercial de Exportación de Servicios |
| **29** | `documents/notaconciliacion` | Nota de Conciliación |
| **30** | `documents/boletoaereo` | Factura Boleto Aéreo |
| **31** | `documents/suministroenergia` | Factura de Suministro de Energía Eléctrica |
| **34** | `documents/seguros` | Factura de Seguros |
| **35** | `documents/compraventabonificaciones` | Factura Compra Venta con Bonificaciones |
| **36** | `documents/prevaloradasd` | Factura Prevalorada Sin Derecho a Crédito Fiscal |
| **37** | `documents/gnv` | Factura Comercialización de GNV |
| **38** | `documents/hidrocarburosnoiehd` | Factura Hidrocarburos No Alcanzada IEHD |
| **39** | `documents/gnglsp` | Factura Comercialización de GN y GLP |
| **40** | `documents/serviciosbasicoszf` | Factura Servicios Básicos Zona Franca |
| **41** | `documents/vuelocharter` | Factura de Vuelo Chárter |
| **42** | `documents/alquilerzf` | Factura Alquiler Zona Franca |
| **43** | `documents/exportacionhidrocarburos` | Factura Comercial de Exportación de Hidrocarburos |
| **44** | `documents/lubricantes` | Factura Importación y Comercialización de Lubricantes |
| **45** | `documents/exportacionprecioventa` | Factura Comercial de Exportación a Precio de Venta |
| **46** | `documents/sectoreducativozf` | Factura Sector Educativo Zona Franca |
| **48** | `documents/notacreditodebitoice` | Nota de Crédito - Débito Alcanzada por ICE |
| **49** | `documents/telecomunicacioneszf` | Factura Telecomunicaciones Zona Franca |
| **50** | `documents/hospitaleszf` | Factura Hospitales Zona Franca |
| **52** | `documents/ventamineralbcb` | Factura Venta de Minerales BCB |
| **53** | `documents/lubricantesiehd` | Factura Lubricantes Alcanzada IEHD |
| **54** | `documents/biodiesel` | Factura de Biodiésel |
| **55** | `documents/combustiblesinsubvencion` | Factura Venta de Combustible Sin Subvención |

See the [complete sector inventory](docs/es/reference/documentos-v3.md) for detailed struct definitions.

---

## 🛠️ Testing & Verification

Run the entire offline test suite (no network credentials or SIAT connection required):

```bash
# Run all unit tests across all packages
go test ./...

# Verify clean typing and strict formatting
go vet ./...
```

---

## 📄 License

This repository is distributed under the MIT License; see [`LICENSE`](LICENSE).
