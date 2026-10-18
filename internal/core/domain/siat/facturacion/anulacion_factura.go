package facturacion

import "encoding/xml"

type AnulacionFactura struct {
	XMLName            xml.Name           `xml:"ns:anulacionFactura" json:"-"`
	SolicitudAnulacion SolicitudAnulacion `xml:"SolicitudServicioAnulacionFactura" json:"solicitudservicioanulacionfactura"`
}

type SolicitudAnulacion struct {
	SolicitudRecepcion
	CodigoMotivo int    `xml:"codigoMotivo" json:"codigomotivo"`
	Cuf          string `xml:"cuf" json:"cuf"`
}

type AnulacionFacturaResponse struct {
	XMLName                      xml.Name           `xml:"anulacionFacturaResponse" json:"-"`
	RespuestaServicioFacturacion RespuestaRecepcion `xml:"RespuestaServicioFacturacion" json:"respuestadelserviciofacturacion"`
}
