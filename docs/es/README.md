<h1 align="center">
  <img src="../../.github/logo.svg" alt="go-siat logo" width="250">
  <br>
  Documentación de go-siat
</h1>

<p align="center">
  <a href="../en/README.md">
    <img src="https://img.shields.io/badge/lang-english-red?style=flat-square" alt="English Version">
  </a>
  <a href="https://go.dev/">
    <img src="https://img.shields.io/badge/go-1.25+-00ADD8?style=flat-square" alt="Go Version">
  </a>
  <a href="../../LICENSE">
    <img src="https://img.shields.io/badge/license-MIT-green?style=flat-square" alt="Licencia">
  </a>
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
| 3 | [**Referencia de API**](referencia-api.md) | Referencia completa de los 7 servicios SIAT y 67+ métodos. |
| 4 | [**Guía de Facturación**](guia-facturacion.md) | Ciclo de vida de facturas, 35 sectores, firmas digitales y procesamiento masivo. |
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
  <sub>Copyright © 2026 Ronaldo Rua — Licenciado bajo MIT</sub>
</p>
