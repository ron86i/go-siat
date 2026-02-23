package compra_venta

import "encoding/xml"

type AnulacionFactura struct {
	XMLName                           xml.Name                          `xml:"ns:anulacionFactura"`
	SolicitudServicioAnulacionFactura SolicitudServicioAnulacionFactura `xml:"SolicitudServicioAnulacionFactura"`
}

type SolicitudServicioAnulacionFactura struct {
	CodigoAmbiente        int    `xml:"codigoAmbiente"`
	CodigoDocumentoSector int    `xml:"codigoDocumentoSector"`
	CodigoEmision         uint64 `xml:"codigoEmision"`
	CodigoModalidad       uint64 `xml:"codigoModalidad"`
	CodigoPuntoVenta      uint64 `xml:"codigoPuntoVenta"`
	CodigoSistema         string `xml:"codigoSistema"`
	CodigoSucursal        uint64 `xml:"codigoSucursal"`
	Cufd                  string `xml:"cufd"`
	Cuf                   string `xml:"cuf"`
	Cuis                  string `xml:"cuis"`
	Nit                   uint64 `xml:"nit"`
	TipoFacturaDocumento  uint64 `xml:"tipoFacturaDocumento"`
	CodigoMotivo          uint64 `xml:"codigoMotivo"`
}
