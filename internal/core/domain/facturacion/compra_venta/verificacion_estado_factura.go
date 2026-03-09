package compra_venta

import "encoding/xml"

type VerificacionEstadoFactura struct {
	XMLName                                    xml.Name                    `xml:"ns:verificacionEstadoFactura" json:"-"`
	SolicitudServicioVerificacionEstadoFactura SolicitudVerificacionEstado `xml:"SolicitudServicioVerificacionEstado" json:"solicitudServicioVerificacionEstado"`
}

type SolicitudVerificacionEstado struct {
	SolicitudRecepcion
	Cuf string `xml:"cuf" json:"cuf"`
}

type VerificacionEstadoFacturaResponse struct {
	XMLName                      xml.Name           `xml:"verificacionEstadoFacturaResponse" json:"-"`
	RespuestaServicioFacturacion RespuestaRecepcion `xml:"RespuestaServicioFacturacion" json:"respuestaServicioFacturacion"`
}
