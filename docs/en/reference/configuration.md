# Configuration

Create one client for the taxpayer and environment used by the caller.

    client, err := siat.NewClient(siat.ClientConfig{
        BaseURL:        "https://pilotosiatservicios.impuestos.gob.bo/sfd",
        Token:          os.Getenv("SIAT_TOKEN"),
        Nit:            123456789,
        CodigoSistema:  os.Getenv("SIAT_CODIGO_SISTEMA"),
        CodigoAmbiente: 2,
    })
    if err != nil {
        return err
    }

## Fields

| Field | Meaning |
| --- | --- |
| BaseURL | Base URL of the SIAT service family. For SIAT 3.0, use an official /sfd endpoint. |
| Token | Delegated SIAT token used for SOAP authentication. |
| Nit | Taxpayer NIT sent by the service clients. |
| CodigoSistema | System code registered with the SIN. |
| CodigoAmbiente | SIAT environment code configured for the taxpayer. |
| HTTPClient | Optional HTTP client for custom transport, timeouts, or test servers. |
| MaxRespuestaBytes | Optional maximum SOAP response size. |

Keep tokens and signing credentials outside source control. Supply them through the application's secret store or environment.

The same configuration is shared by the service clients returned by NewClient. It does not set fiscal document values such as CUF, CUFD, CUIS, amounts, or customer data.

See the [API reference](api.md) and [signing guide](../how-to/sign-invoices.md).
