package compra_venta

import "encoding/xml"

type RecepcionFacturaResponse struct {
	XMLName                      xml.Name                     `xml:"recepcionFacturaResponse" json:"-"`
	XmlnsNs2                     string                       `xml:"xmlns:ns2,attr" json:"-"`
	RespuestaServicioFacturacion RespuestaServicioFacturacion `xml:"RespuestaServicioFacturacion,omitempty"`
}

type AnulacionFacturaResponse struct {
	XMLName                      xml.Name                     `xml:"anulacionFacturaResponse" json:"-"`
	XmlnsNs2                     string                       `xml:"xmlns:ns2,attr" json:"-"`
	RespuestaServicioFacturacion RespuestaServicioFacturacion `xml:"RespuestaServicioFacturacion,omitempty"`
}
