<p align="center">
  <img src="../../.github/logo.svg" alt="go-siat logo" width="300">
</p>

<p align="center">
  <a href="https://masterminds.github.io/stability/active.html"><img src="https://masterminds.github.io/stability/active.svg" alt="Stability: Active"></a>
  <a href="https://goreportcard.com/report/github.com/ron86i/go-siat"><img src="https://goreportcard.com/badge/github.com/ron86i/go-siat" alt="Go Report Card"></a>
  <br>
  <a href="https://pkg.go.dev/github.com/ron86i/go-siat"><img src="https://pkg.go.dev/badge/github.com/ron86i/go-siat.svg" alt="Go Reference"></a>
  <a href="https://go.dev/"><img src="https://img.shields.io/github/go-mod/go-version/ron86i/go-siat?style=flat" alt="Go Version"></a>
  <a href="../../LICENSE"><img src="https://img.shields.io/github/license/ron86i/go-siat?style=flat" alt="Licencia"></a>
  <a href="https://github.com/ron86i/go-siat/releases"><img src="https://img.shields.io/github/v/release/ron86i/go-siat?style=flat&label=release" alt="Última Versión"></a>
  <a href="../../README.md"><img src="https://img.shields.io/badge/lang-english-red?style=flat" alt="English Version"></a>
</p>


<p align="center">
  <em><b>go-siat</b> es un SDK profesional desarrollado en Go, diseñado para simplificar la integración con los servicios web SOAP del <b>SIAT (Sistema Integrado de Administración Tributaria)</b>.</em>
</p>

> [!IMPORTANT]
> **Documentación Completa**: [Español](../../docs/es/README.md) | [English](../../docs/en/README.md)
>
> Arquitectura, Referencia de API, Guía de Facturación, Códigos de Error, Utilidades y Configuración.

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
| **Sectores Especializados** | Soporte para servicios de Telecomunicaciones, Servicios Básicos, Entidades Financieras y Boletos Aéreos. |
| **Sectores Especiales** | Soporte verificado para los **35 sectores** reglamentarios del SIAT. |

---

## 🚀 Inicio Rápido

### Instalación

```bash
go get github.com/ron86i/go-siat
```

### Requisitos

- Go 1.25 o superior.
- Certificado digital válido (p12/pfx) y clave privada (para modalidad Electrónica).

> [!TIP]
> **Mejores Prácticas de Contexto**: Proporcione siempre un contexto con timeout (ej. 30s) en todas las llamadas al SDK. Evite el uso directo de `context.Background()` para prevenir peticiones suspendidas si el servidor del SIAT está lento.

### Ejemplo: Verificar NIT

```go
package main

import (
    "context"
    "fmt"
    "github.com/ron86i/go-siat"
    "github.com/ron86i/go-siat/pkg/models"
)

func main() {
    // 1. Inicializar cliente
    s, _ := siat.New("YOUR_SIAT_URL", nil)

    // 2. Preparar solicitud usando builders
    req := models.Codigos().NewVerificarNitBuilder().
        WithNit(123456789).
        Build()

    // 3. Ejecutar llamada
    resp, err := s.Codigos().VerificarNit(context.Background(), req)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Estado de transacción NIT: %v\n", resp.Body.Content.RespuestaVerificarNit.Transaccion)
}
```

> Para guías avanzadas (flujo completo de facturación, firmas digitales, manejo de errores, middleware), consulta la [**Documentación Completa**](../../docs/es/README.md).

---

## 🛠️ Referencia de Uso (Tests)

Los **Tests de Integración** actúan como la documentación técnica viviente del SDK:

| Categoría | Archivo de Test |
| :--- | :--- |
| **Códigos** | [`siat_codigos_service_test.go`](../../internal/adapter/services/siat_codigos_service_test.go) |
| **Sincronización** | [`siat_sincronizacion_service_test.go`](../../internal/adapter/services/siat_sincronizacion_service_test.go) |
| **Operaciones** | [`siat_operaciones_service_test.go`](../../internal/adapter/services/siat_operaciones_service_test.go) |
| **Compra-Venta** | [`siat_compra_venta_service_test.go`](../../internal/adapter/services/siat_compra_venta_service_test.go) |
| **Electrónica** | [`siat_electronica_service_test.go`](../../internal/adapter/services/siat_electronica_service_test.go) |
| **Computarizada** | [`siat_computarizada_service_test.go`](../../internal/adapter/services/siat_computarizada_service_test.go) |
| **Documentos de Ajuste** | [`siat_documento_ajuste_service_test.go`](../../internal/adapter/services/siat_documento_ajuste_service_test.go) |
| **Telecomunicaciones** | [`siat_telecomunicaciones_service_test.go`](../../internal/adapter/services/siat_telecomunicaciones_service_test.go) |
| **Servicios Básicos** | [`siat_servicio_basico_service_test.go`](../../internal/adapter/services/siat_servicio_basico_service_test.go) |
| **Entidades Financieras** | [`siat_entidad_financiera_service_test.go`](../../internal/adapter/services/siat_entidad_financiera_service_test.go) |
| **Boletos Aéreos** | [`siat_boleto_aereo_service_test.go`](../../internal/adapter/services/siat_boleto_aereo_service_test.go) |
| **Facturación (Sectores)** | [`pkg/models/invoices/`](../../pkg/models/invoices/) |
| **Flujos Completos** | [`siat_test.go`](../../siat_test.go) |

---

## 👍 Contribución y Soporte

¡Las contribuciones son bienvenidas! Si deseas ayudar a mejorar `go-siat`, puedes:

1. Abrir un **Issue** para reportar bugs o solicitar nuevas características (por favor revisa el [`CONTRIBUTING.md`](../../.github/CONTRIBUTING.md)).
2. Enviar un **Pull Request** con mejoras o correcciones.
3. Dejar una ⭐️ en el repositorio si este SDK te ha ahorrado horas de lidiar con SOAP.

Si necesitas ayuda técnica o soporte comercial para integrar la facturación electrónica en tu empresa, revisa nuestro [`SUPPORT.md`](../../.github/SUPPORT.md).

---

## 🧾 Licencia

Distribuido bajo la **Licencia MIT**. Consulte el archivo [`LICENSE`](../../LICENSE) para más detalles.
