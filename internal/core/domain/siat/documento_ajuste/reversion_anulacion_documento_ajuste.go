package documento_ajuste

import (
	"encoding/xml"

	"github.com/ron86i/go-siat/internal/core/domain/siat/facturacion"
)

type ReversionAnulacionDocumentoAjuste struct {
	XMLName                                            xml.Name                                `xml:"ns:reversionAnulacionDocumentoAjuste" json:"-"`
	SolicitudServicioReversionAnulacionDocumentoAjuste facturacion.SolicitudReversionAnulacion `xml:"SolicitudServicioReversionAnulacionDocumentoAjuste" json:"solicitudServicioReversionAnulacionDocumentoAjuste"`
}

type ReversionAnulacionDocumentoAjusteResponse struct {
	XMLName                      xml.Name                       `xml:"ns:reversionAnulacionDocumentoAjusteResponse" json:"-"`
	RespuestaServicioFacturacion facturacion.RespuestaRecepcion `xml:"RespuestaServicioFacturacion" json:"respuestaServicioFacturacion"`
}
