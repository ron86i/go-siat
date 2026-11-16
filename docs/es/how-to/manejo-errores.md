# Cómo manejar errores

<p align="right">
  <a href="../../en/how-to/handle-errors.md">🇬🇧 English</a> · <a href="../README.md">Índice de documentación</a>
</p>

Toda llamada al SIAT puede fallar de dos formas independientes, y el SDK las reporta por separado. Esta guía cubre cómo distinguirlas, cómo clasificar un rechazo del SIAT y cuándo tiene sentido reintentar.

---

## La regla: verificar dos veces

```go
resp, err := s.Electronica().RecepcionFactura(ctx, req)
if err != nil {
    // 1. Transporte: no hubo respuesta, o no se pudo parsear.
    return err
}
if err := siat.Verify(resp.Body.Content.RespuestaServicioFacturacion); err != nil {
    // 2. El SIAT respondió — y dijo que no.
    return err
}
```

Un error `nil` de la llamada significa que el viaje de ida y vuelta funcionó. No dice nada sobre si el SIAT aceptó tu documento. Saltear la segunda verificación es el error de integración más común contra esta API: las facturas parecen enviarse bien y nunca existen del lado del SIAT.

`siat.Verify` acepta cualquier payload de respuesta que tenga un flag `Transaccion` y una lista de mensajes. Devuelve error cuando `Transaccion` es false **o** cuando algún mensaje no es una advertencia — así que una respuesta aceptada-con-advertencias pasa.

---

## Qué te devuelve

Las dos verificaciones devuelven `*siat.SiatError`, así que un solo tipo cubre todo:

```go
var siatErr *siat.SiatError
if errors.As(err, &siatErr) {
    siatErr.Code           // string  — "NETWORK_ERROR", "SIAT_SERVER_ERROR", "AUTH_FAILED", "TIMEOUT"
    siatErr.SiatCode       // int     — el código del SIAT, 0 en errores de transporte
    siatErr.Message        // string  — los mensajes "[código] descripción" concatenados
    siatErr.IsNetworkError // bool
    siatErr.IsRetryable    // bool
    siatErr.Mensajes       // []MensajeServicio — todos los mensajes que devolvió el SIAT
    siatErr.StatusCode     // int     — estado HTTP, cuando aplica
}
```

`SiatError` implementa `Unwrap`, así que `errors.Is` llega al error de transporte de abajo:

```go
if errors.Is(err, context.DeadlineExceeded) {
    // se disparó el timeout de tu ctx
}
```

Dos helpers a nivel de paquete cubren las ramas comunes sin desenvolver nada:

```go
siat.IsRetryable(err)     // vale la pena reintentar
siat.IsNetworkError(err)  // conectividad, no rechazo
```

---

## Clasificar un código del SIAT

`SiatCode` es el campo accionable. El SDK trae la tabla completa de códigos más cuatro clasificadores, así que podés ramificar por categoría en vez de memorizar números:

| Helper | Verdadero para | Qué significa para vos |
| :--- | :--- | :--- |
| `siat.IsWarningCode(c)` | 2000–2019, 3008 | Aceptado. Registralo y seguí. |
| `siat.IsConfigCode(c)` | 910, 911, 912, 917, 958, 959, 975, 989 | Tus credenciales o tu setup están mal. Nunca reintentar. |
| `siat.IsValidationCode(c)` | 910–966, 968–985, 996–998, 1000–1061 | El documento está mal. Corregí los datos y reenviá. |
| `siat.IsRetryableCode(c)` | 123, 967, 991, 995, 999 | Problema transitorio del lado del SIAT. Esperá y reintentá. |

`siat.GetMensaje(code)` devuelve la descripción oficial de cualquier código, incluidos los que no están listados arriba.

```go
var siatErr *siat.SiatError
if errors.As(err, &siatErr) {
    switch {
    case siat.IsConfigCode(siatErr.SiatCode):
        log.Fatal("problema de configuración, reintentar no ayuda: ", siatErr)
    case siat.IsRetryableCode(siatErr.SiatCode):
        return reintentarDespues(req)
    case siat.IsValidationCode(siatErr.SiatCode):
        return fmt.Errorf("factura rechazada: %w", err)
    }
}
```

### Inspeccionar mensajes individuales

El SIAT suele devolver varios mensajes juntos. `SiatCode` guarda solo el primero que no sea advertencia, así que leé la lista cuando necesites el panorama completo:

```go
if siatErr.HasCode(siat.CodeCufdNoVigente) {
    cufd = renovarCufd()
}

for _, w := range siatErr.GetWarnings() {
    log.Printf("advertencia %d: %s", w.Codigo, w.Descripcion)
}
```

`GetWarnings()` filtra los códigos 2000–2019 y 3008. Ojo con la asimetría: si una respuesta trae *solo* advertencias, `Verify` devuelve `nil` y no hay error que inspeccionar — en ese caso leé `MensajesList` directamente de la respuesta.

---

## Códigos con los que te vas a cruzar de verdad

Hay constantes con nombre para todos los códigos (`siat.CodeCuisNoVigente` y demás). Estos son los que conviene manejar de forma explícita:

### Configuración y credenciales

| Código | Constante | Solución |
| ---: | :--- | :--- |
| 989 | `CodeTokenInvalido` | `Config.Token` incorrecto o vencido |
| 912 | `CodeSistemaNoAsociadoAlContribuyente` | `CodigoSistema` no registrado para ese NIT |
| 911 | `CodeCodigoSistemaInvalido` | `CodigoSistema` mal formado |
| 910 | `CodeAmbienteInvalido` | `CodigoAmbiente` no coincide con tu `BaseURL` |
| 958 | `CodeUsuarioNoAutorizado` | Sin autorización para ese servicio |

El código 910 casi siempre significa apuntar un `CodigoAmbiente` de producción a la URL del piloto o al revés. Los dos tienen que coincidir.

### Códigos que vencen

| Código | Constante | Solución |
| ---: | :--- | :--- |
| 3008 | `CodeWarnCuisExpira` | **Advertencia.** Pedí un CUIS nuevo pronto |
| 973 / 929 | `CodeCuisNoSeEncuentraVigente` / `CodeCuisNoVigente` | Pedí un CUIS nuevo |
| 953 | `CodeCufdNoVigente` | CUFD vencido — pedí el del día |
| 123 | `CodeCufdFueraDeTolerancia` | Desfase de reloj, o un CUFD viejo. Reintentable |

El CUFD vale un día. Un proceso de larga duración que cachea uno al arrancar empieza a fallar con 953 pasada la medianoche — renovalo de forma programada, no una sola vez.

### Firma

| Código | Constante | Solución |
| ---: | :--- | :--- |
| 921 | `CodeFirmadoXmlIncorrecto` | Documento sin firmar o firma mal formada |
| 922 | `CodeFirmaXmlNoCorrespondeAlContribuyente` | Certificado emitido para otro NIT |
| 927 | `CodeCertificadoFirmaInvalido` | Certificado inválido o vencido |
| 928 | `CodeCertificadoRevocado` | Certificado revocado |
| 965 | `CodeContribuyenteSinFirmaVigente` | Sin firma vigente registrada en el SIAT |

Un 921 indica que el SIN rechazó el XML firmado. Define `CodigoModalidad` en la solicitud tipada antes de llamar a `PrepararFactura` y pasa un XMLSigner para modalidad electrónica; consulta [Firmar facturas](firmar-facturas.md).

### Contenido del documento

| Código | Constante | Solución |
| ---: | :--- | :--- |
| 939 | `CodeFacturaNoCumpleXsd` | Campo obligatorio faltante o mal formado |
| 1002 | `CodeCufInvalido` | CUF mal — normalmente por `fechaEmision` que no coincide |
| 1000 / 952 | `CodeCufYaExisteEnSin` / `CodeCufYaRegistradoEnSin` | Ese CUF ya fue enviado |
| 969 | `CodeHashInvalido` | El hash no coincide con el archivo |
| 1013 / 1024 | `CodeCalculoMontoTotalErroneo` / `CodeSumatoriaDetallesErronea` | Los totales no cierran |
| 931 / 932 / 940 | Códigos de sector | Mirá [Sectores](../explanation/sectores.md) |

Para el 1002, la causa habitual es generar el CUF con un timestamp y poner otro distinto en `fechaEmision`. Calculá la hora una sola vez y reusá la variable.

### Transitorios

| Código | Constante |
| ---: | :--- |
| 967 | `CodeTiempoEsperaAgotadoDB` |
| 991 | `CodeErrorBaseDeDatos` |
| 995 | `CodeServicioNoDisponible` |
| 999 | `CodeErrorEjecucionServicio` |

---

## Reintentar con criterio

Reintentá solo lo que `IsRetryable` marca. Los errores de red y los timeouts califican; un documento rechazado no — reenviar los mismos datos malos produce el mismo rechazo.

```go
func conReintentos(ctx context.Context, intentos int, fn func() error) error {
    var err error
    for i := range intentos {
        if err = fn(); err == nil {
            return nil
        }
        if !siat.IsRetryable(err) {
            return err
        }
        select {
        case <-time.After(time.Second << i):   // backoff exponencial
        case <-ctx.Done():
            return ctx.Err()
        }
    }
    return err
}
```

Dos advertencias propias de esta API:

**Reintentar un envío no es idempotente.** Si el primer intento llegó al SIAT y solo se perdió la respuesta, reenviar el mismo CUF devuelve 1000 (`CufYaExisteEnSin`). Tratá ese código como "en realidad sí entró", no como una falla.

**Acotá siempre la llamada con un deadline de contexto.** El SIAT puede ponerse lento bajo carga, y sin deadline una petición trabada bloquea indefinidamente:

```go
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()
```

---

## Faults SOAP

Un sobre mal formado vuelve como fault SOAP en vez de rechazo normal. Aparece en la respuesta, no en `err`:

```go
if resp.Body.Fault != nil {
    log.Fatal("fault SOAP: ", resp.Body.Fault.String())
}
```

Un fault significa que la petición nunca llegó a la lógica de negocio del SIAT — normalmente un endpoint equivocado o un payload corrupto. Verificalo antes de `Verify` cuando estés depurando una integración nueva.

---

## Relacionado

| | |
| :--- | :--- |
| Fallas de firma en detalle | [Guía: Firmar facturas](firmar-facturas.md) |
| Códigos de rechazo por sector | [Explicación: Sectores](../explanation/sectores.md) |
| Timeouts y ajuste de HTTP | [Referencia: Configuración](../reference/configuracion.md) |
| API completa de errores | [Referencia: API](../reference/api.md) |
