package facturacion

import (
	"context"
	"encoding/base64"
	"encoding/xml"
	"fmt"

	"github.com/ron86i/go-siat/v3/internal/siatpolicy"
)

// PrepararMasiva prepara un TAR.GZ de facturas de cualquier sector para emisión masiva en SIAT v3.
func (r SolicitudRecepcionMasiva) PrepararMasiva(ctx context.Context, facturas []Documento, signer XMLSigner) (SolicitudRecepcionMasiva, error) {
	prepared, count, err := prepararDocumentos(ctx, r.SolicitudRecepcionFactura, facturas, signer, 3, siatpolicy.MaxFacturasMasiva)
	if err != nil {
		return r, err
	}
	r.SolicitudRecepcionFactura = prepared
	r.CantidadFacturas = count
	return r, nil
}

// WithFacturas prepara un TAR.GZ de documentos v3 seleccionados para emisión masiva.
func (r SolicitudRecepcionMasiva) WithFacturas(ctx context.Context, facturas []Documento, signer XMLSigner) (SolicitudRecepcionMasiva, error) {
	return r.PrepararMasiva(ctx, facturas, signer)
}

// RecepcionMasivaFactura envía el lote de facturas masivas preparadas.
func (c *Client) RecepcionMasivaFactura(ctx context.Context, r SolicitudRecepcionMasiva) (*RespuestaRecepcion, error) {
	if c == nil || c.soap == nil {
		return nil, fmt.Errorf("cliente SIAT no inicializado")
	}
	p := r.SolicitudRecepcionFactura
	if p.CodigoDocumentoSector <= 0 || p.CodigoEmision != 3 || p.CodigoModalidad <= 0 || p.CodigoSucursal < 0 || (p.CodigoPuntoVenta != nil && *p.CodigoPuntoVenta < 0) || p.Cufd == "" || p.Cuis == "" || p.TipoFacturaDocumento <= 0 || len(p.Archivo) == 0 || p.FechaEnvio.IsZero() || p.HashArchivo == "" || r.CantidadFacturas <= 0 || r.CantidadFacturas > siatpolicy.MaxFacturasMasiva {
		return nil, fmt.Errorf("solicitud de recepción masiva incompleta o inválida")
	}
	if err := c.validateDestination(p.CodigoDocumentoSector, p.CodigoModalidad); err != nil {
		return nil, err
	}
	cfg := c.soap.Config()
	datos := datosMasivaXML{
		XMLName:               xml.Name{Local: "SolicitudServicioRecepcionMasiva"},
		CodigoAmbiente:        cfg.CodigoAmbiente,
		CodigoDocumentoSector: p.CodigoDocumentoSector,
		CodigoEmision:         p.CodigoEmision,
		CodigoModalidad:       p.CodigoModalidad,
		CodigoPuntoVenta:      p.CodigoPuntoVenta,
		CodigoSistema:         cfg.CodigoSistema,
		CodigoSucursal:        p.CodigoSucursal,
		Cufd:                  p.Cufd,
		Cuis:                  p.Cuis,
		Nit:                   cfg.Nit,
		TipoFacturaDocumento:  p.TipoFacturaDocumento,
		Archivo:               base64.StdEncoding.EncodeToString(p.Archivo),
		FechaEnvio:            p.FechaEnvio.Format("2006-01-02T15:04:05.000"),
		HashArchivo:           p.HashArchivo,
		CantidadFacturas:      r.CantidadFacturas,
	}
	request := struct {
		XMLName   xml.Name
		Solicitud datosMasivaXML
	}{XMLName: xml.Name{Local: "ns:recepcionMasivaFactura"}, Solicitud: datos}
	var resp struct {
		XMLName   xml.Name            `xml:"recepcionMasivaFacturaResponse"`
		Resultado *RespuestaRecepcion `xml:"RespuestaServicioFacturacion"`
	}
	if err := c.soap.Call(ctx, request, &resp); err != nil {
		return nil, err
	}
	if resp.Resultado == nil {
		return nil, fmt.Errorf("respuesta SIAT sin resultado de recepción masiva")
	}
	return resp.Resultado, nil
}

// ValidacionRecepcionMasivaFactura consulta el resultado de validación de un lote masivo.
func (c *Client) ValidacionRecepcionMasivaFactura(ctx context.Context, r SolicitudValidacionMasiva) (*RespuestaRecepcion, error) {
	if c == nil || c.soap == nil {
		return nil, fmt.Errorf("cliente SIAT no inicializado")
	}
	if r.CodigoDocumentoSector <= 0 || r.CodigoEmision != 3 || r.CodigoModalidad <= 0 || r.CodigoSucursal < 0 || (r.CodigoPuntoVenta != nil && *r.CodigoPuntoVenta < 0) || r.Cufd == "" || r.Cuis == "" || r.TipoFacturaDocumento <= 0 || r.CodigoRecepcion == "" {
		return nil, fmt.Errorf("solicitud de validación masiva incompleta o inválida")
	}
	if err := c.validateDestination(r.CodigoDocumentoSector, r.CodigoModalidad); err != nil {
		return nil, err
	}
	cfg := c.soap.Config()
	datos := datosValidacionPaqueteXML{
		XMLName:               xml.Name{Local: "SolicitudServicioValidacionRecepcionMasiva"},
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
		CodigoRecepcion:       r.CodigoRecepcion,
	}
	request := struct {
		XMLName   xml.Name
		Solicitud datosValidacionPaqueteXML
	}{XMLName: xml.Name{Local: "ns:validacionRecepcionMasivaFactura"}, Solicitud: datos}
	var resp struct {
		XMLName   xml.Name            `xml:"validacionRecepcionMasivaFacturaResponse"`
		Resultado *RespuestaRecepcion `xml:"RespuestaServicioFacturacion"`
	}
	if err := c.soap.Call(ctx, request, &resp); err != nil {
		return nil, err
	}
	if resp.Resultado == nil {
		return nil, fmt.Errorf("respuesta SIAT sin resultado de validación masiva")
	}
	return resp.Resultado, nil
}
