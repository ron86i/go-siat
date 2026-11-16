# Utilidades

La mayoría de las aplicaciones debe usar los modelos sectoriales y clientes de protocolo tipados. Estas utilidades están disponibles cuando se necesita ejecutar una operación de nivel bajo.

## CUF

Usa CUFParams con campos nombrados para evitar confundir argumentos numéricos posicionales.

    cuf, err := (utils.CUFParams{
        Nit:                 nit,
        FechaHora:           fechaEmision,
        Sucursal:            sucursal,
        Modalidad:           modalidad,
        TipoEmision:         tipoEmision,
        TipoFactura:         tipoFactura,
        TipoDocumentoSector: sector,
        NumeroFactura:       numeroFactura,
        PuntoVenta:          puntoVenta,
        CodigoControl:       codigoControl,
    }).Generate()

Los datos del CUF deben coincidir con el documento y el CUFD vigente. Consulta la [guía de facturas](../how-to/envio-facturas.md).

## XML y compresión

| Función | Uso |
| --- | --- |
| MarshalIndentedXML | Serializar un documento Go a XML legible cuando el modelo admite serialización XML ordinaria. |
| Gzip | Comprimir un conjunto de bytes con GZIP. |
| CreateTarGz | Crear un archivo TAR.GZ a partir de nombres y contenidos. |
| CompressAndHash | Devolver el hash SHA-256 y los datos GZIP codificados en Base64. |
| SHA256Hex | Calcular un digest SHA-256 hexadecimal. |

Los métodos de preparación de facturación realizan la serialización sectorial y preparan la solicitud para las operaciones compatibles. Prefiérelos a reconstruir manualmente el payload.

## Firma

NewXMLDocumentSignerFromP12 crea un firmador reutilizable desde bytes PKCS#12. SignXMLBytes y SignWithP12Bytes son utilidades de nivel bajo para aplicaciones que gestionan XML por su cuenta.

Consulta la [guía de firma](../how-to/firmar-facturas.md) y la documentación del paquete para ver todas las funciones públicas.
