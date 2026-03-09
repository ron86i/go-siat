package compra_venta

import "encoding/xml"

type RecepcionMasivaFactura struct {
	XMLName                          xml.Name                 `xml:"ns:recepcionMasivaFactura" json:"-"`
	SolicitudServicioRecepcionMasiva SolicitudRecepcionMasiva `xml:"SolicitudServicioRecepcionMasiva" json:"solicitudServicioRecepcionMasiva"`
}

type SolicitudRecepcionMasiva struct {
	SolicitudRecepcionFactura
	CantidadFacturas int `xml:"cantidadFacturas" json:"cantidadFacturas"`
}

type RecepcionMasivaFacturaResponse struct {
	XMLName                      xml.Name           `xml:"recepcionMasivaFacturaResponse" json:"-"`
	RespuestaServicioFacturacion RespuestaRecepcion `xml:"RespuestaServicioFacturacion" json:"respuestaServicioFacturacion"`
}
