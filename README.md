# go-siat

[![Status](https://img.shields.io/badge/status-active-success)](https://github.com/ron86i/go-siat)
[![Go Version](https://img.shields.io/badge/go-1.22+-00ADD8?logo=go)](https://go.dev/)
[![Architecture](https://img.shields.io/badge/architecture-hexagonal-blue)](#-arquitectura-del-servicio)

**go-siat** es un **servicio backend especializado** diseñado para gestionar la integración robusta con los servicios del **SIAT (Servicio de Impuestos Nacionales de Bolivia)**. Implementado en Go, centraliza la comunicación SOAP de alta complejidad mediante una arquitectura de Puertos y Adaptadores.

## 🚀 Capacidades Implementadas

El core actualmente centraliza las operaciones del **Servicio de Códigos** del SIAT:

### Gestión de Códigos (`codigos`)
- ✅ **CUIS**: Generación y obtención masiva e individual del Código Único de Inicio de Sistemas.
- ✅ **CUFD**: Gestión masiva e individual de Códigos Únicos de Facturación Diaria.
- ✅ **Validación de NIT**: Verificación en tiempo real de la vigencia de contribuyentes.
- ✅ **Comunicación**: Test de comunicación oficial con los servidores del SIAT.

## 🛠️ Arquitectura del Servicio

El proyecto sigue una estructura de **Arquitectura Hexagonal**, separando estrictamente la definición del dominio de las implementaciones técnicas (SOAP, firmadores, etc.):

```text
internal/
├── core/
│   ├── domain/
│   │   └── facturacion/
│   │       ├── codigos/        # Lógica de CUIS/CUFD/NIT
│   │       ├── sincronizacion/ # Catálogos y parámetros
│   │       └── compra_venta/   # Recepción y Anulación
│   └── port/                   # Interfaces (Contratos del SIAT)
├── adapter/
│   └── service/siat/           # Implementación clientes SOAP
```

## ⚙️ Configuración

Requiere un archivo `.env` configurado con las credenciales de entorno del SIN:

| Variable | Propósito |
| :--- | :--- |
| `SIAT_TOKEN` | Token delegado del SIN. |
| `SIAT_NIT` | NIT del emisor. |
| `SIAT_CODIGO_SISTEMA` | Código del sistema certificado. |
| `SIAT_URL` | Endpoint base del SIAT. |

---

## 🗺️ Roadmap de Implementación

Los siguientes módulos se encuentran en fase de definición de modelos de dominio y estructuración SOAP:

### 1. Sincronización de Catálogos (`sincronizacion`)
- 🔄 **Catálogos Paramétricos**: Sincronización de eventos, motivos de anulación, métodos de pago, etc.
- 🔄 **Actividades Económicas**: Obtención y mapeo de actividades y documentos sector asociados.
- 🔄 **Leyendas y Productos**: Homologación de códigos y textos legales.

### 2. Emisión y Recepción (`compra_venta`)
- 🔄 **Recepción de Facturas**: Protocolo de envío de paquetes de facturas electrónicas.
- 🔄 **Anulación**: Gestión de estados y motivos de anulación.

### 3. Core Técnico
- [ ] **Firma Digital XML**: Implementación de firma con estándar DSIG.
- [ ] **Persistencia**: Integración con PostgreSQL (`pgx`) para auditoría y logs.

## 📄 Licencia

Licencia MIT. Consulte `LICENSE` para detalles.
