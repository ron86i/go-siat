package documento_ajuste

import (
	"encoding/xml"

	"github.com/ron86i/go-siat/internal/core/domain/siat/facturacion"
)

type VerificacionEstadoDocumentoAjuste struct {
	XMLName                             xml.Name                                `xml:"ns:verificacionEstadoDocumentoAjuste" json:"-"`
	SolicitudServicioVerificacionEstado facturacion.SolicitudVerificacionEstado `xml:"SolicitudServicioVerificacionEstadoDocumentoAjuste" json:"solicitudServicioVerificacionEstadoDocumentoAjuste"`
}

type VerificacionEstadoDocumentoAjusteResponse struct {
	XMLName                      xml.Name                       `xml:"ns:verificacionEstadoDocumentoAjusteResponse" json:"-"`
	RespuestaServicioFacturacion facturacion.RespuestaRecepcion `xml:"RespuestaServicioFacturacion" json:"respuestaServicioFacturacion"`
}
