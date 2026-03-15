package facturacion

import "encoding/xml"

type AnulacionFactura struct {
	XMLName            xml.Name           `xml:"ns:anulacionFactura" json:"-"`
	SolicitudAnulacion SolicitudAnulacion `xml:"SolicitudServicioAnulacionFactura" json:"solicitudservicioanulacionfactura"`
}

type SolicitudAnulacion struct {
	SolicitudRecepcion
	Cuf          string `xml:"cuf" json:"cuf"`
	CodigoMotivo int    `xml:"codigoMotivo" json:"codigomotivo"`
}

type AnulacionFacturaResponse struct {
	XMLName                      xml.Name           `xml:"anulacionFacturaResponse" json:"-"`
	RespuestaServicioFacturacion RespuestaRecepcion `xml:"RespuestaServicioFacturacion" json:"respuestadelserviciofacturacion"`
}
