package facturacion

import (
	"context"
	"encoding/xml"
	"fmt"
)

func (c *Client) operateInvoice(ctx context.Context, operacion, solicitudLocal string, r SolicitudEstadoFactura, motivo *int) (*RespuestaRecepcion, error) {
	if c == nil || c.soap == nil {
		return nil, fmt.Errorf("cliente SIAT no inicializado")
	}
	if r.CodigoDocumentoSector <= 0 || r.CodigoEmision < 1 || r.CodigoEmision > 3 || r.CodigoModalidad <= 0 || r.CodigoSucursal < 0 || (r.CodigoPuntoVenta != nil && *r.CodigoPuntoVenta < 0) || r.Cufd == "" || r.Cuis == "" || r.TipoFacturaDocumento <= 0 || r.Cuf == "" || (motivo != nil && *motivo <= 0) {
		return nil, fmt.Errorf("solicitud de facturación incompleta o inválida")
	}
	if err := c.validateDestination(r.CodigoDocumentoSector, r.CodigoModalidad); err != nil {
		return nil, err
	}
	cfg := c.soap.Config()
	datos := datosOperacionFacturaXML{
		XMLName:               xml.Name{Local: solicitudLocal},
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
		CodigoMotivo:          motivo,
		Cuf:                   r.Cuf,
	}
	req := struct {
		XMLName   xml.Name
		Solicitud datosOperacionFacturaXML
	}{
		XMLName:   xml.Name{Local: "ns:" + operacion},
		Solicitud: datos,
	}
	var resp struct {
		XMLName   xml.Name
		Resultado *RespuestaRecepcion `xml:"RespuestaServicioFacturacion"`
	}
	if err := c.soap.Call(ctx, req, &resp); err != nil {
		return nil, err
	}
	if resp.XMLName.Local != operacion+"Response" || resp.Resultado == nil {
		return nil, fmt.Errorf("respuesta SIAT sin resultado de %s", operacion)
	}
	return resp.Resultado, nil
}

// VerificacionEstadoFactura consulta el estado de una factura mediante su CUF.
func (c *Client) VerificacionEstadoFactura(ctx context.Context, r SolicitudEstadoFactura) (*RespuestaRecepcion, error) {
	return c.operateInvoice(ctx, "verificacionEstadoFactura", "SolicitudServicioVerificacionEstadoFactura", r, nil)
}

// AnulacionFactura solicita la anulación de una factura indicando el motivo clasificado.
func (c *Client) AnulacionFactura(ctx context.Context, r SolicitudAnulacionFactura) (*RespuestaRecepcion, error) {
	return c.operateInvoice(ctx, "anulacionFactura", "SolicitudServicioAnulacionFactura", r.SolicitudEstadoFactura, &r.CodigoMotivo)
}

// ReversionAnulacionFactura solicita revertir la anulación previa de una factura.
func (c *Client) ReversionAnulacionFactura(ctx context.Context, r SolicitudEstadoFactura) (*RespuestaRecepcion, error) {
	return c.operateInvoice(ctx, "reversionAnulacionFactura", "SolicitudServicioReversionAnulacionFactura", r, nil)
}
