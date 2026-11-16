# Configuración

Crea un cliente por contribuyente y ambiente que usará la aplicación.

    cliente, err := siat.NewClient(siat.ClientConfig{
        BaseURL:        "https://pilotosiatservicios.impuestos.gob.bo/sfd",
        Token:          os.Getenv("SIAT_TOKEN"),
        Nit:            123456789,
        CodigoSistema:  os.Getenv("SIAT_CODIGO_SISTEMA"),
        CodigoAmbiente: 2,
    })
    if err != nil {
        return err
    }

## Campos

| Campo | Significado |
| --- | --- |
| BaseURL | URL base de la familia de servicios SIAT. Para SIAT 3.0 usa una ruta oficial /sfd. |
| Token | Token delegado SIAT para autenticación SOAP. |
| Nit | NIT que los clientes de servicio envían al SIN. |
| CodigoSistema | Código del sistema registrado ante el SIN. |
| CodigoAmbiente | Código del ambiente SIAT configurado para el contribuyente. |
| HTTPClient | Cliente HTTP opcional para transporte, timeouts o servidores de prueba. |
| MaxRespuestaBytes | Tamaño máximo opcional de la respuesta SOAP. |

Mantén los tokens y credenciales de firma fuera del código fuente. Cárgalos desde el gestor de secretos o el entorno de la aplicación.

NewClient comparte la configuración con los clientes de servicios. No establece valores fiscales del documento como CUF, CUFD, CUIS, importes ni datos del cliente.

Consulta la [referencia de API](api.md) y la [guía de firma](../how-to/firmar-facturas.md).
