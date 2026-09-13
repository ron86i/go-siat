<h1 align="center">
  <img src="../../.github/logo.svg" alt="logo de go-siat" width="250">
  <br>
  Documentación de go-siat
</h1>

<p align="center">
  <a href="../en/README.md"><img src="https://img.shields.io/badge/lang-english-red?style=flat" alt="English Version"></a>
  <a href="https://masterminds.github.io/stability/active.html"><img src="https://masterminds.github.io/stability/active.svg" alt="Stability: Active"></a>
  <br>
  <a href="https://pkg.go.dev/github.com/ron86i/go-siat/v2"><img src="https://pkg.go.dev/badge/github.com/ron86i/go-siat/v2.svg" alt="Go Reference"></a>
  <a href="https://go.dev/"><img src="https://img.shields.io/github/go-mod/go-version/ron86i/go-siat?style=flat" alt="Go Version"></a>
  <a href="../../LICENSE"><img src="https://img.shields.io/github/license/ron86i/go-siat?style=flat" alt="License"></a>
  <a href="https://github.com/ron86i/go-siat/releases"><img src="https://img.shields.io/github/v/release/ron86i/go-siat?style=flat&label=release" alt="Latest Release"></a>
  <a href="https://deepwiki.com/ron86i/go-siat"><img src="https://deepwiki.com/badge.svg" alt="Ask DeepWiki"></a>
</p>

<p align="center">
  <em>SDK profesional para integrarse con los servicios web SOAP del <b>SIAT (Sistema Integrado de Administración Tributaria)</b> de Bolivia.</em>
</p>

---

## Empezá por acá

**¿Es tu primera vez con el SDK?** Andá directo al [**tutorial**](tutorial/primera-factura.md). Te lleva desde `go get` hasta una factura aceptada por el SIAT en siete pasos, y todo lo demás se entiende mejor después.

---

## Cómo está organizada esta documentación

Estos documentos siguen [Diátaxis](https://diataxis.fr/): cuatro tipos de material, cada uno responde una pregunta distinta. Elegí la columna que corresponde a lo que necesitás ahora.

|  | **Pasos prácticos** | **Conocimiento teórico** |
| :--- | :--- | :--- |
| **Aprender**<br><sub>cuando recién empezás</sub> | 📚 **[Tutorial](tutorial/)**<br>Seguí los pasos y construí algo.<br><br>· [Tu primera factura](tutorial/primera-factura.md) | 💡 **[Explicación](explanation/)**<br>Entendé por qué funciona así.<br><br>· [Arquitectura](explanation/arquitectura.md)<br>· [Sectores](explanation/sectores.md) |
| **Trabajar**<br><sub>cuando tenés un objetivo</sub> | 🔧 **[Guías prácticas](how-to/)**<br>Resolvé un problema puntual.<br><br>· [Envío de facturas](how-to/envio-facturas.md)<br>· [Envío por lotes](how-to/envio-lotes.md)<br>· [Firmar facturas](how-to/firmar-facturas.md)<br>· [Manejo de errores](how-to/manejo-errores.md) | 📖 **[Referencia](reference/)**<br>Consultá detalles exactos.<br><br>· [API](reference/api.md)<br>· [Configuración](reference/configuracion.md)<br>· [Utilidades](reference/utilidades.md) |

---

## Buscá por lo que querés hacer

| Objetivo | Andá a |
| :--- | :--- |
| Instalar el SDK y hacer mi primera llamada | [Tutorial](tutorial/primera-factura.md) |
| Enviar una factura electrónica individual | [Tutorial → Paso 6](tutorial/primera-factura.md#paso-6-empaquetar-y-enviar) |
| Firmar facturas con mi certificado `.p12` | [Guía: Firmar facturas](how-to/firmar-facturas.md) |
| Enviar cientos de facturas en una sola petición | [Guía: Envío por lotes](how-to/envio-lotes.md) |
| Anular una factura o emitir una nota de crédito | [Guía: Envío de facturas](how-to/envio-facturas.md) |
| Saber qué builder y qué fachada usa mi rubro | [Explicación: Sectores](explanation/sectores.md) |
| Distinguir una falla de red de un rechazo del SIAT | [Guía: Manejo de errores](how-to/manejo-errores.md) |
| Consultar una firma de método o un campo de `Config` | [Referencia: API](reference/api.md) |
| Ajustar timeouts HTTP, pooling o agregar middleware | [Referencia: Configuración](reference/configuracion.md) |
| Generar un CUF, comprimir, hashear o firmar XML | [Referencia: Utilidades](reference/utilidades.md) |
| Entender el diseño hexagonal | [Explicación: Arquitectura](explanation/arquitectura.md) |

---

## Lo esencial en 60 segundos

Tres cosas sostienen casi todo el diseño del SDK. Si solo te acordás de estas, el resto lo vas leyendo sobre la marcha.

**1. La identidad se configura una sola vez.** Tu `Token`, `Nit`, `CodigoSistema`, `CodigoAmbiente` y `BaseURL` viven en `siat.Config` y se pasan a `siat.New` una única vez. Los métodos de servicio solo reciben `(ctx, req)`.

```go
s, err := siat.New(siat.Config{
    Token:          "...",
    Nit:            123456789,
    CodigoSistema:  "...",
    CodigoAmbiente: siat.AmbientePruebas,
    BaseURL:        "https://pilotosiatservicios.impuestos.gob.bo/v2",
})
```

**2. Cada llamada necesita dos verificaciones de error.** Un error `nil` no significa que el SIAT haya aceptado nada.

```go
resp, err := s.Codigos().SolicitudCuis(ctx, req)
if err != nil { /* falló la red / el transporte */ }
if err := siat.Verify(resp.Body.Content.RespuestaCuis); err != nil { /* el SIAT lo rechazó */ }
```

**3. Las solicitudes se arman con builders.** Cada tipo de solicitud tiene un constructor plano en `pkg/models` y se cierra con `.Build()`.

```go
req := models.NewCuisBuilder().
    WithCodigoSucursal(0).
    WithCodigoModalidad(siat.ModalidadElectronica).
    Build()
```

---

## Recursos adicionales

| Recurso | Ubicación |
| :--- | :--- |
| README raíz | [`README.md`](../../README.md) |
| Guía de contribución | [`CONTRIBUTING.md`](../../i18n/es/CONTRIBUTING.md) |
| Soporte y consultoría | [`SUPPORT.md`](../../i18n/es/SUPPORT.md) |
| Código de conducta | [`CODE_OF_CONDUCT.md`](../../.github/CODE_OF_CONDUCT.md) |
| Changelog | [`CHANGELOG.md`](../../CHANGELOG.md) |
| Tests de integración (ejemplos vivos) | [`internal/adapter/services/`](../../internal/adapter/services/) |
| Tests de facturas por sector | [`pkg/models/invoices/`](../../pkg/models/invoices/) |
| Licencia (MIT) | [`LICENSE`](../../LICENSE) |

---

<p align="center">
  <sub>Copyright © 2026 Ronaldo Rua — Licenciado bajo MIT</sub>
</p>
