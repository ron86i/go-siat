package facturacion

import "encoding/xml"

type ReversionAnulacionFactura struct {
	XMLName                     xml.Name                    `xml:"ns:reversionAnulacionFactura" json:"-"`
	SolicitudReversionAnulacion SolicitudReversionAnulacion `xml:"SolicitudServicioReversionAnulacionFactura" json:"solicitudReversionAnulacion"`
}

type SolicitudReversionAnulacion struct {
	SolicitudRecepcion
	Cuf string `xml:"cuf" json:"cuf"`
}

type ReversionAnulacionFacturaResponse struct {
	XMLName                      xml.Name           `xml:"reversionAnulacionFacturaResponse" json:"-"`
	RespuestaServicioFacturacion RespuestaRecepcion `xml:"RespuestaServicioFacturacion" json:"respuestaServicioFacturacion"`
}
