package facturacion

import (
	"encoding/xml"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
)

type RecepcionPaqueteFactura struct {
	XMLName                           xml.Name                  `xml:"ns:recepcionPaqueteFactura" json:"-"`
	SolicitudServicioRecepcionPaquete SolicitudRecepcionPaquete `xml:"SolicitudServicioRecepcionPaquete" json:"solicitudServicioRecepcionPaquete"`
}

type SolicitudRecepcionPaquete struct {
	SolicitudRecepcionFactura
	Cafc             datatype.Nilable[string] `xml:"cafc" json:"cafc"`
	CantidadFacturas int                      `xml:"cantidadFacturas" json:"cantidadFacturas"`
	CodigoEvento     int64                    `xml:"codigoEvento" json:"codigoEvento"`
}

type RecepcionPaqueteFacturaResponse struct {
	XMLName                      xml.Name           `xml:"recepcionPaqueteFacturaResponse" json:"-"`
	RespuestaServicioFacturacion RespuestaRecepcion `xml:"RespuestaServicioFacturacion" json:"respuestaServicioFacturacion"`
}
