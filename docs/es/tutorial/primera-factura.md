# Tutorial: De cero a tu primera factura

<p align="right">
  <a href="../../en/tutorial/first-invoice.md">🇬🇧 English</a>
</p>

En este tutorial vas a enviar una factura real al ambiente piloto del SIAT y vas a recibir un código de aceptación.

Lo harás en siete pasos, y al terminar vas a entender la cadena que toda factura electrónica boliviana debe seguir: **CUIS → CUFD → CUF → factura → envío**.

> Este es un ejercicio de *aprendizaje*. Usamos el ambiente piloto, así que nada de lo que hagas acá tiene validez tributaria. Cuando termines, las [guías prácticas](../how-to/) te muestran cómo manejar estos mismos pasos en producción.

---

## Qué necesitás antes de empezar

| Requisito | Notas |
| :-------- | :---- |
| Go 1.26+ | Verificá con `go version` |
| Un token del SIAT | Lo obtenés desde tu portal de contribuyente |
| Tu NIT | El número de contribuyente al que pertenece el token |
| Un código de sistema autorizado | `CodigoSistema`, que el SIAT emite al registrar tu sistema |
| Un certificado digital | Solo para la modalidad *electrónica*. Empezamos con **computarizada**, que no lo necesita |

Arrancamos a propósito con la modalidad **computarizada** para que puedas completar todo el flujo sin certificado. Agregar la firma es un cambio de una línea que hacemos al final.

---

## Paso 1: Instalar y configurar el cliente

```bash
go get github.com/ron86i/go-siat/v2
```

Creá `main.go`. La idea más importante de este SDK: **tu identidad se configura una sola vez**, no se repite en cada llamada.

```go
package main

import (
    "context"
    "log"
    "time"

    "github.com/ron86i/go-siat/v2"
)

const urlPiloto = "https://pilotosiatservicios.impuestos.gob.bo/v2"

func main() {
    s, err := siat.New(siat.Config{
        Token:          "TU_TOKEN_SIAT",
        Nit:            123456789,
        CodigoSistema:  "TU_CODIGO_SISTEMA",
        CodigoAmbiente: siat.AmbientePruebas,
        BaseURL:        urlPiloto,
    })
    if err != nil {
        log.Fatal(err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    _ = s // lo empezamos a usar en el siguiente paso
}
```

`siat.New` valida la configuración de inmediato. Si olvidás el token o pasás un código de ambiente inválido, el error aparece acá y no como una falla confusa más adelante.

> **Siempre pasá un contexto con timeout.** Los servidores del SIAT pueden responder lento, y un `context.Background()` pelado deja la petición colgada indefinidamente.

---

## Paso 2: Solicitar el CUIS

El **CUIS** (*Código Único de Inicio de Sistemas*) identifica tu sistema ante el SIAT. Todas las llamadas posteriores lo necesitan.

```go
cuisReq := models.NewCuisBuilder().
    WithCodigoSucursal(0).      // 0 = casa matriz
    WithCodigoPuntoVenta(0).    // 0 = punto de venta por defecto
    WithCodigoModalidad(siat.ModalidadComputarizada).
    Build()

cuisResp, err := s.Codigos().SolicitudCuis(ctx, cuisReq)
if err != nil {
    log.Fatal("error de transporte:", err)
}
if err := siat.Verify(cuisResp.Body.Content.RespuestaCuis); err != nil {
    log.Fatal("el SIAT rechazó la solicitud de CUIS:", err)
}

cuis := cuisResp.Body.Content.RespuestaCuis.Codigo
log.Println("CUIS:", cuis)
```

Agregá `"github.com/ron86i/go-siat/v2/pkg/models"` a tus imports.

Fijate en la **verificación de error en dos etapas**. Este patrón se repite en absolutamente todas las llamadas que vas a hacer:

1. `err != nil` significa que la petición nunca se completó — falla de red, timeout, URL mal.
2. `siat.Verify(...)` significa que la petición se completó, pero el SIAT la *rechazó* — NIT inválido, código vencido, sistema no autorizado.

Saltarse la segunda verificación es el error más común al integrar con el SIAT: tu programa parece funcionar mientras el SIAT en realidad rechazó la operación.

---

## Paso 3: Solicitar el CUFD

El **CUFD** (*Código Único de Facturación Diaria*) es un código diario. Vence, así que los sistemas en producción piden uno nuevo cada día y lo cachean.

```go
cufdReq := models.NewCufdBuilder().
    WithCuis(cuis).
    WithCodigoSucursal(0).
    WithCodigoPuntoVenta(0).
    WithCodigoModalidad(siat.ModalidadComputarizada).
    Build()

cufdResp, err := s.Codigos().SolicitudCufd(ctx, cufdReq)
if err != nil {
    log.Fatal(err)
}
if err := siat.Verify(cufdResp.Body.Content.RespuestaCufd); err != nil {
    log.Fatal(err)
}

cufd := cufdResp.Body.Content.RespuestaCufd.Codigo
codigoControl := cufdResp.Body.Content.RespuestaCufd.CodigoControl
```

Esta llamada devuelve **dos** valores que necesitás: el código en sí y un `CodigoControl`. Guardá ambos — el código de control alimenta directamente el siguiente paso.

---

## Paso 4: Generar el CUF

El **CUF** (*Código Único de Facturación*) es un dígito verificador que identifica unívocamente una factura. A diferencia del CUIS y el CUFD, este lo calculás localmente — sin llamada de red.

```go
import "github.com/ron86i/go-siat/v2/pkg/utils"

fechaEmision := time.Now()

cuf, err := utils.NewCUF().
    WithNit(123456789).
    WithFechaHora(fechaEmision).
    WithSucursal(0).
    WithModalidad(siat.ModalidadComputarizada).
    WithTipoEmision(siat.EmisionOnline).
    WithTipoFactura(1).            // 1 = factura con derecho a crédito fiscal
    WithTipoDocumentoSector(1).    // 1 = sector compra-venta
    WithNumeroFactura(1).          // tu propio correlativo
    WithPuntoVenta(0).
    WithCodigoControl(codigoControl).
    Generate()
if err != nil {
    log.Fatal(err)
}
```

`WithNumeroFactura` es **tu** número correlativo de factura. Ese contador es tuyo — el SIAT no lo asigna. Nunca debe repetirse para el mismo punto de venta.

> La `fechaEmision` exacta que pasás acá tiene que ser el mismo timestamp que pongas en la cabecera de la factura en el siguiente paso. Si difieren, el CUF queda inválido.

---

## Paso 5: Construir el documento de la factura

Ahora la factura en sí. Tiene una **cabecera** (quién, cuándo, totales) y una o más **líneas de detalle** (qué se vendió).

```go
import "github.com/ron86i/go-siat/v2/pkg/models/invoices"

nombreCliente := "JUAN PEREZ"
codigoPuntoVenta := 0

cabecera := invoices.NewCompraVentaCabeceraBuilder().
    WithNitEmisor(123456789).
    WithRazonSocialEmisor("NOMBRE DE TU EMPRESA").
    WithMunicipio("Tarija").
    WithNumeroFactura(1).                 // el mismo número usado en el CUF
    WithCuf(cuf).
    WithCufd(cufd).
    WithCodigoSucursal(0).
    WithDireccion("AVENIDA LA PAZ #123").
    WithCodigoPuntoVenta(&codigoPuntoVenta).
    WithFechaEmision(fechaEmision).       // el mismo timestamp usado en el CUF
    WithNombreRazonSocial(&nombreCliente).
    WithCodigoTipoDocumentoIdentidad(1).  // 1 = CI
    WithNumeroDocumento("5115889").
    WithCodigoCliente("1").
    WithCodigoMetodoPago(1).              // 1 = efectivo
    WithMontoTotal(100).
    WithMontoTotalSujetoIva(100).
    WithCodigoMoneda(1).                  // 1 = BOB
    WithTipoCambio(1).
    WithMontoTotalMoneda(100).
    WithLeyenda("Ley N° 453: Tienes derecho a recibir información...").
    WithUsuario("cajero01").
    WithCodigoDocumentoSector(1).
    Build()

detalle := invoices.NewCompraVentaDetalleBuilder().
    WithActividadEconomica("477300").
    WithCodigoProductoSin(622539).
    WithCodigoProducto("abc123").
    WithDescripcion("GASA").
    WithCantidad(1).
    WithUnidadMedida(1).
    WithPrecioUnitario(100).
    WithSubTotal(100).
    Build()

factura := invoices.NewCompraVentaBuilder().
    WithModalidad(siat.ModalidadComputarizada).
    WithCabecera(cabecera).
    AddDetalle(detalle).
    Build()
```

Varios campos acá son **códigos de catálogo**, no texto libre: `WithCodigoMetodoPago`, `WithUnidadMedida`, `WithActividadEconomica`, `WithCodigoProductoSin`. El SIAT publica los valores válidos y los obtenés con el servicio `Sincronizacion()`. Usar un valor que no está en el catálogo hace que la factura sea rechazada.

---

## Paso 6: Empaquetar y enviar

Acá es donde el SDK te ahorra más trabajo. `WithFactura` serializa la factura a XML, la firma cuando corresponde, la comprime en gzip, la codifica en base64 y calcula el hash SHA-256 — todo en una sola llamada.

```go
builder := models.NewRecepcionFacturaBuilder().
    WithCodigoModalidad(siat.ModalidadComputarizada).
    WithCodigoSucursal(0).
    WithCodigoPuntoVenta(0).
    WithCodigoDocumentoSector(1).
    WithCodigoEmision(siat.EmisionOnline).
    WithTipoFacturaDocumento(1).
    WithCuis(cuis).
    WithCufd(cufd).
    WithFechaEnvio(fechaEmision)

// Adjuntar la factura. s.Config() actúa como firmante.
if err := builder.WithFactura(factura, s.Config()); err != nil {
    log.Fatal("no se pudo empaquetar la factura:", err)
}

req := builder.Build()

resp, err := s.CompraVenta().RecepcionFactura(ctx, req)
if err != nil {
    log.Fatal(err)
}
if err := siat.Verify(resp.Body.Content.RespuestaServicioFacturacion); err != nil {
    log.Fatal("el SIAT rechazó la factura:", err)
}

log.Println("Aceptada. Código de recepción:", resp.Body.Content.RespuestaServicioFacturacion.CodigoRecepcion)
```

Dos detalles que vale la pena mirar con calma:

**`WithFactura` tiene que ir después de `WithCodigoModalidad`.** Lee la modalidad que configuraste para decidir si el documento necesita firma digital. Llamarlo antes significa que nunca firma.

**`s.Config()` es el firmante.** `Config` lleva tu `CredentialSign` e implementa la interfaz de firma, así que pasarlo entrega las credenciales al paso de empaquetado. En modalidad computarizada no se firma nada, pero pasarlo deja el código idéntico para cuando cambies a electrónica.

---

## Paso 7: Cambiar a electrónica (con firma digital)

Ya tenés un flujo funcionando. Convertirlo en una factura electrónica firmada legalmente son dos cambios.

Primero, agregá tu certificado a la configuración:

```go
s, err := siat.New(siat.Config{
    Token:          "TU_TOKEN_SIAT",
    Nit:            123456789,
    CodigoSistema:  "TU_CODIGO_SISTEMA",
    CodigoAmbiente: siat.AmbientePruebas,
    BaseURL:        urlPiloto,

    // archivo .p12 / .pfx más su contraseña
    CredentialSign: siat.NewP12Credential("cert.p12", "contrasena-p12"),
})
```

Segundo, reemplazá cada `siat.ModalidadComputarizada` por `siat.ModalidadElectronica`, y llamá a `s.Electronica()` en lugar de `s.CompraVenta()`.

Esa es toda la diferencia. Como `WithFactura` ya recibe `s.Config()`, ahora encuentra una credencial, ve la modalidad electrónica y firma el XML con XMLDSig antes de comprimirlo.

> Si tu certificado es un par certificado/llave separado en vez de un `.p12`, usá `siat.NewPEMCredential("cert.crt", "key.pem")`.

---

## Qué aprendiste

Construiste la cadena completa que exige toda factura electrónica boliviana:

```
Config ──► CUIS ──► CUFD ──► CUF ──► XML factura ──► Envío
una vez   una vez   diario   por factura  por factura
```

También conociste los tres hábitos que hacen confiable una integración con el SIAT:

- **Verificá el error dos veces** — error de transporte, después `siat.Verify`.
- **Reutilizá el timestamp** — la misma `fechaEmision` en el CUF y en la cabecera.
- **El orden importa** — `WithCodigoModalidad` antes de `WithFactura`.

## Hacia dónde seguir

| Si querés… | Leé |
| :--------- | :-- |
| Firmar con certificado, en detalle | [Guía: Firmar facturas](../how-to/firmar-facturas.md) |
| Enviar cientos de facturas de una vez | [Guía: Envío por lotes](../how-to/envio-lotes.md) |
| Manejar rechazos y reintentos bien | [Guía: Manejo de errores](../how-to/manejo-errores.md) |
| Facturar en un sector distinto a compra-venta | [Explicación: Sectores](../explanation/sectores.md) |
| Consultar una firma de método exacta | [Referencia: API](../reference/api.md) |
| Entender cómo está armado el SDK | [Explicación: Arquitectura](../explanation/arquitectura.md) |
