# go-siat

[![Status](https://img.shields.io/badge/status-active-success)](https://github.com/ron86i/go-siat)
[![Go Version](https://img.shields.io/badge/go-1.26+-00ADD8?logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

**go-siat** es un SDK robusto escrito en Go, diseñado para facilitar la integración con los servicios web SOAP del **SIAT (Sistema de Facturación de Impuestos Nacionales de Bolivia)**.

---

## 📋 Tabla de Contenidos

- [Capacidades Implementadas](#-capacidades-implementadas)
- [Guía de Inicio Rápido](#-guía-de-inicio-rápido)
- [Referencia de Uso (Tests)](#-referencia-de-uso-tests)
- [Licencia](#-licencia)

---

## 🚀 Capacidades Implementadas

El SDK ya cuenta con las siguientes implementaciones completas:

### ✅ Gestión de Códigos
- Solicitud de **CUIS** y **CUFD** (Individual y Masivo).
- Validación de **NIT**.
- Prueba de Comunicación y Lista de Certificados Revocados.

### ✅ Sincronización de Catálogos
- Sincronización de Actividades, Paramétricas, Productos y Servicios, y Documentos Sector.

### ✅ Operaciones de Punto de Venta
- Registro y Cierre de Puntos de Venta.
- Gestión de **Eventos Significativos**.

### ✅ Compra y Venta
- Recepción de Facturas (XML firmado, comprimido y codificado).
- Anulación de Facturas.
- Generación de **CUF** y **Firma Digital XML**.

---

## 🛠 Guía de Inicio Rápido

La forma más sencilla de utilizar el SDK es a través del paquete unificado `siat`.

### Instalación

```bash
go get github.com/ron86i/go-siat@v0.3.0
```

### Uso Básico

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
    // 1. Inicializar el SDK
    s, err := siat.New("https://pilotosiatservicios.impuestos.gob.bo/v2", nil)
    if err != nil {
        log.Fatal(err)
    }

    // 2. Usar un Builder para crear una solicitud
    req := models.Codigos().NewCuisBuilder().
		WithCodigoAmbiente(1).
		WithCodigoModalidad(1).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithCodigoSistema("ABC123DEF").
		WithNit(123456789).
		Build()

    // 3. Ejecutar la operación
    ctx := context.Background()
    cfg := config.Config{Token: "TU_TOKEN_API"}
    
    resp, err := s.Codigos().SolicitudCuis(ctx, cfg, req)
    if err == nil {
        log.Println("CUIS:", resp.Body.Content.RespuestaCuis.Codigo)
    }
}
```

---

## 🧪 Referencia de Uso (Tests)

La mejor referencia para entender el uso detallado de cada servicio son los **Tests de Integración**. Estos tests muestran cómo construir solicitudes usando los *Builders* y cómo procesar las respuestas del SIAT.

Puedes consultar los siguientes archivos en `internal/adapter/service/`:

- **[Códigos](./internal/adapter/service/siat_codigos_service_test.go)**: Ejemplos de CUIS, CUFD, validación de NIT y más.
- **[Sincronización](./internal/adapter/service/siat_sincronizacion_service_test.go)**: Ejemplos de sincronización de todas las paramétricas.
- **[Operaciones](./internal/adapter/service/siat_operaciones_service_test.go)**: Registro de puntos de venta y eventos significativos.
- **[Compra y Venta](./internal/adapter/service/siat_compra_venta_service_test.go)**: Recepción y anulación de facturas.

> **TIP**
> Para ejecutar estos tests, asegúrate de configurar las variables de entorno necesarias en un archivo `.env` basado en la configuración de tu ambiente del SIAT.

---

## 📄 Licencia

Este proyecto está bajo la Licencia MIT. Más información en [LICENSE](./LICENSE).
