package facturacion

import (
	"context"
	"encoding/base64"
	"encoding/xml"
	"fmt"

	"github.com/ron86i/go-siat/v3/internal/siatpolicy"
)

// RecepcionPaqueteFactura envía un paquete de contingencia (TAR.GZ). Una recepción no implica validación.
func (c *Client) RecepcionPaqueteFactura(ctx context.Context, r SolicitudRecepcionPaquete) (*RespuestaRecepcion, error) {
	if c == nil || c.soap == nil {
		return nil, fmt.Errorf("cliente SIAT no inicializado")
	}
	p := r.SolicitudRecepcionFactura
	if p.CodigoDocumentoSector <= 0 || p.CodigoEmision != 2 || p.CodigoModalidad <= 0 || p.CodigoSucursal < 0 || (p.CodigoPuntoVenta != nil && *p.CodigoPuntoVenta < 0) || p.Cufd == "" || p.Cuis == "" || p.TipoFacturaDocumento <= 0 || len(p.Archivo) == 0 || p.FechaEnvio.IsZero() || p.HashArchivo == "" || r.CantidadFacturas <= 0 || r.CantidadFacturas > siatpolicy.MaxFacturasPaquete || r.CodigoEvento <= 0 {
		return nil, fmt.Errorf("solicitud de paquete incompleta o inválida")
	}
	if err := c.validateDestination(p.CodigoDocumentoSector, p.CodigoModalidad); err != nil {
		return nil, err
	}
	cfg := c.soap.Config()
	datos := datosPaqueteXML{
		XMLName:               xml.Name{Local: "SolicitudServicioRecepcionPaquete"},
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
		CodigoEvento:          r.CodigoEvento,
	}
	request := struct {
		XMLName   xml.Name
		Solicitud datosPaqueteXML
	}{XMLName: xml.Name{Local: "ns:recepcionPaqueteFactura"}, Solicitud: datos}
	var response struct {
		XMLName   xml.Name            `xml:"recepcionPaqueteFacturaResponse"`
		Resultado *RespuestaRecepcion `xml:"RespuestaServicioFacturacion"`
	}
	if err := c.soap.Call(ctx, request, &response); err != nil {
		return nil, err
	}
	if response.Resultado == nil {
		return nil, fmt.Errorf("respuesta SIAT sin resultado de recepción de paquete")
	}
	return response.Resultado, nil
}

// ValidacionRecepcionPaqueteFactura consulta el resultado de procesamiento de un paquete mediante su código de recepción.
func (c *Client) ValidacionRecepcionPaqueteFactura(ctx context.Context, r SolicitudValidacionPaquete) (*RespuestaRecepcion, error) {
	if c == nil || c.soap == nil {
		return nil, fmt.Errorf("cliente SIAT no inicializado")
	}
	if r.CodigoDocumentoSector <= 0 || r.CodigoEmision != 2 || r.CodigoModalidad <= 0 || r.CodigoSucursal < 0 || (r.CodigoPuntoVenta != nil && *r.CodigoPuntoVenta < 0) || r.Cufd == "" || r.Cuis == "" || r.TipoFacturaDocumento <= 0 || r.CodigoRecepcion == "" {
		return nil, fmt.Errorf("solicitud de validación de paquete incompleta o inválida")
	}
	if err := c.validateDestination(r.CodigoDocumentoSector, r.CodigoModalidad); err != nil {
		return nil, err
	}
	cfg := c.soap.Config()
	datos := datosValidacionPaqueteXML{
		XMLName:               xml.Name{Local: "SolicitudServicioValidacionRecepcionPaquete"},
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
	}{XMLName: xml.Name{Local: "ns:validacionRecepcionPaqueteFactura"}, Solicitud: datos}
	var response struct {
		XMLName   xml.Name            `xml:"validacionRecepcionPaqueteFacturaResponse"`
		Resultado *RespuestaRecepcion `xml:"RespuestaServicioFacturacion"`
	}
	if err := c.soap.Call(ctx, request, &response); err != nil {
		return nil, err
	}
	if response.Resultado == nil {
		return nil, fmt.Errorf("respuesta SIAT sin resultado de validación de paquete")
	}
	return response.Resultado, nil
}
