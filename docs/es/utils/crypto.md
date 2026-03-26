# Criptografía y Utilidades

Una de las capacidades más potentes de `go-siat` es el manejo nativo de los requisitos de seguridad específicos del SIAT (Firma Digital, generación de CUF y Códigos de Control).

## 🔏 Firma Digital XML (XMLDSig)

El SDK utiliza un motor embebido para firmar y verificar facturas electrónicas utilizando certificados estándar PKCS#12 (`.p12` o `.pfx`).

```go
signedXML, err := utils.SignXML(rawXML, certData, certPassword)
if err != nil {
    // Manejar error de firma
}
```

## 🆔 Generación de CUF

La generación del **Código Único de Factura (CUF)** sigue estrictamente los requerimientos algorítmicos y de codificación Base11 del SIAT.

```go
cuf, err := utils.GenerateCUF(nit, fechaCuf, sucursal, modalidad, tipoEmision, tipoFactura, puntoVenta, codigoControl)
```

## 🔐 Código de Control (Computarizada)

Para facturas computarizadas, el SDK genera el hash SHA-256 obligatorio para el envío sin firma digital.

```go
hash := utils.ComputeHash(xmlContent)
```

## 📅 Formato de Fechas SIAT

Utilidad para formatear objetos Go `time.Time` al estándar ISO 8601 exigido por el SIAT con milisegundos: `2026-03-26T14:30:15.000`.

```go
formated := utils.FormatSiatTime(time.Now()) // Ejem: "2026-03-26T14:30:15.000"
```
