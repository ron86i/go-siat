# Utilidades

[← Volver al Índice](README.md)

> Referencia del paquete `pkg/utils`, que proporciona firma criptográfica, generación de CUF, compresión, hash y helpers de parseo.

---

## Tabla de Contenidos

1. [Firma Digital XML](#firma-digital-xml)
2. [Generación de CUF](#generación-de-cuf)
3. [Compresión y Codificación](#compresión-y-codificación)
4. [Hash Criptográfico](#hash-criptográfico)
5. [Helpers de Parseo](#helpers-de-parseo)

---

## Firma Digital XML

**Paquete**: `github.com/ron86i/go-siat/pkg/utils`

El SDK proporciona múltiples formas de firmar digitalmente documentos XML de facturas. Todos los métodos usan **RSA-SHA256** con **Firma Envolvente** (C14N 1.0 con Comentarios), según lo exigido por el SIAT.

### `SignXML` - Firmar desde Archivos PEM

Lee certificado y clave del sistema de archivos.

```go
func SignXML(xmlBytes []byte, keyPath, certPath string) ([]byte, error)
```

| Parámetro | Tipo | Descripción |
|:----------|:-----|:------------|
| `xmlBytes` | `[]byte` | XML de factura serializado (de `xml.Marshal()`) |
| `keyPath` | `string` | Ruta al archivo de clave privada PEM |
| `certPath` | `string` | Ruta al archivo de certificado PEM |

```go
signedXML, err := utils.SignXML(xmlData, "key.pem", "cert.crt")
```

### `SignXMLBytes` - Firmar desde Bytes PEM

Recibe certificado y clave directamente como slices de bytes. Ideal para cargar desde bases de datos o vaults de secretos.

```go
func SignXMLBytes(xmlBytes, keyBytes, certBytes []byte) ([]byte, error)
```

### `SignWithP12` - Firmar desde Archivo P12/PFX

Usa un contenedor PKCS#12 (combina clave y certificado).

```go
func SignWithP12(xmlBytes []byte, p12Path, password string) ([]byte, error)
```

```go
signedXML, err := utils.SignWithP12(xmlData, "cert.p12", "mi_contraseña")
```

### `SignWithP12Bytes` - Firmar desde Bytes P12

Recibe los datos P12 directamente en memoria. Ideal cuando los certificados se almacenan como BLOBs en bases de datos o se obtienen de un vault.

```go
func SignWithP12Bytes(xmlBytes, p12Data []byte, password string) ([]byte, error)
```

### `VerifyP12Expiry` - Verificar Validez del Certificado

Valida que el certificado dentro de un archivo P12 no haya expirado.

```go
func VerifyP12Expiry(p12Data []byte, password string) error
```

```go
if err := utils.VerifyP12Expiry(p12Data, "contraseña"); err != nil {
    log.Printf("Problema con certificado: %v", err)
    // "certificate expired on: 2026-01-15" o
    // "certificate is not yet valid (starts: 2026-06-01)"
}
```

### `VerifyCertificateValidity` - Verificar Certificado x509

Valida un certificado x509 ya parseado.

```go
func VerifyCertificateValidity(cert *x509.Certificate) error
```

> [!IMPORTANT]
> Todas las funciones de firma validan automáticamente el certificado antes de firmar. Si el certificado está expirado o aún no es válido, la firma fallará con un error descriptivo.

### Formatos de Clave Soportados

| Formato | Soporte |
|:--------|:--------|
| PKCS#1 (RSA PRIVATE KEY) | ✅ |
| PKCS#8 (PRIVATE KEY) | ✅ |
| PKCS#12 (.p12 / .pfx) | ✅ |
| Claves EC | ❌ (solo RSA) |

---

## Generación de CUF

El CUF (Código Único de Factura) es un identificador único para cada factura, calculado usando un algoritmo específico exigido por el SIAT.

### `GenerarCUF`

```go
func GenerarCUF(
    nit int64,
    fechaHora time.Time,
    sucursal, modalidad, tipoEmision, tipoFactura,
    tipoDocumentoSector, numeroFactura, puntoVenta int,
    codigoControl string,
) (string, error)
```

| Parámetro | Tipo | Dígitos | Descripción |
|:----------|:-----|:--------|:------------|
| `nit` | `int64` | 13 | NIT del contribuyente (relleno con ceros) |
| `fechaHora` | `time.Time` | 17 | Fecha/hora de emisión (yyyyMMddHHmmssSSS) |
| `sucursal` | `int` | 4 | Número de sucursal (relleno con ceros) |
| `modalidad` | `int` | 1 | 1=Electrónica, 2=Computarizada |
| `tipoEmision` | `int` | 1 | 1=En línea, 2=Fuera de línea, 3=Masiva |
| `tipoFactura` | `int` | 1 | Tipo de factura |
| `tipoDocumentoSector` | `int` | 2 | Tipo de documento sector |
| `numeroFactura` | `int` | 10 | Número de factura (relleno con ceros) |
| `puntoVenta` | `int` | 4 | Punto de venta (relleno con ceros) |
| `codigoControl` | `string` | Variable | Código de control de la respuesta CUFD |

### Algoritmo

1. **Formatear** cada campo a su longitud fija con relleno de ceros.
2. **Concatenar** todos los campos (53 dígitos en total).
3. **Calcular** dígito verificador Módulo 11 (pesos 2-9).
4. **Convertir** la cadena numérica a hexadecimal (Base16).
5. **Agregar** el código de control del CUFD.

```go
cuf, err := utils.GenerarCUF(
    123456789,                         // NIT
    time.Now(),                        // Fecha/hora de emisión
    0,                                 // Sucursal
    siat.ModalidadElectronica,         // Modalidad (1)
    siat.EmisionOnline,                // Tipo emisión (1)
    1,                                 // Tipo factura
    1,                                 // Tipo documento sector
    42,                                // Número de factura
    0,                                 // Punto de venta
    cufdControl,                       // Código de control del CUFD
)
```

---

## Compresión y Codificación

### `Gzip` - Comprimir Datos

```go
func Gzip(data []byte) ([]byte, error)
```

Compresión Gzip estándar.

### `CompressAndHash` - Comprimir + SHA256 + Base64

La función todo-en-uno para preparar datos de factura para transmisión. Combina los tres pasos que requiere el SIAT:

```go
func CompressAndHash(data []byte) (hash, encoded string, err error)
```

| Retorno | Tipo | Descripción |
|:--------|:-----|:------------|
| `hash` | `string` | Hash SHA-256 hexadecimal de los datos **comprimidos** |
| `encoded` | `string` | Codificación Base64 de los datos comprimidos |
| `err` | `error` | Cualquier error de compresión |

```go
hash, archivoBase64, err := utils.CompressAndHash(signedXML)
// hash → se usa en WithHashArchivo()
// archivoBase64 → se usa en WithArchivo()
```

### `CreateTarGz` - Crear Archivo TAR.GZ

Crea un archivo TAR.GZ desde un mapa de nombres de archivo a contenidos.

```go
func CreateTarGz(files map[string][]byte) ([]byte, error)
```

```go
files := map[string][]byte{
    "factura_001.xml": xmlBytes1,
    "factura_002.xml": xmlBytes2,
}
archive, err := utils.CreateTarGz(files)
```

---

## Hash Criptográfico

### `SHA256Hex`

```go
func SHA256Hex(data []byte) string
```

Retorna el hash SHA-256 como cadena hexadecimal en minúsculas.

### `SHA512Hex`

```go
func SHA512Hex(data []byte) string
```

Retorna el hash SHA-512 como cadena hexadecimal en minúsculas.

---

## Helpers de Parseo

### `ParseIntSafe`

Convierte de forma segura un string a `int`, eliminando espacios y proporcionando errores descriptivos.

```go
func ParseIntSafe(valStr string) (int, error)
```

```go
ambiente, err := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
```

### `ParseInt64Safe`

Convierte de forma segura un string a `int64`. Ideal para valores de NIT.

```go
func ParseInt64Safe(valStr string) (int64, error)
```

```go
nit, err := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
```

### Helpers de Punteros

Funciones que retornan punteros a valores. Útiles para campos opcionales en builders de facturas:

```go
func Float64Ptr(v float64) *float64
func Int64Ptr(v int64) *int64
func IntPtr(v int) *int
```

```go
cabecera := invoices.NewCompraVentaCabeceraBuilder().
    WithMontoDescuento(utils.Float64Ptr(10.50)).   // Campo opcional
    WithNombreRazonSocial(nil).                     // Explícitamente nil
    Build()
```

---

[← Volver al Índice](README.md) | [Siguiente: Configuración →](configuracion.md)
