package compra_venta

import "encoding/xml"

type ValidacionRecepcionMasivaFactura struct {
	XMLName                                           xml.Name                           `xml:"ns:validacionRecepcionMasivaFactura" json:"-"`
	SolicitudServicioValidacionRecepcionMasivaFactura SolicitudValidacionRecepcionMasiva `xml:"SolicitudServicioValidacionRecepcionMasiva" json:"solicitudServicioValidacionRecepcionMasiva"`
}

type SolicitudValidacionRecepcionMasiva struct {
	SolicitudRecepcionFactura
	CodigoRecepcion string `xml:"codigoRecepcion" json:"codigoRecepcion"`
}

type ValidacionRecepcionMasivaFacturaResponse struct {
	XMLName                      xml.Name           `xml:"validacionRecepcionMasivaFacturaResponse" json:"-"`
	RespuestaServicioFacturacion RespuestaRecepcion `xml:"RespuestaServicioFacturacion" json:"respuestaServicioFacturacion"`
}
