# Cómo firmar facturas con certificado digital

<p align="right">
  <a href="../../en/how-to/sign-invoices.md">🇬🇧 English</a> · <a href="../README.md">Índice de documentación</a>
</p>

Las facturas en modalidad electrónica deben llevar una firma **XMLDSig** hecha con el certificado que emite el SIAT. Esta guía muestra cómo cargar ese certificado, cómo se aplica la firma y cómo manejar los casos que es fácil equivocar.

> La modalidad computarizada **no** necesita firma. Si ese es tu caso, podés saltear esta guía por completo.

---

## Cargar tu certificado

El SDK acepta los dos formatos que emite el SIAT. Ambos producen un `CredentialSign` que ponés en tu `Config`.

### PKCS#12 (`.p12` / `.pfx`)

```go
s, err := siat.New(siat.Config{
    Token:          "...",
    Nit:            123456789,
    CodigoSistema:  "...",
    CodigoAmbiente: siat.AmbientePruebas,
    BaseURL:        "https://pilotosiatservicios.impuestos.gob.bo/sfd",

    CredentialSign: siat.NewP12Credential("cert.p12", "tu-contrasena"),
})
```

### Certificado y llave separados (PEM)

```go
CredentialSign: siat.NewPEMCredential("cert.crt", "key.pem"),
```

### Cargar desde memoria en vez de disco

Ambos constructores aceptan una **ruta de archivo** (`string`) o **bytes crudos** (`[]byte`). Esto importa cuando tu certificado viene de un gestor de secretos, una variable de entorno o una base de datos en lugar del sistema de archivos:

```go
p12Bytes := obtenerDesdeVault("siat/cert")             // []byte
cred := siat.NewP12Credential(p12Bytes, os.Getenv("P12_PASSWORD"))

certBytes, keyBytes := cargarDesdeSecretos()           // []byte, []byte
cred := siat.NewPEMCredential(certBytes, keyBytes)
```

---

## Cómo ocurre la firma realmente

Rara vez vas a llamar al firmante vos mismo. `WithFactura` lo hace por vos:

```go
builder := models.NewRecepcionFacturaBuilder().
    WithCodigoModalidad(siat.ModalidadElectronica).   // ← tiene que ir primero
    WithCodigoSucursal(0).
    WithCufd(cufd).
    WithCuis(cuis)

// serializa → firma → gzip → base64 → hash, en una sola llamada
if err := builder.WithFactura(factura, s.Config()); err != nil {
    log.Fatal(err)
}
```

Dos reglas gobiernan esto, y las dos son silenciosas cuando se rompen:

**`WithCodigoModalidad` tiene que llamarse antes que `WithFactura`.** El paso de firma lee la modalidad desde la solicitud que se está construyendo. Si la modalidad todavía está en cero cuando corre `WithFactura`, concluye que no hace falta firma y envía un documento sin firmar — que el SIAT rechaza.

**`s.Config()` es lo que lleva las credenciales.** `Config` implementa la interfaz de firmante delegando en su `CredentialSign`. Pasar cualquier otra cosa, o pasar un `Config` sin credencial, significa que no hay firma.

---

## Verificar un certificado antes de depender de él

Un certificado vencido falla al momento del envío, que es el peor momento para enterarse. Verificalo al arrancar:

```go
import "github.com/ron86i/go-siat/v3/pkg/utils"

p12Bytes, err := os.ReadFile("cert.p12")
if err != nil {
    log.Fatal(err)
}
if err := utils.VerifyP12Expiry(p12Bytes, "tu-contrasena"); err != nil {
    log.Fatal("certificado inutilizable:", err)
}
```

También podés confirmar en qué formato terminó una credencial — útil cuando el origen es dinámico:

```go
cred := siat.NewP12Credential(datos, contrasena)
switch cred.GetType() {
case "P12": // PKCS#12 cargado
case "PEM": // certificado + llave cargados
case "UNKNOWN":
    log.Fatal("no se cargó ninguna credencial usable")
}
```

Que `GetType()` devuelva `"UNKNOWN"` es la señal de que una ruta estaba mal o los bytes vinieron vacíos. Los constructores no entran en pánico ante un archivo faltante — registran el error y lo exponen recién cuando intentás firmar por primera vez.

---

## Firmar un documento XML directamente

Cuando necesitás la firma fuera del flujo de facturación — para archivar una copia firmada, o para inspeccionar lo que se envía — llamá al firmante vos mismo:

```go
firmado, err := s.Config().SignXML(xmlBytes)
```

O usá los helpers de `utils`, que reciben el certificado de forma explícita:

| Función | Usala cuando |
| :------ | :----------- |
| `utils.SignWithP12(xmlBytes, p12Path, password)` | el `.p12` está en disco |
| `utils.SignWithP12Bytes(xmlBytes, p12Data, password)` | el `.p12` ya está en memoria |
| `utils.SignXML(xmlBytes, keyPath, certPath)` | el par PEM está en disco |
| `utils.SignXMLBytes(xmlBytes, keyBytes, certBytes)` | el par PEM está en memoria |

---

## Exportar una factura firmada a un archivo

Para guardar una copia firmada en disco — para archivo o para envío offline:

```go
// una factura, firmada
err := utils.ExportSignedXML(factura, s.Config(), "factura-001.xml")

// sin firmar
err := utils.ExportXML(factura, "factura-001.xml")

// varias facturas como paquete .tar.gz
err := utils.ExportTarGz([]any{f1, f2, f3}, s.Config(), "lote.tar.gz")
```

Pasá `nil` como firmante a `ExportTarGz` cuando quieras el paquete sin firmar.

---

## Resolución de problemas

| Síntoma | Causa |
| :------ | :---- |
| El SIAT rechaza por error de firma, sin error local | `WithFactura` corrió antes que `WithCodigoModalidad` |
| `no se configuraron credenciales válidas de firma digital` | `Config.CredentialSign` está en su valor cero, o fallaron ambas cargas |
| `error al leer archivo P12` | Ruta equivocada — el error aparece al firmar, no al construir |
| La firma funciona localmente pero el SIAT igual rechaza | Certificado vencido, o emitido para un NIT distinto al de `Config.Nit` |

---

## Relacionado

| | |
| :--- | :--- |
| Flujo completo incluyendo firma | [Tutorial → Paso 7](../tutorial/primera-factura.md#paso-7-cambiar-a-electrónica-con-firma-digital) |
| Helpers de firma y hashing | [Referencia: Utilidades](../reference/utilidades.md) |
| Interpretar rechazos del SIAT | [Guía: Manejo de errores](manejo-errores.md) |
