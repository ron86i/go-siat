<h1 align="center">
  <img src="../../.github/logo.svg" alt="go-siat logo" width="300">
  <br>
  <a href="https://github.com/ron86i/go-siat">
    <img src="https://img.shields.io/badge/status-active-success?style=flat-square" alt="Status">
  </a>
  <a href="https://go.dev/">
    <img src="https://img.shields.io/badge/go-1.25+-00ADD8?style=flat-square" alt="Go Version">
  </a>
  <a href="../../LICENSE">
    <img src="https://img.shields.io/badge/license-MIT-green?style=flat-square" alt="License">
  </a>
  <a href="../../README.md">
    <img src="https://img.shields.io/badge/lang-english-red?style=flat-square" alt="English Version">
  </a>
</h1>

<p align="center">
  <em><b>go-siat</b> es un SDK profesional desarrollado en Go, diseñado para simplificar la integración con los servicios web SOAP del <b>SIAT (Sistema Integrado de Administración Tributaria)</b>.</em>
</p>

## 💡 ¿Por qué go-siat?

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
- 🧩 **Modular**: Múltiples servicios (`Codigos`, `Sincronizacion`, `Operaciones`, `CompraVenta`, `Electronica`, `Computarizada`) claramente separados.
- 🏢 **Multi-Sector**: Soporte nativo y verificado para **35 sectores** distintos (Compra y Venta, Hoteles, Minería, Hospitales, Hidrocarburos, etc.).

---

## 📖 Tabla de Contenidos

1. [¿Por qué go-siat?](#-por-qué-go-siat)
2. [Características](#-características)
3. [Guía de Inicio Rápido](#guía-de-inicio-rápido)
4. [Ejemplos Avanzados](#-ejemplos-avanzados)
5. [Referencia de Uso (Tests)](#-referencia-de-uso-tests)
6. [Contribución y Soporte](#-contribución-y-soporte)
7. [Licencia](#-licencia)

---


## Capacidades Implementadas

El SDK cubre los servicios críticos del ecosistema SIAT:

| Servicios | Funcionalidades Clave |
| :--- | :--- |
| **Códigos** | Solicitud de CUIS/CUFD (Individual y Masivo), Validación de NIT, Comunicación. |
| **Sincronización** | Catálogos de actividades, paramétricas, productos, servicios y documentos sector. |
| **Operaciones** | Registro/Cierre de Puntos de Venta, Gestión de Eventos Significativos. |
| **Compra-Venta** | Servicio específico para compra venta, bonificaciones y tasas. |
| **Documentos de Ajuste** | Gestión de notas de Crédito/Débito, Conciliación y reversiones. |
| **Electrónica en Línea** | Soporte completo para facturación con firma digital. |
| **Computarizada en Línea** | Soporte para modalidades sin firma digital. |
| **Sectores Especiales** | Soporte verificado para los **35 sectores** reglamentarios del SIAT. |

---

## Sectores Soportados

`go-siat` incluye modelos de dominio, builders y **tests de integración** para los **35 sectores** reglamentarios del SIAT (ubicados en `pkg/models/invoices/`):

### 🏢 Estándar y Servicios
- **Compra-Venta (Sector 1)**: El sector estándar para la mayoría de comercios.
- **Alquiler de Bienes Inmuebles**: Para el sector inmobiliario y arrendamientos.
- **Seguros**: Emisión de pólizas y servicios de aseguradoras.
- **Servicios Básicos**: Suministro de energía eléctrica, agua, gas y telecomunicaciones.
- **Servicios Turísticos y Hospedaje / Hoteles**: Para el sector hotelero y operadores turísticos.
- **Hospitales y Clínicas**: Servicios de salud (Nacional y Zona Franca).
- **Seguridad Alimentaria**: Comercialización de productos de la canasta básica.

### 🏺 Exportación y Zona Franca
- **Exportación de Bienes y Servicios**: Comercial de Exportación, Servicios y Libre Consignación.
- **Zona Franca**: Facturas de Zona Franca, Alquiler ZF y Servicios Hospitalarios ZF.
- **Duty Free**: Facturación para tiendas libres de impuestos en aeropuertos.

### ⛽ Hidrocarburos y Energía
- **Comercialización de Hidrocarburos**: Combustibles, Lubricantes (con y sin IEHD).
- **Engarrafadoras**: Sector de distribución de GLP.
- **GNC y GNV**: Comercialización de Gas Natural Vehicular.
- **Combustible No Subvencionado**: Para la venta a precio internacional.

### ⛰️ Minería y Metales
- **Venta de Minerales**: Venta Interna y Exportación de Minerales.
- **Venta al BCB**: Venta de oro y minerales al **Banco Central de Bolivia**.

### 🎓 Educación
- **Sectores Educativos**: Colegios, Universidades e Institutos (Nacional y Zona Franca).

### 🔄 Documentos de Ajuste
- **Notas de Crédito / Débito**: Notas estándar, ICE y Fiscales.
- **Notas de Conciliación**: Notas de conciliación para ajustes de facturación.

### 🎲 Otros Sectores Especiales
- **Juegos de Azar**: Casinos y salas de entretenimiento.
- **Tasa Cero**: Libros y transporte internacional de carga.
- **Productos ICE**: Artículos alcanzados por el Impuesto al Consumo Específico.
- **Pagos Anticipados y Factura Compartida**: Flujos de facturación complejos.
- **Prevalorada**: Facturas con precio fijo y servicio tributario recurrente.
- **Compra y Venta de Moneda Extranjera**: Casas de cambio y entidades financieras.


---
```bash
go get github.com/ron86i/go-siat
```

---
## Guía de Inicio Rápido

### Instalación

- Go 1.25 o superior.
- Certificado digital válido (p12/pfx) y clave privada.

> [!TIP]
> **Mejores Prácticas de Contexto**: Proporcione siempre un contexto con timeout (ej. 30s) en todas las llamadas al SDK. Evite el uso directo de `context.Background()` para prevenir peticiones suspendidas si el servidor del SIAT está lento.

### Uso Básico

El siguiente ejemplo demuestra cómo inicializar el cliente y realizar una solicitud de código CUIS:

```go
package main

import (
    "context"
    "log"
    "os"

    "github.com/joho/godotenv"
    "github.com/ron86i/go-siat"
    "github.com/ron86i/go-siat/pkg/models"
    "github.com/ron86i/go-siat/pkg/utils"
)

func main() {
    // 1. Cargar configuración desde .env
    godotenv.Load()

    siatURL := os.Getenv("SIAT_URL")
    if siatURL == "" {
        siatURL = "https://pilotosiatservicios.impuestos.gob.bo/v2"
    }

    // 2. Inicializar cliente SIAT
    s, err := siat.New(siatURL, nil)
    if err != nil {
        log.Fatalf("Error inicializando SDK: %v", err)
    }

    // 3. Preparar configuración con token
    cfg := siat.Config{
        Token: os.Getenv("SIAT_TOKEN"),
    }

    // 4. Parsear valores de configuración de forma segura
    nit, err := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
    if err != nil {
        log.Fatalf("NIT inválido: %v", err)
    }

    codAmbiente, err := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
    if err != nil {
        log.Fatalf("Código de ambiente inválido: %v", err)
    }

    // 5. Construir solicitud CUIS usando el Builder
    req := models.Codigos().NewCuisBuilder().
        WithCodigoAmbiente(codAmbiente).
        WithCodigoModalidad(siat.ModalidadElectronica).
        WithCodigoPuntoVenta(0).
        WithCodigoSucursal(0).
        WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
        WithNit(nit).
        Build()

    // 6. Ejecutar solicitud
    ctx := context.Background()
    resp, err := s.Codigos().SolicitudCuis(ctx, cfg, req)
    if err != nil {
        log.Fatalf("Error solicitando CUIS: %v", err)
    }

    // 7. Procesar respuesta
    if resp.Body.Fault != nil {
        log.Fatalf("Error SIAT: %s", resp.Body.Fault.String)
    }

    cuis := resp.Body.Content.RespuestaCuis.Codigo
    vigencia := resp.Body.Content.RespuestaCuis.FechaVigCuis

    log.Printf("✓ CUIS obtenido: %s (válido hasta %s)", cuis, vigencia)
}
```

**Archivo .env requerido:**
```
SIAT_URL=https://pilotosiatservicios.impuestos.gob.bo/v2
SIAT_TOKEN=tu_token_api
SIAT_NIT=123456789
SIAT_CODIGO_AMBIENTE=1
SIAT_CODIGO_SISTEMA=ABC123DEF
```

---

## 👀 Ejemplos Avanzados

A continuación, mostramos algunos de los flujos más comunes. Para más información del diseño arquitectónico, consulte [ARCHITECTURE.md](ARCHITECTURE.md). Si desea ver la documentación técnica completa, revise nuestro repositorio de [Tests de Integración](#referencia-de-uso-tests).

<details>
  <summary>📚 Emitir y Enviar una Factura (Flujo Completo)</summary>

Este ejemplo muestra el flujo completo: obtener CUIS y CUFD, construir una factura, firmarla y enviarla:

```go
package main

import (
    "context"
    "encoding/xml"
    "log"
    "net/http"
    "os"
    "time"

    "github.com/joho/godotenv"
    "github.com/ron86i/go-siat"
    "github.com/ron86i/go-siat/pkg/models"
    "github.com/ron86i/go-siat/pkg/models/invoices"
    "github.com/ron86i/go-siat/pkg/utils"
)

func main() {
    godotenv.Load()

    // Configuración
    siatURL := os.Getenv("SIAT_URL")
    if siatURL == "" {
        siatURL = "https://pilotosiatservicios.impuestos.gob.bo/v2"
    }

    cfg := siat.Config{Token: os.Getenv("SIAT_TOKEN")}
    nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
    codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
    codModalidad := siat.ModalidadElectronica

    // Inicializar cliente
    client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}
    s, err := siat.New(siatURL, client)
    if err != nil {
        log.Fatalf("Error inicializando cliente: %v", err)
    }

    serviceCodigos := s.Codigos()
    serviceElectronica := s.Electronica()

    // ====== PASO 1: Obtener CUIS ======
    cuisReq := models.Codigos().NewCuisBuilder().
        WithCodigoAmbiente(codAmbiente).
        WithCodigoModalidad(codModalidad).
        WithCodigoPuntoVenta(0).
        WithCodigoSucursal(0).
        WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
        WithNit(nit).
        Build()

    cuisResp, err := serviceCodigos.SolicitudCuis(context.Background(), cfg, cuisReq)
    if err != nil {
        log.Fatalf("Error obteniendo CUIS: %v", err)
    }
    if cuisResp.Body.Fault != nil {
        log.Fatalf("Error SIAT CUIS: %s", cuisResp.Body.Fault.String)
    }
    cuis := cuisResp.Body.Content.RespuestaCuis.Codigo
    log.Printf("✓ CUIS obtenido: %s", cuis)

    // ====== PASO 2: Obtener CUFD ======
    cufdReq := models.Codigos().NewCufdBuilder().
        WithCodigoAmbiente(codAmbiente).
        WithCodigoModalidad(codModalidad).
        WithCodigoPuntoVenta(0).
        WithCodigoSucursal(0).
        WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
        WithNit(nit).
        WithCuis(cuis).
        Build()

    cufdResp, err := serviceCodigos.SolicitudCufd(context.Background(), cfg, cufdReq)
    if err != nil {
        log.Fatalf("Error obteniendo CUFD: %v", err)
    }
    if cufdResp.Body.Fault != nil {
        log.Fatalf("Error SIAT CUFD: %s", cufdResp.Body.Fault.String)
    }
    cufd := cufdResp.Body.Content.RespuestaCufd.Codigo
    cufdControl := cufdResp.Body.Content.RespuestaCufd.CodigoControl
    log.Printf("✓ CUFD obtenido: %s", cufd)

    // ====== PASO 3: Generar CUF ======
    fechaEmision := time.Now()
    cuf, err := utils.GenerarCUF(nit, fechaEmision, 0, 1, 1, 1, 1, 1, 0, cufdControl)
    if err != nil {
        log.Fatalf("Error generando CUF: %v", err)
    }
    log.Printf("✓ CUF generado: %s", cuf)

    // ====== PASO 4: Construir Factura ======
    nombre := "CLIENTE PRUEBA"
    cabecera := invoices.NewCompraVentaCabeceraBuilder().
        WithNitEmisor(nit).
        WithRazonSocialEmisor("MI EMPRESA S.R.L.").
        WithMunicipio("La Paz").
        WithDireccion("Calle Principal 123").
        WithNumeroFactura(1).
        WithCuf(cuf).
        WithCufd(cufd).
        WithFechaEmision(fechaEmision).
        WithNombreRazonSocial(&nombre).
        WithMontoTotal(100).
        WithCodigoDocumentoSector(1).
        Build()

    detalle := invoices.NewCompraVentaDetalleBuilder().
        WithActividadEconomica("477300").
        WithCodigoProductoSin(622539).
        WithDescripcion("PRODUCTO DE PRUEBA").
        WithCantidad(1).
        WithPrecioUnitario(100).
        WithSubTotal(100).
        Build()

    factura := invoices.NewCompraVentaBuilder().
        WithModalidad(codModalidad).
        WithCabecera(cabecera).
        AddDetalle(detalle).
        Build()

    log.Printf("✓ Factura construida")

    // ====== PASO 5: Serializar, Firmar y Preparar ======
    xmlData, err := xml.Marshal(factura)
    if err != nil {
        log.Fatalf("Error serializando XML: %v", err)
    }

    // Nota: Requiere certificados válidos (key.pem y cert.crt)
    signedXML, err := utils.SignXML(xmlData, "key.pem", "cert.crt")
    if err != nil {
        log.Printf("Advertencia: Firma digital no disponible (certificados no encontrados): %v", err)
        signedXML = xmlData
    }

    hash, archivoBase64, err := utils.CompressAndHash(signedXML)
    if err != nil {
        log.Fatalf("Error comprimiendo/hasheando: %v", err)
    }
    log.Printf("✓ Factura preparada (hash: %s)", hash[:16]+"...")

    // ====== PASO 6: Enviar Factura al SIAT ======
    recepcionReq := models.Electronica().NewRecepcionFacturaBuilder().
        WithCodigoAmbiente(codAmbiente).
        WithNit(nit).
        WithCufd(cufd).
        WithCuis(cuis).
        WithTipoFacturaDocumento(1).
        WithArchivo(archivoBase64).
        WithFechaEnvio(fechaEmision).
        WithHashArchivo(hash).
        Build()

    recepcionResp, err := serviceElectronica.RecepcionFactura(context.Background(), cfg, recepcionReq)
    if err != nil {
        log.Fatalf("Error enviando factura: %v", err)
    }
    if recepcionResp.Body.Fault != nil {
        log.Fatalf("Error SIAT RecepcionFactura: %s", recepcionResp.Body.Fault.String)
    }

    estado := recepcionResp.Body.Content.RespuestaServicioFacturacion.CodigoEstado
    log.Printf("✓ Factura enviada - Estado: %d", estado)
}
```
</details>

<details>
  <summary>✅ Verificar NIT y Validar Estado</summary>

Antes de emitir facturas, es recomendable verificar que el NIT del cliente esté activo:

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    "github.com/ron86i/go-siat"
    "github.com/ron86i/go-siat/pkg/models"
    "github.com/ron86i/go-siat/pkg/utils"
)

func verificarNIT(s *siat.SiatServices, cfg siat.Config, nit int64) error {
    req := models.Codigos().NewVerificarNitBuilder().
        WithNit(nit).
        Build()

    resp, err := s.Codigos().VerificarNit(context.Background(), cfg, req)
    if err != nil {
        return fmt.Errorf("error verificando NIT: %w", err)
    }

    if resp.Body.Fault != nil {
        return fmt.Errorf("NIT no válido: %s", resp.Body.Fault.String)
    }

    resultado := resp.Body.Content.VerificarNitRespuesta
    log.Printf("✓ NIT %d verificado correctamente", nit)
    log.Printf("  Estado: %d", resultado.CodigoError)
    
    return nil
}

func main() {
    godotenv.Load()

    siatURL := os.Getenv("SIAT_URL")
    if siatURL == "" {
        siatURL = "https://pilotosiatservicios.impuestos.gob.bo/v2"
    }

    s, err := siat.New(siatURL, nil)
    if err != nil {
        log.Fatalf("Error inicializando SDK: %v", err)
    }

    cfg := siat.Config{Token: os.Getenv("SIAT_TOKEN")}
    nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))

    // Verificar NIT antes de emitir factura
    if err := verificarNIT(s, cfg, nit); err != nil {
        log.Printf("✗ %v", err)
    }
}
```
</details>

<details>
  <summary>🔄 Renovar CUFD (Código Único de Facturación Diaria)</summary>

El CUFD tiene vigencia diaria. Este ejemplo muestra cómo renovarlo:

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    "github.com/ron86i/go-siat"
    "github.com/ron86i/go-siat/pkg/models"
    "github.com/ron86i/go-siat/pkg/utils"
)

func renovarCUFD(s *siat.SiatServices, cfg siat.Config, nit int64, codAmbiente int, cuis string) (string, error) {
    req := models.Codigos().NewCufdBuilder().
        WithNit(nit).
        WithCodigoAmbiente(codAmbiente).
        WithCodigoSucursal(0).
        WithCodigoPuntoVenta(0).
        WithCodigoModalidad(siat.ModalidadElectronica).
        WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
        WithCuis(cuis).
        Build()

    resp, err := s.Codigos().SolicitudCufd(context.Background(), cfg, req)
    if err != nil {
        return "", fmt.Errorf("error renovando CUFD: %w", err)
    }

    if resp.Body.Fault != nil {
        return "", fmt.Errorf("error SIAT: %s", resp.Body.Fault.String)
    }

    cufdCode := resp.Body.Content.RespuestaCufd.Codigo
    cufdDate := resp.Body.Content.RespuestaCufd.FechaVigCufd

    log.Printf("✓ CUFD renovado: %s (válido hasta %s)", cufdCode, cufdDate)
    return cufdCode, nil
}

func main() {
    godotenv.Load()

    s, _ := siat.New(os.Getenv("SIAT_URL"), nil)
    cfg := siat.Config{Token: os.Getenv("SIAT_TOKEN")}

    nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
    codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))

    cufd, err := renovarCUFD(s, cfg, nit, codAmbiente, "197C8240")
    if err != nil {
        log.Fatalf("Error: %v", err)
    }

    log.Printf("Usar este CUFD en facturas de hoy: %s", cufd)
}
```
</details>

<details>
  <summary>🏪 Obtener Catálogos (Sincronización)</summary>

Sincronizar catálogos de actividades económicas, monedas, tipos de cambio, etc.:

```go
package main

import (
    "context"
    "log"
    "os"

    "github.com/joho/godotenv"
    "github.com/ron86i/go-siat"
    "github.com/ron86i/go-siat/pkg/models"
    "github.com/ron86i/go-siat/pkg/utils"
)

func main() {
    godotenv.Load()

    s, _ := siat.New(os.Getenv("SIAT_URL"), nil)
    cfg := siat.Config{Token: os.Getenv("SIAT_TOKEN")}

    // Parsear valores necesarios
    nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
    codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))

    // Construcción de solicitud común para sincronización
    baseReq := models.Sincronizacion().NewSincronizarActividadesBuilder().
        WithNit(nit).
        WithCodigoAmbiente(codAmbiente).
        WithCodigoSucursal(0).
        WithCodigoPuntoVenta(0).
        WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
        WithCuis("C2FC682C")

    // 1. Obtener actividades económicas
    actReq := baseReq.Build()
    actResp, err := s.Sincronizacion().SincronizarActividades(context.Background(), cfg, actReq)
    if err != nil {
        log.Fatalf("Error sincronizando actividades: %v", err)
    }
    if actResp != nil {
        log.Printf("✓ Actividades económicas sincronizadas")
    }

    // 2. Obtener monedas vigentes
    monedasReq := models.Sincronizacion().NewSincronizarMonedasBuilder().
        WithNit(nit).
        WithCodigoAmbiente(codAmbiente).
        WithCodigoSucursal(0).
        WithCodigoPuntoVenta(0).
        WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
        WithCuis("C2FC682C").
        Build()

    monedasResp, err := s.Sincronizacion().SincronizarMonedas(context.Background(), cfg, monedasReq)
    if err != nil {
        log.Fatalf("Error sincronizando monedas: %v", err)
    }
    if monedasResp != nil {
        log.Printf("✓ Monedas vigentes sincronizadas")
    }

    // 3. Obtener documentos fiscales por sector
    docsReq := models.Sincronizacion().NewSincronizarListaActividadesDocumentoSectorBuilder().
        WithNit(nit).
        WithCodigoAmbiente(codAmbiente).
        WithCodigoSucursal(0).
        WithCodigoPuntoVenta(0).
        WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
        WithCuis("C2FC682C").
        Build()

    docsResp, err := s.Sincronizacion().SincronizarListaActividadesDocumentoSector(context.Background(), cfg, docsReq)
    if err != nil {
        log.Fatalf("Error sincronizando documentos: %v", err)
    }
    if docsResp != nil {
        log.Printf("✓ Documentos por sector sincronizados")
    }
}
```
</details>

<details>
  <summary>⚙️ Manejo de Errores y Reintentos</summary>

Implementar lógica robusta de manejo de errores con reintentos exponenciales:

```go
package main

import (
    "context"
    "fmt"
    "log"
    "math"
    "os"
    "time"

    "github.com/joho/godotenv"
    "github.com/ron86i/go-siat"
    "github.com/ron86i/go-siat/pkg/models"
    "github.com/ron86i/go-siat/pkg/utils"
)

// operarConReintentos ejecuta una operación con reintentos exponenciales
func operarConReintentos(operacion func() error, maxIntentos int) error {
    var lastErr error

    for intento := 0; intento < maxIntentos; intento++ {
        if intento > 0 {
            // Backoff: 1s, 2s, 4s, 8s...
            backoff := time.Duration(math.Pow(2, float64(intento-1))) * time.Second
            log.Printf("Reintentando en %v (intento %d/%d)...", backoff, intento+1, maxIntentos)
            time.Sleep(backoff)
        }

        lastErr = operacion()
        if lastErr == nil {
            return nil
        }

        // No reintentar errores de configuración
        errMsg := lastErr.Error()
        if errMsg == "Token no válido" || errMsg == "NIT no existe" {
            return lastErr
        }

        log.Printf("Intento %d falló: %v", intento+1, lastErr)
    }

    return fmt.Errorf("operación falló después de %d intentos: %w", maxIntentos, lastErr)
}

func main() {
    godotenv.Load()

    s, _ := siat.New(os.Getenv("SIAT_URL"), nil)
    cfg := siat.Config{Token: os.Getenv("SIAT_TOKEN")}
    nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
    codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))

    // Solicitar CUFD con reintentos
    err := operarConReintentos(func() error {
        req := models.Codigos().NewCufdBuilder().
            WithNit(nit).
            WithCodigoAmbiente(codAmbiente).
            WithCodigoSucursal(0).
            WithCodigoPuntoVenta(0).
            WithCodigoModalidad(siat.ModalidadElectronica).
            WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
            WithCuis("197C8240").
            Build()

        resp, err := s.Codigos().SolicitudCufd(context.Background(), cfg, req)
        if err != nil {
            return err
        }
        
        if resp.Body.Fault != nil {
            return fmt.Errorf("SIAT error: %s", resp.Body.Fault.String)
        }

        log.Printf("✓ CUFD obtenido: %s", resp.Body.Content.RespuestaCufd.Codigo)
        return nil
    }, 3)

    if err != nil {
        log.Printf("✗ Error final: %v", err)
    } else {
        log.Printf("✓ Operación exitosa después de reintentos")
    }
}
```

Para más ejemplos y consultar los tests de integración, revise el repositorio de [Tests](#referencia-de-uso-tests).
</details>

---

## 🛠️ Referencia de Uso (Tests)

Para una comprensión profunda de cada servicio, los **Tests de Integración** actúan como la documentación técnica principal.

| Categoría | Archivo de Test |
| :--- | :--- |
| **Códigos** | [`siat_codigos_service_test.go`](../../internal/adapter/services/siat_codigos_service_test.go) |
| **Sincronización** | [`siat_sincronizacion_service_test.go`](../../internal/adapter/services/siat_sincronizacion_service_test.go) |
| **Operaciones** | [`siat_operaciones_service_test.go`](../../internal/adapter/services/siat_operaciones_service_test.go) |
| **Compra-Venta** | [`siat_compra_venta_service_test.go`](../../internal/adapter/services/siat_compra_venta_service_test.go) |
| **Electrónica** | [`siat_electronica_service_test.go`](../../internal/adapter/services/siat_electronica_service_test.go) |
| **Computarizada** | [`siat_computarizada_service_test.go`](../../internal/adapter/services/siat_computarizada_service_test.go) |
| **Documentos de Ajuste** | [`siat_documento_ajuste_service_test.go`](../../internal/adapter/services/siat_documento_ajuste_service_test.go) |
| **Facturación (Sectores)** | [`pkg/models/invoices/`](../../pkg/models/invoices/) |
| **Flujos Completos** | [`siat_test.go`](../../siat_test.go) |


> **Configuración de Ambiente**
> Antes de ejecutar los tests, asegúrese de crear un archivo `.env` configurado con sus credenciales del ambiente de pruebas del SIAT.


---

## 👍 Contribución y Soporte

¡Las contribuciones son bienvenidas! Si deseas ayudar a mejorar `go-siat`, puedes:

1. Abrir un **Issue** para reportar bugs o solicitar nuevas características (por favor revisa el [`CONTRIBUTING.md`](CONTRIBUTING.md)).
2. Enviar un **Pull Request** con mejoras o correcciones.
3. Dejar una ⭐️ en el repositorio si este SDK te ha ahorrado horas de lidiar con SOAP.

Si necesitas ayuda técnica o soporte comercial para integrar la facturación electrónica en tu empresa, revisa nuestro [`SUPPORT.md`](SUPPORT.md).

---

## 🧾 Licencia

Distribuido bajo la **Licencia MIT**. Consulte el archivo [`LICENSE`](../../LICENSE) para más detalles.
