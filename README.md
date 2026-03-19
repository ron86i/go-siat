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
</h1>

<p align="center">
  <em><b>go-siat</b> es un SDK profesional desarrollado en Go, diseñado para simplificar la integración con los servicios web SOAP del <b>SIAT (Sistema Integrado de Administración Tributaria)</b>.</em>
</p>

## 💡¿Por qué go-siat?

Integrarse con los servicios web SOAP del SIAT para la facturación electrónica en Bolivia suele ser un proceso complejo que involucra el manejo manual de XML, firmas digitales (XMLDSig) y estructuras de datos anidadas propensas a errores.

**go-siat** abstrae toda esta complejidad detrás de un SDK moderno, idiomático y seguro frente a tipos (type-safe). Nuestro objetivo es permitir a los desarrolladores bolivianos concentrarse en la lógica de negocio de sus aplicaciones (puntos de venta, ERPs), mientras el SDK se encarga de:

- La construcción de sobres SOAP perfectos.
- La firma digital exigida por el fisco.
- La compresión y codificación de paquetes de facturas.
- La gestión estructurada de catálogos y operaciones.

---

## 🎯 Características

- 🛡️ **Type-Safe**: Estructuras de datos rigurosas para TODAS las solicitudes y respuestas (adiós a los mapas genéricos y strings hardcodeados).
- 🏗️ **Builder Pattern**: Construcción intuitiva de solicitudes complejas (como facturas y anulaciones) mediante interfaces fluidas.
- 📦 **Abstracción SOAP Total**: Gestión transparente de la capa SOAP. El desarrollador interactúa con structs, no con XML.
- ✍️ **Firma Digital (XMLDSig) Integrada**: Utilidades para firmar facturas automáticamente con su certificado digital.
- 🚀 **Alto Rendimiento**: Cero dependencias innecesarias, aprovechando la velocidad nativa de Go para la manipulación y compresión de bytes.
- 🧩 **Modular**: Múltiples servicios (`Codigos`, `Sincronizacion`, `Operaciones`, `CompraVenta`, `Computarizada`) claramente separados.
- 🏢 **Multi-Sector**: Soporte nativo para diversos documentos sector (Compra y Venta, Hoteles, Duty Free, Hidrocarburos, etc.).

---

## Tabla de Contenidos

1. [Capacidades Implementadas](#capacidades-implementadas)
2. [Guía de Inicio Rápido](#guía-de-inicio-rápido)
3. [Referencia de Uso (Tests)](#referencia-de-uso-tests)
4. [Licencia](#licencia)

---

## Capacidades Implementadas

El SDK cubre los servicios críticos del ecosistema SIAT:

| Servicios | Funcionalidades Clave |
| :--- | :--- |
| **Códigos** | Solicitud de CUIS/CUFD (Individual y Masivo), Validación de NIT, Comunicación. |
| **Sincronización** | Catálogos de actividades, paramétricas, productos, servicios y documentos sector. |
| **Operaciones** | Registro/Cierre de Puntos de Venta, Gestión de Eventos Significativos. |
| **Compra y Venta** | Generación de CUF, Firma Digital XML, Recepción y Anulación de Facturas. |
| **Computarizada** | Recepción de Facturas y Paquetes para modalidades de facturación computarizada. |

---

## Guía de Inicio Rápido

### Instalación

```bash
go get github.com/ron86i/go-siat
```

### Uso Básico

El siguiente ejemplo demuestra cómo inicializar el cliente y realizar una solicitud de código CUIS:

```go
package main

import (
    "context"
    "log"
    "github.com/ron86i/go-siat"
    "github.com/ron86i/go-siat/pkg/config"
    "github.com/ron86i/go-siat/pkg/models"
)

func main() {
    // 1. Configurar cliente unificado
    s, err := siat.New("https://pilotosiatservicios.impuestos.gob.bo/v2", nil)
    if err != nil {
        log.Fatal("Error al inicializar SDK:", err)
    }

    // 2. Construir solicitud usando el Builder
    req := models.Codigos().NewCuisBuilder().
		WithCodigoAmbiente(1).
		WithCodigoModalidad(1).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithCodigoSistema("ABC123DEF").
		WithNit(123456789).
		Build()

    // 3. Ejecutar operación
    ctx := context.Background()
    cfg := config.Config{Token: "TU_TOKEN_API"}
    resp, err := s.Codigos().SolicitudCuis(ctx, cfg, req)
    if err != nil {
        log.Fatal("Error en la solicitud:", err)
    }
    log.Println("Código CUIS obtenido:", resp.Body.Content.RespuestaCuis.Codigo)
}
```

---

## 👀 Ejemplos Avanzados

A continuación, mostramos algunos de los flujos más comunes. Si desea ver más ejemplos, revise nuestro repositorio de [Tests de Integración](#referencia-de-uso-tests).

<details>
  <summary>📚 Emitir y Enviar una Factura (Flujo Completo)</summary>

Este ejemplo muestra cómo construir una factura, firmarla, prepararla para el SIAT y enviarla.

**📋 Ejemplo: Recepción de Factura**

```go
package main

import (
    "context"
    "encoding/xml"
    "log"
    "time"

    "github.com/ron86i/go-siat"
    "github.com/ron86i/go-siat/pkg/config"
    "github.com/ron86i/go-siat/pkg/models"
    "github.com/ron86i/go-siat/pkg/models/facturas"
    "github.com/ron86i/go-siat/pkg/utils"
)

func main() {
    // 1. Inicializar cliente y credenciales (Asumiendo que ya tiene CUIS y CUFD)
    s, _ := siat.New("https://pilotosiatservicios.impuestos.gob.bo/v2", nil)
    cfg := config.Config{Token: "TU_TOKEN"}
    nit := int64(123456789)
    cufdControl := "CODIGO_CONTROL_CUFD"

    // 2. Generar CUF
    fechaEmision := time.Now()
    cuf, _ := utils.GenerarCUF(nit, fechaEmision, 0, 1, 1, 1, 1, 1, 0, cufdControl)

    // 3. Construir Cabecera y Detalle con el Builder
    nombre := "JUAN PEREZ"
    cabecera := facturas.NewCompraVentaCabeceraBuilder().
        WithNitEmisor(nit).
        WithRazonSocialEmisor("Mi Empresa S.R.L.").
        WithNumeroFactura(1).
        WithCuf(cuf).
        WithCufd("TU_CUFD").
        WithFechaEmision(fechaEmision).
        WithNombreRazonSocial(&nombre).
        WithMontoTotal(100).
        WithCodigoDocumentoSector(1).
        Build()

    detalle := facturas.NewCompraVentaDetalleBuilder().
        WithActividadEconomica("477300").
        WithCodigoProductoSin("622539").
        WithDescripcion("PRODUCTO DEMO").
        WithCantidad(1).
        WithPrecioUnitario(100).
        WithSubTotal(100).
        Build()

    factura := facturas.NewCompraVentaBuilder().
        WithCabecera(cabecera).
        AddDetalle(detalle).
        Build()

    // 4. Serializar, Firmar y Preparar (GZIP -> SHA256 -> Base64)
    xmlData, _ := xml.Marshal(factura)
    signedXML, _ := utils.SignXML(xmlData, "key.pem", "cert.crt")
    hash, archivoBase64, _ := utils.CompressAndHash(signedXML)

    // 5. Enviar al SIAT
    req := models.CompraVenta().NewRecepcionFactura().
        WithCodigoAmbiente(1).
        WithNit(nit).
        WithCufd("TU_CUFD").
        WithCuis("TU_CUIS").
        WithTipoFacturaDocumento(1).
        WithArchivo(archivoBase64).
        WithFechaEnvio(fechaEmision).
        WithHashArchivo(hash).
        Build()

    resp, err := s.CompraVenta().RecepcionFactura(context.Background(), cfg, req)
    if err != nil {
        log.Fatal(err)
    }
    
    log.Printf("Estado de Recepción: %v", resp.Body.Content.RespuestaServicioFacturacion.CodigoEstado)
}
```
</details>

---

## 🛠️ Referencia de Uso (Tests)

Para una comprensión profunda de cada servicio, los **Tests de Integración** actúan como la documentación técnica principal.

| Categoría | Archivo de Test |
| :--- | :--- |
| **Códigos** | [`siat_codigos_service_test.go`](./internal/adapter/service/siat_codigos_service_test.go) |
| **Sincronización** | [`siat_sincronizacion_service_test.go`](./internal/adapter/service/siat_sincronizacion_service_test.go) |
| **Operaciones** | [`siat_operaciones_service_test.go`](./internal/adapter/service/siat_operaciones_service_test.go) |
| **Compra y Venta** | [`siat_compra_venta_service_test.go`](./internal/adapter/service/siat_compra_venta_service_test.go) |
| **Computarizada** | [`siat_computarizada_service_test.go`](./internal/adapter/service/siat_computarizada_service_test.go) |

> **Configuración de Ambiente**
> Antes de ejecutar los tests, asegúrese de crear un archivo `.env` configurado con sus credenciales del ambiente de pruebas del SIAT.


---

## 👍 Contribución y Soporte

¡Las contribuciones son bienvenidas! Si deseas ayudar a mejorar `go-siat`, puedes:

1. Abrir un **Issue** para reportar bugs o solicitar nuevas características (por favor revisa el [`CONTRIBUTING.md`](./.github/CONTRIBUTING.md)).
2. Enviar un **Pull Request** con mejoras o correcciones.
3. Dejar una ⭐️ en el repositorio si este SDK te ha ahorrado horas de lidiar con SOAP.

Si necesitas ayuda técnica o soporte comercial para integrar la facturación electrónica en tu empresa, revisa nuestro [`SUPPORT.md`](./.github/SUPPORT.md).

---

## 🧾 Licencia

Distribuido bajo la **Licencia MIT**. Consulte el archivo [`LICENSE`](./LICENSE) para más detalles.
