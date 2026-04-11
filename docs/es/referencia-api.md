# Referencia de API

[← Volver al Índice](README.md)

> Referencia completa de los 17 servicios SIAT expuestos por el SDK `go-siat`. Cada método incluye su firma, builder y navegación de respuesta.

---

## Tabla de Contenidos

1. [Inicialización del Cliente](#inicialización-del-cliente)
2. [Servicio de Códigos](#servicio-de-códigos)
3. [Servicio de Sincronización](#servicio-de-sincronización)
4. [Servicio de Operaciones](#servicio-de-operaciones)
5. [Servicio de Compra-Venta](#servicio-de-compra-venta)
6. [Servicio Electrónica](#servicio-electrónica)
7. [Servicio Computarizada](#servicio-computarizada)
8. [Servicio Documento de Ajuste](#servicio-documento-de-ajuste)
9. [Servicios de Sectores Especializados](#servicios-de-sectores-especializados)
10. [Verificación de Respuestas](#verificación-de-respuestas)

---

## Inicialización del Cliente

### `siat.New(baseUrl, httpClient)`

Crea el cliente principal del SDK.

```go
s, err := siat.New("https://pilotosiatservicios.impuestos.gob.bo/v2", nil)
```

| Parámetro | Tipo | Requerido | Descripción |
|:----------|:-----|:----------|:------------|
| `baseUrl` | `string` | ✅ | URL base de servicios SIAT |
| `httpClient` | `*http.Client` | ❌ | Cliente HTTP personalizado (nil = defaults optimizados) |

**Retorna**: `(*SiatServices, error)`

### `siat.NewWithMiddleware(baseUrl, httpClient, middlewares...)`

Crea el cliente del SDK con middlewares HTTP para logging, métricas, retry, etc.

```go
s, err := siat.NewWithMiddleware(baseURL, nil, &LoggingMiddleware{})
```

### Accesores de Servicios

```go
s.Codigos()            // → SiatCodigosService
s.Sincronizacion()     // → SiatSincronizacionService
s.Operaciones()        // → SiatOperacionesPort
s.CompraVenta()        // → SiatCompraVentaService
s.Electronica()        // → SiatElectronicaService
s.Computarizada()      // → SiatComputarizadaService
s.DocumentoAjuste()    // → SiatDocumentoAjusteService
s.Telecomunicaciones() // → SiatTelecomunicacionesService
s.ServicioBasico()     // → SiatServicioBasicoService
s.EntidadFinanciera()  // → SiatEntidadFinancieraService
s.BoletoAereo()        // → SiatBoletoAereoService
s.RecepcionCompras()   // → SiatRecepcionComprasService
```

### Métodos Auxiliares

| Método | Descripción |
|:-------|:------------|
| `s.WithTraceID(id)` | Establece un ID de traza para trazabilidad distribuida (header `X-Trace-ID`). Retorna `*SiatServices` para encadenamiento. |
| `s.WithConfig(token)` | Crea un `Config` con el token y el trace ID actual. |

---

## Servicio de Códigos

Gestiona los códigos de facturación SIAT (CUIS, CUFD), validación de NIT y comunicación de certificados.

**Acceso**: `s.Codigos()`

### `SolicitudCuis` - Solicitar CUIS

Obtiene el Código Único de Inicio de Sistemas, requerido para operar con el SIAT.

```go
req := models.Codigos().NewCuisBuilder().
    WithCodigoAmbiente(2).
    WithCodigoModalidad(1).
    WithCodigoPuntoVenta(0).
    WithCodigoSucursal(0).
    WithCodigoSistema("ABC123").
    WithNit(123456789).
    Build()

resp, err := s.Codigos().SolicitudCuis(ctx, cfg, req)
// Navegar: resp.Body.Content.RespuestaCuis.Codigo
```

| Método Builder | Tipo | Descripción |
|:---------------|:-----|:------------|
| `WithCodigoAmbiente` | `int` | 1=Producción, 2=Pruebas |
| `WithCodigoModalidad` | `int` | 1=Electrónica, 2=Computarizada |
| `WithCodigoPuntoVenta` | `int` | Punto de venta (0=principal) |
| `WithCodigoSucursal` | `int` | Número de sucursal (0=principal) |
| `WithCodigoSistema` | `string` | Código de sistema SIAT |
| `WithNit` | `int64` | NIT del contribuyente |

### `SolicitudCuisMasivo` - Solicitud Masiva de CUIS

Solicita múltiples códigos CUIS en una sola operación.

### `SolicitudCufd` - Solicitar CUFD

Obtiene el Código Único de Facturación Diaria. Requerido diariamente para emitir facturas.

```go
req := models.Codigos().NewCufdBuilder().
    WithCodigoAmbiente(2).
    WithCodigoModalidad(1).
    WithCodigoPuntoVenta(0).
    WithCodigoSucursal(0).
    WithCodigoSistema("ABC123").
    WithNit(123456789).
    WithCuis("C2FC682C").
    Build()

resp, err := s.Codigos().SolicitudCufd(ctx, cfg, req)
// Navegar: resp.Body.Content.RespuestaCufd.Codigo
// Navegar: resp.Body.Content.RespuestaCufd.CodigoControl
// Navegar: resp.Body.Content.RespuestaCufd.FechaVigCufd
```

### `SolicitudCufdMasivo` - Solicitud Masiva de CUFD

### `VerificarNit` - Validar NIT

Verifica si un Número de Identificación Tributaria está activo y habilitado.

```go
req := models.Codigos().NewVerificarNitBuilder().
    WithNit(123456789).
    Build()

resp, err := s.Codigos().VerificarNit(ctx, cfg, req)
```

### `VerificarComunicacion` - Prueba de Conectividad

### `NotificaCertificadoRevocado` - Revocación de Certificado

**Tests de integración**: [`siat_codigos_service_test.go`](../../internal/adapter/services/siat_codigos_service_test.go)

---

## Servicio de Sincronización

Sincroniza catálogos maestros: actividades económicas, tablas paramétricas, productos, documentos por sector y más.

**Acceso**: `s.Sincronizacion()`

### Métodos de Sincronización de Catálogos

Todos los métodos de sincronización comparten un patrón de builder similar:

```go
req := models.Sincronizacion().NewSincronizarActividadesBuilder().
    WithNit(nit).
    WithCodigoAmbiente(2).
    WithCodigoSucursal(0).
    WithCodigoPuntoVenta(0).
    WithCodigoSistema("ABC123").
    WithCuis("C2FC682C").
    Build()

resp, err := s.Sincronizacion().SincronizarActividades(ctx, cfg, req)
```

| Método | Descripción |
|:-------|:------------|
| `SincronizarActividades` | Catálogo de actividades económicas |
| `SincronizarListaActividadesDocumentoSector` | Mapeo de actividad a documento sector |
| `SincronizarListaLeyendasFactura` | Leyendas oficiales de facturas |
| `SincronizarListaMensajesServicios` | Catálogo de mensajes de servicio |
| `SincronizarListaProductosServicios` | Catálogo de productos y servicios |
| `SincronizarParametricaEventosSignificativos` | Tipos de eventos significativos |
| `SincronizarParametricaMotivoAnulacion` | Motivos de anulación de facturas |
| `SincronizarParametricaPaisOrigen` | Países de origen |
| `SincronizarParametricaTipoDocumentoIdentidad` | Tipos de documento de identidad |
| `SincronizarParametricaTipoDocumentoSector` | Tipos de documento por sector |
| `SincronizarParametricaTipoEmision` | Modos de emisión |
| `SincronizarParametricaTipoHabitacion` | Tipos de habitación (hotelero) |
| `SincronizarParametricaTipoMetodoPago` | Métodos de pago |
| `SincronizarParametricaTipoMoneda` | Tipos de moneda |
| `SincronizarParametricaTipoPuntoVenta` | Tipos de punto de venta |
| `SincronizarParametricaTiposFactura` | Clasificación de facturas |
| `SincronizarParametricaUnidadMedida` | Unidades de medida |
| `VerificarComunicacion` | Prueba de conectividad |

**Tests de integración**: [`siat_sincronizacion_service_test.go`](../../internal/adapter/services/siat_sincronizacion_service_test.go)

---

## Servicio de Operaciones

Gestiona el registro de puntos de venta (PV), eventos significativos y cierres de sistema.

**Acceso**: `s.Operaciones()`

| Método | Descripción |
|:-------|:------------|
| `RegistroPuntoVenta` | Registrar un nuevo punto de venta |
| `ConsultaPuntoVenta` | Consultar detalles de PV registrados |
| `CierrePuntoVenta` | Cerrar/deshabilitar un PV |
| `RegistroPuntoVentaComisionista` | Registrar un PV de comisionista |
| `RegistroEventosSignificativos` | Reportar eventos de contingencia |
| `ConsultaEventosSignificativos` | Consultar eventos de contingencia reportados |
| `CierreOperacionesSistema` | Cerrar operaciones del sistema para un período |
| `VerificarComunicacion` | Prueba de conectividad |

**Tests de integración**: [`siat_operaciones_service_test.go`](../../internal/adapter/services/siat_operaciones_service_test.go)

---

## Servicio de Compra-Venta

Maneja la facturación estándar de compra-venta - el sector más común para comercios generales.

**Acceso**: `s.CompraVenta()`

| Método | Descripción |
|:-------|:------------|
| `RecepcionFactura` | Enviar una factura individual |
| `AnulacionFactura` | Anular una factura previamente aceptada |
| `ReversionAnulacionFactura` | Revertir una anulación (solo una vez) |
| `RecepcionPaqueteFactura` | Enviar lote de hasta 500 facturas |
| `ValidacionRecepcionPaqueteFactura` | Validar estado de recepción de lote |
| `RecepcionMasivaFactura` | Envío masivo (501–1000 facturas) |
| `ValidacionRecepcionMasivaFactura` | Validar estado de recepción masiva |
| `VerificacionEstadoFactura` | Consultar el estado de una factura específica |
| `RecepcionAnexos` | Enviar anexos (adjuntos) |
| `VerificarComunicacion` | Prueba de conectividad |

**Tests de integración**: [`siat_compra_venta_service_test.go`](../../internal/adapter/services/siat_compra_venta_service_test.go)

---

## Servicio Electrónica

Maneja facturación electrónica (con firma digital) para todos los sectores.

**Acceso**: `s.Electronica()`

| Método | Descripción |
|:-------|:------------|
| `RecepcionFactura` | Enviar una factura firmada digitalmente |
| `AnulacionFactura` | Anular una factura electrónica |
| `ReversionAnulacionFactura` | Revertir una anulación |
| `RecepcionPaqueteFactura` | Enviar lote (hasta 500) |
| `ValidacionRecepcionPaqueteFactura` | Validar estado del lote |
| `RecepcionMasivaFactura` | Envío masivo (501–2000) |
| `ValidacionRecepcionMasivaFactura` | Validar estado masivo |
| `VerificacionEstadoFactura` | Consultar estado de factura |
| `RecepcionAnexosSuministroEnergia` | Anexos de suministro de energía (recargas, gift cards) |
| `VerificarComunicacion` | Prueba de conectividad |

```go
req := models.Electronica().NewRecepcionFacturaBuilder().
    WithCodigoAmbiente(2).
    WithNit(nit).
    WithCufd(cufd).
    WithCuis(cuis).
    WithTipoFacturaDocumento(1).
    WithArchivo(archivoBase64).
    WithFechaEnvio(time.Now()).
    WithHashArchivo(hash).
    Build()

resp, err := s.Electronica().RecepcionFactura(ctx, cfg, req)
```

**Tests de integración**: [`siat_electronica_service_test.go`](../../internal/adapter/services/siat_electronica_service_test.go)

---

## Servicio Computarizada

Maneja facturación computarizada sin firma digital, basada en registradoras fiscales.

**Acceso**: `s.Computarizada()`

| Método | Descripción |
|:-------|:------------|
| `RecepcionFactura` | Enviar una factura computarizada |
| `AnulacionFactura` | Anular una factura computarizada |
| `ReversionAnulacionFactura` | Revertir una anulación |
| `RecepcionPaqueteFactura` | Enviar lote (hasta 500) |
| `ValidacionRecepcionPaqueteFactura` | Validar estado del lote |
| `RecepcionMasivaFactura` | Envío masivo (501–2000) |
| `ValidacionRecepcionMasivaFactura` | Validar estado masivo |
| `VerificacionEstadoFactura` | Consultar estado de factura |
| `RecepcionAnexosSuministroEnergia` | Anexos de suministro de energía |
| `VerificarComunicacion` | Prueba de conectividad |

**Tests de integración**: [`siat_computarizada_service_test.go`](../../internal/adapter/services/siat_computarizada_service_test.go)

---

## Servicio Documento de Ajuste

Gestiona documentos de ajuste: notas de crédito/débito, notas de conciliación y reversiones.

**Acceso**: `s.DocumentoAjuste()`

| Método | Descripción |
|:-------|:------------|
| `RecepcionDocumentoAjuste` | Enviar un documento de ajuste (nota crédito/débito) |
| `AnulacionDocumentoAjuste` | Anular un documento de ajuste previamente emitido |
| `ReversionAnulacionDocumentoAjuste` | Revertir la anulación de un documento de ajuste |
| `VerificacionEstadoDocumentoAjuste` | Consultar estado del documento de ajuste |
| `VerificarComunicacion` | Prueba de conectividad |

**Tests de integración**: [`siat_documento_ajuste_service_test.go`](../../internal/adapter/services/siat_documento_ajuste_service_test.go)

---

## Servicios de Sectores Especializados

Varios sectores del SIAT requieren endpoints SOAP dedicados en lugar de usar los servicios genéricos `Electronica` o `Computarizada`. `go-siat` proporciona clientes nativos para todos ellos.

**Acceso**: 
- `s.Telecomunicaciones()`
- `s.ServicioBasico()`
- `s.EntidadFinanciera()`
- `s.BoletoAereo()`
- `s.RecepcionCompras()`

### Operaciones Comunes

Los servicios especializados siguen los mismos patrones de métodos que los generales:

| Método | Descripción |
|:-------|:------------|
| `RecepcionFactura` | Enviar una factura individual al endpoint especializado |
| `AnulacionFactura` | Anular una factura |
| `ReversionAnulacionFactura` | Revertir una anulación |
| `RecepcionPaqueteFactura` | Enviar lote (hasta 500) |
| `RecepcionMasivaFactura` | Envío masivo |
| `VerificacionEstadoFactura` | Consultar estado de factura |
| `VerificarComunicacion` | Prueba de conectividad |

> [!NOTE]
> No todos los servicios soportan todas las operaciones. Por ejemplo, `BoletoAereo` solo admite `RecepcionMasivaFactura` para envío en lote, como requiere el SIAT para las aerolíneas.

---

## Verificación de Respuestas

### `siat.Verify(response)`

Analiza una respuesta del SIAT y retorna un `*SiatError` si la operación fue rechazada:

```go
resp, err := s.Codigos().SolicitudCuis(ctx, cfg, req)
if err != nil {
    return err // Error de red
}

// Verificar la respuesta de negocio del SIAT
if err := siat.Verify(resp.Body.Content.RespuestaCuis); err != nil {
    var siatErr *siat.SiatError
    if errors.As(err, &siatErr) {
        fmt.Printf("Código SIAT: %d\n", siatErr.SiatCode)
        fmt.Printf("Reintentable: %v\n", siatErr.IsRetryable)
    }
    return err
}
```

`Verify` funciona con **cualquier** tipo de respuesta del SDK. Usa la interfaz `Result` primero, luego hace fallback a reflexión para compatibilidad.

Para detalles completos de manejo de errores, consulte la guía de [Manejo de Errores](manejo-errores.md).

---

[← Volver al Índice](README.md) | [Siguiente: Guía de Facturación →](guia-facturacion.md)
