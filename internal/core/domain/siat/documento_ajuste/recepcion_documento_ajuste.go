package documento_ajuste

import (
	"encoding/xml"

	"github.com/ron86i/go-siat/internal/core/domain/siat/facturacion"
)

type RecepcionDocumentoAjuste struct {
	XMLName                   xml.Name                              `xml:"ns:recepcionDocumentoAjuste" json:"-"`
	SolicitudRecepcionFactura facturacion.SolicitudRecepcionFactura `xml:"SolicitudServicioRecepcionDocumentoAjuste" json:"solicitudServicioRecepcionDocumentoAjuste"`
}

type RecepcionDocumentoAjusteResponse struct {
	XMLName                   xml.Name                       `xml:"ns:recepcionDocumentoAjusteResponse" json:"-"`
	RespuestaRecepcionFactura facturacion.RespuestaRecepcion `xml:"RespuestaServicioFacturacion" json:"respuestaServicioFacturacion"`
}
