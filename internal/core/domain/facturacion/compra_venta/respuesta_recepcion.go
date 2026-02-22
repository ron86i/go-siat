package compra_venta

import "encoding/xml"

type RespuestaServicioFacturacion struct {
	CodigoDescripcion string    `xml:"codigoDescripcion,omitempty" json:"codigoDescripcion"`
	CodigoEstado      int       `xml:"codigoEstado,omitempty" json:"codigoEstado"`
	CodigoRecepcion   string    `xml:"codigoRecepcion,omitempty" json:"codigoRecepcion"`
	MensajesList      []Mensaje `xml:"mensajesList,omitempty" json:"mensajesList"`
	Transaccion       bool      `xml:"transaccion,omitempty" json:"transaccion"`
}

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
