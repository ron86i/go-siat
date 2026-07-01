package documento_ajuste

import (
	"encoding/xml"

	"github.com/ron86i/go-siat/v2/internal/core/domain/siat/facturacion"
)

type AnulacionDocumentoAjuste struct {
	XMLName                                   xml.Name                       `xml:"ns:anulacionDocumentoAjuste" json:"-"`
	SolicitudServicioAnulacionDocumentoAjuste facturacion.SolicitudAnulacion `xml:"SolicitudServicioAnulacionDocumentoAjuste" json:"solicitudServicioAnulacionDocumentoAjuste"`
}

type AnulacionDocumentoAjusteResponse struct {
	XMLName                      xml.Name                       `xml:"anulacionDocumentoAjusteResponse" json:"-"`
	RespuestaServicioFacturacion facturacion.RespuestaRecepcion `xml:"RespuestaServicioFacturacion" json:"respuestaServicioFacturacion"`
}
