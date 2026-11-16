<p align="center">
  <a href="https://masterminds.github.io/stability/active.html"><img src="https://masterminds.github.io/stability/active.svg" alt="Stability: Active"></a>
  <br>
  <a href="https://go.dev/"><img src="https://img.shields.io/badge/Go-1.26.0-00ADD8?logo=go&logoColor=white" alt="Go 1.26.0"></a>
  <a href="../../LICENSE"><img src="https://img.shields.io/badge/licencia-MIT-555555?style=flat" alt="Licencia MIT"></a>
  <a href="../../README.md"><img src="https://img.shields.io/badge/lang-english-red?style=flat" alt="English Version"></a>
</p>

<p align="center">
  <em><b>go-siat</b> es un SDK profesional de alto rendimiento desarrollado en Go, diseñado para simplificar la integración con los servicios web SOAP del <b>SIAT (Servicio de Impuestos Nacionales de Bolivia)</b>.</em>
</p>

> [!IMPORTANT]
> **Documentación**: [Índice de Docs (ES)](../../docs/es/README.md) | [Docs Index (EN)](../../docs/en/README.md)
>
> Ejemplos ejecutables disponibles en [`examples/`](../../examples/).

---

## 💡 ¿Qué es go-siat?

**go-siat** representa un rediseño integral de la API pública del SDK, enfocado en código idiomático en Go, modelos documentales tipados y servicios SIAT 3.0:

- **API v3 tipada**: Los documentos fiscales se declaran con structs sectoriales de Go y las operaciones SIAT usan solicitudes y respuestas tipadas.
- **SIAT 3.0**: Usa los servicios `/sfd/` y los esquemas de documentos con importes netos.
- **Cobertura de documentos**: 47 familias fiscales bajo `documents/<sector>`.
- **Pipeline de Emisión Automatizado**: Validación de documentos, firma digital XMLDSig, compresión GZIP y TAR.GZ, hash SHA-256 y empaquetado en sobres SOAP Base64 de forma transparente.
- **Seguro para Concurrencia y Multitenant**: Cero estado mutable global. La configuración del contribuyente es inmutable por instancia del cliente, garantizando seguridad en entornos de alta concurrencia y sistemas multi-empresa (por ejemplo, FALMUS).
- **Repositorio Limpio**: Cero archivos o fixtures de prueba en Git (`testdata/` eliminado). Todas las pruebas se validan en memoria.

---

## 🎯 Arquitectura

```text
go-siat/
├── siat.go                   # Constructor raíz: NewClient
├── documents/                # 47 paquetes sectoriales SIAT 3.0
├── protocols/                # Clientes SOAP SIAT 3.0 (/sfd/)
│   ├── codigos/
│   ├── facturacion/
│   ├── operaciones/
├── internal/
│   ├── pipeline/             # Validate -> SignXML -> Gzip/TarGz -> Hash SHA-256
│   └── soapclient/           # Cliente HTTP/SOAP aislado con context, timeouts y fault handling
├── pkg/utils/                # Cálculo de CUF, firmador digital XMLDSig, utilitarios de codificación
└── examples/
    └── compraventa/          # Ejemplo funcional de emisión SIAT 3.0
```

---

## 🚀 Inicio Rápido (SIAT 3.0)

### 1. Instalación

```bash
go get github.com/ron86i/go-siat/v3
```


> **Requisitos**: Go 1.26 o superior.

### 2. Flujo Completo de Emisión

El siguiente ejemplo muestra la obtención de códigos, generación del CUF oficial, construcción declarativa de una factura SIAT 3.0 y su envío al SIN:

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

	// 1. Inicializar el cliente SIAT 3.0 (/sfd/)
	client, err := siat.NewClient(siat.ClientConfig{
		BaseURL:        "https://pilotosiatservicios.impuestos.gob.bo/sfd",
		Token:          "TU_TOKEN_DELEGADO_SIAT",
		Nit:            123456789,
		CodigoSistema:  "TU_CODIGO_DE_SISTEMA",
		CodigoAmbiente: 2, // 1: Producción, 2: Piloto / Pruebas
	})
	if err != nil {
		log.Fatalf("Error inicializando cliente: %v", err)
	}

	sucursal := 0
	puntoVenta := 0
	modalidad := 2 // 1: Electrónica en Línea (con firma digital), 2: Computarizada en Línea
	numeroFactura := int64(1001)
	fechaEmision := time.Now()

	// 2. Solicitar CUIS y CUFD
	cuisResp, err := client.Codigos.SolicitarCUIS(ctx, codigos.SolicitudCUIS{
		CodigoSucursal:   sucursal,
		CodigoPuntoVenta: &puntoVenta,
	})
	if err != nil {
		log.Fatalf("Error solicitando CUIS: %v", err)
	}

	cufdResp, err := client.Codigos.SolicitarCUFD(ctx, codigos.SolicitudCUFD{
		CodigoSucursal:   sucursal,
		CodigoPuntoVenta: &puntoVenta,
		Cuis:             cuisResp.Codigo,
	})
	if err != nil {
		log.Fatalf("Error solicitando CUFD: %v", err)
	}

	// 3. Generar el CUF Oficial (Código Único de Factura)
	cuf, err := utils.CUFParams{
		Nit:                 123456789,
		FechaHora:           fechaEmision,
		Sucursal:            sucursal,
		Modalidad:           modalidad,
		TipoEmision:         1, // En línea
		TipoFactura:         1, // Con derecho a crédito fiscal
		TipoDocumentoSector: 1, // Sector 1: Factura Compra y Venta
		NumeroFactura:       numeroFactura,
		PuntoVenta:          puntoVenta,
		CodigoControl:       cufdResp.CodigoControl,
	}.Generate()
	if err != nil {
		log.Fatalf("Error generando CUF: %v", err)
	}

	// 4. Crear y validar la factura con compraventa.New()
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
		CodigoMetodoPago:             1, // Efectivo
		// SIAT 3.0 utiliza campos de precio neto:
		SubTotalNeto:                 100.00,
		TotalPrecioNeto:              100.00,
		Iva:                          13.00,
		MontoAPagar:                  100.00,
		CodigoMoneda:                 1, // BOB (Bolivianos)
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
			UnidadMedida:       58, // Unidad de servicios
			PrecioNetoUnitario: 100.00,
			SubTotal:           100.00,
		},
	}

	factura, err := compraventa.New(cabecera, detalles)
	if err != nil {
		log.Fatalf("Error construyendo factura: %v", err)
	}

	// 5. Pipeline Automatizado: Validación, Compresión (GZIP) y Hash (SHA-256)
	solicitudBase := facturacion.SolicitudRecepcionFactura{
		CodigoModalidad: modalidad, // 1: Electrónica, 2: Computarizada
		Cuis:            cuisResp.Codigo,
	}
	solicitud, err := solicitudBase.PrepararFactura(factura, nil) // En modo electrónico, pasar el firmador
	if err != nil {
		log.Fatalf("Error en preparación de factura: %v", err)
	}

	// 6. Envío al servicio web del SIAT
	resp, err := client.Facturacion.RecepcionFactura(ctx, solicitud)
	if err != nil {
		log.Fatalf("Error SOAP: %v", err)
	}

	fmt.Printf("Estado: %s (Código: %d, Recepción: %s)\n",
		resp.CodigoDescripcion, resp.CodigoEstado, resp.CodigoRecepcion)
}
```

---

Para SIAT v2 (`/v2/`), utilice el módulo separado [`github.com/ron86i/go-siat/v2`](https://pkg.go.dev/github.com/ron86i/go-siat/v2).

---

## ✍️ Firma Digital (XMLDSig)

Para la **Modalidad Electrónica en Línea** (`Modalidad = 1`), los documentos deben firmarse digitalmente con un certificado PKCS#12 (`.p12` / `.pfx`) o certificados PEM válidos antes de su transmisión.

El SDK incluye un firmador XMLDSig thread-safe de alto rendimiento en `pkg/utils`:

```go
// Desde archivos PEM o cadenas de bytes:
signer, err := utils.NewXMLDocumentSigner(privateKeyPEM, certPEM)

// O directamente desde PKCS#12 (.p12):
signer, err := utils.NewXMLDocumentSignerFromP12(p12Bytes, "password_del_certificado")
if err != nil {
    log.Fatal(err)
}

// Pasar el firmador al preparar una factura electrónica:
solicitud, err := solicitudBase.PrepararFactura(factura, signer)
```

---

## 📦 Sectores Compatibles (47 Familias de Documentos)

Las 47 familias fiscales cubiertas por este módulo están implementadas bajo `documents/`:

| Sector | Paquete | Descripción |
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

Consulte el [inventario completo de sectores](../../docs/es/reference/documentos-v3.md) para más detalles.

---

## 🛠️ Pruebas y Verificación

Ejecute la suite completa de pruebas unitarias locales (sin requerir credenciales ni conexión al SIAT):

```bash
# Ejecutar todas las pruebas unitarias
go test ./...

# Verificar tipado y formato estricto
go vet ./...
```

---

## 📄 Licencia

Este repositorio contiene software propietario. Su uso requiere autorización escrita previa de Falmusbo; consulte [`LICENSE`](../../LICENSE).
