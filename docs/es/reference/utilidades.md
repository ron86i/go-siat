# Referencia de utilidades

<p align="right">
  <a href="../../en/reference/utilities.md">🇬🇧 English</a> · <a href="../README.md">Índice de documentación</a>
</p>

Superficie completa de `github.com/ron86i/go-siat/v3/pkg/utils` — generación de CUF, firma, compresión, hash, exportación y helpers chicos de conversión.

```go
import "github.com/ron86i/go-siat/v3/pkg/utils"
```

La mayor parte de esto corre automáticamente dentro de `WithFactura` / `WithFacturas` / `WithDocumento`. Recurrí a estas funciones de forma directa cuando necesitás los artefactos intermedios o estás trabajando fuera del flujo normal de envío.

---

## Generación de CUF

El CUF es la suma de verificación fiscal por factura. Lo calculás localmente y lo ponés en la cabecera de la factura antes de enviarla.

### Builder fluido

```go
cuf, err := utils.NewCUF().
    WithNit(123456789).
    WithFechaHora(fechaEmision).
    WithSucursal(0).
    WithModalidad(siat.ModalidadElectronica).
    WithTipoEmision(siat.EmisionOnline).
    WithTipoFactura(1).
    WithTipoDocumentoSector(1).
    WithNumeroFactura(1).
    WithPuntoVenta(0).
    WithCodigoControl(cufd.CodigoControl).
    Generate()
```

| Método | Tipo |
| :--- | :--- |
| `WithNit` | `int64` |
| `WithFechaHora` | `time.Time` |
| `WithSucursal` | `int` |
| `WithModalidad` | `int` |
| `WithTipoEmision` | `int` |
| `WithTipoFactura` | `int` |
| `WithTipoDocumentoSector` | `int` |
| `WithNumeroFactura` | `int64` |
| `WithPuntoVenta` | `int` |
| `WithCodigoControl` | `string` |
| `Generate` | → `(string, error)` |

**El `fechaHora` que pasás acá tiene que ser exactamente el mismo valor que ponés en el `fechaEmision` de la cabecera.** Se hashea dentro del CUF, y cualquier desfase produce el código 1002 (`CufInvalido`). Calculá el timestamp una sola vez en una variable y reusala.

`WithCodigoControl` recibe el `CodigoControl` de la respuesta del CUFD, no el código CUFD en sí.

### Forma posicional

```go
func GenerarCUF(
    nit int64,
    fechaHora time.Time,
    sucursal, modalidad, tipoEmision, tipoFactura, tipoDocumentoSector int,
    numeroFactura int64,
    puntoVenta int,
    codigoControl string,
) (string, error)
```

Mismo resultado, diez argumentos posicionales. El builder es más seguro — varios de estos son `int` y se transponen sin darte cuenta.

---

## Firma XML

```go
type XMLSigner interface {
    SignXML(xmlBytes []byte) ([]byte, error)
}
```

`siat.Config` la implementa, y por eso `s.Config()` es lo que pasás donde se pide un firmante.

| Función | Firma |
| :--- | :--- |
| `SignWithP12` | `(xmlBytes []byte, p12Path, password string) ([]byte, error)` |
| `SignWithP12Bytes` | `(xmlBytes, p12Data []byte, password string) ([]byte, error)` |
| `SignXML` | `(xmlBytes []byte, keyPath, certPath string) ([]byte, error)` |
| `SignXMLBytes` | `(xmlBytes, keyBytes, certBytes []byte) ([]byte, error)` |

Ojo con el orden de argumentos del par PEM: `SignXML` recibe **primero la llave, después el certificado**; `SignXMLBytes` igual. Invertirlos falla al parsear.

Las firmas son XMLDSig enveloped con RSA-SHA256, el algoritmo que exige el SIAT.

### Validación de certificados

```go
func VerifyP12Expiry(p12Data []byte, password string) error
func VerifyCertificateValidity(cert *x509.Certificate) error
```

Verificá al arrancar en vez de descubrir un certificado vencido en pleno envío:

```go
p12, err := os.ReadFile("cert.p12")
if err != nil {
    log.Fatal(err)
}
if err := utils.VerifyP12Expiry(p12, password); err != nil {
    log.Fatal("certificado inutilizable: ", err)
}
```

`VerifyP12Expiry` recibe bytes, no una ruta.

---

## Compresión y hash

```go
func Gzip(data []byte) ([]byte, error)
func CompressAndHash(data []byte) (hash, encoded string, err error)
func CreateTarGz(files map[string][]byte) ([]byte, error)
func SHA256Hex(data []byte) string
func SHA512Hex(data []byte) string
```

`CompressAndHash` es la que querés para un envío: hace gzip, codifica en base64 y hashea en una sola pasada, devolviendo exactamente las dos cadenas que esperan los builders de solicitud.

```go
hash, encoded, err := utils.CompressAndHash(xmlFirmado)
builder.WithArchivo(encoded).WithHashArchivo(hash)
```

Atención al orden de retorno — primero `hash`, después `encoded`. Las dos son strings, así que intercambiarlas compila y falla recién en el SIAT con el código 969.

`CreateTarGz` recibe un mapa nombre-de-archivo a contenido y devuelve los bytes del archivo comprimido, para armar paquetes de lote a mano.

---

## Exportación

```go
func ExportXML(factura any, path string) error
func ExportSignedXML(factura any, signer XMLSigner, path string) error
func ExportTarGz(facturas []any, signer XMLSigner, path string) error
```

Escriben documentos a disco — para archivar, para envío offline, o para inspeccionar qué se manda realmente.

```go
utils.ExportXML(factura, "factura-001.xml")                        // sin firmar
utils.ExportSignedXML(factura, s.Config(), "factura-001.xml")      // firmada
utils.ExportTarGz([]any{f1, f2, f3}, s.Config(), "lote.tar.gz")    // paquete
```

Pasá `nil` como firmante a `ExportTarGz` para obtener un paquete sin firmar.

---

## Helpers de conversión

```go
func ParseIntSafe(valStr string) (int, error)
func ParseInt64Safe(valStr string) (int64, error)
func Round(v float64, decimals int) float64
```

`Round` importa más de lo que parece: el SIAT recalcula tus totales y rechaza las diferencias con los códigos 1013, 1018 y 1024. Redondeá cada monto a dos decimales de forma consistente en lugar de dejar que el error de coma flotante se acumule a lo largo de una lista de detalles larga.

### Constructores de punteros

Muchos campos de cabecera son opcionales y por eso son punteros. Estos te dan una dirección en línea:

```go
func IntPtr(v int) *int
func Int64Ptr(v int64) *int64
func Float64Ptr(v float64) *float64
```

```go
cabecera := invoices.NewCompraVentaCabeceraBuilder().
    WithCodigoPuntoVenta(utils.IntPtr(0)).
    Build()
```

No hay `StringPtr` — para esos casos tomá la dirección de una variable local:

```go
nombre := "NOMBRE DEL CLIENTE"
cabecera.WithNombreRazonSocial(&nombre)
```

Un puntero nil significa que el elemento se omite del XML. Eso tiene significado: para campos genuinamente opcionales, nil y un valor cero no son el mismo documento.

---

## Relacionado

| | |
| :--- | :--- |
| El CUF dentro de un flujo real | [Tutorial → Paso 4](../tutorial/primera-factura.md) |
| Certificados y firma | [Guía: Firmar facturas](../how-to/firmar-facturas.md) |
| Armar paquetes a mano | [Guía: Envío por lotes](../how-to/envio-lotes.md) |
| Códigos 969, 1002, 1013 | [Guía: Manejo de errores](../how-to/manejo-errores.md) |
