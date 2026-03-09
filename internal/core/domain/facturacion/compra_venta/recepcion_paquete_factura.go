package compra_venta

import (
	"encoding/xml"
)

type RecepcionPaqueteFactura struct {
	XMLName                           xml.Name                  `xml:"ns:recepcionPaqueteFactura" json:"-"`
	SolicitudServicioRecepcionPaquete SolicitudRecepcionPaquete `xml:"SolicitudServicioRecepcionPaquete" json:"solicitudServicioRecepcionPaquete"`
}

type SolicitudRecepcionPaquete struct {
	SolicitudRecepcionFactura
	Cafc             string `xml:"cafc" json:"cafc"`
	CantidadFacturas int    `xml:"cantidadFacturas" json:"cantidadFacturas"`
	CodigoEvento     int64  `xml:"codigoEvento" json:"codigoEvento"`
}

type RecepcionPaqueteFacturaResponse struct {
	XMLName                      xml.Name           `xml:"recepcionPaqueteFacturaResponse" json:"-"`
	RespuestaServicioFacturacion RespuestaRecepcion `xml:"RespuestaServicioFacturacion" json:"respuestaServicioFacturacion"`
}
