# Firmar facturas electrónicas

Los documentos electrónicos deben firmarse antes del envío. El método de preparación recibe un XMLSigner e incluye la firma en la solicitud. Las facturas computarizadas no se firman.

## Cargar un firmador PKCS#12

    p12, err := os.ReadFile(os.Getenv("SIAT_P12_PATH"))
    if err != nil {
        return err
    }
    firmador, err := utils.NewXMLDocumentSignerFromP12(
        p12,
        os.Getenv("SIAT_P12_PASSWORD"),
    )
    if err != nil {
        return err
    }

Guarda el archivo PKCS#12 y su contraseña en un gestor de secretos. No registres esos valores en logs.

## Preparar una factura electrónica

La factura usa el modelo público del sector, por ejemplo documents/compraventa. Define la modalidad en la solicitud antes de prepararla.

    solicitud := facturacion.SolicitudRecepcionFactura{
        CodigoModalidad: facturacion.ModalidadElectronica,
        Cuis:            cuis,
    }
    solicitud, err = cliente.Facturacion.PrepararFactura(solicitud, factura, firmador)
    if err != nil {
        return err
    }
    respuesta, err := cliente.Facturacion.RecepcionFactura(ctx, solicitud)

La preparación valida el documento, serializa el XML sectorial, lo firma, comprime y calcula el hash de la solicitud. Envía el valor preparado que devuelve el método.

## Verificaciones previas

Comprueba que el certificado esté vigente para la firma, los datos del documento correspondan al CUF y la solicitud use modalidad electrónica. Consulta [configuración](../reference/configuracion.md) y el [ejemplo Compra-Venta](../../../examples/compraventa/main.go).
