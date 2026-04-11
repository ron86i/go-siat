<h1 align="center">
  <img src="../../.github/logo.svg" alt="go-siat logo" width="250">
  <br>
  Documentación de go-siat
</h1>

<p align="center">
  <a href="../en/README.md"><img src="https://img.shields.io/badge/lang-english-red?style=flat" alt="English Version"></a>
  <a href="https://masterminds.github.io/stability/active.html"><img src="https://masterminds.github.io/stability/active.svg" alt="Stability: Active"></a>
  <a href="https://goreportcard.com/report/github.com/ron86i/go-siat"><img src="https://goreportcard.com/badge/github.com/ron86i/go-siat" alt="Go Report Card"></a>
  <br>
  <a href="https://pkg.go.dev/github.com/ron86i/go-siat"><img src="https://pkg.go.dev/badge/github.com/ron86i/go-siat.svg" alt="Go Reference"></a>
  <a href="https://go.dev/"><img src="https://img.shields.io/github/go-mod/go-version/ron86i/go-siat?style=flat" alt="Go Version"></a>
  <a href="../../LICENSE"><img src="https://img.shields.io/github/license/ron86i/go-siat?style=flat" alt="Licencia"></a>
  <a href="https://github.com/ron86i/go-siat/releases"><img src="https://img.shields.io/github/v/release/ron86i/go-siat?style=flat&label=release" alt="Última Versión"></a>
</p>

<p align="center">
  <em>SDK profesional para integración con los servicios web SOAP del <b>SIAT (Sistema Integrado de Administración Tributaria)</b> de Bolivia.</em>
</p>

---

## 📚 Índice de Documentación

Bienvenido a la documentación técnica de **go-siat**. Esta guía está organizada para ayudarte a ir de cero a producción con el sistema de facturación electrónica de Bolivia.

| # | Documento | Descripción |
|:--|:----------|:------------|
| 1 | [**Arquitectura**](arquitectura.md) | Diseño hexagonal, capas, patrones de diseño y flujo de datos interno. |
| 2 | [**Inicio Rápido**](inicio-rapido.md) | Instalación, prerrequisitos, configuración de ambiente y tu primera llamada. |
| 3 | [**Referencia de API**](referencia-api.md) | Referencia completa de los 17 servicios SIAT y 100+ métodos. |
| 4 | [**Guía de Facturación**](guia-facturacion.md) | Ciclo de vida de facturas, 48 sectores, firmas digitales y procesamiento masivo. |
| 5 | [**Manejo de Errores**](manejo-errores.md) | Tipos de error, 150+ códigos SIAT, estrategias de reintento y verificación. |
| 6 | [**Utilidades**](utilidades.md) | Generación de CUF, firma XML, compresión, hash y helpers de parseo. |
| 7 | [**Configuración**](configuracion.md) | Configuración de cliente HTTP, middleware, trazabilidad distribuida y constantes. |

---

## 🗺️ Navegación Rápida por Caso de Uso

### "Quiero..."

| Objetivo | Empezar Aquí |
|:---------|:-------------|
| Instalar y hacer mi primera llamada | [Inicio Rápido](inicio-rapido.md) |
| Entender la arquitectura del proyecto | [Arquitectura](arquitectura.md) |
| Enviar una factura electrónica | [Guía de Facturación](guia-facturacion.md) |
| Saber qué sectores están soportados | [Guía de Facturación → Sectores](guia-facturacion.md#referencia-de-sectores-soportados) |
| Manejar errores del SIAT en producción | [Manejo de Errores](manejo-errores.md) |
| Firmar facturas con mi certificado digital | [Utilidades → Firma XML](utilidades.md#firma-digital-xml) |
| Personalizar timeouts HTTP o agregar logging | [Configuración](configuracion.md) |
| Encontrar la firma de un método específico | [Referencia de API](referencia-api.md) |

---

## 🔗 Recursos Adicionales

| Recurso | Ubicación |
|:--------|:----------|
| README Principal | [`README.md`](../../README.md) |
| Guía de Contribución | [`CONTRIBUTING.md`](../../.github/CONTRIBUTING.md) |
| Soporte y Consultoría | [`SUPPORT.md`](../../.github/SUPPORT.md) |
| Código de Conducta | [`CODE_OF_CONDUCT.md`](../../.github/CODE_OF_CONDUCT.md) |
| Tests de Integración | [`internal/adapter/services/*_test.go`](../../internal/adapter/services/) |
| Tests de Facturas por Sector | [`pkg/models/invoices/*_test.go`](../../pkg/models/invoices/) |
| Licencia (MIT) | [`LICENSE`](../../LICENSE) |

---

<p align="center">
  <sub>Copyright © 2026 Ronaldo Rua - Licenciado bajo MIT</sub>
</p>
