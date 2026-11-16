package facturacion

import (
	"context"
	"encoding/base64"
	"encoding/xml"
	"fmt"
)

// RecepcionFactura envía una factura individual preparada al SIN y devuelve la respuesta fiscal.
func (c *Client) RecepcionFactura(ctx context.Context, r SolicitudRecepcionFactura) (*RespuestaRecepcion, error) {
	if c == nil || c.soap == nil {
		return nil, fmt.Errorf("cliente SIAT no inicializado")
	}
	if r.CodigoDocumentoSector <= 0 || r.CodigoEmision != 1 || r.CodigoModalidad <= 0 || r.CodigoSucursal < 0 || (r.CodigoPuntoVenta != nil && *r.CodigoPuntoVenta < 0) || r.Cufd == "" || r.Cuis == "" || r.TipoFacturaDocumento <= 0 || len(r.Archivo) == 0 || r.FechaEnvio.IsZero() || r.HashArchivo == "" {
		return nil, fmt.Errorf("solicitud de recepción incompleta o inválida")
	}
	if err := c.validateDestination(r.CodigoDocumentoSector, r.CodigoModalidad); err != nil {
		return nil, err
	}
	cfg := c.soap.Config()
	payload := solicitudRecepcionFacturaXML{
		CodigoAmbiente:        cfg.CodigoAmbiente,
		CodigoDocumentoSector: r.CodigoDocumentoSector,
		CodigoEmision:         r.CodigoEmision,
		CodigoModalidad:       r.CodigoModalidad,
		CodigoPuntoVenta:      r.CodigoPuntoVenta,
		CodigoSistema:         cfg.CodigoSistema,
		CodigoSucursal:        r.CodigoSucursal,
		Cufd:                  r.Cufd,
		Cuis:                  r.Cuis,
		Nit:                   cfg.Nit,
		TipoFacturaDocumento:  r.TipoFacturaDocumento,
		Archivo:               base64.StdEncoding.EncodeToString(r.Archivo),
		FechaEnvio:            r.FechaEnvio.Format("2006-01-02T15:04:05.000"),
		HashArchivo:           r.HashArchivo,
	}
	req := struct {
		XMLName   xml.Name
		Solicitud solicitudRecepcionFacturaXML
	}{
		XMLName:   xml.Name{Local: "ns:recepcionFactura"},
		Solicitud: payload,
	}
	var resp struct {
		XMLName   xml.Name
		Resultado *RespuestaRecepcion `xml:"RespuestaServicioFacturacion"`
	}
	if err := c.soap.Call(ctx, req, &resp); err != nil {
		return nil, err
	}
	if resp.Resultado == nil {
		return nil, fmt.Errorf("respuesta SIAT sin resultado de recepción")
	}
	return resp.Resultado, nil
}
