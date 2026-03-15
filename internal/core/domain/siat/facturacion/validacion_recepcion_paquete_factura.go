package facturacion

import "encoding/xml"

type ValidacionRecepcionPaqueteFactura struct {
	XMLName                                     xml.Name                     `xml:"ns:validacionRecepcionPaqueteFactura" json:"-"`
	SolicitudServicioValidacionRecepcionPaquete SolicitudValidacionRecepcion `xml:"SolicitudServicioValidacionRecepcionPaquete" json:"solicitudServicioValidacionRecepcionPaquete"`
}

type SolicitudValidacionRecepcion struct {
	SolicitudRecepcion
	CodigoRecepcion string `xml:"codigoRecepcion" json:"codigoRecepcion"`
}

type ValidacionRecepcionPaqueteFacturaResponse struct {
	XMLName                      xml.Name           `xml:"validacionRecepcionPaqueteFacturaResponse" json:"-"`
	RespuestaServicioFacturacion RespuestaRecepcion `xml:"RespuestaServicioFacturacion" json:"respuestaServicioFacturacion"`
}
