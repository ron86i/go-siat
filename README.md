# go-siat

[![Status](https://img.shields.io/badge/status-active-success)](https://github.com/ron86i/go-siat)
[![Go Version](https://img.shields.io/badge/go-1.26+-00ADD8?logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

**go-siat** es un SDK robusto escrito en Go, diseñado para facilitar la integración con los servicios web SOAP del **SIAT (Sistema de Facturación de Impuestos Nacionales de Bolivia)**.

---

## 📋 Tabla de Contenidos

- [Capacidades Implementadas](#-capacidades-implementadas)
- [Guía de Inicio Rápido](#-guía-de-inicio-rápido)
- [Ejemplos por Módulo](#-ejemplos-por-módulo)
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
go get github.com/ron86i/go-siat@latest
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
    req := models.Codigos.NewCuisRequest().
        WithNit(123456789).
        WithCodigoAmbiente(2).
        WithCodigoSistema("ABC123XYZ").
        Build()

    // 3. Ejecutar la operación
    ctx := context.Background()
    cfg := config.Config{Token: "TU_TOKEN_API"}
    
    resp, err := s.Codigos.SolicitudCuis(ctx, cfg, req)
    if err == nil {
        log.Println("CUIS:", resp.Body.Content.RespuestaCuis.Codigo)
    }
}
```

---

## 📦 Ejemplos por Módulo

Puedes encontrar ejemplos detallados y ejecutables en la carpeta [`examples/`](./examples):

- **[Códigos](./examples/codigos/main.go)**: Flujo de CUIS y CUFD.
- **[Sincronización](./examples/sincronizacion/main.go)**: Catálogos y productos.
- **[Operaciones](./examples/operaciones/main.go)**: Puntos de venta y eventos.
- **[Compra y Venta](./examples/compra_venta/main.go)**: Firma y envío de facturas.

---

## 📄 Licencia

Este proyecto está bajo la Licencia MIT.
