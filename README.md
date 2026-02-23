# go-siat

[![Status](https://img.shields.io/badge/status-active-success)](https://github.com/ron86i/go-siat)
[![Go Version](https://img.shields.io/badge/go-1.26+-00ADD8?logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

**go-siat** es un SDK robusto escrito en Go, diseñado para facilitar la integración con los servicios web SOAP del **SIAT (Sistema de Facturación de Impuestos Nacionales de Bolivia)**. 

Actualmente, el SDK ya cuenta con las implementaciones completas para **Gestión de Códigos**, **Sincronización de Catálogos** y **Operaciones de Punto de Venta**, utilizando el cliente HTTP estándar de Go para garantizar ligereza y facilidad de mantenimiento.

---

## 📋 Tabla de Contenidos

- [Capacidades Implementadas](#-capacidades-implementadas)
- [Estructura del Proyecto](#-estructura-del-proyecto)
- [Configuración](#-configuración)
- [Testing](#-testing)
- [Roadmap de Implementación](#-roadmap-de-implementación)
- [Licencia](#-licencia)

---

## 🚀 Capacidades Implementadas

Actualmente, el proyecto soporta las operaciones críticas del **Servicio de Códigos** del SIAT:

### Gestión de Códigos (`codigos`)
- ✅ **CUIS**: Obtención individual y masiva del Código Único de Inicio de Sistemas.
- ✅ **CUFD**: Generación individual y masiva del Código Único de Facturación Diaria (vigencia 24h).
- ✅ **Validación de NIT**: Verificación automatizada de la validez y estado de contribuyentes.
- ✅ **Prueba de Comunicación**: Validación de conectividad y credenciales con los servidores oficiales.
- ✅ **Certificados Revocados**: Notificación de revocación de certificados digitales.

### Sincronización de Catálogos (`sincronizacion`)
- ✅ **Actividades Económicas**: Sincronización completa del catálogo de actividades del contribuyente.
- ✅ **Paramétricas**: Obtención de todos los catálogos paramétricos (Eventos, Motivos, Países, Monedas, etc).
- ✅ **Productos y Servicios**: Homologación y listado de productos y servicios autorizados.
- ✅ **Documentos Sector**: Relación entre actividades y tipos de documentos sector.

### Operaciones de Punto de Venta (`operaciones`)
- ✅ **Registro de Punto de Venta**: Apertura y registro de nuevos puntos de venta/comisionistas.
- ✅ **Cierre de Operaciones**: Gestión de cierre de sistemas y puntos de venta.
- ✅ **Eventos Significativos**: Registro y consulta de eventos (cortes de internet, fallas, etc).

---

## 📂 Estructura del Proyecto

El proyecto está organizado de la siguiente manera:

- **`cmd/`**: Ejemplos de uso del SDK y pruebas rápidas.
- **`internal/`**: Núcleo del SDK.
    - **`core/domain/`**: Modelos de datos y estructuras XML para el SIAT.
    - **`core/port/`**: Definición de interfaces y contratos.
    - **`adapter/service/`**: Implementación de la comunicación SOAP/HTTP con el SIAT.
- **`pkg/`**: Paquetes de utilidad, configuración y modelos auxiliares.
- **`siat.go`**: Punto de entrada principal para inicializar el SDK.

---

## ⚙️ Configuración

Cree un archivo `.env` en la raíz del proyecto basado en la siguiente tabla:

| Variable | Descripción | Ejemplo                                            |
| :--- | :--- |:---------------------------------------------------|
| `SIAT_TOKEN` | Token delegado proporcionado por el SIN | `eyJ0eX...`                                        |
| `SIAT_NIT` | NIT del emisor | `123456789`                                        |
| `SIAT_CODIGO_SISTEMA` | Código del sistema certificado | `ABC123XYZ`                                        |
| `SIAT_CODIGO_AMBIENTE` | Código de ambiente (1: Producción, 2: Pruebas) | `2`                                                |
| `SIAT_CODIGO_MODALIDAD` | Código de modalidad (1: Electrónica, 2: Computarizada) | `1`                                                |
| `SIAT_URL` | Endpoint base del SIAT (Pruebas/Producción) | `https://pilotosiatservicios.impuestos.gob.bo/...` |

---

## 🧪 Testing

El proyecto incluye una suite de pruebas unitarias y de integración para validar la comunicación con el SIAT.

### Ejecutar todas las pruebas
```bash
go test ./...
```

### Ejecutar pruebas del servicio SIAT (con logs)
```bash
go test -v ./internal/adapter/service/siat/...
```

> [!IMPORTANT]
> Para ejecutar las pruebas de integración con el SIAT, asegúrese de tener configuradas las variables de entorno correctas en su archivo `.env`.

## 🤝 Contribución y Soporte

¡Las contribuciones son lo que hacen que la comunidad de código abierto sea un lugar increíble para aprender, inspirar y crear!

- Si deseas colaborar, consulta nuestra [Guía de Contribución](CONTRIBUTING.md).
- Para apoyo financiero o soporte técnico especializado, revisa nuestra sección de [Soporte y Financiación](SUPPORT.md).

## 📄 Licencia

Este proyecto está bajo la Licencia MIT. Para más información, consulte el archivo [LICENSE](LICENSE).
